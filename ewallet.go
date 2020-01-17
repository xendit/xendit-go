package xendit

// EWallet contains data from Xendit's API response of eWallet related request.
// For more details see https://xendit.github.io/apireference/?bash#ewallets.
type EWallet struct {
	EWalletType     string  `json:"ewallet_type"`
	ExternalID      string  `json:"external_id"`
	Amount          float64 `json:"amount"`
	TransactionDate string  `json:"transaction_date,omitempty"`
	CheckoutURL     string  `json:"checkout_url,omitempty"`
	BusinessID      string  `json:"business_id,omitempty"`
}
