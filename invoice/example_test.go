package invoice_test

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

func ExampleCreate() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := invoice.CreateParams{
		ExternalID:  "invoice-example",
		Amount:      200000,
		PayerEmail:  "customer@customer.com",
		Description: "invoice #1",
	}

	resp, err := invoice.Create(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created invoice: %+v\n", resp)
}

func ExampleGet() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := invoice.GetParams{
		ID: "invoice-id",
	}

	resp, err := invoice.Get(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved invoice: %+v\n", resp)
}

func ExampleExpire() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := invoice.ExpireParams{
		ID: "invoice-id",
	}

	resp, err := invoice.Expire(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("expired invoice: %+v\n", resp)
}

func ExampleGetAll() {
	xendit.Opt.SecretKey = "examplesecretkey"

	createdAfter, _ := time.Parse(time.RFC3339, "2016-02-24T23:48:36.697Z")

	data := invoice.GetAllParams{
		Statuses:     []string{"EXPIRED", "SETTLED"},
		Limit:        5,
		CreatedAfter: createdAfter,
	}

	resps, err := invoice.GetAll(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("invoices: %+v\n", resps)
}
