package xendit

import (
	"time"
)

// VirtualAccount contains data from Xendit's API response of virtual account related request.
// For more details see https://xendit.github.io/apireference/?bash#virtual-accounts.
type VirtualAccount struct {
	OwnerID         string     `json:"owner_id"`
	ExternalID      string     `json:"external_id"`
	BankCode        string     `json:"bank_code"`
	MerchantCode    string     `json:"merchant_code"`
	Name            string     `json:"name"`
	AccountNumber   string     `json:"account_number"`
	IsClosed        bool       `json:"is_closed"`
	ID              string     `json:"id"`
	IsSingleUse     bool       `json:"is_single_use"`
	Status          string     `json:"status"`
	ExpirationDate  *time.Time `json:"expiration_date"`
	SuggestedAmount float64    `json:"suggested_amount,omitempty"`
	ExpectedAmount  float64    `json:"expected_amount,omitempty"`
	Description     string     `json:"description,omitempty"`
}
