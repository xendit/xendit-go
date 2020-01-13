package invoice_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

func initTesting(apiRequesterMockObj xendit.APIRequester) {
	xendit.Opt.SecretKey = "xnd_development_REt02KJzkM6AootfKnDrMw1Sse4LlzEDHfKzXoBocqIEiH4bqjHUJXbl6Cfaab"
	xendit.SetAPIRequester(apiRequesterMockObj)
}

func getTestingContext() (context.Context, func()) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	return ctx, cancel
}

type apiRequesterMock struct {
	mock.Mock
}

func (m *apiRequesterMock) Call(ctx context.Context, method string, path string, secretKey string, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, params, result)

	result.(*xendit.Invoice).ID = "123"
	result.(*xendit.Invoice).ExternalID = "invoice-external-id"
	result.(*xendit.Invoice).Amount = 200000
	result.(*xendit.Invoice).PayerEmail = "customer@customer.com"
	result.(*xendit.Invoice).Description = "invoice test #1"

	return nil
}

func TestCreateInvoice(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	expectedResult := &xendit.Invoice{
		ID:          "123",
		ExternalID:  "invoice-external-id",
		Amount:      200000,
		PayerEmail:  "customer@customer.com",
		Description: "invoice test #1",
	}

	data := invoice.CreateParams{
		ExternalID:  "invoice-external-id",
		Amount:      200000,
		PayerEmail:  "customer@customer.com",
		Description: "invoice test #1",
	}

	apiRequesterMockObj.On(
		"Call",
		context.Background(),
		"POST",
		"https://api.xendit.co/v2/invoices",
		xendit.Opt.SecretKey,
		&data,
		&xendit.Invoice{},
	).Return(nil)

	resp, err := invoice.Create(&data)

	apiRequesterMockObj.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, resp)
}

func TestFalseCreateInvoice(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	data := invoice.CreateParams{
		ExternalID: "invoice-external-id",
		Amount:     200000,
	}

	resp, err := invoice.Create(&data)

	apiRequesterMockObj.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestGetInvoice(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	expectedResult := &xendit.Invoice{
		ID:          "123",
		ExternalID:  "invoice-external-id",
		Amount:      200000,
		PayerEmail:  "customer@customer.com",
		Description: "invoice test #1",
	}

	apiRequesterMockObj.On(
		"Call",
		context.Background(),
		"GET",
		"https://api.xendit.co/v2/invoices/123",
		xendit.Opt.SecretKey,
		nil,
		&xendit.Invoice{},
	).Return(nil)

	resp, err := invoice.Get("123")

	apiRequesterMockObj.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, resp)
}

func TestFalseGetInvoice(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	resp, err := invoice.Get("")

	apiRequesterMockObj.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestExpireInvoice(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	expectedResult := &xendit.Invoice{
		ID:          "123",
		ExternalID:  "invoice-external-id",
		Amount:      200000,
		PayerEmail:  "customer@customer.com",
		Description: "invoice test #1",
	}

	apiRequesterMockObj.On(
		"Call",
		context.Background(),
		"POST",
		"https://api.xendit.co/invoices/123/expire!",
		xendit.Opt.SecretKey,
		nil,
		&xendit.Invoice{},
	).Return(nil)

	resp, err := invoice.Expire("123")

	apiRequesterMockObj.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, resp)
}

type apiRequesterGetAllMock struct {
	mock.Mock
}

func (m *apiRequesterGetAllMock) Call(ctx context.Context, method string, path string, secretKey string, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, params, result)

	resultString := `[{
		"id": "123",
		"external_id": "invoice-external-id",
		"amount": 200000,
		"payer_email": "customer@customer.com",
		"description": "invoice test #1"
	}]`

	json.Unmarshal([]byte(resultString), &result)
	return nil
}

func TestGetAllInvoices(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterGetAllMock)
	initTesting(apiRequesterMockObj)

	expectedResult := []xendit.Invoice{
		xendit.Invoice{
			ID:          "123",
			ExternalID:  "invoice-external-id",
			Amount:      200000,
			PayerEmail:  "customer@customer.com",
			Description: "invoice test #1",
		},
	}

	createdAfter, _ := time.Parse(time.RFC3339, "2016-02-24T23:48:36.697Z")
	data := invoice.GetAllParams{
		Statuses:     []string{"EXPIRED", "SETTLED"},
		Limit:        2,
		CreatedAfter: createdAfter,
	}

	apiRequesterMockObj.On(
		"Call",
		context.Background(),
		"GET",
		"https://api.xendit.co/v2/invoices?"+data.QueryString(),
		xendit.Opt.SecretKey,
		&data,
		&[]xendit.Invoice{},
	).Return(nil)

	resp, err := invoice.GetAll(&data)

	apiRequesterMockObj.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, resp)
}
