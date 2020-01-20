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

	createFixedData := virtualaccount.CreateFixedParams{
		ExternalID: "va-" + time.Now().String(),
		BankCode:   availableBanks[0].Code,
		Name:       "Michael Jackson",
	}

	resp, err := virtualaccount.CreateFixed(&createFixedData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created fixed va: %+v\n", resp)

	getFixedData := virtualaccount.GetFixedParams{
		ID: resp.ID,
	}

	resp, err = virtualaccount.GetFixed(&getFixedData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retrieved fixed va: %+v\n", resp)

	expirationDate := time.Now().AddDate(0, 0, 1)

	updateFixedData := virtualaccount.UpdateFixedParams{
		ID:             resp.ID,
		ExpirationDate: &expirationDate,
	}

	resp, err = virtualaccount.UpdateFixed(&updateFixedData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("updated fixed va: %+v\n", resp)

	// payment, err := virtualaccount.GetPayment(&virtualaccount.GetPaymentParams{
	// 	PaymentID: "VA_fixed-1579507045_1579507068112",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("retrieved va payment: %+v\n", payment)
}
