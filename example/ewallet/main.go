package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/ewallet"
)

func main() {
	godotenvErr := godotenv.Load()
	if godotenvErr != nil {
		log.Fatal(godotenvErr)
	}
	xendit.Opt.SecretKey = os.Getenv("SECRET_KEY")

	createPaymentData := ewallet.CreatePaymentParams{
		ExternalID:  "dana-" + time.Now().String(),
		Amount:      20000,
		Phone:       "08123123123",
		EWalletType: xendit.EWalletTypeDANA,
		CallbackURL: "mystore.com/callback",
		RedirectURL: "mystore.com/redirect",
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
