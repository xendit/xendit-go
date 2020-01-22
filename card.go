package xendit

import "time"

// CardCreateAuthorizationResponse contains data from Xendit's API response of credit card's Create Authorization related request.
// For more details see https://xendit.github.io/apireference/?bash#create-authorization.
type CardCreateAuthorizationResponse struct {
	ID                    string     `json:"id"`
	Status                string     `json:"status"`
	MerchantID            string     `json:"merchant_id"`
	Created               *time.Time `json:"created"`
	BusinessID            string     `json:"business_id"`
	AuthorizedAmount      float64    `json:"authorized_amount"`
	ExternalID            string     `json:"external_id"`
	MerchantReferenceCode string     `json:"merchant_reference_code"`
	ChargeType            string     `json:"charge_type"`
	CardBrand             string     `json:"card_brand"`
	MaskedCardNumber      string     `json:"masked_card_number"`
	CaptureAmount         float64    `json:"capture_amount,omitempty"`
	ECI                   string     `json:"eci,omitempty"`
	FailureReason         string     `json:"failure_reason,omitempty"`
	CardType              string     `json:"card_type,omitempty"`
	BankReconciliationID  string     `json:"bank_reconciliation_id,omitempty"`
	Descriptor            string     `json:"descriptor,omitempty"`
	MidLabel              string     `json:"mid_label,omitempty"`
	Currency              string     `json:"currency,omitempty"`
}

// CardReverseAuthorizationResponse contains data from Xendit's API response of credit card's Create Authorization related request.
// For more details see https://xendit.github.io/apireference/?bash#reverse-authorization.
type CardReverseAuthorizationResponse struct {
	ID                 string     `json:"id"`
	ExternalID         string     `json:"external_id"`
	CreditCardChargeID string     `json:"credit_card_charge_id"`
	BusinessID         string     `json:"business_id"`
	Amount             float64    `json:"amount"`
	Status             string     `json:"status"`
	Created            *time.Time `json:"created"`
	Currency           string     `json:"currency,omitempty"`
}
