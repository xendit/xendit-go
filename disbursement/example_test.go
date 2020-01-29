package disbursement_test

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/disbursement"
)

func ExampleCreate() {
	xendit.Opt.SecretKey = "examplesecretkey"

	createData := disbursement.CreateParams{
		IdempotencyKey:    "disbursement" + time.Now().String(),
		ExternalID:        "disbursement-external",
		BankCode:          "BRI",
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

func ExampleGetByID() {
	xendit.Opt.SecretKey = "examplesecretkey"

	getByIDData := disbursement.GetByIDParams{
		DisbursementID: "123",
	}

	resp, err := disbursement.GetByID(&getByIDData)
	if err != nil {
		log.Fatal(err.ErrorCode, err.Message, err.GetStatus())
	}

	fmt.Printf("retrieved disbursement: %+v\n", resp)
}

func ExampleGetByExternalID() {
	xendit.Opt.SecretKey = "examplesecretkey"

	getByExternalIDData := disbursement.GetByExternalIDParams{
		ExternalID: "disbursement-external-id",
	}

	resps, err := disbursement.GetByExternalID(&getByExternalIDData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved disbursements: %+v\n", resps)
}

func ExampleGetAvailableBanks() {
	xendit.Opt.SecretKey = "examplesecretkey"

	availableBanks, err := disbursement.GetAvailableBanks()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("available disbursement banks: %+v\n", availableBanks)
}
