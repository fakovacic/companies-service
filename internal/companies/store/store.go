package store

import (
	"database/sql"

	"github.com/fakovacic/companies-service/internal/companies"
	"github.com/fakovacic/companies-service/internal/companies/store/postgres"
)

func NewStore(db *sql.DB) companies.Store {
	return postgres.New(db)
}
