package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/ewallet"
)

func main() {
	xendit.Opt.SecretKey = "xnd_development_REt02KJzkM6AootfKnDrMw1Sse4LlzEDHfKzXoBocqIEiH4bqjHUJXbl6Cfaab"

	createPaymentData := ewallet.CreatePaymentParams{
		ExternalID:     "dana-" + time.Now().String(),
		Amount:         20000,
		Phone:          "08123123123",
		EWalletType:    "DANA",
		CallbackURL:    "mystore.com/callback",
		RedirectURL:    "mystore.com/redirect",
		ExpirationDate: time.Now().Local().Add(60 * time.Hour),
	}

	resp, err := ewallet.CreatePayment(&createPaymentData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created payment: %+v\n", resp)

	getPaymentStatusData := ewallet.GetPaymentStatusParams{
		ExternalID:  resp.ExternalID,
		EWalletType: resp.EWalletType,
	}

	resp, err = ewallet.GetPaymentStatus(&getPaymentStatusData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retrieved payment: %+v\n", resp)
}
