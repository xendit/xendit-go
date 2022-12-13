package pmsv2

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

type CreatePaymentMethodParams struct {
	Type        constant.PaymentMethodTypeEnum `json:"type"`
	Country     constant.CountryEnum           `json:"country"`
	CustomerID  *string                        `json:"customer_id"`
	ReferenceID string                         `json:"reference_id,omitempty"`
	Reusability constant.ReusabilityEnum       `json:"reusability"`
	Description *string                        `json:"description"`
	Metadata    map[string]interface{}         `json:"metadata"`

	XCallbackURL string `json:"-"`
	MerchantID   string `json:"-"`

	Card           *card.CreateMethod           `json:"card"`
	DirectDebit    *directdebit.CreateMethod    `json:"direct_debit"`
	Ewallet        *ewallet.CreateMethod        `json:"ewallet"`
	OverTheCounter *overthecounter.CreateMethod `json:"over_the_counter"`
	VirtualAccount *virtualaccount.CreateMethod `json:"virtual_account"`
	QRCode         *qrcode.CreateMethod         `json:"qr_code"`
}

type ValidateOTPRequest struct {
	ID   string `param:"id"`
	Code string `json:"auth_code"`

	MerchantID string `json:"-"`
}

type ExpireRequest struct {
	ID                     string  `json:"-"`
	UnlinkSuccessReturnURL *string `json:"unlink_success_return_url"`

	MerchantID string `json:"-"`
}

type RetrievePaymentMethodRequest struct {
	ID         string `json:"-"`
	MerchantID string `json:"-"`
}

type RetrieveAllPaymentMethodsRequest struct {
	ID          []string                           `query:"id"`
	Type        constant.PaymentMethodTypeEnum     `query:"type"`
	Status      []constant.PaymentMethodStatusEnum `query:"status"`
	Reusability constant.ReusabilityEnum           `query:"reusability"`
	CustomerID  string                             `query:"customer_id"`
	ReferenceID string                             `query:"reference_id"`
	MerchantID  string                             `query:"business_id" header:"business-id"`
	AfterID     string                             `query:"after_id"`
	BeforeID    string                             `query:"before_id"`
	Limit       int                                `query:"limit"`
}

type UpdateRequest struct {
	ReferenceID    string                           `json:"reference_id"`
	Description    string                           `json:"description"`
	Status         constant.PaymentMethodStatusEnum `json:"status"`
	Reusability    constant.ReusabilityEnum         `json:"reusability"`
	OverTheCounter *overthecounter.MutableMethod    `json:"over_the_counter"`
	VirtualAccount *virtualaccount.MutableMethod    `json:"virtual_account"`

	MerchantID string `json:"-"`
	ID         string `json:"-"`
}

type RetrievePaymentsRequest struct {
	ID                string    `json:"-"`
	CreatedGTE        time.Time `query:"created[gte]" url:"created[gte],omitempty"`
	CreatedLTE        time.Time `query:"created[lte]" url:"created[lte],omitempty"`
	UpdatedGTE        time.Time `query:"updated[gte]" url:"updated[gte],omitempty"`
	UpdatedLTE        time.Time `query:"updated[lte]" url:"updated[lte],omitempty"`
	PaymentMethodType []string  `query:"payment_method_type" url:"payment_method_type,omitempty"`
	PaymentMethodID   []string  `query:"payment_method_id" url:"payment_method_id,omitempty"`
	PaymentRequestID  []string  `query:"payment_request_id" url:"payment_request_id,omitempty"`
	ChannelCode       []string  `query:"channel_code" url:"channel_code,omitempty"`
	Status            []string  `query:"status" url:"status,omitempty"`
	Currency          []string  `query:"currency" url:"currency,omitempty"`
	ReferenceID       []string  `query:"reference_id" url:"reference_id,omitempty"`
	MerchantID        string    `header:"business_id" url:"-"`
	PaginationFilters
}

type PaginationFilters struct {
	BeforeID string `query:"before_id" url:"before_id,omitempty"`
	AfterID  string `query:"after_id" url:"after_id,omitempty"`
	Limit    int    `query:"limit" url:"limit,omitempty"`
}
