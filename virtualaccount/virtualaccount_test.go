package virtualaccount_test

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
	"github.com/xendit/xendit-go/utils/validator"
	"github.com/xendit/xendit-go/virtualaccount"
)

func initTesting(apiRequesterMockObj xendit.APIRequester) {
	xendit.Opt.SecretKey = "xnd_development_REt02KJzkM6AootfKnDrMw1Sse4LlzEDHfKzXoBocqIEiH4bqjHUJXbl6Cfaab"
	xendit.SetAPIRequester(apiRequesterMockObj)
}

type apiRequesterMock struct {
	mock.Mock
}

func (m *apiRequesterMock) Call(ctx context.Context, method string, path string, secretKey string, header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	expirationDate, _ := time.Parse(time.RFC3339, "2051-01-19T17:00:00.000Z")

	result.(*xendit.VirtualAccount).IsClosed = false
	result.(*xendit.VirtualAccount).Status = "PENDING"
	result.(*xendit.VirtualAccount).Currency = "IDR"
	result.(*xendit.VirtualAccount).OwnerID = "owner-id"
	result.(*xendit.VirtualAccount).ExternalID = "va-external-id"
	result.(*xendit.VirtualAccount).BankCode = "MANDIRI"
	result.(*xendit.VirtualAccount).MerchantCode = "88888"
	result.(*xendit.VirtualAccount).Name = "Michael Jackson"
	result.(*xendit.VirtualAccount).AccountNumber = "8888888888888"
	result.(*xendit.VirtualAccount).IsSingleUse = false
	result.(*xendit.VirtualAccount).ExpirationDate = &expirationDate
	result.(*xendit.VirtualAccount).ID = "123"

	return nil
}

func TestCreateFixed(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	expirationDate, _ := time.Parse(time.RFC3339, "2051-01-19T17:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *virtualaccount.CreateFixedParams
		expectedRes *xendit.VirtualAccount
		expectedErr *xendit.Error
	}{
		{
			desc: "should create a fixed va",
			data: &virtualaccount.CreateFixedParams{
				ExternalID: "va-external-id",
				BankCode:   "MANDIRI",
				Name:       "Michael Jackson",
			},
			expectedRes: &xendit.VirtualAccount{
				IsClosed:       false,
				Status:         "PENDING",
				Currency:       "IDR",
				OwnerID:        "owner-id",
				ExternalID:     "va-external-id",
				BankCode:       "MANDIRI",
				MerchantCode:   "88888",
				Name:           "Michael Jackson",
				AccountNumber:  "8888888888888",
				IsSingleUse:    false,
				ExpirationDate: &expirationDate,
				ID:             "123",
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &virtualaccount.CreateFixedParams{
				ExternalID: "va-external-id",
				BankCode:   "MANDIRI",
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
				"https://api.xendit.co/callback_virtual_accounts",
				xendit.Opt.SecretKey,
				nil,
				tC.data,
				&xendit.VirtualAccount{},
			).Return(nil)

			resp, err := virtualaccount.CreateFixed(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestGetFixed(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	expirationDate, _ := time.Parse(time.RFC3339, "2051-01-19T17:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *virtualaccount.GetFixedParams
		expectedRes *xendit.VirtualAccount
		expectedErr *xendit.Error
	}{
		{
			desc: "should get a fixed va",
			data: &virtualaccount.GetFixedParams{
				ID: "123",
			},
			expectedRes: &xendit.VirtualAccount{
				IsClosed:       false,
				Status:         "PENDING",
				Currency:       "IDR",
				OwnerID:        "owner-id",
				ExternalID:     "va-external-id",
				BankCode:       "MANDIRI",
				MerchantCode:   "88888",
				Name:           "Michael Jackson",
				AccountNumber:  "8888888888888",
				IsSingleUse:    false,
				ExpirationDate: &expirationDate,
				ID:             "123",
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &virtualaccount.GetFixedParams{},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'ID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"GET",
				"https://api.xendit.co/callback_virtual_accounts/"+tC.data.ID,
				xendit.Opt.SecretKey,
				nil,
				nil,
				&xendit.VirtualAccount{},
			).Return(nil)

			resp, err := virtualaccount.GetFixed(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestUpdateFixed(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	expirationDate, _ := time.Parse(time.RFC3339, "2051-01-19T17:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *virtualaccount.UpdateFixedParams
		expectedRes *xendit.VirtualAccount
		expectedErr *xendit.Error
	}{
		{
			desc: "should update a fixed va",
			data: &virtualaccount.UpdateFixedParams{
				ID:             "123",
				ExpirationDate: &expirationDate,
			},
			expectedRes: &xendit.VirtualAccount{
				IsClosed:       false,
				Status:         "PENDING",
				Currency:       "IDR",
				OwnerID:        "owner-id",
				ExternalID:     "va-external-id",
				BankCode:       "MANDIRI",
				MerchantCode:   "88888",
				Name:           "Michael Jackson",
				AccountNumber:  "8888888888888",
				IsSingleUse:    false,
				ExpirationDate: &expirationDate,
				ID:             "123",
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &virtualaccount.UpdateFixedParams{},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'ID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"PATCH",
				"https://api.xendit.co/callback_virtual_accounts/"+tC.data.ID,
				xendit.Opt.SecretKey,
				nil,
				tC.data,
				&xendit.VirtualAccount{},
			).Return(nil)

			resp, err := virtualaccount.UpdateFixed(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

type apiRequesterGetAvailableBanksMock struct {
	mock.Mock
}

func (m *apiRequesterGetAvailableBanksMock) Call(ctx context.Context, method string, path string, secretKey string, header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	resultString := `[
		{
			"name": "Bank Mandiri",
			"code": "MANDIRI"
		},
		{
			"name": "Bank Negara Indonesia",
			"code": "BNI"
		}
	]`

	_ = json.Unmarshal([]byte(resultString), &result)

	return nil
}

func TestGetAvailableBanks(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterGetAvailableBanksMock)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		expectedRes []xendit.VirtualAccountBank
		expectedErr *xendit.Error
	}{
		{
			desc: "should get a fixed va",
			expectedRes: []xendit.VirtualAccountBank{
				xendit.VirtualAccountBank{
					Name: "Bank Mandiri",
					Code: "MANDIRI",
				},
				xendit.VirtualAccountBank{
					Name: "Bank Negara Indonesia",
					Code: "BNI",
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
				"https://api.xendit.co/available_virtual_account_banks",
				xendit.Opt.SecretKey,
				nil,
				nil,
				&[]xendit.VirtualAccountBank{},
			).Return(nil)

			resp, err := virtualaccount.GetAvailableBanks()

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

type apiRequesterGetPaymentMock struct {
	mock.Mock
}

func (m *apiRequesterGetPaymentMock) Call(ctx context.Context, method string, path string, secretKey string, header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	transactionTimestamp, _ := time.Parse(time.RFC3339, "2020-01-01T17:00:00.000Z")

	result.(*xendit.VirtualAccountPayment).ID = "123"
	result.(*xendit.VirtualAccountPayment).PaymentID = "payment123"
	result.(*xendit.VirtualAccountPayment).CallbackVirtualAccountID = "callback"
	result.(*xendit.VirtualAccountPayment).AccountNumber = "8888888888888"
	result.(*xendit.VirtualAccountPayment).BankCode = "MANDIRI"
	result.(*xendit.VirtualAccountPayment).Amount = 200000
	result.(*xendit.VirtualAccountPayment).TransactionTimestamp = &transactionTimestamp
	result.(*xendit.VirtualAccountPayment).MerchantCode = "88888"
	result.(*xendit.VirtualAccountPayment).Currency = "IDR"

	return nil
}

func TestGetPayment(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterGetPaymentMock)
	initTesting(apiRequesterMockObj)

	transactionTimestamp, _ := time.Parse(time.RFC3339, "2020-01-01T17:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *virtualaccount.GetPaymentParams
		expectedRes *xendit.VirtualAccountPayment
		expectedErr *xendit.Error
	}{
		{
			desc: "should gets a va payment",
			data: &virtualaccount.GetPaymentParams{
				PaymentID: "payment123",
			},
			expectedRes: &xendit.VirtualAccountPayment{
				ID:                       "123",
				PaymentID:                "payment123",
				CallbackVirtualAccountID: "callback",
				AccountNumber:            "8888888888888",
				BankCode:                 "MANDIRI",
				Amount:                   200000,
				TransactionTimestamp:     &transactionTimestamp,
				MerchantCode:             "88888",
				Currency:                 "IDR",
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &virtualaccount.GetPaymentParams{},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'PaymentID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"GET",
				"https://api.xendit.co/callback_virtual_account_payments/payment_id="+tC.data.PaymentID,
				xendit.Opt.SecretKey,
				nil,
				nil,
				&xendit.VirtualAccountPayment{},
			).Return(nil)

			resp, err := virtualaccount.GetPayment(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}