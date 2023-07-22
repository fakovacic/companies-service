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

func TestDelete(t *testing.T) {
	cases := []struct {
		it string

		id string

		// Store
		companyDeleteInput string
		companyDeleteError error

		expectedError string
	}{
		{
			it: "it delete company",

			id:                 "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
			companyDeleteInput: "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
		},
		{
			it: "it return error on store Delete",

			id:                 "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
			companyDeleteInput: "bbe20b5d-929e-4280-9a0a-7ee7ab3c7c1d",
			companyDeleteError: errors.New("mock-error"),
			expectedError:      "delete company: mock-error",
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
				DeleteFunc: func(ctx context.Context, id string) error {
					checkIs.Equal(id, tc.companyDeleteInput)

					return tc.companyDeleteError
				},
			}

			service := companies.New(
				companies.NewConfig(""),
				store,
				tests.GenTime,
				tests.GenUUID,
			)

			err := service.Delete(context.Background(), tc.id)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}
		})
	}
}
