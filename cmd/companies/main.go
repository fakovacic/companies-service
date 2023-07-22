package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fakovacic/companies-service/cmd/companies/config"
	"github.com/fakovacic/companies-service/internal/companies"
	handlers "github.com/fakovacic/companies-service/internal/companies/handlers/http"
	"github.com/fakovacic/companies-service/internal/companies/handlers/http/middleware"
	"github.com/fakovacic/companies-service/internal/companies/integrations/notifier"
	svcMiddleware "github.com/fakovacic/companies-service/internal/companies/middleware"
	"github.com/fakovacic/companies-service/internal/companies/store"
	"github.com/fakovacic/companies-service/internal/health"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const errorChan int = 10

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal("config error:", err)
	}

	dbConn, err := config.NewDBConn(c)
	if err != nil {
		log.Fatal("db conn error:", err)
	}

	dbStore := store.NewStore(dbConn)

	kafkaConn, err := config.NewKafkaConn()
	if err != nil {
		log.Fatal("kafka conn error:", err)
	}

	defer func() {
		er := dbConn.Close()
		if er != nil {
			log.Fatal("db close error:", er)
		}
	}()

	defer func() {
		er := kafkaConn.Close()
		if er != nil {
			log.Fatal("kafka close error:", er)
		}
	}()

	notifier := notifier.New(kafkaConn)

	service := companies.New(c, dbStore, time.Now, uuid.New)
	service = svcMiddleware.NewNotificationMiddleware(service, c, notifier)
	service = svcMiddleware.NewLoggingMiddleware(service, c)

	h := handlers.New(c, service)

	app := fiber.New()
	app.Use(middleware.Logger(c))
	app.Use(middleware.ReqID())

	app.Post("/login", h.Login())
	app.Get("/companies/:id", h.Get())

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(c.JWTSigningKey),
		},
	}))

	app.Post("/companies", h.Create())
	app.Patch("/companies/:id", h.Update())
	app.Delete("/companies/:id", h.Delete())

	var (
		httpAddr   = "0.0.0.0:8080"
		healthAddr = "0.0.0.0:8081"
	)

	// health router
	healthServer := health.StartServer()

	errChan := make(chan error, errorChan)

	go func() {
		c.Log.Info().Msgf("Health service listening on %s", healthAddr)
		errChan <- healthServer.Listen(healthAddr)
	}()
	go func() {
		c.Log.Info().Msgf("HTTP service listening on %s", httpAddr)
		errChan <- app.Listen(httpAddr)
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case e := <-errChan:
			if e != nil {
				c.Log.Fatal().Msg(e.Error())
			}

			return
		case s := <-signalChan:
			c.Log.Info().Msgf("Captured %v. Exiting...", s)
			health.SetHealthStatus(http.StatusServiceUnavailable)

			err = app.Shutdown()
			if err != nil {
				c.Log.Fatal().Msg(err.Error())
			}

			return
		}
	}
}
