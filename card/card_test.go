package card_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/card"
	"github.com/xendit/xendit-go/utils/validator"
)

func initTesting(apiRequesterMockObj xendit.APIRequester) {
	xendit.Opt.SecretKey = "examplesecretkey"
	xendit.SetAPIRequester(apiRequesterMockObj)
}

/* Charge */

type apiRequesterChargeMock struct {
	mock.Mock
}

func (m *apiRequesterChargeMock) Call(ctx context.Context, method string, path string, secretKey string, header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	created, _ := time.Parse(time.RFC3339, "2020-02-02T00:00:00.000Z")

	result.(*xendit.CardCharge).ID = "123"
	result.(*xendit.CardCharge).Status = "AUTHORIZED"
	result.(*xendit.CardCharge).MerchantID = "example-merchant-id"
	result.(*xendit.CardCharge).Created = &created
	result.(*xendit.CardCharge).BusinessID = "example-business-id"
	result.(*xendit.CardCharge).AuthorizedAmount = 200000
	result.(*xendit.CardCharge).ExternalID = "example-external-id"
	result.(*xendit.CardCharge).MerchantReferenceCode = "123123"
	result.(*xendit.CardCharge).ChargeType = "SINGLE_USE_TOKEN"
	result.(*xendit.CardCharge).CardBrand = "VISA"
	result.(*xendit.CardCharge).MaskedCardNumber = "400000XXXXXX0002"
	result.(*xendit.CardCharge).ECI = "05"
	result.(*xendit.CardCharge).CardType = "CREDIT"
	result.(*xendit.CardCharge).BankReconciliationID = "123123123"
	result.(*xendit.CardCharge).Currency = "IDR"

	return nil
}

func TestCreateCharge(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterChargeMock)
	initTesting(apiRequesterMockObj)

	created, _ := time.Parse(time.RFC3339, "2020-02-02T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *card.CreateChargeParams
		expectedRes *xendit.CardCharge
		expectedErr *xendit.Error
	}{
		{
			desc: "should create a charge",
			data: &card.CreateChargeParams{
				TokenID:    "charge-token-id",
				ExternalID: "charge-external-id",
				Amount:     200000,
			},
			expectedRes: &xendit.CardCharge{
				ID:                    "123",
				Status:                "AUTHORIZED",
				MerchantID:            "example-merchant-id",
				Created:               &created,
				BusinessID:            "example-business-id",
				AuthorizedAmount:      200000,
				ExternalID:            "example-external-id",
				MerchantReferenceCode: "123123",
				ChargeType:            "SINGLE_USE_TOKEN",
				CardBrand:             "VISA",
				MaskedCardNumber:      "400000XXXXXX0002",
				ECI:                   "05",
				CardType:              "CREDIT",
				BankReconciliationID:  "123123123",
				Currency:              "IDR",
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &card.CreateChargeParams{
				ExternalID: "example-external-id",
				Amount:     200000,
			},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'TokenID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"POST",
				xendit.Opt.XenditURL+"/credit_card_charges",
				xendit.Opt.SecretKey,
				nil,
				tC.data,
				&xendit.CardCharge{},
			).Return(nil)

			resp, err := card.CreateCharge(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestGetCharge(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterChargeMock)
	initTesting(apiRequesterMockObj)

	created, _ := time.Parse(time.RFC3339, "2020-02-02T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *card.GetChargeParams
		expectedRes *xendit.CardCharge
		expectedErr *xendit.Error
	}{
		{
			desc: "should get a charge",
			data: &card.GetChargeParams{
				ChargeID: "123",
			},
			expectedRes: &xendit.CardCharge{
				ID:                    "123",
				Status:                "AUTHORIZED",
				MerchantID:            "example-merchant-id",
				Created:               &created,
				BusinessID:            "example-business-id",
				AuthorizedAmount:      200000,
				ExternalID:            "example-external-id",
				MerchantReferenceCode: "123123",
				ChargeType:            "SINGLE_USE_TOKEN",
				CardBrand:             "VISA",
				MaskedCardNumber:      "400000XXXXXX0002",
				ECI:                   "05",
				CardType:              "CREDIT",
				BankReconciliationID:  "123123123",
				Currency:              "IDR",
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &card.GetChargeParams{},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'ChargeID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"GET",
				xendit.Opt.XenditURL+"/credit_card_charges/"+tC.data.ChargeID,
				xendit.Opt.SecretKey,
				nil,
				nil,
				&xendit.CardCharge{},
			).Return(nil)

			resp, err := card.GetCharge(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestCaptureCharge(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterChargeMock)
	initTesting(apiRequesterMockObj)

	created, _ := time.Parse(time.RFC3339, "2020-02-02T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *card.CaptureChargeParams
		expectedRes *xendit.CardCharge
		expectedErr *xendit.Error
	}{
		{
			desc: "should capture a charge",
			data: &card.CaptureChargeParams{
				ChargeID: "123",
				Amount:   200000,
			},
			expectedRes: &xendit.CardCharge{
				ID:                    "123",
				Status:                "AUTHORIZED",
				MerchantID:            "example-merchant-id",
				Created:               &created,
				BusinessID:            "example-business-id",
				AuthorizedAmount:      200000,
				ExternalID:            "example-external-id",
				MerchantReferenceCode: "123123",
				ChargeType:            "SINGLE_USE_TOKEN",
				CardBrand:             "VISA",
				MaskedCardNumber:      "400000XXXXXX0002",
				ECI:                   "05",
				CardType:              "CREDIT",
				BankReconciliationID:  "123123123",
				Currency:              "IDR",
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &card.CaptureChargeParams{
				Amount: 200000,
			},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'ChargeID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"POST",
				xendit.Opt.XenditURL+"/credit_card_charges/"+tC.data.ChargeID+"/capture",
				xendit.Opt.SecretKey,
				nil,
				tC.data,
				&xendit.CardCharge{},
			).Return(nil)

			resp, err := card.CaptureCharge(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

/* Refund */

type apiRequesterRefundMock struct {
	mock.Mock
}

func (m *apiRequesterRefundMock) Call(ctx context.Context, method string, path string, secretKey string, header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, header, params, result)

	created, _ := time.Parse(time.RFC3339, "2020-02-02T00:00:00.000Z")

	result.(*xendit.CardRefund).ID = "1234"
	result.(*xendit.CardRefund).Updated = &created
	result.(*xendit.CardRefund).Created = &created
	result.(*xendit.CardRefund).CreditCardChargeID = "123"
	result.(*xendit.CardRefund).UserID = "123123123"
	result.(*xendit.CardRefund).Amount = 10000
	result.(*xendit.CardRefund).ExternalID = "example-external-id"
	result.(*xendit.CardRefund).Currency = "IDR"
	result.(*xendit.CardRefund).Status = "AUTHORIZED"
	result.(*xendit.CardRefund).FeeRefundAmount = 100

	return nil
}

func TestCreateRefund(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterRefundMock)
	initTesting(apiRequesterMockObj)

	created, _ := time.Parse(time.RFC3339, "2020-02-02T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *card.CreateRefundParams
		expectedRes *xendit.CardRefund
		expectedErr *xendit.Error
	}{
		{
			desc: "should create a refund",
			data: &card.CreateRefundParams{
				IdempotencyKey: "unique-idempotency-key",
				ChargeID:       "123",
				Amount:         10000,
				ExternalID:     "example-external-id",
			},
			expectedRes: &xendit.CardRefund{
				ID:                 "1234",
				Updated:            &created,
				Created:            &created,
				CreditCardChargeID: "123",
				UserID:             "123123123",
				Amount:             10000,
				ExternalID:         "example-external-id",
				Currency:           "IDR",
				Status:             "AUTHORIZED",
				FeeRefundAmount:    100,
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &card.CreateRefundParams{
				Amount:     10000,
				ExternalID: "example-external-id",
			},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'ChargeID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"POST",
				xendit.Opt.XenditURL+"/credit_card_charges/"+tC.data.ChargeID+"/refunds",
				xendit.Opt.SecretKey,
				mock.Anything,
				tC.data,
				&xendit.CardRefund{},
			).Return(nil)

			resp, err := card.CreateRefund(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

/* Authorization */

type apiRequesterReverseAuthorizationMock struct {
	mock.Mock
}

func (m *apiRequesterReverseAuthorizationMock) Call(ctx context.Context, method string, path string, secretKey string, header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	created, _ := time.Parse(time.RFC3339, "2020-02-02T00:00:00.000Z")

	result.(*xendit.CardReverseAuthorization).ID = "123123"
	result.(*xendit.CardReverseAuthorization).ExternalID = "example-external-id"
	result.(*xendit.CardReverseAuthorization).CreditCardChargeID = "123"
	result.(*xendit.CardReverseAuthorization).BusinessID = "123123123"
	result.(*xendit.CardReverseAuthorization).Amount = 10000
	result.(*xendit.CardReverseAuthorization).Status = "PENDING"
	result.(*xendit.CardReverseAuthorization).Created = &created
	result.(*xendit.CardReverseAuthorization).Currency = "IDR"

	return nil
}

func TestReverseAuthorization(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterReverseAuthorizationMock)
	initTesting(apiRequesterMockObj)

	created, _ := time.Parse(time.RFC3339, "2020-02-02T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *card.ReverseAuthorizationParams
		expectedRes *xendit.CardReverseAuthorization
		expectedErr *xendit.Error
	}{
		{
			desc: "should reverse authorization",
			data: &card.ReverseAuthorizationParams{
				ChargeID:   "123",
				ExternalID: "example-external-id",
			},
			expectedRes: &xendit.CardReverseAuthorization{
				ID:                 "123123",
				ExternalID:         "example-external-id",
				CreditCardChargeID: "123",
				BusinessID:         "123123123",
				Amount:             10000,
				Status:             "PENDING",
				Created:            &created,
				Currency:           "IDR",
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &card.ReverseAuthorizationParams{
				ExternalID: "example-external-id",
			},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'ChargeID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"POST",
				xendit.Opt.XenditURL+"/credit_card_charges/"+tC.data.ChargeID+"/auth_reversal",
				xendit.Opt.SecretKey,
				nil,
				tC.data,
				&xendit.CardReverseAuthorization{},
			).Return(nil)

			resp, err := card.ReverseAuthorization(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
