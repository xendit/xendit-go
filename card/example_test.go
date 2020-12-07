package card_test

import (
	"fmt"
	"log"

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
