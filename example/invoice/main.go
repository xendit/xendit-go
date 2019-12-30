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

func createInvoiceWithClientTest(ctx context.Context, x client.API) {
	data := xendit.CreateInvoiceParams{
		ExternalID:  "invoice-" + time.Now().String(),
		Amount:      200000,
		PayerEmail:  "customer@customer.com",
		Description: "invoice test #1",
	}

	resp, err := x.Invoice.CreateInvoice(ctx, &data)

	fmt.Println(resp, err)
}

func getInvoiceWithClientTest(ctx context.Context, x client.API) {
	resp, err := x.Invoice.GetInvoice(ctx, "5e058d9af4d38b20d542012f")

	fmt.Println(resp, err)
}

func expireInvoiceWithClientTest(ctx context.Context, x client.API) {
	resp, err := x.Invoice.ExpireInvoice(ctx, "5e058d9af4d38b20d542012f")

	fmt.Println(resp, err)
}

func getAllInvoicesWithClientTest(ctx context.Context, x client.API) {
	data := xendit.GetAllInvoicesParams{
		Statuses:     []string{"EXPIRED", "SETTLED"},
		Limit:        2,
		CreatedAfter: "2016-02-24T23:48:36.697Z",
	}

	resp, err := x.Invoice.GetAllInvoices(ctx, &data)

	fmt.Println(resp, err)
}

/* Without Client */

func createInvoiceWithoutClientTest(ctx context.Context) {
	data := xendit.CreateInvoiceParams{
		ExternalID:  "invoice-" + time.Now().String(),
		Amount:      200000,
		PayerEmail:  "customer@customer.com",
		Description: "invoice test #1",
	}

	resp, err := invoice.CreateInvoice(ctx, &data)

	fmt.Println(resp, err)
}

func getInvoiceWithoutClientTest(ctx context.Context) {
	resp, err := invoice.GetInvoice(ctx, "5e058d9af4d38b20d542012f")

	fmt.Println(resp, err)
}

func expireInvoiceWithoutClientTest(ctx context.Context) {
	resp, err := invoice.ExpireInvoice(ctx, "5e058d9af4d38b20d542012f")

	fmt.Println(resp, err)
}

func getAllInvoicesWithoutClientTest(ctx context.Context) {
	data := xendit.GetAllInvoicesParams{
		Statuses:     []string{"EXPIRED", "SETTLED"},
		Limit:        2,
		CreatedAfter: "2016-02-24T23:48:36.697Z",
	}

	resp, err := invoice.GetAllInvoices(ctx, &data)

	fmt.Println(resp, err)
}

/* Wrapper */

func invoiceWithClientTest(ctx context.Context) {
	x := client.New("", "xnd_development_REt02KJzkM6AootfKnDrMw1Sse4LlzEDHfKzXoBocqIEiH4bqjHUJXbl6Cfaab", "", nil)

	createInvoiceWithClientTest(ctx, *x)
	getInvoiceWithClientTest(ctx, *x)
	expireInvoiceWithClientTest(ctx, *x)
	getAllInvoicesWithClientTest(ctx, *x)
}

func invoiceWithoutClientTest(ctx context.Context) {
	xendit.Opt.SecretKey = "xnd_development_REt02KJzkM6AootfKnDrMw1Sse4LlzEDHfKzXoBocqIEiH4bqjHUJXbl6Cfaab"

	createInvoiceWithoutClientTest(ctx)
	getInvoiceWithoutClientTest(ctx)
	expireInvoiceWithoutClientTest(ctx)
	getAllInvoicesWithoutClientTest(ctx)
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	invoiceWithClientTest(ctx)
	invoiceWithoutClientTest(ctx)
}
