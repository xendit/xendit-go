package promotion

import (
	"net/url"
	"time"
)

// CreatePromotionParams contains parameters for CreatePromotion.
type CreatePromotionParams struct {
	ReferenceID       string     `json:"reference_id" validate:"required"`
	Description       string     `json:"description" validate:"required"`
	PromoCode         string     `json:"promo_code,omitempty"`
	BinList           []string   `json:"bin_list,omitempty"`
	ChannelCode       string     `json:"channel_code,omitempty"`
	DiscountPercent   float64    `json:"discount_percent,omitempty"`
	DiscountAmount    float64    `json:"discount_amount,omitempty"`
	Currency          string     `json:"currency" validate:"required"`
	StartTime         *time.Time `json:"start_time" validate:"required"`
	EndTime           *time.Time `json:"end_time" validate:"required"`
	MinOriginalAmount float64    `json:"min_original_amount,omitempty"`
	MaxDiscountAmount float64    `json:"max_discount_amount,omitempty"`
}

// GetPromotionsParams contains parameters for GetPromotions.
type GetPromotionsParams struct {
	ReferenceID string
	Status      string
	Bin         string
	ChannelCode string
	Currency    string
}

// QueryString creates query string from GetPromotionsParams, ignores nil values.
func (p *GetPromotionsParams) QueryString() string {
	urlValues := &url.Values{}

	if p.ReferenceID != "" {
		urlValues.Add("reference_id", p.ReferenceID)
	}
	if p.Status != "" {
		urlValues.Add("status", p.Status)
	}
	if p.Bin != "" {
		urlValues.Add("bin", p.Bin)
	}
	if p.ChannelCode != "" {
		urlValues.Add("channel_code", p.ChannelCode)
	}
	if p.Currency != "" {
		urlValues.Add("currency", p.Currency)
	}

	return urlValues.Encode()
}

// UpdatePromotionParams contains parameters for UpdatePromotion.
type UpdatePromotionParams struct {
	PromotionID       string     `json:"-" validate:"required"`
	Description       string     `json:"description,omitempty"`
	PromoCode         string     `json:"promo_code,omitempty"`
	BinList           []string   `json:"bin_list,omitempty"`
	ChannelCode       string     `json:"channel_code,omitempty"`
	DiscountPercent   float64    `json:"discount_percent,omitempty"`
	DiscountAmount    float64    `json:"discount_amount,omitempty"`
	Currency          string     `json:"currency,omitempty"`
	StartTime         *time.Time `json:"start_time,omitempty"`
	EndTime           *time.Time `json:"end_time,omitempty"`
	MinOriginalAmount float64    `json:"min_original_amount,omitempty"`
	MaxDiscountAmount float64    `json:"max_discount_amount,omitempty"`
}

// DeletePromotionParams contains parameters for DeletePromotion.
type DeletePromotionParams struct {
	PromotionID string `validate:"required"`
}
