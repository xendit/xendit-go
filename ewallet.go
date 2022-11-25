package xendit

import "time"

//IMPORTANT: EWALLET PAYMENT WILL BE DEPRECATED. PLEASE USE EWALLET CHARGE.

// EWalletTypeEnum constants are the available e-wallet type for Ewallet Payment
type EWalletTypeEnum string

// This consists the values that EWalletTypeEnum can take
const (
	EWalletTypeOVO       EWalletTypeEnum = "OVO"
	EWalletTypeDANA      EWalletTypeEnum = "DANA"
	EWalletTypeLINKAJA   EWalletTypeEnum = "LINKAJA"
	EWalletTypeSHOPEEPAY EWalletTypeEnum = "SHOPEEPAY"
	EWalletTypeGCASH     EWalletTypeEnum = "GCASH"
	EWalletTypeGRABPAY   EWalletTypeEnum = "GRABPAY"
	EWalletTypePAYMAYA   EWalletTypeEnum = "PAYMAYA"
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

type EWalletBasketItem struct {
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

// ChargeOutputStatus represents the enum for charge output status for an Ewallet Charge
type ChargeOutputStatus string

// These are the available ChargeOutputStatus values for an Ewallet Charge
const (
	ChargeOutputStatusPending   ChargeOutputStatus = "PENDING"
	ChargeOutputStatusFailed    ChargeOutputStatus = "FAILED"
	ChargeOutputStatusSucceeded ChargeOutputStatus = "SUCCEEDED"
	ChargeOutputStatusVoided    ChargeOutputStatus = "VOIDED"
	ChargeOutputStatusRefunded  ChargeOutputStatus = "REFUNDED"
)

// Actions represents the actions of an Ewallet Charge.
type Actions struct {
	DesktopWebCheckoutURL     *string `json:"desktop_web_checkout_url"`
	MobileWebCheckoutURL      *string `json:"mobile_web_checkout_url"`
	MobileDeeplinkCheckoutURL *string `json:"mobile_deeplink_checkout_url"`
	QrCheckoutString          *string `json:"qr_checkout_string"`
}

// ShippingInformation struct for Ewallet Charge ShippingInformation
type ShippingInformation struct {
	Country       string  `json:"country"` //two letter country code eg. ID/PH
	StreetLine1   string  `json:"street_line1"`
	StreetLine2   *string `json:"street_line2,omitempty"`
	City          string  `json:"city"`
	ProvinceState string  `json:"province_state"`
	PostalCode    string  `json:"postal_code"`
}

//
type EwalletCustomer struct {
	ReferenceId            *string `json:"reference_id,omitempty"`
	GivenNames             *string `json:"given_names,omitempty"`
	Surname                *string `json:"surname,omitempty"`
	Email                  *string `json:"email,omitempty"`
	MobileNumber           *string `json:"mobile_number,omitempty"`
	DomicileOfRegistration *string `json:"domicile_of_registration,omitempty"`
	DateOfRegistration     *string `json:"date_of_registration,omitempty"`
}

//EWalletCharge represents the response from Xendit Ewallet Charge API
type EWalletCharge struct {
	ID                   string                 `json:"id"`
	BusinessID           string                 `json:"business_id"`
	ReferenceID          string                 `json:"reference_id"`
	Status               ChargeOutputStatus     `json:"status"`
	Currency             string                 `json:"currency"`
	ChargeAmount         float64                `json:"charge_amount"`
	CaptureAmount        float64                `json:"capture_amount"`
	PayerChargedCurrency string                 `json:"payer_charged_currency,omitempty"`
	PayerChargedAmount   float64                `json:"payer_charged_amount,omitempty"`
	RefundedAmount       float64                `json:"refunded_amount,omitempty"`
	CheckoutMethod       string                 `json:"checkout_method"`
	ChannelCode          string                 `json:"channel_code"`
	ChannelProperties    map[string]string      `json:"channel_properties"`
	Actions              Actions                `json:"actions"`
	IsRedirectRequired   bool                   `json:"is_redirect_required"`
	CallbackURL          string                 `json:"callback_url"`
	Created              string                 `json:"created"`
	Updated              string                 `json:"updated"`
	VoidStatus           string                 `json:"void_status,omitempty"`
	VoidedAt             string                 `json:"voided_at,omitempty"`
	CaptureNow           bool                   `json:"capture_now"`
	CustomerID           string                 `json:"customer_id,omitempty"`
	Customer             *EwalletCustomer       `json:"customer,omitempty"`
	PaymentMethodID      string                 `json:"payment_method_id,omitempty"`
	FailureCode          string                 `json:"failure_code,omitempty"`
	Basket               []EWalletBasketItem    `json:"basket,omitempty"`
	Metadata             map[string]interface{} `json:"metadata,omitempty"`
	ShippingInformation  *ShippingInformation   `json:"shipping_information,omitempty"`
}
