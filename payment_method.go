package xendit

import (
	"time"

	"github.com/xendit/xendit-go/pmsv2/card"
	"github.com/xendit/xendit-go/pmsv2/constant"
	"github.com/xendit/xendit-go/pmsv2/directdebit"
	"github.com/xendit/xendit-go/pmsv2/ewallet"
	"github.com/xendit/xendit-go/pmsv2/overthecounter"
	"github.com/xendit/xendit-go/pmsv2/qrcode"
	"github.com/xendit/xendit-go/pmsv2/virtualaccount"
)

type PaymentMethodActions []Action

type Action struct {
	Action     ActionType `json:"action"`
	URL        string     `json:"url"`
	URLType    URLType    `json:"url_type"`
	HTTPMethod string     `json:"method"`
}

type ActionType string

const (
	Auth       ActionType = "AUTH"
	ResendAuth ActionType = "RESEND_AUTH"
	Capture    ActionType = "CAPTURE"
	Cancel     ActionType = "CANCEL"
)

type URLType string

const (
	API      URLType = "API"
	Web      URLType = "WEB"
	Mobile   URLType = "MOBILE"
	Deeplink URLType = "DEEPLINK"
)

type Metadata map[string]interface{}

type BillingInformation struct {
	Country       constant.CountryEnum `json:"country"`
	StreetLine1   *string              `json:"street_line1"`
	StreetLine2   *string              `json:"street_line2"`
	City          *string              `json:"city"`
	ProvinceState *string              `json:"province_state"`
	PostalCode    *string              `json:"postal_code"`
}

type EwalletMethod struct {
	ChannelCode       ewallet.ChannelCode       `json:"channel_code"`
	ChannelProperties ewallet.ChannelProperties `json:"channel_properties"`
	Account           ewallet.Account           `json:"account"`
}

type DirectDebitMethod struct {
	LinkedAccountTokenID string                        `json:"-"`
	LinkedAccountID      string                        `json:"-"`
	ChannelCode          directdebit.ChannelCode       `json:"channel_code"`
	ChannelProperties    directdebit.ChannelProperties `json:"channel_properties"`
	Type                 directdebit.Type              `json:"type"`
	BankAccount          *directdebit.BankAccount      `json:"bank_account"`
	DebitCard            *directdebit.DebitCard        `json:"debit_card"`
}

type OverTheCounterMethod struct {
	Amount            float64                          `json:"amount"`
	Currency          constant.CurrencyEnum            `json:"currency"`
	ChannelCode       overthecounter.ChannelCode       `json:"channel_code"`
	ChannelProperties overthecounter.ChannelProperties `json:"channel_properties"`
}

type VirtualAccountMethod struct {
	Amount            float64                          `json:"amount"`
	MinimumAmount     float64                          `json:"min_amount,omitempty"`
	MaximumAmount     float64                          `json:"max_amount,omitempty"`
	Currency          constant.CurrencyEnum            `json:"currency"`
	ChannelCode       virtualaccount.ChannelCode       `json:"channel_code"`
	ChannelProperties virtualaccount.ChannelProperties `json:"channel_properties"`
}

type QRCodeMethod struct {
	Amount            float64                  `json:"amount"`
	Currency          constant.CurrencyEnum    `json:"currency"`
	ChannelCode       qrcode.ChannelCode       `json:"channel_code"`
	ChannelProperties qrcode.ChannelProperties `json:"channel_properties"`
}

type PaymentMethodResponse struct {
	ID                 string                           `json:"id"`
	Type               constant.PaymentMethodTypeEnum   `json:"type"`
	Country            constant.CountryEnum             `json:"country"`
	MerchantID         string                           `json:"business_id"`
	CustomerID         *string                          `json:"customer_id"`
	ReferenceID        string                           `json:"reference_id,omitempty"`
	Reusability        constant.ReusabilityEnum         `json:"reusability"`
	Status             constant.PaymentMethodStatusEnum `json:"status"`
	Actions            PaymentMethodActions             `json:"actions"`
	Description        *string                          `json:"description"`
	CreatedAt          time.Time                        `json:"created"`
	UpdatedAt          time.Time                        `json:"updated"`
	Metadata           *Metadata                        `json:"metadata"`
	BillingInformation *BillingInformation              `json:"billing_information"`
	FailureCode        *string                          `json:"failure_code"`

	Ewallet        *EwalletMethod        `json:"ewallet"`
	DirectDebit    *DirectDebitMethod    `json:"direct_debit"`
	Card           *card.Method          `json:"card"`
	OverTheCounter *OverTheCounterMethod `json:"over_the_counter"`
	QRCode         *QRCodeMethod         `json:"qr_code"`
	VirtualAccount *VirtualAccountMethod `json:"virtual_account"`
}
