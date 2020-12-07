package card_test

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/card"
)

func ExampleCreateCharge() {
	xendit.Opt.SecretKey = "examplesecretkey"

	createChargeData := card.CreateChargeParams{
		TokenID:          "example-token-id",
		AuthenticationID: "example-authentication-id",
		ExternalID:       "cardAuth-1",
		Amount:           10000,
		Capture:          new(bool), // false
	}

	chargeResp, err := card.CreateCharge(&createChargeData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created charge: %+v\n", chargeResp)
}

func ExampleGetCharge() {
	xendit.Opt.SecretKey = "examplesecretkey"

	getChargeData := card.GetChargeParams{
		ChargeID: "123",
	}

	chargeResp, err := card.GetCharge(&getChargeData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved charge: %+v\n", chargeResp)
}

func ExampleCaptureCharge() {
	xendit.Opt.SecretKey = "examplesecretkey"

	captureChargeData := card.CaptureChargeParams{
		ChargeID: "123",
		Amount:   10000,
	}

	chargeResp, err := card.CaptureCharge(&captureChargeData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("captured charge: %+v\n", chargeResp)
}

func ExampleCreateRefund() {
	xendit.Opt.SecretKey = "examplesecretkey"

	createRefundData := card.CreateRefundParams{
		IdempotencyKey: "unique-idempotency-key",
		ChargeID:       "123",
		Amount:         10000,
		ExternalID:     "example-external-id",
	}

	refundResp, err := card.CreateRefund(&createRefundData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("refunded charge: %+v\n", refundResp)
}

func ExampleReverseAuthorization() {
	xendit.Opt.SecretKey = "examplesecretkey"

	reverseAuthorizationData := card.ReverseAuthorizationParams{
		ChargeID:   "123",
		ExternalID: "reverseAuth-id",
	}

	reverseAuthorizationResp, err := card.ReverseAuthorization(&reverseAuthorizationData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("reversed authorization: %+v\n", reverseAuthorizationResp)
}

func ExampleCreatePromotion() {
	xendit.Opt.SecretKey = "examplesecretkey"

	startTime := time.Now().Add(time.Hour)
	endTime := startTime.Add(time.Hour)

	createPromotionData := card.CreatePromotionParams{
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

	promotionResp, err := card.CreatePromotion(&createPromotionData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created promotion: %+v\n", promotionResp)
}
