package customer

import (
	"net/url"

	"github.com/xendit/xendit-go"
)

// CreateCustomerParams contains parameters for CreateCustomer
type CreateCustomerParams struct {
	ForUserID    string                   `json:"-"`
	ReferenceID  string                   `json:"reference_id" validate:"required"`
	MobileNumber string                   `json:"mobile_number,omitempty"`
	Email        string                   `json:"email,omitempty"`
	GivenNames   string                   `json:"given_names" validate:"required"`
	MiddleName   string                   `json:"middle_name,omitempty"`
	Surname      string                   `json:"surname,omitempty"`
	Description  string                   `json:"description,omitempty"`
	PhoneNumber  string                   `json:"phone_number,omitempty"`
	Nationality  string                   `json:"nationality,omitempty"`
	Addresses    []xendit.CustomerAddress `json:"addresses,omitempty"`
	DateOfBirth  string                   `json:"date_of_birth,omitempty"`
	Metadata     map[string]interface{}   `json:"metadata,omitempty"`
}

// GetCustomerByReferenceIDParams contains parameters for GetCustomerByReferenceID
type GetCustomerByReferenceIDParams struct {
	ReferenceID string `json:"reference_id" validate:"required"`
}

// QueryString creates query string from GetCustomerByReferenceIDParams, ignores nil values
func (p *GetCustomerByReferenceIDParams) QueryString() string {
	urlValues := &url.Values{}

	urlValues.Add("reference_id", p.ReferenceID)

	return urlValues.Encode()
}
