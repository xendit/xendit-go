package retailoutlet_test

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/retailoutlet"
)

func ExampleCreateFixedPaymentCode() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := retailoutlet.CreateFixedPaymentCodeParams{
		ExternalID:       "retailoutlet-external-id",
		RetailOutletName: xendit.RetailOutletNameAlfamart,
		Name:             "Michael Jackson",
		ExpectedAmount:   200000,
	}

	resp, err := retailoutlet.CreateFixedPaymentCode(&data)
	if err != nil {
		log.Fatal(err.ErrorCode)
	}

	fmt.Printf("created retail outlet fixed payment code: %+v\n", resp)
}

func ExampleGetFixedPaymentCode() {
	xendit.Opt.SecretKey = "examplesecretkey"

	getFixedPaymentCodeData := retailoutlet.GetFixedPaymentCodeParams{
		FixedPaymentCodeID: "123",
	}

	resp, err := retailoutlet.GetFixedPaymentCode(&getFixedPaymentCodeData)
	if err != nil {
		log.Fatal(err.ErrorCode)
	}

	fmt.Printf("retrieved retail outlet fixed payment code: %+v\n", resp)
}

func ExampleUpdateFixedPaymentCode() {
	xendit.Opt.SecretKey = "examplesecretkey"

	expirationDate := time.Now().AddDate(0, 0, 1)

	updateFixedPaymentCodeData := retailoutlet.UpdateFixedPaymentCodeParams{
		FixedPaymentCodeID: "123",
		Name:               "Billy Jackson",
		ExpectedAmount:     2000000,
		ExpirationDate:     &expirationDate,
	}

	resp, err := retailoutlet.UpdateFixedPaymentCode(&updateFixedPaymentCodeData)
	if err != nil {
		log.Fatal(err.ErrorCode)
	}
	fmt.Printf("updated retail outlet fixed payment code: %+v\n", resp)
}
