package invoice

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"
)

func initTesting(xdAPIRequesterMockObj xendit.XdAPIRequester) {
	xendit.Opt.SecretKey = "xnd_development_REt02KJzkM6AootfKnDrMw1Sse4LlzEDHfKzXoBocqIEiH4bqjHUJXbl6Cfaab"
	xendit.SetXdAPIRequester(xdAPIRequesterMockObj)
}

func getTestingContext() (context.Context, func()) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	return ctx, cancel
}

type xdAPIRequesterMock struct {
	mock.Mock
}

func (m *xdAPIRequesterMock) Call(ctx context.Context, method string, path string, secretKey string, params interface{}, result interface{}) error {
	args := m.Called(ctx, method, path, secretKey, params, result)

	result.(*xendit.Invoice).ID = "123"
	result.(*xendit.Invoice).ExternalID = "invoice-external-id"
	result.(*xendit.Invoice).Amount = 200000
	result.(*xendit.Invoice).PayerEmail = "customer@customer.com"
	result.(*xendit.Invoice).Description = "invoice test #1"

	return args.Error(0)
}

func TestCreateInvoice(t *testing.T) {
	xdAPIRequesterMockObj := new(xdAPIRequesterMock)
	initTesting(xdAPIRequesterMockObj)

	expectedResult := &xendit.Invoice{
		ID:          "123",
		ExternalID:  "invoice-external-id",
		Amount:      200000,
		PayerEmail:  "customer@customer.com",
		Description: "invoice test #1",
	}

	data := xendit.CreateInvoiceParams{
		ExternalID:  "invoice-external-id",
		Amount:      200000,
		PayerEmail:  "customer@customer.com",
		Description: "invoice test #1",
	}

	xdAPIRequesterMockObj.On(
		"Call",
		nil,
		"POST",
		"https://api.xendit.co/v2/invoices",
		xendit.Opt.SecretKey,
		&data,
		&xendit.Invoice{},
	).Return(nil)

	resp, err := CreateInvoice(&data)

	xdAPIRequesterMockObj.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, resp)
}

func TestFalseCreateInvoice(t *testing.T) {
	xdAPIRequesterMockObj := new(xdAPIRequesterMock)
	initTesting(xdAPIRequesterMockObj)

	data := xendit.CreateInvoiceParams{
		ExternalID: "invoice-external-id",
		Amount:     200000,
	}

	resp, err := CreateInvoice(&data)

	xdAPIRequesterMockObj.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestGetInvoice(t *testing.T) {
	xdAPIRequesterMockObj := new(xdAPIRequesterMock)
	initTesting(xdAPIRequesterMockObj)

	expectedResult := &xendit.Invoice{
		ID:          "123",
		ExternalID:  "invoice-external-id",
		Amount:      200000,
		PayerEmail:  "customer@customer.com",
		Description: "invoice test #1",
	}

	xdAPIRequesterMockObj.On(
		"Call",
		nil,
		"GET",
		"https://api.xendit.co/v2/invoices/123",
		xendit.Opt.SecretKey,
		nil,
		&xendit.Invoice{},
	).Return(nil)

	resp, err := GetInvoice("123")

	xdAPIRequesterMockObj.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, resp)
}

func TestFalseGetInvoice(t *testing.T) {
	xdAPIRequesterMockObj := new(xdAPIRequesterMock)
	initTesting(xdAPIRequesterMockObj)

	resp, err := GetInvoice("")

	xdAPIRequesterMockObj.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestExpireInvoice(t *testing.T) {
	xdAPIRequesterMockObj := new(xdAPIRequesterMock)
	initTesting(xdAPIRequesterMockObj)

	expectedResult := &xendit.Invoice{
		ID:          "123",
		ExternalID:  "invoice-external-id",
		Amount:      200000,
		PayerEmail:  "customer@customer.com",
		Description: "invoice test #1",
	}

	xdAPIRequesterMockObj.On(
		"Call",
		nil,
		"POST",
		"https://api.xendit.co/invoices/123/expire!",
		xendit.Opt.SecretKey,
		nil,
		&xendit.Invoice{},
	).Return(nil)

	resp, err := ExpireInvoice("123")

	xdAPIRequesterMockObj.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, resp)
}

func TestFalseExpireInvoice(t *testing.T) {
	xdAPIRequesterMockObj := new(xdAPIRequesterMock)
	initTesting(xdAPIRequesterMockObj)

	resp, err := ExpireInvoice("123")

	xdAPIRequesterMockObj.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

type xdAPIRequesterGetAllMock struct {
	mock.Mock
}

func (m *xdAPIRequesterGetAllMock) Call(ctx context.Context, method string, path string, secretKey string, params interface{}, result interface{}) error {
	args := m.Called(ctx, method, path, secretKey, params, result)

	resultString := `[{
		"id": "123",
		"external_id": "invoice-external-id",
		"amount": 200000,
		"payer_email": "customer@customer.com",
		"description": "invoice test #1"
	}]`

	if err := json.Unmarshal([]byte(resultString), &result); err != nil {
		return err
	}

	return args.Error(0)
}

func TestGetAllInvoices(t *testing.T) {
	xdAPIRequesterMockObj := new(xdAPIRequesterGetAllMock)
	initTesting(xdAPIRequesterMockObj)

	expectedResult := []xendit.Invoice{
		xendit.Invoice{
			ID:          "123",
			ExternalID:  "invoice-external-id",
			Amount:      200000,
			PayerEmail:  "customer@customer.com",
			Description: "invoice test #1",
		},
	}

	data := xendit.GetAllInvoicesParams{
		Statuses:     []string{"EXPIRED", "SETTLED"},
		Limit:        2,
		CreatedAfter: "2016-02-24T23:48:36.697Z",
	}

	xdAPIRequesterMockObj.On(
		"Call",
		nil,
		"GET",
		"https://api.xendit.co/v2/invoices",
		xendit.Opt.SecretKey,
		&data,
		&[]xendit.Invoice{},
	).Return(nil)

	resp, err := GetAllInvoices(&data)

	xdAPIRequesterMockObj.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, resp)
}
