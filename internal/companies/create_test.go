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

func TestCreate(t *testing.T) {
	cases := []struct {
		it  string
		req *companies.Company

		// Store
		companyCreateInput *companies.Company
		companyCreateError error

		expectedError  string
		expectedResult *companies.Company
	}{
		{
			it: "it create company",

			req: &companies.Company{
				Name:            "mock-name",
				EmployeesAmount: 10,
			},

			companyCreateInput: &companies.Company{
				ID:              tests.GenUUID().String(),
				Name:            "mock-name",
				EmployeesAmount: 10,
				CreatedAt:       tests.GenTime(),
				UpdatedAt:       tests.GenTime(),
			},

			expectedResult: &companies.Company{
				ID:              tests.GenUUID().String(),
				Name:            "mock-name",
				EmployeesAmount: 10,
				CreatedAt:       tests.GenTime(),
				UpdatedAt:       tests.GenTime(),
			},
		},
		{
			it: "it return error on store Create",

			req: &companies.Company{
				Name:            "mock-name",
				EmployeesAmount: 10,
			},

			companyCreateInput: &companies.Company{
				ID:              tests.GenUUID().String(),
				Name:            "mock-name",
				EmployeesAmount: 10,
				CreatedAt:       tests.GenTime(),
				UpdatedAt:       tests.GenTime(),
			},
			companyCreateError: errors.New("mock-error"),

			expectedError: "create company: mock-error",
		},
		{
			it: "it return error on validation, employees amount is required",

			req: &companies.Company{
				Name: "mock-name",
			},

			expectedError: "validation: employees amount is required",
		},
		{
			it: "it return error on validation, name is too long",

			req: &companies.Company{
				Name: "mock-very-long-name",
			},

			expectedError: "validation: name is too long",
		},
		{
			it: "it return error on validation, name empty",

			req: &companies.Company{},

			expectedError: "validation: name is required",
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			store := &mocks.StoreMock{
				CreateFunc: func(ctx context.Context, model *companies.Company) error {
					checkIs.Equal(model, tc.companyCreateInput)

					return tc.companyCreateError
				},
			}

			service := companies.New(
				companies.NewConfig("", ""),
				store,
				tests.GenTime,
				tests.GenUUID,
			)

			res, err := service.Create(context.Background(), tc.req)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}
			checkIs.Equal(res, tc.expectedResult)
		})
	}
}
