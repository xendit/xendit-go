package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

func main() {
	xendit.Opt.SecretKey = "xnd_development_REt02KJzkM6AootfKnDrMw1Sse4LlzEDHfKzXoBocqIEiH4bqjHUJXbl6Cfaab"

	data := xendit.CreateInvoiceParams{
		ExternalID:  "invoice-" + time.Now().String(),
		Amount:      200000,
		PayerEmail:  "customer@customer.com",
		Description: "invoice  #1",
	}

	resp, err := invoice.CreateInvoice(&data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created invoice: %+v\n", resp)

	resp, err = invoice.GetInvoice(resp.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retrieved invoice: %+v\n", resp)

	resp, err = invoice.ExpireInvoice(resp.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("expired invoice: %+v\n", resp)

	resps, err := invoice.GetAllInvoices(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("first 10 invoices %+v\n", resps)
}
