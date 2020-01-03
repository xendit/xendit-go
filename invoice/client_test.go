package invoice

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"
)

type httpRequesterMock struct {
	mock.Mock
}

func (m *httpRequesterMock) Call(ctx context.Context, method string, path string, secretKey string, params interface{}, result interface{}) error {
	args := m.Called(ctx, method, path, secretKey, params, result)

	return args.Error(0)
}

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
	httpRequesterMockObj := new(httpRequesterMock)
	xendit.SetHTTPRequester(httpRequesterMockObj)

	data := xendit.CreateInvoiceParams{
		ExternalID:  "invoice-" + time.Now().String(),
		Amount:      200000,
		PayerEmail:  "customer@customer.com",
		Description: "invoice test #1",
	}

	httpRequesterMockObj.On(
		"Call",
		nil,
		"POST",
		"https://api.xendit.co/v2/invoices",
		xendit.Opt.SecretKey,
		&data,
		&xendit.Invoice{},
	).Return(nil)

	resp, err := CreateInvoice(&data)

	httpRequesterMockObj.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Nil(t, resp)
}

// func TestGetInvoice(t *testing.T) {
// 	initTesting()

// 	ctx, cancel := getTestingContext()
// 	defer cancel()

// 	resp, err := GetInvoice("5e058d9af4d38b20d542012f")

// 	assert.Nil(t, err)
// 	assert.NotNil(t, resp)
// }

// func TestExpireInvoice(t *testing.T) {
// 	initTesting()

// 	ctx, cancel := getTestingContext()
// 	defer cancel()

// 	resp, err := ExpireInvoice("5e058d9af4d38b20d542012f")

// 	assert.Nil(t, err)
// 	assert.NotNil(t, resp)
// }

// func TestGetAllInvoices(t *testing.T) {
// 	initTesting()

// 	ctx, cancel := getTestingContext()
// 	defer cancel()

// 	data := xendit.GetAllInvoicesParams{
// 		Statuses:     []string{"EXPIRED", "SETTLED"},
// 		Limit:        2,
// 		CreatedAfter: "2016-02-24T23:48:36.697Z",
// 	}

// 	resp, err := GetAllInvoices(&data)

// 	assert.Nil(t, err)
// 	assert.NotNil(t, resp)
// }

// func TestInvoiceWithoutSecretKey(t *testing.T) {
// 	ctx, cancel := getTestingContext()
// 	defer cancel()

// 	resp, err := GetInvoice("5e058d9af4d38b20d542012f")

// 	assert.NotNil(t, err)
// 	assert.Nil(t, resp)
// }
