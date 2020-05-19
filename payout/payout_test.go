package payout_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/payout"
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

	date, _ := time.Parse(time.RFC3339, "2050-01-01T00:00:00.000Z")

	result.(*xendit.Payout).ID = "123"
	result.(*xendit.Payout).ExternalID = "payout-external-id"
	result.(*xendit.Payout).Amount = 200000
	result.(*xendit.Payout).Status = "ISSUED"
	result.(*xendit.Payout).Created = &date
	result.(*xendit.Payout).ExpirationTimestamp = &date
	result.(*xendit.Payout).MerchantName = "Business Name"
	result.(*xendit.Payout).PayoutURL = "https://payout-url.com/123"

	return nil
}

func TestCreatePayout(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	date, _ := time.Parse(time.RFC3339, "2050-01-01T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *payout.CreateParams
		expectedRes *xendit.Payout
		expectedErr *xendit.Error
	}{
		{
			desc: "should create an payout",
			data: &payout.CreateParams{
				ExternalID: "payout-external-id",
				Amount:     200000,
				Email:      "customer@customer.com",
			},
			expectedRes: &xendit.Payout{
				ID:                  "123",
				ExternalID:          "payout-external-id",
				Amount:              200000,
				Status:              "ISSUED",
				Created:             &date,
				ExpirationTimestamp: &date,
				MerchantName:        "Business Name",
				PayoutURL:           "https://payout-url.com/123",
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &payout.CreateParams{
				ExternalID: "payout-external-id",
			},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'Amount', 'Email'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"POST",
				xendit.Opt.XenditURL+"/payouts",
				xendit.Opt.SecretKey,
				nil,
				tC.data,
				&xendit.Payout{},
			).Return(nil)

			resp, err := payout.Create(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestGetPayout(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	date, _ := time.Parse(time.RFC3339, "2050-01-01T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *payout.GetParams
		expectedRes *xendit.Payout
		expectedErr *xendit.Error
	}{
		{
			desc: "should get an payout",
			data: &payout.GetParams{
				ID: "123",
			},
			expectedRes: &xendit.Payout{
				ID:                  "123",
				ExternalID:          "payout-external-id",
				Amount:              200000,
				Status:              "ISSUED",
				Created:             &date,
				ExpirationTimestamp: &date,
				MerchantName:        "Business Name",
				PayoutURL:           "https://payout-url.com/123",
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &payout.GetParams{},
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
				xendit.Opt.XenditURL+"/payouts/123",
				xendit.Opt.SecretKey,
				nil,
				nil,
				&xendit.Payout{},
			).Return(nil)

			resp, err := payout.Get(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestVoidPayout(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	date, _ := time.Parse(time.RFC3339, "2050-01-01T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *payout.VoidParams
		expectedRes *xendit.Payout
		expectedErr *xendit.Error
	}{
		{
			desc: "should void an payout",
			data: &payout.VoidParams{
				ID: "123",
			},
			expectedRes: &xendit.Payout{
				ID:                  "123",
				ExternalID:          "payout-external-id",
				Amount:              200000,
				Status:              "ISSUED",
				Created:             &date,
				ExpirationTimestamp: &date,
				MerchantName:        "Business Name",
				PayoutURL:           "https://payout-url.com/123",
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &payout.VoidParams{},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'ID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"POST",
				xendit.Opt.XenditURL+"/payouts/"+tC.data.ID+"/void",
				xendit.Opt.SecretKey,
				nil,
				nil,
				&xendit.Payout{},
			).Return(nil)

			resp, err := payout.Void(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
