package http_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/fakovacic/companies-service/internal/companies"
	handlers "github.com/fakovacic/companies-service/internal/companies/handlers/http"
	"github.com/fakovacic/companies-service/internal/companies/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/matryer/is"
)

func TestDelete(t *testing.T) {
	cases := []struct {
		it string

		id string

		deleteInput string
		deleteError error

		expectedError  string
		expectedResult string
		expectedStatus int
	}{
		{
			it: "it delete comapny",

			id: `mock-id`,

			deleteInput: "mock-id",

			expectedResult: ``,
			expectedStatus: http.StatusNoContent,
		},
		{
			it: "it return error on delete company",

			id: `mock-id`,

			deleteInput: "mock-id",
			deleteError: errors.New("mock-error"),

			expectedResult: `{"message":"mock-error"}`,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			service := &mocks.ServiceMock{
				DeleteFunc: func(ctx context.Context, id string) error {
					checkIs.Equal(id, tc.deleteInput)

					return tc.deleteError
				},
			}

			req, err := http.NewRequestWithContext(
				context.Background(),
				http.MethodDelete,
				fmt.Sprintf("/%s", tc.id),
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			h := handlers.New(companies.NewConfig("", ""), service)

			app := fiber.New()
			app.Delete("/:id", h.Delete())

			response, err := app.Test(req)
			if err != nil {
				t.Fatal(err)
			}

			body, err := io.ReadAll(response.Body)
			if err != nil {
				t.Fatal(err)
			}

			checkIs.Equal(string(body), tc.expectedResult)
			checkIs.Equal(response.StatusCode, tc.expectedStatus)
		})
	}
}
