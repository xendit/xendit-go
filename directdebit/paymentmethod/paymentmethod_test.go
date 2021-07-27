package paymentmethod

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"
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

	result.(*xendit.PaymentMethod).ID = "pm-ebb1c863-c7b5-4f20-b116-b3071b1d3aef"
	result.(*xendit.PaymentMethod).CustomerID = "4b7b6050-0830-440a-903b-37d527dbbaa9"
	result.(*xendit.PaymentMethod).Type = xendit.DEBIT_CARD
	result.(*xendit.PaymentMethod).Status = "ACTIVE"
	result.(*xendit.PaymentMethod).Created = "2021-07-15T03:17:53.989Z"
	result.(*xendit.PaymentMethod).Updated = "2021-07-15T03:17:53.989Z"
	result.(*xendit.PaymentMethod).Properties = map[string]interface{}{
		"id": "la-55048b41-a7ab-4799-9f33-6ec5cc078db0",
		"currency": "IDR",
		"card_expiry": "11/23",
		"description": "",
		"channel_code": "DC_BRI",
		"card_last_four": "8888",
	}
	result.(*xendit.PaymentMethod).Metadata = map[string]interface{}{
		"meta": "data",
	}

	return nil
}

func TestCreatePaymentMethod(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	requestProperties := map[string]interface{}{
		"id": "la-55048b41-a7ab-4799-9f33-6ec5cc078db0",
	}

	responseProperties := map[string]interface{}{
		"id": "la-55048b41-a7ab-4799-9f33-6ec5cc078db0",
		"currency": "IDR",
		"card_expiry": "11/23",
		"description": "",
		"channel_code": "DC_BRI",
		"card_last_four": "8888",
	}

	metadata := map[string]interface{}{
		"meta": "data",
	}

	testCases := []struct {
		desc        string
		data        *CreatePaymentMethodParams
		expectedRes *xendit.PaymentMethod
		expectedErr *xendit.Error
	}{
		{
			desc: "should create a payment method",
			data: &CreatePaymentMethodParams{
				CustomerID: 	"4b7b6050-0830-440a-903b-37d527dbbaa9",
				Type:			xendit.DEBIT_CARD,
				Properties:		requestProperties,
				Metadata:		metadata,
			},
			expectedRes: &xendit.PaymentMethod{
				ID:				"pm-ebb1c863-c7b5-4f20-b116-b3071b1d3aef",
				CustomerID:  	"4b7b6050-0830-440a-903b-37d527dbbaa9",
				Type:     		xendit.DEBIT_CARD,
				Status: 		"ACTIVE",
				Created: 		"2021-07-15T03:17:53.989Z",
				Updated: 		"2021-07-15T03:17:53.989Z",
				Properties: 	responseProperties,
				Metadata: 		metadata,
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &CreatePaymentMethodParams{
				Type:			xendit.DEBIT_CARD,
				Properties:		requestProperties,
				Metadata:		metadata,
			},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'CustomerID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"POST",
				xendit.Opt.XenditURL+"/payment_methods",
				xendit.Opt.SecretKey,
				&http.Header{},
				tC.data,
				&xendit.PaymentMethod{},
			).Return(nil)

			resp, err := CreatePaymentMethod(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

type apiRequesterMockGet struct {
	mock.Mock
}

func (m *apiRequesterMockGet) Call(ctx context.Context, method string, path string, secretKey string, header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	resultString := `[{
		"id": "pm-ebb1c863-c7b5-4f20-b116-b3071b1d3aef",
		"customer_id": "4b7b6050-0830-440a-903b-37d527dbbaa9",
		"type": "DEBIT_CARD",
		"status": "ACTIVE",
		"properties": {
			"id": "la-55048b41-a7ab-4799-9f33-6ec5cc078db0",
			"currency": "IDR",
			"card_expiry": "11/23",
			"description": "",
			"channel_code": "DC_BRI",
			"card_last_four": "8888"
		},
		"metadata": {
			"meta": "data"
		},
		"created": "2021-07-15T03:17:53.989Z",
		"updated": "2021-07-15T03:17:53.989Z"
	}]`

	_ = json.Unmarshal([]byte(resultString), &result)

	return nil
}

func TestGetPaymentMethodsByCustomerID(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMockGet)
	initTesting(apiRequesterMockObj)

	properties := map[string]interface{}{
		"id": "la-55048b41-a7ab-4799-9f33-6ec5cc078db0",
		"currency": "IDR",
		"card_expiry": "11/23",
		"description": "",
		"channel_code": "DC_BRI",
		"card_last_four": "8888",
	}

	metadata := map[string]interface{}{
		"meta": "data",
	}

	testCases := []struct {
		desc        string
		data        *GetPaymentMethodsByCustomerIDParams
		expectedRes []xendit.PaymentMethod
		expectedErr *xendit.Error
	}{
		{
			desc: "should get payment methods by customer ID",
			data: &GetPaymentMethodsByCustomerIDParams{
				CustomerID: "4b7b6050-0830-440a-903b-37d527dbbaa9",
			},
			expectedRes: []xendit.PaymentMethod{
				{
					ID:				"pm-ebb1c863-c7b5-4f20-b116-b3071b1d3aef",
					CustomerID:  	"4b7b6050-0830-440a-903b-37d527dbbaa9",
					Type:     		xendit.DEBIT_CARD,
					Status: 		"ACTIVE",
					Created: 		"2021-07-15T03:17:53.989Z",
					Updated: 		"2021-07-15T03:17:53.989Z",
					Properties: 	properties,
					Metadata: 		metadata,
				},
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &GetPaymentMethodsByCustomerIDParams{},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'CustomerID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"GET",
				xendit.Opt.XenditURL+"/payment_methods?"+tC.data.QueryString(),
				xendit.Opt.SecretKey,
				nil,
				nil,
				&[]xendit.PaymentMethod{},
			).Return(nil)

			resp, err := GetPaymentMethodsByCustomerID(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
