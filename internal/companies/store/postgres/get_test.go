package postgres_test

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fakovacic/companies-service/internal/companies"
	"github.com/fakovacic/companies-service/internal/companies/errors"
	"github.com/fakovacic/companies-service/internal/companies/store/postgres"
	"github.com/matryer/is"
)

func TestGet(t *testing.T) {
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
		it             string
		id             string
		r              *row
		sqlError       error
		rowError       error
		expectedError  string
		expectedResult *companies.Company
	}{
		{
			it: "it returns company",
			id: "mock-id",
			r: &row{
				id:              "mock-id",
				name:            "mock-name",
				description:     "mock-description",
				employeesAmount: 10,
				registered:      true,
				companyType:     "corporations",
			},
			expectedResult: &companies.Company{
				ID:              "mock-id",
				Name:            "mock-name",
				Description:     "mock-description",
				EmployeesAmount: 10,
				Registered:      true,
				Type:            companies.TypeCorporations,
			},
		},
		{
			it:            "it returns error on row error",
			rowError:      errors.New("expected 6 destination arguments in Scan, not 8"),
			expectedError: "get company: sql: expected 6 destination arguments in Scan, not 8",
		},
		{
			it:            "it returns error db error",
			sqlError:      errors.New("mock-error"),
			expectedError: "get company: mock-error",
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

			query := mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, description, employees_amount, registered, type, created_at, updated_at FROM companies WHERE id=$1`))

			switch {
			case tc.r != nil:
				mockRow := sqlmock.NewRows([]string{
					"id",
					"name",
					"description",
					"employees_amount",
					"registered",
					"type",
					"created_at",
					"updated_at",
				}).AddRow(
					tc.r.id,
					tc.r.name,
					tc.r.description,
					tc.r.employeesAmount,
					tc.r.registered,
					tc.r.companyType,
					tc.r.createdAt,
					tc.r.updatedAt,
				)
				query.WillReturnRows(mockRow)
			case tc.rowError != nil:
				mockRow := sqlmock.NewRows([]string{
					"id",
					"name",
					"description",
					"employees_amount",
					"registered",
					"type",
				}).AddRow("", "", "", 0, false, "").RowError(1, tc.rowError)
				query.WillReturnRows(mockRow)
			default:
				query.WillReturnRows(sqlmock.NewRows([]string{
					"id",
					"name",
					"description",
					"employees_amount",
					"registered",
					"type",
					"created_at",
					"updated_at",
				}).AddRow("", "", "", 0, false, "", time.Now(), time.Now()))
				query.WillReturnError(tc.sqlError)
			}

			service := postgres.New(db)

			res, err := service.Get(context.Background(), tc.id)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}
			checkIs.Equal(res, tc.expectedResult)
		})
	}
}
