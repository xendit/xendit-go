package account_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"

	"github.com/xendit/xendit-go/account"
	"github.com/xendit/xendit-go/utils/validator"
)

func initTesting(apiRequesterMockObj xendit.APIRequester) {
	xendit.Opt.SecretKey = "examplesecretkey"
	xendit.SetAPIRequester(apiRequesterMockObj)
}

type apiRequesterMock struct {
	mock.Mock
}

func (m *apiRequesterMock) Call(ctx context.Context, method string, path string, secretKey string,
	header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, header, params, result)

	result.(*xendit.Account).ID = "123"
	result.(*xendit.Account).Type = xendit.OWNED
	result.(*xendit.Account).Email = "customer@customer.com"

	return nil
}

func TestCreate(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *account.CreateParams
		expectedRes *xendit.Account
		expectedErr *xendit.Error
	}{
		{
			desc: "should create an account",
			data: &account.CreateParams{
				Email: "customer@customer.com",
				Type:  xendit.OWNED,
			},
			expectedRes: &xendit.Account{
				ID:    "123",
				Email: "customer@customer.com",
				Type:  xendit.OWNED,
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &account.CreateParams{
				Type: xendit.OWNED,
			},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'Email'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"POST",
				xendit.Opt.XenditURL+"/v2/accounts",
				xendit.Opt.SecretKey,
				&http.Header{},
				tC.data,
				&xendit.Account{},
			).Return(nil)

			resp, err := account.Create(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestGet(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *account.GetParams
		expectedRes *xendit.Account
		expectedErr *xendit.Error
	}{
		{
			desc: "should get an account",
			data: &account.GetParams{
				ID: "123",
			},
			expectedRes: &xendit.Account{
				ID:    "123",
				Email: "customer@customer.com",
				Type:  xendit.OWNED,
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &account.GetParams{},
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
				xendit.Opt.XenditURL+"/v2/accounts/123",
				xendit.Opt.SecretKey,
				&http.Header{},
				nil,
				&xendit.Account{},
			).Return(nil)

			resp, err := account.Get(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
