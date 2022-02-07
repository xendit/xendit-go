package linkedaccount

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

func (m *apiRequesterMock) Call(ctx context.Context, method string, path string, secretKey string, header http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, header, params, result)

	result.(*xendit.InitializedLinkedAccount).ID = "lat-f9dc34e7-153a-444e-b337-cd2599e7f670"
	result.(*xendit.InitializedLinkedAccount).CustomerID = "6778b829-1936-4c4a-a321-9a0178840571"
	result.(*xendit.InitializedLinkedAccount).ChannelCode = xendit.DC_BRI
	result.(*xendit.InitializedLinkedAccount).Status = "PENDING"

	return nil
}

func TestInitializeLinkedAccountTokenization(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	properties := map[string]interface{}{
		"account_mobile_number": "+62818555988",
		"card_last_four":        "8888",
		"card_expiry":           "06/24",
		"account_email":         "test.email@xendit.co",
	}

	testCases := []struct {
		desc        string
		data        *InitializeLinkedAccountTokenizationParams
		expectedRes *xendit.InitializedLinkedAccount
		expectedErr *xendit.Error
	}{
		{
			desc: "should initialize linked account tokenization",
			data: &InitializeLinkedAccountTokenizationParams{
				CustomerID:  "6778b829-1936-4c4a-a321-9a0178840571",
				ChannelCode: xendit.DC_BRI,
				Properties:  properties,
			},
			expectedRes: &xendit.InitializedLinkedAccount{
				ID:          "lat-f9dc34e7-153a-444e-b337-cd2599e7f670",
				CustomerID:  "6778b829-1936-4c4a-a321-9a0178840571",
				ChannelCode: xendit.DC_BRI,
				Status:      "PENDING",
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &InitializeLinkedAccountTokenizationParams{
				ChannelCode: xendit.DC_BRI,
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
				xendit.Opt.XenditURL+"/linked_account_tokens/auth",
				xendit.Opt.SecretKey,
				http.Header{},
				tC.data,
				&xendit.InitializedLinkedAccount{},
			).Return(nil)

			resp, err := InitializeLinkedAccountTokenization(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

type apiRequesterMockValidate struct {
	mock.Mock
}

func (m *apiRequesterMockValidate) Call(ctx context.Context, method string, path string, secretKey string, header http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, header, params, result)

	result.(*xendit.ValidatedLinkedAccount).ID = "lat-f9dc34e7-153a-444e-b337-cd2599e7f670"
	result.(*xendit.ValidatedLinkedAccount).CustomerID = "6778b829-1936-4c4a-a321-9a0178840571"
	result.(*xendit.ValidatedLinkedAccount).ChannelCode = xendit.DC_BRI
	result.(*xendit.ValidatedLinkedAccount).Status = "SUCCESS"

	return nil
}

func TestValidateOTPForLinkedAccount(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMockValidate)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *ValidateOTPForLinkedAccountParams
		expectedRes *xendit.ValidatedLinkedAccount
		expectedErr *xendit.Error
	}{
		{
			desc: "should validate OTP for linked account",
			data: &ValidateOTPForLinkedAccountParams{
				LinkedAccountTokenID: "lat-f9dc34e7-153a-444e-b337-cd2599e7f670",
				OTPCode:              "333000",
			},
			expectedRes: &xendit.ValidatedLinkedAccount{
				ID:          "lat-f9dc34e7-153a-444e-b337-cd2599e7f670",
				CustomerID:  "6778b829-1936-4c4a-a321-9a0178840571",
				ChannelCode: xendit.DC_BRI,
				Status:      "SUCCESS",
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &ValidateOTPForLinkedAccountParams{
				LinkedAccountTokenID: "lat-f9dc34e7-153a-444e-b337-cd2599e7f670",
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
				xendit.Opt.XenditURL+"/linked_account_tokens/"+tC.data.LinkedAccountTokenID+"/validate_otp",
				xendit.Opt.SecretKey,
				http.Header{},
				tC.data,
				&xendit.ValidatedLinkedAccount{},
			).Return(nil)

			resp, err := ValidateOTPForLinkedAccount(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

type apiRequesterMockRetrieve struct {
	mock.Mock
}

func (m *apiRequesterMockRetrieve) Call(ctx context.Context, method string, path string, secretKey string, header http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	resultString := `[{
		"channel_code": "DC_BRI",
		"id": "la-958b3e15-a354-4ac1-80d5-d7f68061fda2",
		"properties": {
			"card_expiry": "11/23",
			"card_last_four": "8888",
			"currency": "IDR",
			"description": ""
		},
		"type": "DEBIT_CARD"
	}]`

	_ = json.Unmarshal([]byte(resultString), &result)

	return nil
}

func TestRetrieveAccessibleLinkedAccounts(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMockRetrieve)
	initTesting(apiRequesterMockObj)

	properties := map[string]interface{}{
		"card_expiry":    "11/23",
		"card_last_four": "8888",
		"currency":       "IDR",
		"description":    "",
	}

	testCases := []struct {
		desc        string
		data        *RetrieveAccessibleLinkedAccountParams
		expectedRes []xendit.AccessibleLinkedAccount
		expectedErr *xendit.Error
	}{
		{
			desc: "should retrieve accessible linked accounts",
			data: &RetrieveAccessibleLinkedAccountParams{
				LinkedAccountTokenID: "lat-f9dc34e7-153a-444e-b337-cd2599e7f670",
			},
			expectedRes: []xendit.AccessibleLinkedAccount{
				{
					ID:          "la-958b3e15-a354-4ac1-80d5-d7f68061fda2",
					ChannelCode: xendit.DC_BRI,
					AccountType: xendit.DEBIT_CARD,
					Properties:  properties,
				},
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &RetrieveAccessibleLinkedAccountParams{},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'LinkedAccountTokenID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"GET",
				xendit.Opt.XenditURL+"/linked_account_tokens/"+tC.data.LinkedAccountTokenID+"/accounts",
				xendit.Opt.SecretKey,
				nil,
				nil,
				&[]xendit.AccessibleLinkedAccount{},
			).Return(nil)

			resp, err := RetrieveAccessibleLinkedAccounts(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

type apiRequesterMockUnbind struct {
	mock.Mock
}

func (m *apiRequesterMockUnbind) Call(ctx context.Context, method string, path string, secretKey string, header http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, header, params, result)

	result.(*xendit.UnbindedLinkedAccount).ID = "lat-f9dc34e7-153a-444e-b337-cd2599e7f670"
	result.(*xendit.UnbindedLinkedAccount).IsDeleted = true

	return nil
}

func TestUnbindLinkedAccountToken(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMockUnbind)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *UnbindLinkedAccountTokenParams
		expectedRes *xendit.UnbindedLinkedAccount
		expectedErr *xendit.Error
	}{
		{
			desc: "should unbind linked account token",
			data: &UnbindLinkedAccountTokenParams{
				LinkedAccountTokenID: "lat-f9dc34e7-153a-444e-b337-cd2599e7f670",
			},
			expectedRes: &xendit.UnbindedLinkedAccount{
				ID:        "lat-f9dc34e7-153a-444e-b337-cd2599e7f670",
				IsDeleted: true,
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &UnbindLinkedAccountTokenParams{},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'LinkedAccountTokenID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"DELETE",
				xendit.Opt.XenditURL+"/linked_account_tokens/"+tC.data.LinkedAccountTokenID,
				xendit.Opt.SecretKey,
				http.Header{},
				tC.data,
				&xendit.UnbindedLinkedAccount{},
			).Return(nil)

			resp, err := UnbindLinkedAccountToken(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
