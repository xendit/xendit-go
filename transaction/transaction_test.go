package transaction_test

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
	"github.com/xendit/xendit-go/retailoutlet"
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
	m.Called(ctx, method, path, secretKey, nil, params, result)

	rawJsonString := `{
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
		}
	}`

	var transaction xendit.Transaction

	json.Unmarshal([]byte(rawJsonString), &transaction)

	result = transaction

	return nil
}

func TestGetTransaction(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	expirationDate, _ := time.Parse(time.RFC3339, "2050-01-01T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *retailoutlet.GetFixedPaymentCodeParams
		expectedRes *xendit.RetailOutlet
		expectedErr *xendit.Error
	}{
		{
			desc: "should gets a retail outlet fixed payment code",
			data: &retailoutlet.GetFixedPaymentCodeParams{
				FixedPaymentCodeID: "123",
			},
			expectedRes: &xendit.RetailOutlet{
				IsSingleUse:      false,
				Status:           "ACTIVE",
				OwnerID:          "someone-owner-id",
				ExternalID:       "retailoutlet-external-id",
				RetailOutletName: xendit.RetailOutletNameAlfamart,
				Prefix:           "TEST",
				Name:             "Michael Jackson",
				PaymentCode:      "TEST123",
				Type:             "USER",
				ExpectedAmount:   200000,
				ExpirationDate:   &expirationDate,
				ID:               "123",
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &retailoutlet.GetFixedPaymentCodeParams{},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'FixedPaymentCodeID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"GET",
				xendit.Opt.XenditURL+"/fixed_payment_code/"+tC.data.FixedPaymentCodeID,
				xendit.Opt.SecretKey,
				nil,
				nil,
				&xendit.RetailOutlet{},
			).Return(nil)

			resp, err := retailoutlet.GetFixedPaymentCode(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestUpdateFixedPaymentCode(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	expirationDate, _ := time.Parse(time.RFC3339, "2050-01-01T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *retailoutlet.UpdateFixedPaymentCodeParams
		expectedRes *xendit.RetailOutlet
		expectedErr *xendit.Error
	}{
		{
			desc: "should update a retail outlet fixed payment code",
			data: &retailoutlet.UpdateFixedPaymentCodeParams{
				FixedPaymentCodeID: "123",
				ExpirationDate:     &expirationDate,
				Name:               "Michael Jackson",
				ExpectedAmount:     200000,
			},
			expectedRes: &xendit.RetailOutlet{
				IsSingleUse:      false,
				Status:           "ACTIVE",
				OwnerID:          "someone-owner-id",
				ExternalID:       "retailoutlet-external-id",
				RetailOutletName: xendit.RetailOutletNameAlfamart,
				Prefix:           "TEST",
				Name:             "Michael Jackson",
				PaymentCode:      "TEST123",
				Type:             "USER",
				ExpectedAmount:   200000,
				ExpirationDate:   &expirationDate,
				ID:               "123",
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &retailoutlet.UpdateFixedPaymentCodeParams{
				ExpectedAmount: 200000,
			},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'FixedPaymentCodeID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"PATCH",
				xendit.Opt.XenditURL+"/fixed_payment_code/"+tC.data.FixedPaymentCodeID,
				xendit.Opt.SecretKey,
				nil,
				tC.data,
				&xendit.RetailOutlet{},
			).Return(nil)

			resp, err := retailoutlet.UpdateFixedPaymentCode(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
