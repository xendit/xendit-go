package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

func main() {
	godotenvErr := godotenv.Load()
	if godotenvErr != nil {
		log.Fatal(godotenvErr)
	}
	xendit.Opt.SecretKey = os.Getenv("SECRET_KEY")

	data := invoice.CreateParams{
		ExternalID:  "invoice-" + time.Now().String(),
		Amount:      200000,
		PayerEmail:  "customer@customer.com",
		Description: "invoice  #1",
	}

	resp, err := invoice.Create(&data)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Printf("created invoice: %+v\n", resp)

	resp, err = invoice.Get(&invoice.GetParams{
		ID: resp.ID,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retrieved invoice: %+v\n", resp)

	resp, err = invoice.Expire(&invoice.ExpireParams{
		ID: resp.ID,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("expired invoice: %+v\n", resp)

	resps, err := invoice.GetAll(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("first 10 invoices %+v\n", resps)
}
