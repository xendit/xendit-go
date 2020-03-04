package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/retailoutlet"
)

func retailoutletTest() {
	createFixedPaymentCodeData := retailoutlet.CreateFixedPaymentCodeParams{
		ExternalID:       "retailoutlet-" + time.Now().String(),
		RetailOutletName: xendit.RetailOutletNameAlfamart,
		Name:             "Michael Jackson",
		ExpectedAmount:   200000,
	}

	resp, err := retailoutlet.CreateFixedPaymentCode(&createFixedPaymentCodeData)
	if err != nil {
		log.Fatal(err)
	}

	getFixedPaymentCodeData := retailoutlet.GetFixedPaymentCodeParams{
		FixedPaymentCodeID: resp.ID,
	}

	resp, err = retailoutlet.GetFixedPaymentCode(&getFixedPaymentCodeData)
	if err != nil {
		log.Fatal(err)
	}

	expirationDate := time.Now().AddDate(0, 0, 1)
	updateFixedPaymentCodeData := retailoutlet.UpdateFixedPaymentCodeParams{
		FixedPaymentCodeID: resp.ID,
		Name:               "Billy Jackson",
		ExpectedAmount:     2000000,
		ExpirationDate:     &expirationDate,
	}

	_, err = retailoutlet.UpdateFixedPaymentCode(&updateFixedPaymentCodeData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Retail Outlet integration tests done!")
}
