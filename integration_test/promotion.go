package main

import (
	"fmt"
	"github.com/xendit/xendit-go/promotion"
	"log"
	"time"
)

func promotionTest() {
	startTime := time.Now().Add(time.Hour)
	endTime := startTime.Add(time.Hour)

	created, err := promotion.CreatePromotion(&promotion.CreatePromotionParams{
		ReferenceID: "BRI_20_JAN",
		Description: "20% discount applied for all BRI cards",
		BinList: []string{
			"400000",
			"460000",
		},
		DiscountPercent:   20,
		Currency:          "IDR",
		StartTime:         &startTime,
		EndTime:           &endTime,
		ChannelCode:       "BRI",
		MinOriginalAmount: 25000,
		MaxDiscountAmount: 5000,
	})
	if err != nil {
		log.Panic(err)
	}

	_, err = promotion.GetPromotions(&promotion.GetPromotionsParams{
		ReferenceID: created.ReferenceID,
	})
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Promotion's integration tests done!")
}
