package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Companies interface {
	Login(context.Context, string) (*LoginResponse, error)
	Get(context.Context, string, string) (*Company, error)
	Create(context.Context, string, *CreateRequest) (*Company, error)
	Update(context.Context, string, string, *UpdateRequest) (*Company, error)
	Delete(context.Context, string, string) error
}

type LoginResponse struct {
	Token string `json:"token"`
}

type CreateRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description,omitempty"`
	EmployeesAmount uint32 `json:"employeesAmount"`
	Registered      bool   `json:"registered"`
	Type            string `json:"type"`
}

type UpdateRequest struct {
	Fields []string      `json:"fields"`
	Data   UpdateCompany `json:"data"`
}

type UpdateCompany struct {
	Name            string `json:"name"`
	Description     string `json:"description,omitempty"`
	EmployeesAmount uint32 `json:"employeesAmount"`
	Registered      bool   `json:"registered"`
	Type            string `json:"type"`
}

type Company struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	EmployeesAmount uint32    `json:"employeesAmount"`
	Registered      bool      `json:"registered"`
	Type            string    `json:"type"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type ResponseError struct {
	Message string `json:"message"`
}

func (e ResponseError) Error() string {
	return e.Message
}

func NewClient(host string, client HTTPClient) Companies {
	return &companies{
		host:   host,
		client: client,
	}
}

type companies struct {
	host   string
	client HTTPClient
}

func (r *companies) Login(ctx context.Context, key string) (*LoginResponse, error) {
	req, err := http.NewRequestWithContext(ctx,
		http.MethodPost,
		fmt.Sprintf("%s/%s", r.host, "login"),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("login handler request:%w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth", key)

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request to login:%w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read login body response:%w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var eMsg ResponseError

		err = json.Unmarshal(body, &eMsg)
		if err != nil {
			return nil, fmt.Errorf("unmarshal error body response from login:%w", err)
		}

		return nil, eMsg
	}

	var res *LoginResponse

	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, fmt.Errorf("unmarshal body response from login:%w", err)
	}

	return res, nil
}

func (r *companies) Get(ctx context.Context, token, id string) (*Company, error) {
	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet,
		fmt.Sprintf("%s/%s/%s", r.host, "companies", id),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("get companies handler request:%w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request to get companies:%w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read get companies body response:%w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var eMsg ResponseError

		err = json.Unmarshal(body, &eMsg)
		if err != nil {
			return nil, fmt.Errorf("unmarshal error body response from get companies:%w", err)
		}

		return nil, eMsg
	}

	var res *Company

	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, fmt.Errorf("unmarshal body response from get companies:%w", err)
	}

	return res, nil
}

func (r *companies) Create(ctx context.Context, token string, in *CreateRequest) (*Company, error) {
	requestBody, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("marshal companies create request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx,
		http.MethodPost,
		fmt.Sprintf("%s/%s", r.host, "companies"),
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return nil, fmt.Errorf("create companies handler request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request to create companies: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read create companies body response: %w", err)
	}

	if resp.StatusCode != http.StatusCreated {
		var eMsg ResponseError

		err = json.Unmarshal(body, &eMsg)
		if err != nil {
			return nil, fmt.Errorf("unmarshal error body response from create companies: %w", err)
		}

		return nil, eMsg
	}

	var res *Company

	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, fmt.Errorf("unmarshal body response from get companies: %w", err)
	}

	return res, nil
}

func (r *companies) Update(ctx context.Context, token, id string, in *UpdateRequest) (*Company, error) {
	requestBody, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("marshal companies update request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx,
		http.MethodPatch,
		fmt.Sprintf("%s/%s/%s", r.host, "companies", id),
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return nil, fmt.Errorf("update companies handler request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request to update companies: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read update companies body response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var eMsg ResponseError

		err = json.Unmarshal(body, &eMsg)
		if err != nil {
			return nil, fmt.Errorf("unmarshal error body response from update companies: %w", err)
		}

		return nil, eMsg
	}

	var res *Company

	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, fmt.Errorf("unmarshal body response from update companies: %w", err)
	}

	return res, nil
}

func (r *companies) Delete(ctx context.Context, token, id string) error {
	req, err := http.NewRequestWithContext(ctx,
		http.MethodDelete,
		fmt.Sprintf("%s/%s/%s", r.host, "companies", id),
		nil,
	)
	if err != nil {
		return fmt.Errorf("delete companies handler request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := r.client.Do(req)
	if err != nil {
		return fmt.Errorf("request to delete companies: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read delete companies body response: %w", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		var eMsg ResponseError

		err = json.Unmarshal(body, &eMsg)
		if err != nil {
			return fmt.Errorf("unmarshal error body response from delete companies: %w", err)
		}

		return eMsg
	}

	return nil
}
