package xendit

import "time"

// Promotion contains data from Xendit's API response of promotion-related request.
// For more details see https://xendit.github.io/apireference/?bash#create-promotion.
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
	MinOriginalAmount float64    `json:"min_original_amount"`
	MaxDiscountAmount float64    `json:"max_discount_amount"`
}

// PromotionCalculation contains data from Xendit's API response of get promotion calculation request.
// For more details see https://developers.xendit.co/api-reference/?bash#get-promotions-calculation.
type PromotionCalculation struct {
	ReferenceID       string  `json:"reference_id"`
	OriginalAmount    float64 `json:"original_amount"`
	DiscountPercent   float64 `json:"discount_percent"`
	DiscountAmount    float64 `json:"discount_amount"`
	FinalAmount       float64 `json:"final_amount"`
	Currency          string  `json:"currency"`
	Description       string  `json:"description"`
	MinOriginalAmount float64 `json:"min_original_amount"`
	MaxDiscountAmount float64 `json:"max_discount_amount"`
}
