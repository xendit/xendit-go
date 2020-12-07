package xendit

import "time"

// Promotion contains data from Xendit's API response of promotion-related request.
// For more details see https://xendit.github.io/apireference/?bash#create-promotion.
// For documentation of subpackage card, checkout https://pkg.go.dev/github.com/xendit/xendit-go/card
type Promotion struct {
	ID                string     `json:"id"`
	BusinessID        string     `json:"business_id"`
	Status            string     `json:"status"`
	ReferenceID       string     `json:"reference_id"`
	Description       string     `json:"description"`
	PromoCode         string     `json:"promo_code"`
	BinList           []string   `json:"bin_list"`
	ChannelCode       string     `json:"channel_code"`
	DiscountPercent   float64    `json:"discount_percent"`
	DiscountAmount    float64    `json:"discount_amount"`
	Currency          string     `json:"currency"`
	StartTime         *time.Time `json:"start_time"`
	EndTime           *time.Time `json:"end_time"`
	MinOriginalAmount int        `json:"min_original_amount"`
	MaxDiscountAmount int        `json:"max_discount_amount"`
}
