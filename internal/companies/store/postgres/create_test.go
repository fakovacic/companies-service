package postgres_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fakovacic/companies-service/internal/companies"
	"github.com/fakovacic/companies-service/internal/companies/store/postgres"
	"github.com/fakovacic/companies-service/internal/tests"
	"github.com/matryer/is"
)

func TestCreate(t *testing.T) {
	type row struct {
		id              string
		name            string
		description     string
		employeesAmount uint32
		registered      bool
		companyType     string
		createdAt       time.Time
		updatedAt       time.Time
	}

	cases := []struct {
		it string

		req *companies.Company

		sqlInput *row
		sqlError error

		expectedError string
	}{
		{
			it: "it create a company",
			req: &companies.Company{
				ID:        "mock-id",
				CreatedAt: tests.GenTime(),
				UpdatedAt: tests.GenTime(),
			},

			sqlInput: &row{
				id:        "mock-id",
				createdAt: tests.GenTime(),
				updatedAt: tests.GenTime(),
			},
		},
		{
			it: "it return sql error",
			req: &companies.Company{
				ID:        "mock-id",
				CreatedAt: tests.GenTime(),
				UpdatedAt: tests.GenTime(),
			},

			sqlInput: &row{
				id:        "mock-id",
				createdAt: tests.GenTime(),
				updatedAt: tests.GenTime(),
			},
			sqlError: sql.ErrConnDone,

			expectedError: "create company: sql: connection is already closed",
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO companies (id, name, description, employees_amount, registered, type, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`)).
				WithArgs(
					tc.sqlInput.id,
					tc.sqlInput.name,
					tc.sqlInput.description,
					tc.sqlInput.employeesAmount,
					tc.sqlInput.registered,
					tc.sqlInput.companyType,
					tc.sqlInput.createdAt,
					tc.sqlInput.updatedAt,
				).
				WillReturnResult(sqlmock.NewResult(1, 1)).
				WillReturnError(tc.sqlError)

			service := postgres.New(db)

			err = service.Create(context.Background(), tc.req)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}
		})
	}
}
