package disbursementph

import (
	"net/url"

	"github.com/xendit/xendit-go"
)

// CreateParams contains parameters for Create
type CreateParams struct {
	IdempotencyKey      string                     `json:"xendit_idempotency_key" validate:"required"`
	ReferenceID         string                     `json:"reference_id" validate:"required"`
	Currency            string                     `json:"currency" validate:"required"`
	ChannelCode         string                     `json:"channel_code" validate:"required"`
	AccountName         string                     `json:"account_name" validate:"required"`
	AccountNumber       string                     `json:"account_number" validate:"required"`
	Description         string                     `json:"description" validate:"required"`
	Amount              float64                    `json:"amount" validate:"required"`
	Beneficiary         xendit.Beneficiary         `json:"beneficiary,omitempty"`
	ReceiptNotification xendit.ReceiptNotification `json:"receipt_notification,omitempty"`
	Metadata            map[string]interface{}     `json:"metadata,omitempty"`
}

// GetByIDParams contains parameters for GetByID
type GetByIDParams struct {
	DisbursementID string `json:"disbursement_id" validate:"required"`
}

// GetByReferenceIDParams contains parameters for GetByReferenceID
type GetByReferenceIDParams struct {
	ReferenceID string `json:"reference_id" validate:"required"`
}

// QueryString creates query string from GetByReferenceIDParams, ignores nil values
func (p *GetByReferenceIDParams) QueryString() string {
	urlValues := &url.Values{}

	urlValues.Add("reference_id", p.ReferenceID)

	return urlValues.Encode()
}

// GetByChannelCategoryParams contains parameters for GetDisbursementChannelByChannelCategory
type GetByChannelCategoryParams struct {
	ChannelCategory string `json:"channel_category" validate:"required"`
}

// QueryString creates query string from GetByChannelCategoryParams, ignores nil values
func (p *GetByChannelCategoryParams) QueryString() string {
	urlValues := &url.Values{}

	urlValues.Add("channel_category", p.ChannelCategory)

	return urlValues.Encode()
}

// GetByChannelCodeParams contains parameters for GetDisbursementChannelByChannelCode
type GetByChannelCodeParams struct {
	ChannelCode string `json:"channel_code" validate:"required"`
}

// QueryString creates query string from GetByChannelCodeParams, ignores nil values
func (p *GetByChannelCodeParams) QueryString() string {
	urlValues := &url.Values{}

	urlValues.Add("channel_code", p.ChannelCode)

	return urlValues.Encode()
}
