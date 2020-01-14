package invoice_test

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

/* Without Client */

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

/* With Client */

// This example creates an invoice using client
func ExampleClient_Create() {
	client := invoice.Client{
		Opt: &xendit.Option{
			SecretKey: "examplesecretkey",
		},
		APIRequester: xendit.GetAPIRequester(),
	}

	data := invoice.CreateParams{
		ExternalID:  "invoice-example",
		Amount:      200000,
		PayerEmail:  "customer@customer.com",
		Description: "invoice #1",
	}

	resp, err := client.Create(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created invoice: %+v\n", resp)
}

// This example gets an invoice using client
func ExampleClient_Get() {
	client := invoice.Client{
		Opt: &xendit.Option{
			SecretKey: "examplesecretkey",
		},
		APIRequester: xendit.GetAPIRequester(),
	}

	data := invoice.GetParams{
		ID: "invoice-id",
	}

	resp, err := client.Get(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved invoice: %+v\n", resp)
}

// This example expires an invoice using client
func ExampleClient_Expire() {
	client := invoice.Client{
		Opt: &xendit.Option{
			SecretKey: "examplesecretkey",
		},
		APIRequester: xendit.GetAPIRequester(),
	}

	data := invoice.ExpireParams{
		ID: "invoice-id",
	}

	resp, err := client.Expire(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("expired invoice: %+v\n", resp)
}

// This example gets list of invoices using client
func ExampleClient_GetAll() {
	client := invoice.Client{
		Opt: &xendit.Option{
			SecretKey: "examplesecretkey",
		},
		APIRequester: xendit.GetAPIRequester(),
	}

	createdAfter, _ := time.Parse(time.RFC3339, "2016-02-24T23:48:36.697Z")

	data := invoice.GetAllParams{
		Statuses:     []string{"EXPIRED", "SETTLED"},
		Limit:        5,
		CreatedAfter: createdAfter,
	}

	resps, err := client.GetAll(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("invoices: %+v\n", resps)
}

/* Package-level */

// This examples shows the full usage of the package invoice
func Example_fullUsage() {
	xendit.Opt.SecretKey = "xnd_development_REt02KJzkM6AootfKnDrMw1Sse4LlzEDHfKzXoBocqIEiH4bqjHUJXbl6Cfaab"

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
