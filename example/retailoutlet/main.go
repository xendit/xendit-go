package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/retailoutlet"
)

func main() {
	xendit.Opt.SecretKey = "xnd_development_REt02KJzkM6AootfKnDrMw1Sse4LlzEDHfKzXoBocqIEiH4bqjHUJXbl6Cfaab"

	createFixedPaymentCodeData := retailoutlet.CreateFixedPaymentCodeParams{
		ExternalID:       "retailoutlet-" + time.Now().String(),
		RetailOutletName: "ALFAMART",
		Name:             "Michael Jackson",
		ExpectedAmount:   200000,
	}

	resp, err := retailoutlet.CreateFixedPaymentCode(&createFixedPaymentCodeData)
	if err != nil {
		log.Fatal(err.ErrorCode)
	}
	fmt.Printf("created retail outlet fixed payment code: %+v\n", resp)

	getFixedPaymentCodeData := retailoutlet.GetFixedPaymentCodeParams{
		FixedPaymentCodeID: resp.ID,
	}

	resp, err = retailoutlet.GetFixedPaymentCode(&getFixedPaymentCodeData)
	if err != nil {
		log.Fatal(err.ErrorCode)
	}
	fmt.Printf("retrieved retail outlet fixed payment code: %+v\n", resp)

	expirationDate := time.Now().AddDate(0, 0, 1)

	updateFixedPaymentCodeData := retailoutlet.UpdateFixedPaymentCodeParams{
		FixedPaymentCodeID: resp.ID,
		Name:               "Billy Jackson",
		ExpectedAmount:     2000000,
		ExpirationDate:     &expirationDate,
	}

	resp, err = retailoutlet.UpdateFixedPaymentCode(&updateFixedPaymentCodeData)
	if err != nil {
		log.Fatal(err.ErrorCode)
	}
	fmt.Printf("updated retail outlet fixed payment code: %+v\n", resp)
}
