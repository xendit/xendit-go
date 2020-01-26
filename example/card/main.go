package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/card"
)

func main() {
	xendit.Opt.SecretKey = "xnd_development_REt02KJzkM6AootfKnDrMw1Sse4LlzEDHfKzXoBocqIEiH4bqjHUJXbl6Cfaab"

	createChargeData := card.CreateChargeParams{
		TokenID:          "5e2ab828d97c174c58bcd9f6",
		AuthenticationID: "5e2ab828d97c174c58bcd9f7",
		ExternalID:       "cardAuth-" + time.Now().String(),
		Amount:           10000,
		Capture:          new(bool),
	}

	chargeResp, err := card.CreateCharge(&createChargeData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created charge: %+v\n", chargeResp)

	getChargeData := card.GetChargeParams{
		ChargeID: chargeResp.ID,
	}

	chargeResp, err = card.GetCharge(&getChargeData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retrieved charge: %+v\n", chargeResp)

	captureChargeData := card.CaptureChargeParams{
		ChargeID: chargeResp.ID,
		Amount:   10000,
	}

	chargeResp, err = card.CaptureCharge(&captureChargeData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("captured charge: %+v\n", chargeResp)

	createRefundData := card.CreateRefundParams{
		IdempotencyKey: "idempotency-" + time.Now().String(),
		ChargeID:       "5e2abc61d97c174c58bcda30",
		Amount:         10000,
		ExternalID:     "refund-" + time.Now().String(),
	}

	refundResp, err := card.CreateRefund(&createRefundData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("refunded charge: %+v\n", refundResp)

	reverseAuthorizationData := card.ReverseAuthorizationParams{
		ChargeID:   chargeResp.ID,
		ExternalID: "reverseAuth-" + time.Now().String(),
	}

	reverseAuthorizationResp, err := card.ReverseAuthorization(&reverseAuthorizationData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("reversed authorization: %+v\n", reverseAuthorizationResp)
}
