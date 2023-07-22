package http_test

import (
	"bytes"
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

func TestUpdate(t *testing.T) {
	cases := []struct {
		it string

		id          string
		requestBody string

		updateInputID     string
		updateInputData   *companies.Company
		updateInputFields []string
		updateResponse    *companies.Company
		updateError       error

		expectedError  string
		expectedResult string
		expectedStatus int
	}{
		{
			it: "it update a company",

			id:          `mock-id`,
			requestBody: `{"data":{"name":"mock-name","description":"mock-description","employeesAmount":1,"registered":true,"type":"corporations"},"fields":["name","description","employeesAmount","registered","type"]}`,

			updateInputID: "mock-id",
			updateInputData: &companies.Company{
				Name:            "mock-name",
				Description:     "mock-description",
				EmployeesAmount: 1,
				Registered:      true,
				Type:            "corporations",
			},
			updateInputFields: []string{"name", "description", "employeesAmount", "registered", "type"},
			updateResponse: &companies.Company{
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
			expectedStatus: http.StatusOK,
		},
		{
			it: "it return error on update company",

			id:          `mock-id`,
			requestBody: `{"data":{"name":"mock-name","description":"mock-description","employeesAmount":1,"registered":true,"type":"corporations"},"fields":["name","description","employeesAmount","registered","type"]}`,

			updateInputID: "mock-id",
			updateInputData: &companies.Company{
				Name:            "mock-name",
				Description:     "mock-description",
				EmployeesAmount: 1,
				Registered:      true,
				Type:            "corporations",
			},
			updateInputFields: []string{"name", "description", "employeesAmount", "registered", "type"},
			updateError:       errors.New("mock-error"),

			expectedResult: `{"message":"mock-error"}`,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			service := &mocks.ServiceMock{
				UpdateFunc: func(ctx context.Context, id string, in *companies.Company, fields []string) (*companies.Company, error) {
					checkIs.Equal(id, tc.updateInputID)
					checkIs.Equal(in, tc.updateInputData)
					checkIs.Equal(fields, tc.updateInputFields)

					return tc.updateResponse, tc.updateError
				},
			}

			req, err := http.NewRequestWithContext(
				context.Background(),
				http.MethodPatch,
				fmt.Sprintf("/%s", tc.id),
				bytes.NewReader([]byte(tc.requestBody)),
			)
			if err != nil {
				t.Fatal(err)
			}

			h := handlers.New(companies.NewConfig(""), service)

			app := fiber.New()
			app.Patch("/:id", h.Update())

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
