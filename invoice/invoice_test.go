package invoice_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
	"github.com/xendit/xendit-go/utils/validator"
)

func initTesting(apiRequesterMockObj xendit.APIRequester) {
	xendit.Opt.SecretKey = "examplesecretkey"
	xendit.SetAPIRequester(apiRequesterMockObj)
}

type apiRequesterMock struct {
	mock.Mock
}

func (m *apiRequesterMock) Call(ctx context.Context, method string, path string, secretKey string, header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, header, params, result)

	result.(*xendit.Invoice).ID = "123"
	result.(*xendit.Invoice).ExternalID = "invoice-external-id"
	result.(*xendit.Invoice).Amount = 200000
	result.(*xendit.Invoice).PayerEmail = "customer@customer.com"
	result.(*xendit.Invoice).Description = "invoice test #1"

	return nil
}

func TestCreate(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *invoice.CreateParams
		expectedRes *xendit.Invoice
		expectedErr *xendit.Error
	}{
		{
			desc: "should create an invoice",
			data: &invoice.CreateParams{
				ExternalID:  "invoice-external-id",
				Amount:      200000,
				PayerEmail:  "customer@customer.com",
				Description: "invoice test #1",
			},
			expectedRes: &xendit.Invoice{
				ID:          "123",
				ExternalID:  "invoice-external-id",
				Amount:      200000,
				PayerEmail:  "customer@customer.com",
				Description: "invoice test #1",
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &invoice.CreateParams{
				ExternalID: "invoice-external-id",
				Amount:     200000,
			},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'PayerEmail', 'Description'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"POST",
				xendit.Opt.XenditURL+"/v2/invoices",
				xendit.Opt.SecretKey,
				&http.Header{},
				tC.data,
				&xendit.Invoice{},
			).Return(nil)

			resp, err := invoice.Create(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestGet(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *invoice.GetParams
		expectedRes *xendit.Invoice
		expectedErr *xendit.Error
	}{
		{
			desc: "should get an invoice",
			data: &invoice.GetParams{
				ID: "123",
			},
			expectedRes: &xendit.Invoice{
				ID:          "123",
				ExternalID:  "invoice-external-id",
				Amount:      200000,
				PayerEmail:  "customer@customer.com",
				Description: "invoice test #1",
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &invoice.GetParams{},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'ID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var header *http.Header

			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"GET",
				xendit.Opt.XenditURL+"/v2/invoices/123",
				xendit.Opt.SecretKey,
				header,
				nil,
				&xendit.Invoice{},
			).Return(nil)

			resp, err := invoice.Get(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestExpire(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *invoice.ExpireParams
		expectedRes *xendit.Invoice
		expectedErr *xendit.Error
	}{
		{
			desc: "should expire an invoice",
			data: &invoice.ExpireParams{
				ID: "123",
			},
			expectedRes: &xendit.Invoice{
				ID:          "123",
				ExternalID:  "invoice-external-id",
				Amount:      200000,
				PayerEmail:  "customer@customer.com",
				Description: "invoice test #1",
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &invoice.ExpireParams{},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'ID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var header *http.Header

			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"POST",
				xendit.Opt.XenditURL+"/invoices/123/expire!",
				xendit.Opt.SecretKey,
				header,
				nil,
				&xendit.Invoice{},
			).Return(nil)

			resp, err := invoice.Expire(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

type apiRequesterGetAllMock struct {
	mock.Mock
}

func (m *apiRequesterGetAllMock) Call(ctx context.Context, method string, path string, secretKey string, header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	resultString := `[{
		"id": "123",
		"external_id": "invoice-external-id",
		"amount": 200000,
		"payer_email": "customer@customer.com",
		"description": "invoice test #1"
	}]`

	_ = json.Unmarshal([]byte(resultString), &result)

	return nil
}

func TestGetAll(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterGetAllMock)
	initTesting(apiRequesterMockObj)

	createdAfter, _ := time.Parse(time.RFC3339, "2016-02-24T23:48:36.697Z")

	testCases := []struct {
		desc        string
		data        *invoice.GetAllParams
		expectedRes []xendit.Invoice
		expectedErr *xendit.Error
	}{
		{
			desc: "should get a list of invoices",
			data: &invoice.GetAllParams{},
			expectedRes: []xendit.Invoice{
				xendit.Invoice{
					ID:          "123",
					ExternalID:  "invoice-external-id",
					Amount:      200000,
					PayerEmail:  "customer@customer.com",
					Description: "invoice test #1",
				},
			},
			expectedErr: nil,
		},
		{
			desc: "should get a list of invoices",
			data: &invoice.GetAllParams{
				Statuses:     []string{"EXPIRED", "SETTLED"},
				Limit:        2,
				CreatedAfter: createdAfter,
			},
			expectedRes: []xendit.Invoice{
				xendit.Invoice{
					ID:          "123",
					ExternalID:  "invoice-external-id",
					Amount:      200000,
					PayerEmail:  "customer@customer.com",
					Description: "invoice test #1",
				},
			},
			expectedErr: nil,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"GET",
				xendit.Opt.XenditURL+"/v2/invoices?"+tC.data.QueryString(),
				xendit.Opt.SecretKey,
				nil,
				nil,
				&[]xendit.Invoice{},
			).Return(nil)

			resp, err := invoice.GetAll(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
