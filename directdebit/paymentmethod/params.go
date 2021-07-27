package paymentmethod

import (
	"net/url"

	"github.com/xendit/xendit-go"
)

// CreatePaymentMethodParams contains parameters for CreatePaymentMethod
type CreatePaymentMethodParams struct {
	ForUserID		string						`json:"-"`
	CustomerID		string						`json:"customer_id" validate:"required"`
	Type 			xendit.AccountTypeEnum		`json:"type" validate:"required"`
	Properties		map[string]interface{}		`json:"properties" validate:"required"`
	Metadata		map[string]interface{}		`json:"metadata,omitempty"`
}

// GetPaymentMethodsByCustomerIDParams contains parameters for GetPaymentMethodsByCustomerID
type GetPaymentMethodsByCustomerIDParams struct {
	ForUserID		string						`json:"-"`
	CustomerID 		string 						`json:"customer_id" validate:"required"`
}

// QueryString creates query string from GetPaymentMethodsByCustomerIDParams, ignores nil values
func (p *GetPaymentMethodsByCustomerIDParams) QueryString() string {
	urlValues := &url.Values{}

	urlValues.Add("customer_id", p.CustomerID)

	return urlValues.Encode()
}
