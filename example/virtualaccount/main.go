package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/virtualaccount"
)

func main() {
	xendit.Opt.SecretKey = "xnd_development_REt02KJzkM6AootfKnDrMw1Sse4LlzEDHfKzXoBocqIEiH4bqjHUJXbl6Cfaab"

	availableBanks, err := virtualaccount.GetAvailableBanks()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("available va banks: %+v\n", availableBanks)

	createFixedVAData := virtualaccount.CreateFixedVAParams{
		ExternalID: "va-" + time.Now().String(),
		BankCode:   availableBanks[0].Code,
		Name:       "Michael Jackson",
	}

	resp, err := virtualaccount.CreateFixedVA(&createFixedVAData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created fixed va: %+v\n", resp)

	getFixedVAData := virtualaccount.GetFixedVAParams{
		ID: resp.ID,
	}

	resp, err = virtualaccount.GetFixedVA(&getFixedVAData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retrieved fixed va: %+v\n", resp)

	expirationDate := time.Now().AddDate(0, 0, 1)

	updateFixedVAData := virtualaccount.UpdateFixedVAParams{
		ID:             resp.ID,
		ExpirationDate: &expirationDate,
	}

	resp, err = virtualaccount.UpdateFixedVA(&updateFixedVAData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("updated fixed va: %+v\n", resp)

	payment, err := virtualaccount.GetPayment(&virtualaccount.GetPaymentParams{
		PaymentID: "VA_fixed-1579507045_1579507068112",
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retrieved va payment: %+v\n", payment)
}
