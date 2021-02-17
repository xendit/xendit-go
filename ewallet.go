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

// EWallet contains data from Xendit's API response of e-wallet related requests.
// For more details see https://xendit.github.io/apireference/?bash#ewallets.
// For documentation of subpackage ewallet, checkout https://pkg.go.dev/github.com/xendit/xendit-go/ewallet
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

type BasketItem struct {
	ReferenceID string                 `json:"reference_id" validate:"required"`
	Name        string                 `json:"name" validate:"required"`
	Category    string                 `json:"category" validate:"required"`
	Currency    string                 `json:"currency" validate:"required"`
	Price       float64                `json:"price" validate:"required"`
	Quantity    int                    `json:"quantity" validate:"required"`
	Type        string                 `json:"type" validate:"required"`
	Url         string                 `json:"url,omitempty"`
	Description string                 `json:"description,omitempty"`
	SubCategory string                 `json:"sub_category,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type EWalletCharge struct {
	ID                 string                 `json:"id"`
	BusinessID         string                 `json:"business_id"`
	ReferenceID        string                 `json:"reference_id"`
	Status             string                 `json:"status"`
	Currency           string                 `json:"currency"`
	ChargeAmount       float64                `json:"charge_amount"`
	CaptureAmount      float64                `json:"capture_amount"`
	CheckoutMethod     string                 `json:"checkout_method"`
	ChannelCode        string                 `json:"channel_code"`
	ChannelProperties  map[string]string      `json:"channel_properties"`
	Actions            map[string]string      `json:"actions"`
	IsRedirectRequired bool                   `json:"is_redirect_required"`
	CallbackURL        string                 `json:"callback_url"`
	Created            string                 `json:"created"`
	Updated            string                 `json:"updated"`
	VoidedAt           string                 `json:"voided_at,omitempty"`
	CaptureNow         bool                   `json:"capture_now"`
	CustomerID         string                 `json:"customer_id,omitempty"`
	PaymentMethodID    string                 `json:"payment_method_id,omitempty"`
	FailureCode        string                 `json:"failure_code,omitempty"`
	Basket             []BasketItem           `json:"basket,omitempty"`
	Metadata           map[string]interface{} `json:"metadata,omitempty"`
}
