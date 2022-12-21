package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go/disbursement"
)

func disbursementTest() {
	availableBanks, err := disbursement.GetAvailableBanks()
	if err != nil {
		log.Panic(err)
	}

	createData := disbursement.CreateParams{
		IdempotencyKey:    "disbursement" + time.Now().String(),
		ExternalID:        "disbursement-external",
		BankCode:          availableBanks[0].Code,
		AccountHolderName: "Michael Jackson",
		AccountNumber:     "123124123",
		Description:       "Disbursement from Go",
		Amount:            200000,
	}
	fmt.Println("aaa")
	resp, err := disbursement.Create(&createData)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("aaa")

	getByIDData := disbursement.GetByIDParams{
		DisbursementID: resp.ID,
	}
	resp, err = disbursement.GetByID(&getByIDData)
	if err != nil {
		log.Panic(err)
	}

	getByExternalIDData := disbursement.GetByExternalIDParams{
		ExternalID: resp.ExternalID,
	}
	_, err = disbursement.GetByExternalID(&getByExternalIDData)
	if err != nil {
		log.Panic(err)
	}

	createBatchData := disbursement.CreateBatchParams{
		Reference: "batch_disbursement" + time.Now().String(),
		Disbursements: []disbursement.DisbursementItem{
			{
				Amount:            200000,
				BankCode:          availableBanks[0].Code,
				BankAccountName:   "Michael Jackson",
				BankAccountNumber: "1234567890",
				Description:       "Batch disbursement from Go",
			},
			{
				Amount:            100000,
				BankCode:          availableBanks[1].Code,
				BankAccountName:   "Michael Jackson",
				BankAccountNumber: "1234567890",
				Description:       "Batch disbursement from Go 2",
			},
		},
	}

	_, err = disbursement.CreateBatch(&createBatchData)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Disbursement integration tests done!")
}
