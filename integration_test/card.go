package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go/card"
)

func cardTest() {
	_, err := card.GetCharge(&card.GetChargeParams{
		ChargeID: "5e046a736113354249aab8bd",
	})
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Card integration tests done!")
}

func cardPromotionTest() {
	startTime := time.Now().Add(time.Hour)
	endTime := startTime.Add(time.Hour)

	_, err := card.CreatePromotion(&card.CreatePromotionParams{
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
	fmt.Println("Card's promotion integration tests done!")
}
