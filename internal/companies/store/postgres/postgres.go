package postgres

import (
	"database/sql"

	"github.com/fakovacic/companies-service/internal/companies"
)

func New(db *sql.DB) companies.Store {
	s := &store{
		db: db,
	}

	return s
}

type store struct {
	db *sql.DB
}

func (s *store) DB() *sql.DB {
	return s.db
}
