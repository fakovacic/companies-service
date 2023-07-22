package integrationtests_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/fakovacic/companies-service/sdk"
)

const host = "http://localhost:8080"

var httpClient = &http.Client{}

func TestCompanies(t *testing.T) {
	ctx := context.Background()

	client := sdk.NewClient(host, httpClient)

	tokenRes, err := client.Login(ctx, "test")
	if err != nil {
		t.Fatal("login:", err)
	}

	createReq := &sdk.CreateRequest{
		Name:            "Mock company",
		Description:     "Mock company description",
		EmployeesAmount: 21,
		Registered:      true,
		Type:            "corporations",
	}

	createResp, err := client.Create(ctx, tokenRes.Token, createReq)
	if err != nil {
		t.Fatal("create company:", err)
	}

	if createResp.Name != createReq.Name {
		t.Fatal("create company name does not match")
	}

	if createResp.Description != createReq.Description {
		t.Fatal("create company description does not match")
	}

	if createResp.EmployeesAmount != createReq.EmployeesAmount {
		t.Fatal("create company employees amount does not match")
	}

	if createResp.Registered != createReq.Registered {
		t.Fatal("create company registered does not match")
	}

	getResp, err := client.Get(ctx, tokenRes.Token, createResp.ID)
	if err != nil {
		t.Fatal("get company:", err)
	}

	if createResp.ID != getResp.ID {
		t.Fatal("create and get response IDs do not match")
	}

	updateData := sdk.UpdateCompany{
		Name: "Mock new",
	}

	updateResp, err := client.Update(ctx, tokenRes.Token, createResp.ID, &sdk.UpdateRequest{
		Fields: []string{"name"},
		Data:   updateData,
	})
	if err != nil {
		t.Fatal("update company:", err)
	}

	if updateResp.Name != updateData.Name {
		t.Fatal("update company name does not match")
	}

	err = client.Delete(ctx, tokenRes.Token, createResp.ID)
	if err != nil {
		t.Fatal("delete company:", err)
	}

	_, err = client.Get(ctx, tokenRes.Token, createResp.ID)
	if err == nil {
		t.Fatal("get company should fail after delete")
	}
}
