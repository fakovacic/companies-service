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
	"github.com/fakovacic/companies-service/internal/tests"
	"github.com/gofiber/fiber/v2"
	"github.com/matryer/is"
)

func TestGet(t *testing.T) {
	cases := []struct {
		it string

		id string

		getInput    string
		getResponse *companies.Company
		getError    error

		expectedError  string
		expectedResult string
		expectedStatus int
	}{
		{
			it: "it return a company",

			id: `mock-id`,

			getInput: "mock-id",
			getResponse: &companies.Company{
				ID:        "df65c534-0e8c-4f06-9260-24d07ecd6127",
				CreatedAt: tests.GenTime(),
				UpdatedAt: tests.GenTime(),
			},

			expectedResult: `{"id":"df65c534-0e8c-4f06-9260-24d07ecd6127","name":"","description":"","employeesAmount":0,"registered":false,"type":"","createdAt":"2020-01-02T03:04:05Z","updatedAt":"2020-01-02T03:04:05Z"}`,
			expectedStatus: http.StatusOK,
		},
		{
			it: "it return error on get company",

			id: `mock-id`,

			getInput: "mock-id",
			getError: errors.New("mock-error"),

			expectedResult: `{"message":"mock-error"}`,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			service := &mocks.ServiceMock{
				GetFunc: func(ctx context.Context, id string) (*companies.Company, error) {
					checkIs.Equal(id, tc.getInput)

					return tc.getResponse, tc.getError
				},
			}

			req, err := http.NewRequestWithContext(
				context.Background(),
				http.MethodGet,
				fmt.Sprintf("/%s", tc.id),
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			h := handlers.New(companies.NewConfig("", ""), service)

			app := fiber.New()
			app.Get("/:id", h.Get())

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
