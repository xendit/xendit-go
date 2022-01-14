package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/retailoutlet"
)

func main() {
	godotenvErr := godotenv.Load()
	if godotenvErr != nil {
		log.Fatal(godotenvErr)
	}
	xendit.Opt.SecretKey = os.Getenv("SECRET_KEY")

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
	fmt.Printf("created retail outlet fixed payment code: %+v\n", resp)

	getFixedPaymentCodeData := retailoutlet.GetFixedPaymentCodeParams{
		FixedPaymentCodeID: resp.ID,
	}

	resp, err = retailoutlet.GetFixedPaymentCode(&getFixedPaymentCodeData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retrieved retail outlet fixed payment code: %+v\n", resp)

	getPaymentByFixedPaymentCodeData := retailoutlet.GetPaymentByFixedPaymentCodeParams{
		FixedPaymentCodeID: resp.ID,
	}

	respPayments, err := retailoutlet.GetPaymentByFixedPaymentCode(&getPaymentByFixedPaymentCodeData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retrieved retail outlet payments: %+v\n", respPayments)

	expirationDate := time.Now().AddDate(0, 0, 1)

	updateFixedPaymentCodeData := retailoutlet.UpdateFixedPaymentCodeParams{
		FixedPaymentCodeID: resp.ID,
		Name:               "Billy Jackson",
		ExpectedAmount:     2000000,
		ExpirationDate:     &expirationDate,
	}

	resp, err = retailoutlet.UpdateFixedPaymentCode(&updateFixedPaymentCodeData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("updated retail outlet fixed payment code: %+v\n", resp)
}
