package main

import (
	"fmt"
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

	resp, _ := invoice.CreateInvoice(&data)
	fmt.Println("created invoice:", resp)

	resp, _ = invoice.GetInvoice(resp.ID)
	fmt.Println("retrieved invoice:", resp)

	resp, _ = invoice.ExpireInvoice(resp.ID)
	fmt.Println("expired invoice:", resp)

	resps, _ := invoice.GetAllInvoices(nil)
	fmt.Println("first 10 invoices", resps)
}
