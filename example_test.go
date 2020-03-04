package xendit_test

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/client"
	"github.com/xendit/xendit-go/disbursement"
)

func ExampleDisbursement_createWithoutClient() {
	xendit.Opt.SecretKey = "xnd_..."

	availableBanks, err := disbursement.GetAvailableBanks()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("available disbursement banks: %+v\n", availableBanks)

	createData := disbursement.CreateParams{
		IdempotencyKey:    "disbursement" + time.Now().String(),
		ExternalID:        "disbursement-external",
		BankCode:          availableBanks[0].Code,
		AccountHolderName: "Michael Jackson",
		AccountNumber:     "1234567890",
		Description:       "Disbursement from Go",
		Amount:            200000,
	}

	resp, err := disbursement.Create(&createData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created disbursement: %+v\n", resp)
}

func ExampleDisbursement_createWithClient() {
	cli := client.New("xnd_...")

	availableBanks, err := cli.Disbursement.GetAvailableBanks()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("available disbursement banks: %+v\n", availableBanks)

	createData := disbursement.CreateParams{
		IdempotencyKey:    "disbursement" + time.Now().String(),
		ExternalID:        "disbursement-external",
		BankCode:          availableBanks[0].Code,
		AccountHolderName: "Michael Jackson",
		AccountNumber:     "1234567890",
		Description:       "Disbursement from Go",
		Amount:            200000,
	}

	resp, err := cli.Disbursement.Create(&createData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created disbursement: %+v\n", resp)
}
