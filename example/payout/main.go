package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/xendit/xendit-go/payout"

	"github.com/xendit/xendit-go"
)

func main() {
	godotenvErr := godotenv.Load()
	if godotenvErr != nil {
		log.Fatal(godotenvErr)
	}
	xendit.Opt.SecretKey = os.Getenv("SECRET_KEY")

	createData := payout.CreateParams{
		ExternalID: "payout-" + time.Now().String(),
		Amount:     200000,
		Email:      "customer@customer.com",
	}

	resp, err := payout.Create(&createData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created payout: %+v\n", resp)

	resp, err = payout.Get(&payout.GetParams{
		ID: resp.ID,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retrieved payout: %+v\n", resp)

	resp, err = payout.Void(&payout.VoidParams{
		ID: resp.ID,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("voided payout: %+v\n", resp)
}
