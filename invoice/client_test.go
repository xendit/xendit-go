package invoice

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xendit/xendit-go"
)

func initTesting() {
	xendit.Opt.SecretKey = "xnd_development_REt02KJzkM6AootfKnDrMw1Sse4LlzEDHfKzXoBocqIEiH4bqjHUJXbl6Cfaab"
}

func getTestingContext() (context.Context, func()) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	return ctx, cancel
}

func TestCreateInvoice(t *testing.T) {
	initTesting()

	ctx, cancel := getTestingContext()
	defer cancel()

	data := xendit.CreateInvoiceParams{
		ExternalID:  "invoice-" + time.Now().String(),
		Amount:      200000,
		PayerEmail:  "customer@customer.com",
		Description: "invoice test #1",
	}

	resp, err := CreateInvoice(ctx, &data)

	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestGetInvoice(t *testing.T) {
	initTesting()

	ctx, cancel := getTestingContext()
	defer cancel()

	resp, err := GetInvoice(ctx, "5e058d9af4d38b20d542012f")

	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestExpireInvoice(t *testing.T) {
	initTesting()

	ctx, cancel := getTestingContext()
	defer cancel()

	resp, err := ExpireInvoice(ctx, "5e058d9af4d38b20d542012f")

	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestGetAllInvoices(t *testing.T) {
	initTesting()

	ctx, cancel := getTestingContext()
	defer cancel()

	data := xendit.GetAllInvoicesParams{
		Statuses:     []string{"EXPIRED", "SETTLED"},
		Limit:        2,
		CreatedAfter: "2016-02-24T23:48:36.697Z",
	}

	resp, err := GetAllInvoices(ctx, &data)

	assert.Nil(t, err)
	assert.NotNil(t, resp)
}
