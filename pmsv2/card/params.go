package card

import (
	"github.com/xendit/xendit-go/pmsv2/constant"
)

type CardOnFileType string

const (
	CustomerUnscheduledCOF CardOnFileType = "CUSTOMER_UNSCHEDULED"
	MerchantUnscheduledCOF CardOnFileType = "MERCHANT_UNSCHEDULED"
	RecurringCOF           CardOnFileType = "RECURRING"
)

type CreateMethod struct {
	Currency          *constant.CurrencyEnum `json:"currency"`
	ChannelProperties *CardChannelProperties `json:"channel_properties"`
	Token             *Token                 `json:"token"`
}

type Method CreateMethod

type CardChannelProperties struct {
	SkipThreeDSecure *bool   `json:"skip_three_d_secure"`
	SuccessReturnURL *string `json:"success_return_url"`
	FailureReturnURL *string `json:"failure_return_url"`
	CardOnFileType   *string `json:"cardonfile_type"`
}

type Token struct {
	ID               string                `json:"id"`
	PaymentMethodID  string                `json:"payment_method_id"`
	ReferenceID      string                `json:"reference_id"`
	Status           constant.TokenStatus  `json:"status"`
	Currency         constant.CurrencyEnum `json:"currency"`
	Authentication   *Authentication       `json:"authentication"`
	Authorization    *Authorization        `json:"authorization"`
	RedirectURL      *string               `json:"redirect_url"`
	MaskedCardNumber *string               `json:"masked_card_number"`
	ExpiryMonth      *string               `json:"expiry_month"`
	ExpiryYear       *string               `json:"expiry_year"`
	CardInfo         *CardInfo             `json:"card_info"`
	BillingDetails   *BillingDetails       `json:"billing_details"`
	CofType          *CardOnFileType       `json:"cof_type"`
}

type Authentication struct {
	ID            *string `json:"id"`
	Status        *string `json:"status"`
	ThreeDSResult *string `json:"three_ds_result"`
}

type Authorization struct {
	ID      *string         `json:"id"`
	Status  *string         `json:"status"`
	Result  *string         `json:"result"`
	CofType *CardOnFileType `json:"cof_type"`
}

type CardInfo struct {
	Bank        *string               `json:"bank"`
	Country     *constant.CountryEnum `json:"country"`
	Type        *string               `json:"type"`
	Brand       *string               `json:"brand"`
	Fingerprint *string               `json:"fingerprint"`
}
type BillingDetails struct {
	FirstName         *string `json:"first_name"`
	AddressLine1      *string `json:"address_line_1,omitempty"`
	AddressLine2      *string `json:"address_line_2,omitempty"`
	AddressCity       *string `json:"address_city,omitempty"`
	AddressCountry    *string `json:"address_country,omitempty"`
	AddressPostalCode *string `json:"address_postal_code,omitempty"`
	AddressState      *string `json:"address_state,omitempty"`
}
