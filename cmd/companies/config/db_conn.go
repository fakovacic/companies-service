package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fakovacic/companies-service/internal/companies"
	"github.com/fakovacic/companies-service/internal/companies/errors"
	migrate "github.com/rubenv/sql-migrate"
)

const (
	retryDBTimeout time.Duration = 5 * time.Second
	postgresPort   int           = 5432
)

func NewDBConn(c *companies.Config) (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		return nil, errors.New("database host is empty")
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		return nil, errors.New("database user is empty")
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		return nil, errors.New("database password is empty")
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		return nil, errors.New("database name is empty")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, postgresPort, user, password, dbname)

	db, err := retryDBConn(psqlInfo)
	if err != nil {
		return nil, errors.Wrap(err, "db connection error")
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "/sql",
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return nil, errors.Wrap(err, "migration error")
	}

	c.Log.Info().Msgf("applied %d migrations", n)

	return db, nil
}

func retryDBConn(psqlInfo string) (*sql.DB, error) {
	for i := 0; i <= 3; i++ {
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Printf("postgres connection error: %v", err)
			time.Sleep(retryDBTimeout)

			continue
		}

		err = db.Ping()
		if err == nil {
			return db, nil
		}

		time.Sleep(retryDBTimeout)
	}

	return nil, errors.New("database connection retry exceded")
}
