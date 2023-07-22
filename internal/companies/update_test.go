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

func TestUpdate(t *testing.T) {
	cases := []struct {
		it string

		id           string
		model        *companies.Company
		updateFields []string

		// Store
		companyGetInput  string
		companyGetResult *companies.Company
		companyGetError  error

		companyUpdateInputID    string
		companyUpdateInputModel *companies.Company
		companyUpdateError      error

		expectedError  string
		expectedResult *companies.Company
	}{
		{
			it: "it update and return company",

			id: "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
			model: &companies.Company{
				Name: "mock-name",
			},
			updateFields: []string{
				"name",
			},

			companyGetInput: "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
			companyGetResult: &companies.Company{
				ID:              "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
				EmployeesAmount: 10,
			},

			companyUpdateInputID: "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
			companyUpdateInputModel: &companies.Company{
				ID:              "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
				Name:            "mock-name",
				EmployeesAmount: 10,
				UpdatedAt:       tests.GenTime(),
			},

			expectedResult: &companies.Company{
				ID:              "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
				Name:            "mock-name",
				EmployeesAmount: 10,
				UpdatedAt:       tests.GenTime(),
			},
		},
		{
			it: "it returns error on store Update",

			id: "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
			model: &companies.Company{
				Name: "mock-name",
			},
			updateFields: []string{
				"name",
			},

			companyGetInput: "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
			companyGetResult: &companies.Company{
				ID:              "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
				EmployeesAmount: 10,
			},

			companyUpdateInputID: "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
			companyUpdateInputModel: &companies.Company{
				ID:              "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
				Name:            "mock-name",
				EmployeesAmount: 10,
				UpdatedAt:       tests.GenTime(),
			},

			companyUpdateError: errors.New("mock-error"),
			expectedError:      "update company: mock-error",
		},
		{
			it: "it returns error on model validation",

			id: "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
			model: &companies.Company{
				Name: "mock-very-long-name",
			},
			updateFields: []string{
				"name",
			},

			companyGetInput: "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
			companyGetResult: &companies.Company{
				ID:              "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
				EmployeesAmount: 10,
			},

			expectedError: "validation: name is too long",
		},
		{
			it: "it returns error on store Get",

			id: "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
			model: &companies.Company{
				Name: "mock-name",
			},
			updateFields: []string{
				"name",
			},

			companyGetInput: "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
			companyGetError: errors.New("mock-error"),

			expectedError: "get company: mock-error",
		},
		{
			it: "it returns error on check fields",

			id: "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
			updateFields: []string{
				"id",
			},

			expectedError: "field 'id' cannot be updated",
		},
		{
			it: "it returns error on field not found",

			id: "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
			updateFields: []string{
				"mock",
			},

			expectedError: "field 'mock' not exist",
		},
		{
			it:            "it returns error on empty fields",
			id:            "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
			expectedError: "update fields empty",
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
				UpdateFunc: func(ctx context.Context, id string, model *companies.Company) error {
					checkIs.Equal(id, tc.companyUpdateInputID)
					checkIs.Equal(model, tc.companyUpdateInputModel)

					return tc.companyUpdateError
				},
			}

			service := companies.New(
				companies.NewConfig("", ""),
				store,
				tests.GenTime,
				nil,
			)

			res, err := service.Update(context.Background(), tc.id, tc.model, tc.updateFields)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}
			checkIs.Equal(res, tc.expectedResult)
		})
	}
}
