package ewallet

import (
	"context"
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
				xendit.Opt.XenditURL+"/ewallets",
				xendit.Opt.SecretKey,
				&http.Header{},
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

func (m *apiRequesterMockGet) Call(ctx context.Context, method string, path string, secretKey string, header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

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
				xendit.Opt.XenditURL+"/ewallets?"+tC.data.QueryString(),
				xendit.Opt.SecretKey,
				nil,
				nil,
				&getPaymentStatusResponse{},
			).Return(nil)

			resp, err := GetPaymentStatus(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

type apiRequesterMockPostCharge struct {
	mock.Mock
}

func (m *apiRequesterMockPostCharge) Call(ctx context.Context, method string, path string, secretKey string, header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, header, params, result)

	result.(*xendit.EWalletCharge).ID = "ewc_f3925450-5c54-4777-98c1-fcf22b0d1e1c"
	result.(*xendit.EWalletCharge).BusinessID = "business-id-example"
	result.(*xendit.EWalletCharge).ReferenceID = "test-reference-id"
	result.(*xendit.EWalletCharge).Currency = "IDR"
	result.(*xendit.EWalletCharge).ChargeAmount = 10000
	result.(*xendit.EWalletCharge).CaptureAmount = 10000
	result.(*xendit.EWalletCharge).CheckoutMethod = "ONE_TIME_PAYMENT"
	result.(*xendit.EWalletCharge).ChannelCode = "ID_SHOPEEPAY"
	result.(*xendit.EWalletCharge).ChannelProperties = map[string]string{
		"success_redirect_url": "https://yourwebsite.com/order/123",
	}
	result.(*xendit.EWalletCharge).Actions = map[string]string{
		"desktop_web_checkout_url":     "https://desktop.web.checkout.url",
		"mobile_web_checkout_url":      "https://mobile.web.checkout.url",
		"mobile_deeplink_checkout_url": "https://mobile.deeplink.checkout.url",
		"qr_checkout_string":           "test-qr-string",
	}
	result.(*xendit.EWalletCharge).IsRedirectRequired = true
	result.(*xendit.EWalletCharge).CallbackURL = "https://yourwebsite.com/order/123"
	result.(*xendit.EWalletCharge).Created = "2021-02-09T06:22:35.064408Z"
	result.(*xendit.EWalletCharge).Updated = "2021-02-09T06:22:35.064408Z"
	result.(*xendit.EWalletCharge).CaptureNow = true

	return nil
}

func TestCreateEWalletCharge(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMockPostCharge)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *CreateEWalletChargeParams
		expectedRes *xendit.EWalletCharge
		expectedErr *xendit.Error
	}{
		{
			desc: "should create an e-wallet charge",
			data: &CreateEWalletChargeParams{
				ReferenceID:    "test-reference-id",
				Currency:       "IDR",
				Amount:         10000,
				CheckoutMethod: "ONE_TIME_PAYMENT",
				ChannelCode:    "ID_SHOPEEPAY",
				ChannelProperties: map[string]string{
					"success_redirect_url": "https://yourwebsite.com/order/123",
				},
			},
			expectedRes: &xendit.EWalletCharge{
				ID:             "ewc_f3925450-5c54-4777-98c1-fcf22b0d1e1c",
				BusinessID:     "business-id-example",
				ReferenceID:    "test-reference-id",
				Currency:       "IDR",
				ChargeAmount:   10000,
				CaptureAmount:  10000,
				CheckoutMethod: "ONE_TIME_PAYMENT",
				ChannelCode:    "ID_SHOPEEPAY",
				ChannelProperties: map[string]string{
					"success_redirect_url": "https://yourwebsite.com/order/123",
				},
				Actions: map[string]string{
					"desktop_web_checkout_url":     "https://desktop.web.checkout.url",
					"mobile_web_checkout_url":      "https://mobile.web.checkout.url",
					"mobile_deeplink_checkout_url": "https://mobile.deeplink.checkout.url",
					"qr_checkout_string":           "test-qr-string",
				},
				IsRedirectRequired: true,
				CallbackURL:        "https://yourwebsite.com/order/123",
				Created:            "2021-02-09T06:22:35.064408Z",
				Updated:            "2021-02-09T06:22:35.064408Z",
				CaptureNow:         true,
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &CreateEWalletChargeParams{
				Currency:       "IDR",
				Amount:         10000,
				CheckoutMethod: "ONE_TIME_PAYMENT",
				ChannelCode:    "ID_SHOPEEPAY",
				ChannelProperties: map[string]string{
					"success_redirect_url": "https://yourwebsite.com/order/123",
				},
			},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'ReferenceID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"POST",
				xendit.Opt.XenditURL+"/ewallets/charges",
				xendit.Opt.SecretKey,
				&http.Header{},
				tC.data,
				&xendit.EWalletCharge{},
			).Return(nil)

			resp, err := CreateEWalletCharge(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

type apiRequesterMockGetCharge struct {
	mock.Mock
}

func (m *apiRequesterMockGetCharge) Call(ctx context.Context, method string, path string, secretKey string, header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	result.(*EWalletChargeResponse).ID = "ewc_f3925450-5c54-4777-98c1-fcf22b0d1e1c"
	result.(*EWalletChargeResponse).BusinessID = "business-id-example"
	result.(*EWalletChargeResponse).ReferenceID = "test-reference-id"
	result.(*EWalletChargeResponse).Currency = "IDR"
	result.(*EWalletChargeResponse).ChargeAmount = 10000
	result.(*EWalletChargeResponse).CaptureAmount = 10000
	result.(*EWalletChargeResponse).CheckoutMethod = "ONE_TIME_PAYMENT"
	result.(*EWalletChargeResponse).ChannelCode = "ID_SHOPEEPAY"
	result.(*EWalletChargeResponse).ChannelProperties = map[string]string{
		"success_redirect_url": "https://yourwebsite.com/order/123",
	}
	result.(*EWalletChargeResponse).Actions = map[string]string{
		"desktop_web_checkout_url":     "https://desktop.web.checkout.url",
		"mobile_web_checkout_url":      "https://mobile.web.checkout.url",
		"mobile_deeplink_checkout_url": "https://mobile.deeplink.checkout.url",
		"qr_checkout_string":           "test-qr-string",
	}
	result.(*EWalletChargeResponse).IsRedirectRequired = true
	result.(*EWalletChargeResponse).CallbackURL = "https://yourwebsite.com/order/123"
	result.(*EWalletChargeResponse).Created = "2021-02-09T06:22:35.064408Z"
	result.(*EWalletChargeResponse).Updated = "2021-02-09T06:22:35.064408Z"
	result.(*EWalletChargeResponse).CaptureNow = true

	return nil
}

func TestGetEWalletChargeStatus(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMockGetCharge)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *GetEWalletChargeStatusParams
		expectedRes *xendit.EWalletCharge
		expectedErr *xendit.Error
	}{
		{
			desc: "should get an e-wallet charge status",
			data: &GetEWalletChargeStatusParams{
				ChargeID: "ewc_f3925450-5c54-4777-98c1-fcf22b0d1e1c",
			},
			expectedRes: &xendit.EWalletCharge{
				ID:             "ewc_f3925450-5c54-4777-98c1-fcf22b0d1e1c",
				BusinessID:     "business-id-example",
				ReferenceID:    "test-reference-id",
				Currency:       "IDR",
				ChargeAmount:   10000,
				CaptureAmount:  10000,
				CheckoutMethod: "ONE_TIME_PAYMENT",
				ChannelCode:    "ID_SHOPEEPAY",
				ChannelProperties: map[string]string{
					"success_redirect_url": "https://yourwebsite.com/order/123",
				},
				Actions: map[string]string{
					"desktop_web_checkout_url":     "https://desktop.web.checkout.url",
					"mobile_web_checkout_url":      "https://mobile.web.checkout.url",
					"mobile_deeplink_checkout_url": "https://mobile.deeplink.checkout.url",
					"qr_checkout_string":           "test-qr-string",
				},
				IsRedirectRequired: true,
				CallbackURL:        "https://yourwebsite.com/order/123",
				Created:            "2021-02-09T06:22:35.064408Z",
				Updated:            "2021-02-09T06:22:35.064408Z",
				CaptureNow:         true,
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &GetEWalletChargeStatusParams{},
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
				xendit.Opt.XenditURL+"/ewallets/charges/"+tC.data.ChargeID,
				xendit.Opt.SecretKey,
				nil,
				nil,
				&EWalletChargeResponse{},
			).Return(nil)

			resp, err := GetEWalletChargeStatus(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
