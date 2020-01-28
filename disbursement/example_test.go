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

func ExampleCreateBatch() {
	xendit.Opt.SecretKey = "examplesecretkey"

	createBatchData := disbursement.CreateBatchParams{
		Reference: "batch_disbursement" + time.Now().String(),
		Disbursements: []disbursement.DisbursementItem{
			{
				Amount:            200000,
				BankCode:          "BNI",
				BankAccountName:   "Michael Jackson",
				BankAccountNumber: "1234567890",
				Description:       "Batch disbursement from Go",
			},
			{
				Amount:            100000,
				BankCode:          "BRI",
				BankAccountName:   "Michael Jackson",
				BankAccountNumber: "1234567890",
				Description:       "Batch disbursement from Go 2",
			},
		},
	}

	batchDisbursementResp, err := disbursement.CreateBatch(&createBatchData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created batch disbursement: %+v\n", batchDisbursementResp)
}
