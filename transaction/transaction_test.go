package transaction_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/transaction"
	"github.com/xendit/xendit-go/utils/validator"
)

func initTesting(apiRequesterMockObj xendit.APIRequester) {
	xendit.Opt.SecretKey = "examplesecretkey"
	xendit.SetAPIRequester(apiRequesterMockObj)
}

type apiRequesterMock struct {
	mock.Mock
}

func (m *apiRequesterMock) Call(ctx context.Context, method string, path string, secretKey string, header http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	resultString := `{
		"id": "txn_13dd178d-41fa-40b7-8fd3-f83675d6f413",
		"product_id": "d290f1ee-6c54-4b01-90e6-d701748f0701",
		"type": "PAYMENT",
		"channel_category": "RETAIL_OUTLET",
		"channel_code": "ALFAMART",
		"reference_id": "ref23232",
		"account_identifier": null,
		"currency": "IDR",
		"amount": 500000,
		"cashflow": "MONEY_IN",
		"status": "SUCCESS",
		"business_id": "5fc9f5b246f820517e38c84d",
		"created": "2021-06-23T02:42:15.601Z",
		"updated": "2021-06-23T02:42:15.601Z",
		"fee":{
			"xendit_fee": 1500,
			"value_added_tax": 500,
			"xendit_withholding_tax": 0,
			"third_party_withholding_tax": 0,
			"status": "COMPLETED"
		},
		"settlement_status": "PENDING",
		"estimated_settlement_time": "2021-06-23T02:42:15.601Z"
	}`

	_ = json.Unmarshal([]byte(resultString), &result)

	return nil
}

func TestGetTransaction(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	created := time.Date(2021, 6, 23, 2, 42, 15, 601000000, time.UTC)
	updated := time.Date(2021, 6, 23, 2, 42, 15, 601000000, time.UTC)
	estimatedSettlementTime := time.Date(2021, 6, 23, 2, 42, 15, 601000000, time.UTC)

	testCases := []struct {
		desc        string
		data        *transaction.GetTransactionParams
		expectedRes *xendit.Transaction
		expectedErr *xendit.Error
	}{
		{
			desc: "should gets a single transaction",
			data: &transaction.GetTransactionParams{
				TransactionID: "txn_13dd178d-41fa-40b7-8fd3-f83675d6f413",
			},
			expectedRes: &xendit.Transaction{
				ID:                "txn_13dd178d-41fa-40b7-8fd3-f83675d6f413",
				ProductID:         "d290f1ee-6c54-4b01-90e6-d701748f0701",
				Type:              "PAYMENT",
				ChannelCategory:   "RETAIL_OUTLET",
				ChannelCode:       "ALFAMART",
				ReferenceID:       "ref23232",
				AccountIdentifier: "",
				Currency:          "IDR",
				Amount:            500000,
				Cashflow:          "MONEY_IN",
				Status:            "SUCCESS",
				BusinessID:        "5fc9f5b246f820517e38c84d",
				Created:           &created,
				Updated:           &updated,
				Fee: xendit.TransactionFee{
					XenditFee:                1500,
					ValueAddedTax:            500,
					XenditWithholdingTax:     0,
					ThirdPartyWithholdingTax: 0,
					Status:                   "COMPLETED",
				},
				SettlementStatus:        "PENDING",
				EstimatedSettlementTime: &estimatedSettlementTime,
			},
			expectedErr: nil,
		},

		{
			desc:        "should report missing required fields",
			data:        &transaction.GetTransactionParams{},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'TransactionID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"GET",
				xendit.Opt.XenditURL+"/transactions/"+tC.data.TransactionID,
				xendit.Opt.SecretKey,
				nil,
				nil,
				&xendit.Transaction{},
			).Return(nil)

			resp, err := transaction.GetTransaction(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

type apiRequesterGetListMock struct {
	mock.Mock
}

func (m *apiRequesterGetListMock) Call(ctx context.Context, method string, path string, secretKey string, header http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	resultString := `{
		"has_more": true,
		"data": [
			{
				"id": "txn_13dd178d-41fa-40b7-8fd3-f83675d6f413",
				"product_id": "d290f1ee-6c54-4b01-90e6-d701748f0701",
				"type": "PAYMENT",
				"status": "SUCCESS",
				"channel_category": "RETAIL_OUTLET",
				"channel_code": "ALFAMART",
				"reference_id": "ref23244",
				"account_identifier": null,
				"currency": "IDR",
				"amount": 1,
				"cashflow": "MONEY_IN",
				"business_id": "5fc9f5b246f820517e38c84d",
				"created": "2021-06-23T02:42:15.601Z",
				"updated": "2021-06-23T02:42:15.601Z",
				"settlement_status": "PENDING",
				"estimated_settlement_time": "2021-06-23T02:42:15.601Z"
			},
			{
				"id": "txn_a765a3f0-34c0-41ee-8686-bca11835ebdc",
				"product_id": "d290f1ee-6c54-4b01-90e6-d701748f0700",
				"type": "PAYMENT",
				"status": "SUCCESS",
				"channel_category": "RETAIL_OUTLET",
				"channel_code": "ALFAMART",
				"reference_id": "ref242424",
				"account_identifier": null,
				"currency": "IDR",
				"amount": 1,
				"cashflow": "MONEY_IN",
				"business_id": "5fc9f5b246f820517e38c84d",
				"created": "2021-06-23T02:39:23.176Z",
				"updated": "2021-06-23T02:39:23.176Z",
				"settlement_status": "PENDING",
				"estimated_settlement_time": "2021-06-23T02:42:15.601Z"
			}
		],
		"links": [
			{
				"href": "/transactions?types=PAYMENT&statuses=SUCCESS&channel_categories=EWALLET&channel_categories=RETAIL_OUTLET&limit=2&after_id=txn_a765a3f0-34c0-41ee-8686-bca11835ebdc",
				"method": "GET",
				"rel": "next"
			}
		]
	}`

	_ = json.Unmarshal([]byte(resultString), &result)

	return nil
}

func TestGetListTransaction(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterGetListMock)
	initTesting(apiRequesterMockObj)

	created := time.Date(2021, 6, 23, 2, 42, 15, 601000000, time.UTC)
	updated := time.Date(2021, 6, 23, 2, 42, 15, 601000000, time.UTC)
	estimatedSettlementTime := time.Date(2021, 6, 23, 2, 42, 15, 601000000, time.UTC)

	created2 := time.Date(2021, 6, 23, 2, 39, 23, 176000000, time.UTC)
	updated2 := time.Date(2021, 6, 23, 2, 39, 23, 176000000, time.UTC)
	estimatedSettlementTime2 := time.Date(2021, 6, 23, 2, 42, 15, 601000000, time.UTC)

	testCases := []struct {
		desc        string
		data        *transaction.GetListTransactionParams
		expectedRes *xendit.ListTransactions
		expectedErr *xendit.Error
	}{
		{
			desc: "should get a list of transaction",
			data: &transaction.GetListTransactionParams{},
			expectedRes: &xendit.ListTransactions{
				HasMore: true,
				Data: []xendit.Transaction{
					{
						ID:                      "txn_13dd178d-41fa-40b7-8fd3-f83675d6f413",
						ProductID:               "d290f1ee-6c54-4b01-90e6-d701748f0701",
						Type:                    "PAYMENT",
						Status:                  "SUCCESS",
						ChannelCategory:         "RETAIL_OUTLET",
						ChannelCode:             "ALFAMART",
						ReferenceID:             "ref23244",
						AccountIdentifier:       "",
						Currency:                "IDR",
						Amount:                  1,
						Cashflow:                "MONEY_IN",
						BusinessID:              "5fc9f5b246f820517e38c84d",
						Created:                 &created,
						Updated:                 &updated,
						SettlementStatus:        "PENDING",
						EstimatedSettlementTime: &estimatedSettlementTime,
					},
					{
						ID:                      "txn_a765a3f0-34c0-41ee-8686-bca11835ebdc",
						ProductID:               "d290f1ee-6c54-4b01-90e6-d701748f0700",
						Type:                    "PAYMENT",
						Status:                  "SUCCESS",
						ChannelCategory:         "RETAIL_OUTLET",
						ChannelCode:             "ALFAMART",
						ReferenceID:             "ref242424",
						AccountIdentifier:       "",
						Currency:                "IDR",
						Amount:                  1,
						Cashflow:                "MONEY_IN",
						BusinessID:              "5fc9f5b246f820517e38c84d",
						Created:                 &created2,
						Updated:                 &updated2,
						SettlementStatus:        "PENDING",
						EstimatedSettlementTime: &estimatedSettlementTime2,
					},
				},
				Links: []xendit.ListTransactionsLinks{
					{
						Href:   "/transactions?types=PAYMENT&statuses=SUCCESS&channel_categories=EWALLET&channel_categories=RETAIL_OUTLET&limit=2&after_id=txn_a765a3f0-34c0-41ee-8686-bca11835ebdc",
						Method: "GET",
						Rel:    "next",
					},
				},
			},
			expectedErr: nil,
		},
	}

	for _, tC := range testCases {

		qs, _ := query.Values(tC.data)

		t.Run(tC.desc, func(t *testing.T) {

			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"GET",
				xendit.Opt.XenditURL+"/transactions?"+qs.Encode(),
				xendit.Opt.SecretKey,
				nil,
				nil,
				&xendit.ListTransactions{},
			).Return(nil)

			resp, err := transaction.GetListTransaction(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
