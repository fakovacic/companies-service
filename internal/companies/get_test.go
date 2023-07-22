package companies_test

import (
	"context"
	"testing"

	"github.com/fakovacic/companies-service/internal/companies"
	"github.com/fakovacic/companies-service/internal/companies/errors"
	"github.com/fakovacic/companies-service/internal/companies/mocks"
	"github.com/fakovacic/companies-service/internal/tests"
	"github.com/matryer/is"
)

func TestGet(t *testing.T) {
	cases := []struct {
		it string

		id string

		// Store
		companyGetInput  string
		companyGetResult *companies.Company
		companyGetError  error

		expectedError  string
		expectedResult *companies.Company
	}{
		{
			it: "it returns company",

			id: "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",

			companyGetInput: "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
			companyGetResult: &companies.Company{
				ID: "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
			},

			expectedResult: &companies.Company{
				ID: "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
			},
		},
		{
			it: "it return error on store Get",

			id: "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",

			companyGetInput: "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
			companyGetError: errors.New("mock-error"),

			expectedError: "get company: mock-error",
		},
		{
			it: "it return error, id invalid format",

			id:            "mock-id",
			expectedError: "id invalid format",
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			store := &mocks.StoreMock{
				GetFunc: func(ctx context.Context, id string) (*companies.Company, error) {
					checkIs.Equal(id, tc.companyGetInput)

					return tc.companyGetResult, tc.companyGetError
				},
			}

			service := companies.New(
				companies.NewConfig(""),
				store,
				tests.GenTime,
				tests.GenUUID,
			)

			res, err := service.Get(context.Background(), tc.id)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}
			checkIs.Equal(res, tc.expectedResult)
		})
	}
}
