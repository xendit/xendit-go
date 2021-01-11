package promotion_test

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/promotion"
)

func ExampleCreatePromotion() {
	xendit.Opt.SecretKey = "examplesecretkey"

	startTime := time.Now().Add(time.Hour)
	endTime := startTime.Add(time.Hour)

	createPromotionData := promotion.CreatePromotionParams{
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
	}

	promotionResp, err := promotion.CreatePromotion(&createPromotionData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created promotion: %+v\n", promotionResp)
}

func ExampleGetPromotions() {
	xendit.Opt.SecretKey = "examplesecretkey"

	promotions, err := promotion.GetPromotions(&promotion.GetPromotionsParams{
		ReferenceID: "BRI_20_JAN",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved promotions: %+v\n", promotions)
}

func ExampleUpdatePromotion() {
	xendit.Opt.SecretKey = "examplesecretkey"

	updatePromotionData := promotion.UpdatePromotionParams{
		PromotionID: "36ab1517-208a-4f22-b155-96fb101cb378",
		Description: "20% discount applied for all BCA cards",
		BinList: []string{
			"123123",
			"456456",
		},
		DiscountPercent: 20,
		Currency:        "IDR",
		ChannelCode:     "BCA",
	}

	promotionResp, err := promotion.UpdatePromotion(&updatePromotionData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("updated promotion: %+v\n", promotionResp)
}

func ExampleDeletePromotion() {
	xendit.Opt.SecretKey = "examplesecretkey"

	deletePromotionData := promotion.DeletePromotionParams{
		PromotionID: "36ab1517-208a-4f22-b155-96fb101cb378",
	}

	promotionResp, err := promotion.DeletePromotion(&deletePromotionData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("deleted promotion: %+v\n", promotionResp)
}
