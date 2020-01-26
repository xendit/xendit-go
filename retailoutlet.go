package xendit

import "time"

// RetailOutlet contains data from Xendit's API response of retail outlet related requests.
// For more details see https://xendit.github.io/apireference/?bash#retail-outlets
type RetailOutlet struct {
	IsSingleUse      bool       `json:"is_single_use"`
	Status           string     `json:"status"`
	OwnerID          string     `json:"owner_id"`
	ExternalID       string     `json:"external_id"`
	RetailOutletName string     `json:"retail_outlet_name"`
	Prefix           string     `json:"prefix"`
	Name             string     `json:"name"`
	PaymentCode      string     `json:"payment_code"`
	Type             string     `json:"type"`
	ExpectedAmount   float64    `json:"expected_amount"`
	ExpirationDate   *time.Time `json:"expiration_date"`
	ID               string     `json:"id"`
}
