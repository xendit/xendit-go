package promotion

import (
	"net/url"
	"time"
)

// CreatePromotionParams contains parameters for CreatePromotion
type CreatePromotionParams struct {
	ReferenceID       string     `json:"reference_id" validate:"required"`
	Description       string     `json:"description" validate:"required"`
	PromoCode         string     `json:"promo_code"`
	BinList           []string   `json:"bin_list"`
	ChannelCode       string     `json:"channel_code"`
	DiscountPercent   float64    `json:"discount_percent"`
	DiscountAmount    float64    `json:"discount_amount"`
	Currency          string     `json:"currency" validate:"required"`
	StartTime         *time.Time `json:"start_time" validate:"required"`
	EndTime           *time.Time `json:"end_time" validate:"required"`
	MinOriginalAmount int        `json:"min_original_amount"`
	MaxDiscountAmount int        `json:"max_discount_amount"`
}

// GetPromotionsParams contains parameters for GetPromotions
type GetPromotionsParams struct {
	ReferenceID string `json:"reference_id"`
	Status      string `json:"status"`
	Bin         string `json:"bin"`
	ChannelCode string `json:"channel_code"`
	Currency    string `json:"currency"`
}

// QueryString creates query string from GetPromotionsParams, ignores nil values
func (p *GetPromotionsParams) QueryString() string {
	urlValues := &url.Values{}

	urlValues.Add("reference_id", p.ReferenceID)
	urlValues.Add("status", p.Status)
	urlValues.Add("bin", p.Bin)
	urlValues.Add("channel_code", p.ChannelCode)
	urlValues.Add("currency", p.Currency)

	return urlValues.Encode()
}
