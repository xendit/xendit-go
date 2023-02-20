package qrcode

import (
	"time"

	"github.com/xendit/xendit-go/pmsv2/constant"
)

type ChannelCode string

const (
	QRIS      ChannelCode = "QRIS"
	DANA      ChannelCode = "DANA"
	RCBC      ChannelCode = "RCBC"
	PromptPay ChannelCode = "PROMPTPAY"
)

const (
	EMPTY   ChannelCode = ""
	LinkAja ChannelCode = "LINKAJA"
)

type ChannelProperties struct {
	QRString  string     `json:"qr_string"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

type CreateMethod struct {
	Amount            *float64               `json:"amount,omitempty"`
	Currency          *constant.CurrencyEnum `json:"currency,omitempty"`
	ChannelCode       ChannelCode            `json:"channel_code"`
	ChannelProperties ChannelProperties      `json:"channel_properties"`
}
