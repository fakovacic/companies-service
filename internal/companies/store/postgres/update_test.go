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

func TestUpdate(t *testing.T) {
	type row struct {
		id              string
		name            string
		description     string
		employeesAmount uint32
		registered      bool
		companyType     string
		updatedAt       time.Time
	}

	cases := []struct {
		it string

		id  string
		req *companies.Company

		sqlInput *row
		sqlError error

		expectedError string
	}{
		{
			it: "it update a company",

			id: "mock-id",
			req: &companies.Company{
				UpdatedAt: tests.GenTime(),
			},

			sqlInput: &row{
				id:        "mock-id",
				updatedAt: tests.GenTime(),
			},
		},
		{
			it: "it return sql error",

			id: "mock-id",
			req: &companies.Company{
				UpdatedAt: tests.GenTime(),
			},

			sqlInput: &row{
				id:        "mock-id",
				updatedAt: tests.GenTime(),
			},
			sqlError: sql.ErrConnDone,

			expectedError: "update company: sql: connection is already closed",
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

			mock.ExpectExec(regexp.QuoteMeta(`UPDATE companies SET name = $1, description = $2, employees_amount = $3, registered = $4, type = $5, updated_at = $6 WHERE id = $7`)).
				WithArgs(
					tc.sqlInput.name,
					tc.sqlInput.description,
					tc.sqlInput.employeesAmount,
					tc.sqlInput.registered,
					tc.sqlInput.companyType,
					tc.sqlInput.updatedAt,
					tc.sqlInput.id,
				).
				WillReturnResult(sqlmock.NewResult(1, 1)).
				WillReturnError(tc.sqlError)

			service := postgres.New(db)

			err = service.Update(context.Background(), tc.id, tc.req)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}
		})
	}
}
