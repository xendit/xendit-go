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

// EWalletXApiVersionEnum constants are the available e-wallet "X-API-VERSION"s
type EWalletXApiVersionEnum string

// This consists the values that EWalletXApiVersionEnum can take
const (
	EWalletXApiVersionEnum20190204 EWalletXApiVersionEnum = "2019-02-04"
	EWalletXApiVersionEnum20200201 EWalletXApiVersionEnum = "2020-02-01"
)

// EWallet contains data from Xendit's API response of e-wallet related requests.
// For more details see https://xendit.github.io/apireference/?bash#ewallets.
type EWallet struct {
	EWalletType          EWalletTypeEnum `json:"ewallet_type"`
	ExternalID           string          `json:"external_id"`
	Status               string          `json:"status"`
	Amount               float64         `json:"amount"`
	TransactionDate      *time.Time      `json:"transaction_date,omitempty"`
	CheckoutURL          string          `json:"checkout_url,omitempty"`
	BusinessID           string          `json:"business_id,omitempty"`
	Created              *time.Time      `json:"created,omitempty"`
	EWalletTransactionID string          `json:"e_wallet_transaction_id,omitempty"`
}
