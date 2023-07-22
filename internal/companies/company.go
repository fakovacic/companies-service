package companies

import (
	"strings"
	"time"
	"unicode/utf8"

	"github.com/fakovacic/companies-service/internal/companies/errors"
)

const (
	maxNameLength = 15
)

type Company struct {
	ID              string
	Name            string
	Description     string
	EmployeesAmount uint32
	Registered      bool
	Type            CompanyType
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (c Company) Validate() error {
	if c.Name == "" {
		return errors.BadRequest("name is required")
	}

	if utf8.RuneCountInString(c.Name) > maxNameLength {
		return errors.BadRequest("name is too long")
	}

	if c.EmployeesAmount == 0 {
		return errors.BadRequest("employees amount is required")
	}

	return nil
}

var CompanyField = struct {
	ID              string
	Name            string
	Description     string
	EmployeesAmount string
	Registered      string
	Type            string
	CreatedAt       string
	UpdatedAt       string
}{
	ID:              "id",
	Name:            "name",
	Description:     "description",
	EmployeesAmount: "employeesAmount",
	Registered:      "registered",
	Type:            "type",
	CreatedAt:       "createdAt",
	UpdatedAt:       "updatedAt",
}

var CompanyUpdateFields = map[string]bool{
	CompanyField.ID:              false,
	CompanyField.Name:            true,
	CompanyField.Description:     true,
	CompanyField.EmployeesAmount: true,
	CompanyField.Registered:      true,
	CompanyField.Type:            true,
	CompanyField.CreatedAt:       false,
	CompanyField.UpdatedAt:       false,
}

const (
	TypeCorporations       CompanyType = "corporations"
	TypeNonProfit          CompanyType = "nonProfit"
	TypeCooperative        CompanyType = "cooperative"
	TypeSoleProprietorship CompanyType = "soleProprietorship"
)

type CompanyType string

func (t *CompanyType) Parse(s string) error {
	s = strings.Trim(s, "\"")
	switch s {
	case "corporations":
		*t = TypeCorporations
	case "nonProfit":
		*t = TypeNonProfit
	case "cooperative":
		*t = TypeCooperative
	case "soleProprietorship":
		*t = TypeSoleProprietorship
	case "":
	default:
		return errors.BadRequest("invalid type '%s'", s)
	}

	return nil
}

func (t CompanyType) String() string {
	return string(t)
}

func (t *CompanyType) UnmarshalJSON(b []byte) error {
	return t.Parse(string(b))
}
