package main

import (
	"context"
	"fmt"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/client"
	"github.com/xendit/xendit-go/invoice"
)

/* With Client */

func createInvoiceWithClient(ctx context.Context, x client.API) {
	data := xendit.CreateInvoiceParams{
		ExternalID:  "invoice-external-id",
		Amount:      200000,
		PayerEmail:  "customer@customer.com",
		Description: "invoice  #1",
	}

	resp, err := x.Invoice.CreateInvoice(&data)
	fmt.Println(resp, err)

	resp, err = x.Invoice.CreateInvoiceWithContext(ctx, &data)
	fmt.Println(resp, err)
}

func getInvoiceWithClient(ctx context.Context, x client.API) {
	resp, err := x.Invoice.GetInvoice("5e058d9af4d38b20d542012f")
	fmt.Println(resp, err)

	resp, err = x.Invoice.GetInvoiceWithContext(ctx, "5e058d9af4d38b20d542012f")
	fmt.Println(resp, err)
}

func expireInvoiceWithClient(ctx context.Context, x client.API) {
	resp, err := x.Invoice.ExpireInvoice("5e0ec2bff4d38b20d5422231")
	fmt.Println(resp, err)

	resp, err = x.Invoice.ExpireInvoiceWithContext(ctx, "5e0ec2bff4d38b20d5422231")
	fmt.Println(resp, err)
}

func getAllInvoicesWithClient(ctx context.Context, x client.API) {
	data := xendit.GetAllInvoicesParams{
		Statuses:     []string{"EXPIRED", "SETTLED"},
		Limit:        2,
		CreatedAfter: "2016-02-24T23:48:36.697Z",
	}

	resp, err := x.Invoice.GetAllInvoices(&data)
	fmt.Println(resp, err)

	resp, err = x.Invoice.GetAllInvoicesWithContext(ctx, &data)
	fmt.Println(resp, err)
}

/* Without Client */

func createInvoiceWithoutClient(ctx context.Context) {
	data := xendit.CreateInvoiceParams{
		ExternalID:  "invoice-" + time.Now().String(),
		Amount:      200000,
		PayerEmail:  "customer@customer.com",
		Description: "invoice  #1",
	}

	resp, err := invoice.CreateInvoice(&data)
	fmt.Println(resp, err)

	resp, err = invoice.CreateInvoiceWithContext(ctx, &data)
	fmt.Println(resp, err)
}

func getInvoiceWithoutClient(ctx context.Context) {
	resp, err := invoice.GetInvoice("5e058d9af4d38b20d542012f")
	fmt.Println(resp, err)

	resp, err = invoice.GetInvoiceWithContext(ctx, "5e058d9af4d38b20d542012f")
	fmt.Println(resp, err)
}

func expireInvoiceWithoutClient(ctx context.Context) {
	resp, err := invoice.ExpireInvoice("5e058d9af4d38b20d542012f")
	fmt.Println(resp, err)

	resp, err = invoice.ExpireInvoiceWithContext(ctx, "5e058d9af4d38b20d542012f")
	fmt.Println(resp, err)
}

func getAllInvoicesWithoutClient(ctx context.Context) {
	data := xendit.GetAllInvoicesParams{
		Statuses:     []string{"EXPIRED", "SETTLED"},
		Limit:        2,
		CreatedAfter: "2016-02-24T23:48:36.697Z",
	}

	resp, err := invoice.GetAllInvoices(&data)
	fmt.Println(resp, err)

	resp, err = invoice.GetAllInvoicesWithContext(ctx, &data)
	fmt.Println(resp, err)
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	x := client.New("", "xnd_development_REt02KJzkM6AootfKnDrMw1Sse4LlzEDHfKzXoBocqIEiH4bqjHUJXbl6Cfaab", "", nil)

	createInvoiceWithClient(ctx, *x)
	getInvoiceWithClient(ctx, *x)
	expireInvoiceWithClient(ctx, *x)
	getAllInvoicesWithClient(ctx, *x)

	xendit.Opt.SecretKey = "xnd_development_REt02KJzkM6AootfKnDrMw1Sse4LlzEDHfKzXoBocqIEiH4bqjHUJXbl6Cfaab"

	createInvoiceWithoutClient(ctx)
	getInvoiceWithoutClient(ctx)
	expireInvoiceWithoutClient(ctx)
	getAllInvoicesWithoutClient(ctx)
}
