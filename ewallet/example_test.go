package ewallet_test

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/ewallet"
)

func ExampleCreatePayment() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := ewallet.CreatePaymentParams{
		ExternalID:     "dana-ewallet",
		Amount:         20000,
		Phone:          "08123123123",
		EWalletType:    "DANA",
		CallbackURL:    "mystore.com/callback",
		RedirectURL:    "mystore.com/redirect",
		ExpirationDate: time.Now().Local().Add(60 * time.Hour),
	}

	resp, err := ewallet.CreatePayment(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created payment: %+v\n", resp)
}

func ExampleGetPaymentStatus() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := ewallet.GetPaymentStatusParams{
		ExternalID:  "data-ewallet",
		EWalletType: "DANA",
	}

	resp, err := ewallet.GetPaymentStatus(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved payment: %+v\n", resp)
}
