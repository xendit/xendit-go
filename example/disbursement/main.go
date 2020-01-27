package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/disbursement"
)

func main() {
	xendit.Opt.SecretKey = "xnd_development_REt02KJzkM6AootfKnDrMw1Sse4LlzEDHfKzXoBocqIEiH4bqjHUJXbl6Cfaab"

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

	getByIDData := disbursement.GetByIDParams{
		DisbursementID: resp.ID,
	}

	resp, err = disbursement.GetByID(&getByIDData)
	if err != nil {
		log.Fatal(err.ErrorCode, err.Message, err.GetStatus())
	}
	fmt.Printf("retrieved disbursement: %+v\n", resp)

	getByExternalIDData := disbursement.GetByExternalIDParams{
		ExternalID: resp.ExternalID,
	}

	resps, err := disbursement.GetByExternalID(&getByExternalIDData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retrieved disbursements: %+v\n", resps)
}
