package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/disbursement"
)

func main() {
	godotenvErr := godotenv.Load()
	if godotenvErr != nil {
		log.Fatal(godotenvErr)
	}
	xendit.Opt.SecretKey = os.Getenv("SECRET_KEY")

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
		log.Fatal(err)
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
