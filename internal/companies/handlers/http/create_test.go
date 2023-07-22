package http_test

import (
	"bytes"
	"context"
	"errors"
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

func TestCreate(t *testing.T) {
	cases := []struct {
		it string

		requestBody string

		createInput    *companies.Company
		createResponse *companies.Company
		createError    error

		expectedError  string
		expectedResult string
		expectedStatus int
	}{
		{
			it: "it create a company",

			requestBody: `{"name":"mock-name","description":"mock-description","employeesAmount":1,"registered":true,"type":"corporations"}`,

			createInput: &companies.Company{
				Name:            "mock-name",
				Description:     "mock-description",
				EmployeesAmount: 1,
				Registered:      true,
				Type:            "corporations",
			},
			createResponse: &companies.Company{
				ID:              "df65c534-0e8c-4f06-9260-24d07ecd6127",
				Name:            "mock-name",
				Description:     "mock-description",
				EmployeesAmount: 1,
				Registered:      true,
				Type:            "mock-type",
				CreatedAt:       tests.GenTime(),
				UpdatedAt:       tests.GenTime(),
			},

			expectedResult: `{"id":"df65c534-0e8c-4f06-9260-24d07ecd6127","name":"mock-name","description":"mock-description","employeesAmount":1,"registered":true,"type":"mock-type","createdAt":"2020-01-02T03:04:05Z","updatedAt":"2020-01-02T03:04:05Z"}`,
			expectedStatus: http.StatusCreated,
		},
		{
			it: "it return error on create company",

			requestBody: `{"name":"mock-name","description":"mock-description","employeesAmount":1,"registered":true,"type":"corporations"}`,

			createInput: &companies.Company{
				Name:            "mock-name",
				Description:     "mock-description",
				EmployeesAmount: 1,
				Registered:      true,
				Type:            "corporations",
			},
			createError: errors.New("mock-error"),

			expectedResult: `{"message":"mock-error"}`,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			service := &mocks.ServiceMock{
				CreateFunc: func(ctx context.Context, in *companies.Company) (*companies.Company, error) {
					checkIs.Equal(in, tc.createInput)

					return tc.createResponse, tc.createError
				},
			}

			req, err := http.NewRequestWithContext(
				context.Background(),
				http.MethodPost,
				"/",
				bytes.NewReader([]byte(tc.requestBody)),
			)
			if err != nil {
				t.Fatal(err)
			}

			h := handlers.New(companies.NewConfig("", ""), service)

			app := fiber.New()
			app.Post("/", h.Create())

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
