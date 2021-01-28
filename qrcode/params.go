package qrcode

import (
	"net/url"
	"strconv"

	"github.com/xendit/xendit-go"
)

// CreateQRCodeParams contains parameters for CreateQRCode
type CreateQRCodeParams struct {
	ForUserID   string            `json:"-"`
	ExternalID  string            `json:"external_id" validate:"required"`
	Type        xendit.QRCodeType `json:"type" validate:"required"`
	CallbackURL string            `json:"callback_url" validate:"required"`
	Amount      float64           `json:"amount,omitempty"`
}

// GetQRCodeParams contains parameters for GetQRCode
type GetQRCodeParams struct {
	ForUserID  string `json:"-"`
	ExternalID string `json:"external_id" validate:"required"`
}

// GetQRCodePaymentsParams contains parameters for GetQRCodePayments
type GetQRCodePaymentsParams struct {
	ExternalID string `json:"external_id" validate:"required"`
	Limit      int    `json:"limit,omitempty"`
	From       string `json:"from,omitempty"`
	To         string `json:"to,omitempty"`
}

// QueryString creates query string from GetQRCodePaymentsParams, ignores nil values
func (p *GetQRCodePaymentsParams) QueryString() string {
	urlValues := &url.Values{}

	urlValues.Add("external_id", p.ExternalID)
	if p.Limit > 0 {
		urlValues.Add("limit", strconv.Itoa(p.Limit))
	}

	if p.From != "" {
		urlValues.Add("from", p.From)
	}

	if p.To != "" {
		urlValues.Add("to", p.To)
	}

	return urlValues.Encode()
}
