package pmsv2

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/pmsv2/constant"
	"github.com/xendit/xendit-go/pmsv2/ewallet"
)

func initTesting(apiRequesterMockObj xendit.APIRequester) {
	xendit.Opt.SecretKey = "examplesecretkey"
	xendit.SetAPIRequester(apiRequesterMockObj)
}

type apiRequesterMock struct {
	mock.Mock
}

func (m *apiRequesterMock) Call(ctx context.Context, method string, path string, secretKey string, header http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, header, params, result)

	result.(*xendit.PaymentMethodResponse).Type = constant.PaymentMethodTypeEwallet
	result.(*xendit.PaymentMethodResponse).Country = constant.CountryPH
	result.(*xendit.PaymentMethodResponse).ReferenceID = "reference_id"
	result.(*xendit.PaymentMethodResponse).Reusability = constant.ReusabilityMultipleUse

	ewallet := &xendit.EwalletMethod{
		ChannelCode:       ewallet.Astrapay,
		ChannelProperties: ewallet.ChannelProperties{},
		Account:           ewallet.Account{},
	}

	result.(*xendit.PaymentMethodResponse).Ewallet = ewallet

	return nil
}

func TestCreatePaymentMethod(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	referenceID := "reference_id"

	testCases := []struct {
		desc        string
		data        *CreatePaymentMethodParams
		expectedRes *xendit.PaymentMethodResponse
		expectedErr *xendit.Error
	}{
		{
			desc: "should create a payment",
			data: &CreatePaymentMethodParams{
				Type:        constant.PaymentMethodTypeEwallet,
				Country:     constant.CountryPH,
				CustomerID:  nil,
				ReferenceID: &referenceID,
				Reusability: constant.ReusabilityMultipleUse,
				Description: nil,
				Metadata:    nil,

				Ewallet: &ewallet.CreateMethod{
					ChannelCode:       ewallet.Astrapay,
					ChannelProperties: ewallet.ChannelProperties{},
				},
			},
			expectedRes: &xendit.PaymentMethodResponse{
				Type:        constant.PaymentMethodTypeEwallet,
				Country:     constant.CountryPH,
				CustomerID:  nil,
				ReferenceID: "reference_id",
				Reusability: constant.ReusabilityMultipleUse,
				Description: nil,
				Metadata:    nil,

				Ewallet: &xendit.EwalletMethod{
					ChannelCode:       ewallet.Astrapay,
					ChannelProperties: ewallet.ChannelProperties{},
					Account:           ewallet.Account{},
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
				"POST",
				xendit.Opt.XenditURL+"/v2/payment_methods",
				xendit.Opt.SecretKey,
				mock.AnythingOfType("http.Header"),
				tC.data,
				&xendit.PaymentMethodResponse{},
			).Return(nil)

			resp, err := CreatePaymentMethod(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestAuthPaymentMethod(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *ValidateOTPRequest
		expectedRes *xendit.PaymentMethodResponse
		expectedErr *xendit.Error
	}{
		{
			desc: "should auth",
			data: &ValidateOTPRequest{
				ID:   "id",
				Code: "auth",
			},
			expectedRes: &xendit.PaymentMethodResponse{
				Type:        constant.PaymentMethodTypeEwallet,
				Country:     constant.CountryPH,
				CustomerID:  nil,
				ReferenceID: "reference_id",
				Reusability: constant.ReusabilityMultipleUse,
				Description: nil,
				Metadata:    nil,

				Ewallet: &xendit.EwalletMethod{
					ChannelCode:       ewallet.Astrapay,
					ChannelProperties: ewallet.ChannelProperties{},
					Account:           ewallet.Account{},
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
				"POST",
				xendit.Opt.XenditURL+"/v2/payment_methods/id/auth",
				xendit.Opt.SecretKey,
				mock.AnythingOfType("http.Header"),
				tC.data,
				&xendit.PaymentMethodResponse{},
			).Return(nil)

			resp, err := AuthPaymentMethod(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestExpirePaymentMethod(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *ExpireRequest
		expectedRes *xendit.PaymentMethodResponse
		expectedErr *xendit.Error
	}{
		{
			desc: "should auth",
			data: &ExpireRequest{
				ID: "id",
			},
			expectedRes: &xendit.PaymentMethodResponse{
				Type:        constant.PaymentMethodTypeEwallet,
				Country:     constant.CountryPH,
				CustomerID:  nil,
				ReferenceID: "reference_id",
				Reusability: constant.ReusabilityMultipleUse,
				Description: nil,
				Metadata:    nil,

				Ewallet: &xendit.EwalletMethod{
					ChannelCode:       ewallet.Astrapay,
					ChannelProperties: ewallet.ChannelProperties{},
					Account:           ewallet.Account{},
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
				"POST",
				xendit.Opt.XenditURL+"/v2/payment_methods/id/expire",
				xendit.Opt.SecretKey,
				mock.AnythingOfType("http.Header"),
				tC.data,
				&xendit.PaymentMethodResponse{},
			).Return(nil)

			resp, err := ExpirePaymentMethod(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestGetPaymentMethod(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *RetrievePaymentMethodRequest
		expectedRes *xendit.PaymentMethodResponse
		expectedErr *xendit.Error
	}{
		{
			desc: "should get a payment method",
			data: &RetrievePaymentMethodRequest{
				ID: "id",
			},
			expectedRes: &xendit.PaymentMethodResponse{
				Type:        constant.PaymentMethodTypeEwallet,
				Country:     constant.CountryPH,
				CustomerID:  nil,
				ReferenceID: "reference_id",
				Reusability: constant.ReusabilityMultipleUse,
				Description: nil,
				Metadata:    nil,

				Ewallet: &xendit.EwalletMethod{
					ChannelCode:       ewallet.Astrapay,
					ChannelProperties: ewallet.ChannelProperties{},
					Account:           ewallet.Account{},
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
				xendit.Opt.XenditURL+"/v2/payment_methods/id",
				xendit.Opt.SecretKey,
				mock.AnythingOfType("http.Header"),
				nil,
				&xendit.PaymentMethodResponse{},
			).Return(nil)

			resp, err := RetrievePaymentMethod(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestGetPaymentMethods(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *RetrieveAllPaymentMethodsRequest
		expectedRes *xendit.PaymentMethodResponse
		expectedErr *xendit.Error
	}{
		{
			desc: "should get payment methods",
			data: &RetrieveAllPaymentMethodsRequest{},
			expectedRes: &xendit.PaymentMethodResponse{
				Type:        constant.PaymentMethodTypeEwallet,
				Country:     constant.CountryPH,
				CustomerID:  nil,
				ReferenceID: "reference_id",
				Reusability: constant.ReusabilityMultipleUse,
				Description: nil,
				Metadata:    nil,

				Ewallet: &xendit.EwalletMethod{
					ChannelCode:       ewallet.Astrapay,
					ChannelProperties: ewallet.ChannelProperties{},
					Account:           ewallet.Account{},
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
				xendit.Opt.XenditURL+"/v2/payment_methods?",
				xendit.Opt.SecretKey,
				mock.AnythingOfType("http.Header"),
				nil,
				&xendit.PaymentMethodResponse{},
			).Return(nil)

			resp, err := RetrieveAllPaymentMethods(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestUpdatePaymentMethod(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	status := constant.Active
	testCases := []struct {
		desc        string
		data        *UpdateRequest
		expectedRes *xendit.PaymentMethodResponse
		expectedErr *xendit.Error
	}{
		{
			desc: "should update",
			data: &UpdateRequest{
				ID:     "id",
				Status: &status,
			},
			expectedRes: &xendit.PaymentMethodResponse{
				Type:        constant.PaymentMethodTypeEwallet,
				Country:     constant.CountryPH,
				CustomerID:  nil,
				ReferenceID: "reference_id",
				Reusability: constant.ReusabilityMultipleUse,
				Description: nil,
				Metadata:    nil,

				Ewallet: &xendit.EwalletMethod{
					ChannelCode:       ewallet.Astrapay,
					ChannelProperties: ewallet.ChannelProperties{},
					Account:           ewallet.Account{},
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
				"PATCH",
				xendit.Opt.XenditURL+"/v2/payment_methods/id",
				xendit.Opt.SecretKey,
				mock.AnythingOfType("http.Header"),
				tC.data,
				&xendit.PaymentMethodResponse{},
			).Return(nil)

			resp, err := UpdatePaymentMethod(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestRetrievePayments(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *RetrievePaymentsRequest
		expectedRes *xendit.PaymentMethodResponse
		expectedErr *xendit.Error
	}{
		{
			desc: "should get a payment method",
			data: &RetrievePaymentsRequest{
				ID: "id",
				PaginationFilters: PaginationFilters{
					Limit: 5,
				},
			},
			expectedRes: &xendit.PaymentMethodResponse{
				Type:        constant.PaymentMethodTypeEwallet,
				Country:     constant.CountryPH,
				CustomerID:  nil,
				ReferenceID: "reference_id",
				Reusability: constant.ReusabilityMultipleUse,
				Description: nil,
				Metadata:    nil,

				Ewallet: &xendit.EwalletMethod{
					ChannelCode:       ewallet.Astrapay,
					ChannelProperties: ewallet.ChannelProperties{},
					Account:           ewallet.Account{},
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
				xendit.Opt.XenditURL+"/v2/payment_methods/id/payments?limit=5",
				xendit.Opt.SecretKey,
				mock.AnythingOfType("http.Header"),
				nil,
				&xendit.PaymentMethodResponse{},
			).Return(nil)

			resp, err := RetrievePayments(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
