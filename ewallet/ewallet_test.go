package ewallet

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/utils/validator"
)

func initTesting(apiRequesterMockObj xendit.APIRequester) {
	xendit.Opt.SecretKey = "xnd_development_REt02KJzkM6AootfKnDrMw1Sse4LlzEDHfKzXoBocqIEiH4bqjHUJXbl6Cfaab"
	xendit.SetAPIRequester(apiRequesterMockObj)
}

type apiRequesterMock struct {
	mock.Mock
}

func (m *apiRequesterMock) Call(ctx context.Context, method string, path string, secretKey string, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, params, result)

	result.(*xendit.EWallet).EWalletType = xendit.EWalletTypeDANA
	result.(*xendit.EWallet).ExternalID = "dana-ewallet"
	result.(*xendit.EWallet).Amount = 200000
	result.(*xendit.EWallet).CheckoutURL = "mystore.com/callback"

	return nil
}

func TestCreatePayment(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *CreatePaymentParams
		expectedRes *xendit.EWallet
		expectedErr *xendit.Error
	}{
		{
			desc: "should create a payment",
			data: &CreatePaymentParams{
				ExternalID:  "dana-ewallet",
				Amount:      200000,
				Phone:       "08123123123",
				EWalletType: xendit.EWalletTypeDANA,
				CallbackURL: "mystore.com/callback",
				RedirectURL: "mystore.com/redirect",
			},
			expectedRes: &xendit.EWallet{
				EWalletType: xendit.EWalletTypeDANA,
				ExternalID:  "dana-ewallet",
				Amount:      200000,
				CheckoutURL: "mystore.com/callback",
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &CreatePaymentParams{
				EWalletType: xendit.EWalletTypeDANA,
				ExternalID:  "dana-ewallet",
			},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'Amount'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"POST",
				"https://api.xendit.co/ewallets",
				xendit.Opt.SecretKey,
				tC.data,
				&xendit.EWallet{},
			).Return(nil)

			resp, err := CreatePayment(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

type apiRequesterMockGet struct {
	mock.Mock
}

func (m *apiRequesterMockGet) Call(ctx context.Context, method string, path string, secretKey string, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, params, result)

	result.(*getPaymentStatusResponse).EWalletType = xendit.EWalletTypeDANA
	result.(*getPaymentStatusResponse).ExternalID = "dana-ewallet"
	result.(*getPaymentStatusResponse).Amount = 200000
	result.(*getPaymentStatusResponse).CheckoutURL = "mystore.com/callback"

	return nil
}

func TestGetPaymentStatus(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMockGet)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *GetPaymentStatusParams
		expectedRes *xendit.EWallet
		expectedErr *xendit.Error
	}{
		{
			desc: "should get a payment status",
			data: &GetPaymentStatusParams{
				ExternalID:  "dana-ewallet",
				EWalletType: xendit.EWalletTypeDANA,
			},
			expectedRes: &xendit.EWallet{
				EWalletType: xendit.EWalletTypeDANA,
				ExternalID:  "dana-ewallet",
				Amount:      200000,
				CheckoutURL: "mystore.com/callback",
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &GetPaymentStatusParams{
				EWalletType: xendit.EWalletTypeDANA,
			},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'ExternalID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"GET",
				"https://api.xendit.co/ewallets?"+tC.data.QueryString(),
				xendit.Opt.SecretKey,
				nil,
				&getPaymentStatusResponse{},
			).Return(nil)

			resp, err := GetPaymentStatus(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
