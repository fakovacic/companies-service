package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/fakovacic/companies-service/internal/companies"
	"github.com/fakovacic/companies-service/internal/companies/errors"
	"github.com/fakovacic/companies-service/internal/companies/store"
	migrate "github.com/rubenv/sql-migrate"
)

const (
	retryTimeout time.Duration = 5 * time.Second
	postgresPort int           = 5432
)

func NewStore(c *companies.Config) (companies.Store, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" {
		return nil, errors.New("database host is empty")
	}

	if user == "" {
		return nil, errors.New("database user is empty")
	}

	if password == "" {
		return nil, errors.New("database password is empty")
	}

	if dbname == "" {
		return nil, errors.New("database name is empty")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, postgresPort, user, password, dbname)

	db, err := retryConn(c, psqlInfo)
	if err != nil {
		return nil, errors.Wrap(err, "store connection error")
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "/sql",
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return nil, errors.Wrap(err, "migration error")
	}

	c.Log.Info().Msg(fmt.Sprintf("applied %d migrations", n))

	return store.NewStore(db), nil
}

func retryConn(c *companies.Config, psqlInfo string) (*sql.DB, error) {
	for i := 0; i <= 3; i++ {
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			c.Log.Warn().Msg(fmt.Sprintf("postgres connection error: %v", err))
			time.Sleep(retryTimeout)

			continue
		}

		err = db.Ping()
		if err == nil {
			return db, nil
		}

		time.Sleep(retryTimeout)
	}

	return nil, errors.New("database connection retry exceded")
}
