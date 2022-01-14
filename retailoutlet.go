package xendit

import "time"

// RetailOutletNameEnum constants are the available retail outlet names
type RetailOutletNameEnum string

// This consists the values that RetailOutletNameEnum can take
const (
	RetailOutletNameAlfamart  RetailOutletNameEnum = "ALFAMART"
	RetailOutletNameIndomaret RetailOutletNameEnum = "INDOMARET"
)

// RetailOutlet contains data from Xendit's API response of retail outlet related requests.
// For more details see https://xendit.github.io/apireference/?bash#retail-outlets.
// For documentation of subpackage retailoutlet, checkout https://pkg.go.dev/github.com/xendit/xendit-go/retailoutlet
type RetailOutlet struct {
	IsSingleUse      bool                 `json:"is_single_use"`
	Status           string               `json:"status"`
	OwnerID          string               `json:"owner_id"`
	ExternalID       string               `json:"external_id"`
	RetailOutletName RetailOutletNameEnum `json:"retail_outlet_name"`
	Prefix           string               `json:"prefix"`
	Name             string               `json:"name"`
	PaymentCode      string               `json:"payment_code"`
	Type             string               `json:"type"`
	ExpectedAmount   float64              `json:"expected_amount"`
	ExpirationDate   *time.Time           `json:"expiration_date"`
	ID               string               `json:"id"`
}

// RetailOutletPayments contains data from Xendit's API response of Retail Outlet Get Payments By Fixed Payment Code ID requests.
// For more details see https://developers.xendit.co/api-reference/#get-payments-by-fixed-payment-code-id.
// For documentation of subpackage direct debit payment, checkout https://pkg.go.dev/github.com/xendit/xendit-go/retailoutlet/
type RetailOutletPayments struct {
	Data    []RetailOutletPayment     `json:"data"`
	HasMore bool                      `json:"has_more"`
	Links   RetailOutletPaymentsLinks `json:"links"`
}

// RetailOutletPaymentsLinks is data that contained in `RetailOutletPayments` at Links field.
type RetailOutletPaymentsLinks struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}

// RetailOutletPayments contains data from Xendit's API response of Retail Outlet Get Payments By Fixed Payment Code ID requests.
// For more details see https://developers.xendit.co/api-reference/#get-payments-by-fixed-payment-code-id.
// For documentation of subpackage direct debit payment, checkout https://pkg.go.dev/github.com/xendit/xendit-go/retailoutlet/

type RetailOutletPayment struct {
	ID                        string               `json:"id"`
	FixedPaymentCodeID        string               `json:"fixed_payment_code_id"`
	FixedPaymentCodePaymentID string               `json:"fixed_payment_code_payment_id"`
	PaymentCode               string               `json:"payment_code"`
	RetailOutletName          RetailOutletNameEnum `json:"retail_outlet_name"`
	Amount                    float64              `json:"amount"`
	Status                    string               `json:"status"`
	OwnerID                   string               `json:"owner_id"`
	Name                      string               `json:"name"`
	Prefix                    string               `json:"prefix"`
	PaymentID                 string               `json:"payment_id"`
	ExternalID                string               `json:"external_id"`
	TransactionTimestamp      *time.Time           `json:"transaction_timestamp"`
}
