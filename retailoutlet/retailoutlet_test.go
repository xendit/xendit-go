package retailoutlet_test

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

func (m *apiRequesterMock) Call(ctx context.Context, method string, path string, secretKey string, header http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	expirationDate, _ := time.Parse(time.RFC3339, "2050-01-01T00:00:00.000Z")

	result.(*xendit.RetailOutlet).IsSingleUse = false
	result.(*xendit.RetailOutlet).Status = "ACTIVE"
	result.(*xendit.RetailOutlet).OwnerID = "someone-owner-id"
	result.(*xendit.RetailOutlet).ExternalID = "retailoutlet-external-id"
	result.(*xendit.RetailOutlet).RetailOutletName = xendit.RetailOutletNameAlfamart
	result.(*xendit.RetailOutlet).Prefix = "TEST"
	result.(*xendit.RetailOutlet).Name = "Michael Jackson"
	result.(*xendit.RetailOutlet).PaymentCode = "TEST123"
	result.(*xendit.RetailOutlet).Type = "USER"
	result.(*xendit.RetailOutlet).ExpectedAmount = 200000
	result.(*xendit.RetailOutlet).ExpirationDate = &expirationDate
	result.(*xendit.RetailOutlet).ID = "123"

	return nil
}

func TestCreateFixedPaymentCode(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	expirationDate, _ := time.Parse(time.RFC3339, "2050-01-01T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *retailoutlet.CreateFixedPaymentCodeParams
		expectedRes *xendit.RetailOutlet
		expectedErr *xendit.Error
	}{
		{
			desc: "should create a retail outlet fixed payment code",
			data: &retailoutlet.CreateFixedPaymentCodeParams{
				ExternalID:       "retailoutlet-external-id",
				RetailOutletName: xendit.RetailOutletNameAlfamart,
				Name:             "Michael Jackson",
				ExpectedAmount:   200000,
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
			data: &retailoutlet.CreateFixedPaymentCodeParams{
				ExternalID:       "retailoutlet-external-id",
				RetailOutletName: xendit.RetailOutletNameAlfamart,
				ExpectedAmount:   200000,
			},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'Name'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"POST",
				xendit.Opt.XenditURL+"/fixed_payment_code",
				xendit.Opt.SecretKey,
				nil,
				tC.data,
				&xendit.RetailOutlet{},
			).Return(nil)

			resp, err := retailoutlet.CreateFixedPaymentCode(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestGetFixedPaymentCode(t *testing.T) {
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

type apiRequesterGetPaymentByFixedPaymentCodeMock struct {
	mock.Mock
}

func (m *apiRequesterGetPaymentByFixedPaymentCodeMock) Call(ctx context.Context, method string, path string, secretKey string, header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	resultString := `{
		"data": [
			{
				"status": "COMPLETED",
				"fixed_payment_code_payment_id": "61c53c4fdc1b825d9a58ff54",
				"fixed_payment_code_id": "61c53c3727c7a679826dd90a",
				"amount": 2500000,
				"name": "JOHN DOE",
				"prefix": "TEST",
				"payment_code": "TEST892185",
				"payment_id": "1640315983260",
				"external_id": "FPC-1640315959",
				"retail_outlet_name": "ALFAMART",
				"transaction_timestamp": "2021-12-24T03:19:43.260Z",
				"id": "61c53c4f6cc577e4038ab099",
				"owner_id": "60ca10b83ffd534ece8aa856"
			}
		],
		"has_more": true,
		"links": {
			"href": "https://api.xendit.co/fixed_payment_code/61c53c3727c7a679826dd90a/payments?limit=1&after_id=61c53c4f6cc577e4038ab099",
			"rel": "next",
			"method": "GET"
		}
	}`

	_ = json.Unmarshal([]byte(resultString), &result)

	return nil
}

func TestGetListTransaction(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterGetPaymentByFixedPaymentCodeMock)
	initTesting(apiRequesterMockObj)

	transactionTimestamp := time.Date(2021, 12, 24, 3, 19, 43, 260000000, time.UTC)

	testCases := []struct {
		desc        string
		data        *retailoutlet.GetPaymentByFixedPaymentCodeParams
		expectedRes *xendit.RetailOutletPayments
		expectedErr *xendit.Error
	}{
		{
			desc: "should get a list of transaction",
			data: &retailoutlet.GetPaymentByFixedPaymentCodeParams{
				FixedPaymentCodeID: "61c53c4fdc1b825d9a58ff54",
			},
			expectedRes: &xendit.RetailOutletPayments{
				Data: []xendit.RetailOutletPayment{
					{
						Status:                    "COMPLETED",
						FixedPaymentCodePaymentID: "61c53c4fdc1b825d9a58ff54",
						FixedPaymentCodeID:        "61c53c3727c7a679826dd90a",
						Amount:                    2500000,
						Name:                      "JOHN DOE",
						Prefix:                    "TEST",
						PaymentCode:               "TEST892185",
						PaymentID:                 "1640315983260",
						ExternalID:                "FPC-1640315959",
						RetailOutletName:          "ALFAMART",
						TransactionTimestamp:      &transactionTimestamp,
						ID:                        "61c53c4f6cc577e4038ab099",
						OwnerID:                   "60ca10b83ffd534ece8aa856",
					},
				},
				HasMore: true,
				Links: xendit.RetailOutletPaymentsLinks{
					Href:   "https://api.xendit.co/fixed_payment_code/61c53c3727c7a679826dd90a/payments?limit=1&after_id=61c53c4f6cc577e4038ab099",
					Rel:    "next",
					Method: "GET",
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
				xendit.Opt.XenditURL+"/fixed_payment_code/"+tC.data.FixedPaymentCodeID+"/payments",
				xendit.Opt.SecretKey,
				nil,
				nil,
				&xendit.RetailOutletPayments{},
			).Return(nil)

			resp, err := retailoutlet.GetPaymentByFixedPaymentCode(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
