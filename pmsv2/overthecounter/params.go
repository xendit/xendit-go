package overthecounter

import (
	"time"

	"github.com/xendit/xendit-go/pmsv2/constant"
)

type ChannelCode string

const (
	// PH
	SevenEleven      ChannelCode = "7ELEVEN"
	SevenElevenCliqq ChannelCode = "7ELEVEN_CLIQQ"
	Cebuana          ChannelCode = "CEBUANA"
	ECPay            ChannelCode = "ECPAY"

	// PH DragonPay
	Palawan         ChannelCode = "PALAWAN"
	MLhuiller       ChannelCode = "MLHUILLIER"
	ECPayDragonLoan ChannelCode = "ECPAY_DRAGONLOAN"
	LBC             ChannelCode = "LBC"
	RDPawnshop      ChannelCode = "RD_PAWNSHOP"
	CVM             ChannelCode = "CVM"
	ECPaySchool     ChannelCode = "ECPAY_SCHOOL"
	USSC            ChannelCode = "USSC"

	// ID
	Alfamart  ChannelCode = "ALFAMART"
	Indomaret ChannelCode = "INDOMARET"
)

type ChannelProperties struct {
	CustomerName string     `json:"customer_name"`
	PaymentCode  *string    `json:"payment_code,omitempty"`
	ExpiresAt    *time.Time `json:"expires_at,omitempty"`
}

type CreateMethod struct {
	Amount            *float64               `json:"amount,omitempty"`
	Currency          *constant.CurrencyEnum `json:"currency,omitempty"`
	ChannelCode       ChannelCode            `json:"channel_code"`
	ChannelProperties ChannelProperties      `json:"channel_properties"`
}

type MutableMethod struct {
	Amount            *float64                  `json:"amount,omitempty"`
	Currency          *string                   `json:"currency,omitempty"`
	ChannelProperties *MutableChannelProperties `json:"channel_properties,omitempty"`
}

type MutableChannelProperties struct {
	CustomerName *string    `json:"customer_name,omitempty"`
	ExpiresAt    *time.Time `json:"expires_at,omitempty"`
}
