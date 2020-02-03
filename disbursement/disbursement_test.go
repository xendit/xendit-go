package disbursement_test

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
	"github.com/xendit/xendit-go/disbursement"
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
				xendit.Opt.XenditURL+"/disbursements",
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
				xendit.Opt.XenditURL+"/disbursements/"+tC.data.DisbursementID,
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
				xendit.Opt.XenditURL+"/disbursements?"+tC.data.QueryString(),
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
				{
					Name:            "Bank Mandiri",
					Code:            "MANDIRI",
					CanDisburse:     true,
					CanNameValidate: true,
				},
				{
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
				xendit.Opt.XenditURL+"/available_disbursements_banks",
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

type apiRequesterBatchMock struct {
	mock.Mock
}

func (m *apiRequesterBatchMock) Call(ctx context.Context, method string, path string, secretKey string, header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, header, params, result)

	date, _ := time.Parse(time.RFC3339, "2050-01-01T00:00:00.000Z")

	result.(*xendit.BatchDisbursement).Created = &date
	result.(*xendit.BatchDisbursement).Reference = "reference_test"
	result.(*xendit.BatchDisbursement).TotalUploadedAmount = 400000
	result.(*xendit.BatchDisbursement).TotalUploadedCount = 2
	result.(*xendit.BatchDisbursement).Status = "NEEDS_APPROVAL"
	result.(*xendit.BatchDisbursement).ID = "123"

	return nil
}

func TestCreateBatch(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterBatchMock)
	initTesting(apiRequesterMockObj)

	date, _ := time.Parse(time.RFC3339, "2050-01-01T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *disbursement.CreateBatchParams
		expectedRes *xendit.BatchDisbursement
		expectedErr *xendit.Error
	}{
		{
			desc: "should create a batch disbursement",
			data: &disbursement.CreateBatchParams{
				Reference: "reference_test",
				Disbursements: []disbursement.DisbursementItem{
					{
						Amount:            200000,
						BankCode:          "BRI",
						BankAccountName:   "Michael Jackson",
						BankAccountNumber: "1234567890",
						Description:       "Batch disbursement test 1",
					},
					{
						Amount:            200000,
						BankCode:          "BNI",
						BankAccountName:   "Michael Jackson",
						BankAccountNumber: "1234567890",
						Description:       "Batch disbursement test 2",
					},
				},
			},
			expectedRes: &xendit.BatchDisbursement{
				Created:             &date,
				Reference:           "reference_test",
				TotalUploadedAmount: 400000,
				TotalUploadedCount:  2,
				Status:              "NEEDS_APPROVAL",
				ID:                  "123",
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &disbursement.CreateBatchParams{
				Reference: "reference_test",
			},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'Disbursements'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"POST",
				"https://api.xendit.co/batch_disbursements",
				xendit.Opt.SecretKey,
				mock.AnythingOfType("*http.Header"),
				tC.data,
				&xendit.BatchDisbursement{},
			).Return(nil)

			resp, err := disbursement.CreateBatch(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
