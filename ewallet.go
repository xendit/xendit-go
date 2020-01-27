package xendit

import "time"

// EWalletTypeEnum constants are the available e-wallet type
type EWalletTypeEnum string

// This consists the values that EWalletTypeEnum can take
const (
	EWalletTypeOVO     EWalletTypeEnum = "OVO"
	EWalletTypeDANA    EWalletTypeEnum = "DANA"
	EWalletTypeLINKAJA EWalletTypeEnum = "LINKAJA"
)

// EWallet contains data from Xendit's API response of eWallet related requests.
// For more details see https://xendit.github.io/apireference/?bash#ewallets.
type EWallet struct {
	EWalletType     EWalletTypeEnum `json:"ewallet_type"`
	ExternalID      string          `json:"external_id"`
	Amount          float64         `json:"amount"`
	TransactionDate *time.Time      `json:"transaction_date,omitempty"`
	CheckoutURL     string          `json:"checkout_url,omitempty"`
	BusinessID      string          `json:"business_id,omitempty"`
}
