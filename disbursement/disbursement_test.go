package disbursement_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/disbursement"
	"github.com/xendit/xendit-go/utils/validator"
)

func initTesting(apiRequesterMockObj xendit.APIRequester) {
	xendit.Opt.SecretKey = "xnd_development_REt02KJzkM6AootfKnDrMw1Sse4LlzEDHfKzXoBocqIEiH4bqjHUJXbl6Cfaab"
	xendit.SetAPIRequester(apiRequesterMockObj)
}

type apiRequesterMock struct {
	mock.Mock
}

func (m *apiRequesterMock) Call(ctx context.Context, method string, path string, secretKey string, header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, header, params, result)

	result.(*xendit.Disbursement).ID = "123"
	result.(*xendit.Disbursement).UserID = "user-123"
	result.(*xendit.Disbursement).ExternalID = "disbursement-external-id"
	result.(*xendit.Disbursement).Amount = 200000
	result.(*xendit.Disbursement).BankCode = "BRI"
	result.(*xendit.Disbursement).AccountHolderName = "Michael Jackson"
	result.(*xendit.Disbursement).DisbursementDescription = "test disbursement"
	result.(*xendit.Disbursement).Status = "PENDING"

	return nil
}

func TestCreate(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *disbursement.CreateParams
		expectedRes *xendit.Disbursement
		expectedErr *xendit.Error
	}{
		{
			desc: "should create a disbursement",
			data: &disbursement.CreateParams{
				ExternalID:        "disbursement-external-id",
				BankCode:          "BRI",
				AccountHolderName: "Michael Jackson",
				AccountNumber:     "1234567890",
				Description:       "test disbursement",
				Amount:            200000,
			},
			expectedRes: &xendit.Disbursement{
				ID:                      "123",
				UserID:                  "user-123",
				ExternalID:              "disbursement-external-id",
				Amount:                  200000,
				BankCode:                "BRI",
				AccountHolderName:       "Michael Jackson",
				DisbursementDescription: "test disbursement",
				Status:                  "PENDING",
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &disbursement.CreateParams{
				ExternalID:        "disbursement-external-id",
				BankCode:          "BRI",
				AccountHolderName: "Michael Jackson",
				AccountNumber:     "1234567890",
				Description:       "test disbursement",
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
				"https://api.xendit.co/disbursements",
				xendit.Opt.SecretKey,
				mock.AnythingOfType("*http.Header"),
				tC.data,
				&xendit.Disbursement{},
			).Return(nil)

			resp, err := disbursement.Create(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestGetByID(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *disbursement.GetByIDParams
		expectedRes *xendit.Disbursement
		expectedErr *xendit.Error
	}{
		{
			desc: "should get an disbursement",
			data: &disbursement.GetByIDParams{
				DisbursementID: "123",
			},
			expectedRes: &xendit.Disbursement{
				ID:                      "123",
				UserID:                  "user-123",
				ExternalID:              "disbursement-external-id",
				Amount:                  200000,
				BankCode:                "BRI",
				AccountHolderName:       "Michael Jackson",
				DisbursementDescription: "test disbursement",
				Status:                  "PENDING",
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &disbursement.GetByIDParams{},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'DisbursementID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"GET",
				"https://api.xendit.co/disbursements/"+tC.data.DisbursementID,
				xendit.Opt.SecretKey,
				mock.AnythingOfType("*http.Header"),
				nil,
				&xendit.Disbursement{},
			).Return(nil)

			resp, err := disbursement.GetByID(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

type apiRequesterGetByExternalIDMock struct {
	mock.Mock
}

func (m *apiRequesterGetByExternalIDMock) Call(ctx context.Context, method string, path string, secretKey string, header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	resultString := `[{
        "status": "PENDING",
        "user_id": "user-123",
        "external_id": "disbursement-external-id",
        "amount": 200000,
        "bank_code": "BRI",
        "account_holder_name": "Michael Jackson",
        "disbursement_description": "test disbursement",
        "id": "123"
	},
	{
        "status": "PENDING",
        "user_id": "user-123",
        "external_id": "disbursement-external-id",
        "amount": 200000,
        "bank_code": "BRI",
        "account_holder_name": "Michael Jackson",
        "disbursement_description": "test disbursement",
        "id": "124"
    }]`

	_ = json.Unmarshal([]byte(resultString), &result)

	return nil
}

func TestGetByExternalID(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterGetByExternalIDMock)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *disbursement.GetByExternalIDParams
		expectedRes []xendit.Disbursement
		expectedErr *xendit.Error
	}{
		{
			desc: "should get a list of disbursements",
			data: &disbursement.GetByExternalIDParams{
				ExternalID: "disbursement-external-id",
			},
			expectedRes: []xendit.Disbursement{
				{
					ID:                      "123",
					UserID:                  "user-123",
					ExternalID:              "disbursement-external-id",
					Amount:                  200000,
					BankCode:                "BRI",
					AccountHolderName:       "Michael Jackson",
					DisbursementDescription: "test disbursement",
					Status:                  "PENDING",
				},
				{
					ID:                      "124",
					UserID:                  "user-123",
					ExternalID:              "disbursement-external-id",
					Amount:                  200000,
					BankCode:                "BRI",
					AccountHolderName:       "Michael Jackson",
					DisbursementDescription: "test disbursement",
					Status:                  "PENDING",
				},
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &disbursement.GetByExternalIDParams{},
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
				"https://api.xendit.co/disbursements?"+tC.data.QueryString(),
				xendit.Opt.SecretKey,
				nil,
				nil,
				&[]xendit.Disbursement{},
			).Return(nil)

			resp, err := disbursement.GetByExternalID(tC.data)

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
			"code": "MANDIRI",
			"can_disburse": true,
			"can_name_validate": true
		},
		{
			"name": "Bank Negara Indonesia",
			"code": "BNI",
			"can_disburse": true,
			"can_name_validate": true
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
		expectedRes []xendit.DisbursementBank
		expectedErr *xendit.Error
	}{
		{
			desc: "should get available disbursement banks",
			expectedRes: []xendit.DisbursementBank{
				xendit.DisbursementBank{
					Name:            "Bank Mandiri",
					Code:            "MANDIRI",
					CanDisburse:     true,
					CanNameValidate: true,
				},
				xendit.DisbursementBank{
					Name:            "Bank Negara Indonesia",
					Code:            "BNI",
					CanDisburse:     true,
					CanNameValidate: true,
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
				"https://api.xendit.co/available_disbursements_banks",
				xendit.Opt.SecretKey,
				nil,
				nil,
				&[]xendit.DisbursementBank{},
			).Return(nil)

			resp, err := disbursement.GetAvailableBanks()

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
