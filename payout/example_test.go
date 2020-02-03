package payout_test

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/payout"
)

func ExampleCreate() {
	xendit.Opt.SecretKey = "examplesecretkey"

	createData := payout.CreateParams{
		ExternalID: "payout-" + time.Now().String(),
		Amount:     200000,
	}

	resp, err := payout.Create(&createData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created payout: %+v\n", resp)
}

func ExampleGet() {
	xendit.Opt.SecretKey = "examplesecretkey"

	resp, err := payout.Get(&payout.GetParams{
		ID: "123",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved payout: %+v\n", resp)
}

func ExampleVoid() {
	xendit.Opt.SecretKey = "examplesecretkey"

	resp, err := payout.Void(&payout.VoidParams{
		ID: "123",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("voided payout: %+v\n", resp)
}
