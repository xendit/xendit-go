package virtualaccount_test

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/virtualaccount"
)

func ExampleCreateFixed() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := virtualaccount.CreateFixedParams{
		ExternalID: "va-example",
		BankCode:   "BRI",
		Name:       "Michael Jackson",
	}

	resp, err := virtualaccount.CreateFixed(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created fixed va: %+v\n", resp)
}

func ExampleGetFixed() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := virtualaccount.GetFixedParams{
		ID: "va-id",
	}

	resp, err := virtualaccount.GetFixed(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved fixed va: %+v\n", resp)
}

func ExampleUpdateFixed() {
	xendit.Opt.SecretKey = "examplesecretkey"

	expirationDate := time.Now().AddDate(0, 0, 1)

	updateFixedData := virtualaccount.UpdateFixedParams{
		ID:             "va-id",
		ExpirationDate: &expirationDate,
		ExpectedAmount: 100000,
	}

	resp, err := virtualaccount.UpdateFixed(&updateFixedData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("updated fixed va: %+v\n", resp)
}

func ExampleGetAvailableBanks() {
	xendit.Opt.SecretKey = "examplesecretkey"

	availableBanks, err := virtualaccount.GetAvailableBanks()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("available va banks: %+v\n", availableBanks)
}

func ExampleGetPayment() {
	xendit.Opt.SecretKey = "examplesecretkey"

	payment, err := virtualaccount.GetPayment(&virtualaccount.GetPaymentParams{
		PaymentID: "va-payment-id",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved va payment: %+v\n", payment)
}
