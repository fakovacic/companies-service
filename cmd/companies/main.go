package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fakovacic/companies-service/cmd/companies/config"
	"github.com/fakovacic/companies-service/cmd/companies/config/integrations"
	"github.com/fakovacic/companies-service/internal/companies"
	"github.com/fakovacic/companies-service/internal/companies/handlers/http/middleware"
	svcMiddleware "github.com/fakovacic/companies-service/internal/companies/middleware"
	"github.com/fakovacic/companies-service/internal/health"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const errorChan int = 10

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbStore, err := config.NewStore(c)
	if err != nil {
		log.Fatal(err)
	}

	notifier := integrations.NewNotifier(c)

	service := companies.New(c, dbStore, time.Now, uuid.New)
	service = svcMiddleware.NewNotificationMiddleware(service, notifier)
	service = svcMiddleware.NewLoggingMiddleware(service, c)

	h := config.NewHandlers(c, service)

	app := fiber.New()
	app.Use(middleware.Logger(c))
	app.Use(middleware.ReqID())

	app.Post("/", h.Create())
	app.Get("/:id", h.Get())
	app.Patch("/:id", h.Update())
	app.Delete("/:id", h.Delete())

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
		case s := <-signalChan:
			c.Log.Info().Msgf("Captured %v. Exiting...", s)
			health.SetHealthStatus(http.StatusServiceUnavailable)

			err = app.Shutdown()
			if err != nil {
				c.Log.Fatal().Msg(err.Error())
			}

			os.Exit(0)
		}
	}
}
