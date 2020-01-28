package virtualaccount_test

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/virtualaccount"
)

func ExampleCreateFixedVA() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := virtualaccount.CreateFixedVAParams{
		ExternalID: "va-example",
		BankCode:   "BRI",
		Name:       "Michael Jackson",
	}

	resp, err := virtualaccount.CreateFixedVA(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created fixed va: %+v\n", resp)
}

func ExampleGetFixedVA() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := virtualaccount.GetFixedVAParams{
		ID: "va-id",
	}

	resp, err := virtualaccount.GetFixedVA(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved fixed va: %+v\n", resp)
}

func ExampleUpdateFixedVA() {
	xendit.Opt.SecretKey = "examplesecretkey"

	expirationDate := time.Now().AddDate(0, 0, 1)

	updateFixedVAData := virtualaccount.UpdateFixedVAParams{
		ID:             "va-id",
		ExpirationDate: &expirationDate,
		ExpectedAmount: 100000,
	}

	resp, err := virtualaccount.UpdateFixedVA(&updateFixedVAData)
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
