package directdebitpayment

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

	result.(*xendit.DirectDebitPayment).ID = "ddpy-7e61b0a7-92f9-4762-a994-c2936306f44c"
	result.(*xendit.DirectDebitPayment).ReferenceID = "direct-debit-ref-id"
	result.(*xendit.DirectDebitPayment).PaymentMethodID = "pm-ebb1c863-c7b5-4f20-b116-b3071b1d3aef"
	result.(*xendit.DirectDebitPayment).ChannelCode = xendit.DC_BRI
	result.(*xendit.DirectDebitPayment).Currency = "IDR"
	result.(*xendit.DirectDebitPayment).Amount = 15000
	result.(*xendit.DirectDebitPayment).IsOTPRequired = true
	result.(*xendit.DirectDebitPayment).Basket = []xendit.DirectDebitBasketItem{
		{
			ReferenceID: "basket-product-ref-id",
			Name:        "product-name",
			Category:    "mechanics",
			Market:      "ID",
			Price:       50000,
			Quantity:    5,
			Type:        "product type",
			SubCategory: "product sub category",
			Description: "product description",
			URL:         "https://product.url",
		},
	}
	result.(*xendit.DirectDebitPayment).Description = "test description"
	result.(*xendit.DirectDebitPayment).Status = "PENDING"
	result.(*xendit.DirectDebitPayment).Metadata = map[string]interface{}{
		"meta": "data",
	}
	result.(*xendit.DirectDebitPayment).Created = "2021-07-16T02:19:07.277466Z"
	result.(*xendit.DirectDebitPayment).Updated = "2021-07-16T02:19:07.277466Z"
	result.(*xendit.DirectDebitPayment).Device = xendit.DirectDebitDevice{
		ID:        "device-id",
		IPAddress: "0.0.0.0",
		UserAgent: "user-agent",
		ADID:      "ad-id",
		Imei:      "123a456b789c",
	}
	result.(*xendit.DirectDebitPayment).RefundedAmount = 0
	result.(*xendit.DirectDebitPayment).Refunds = xendit.DirectDebitRefunds{
		Data:    []string{"a1b", "c2d", "e3f", "g4h"},
		HasMore: false,
		URL:     "https://ref.unds",
	}
	result.(*xendit.DirectDebitPayment).FailureCode = "failure-code"
	result.(*xendit.DirectDebitPayment).OTPMobileNumber = "+6281234567890"
	result.(*xendit.DirectDebitPayment).OTPExpirationTimestamp = "2100-07-16T02:19:07.277466Z"
	result.(*xendit.DirectDebitPayment).SuccessRedirectURL = "https://success-redirect.url"
	result.(*xendit.DirectDebitPayment).CheckoutURL = "https://checkout.url"
	result.(*xendit.DirectDebitPayment).FailureRedirectURL = "https://failure-redirect.url"
	result.(*xendit.DirectDebitPayment).RequiredAction = "VALIDATE_OTP"

	return nil
}

func TestCreateDirectDebitPayment(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *CreateDirectDebitPaymentParams
		expectedRes *xendit.DirectDebitPayment
		expectedErr *xendit.Error
	}{
		{
			desc: "should create direct debit payment",
			data: &CreateDirectDebitPaymentParams{
				IdempotencyKey:  "idem-key",
				ReferenceID:     "test-ref-id",
				PaymentMethodID: "test-pm-id",
				Currency:        "IDR",
				Amount:          15000,
				CallbackURL:     "http://webhook.site",
				EnableOTP:       true,
				Description:     "Test description",
				Basket: []xendit.DirectDebitBasketItem{
					{
						ReferenceID: "basket-product-ref-id",
						Name:        "product-name",
						Category:    "mechanics",
						Market:      "ID",
						Price:       50000,
						Quantity:    5,
						Type:        "product type",
						SubCategory: "product sub category",
						Description: "product description",
						URL:         "https://product.url",
					},
				},
				Device: xendit.DirectDebitDevice{
					ID:        "device-id",
					IPAddress: "0.0.0.0",
					UserAgent: "user-agent",
					ADID:      "ad-id",
					Imei:      "123a456b789c",
				},
				SuccessRedirectURL: "https://success-redirect.url",
				FailureRedirectURL: "https://failure-redirect.url",
				Metadata: map[string]interface{}{
					"meta": "data",
				},
			},
			expectedRes: &xendit.DirectDebitPayment{
				ID:              "ddpy-7e61b0a7-92f9-4762-a994-c2936306f44c",
				ReferenceID:     "direct-debit-ref-id",
				PaymentMethodID: "pm-ebb1c863-c7b5-4f20-b116-b3071b1d3aef",
				ChannelCode:     xendit.DC_BRI,
				Currency:        "IDR",
				Amount:          15000,
				IsOTPRequired:   true,
				Basket: []xendit.DirectDebitBasketItem{
					{
						ReferenceID: "basket-product-ref-id",
						Name:        "product-name",
						Category:    "mechanics",
						Market:      "ID",
						Price:       50000,
						Quantity:    5,
						Type:        "product type",
						SubCategory: "product sub category",
						Description: "product description",
						URL:         "https://product.url",
					},
				},
				Description: "test description",
				Status:      "PENDING",
				Metadata: map[string]interface{}{
					"meta": "data",
				},
				Created: "2021-07-16T02:19:07.277466Z",
				Updated: "2021-07-16T02:19:07.277466Z",
				Device: xendit.DirectDebitDevice{
					ID:        "device-id",
					IPAddress: "0.0.0.0",
					UserAgent: "user-agent",
					ADID:      "ad-id",
					Imei:      "123a456b789c",
				},
				RefundedAmount: 0,
				Refunds: xendit.DirectDebitRefunds{
					Data:    []string{"a1b", "c2d", "e3f", "g4h"},
					HasMore: false,
					URL:     "https://ref.unds",
				},
				FailureCode:            "failure-code",
				OTPMobileNumber:        "+6281234567890",
				OTPExpirationTimestamp: "2100-07-16T02:19:07.277466Z",
				SuccessRedirectURL:     "https://success-redirect.url",
				CheckoutURL:            "https://checkout.url",
				FailureRedirectURL:     "https://failure-redirect.url",
				RequiredAction:         "VALIDATE_OTP",
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &CreateDirectDebitPaymentParams{
				IdempotencyKey:  "idem-key-2",
				PaymentMethodID: "test-pm-id",
				Currency:        "IDR",
				Amount:          15000,
				CallbackURL:     "http://webhook.site",
				EnableOTP:       true,
				Description:     "Test description",
				Basket: []xendit.DirectDebitBasketItem{
					{
						ReferenceID: "basket-product-ref-id",
						Name:        "product-name",
						Category:    "mechanics",
						Market:      "ID",
						Price:       50000,
						Quantity:    5,
						Type:        "product type",
						SubCategory: "product sub category",
						Description: "product description",
						URL:         "https://product.url",
					},
				},
				Device: xendit.DirectDebitDevice{
					ID:        "device-id",
					IPAddress: "0.0.0.0",
					UserAgent: "user-agent",
					ADID:      "ad-id",
					Imei:      "123a456b789c",
				},
				SuccessRedirectURL: "https://success-redirect.url",
				FailureRedirectURL: "https://failure-redirect.url",
				Metadata: map[string]interface{}{
					"meta": "data",
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
				xendit.Opt.XenditURL+"/direct_debits",
				xendit.Opt.SecretKey,
				&http.Header{"Idempotency-Key": []string{tC.data.IdempotencyKey}},
				tC.data,
				&xendit.DirectDebitPayment{},
			).Return(nil)

			resp, err := CreateDirectDebitPayment(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

type apiRequesterMockValidate struct {
	mock.Mock
}

func (m *apiRequesterMockValidate) Call(ctx context.Context, method string, path string, secretKey string, header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, header, params, result)

	result.(*xendit.DirectDebitPayment).ID = "ddpy-7e61b0a7-92f9-4762-a994-c2936306f44c"
	result.(*xendit.DirectDebitPayment).ReferenceID = "direct-debit-ref-id"
	result.(*xendit.DirectDebitPayment).PaymentMethodID = "pm-ebb1c863-c7b5-4f20-b116-b3071b1d3aef"
	result.(*xendit.DirectDebitPayment).ChannelCode = xendit.DC_BRI
	result.(*xendit.DirectDebitPayment).Currency = "IDR"
	result.(*xendit.DirectDebitPayment).Amount = 15000
	result.(*xendit.DirectDebitPayment).IsOTPRequired = true
	result.(*xendit.DirectDebitPayment).Basket = []xendit.DirectDebitBasketItem{
		{
			ReferenceID: "basket-product-ref-id",
			Name:        "product-name",
			Category:    "mechanics",
			Market:      "ID",
			Price:       50000,
			Quantity:    5,
			Type:        "product type",
			SubCategory: "product sub category",
			Description: "product description",
			URL:         "https://product.url",
		},
	}
	result.(*xendit.DirectDebitPayment).Description = "test description"
	result.(*xendit.DirectDebitPayment).Status = "COMPLETED"
	result.(*xendit.DirectDebitPayment).Metadata = map[string]interface{}{
		"meta": "data",
	}
	result.(*xendit.DirectDebitPayment).Created = "2021-07-16T02:19:07.277466Z"
	result.(*xendit.DirectDebitPayment).Updated = "2021-07-16T02:19:07.277466Z"
	result.(*xendit.DirectDebitPayment).Device = xendit.DirectDebitDevice{
		ID:        "device-id",
		IPAddress: "0.0.0.0",
		UserAgent: "user-agent",
		ADID:      "ad-id",
		Imei:      "123a456b789c",
	}
	result.(*xendit.DirectDebitPayment).RefundedAmount = 0
	result.(*xendit.DirectDebitPayment).Refunds = xendit.DirectDebitRefunds{
		Data:    []string{"a1b", "c2d", "e3f", "g4h"},
		HasMore: false,
		URL:     "https://ref.unds",
	}
	result.(*xendit.DirectDebitPayment).FailureCode = "failure-code"
	result.(*xendit.DirectDebitPayment).OTPMobileNumber = "+6281234567890"
	result.(*xendit.DirectDebitPayment).OTPExpirationTimestamp = "2100-07-16T02:19:07.277466Z"
	result.(*xendit.DirectDebitPayment).SuccessRedirectURL = "https://success-redirect.url"
	result.(*xendit.DirectDebitPayment).CheckoutURL = "https://checkout.url"
	result.(*xendit.DirectDebitPayment).FailureRedirectURL = "https://failure-redirect.url"
	result.(*xendit.DirectDebitPayment).RequiredAction = "VALIDATE_OTP"

	return nil
}

func TestValidateOTPForDirectDebitPayment(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMockValidate)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *ValidateOTPForDirectDebitPaymentParams
		expectedRes *xendit.DirectDebitPayment
		expectedErr *xendit.Error
	}{
		{
			desc: "should validate OTP for direct debit payment",
			data: &ValidateOTPForDirectDebitPaymentParams{
				DirectDebitID: "ddpy-7e61b0a7-92f9-4762-a994-c2936306f44c",
				OTPCode:       "222000",
			},
			expectedRes: &xendit.DirectDebitPayment{
				ID:              "ddpy-7e61b0a7-92f9-4762-a994-c2936306f44c",
				ReferenceID:     "direct-debit-ref-id",
				PaymentMethodID: "pm-ebb1c863-c7b5-4f20-b116-b3071b1d3aef",
				ChannelCode:     xendit.DC_BRI,
				Currency:        "IDR",
				Amount:          15000,
				IsOTPRequired:   true,
				Basket: []xendit.DirectDebitBasketItem{
					{
						ReferenceID: "basket-product-ref-id",
						Name:        "product-name",
						Category:    "mechanics",
						Market:      "ID",
						Price:       50000,
						Quantity:    5,
						Type:        "product type",
						SubCategory: "product sub category",
						Description: "product description",
						URL:         "https://product.url",
					},
				},
				Description: "test description",
				Status:      "COMPLETED",
				Metadata: map[string]interface{}{
					"meta": "data",
				},
				Created: "2021-07-16T02:19:07.277466Z",
				Updated: "2021-07-16T02:19:07.277466Z",
				Device: xendit.DirectDebitDevice{
					ID:        "device-id",
					IPAddress: "0.0.0.0",
					UserAgent: "user-agent",
					ADID:      "ad-id",
					Imei:      "123a456b789c",
				},
				RefundedAmount: 0,
				Refunds: xendit.DirectDebitRefunds{
					Data:    []string{"a1b", "c2d", "e3f", "g4h"},
					HasMore: false,
					URL:     "https://ref.unds",
				},
				FailureCode:            "failure-code",
				OTPMobileNumber:        "+6281234567890",
				OTPExpirationTimestamp: "2100-07-16T02:19:07.277466Z",
				SuccessRedirectURL:     "https://success-redirect.url",
				CheckoutURL:            "https://checkout.url",
				FailureRedirectURL:     "https://failure-redirect.url",
				RequiredAction:         "VALIDATE_OTP",
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &ValidateOTPForDirectDebitPaymentParams{
				DirectDebitID: "ddpy-7e61b0a7-92f9-4762-a994-c2936306f44c",
			},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'OTPCode'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"POST",
				xendit.Opt.XenditURL+"/direct_debits/"+tC.data.DirectDebitID+"/validate_otp/",
				xendit.Opt.SecretKey,
				&http.Header{},
				tC.data,
				&xendit.DirectDebitPayment{},
			).Return(nil)

			resp, err := ValidateOTPForDirectDebitPayment(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

type apiRequesterMockGetByID struct {
	mock.Mock
}

func (m *apiRequesterMockGetByID) Call(ctx context.Context, method string, path string, secretKey string, header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	result.(*DirectDebitPaymentResponse).ID = "ddpy-7e61b0a7-92f9-4762-a994-c2936306f44c"
	result.(*DirectDebitPaymentResponse).ReferenceID = "direct-debit-ref-id"
	result.(*DirectDebitPaymentResponse).PaymentMethodID = "pm-ebb1c863-c7b5-4f20-b116-b3071b1d3aef"
	result.(*DirectDebitPaymentResponse).ChannelCode = xendit.DC_BRI
	result.(*DirectDebitPaymentResponse).Currency = "IDR"
	result.(*DirectDebitPaymentResponse).Amount = 15000
	result.(*DirectDebitPaymentResponse).IsOTPRequired = true
	result.(*DirectDebitPaymentResponse).Basket = []xendit.DirectDebitBasketItem{
		{
			ReferenceID: "basket-product-ref-id",
			Name:        "product-name",
			Category:    "mechanics",
			Market:      "ID",
			Price:       50000,
			Quantity:    5,
			Type:        "product type",
			SubCategory: "product sub category",
			Description: "product description",
			URL:         "https://product.url",
		},
	}
	result.(*DirectDebitPaymentResponse).Description = "test description"
	result.(*DirectDebitPaymentResponse).Status = "COMPLETED"
	result.(*DirectDebitPaymentResponse).Metadata = map[string]interface{}{
		"meta": "data",
	}
	result.(*DirectDebitPaymentResponse).Created = "2021-07-16T02:19:07.277466Z"
	result.(*DirectDebitPaymentResponse).Updated = "2021-07-16T02:19:07.277466Z"
	result.(*DirectDebitPaymentResponse).Device = xendit.DirectDebitDevice{
		ID:        "device-id",
		IPAddress: "0.0.0.0",
		UserAgent: "user-agent",
		ADID:      "ad-id",
		Imei:      "123a456b789c",
	}
	result.(*DirectDebitPaymentResponse).RefundedAmount = 0
	result.(*DirectDebitPaymentResponse).Refunds = xendit.DirectDebitRefunds{
		Data:    []string{"a1b", "c2d", "e3f", "g4h"},
		HasMore: false,
		URL:     "https://ref.unds",
	}
	result.(*DirectDebitPaymentResponse).FailureCode = "failure-code"
	result.(*DirectDebitPaymentResponse).OTPMobileNumber = "+6281234567890"
	result.(*DirectDebitPaymentResponse).OTPExpirationTimestamp = "2100-07-16T02:19:07.277466Z"
	result.(*DirectDebitPaymentResponse).SuccessRedirectURL = "https://success-redirect.url"
	result.(*DirectDebitPaymentResponse).CheckoutURL = "https://checkout.url"
	result.(*DirectDebitPaymentResponse).FailureRedirectURL = "https://failure-redirect.url"
	result.(*DirectDebitPaymentResponse).RequiredAction = "VALIDATE_OTP"

	return nil
}

func TestGetDirectDebitPaymentStatusByID(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMockGetByID)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *GetDirectDebitPaymentStatusByIDParams
		expectedRes *xendit.DirectDebitPayment
		expectedErr *xendit.Error
	}{
		{
			desc: "should retrieve direct debit payment by ID",
			data: &GetDirectDebitPaymentStatusByIDParams{
				ID: "ddpy-7e61b0a7-92f9-4762-a994-c2936306f44c",
			},
			expectedRes: &xendit.DirectDebitPayment{
				ID:              "ddpy-7e61b0a7-92f9-4762-a994-c2936306f44c",
				ReferenceID:     "direct-debit-ref-id",
				PaymentMethodID: "pm-ebb1c863-c7b5-4f20-b116-b3071b1d3aef",
				ChannelCode:     xendit.DC_BRI,
				Currency:        "IDR",
				Amount:          15000,
				IsOTPRequired:   true,
				Basket: []xendit.DirectDebitBasketItem{
					{
						ReferenceID: "basket-product-ref-id",
						Name:        "product-name",
						Category:    "mechanics",
						Market:      "ID",
						Price:       50000,
						Quantity:    5,
						Type:        "product type",
						SubCategory: "product sub category",
						Description: "product description",
						URL:         "https://product.url",
					},
				},
				Description: "test description",
				Status:      "COMPLETED",
				Metadata: map[string]interface{}{
					"meta": "data",
				},
				Created: "2021-07-16T02:19:07.277466Z",
				Updated: "2021-07-16T02:19:07.277466Z",
				Device: xendit.DirectDebitDevice{
					ID:        "device-id",
					IPAddress: "0.0.0.0",
					UserAgent: "user-agent",
					ADID:      "ad-id",
					Imei:      "123a456b789c",
				},
				RefundedAmount: 0,
				Refunds: xendit.DirectDebitRefunds{
					Data:    []string{"a1b", "c2d", "e3f", "g4h"},
					HasMore: false,
					URL:     "https://ref.unds",
				},
				FailureCode:            "failure-code",
				OTPMobileNumber:        "+6281234567890",
				OTPExpirationTimestamp: "2100-07-16T02:19:07.277466Z",
				SuccessRedirectURL:     "https://success-redirect.url",
				CheckoutURL:            "https://checkout.url",
				FailureRedirectURL:     "https://failure-redirect.url",
				RequiredAction:         "VALIDATE_OTP",
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &GetDirectDebitPaymentStatusByIDParams{},
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
				xendit.Opt.XenditURL+"/direct_debits/"+tC.data.ID+"/",
				xendit.Opt.SecretKey,
				nil,
				nil,
				&DirectDebitPaymentResponse{},
			).Return(nil)

			resp, err := GetDirectDebitPaymentStatusByID(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

type apiRequesterMockGetByReferenceID struct {
	mock.Mock
}

func (m *apiRequesterMockGetByReferenceID) Call(ctx context.Context, method string, path string, secretKey string, header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	resultString := `[{
		"id": "ddpy-7e61b0a7-92f9-4762-a994-c2936306f44c",
		"reference_id": "direct-debit-ref-id",
		"payment_method_id": "pm-ebb1c863-c7b5-4f20-b116-b3071b1d3aef",
		"channel_code": "DC_BRI",
		"currency": "IDR",
		"amount": 15000,
		"is_otp_required": true,
		"basket": [
			{
				"reference_id":	"basket-product-ref-id",
				"name":	"product-name",
				"category": "mechanics",
				"market": "ID",
				"price": 50000,
				"quantity": 5,
				"type": "product type",
				"sub_category": "product sub category",
				"description": "product description",
				"url": "https://product.url"
			}
		],
		"description": "test description",
		"status": "COMPLETED",
		"metadata": {
			"meta": "data"
		},
		"created": "2021-07-16T02:19:07.277466Z",
		"updated": "2021-07-16T02:19:07.277466Z",
		"device": {
			"id": "device-id",
			"ip_address": "0.0.0.0",
			"user_agent": "user-agent",
			"ad_id": "ad-id",
			"imei": "123a456b789c"
		},
		"refunded_amount": 0,
		"refunds": {
			"data":	["a1b", "c2d", "e3f", "g4h"],
			"has_more": false,
			"url": "https://ref.unds"
		},
		"failure_code": "failure-code",
		"otp_mobile_number": "+6281234567890",
		"otp_expiration_timestamp": "2100-07-16T02:19:07.277466Z",
		"success_redirect_url": "https://success-redirect.url",
		"checkout_url": "https://checkout.url",
		"failure_redirect_url": "https://failure-redirect.url",
		"required_action": "VALIDATE_OTP"
	}]`

	_ = json.Unmarshal([]byte(resultString), &result)

	return nil
}

func TestGetDirectDebitPaymentStatusByReferenceID(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMockGetByReferenceID)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *GetDirectDebitPaymentStatusByReferenceIDParams
		expectedRes []xendit.DirectDebitPayment
		expectedErr *xendit.Error
	}{
		{
			desc: "should gets direct debit payment status by reference ID",
			data: &GetDirectDebitPaymentStatusByReferenceIDParams{
				ReferenceID: "direct-debit-ref-id",
			},
			expectedRes: []xendit.DirectDebitPayment{
				{
					ID:              "ddpy-7e61b0a7-92f9-4762-a994-c2936306f44c",
					ReferenceID:     "direct-debit-ref-id",
					PaymentMethodID: "pm-ebb1c863-c7b5-4f20-b116-b3071b1d3aef",
					ChannelCode:     xendit.DC_BRI,
					Currency:        "IDR",
					Amount:          15000,
					IsOTPRequired:   true,
					Basket: []xendit.DirectDebitBasketItem{
						{
							ReferenceID: "basket-product-ref-id",
							Name:        "product-name",
							Category:    "mechanics",
							Market:      "ID",
							Price:       50000,
							Quantity:    5,
							Type:        "product type",
							SubCategory: "product sub category",
							Description: "product description",
							URL:         "https://product.url",
						},
					},
					Description: "test description",
					Status:      "COMPLETED",
					Metadata: map[string]interface{}{
						"meta": "data",
					},
					Created: "2021-07-16T02:19:07.277466Z",
					Updated: "2021-07-16T02:19:07.277466Z",
					Device: xendit.DirectDebitDevice{
						ID:        "device-id",
						IPAddress: "0.0.0.0",
						UserAgent: "user-agent",
						ADID:      "ad-id",
						Imei:      "123a456b789c",
					},
					RefundedAmount: 0,
					Refunds: xendit.DirectDebitRefunds{
						Data:    []string{"a1b", "c2d", "e3f", "g4h"},
						HasMore: false,
						URL:     "https://ref.unds",
					},
					FailureCode:            "failure-code",
					OTPMobileNumber:        "+6281234567890",
					OTPExpirationTimestamp: "2100-07-16T02:19:07.277466Z",
					SuccessRedirectURL:     "https://success-redirect.url",
					CheckoutURL:            "https://checkout.url",
					FailureRedirectURL:     "https://failure-redirect.url",
					RequiredAction:         "VALIDATE_OTP",
				},
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &GetDirectDebitPaymentStatusByReferenceIDParams{},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'ReferenceID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"GET",
				xendit.Opt.XenditURL+"/direct_debits?"+tC.data.QueryString(),
				xendit.Opt.SecretKey,
				nil,
				nil,
				&[]xendit.DirectDebitPayment{},
			).Return(nil)

			resp, err := GetDirectDebitPaymentStatusByReferenceID(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
