package card

type CardOnFileType string

const (
	CustomerUnscheduledCOF CardOnFileType = "CUSTOMER_UNSCHEDULED"
	MerchantUnscheduledCOF CardOnFileType = "MERCHANT_UNSCHEDULED"
	RecurringCOF           CardOnFileType = "RECURRING"
)

type CreateMethod struct {
	Currency          *string                `json:"currency"`
	ChannelProperties *CardChannelProperties `json:"channel_properties"`
	CardInformation   *CardInformation       `json:"card_information"`
}

type Method CreateMethod

type CardChannelProperties struct {
	SkipThreeDSecure *bool   `json:"skip_three_d_secure"`
	SuccessReturnURL *string `json:"success_return_url"`
	FailureReturnURL *string `json:"failure_return_url"`
	CardOnFileType   *string `json:"cardonfile_type"`
}

type Device struct {
	Fingerprint string `json:"fingerprint"`
}

type CardInformation struct {
	CardNumber     string  `json:"card_number"`
	ExpiryMonth    string  `json:"expiry_month"`
	ExpiryYear     string  `json:"expiry_year"`
	CardCVN        *string `json:"cvv"`
	CardholderName *string `json:"cardholder_name"`
}

type BillingInformation struct {
	Country       string  `json:"country"`
	StreetLine1   *string `json:"street_line1"`
	StreetLine2   *string `json:"street_line2"`
	City          *string `json:"city"`
	ProvinceState *string `json:"province_state"`
	Province      *string `json:"province"`
	PostalCode    *string `json:"postal_code"`
}
