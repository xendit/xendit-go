package disbursementph_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/disbursementph"
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

	result.(*xendit.DisbursementPh).ID = "123"
	result.(*xendit.DisbursementPh).ReferenceID = "disbursement-reference-id"
	result.(*xendit.DisbursementPh).Currency = "currency"
	result.(*xendit.DisbursementPh).Amount = 200000
	result.(*xendit.DisbursementPh).ChannelCode = "PH_BDO"
	result.(*xendit.DisbursementPh).Description = "test disbursement"
	result.(*xendit.DisbursementPh).Status = "PENDING"
	result.(*xendit.DisbursementPh).Currency = "PHP"

	return nil
}

func TestCreate(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *disbursementph.CreateParams
		expectedRes *xendit.DisbursementPh
		expectedErr *xendit.Error
	}{
		{
			desc: "should create a disbursement",
			data: &disbursementph.CreateParams{
				IdempotencyKey: "test",
				ReferenceID:    "disbursement-reference-id",
				ChannelCode:    "PH_BDO",
				AccountName:    "Michael Jackson",
				AccountNumber:  "1234567890",
				Description:    "test disbursement",
				Amount:         200000,
				Currency:       "PHP",
			},
			expectedRes: &xendit.DisbursementPh{
				ID:          "123",
				ReferenceID: "disbursement-reference-id",
				Amount:      200000,
				ChannelCode: "PH_BDO",
				Description: "test disbursement",
				Status:      "PENDING",
				Currency:    "PHP",
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &disbursementph.CreateParams{
				IdempotencyKey: "test",
				ReferenceID:    "disbursement-reference-id",
				ChannelCode:    "PH_BDO",
				AccountName:    "Michael Jackson",
				AccountNumber:  "1234567890",
				Description:    "test disbursement",
				Currency:       "PHP",
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
				mock.AnythingOfType("http.Header"),
				tC.data,
				&xendit.DisbursementPh{},
			).Return(nil)

			resp, err := disbursementph.Create(tC.data)

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
		data        *disbursementph.GetByIDParams
		expectedRes *xendit.DisbursementPh
		expectedErr *xendit.Error
	}{
		{
			desc: "should get an disbursement",
			data: &disbursementph.GetByIDParams{
				DisbursementID: "123",
			},
			expectedRes: &xendit.DisbursementPh{
				ID:          "123",
				ReferenceID: "disbursement-reference-id",
				Currency:    "PHP",
				Amount:      200000,
				ChannelCode: "PH_BDO",
				Description: "test disbursement",
				Status:      "PENDING",
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &disbursementph.GetByIDParams{},
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
				mock.AnythingOfType("http.Header"),
				nil,
				&xendit.DisbursementPh{},
			).Return(nil)

			resp, err := disbursementph.GetByID(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

type apiRequesterGetByRefernceIDMock struct {
	mock.Mock
}

func (m *apiRequesterGetByRefernceIDMock) Call(ctx context.Context, method string, path string, secretKey string, header http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	resultString := `[{
        "status": "PENDING",
        "reference_id": "disbursement-reference-id",
        "amount": 200000,
        "channel_code": "PH_BDO",
        "description": "test disbursement",
        "id": "123",
		"currency":"PHP"
	},
	{
        "status": "PENDING",
        "user_id": "user-123",
        "reference_id": "disbursement-reference-id",
        "amount": 200000,
        "channel_code": "PH_BDO",
        "description": "test disbursement",
        "id": "124",
		"currency":"PHP"
    }]`

	_ = json.Unmarshal([]byte(resultString), &result)

	return nil
}

func TestGetByExternalID(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterGetByRefernceIDMock)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *disbursementph.GetByReferenceIDParams
		expectedRes []xendit.DisbursementPh
		expectedErr *xendit.Error
	}{
		{
			desc: "should get a list of disbursements",
			data: &disbursementph.GetByReferenceIDParams{
				ReferenceID: "disbursement-reference-id",
			},
			expectedRes: []xendit.DisbursementPh{
				{
					ID:          "123",
					ReferenceID: "disbursement-reference-id",
					Amount:      200000,
					ChannelCode: "PH_BDO",
					Description: "test disbursement",
					Status:      "PENDING",
					Currency:    "PHP",
				},
				{
					ID:          "124",
					ReferenceID: "disbursement-reference-id",
					Amount:      200000,
					ChannelCode: "PH_BDO",
					Description: "test disbursement",
					Status:      "PENDING",
					Currency:    "PHP",
				},
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &disbursementph.GetByReferenceIDParams{},
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
				xendit.Opt.XenditURL+"/disbursements?"+tC.data.QueryString(),
				xendit.Opt.SecretKey,
				nil,
				nil,
				&[]xendit.DisbursementPh{},
			).Return(nil)

			resp, err := disbursementph.GetByReferenceID(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

type apiRequesterGetDisbursementChannelssMock struct {
	mock.Mock
}

func (m *apiRequesterGetDisbursementChannelssMock) Call(ctx context.Context, method string, path string, secretKey string, header http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	resultString := `[
		{
			"channel_code": "PH_ONB",
			"channel_category": "BANK",
			"currency": "PHP",
			"name": "BDO Network Bank",
			"amount_limits": {
				"minimum": 1,
				"maximum": 50000,
				"minimum_increment": 0.01
			}
		},
		{
			"channel_code": "PH_PAR",
			"channel_category": "BANK",
			"currency": "PHP",
			"name": "Partner Rural Bank (Cotabato), Inc.",
			"amount_limits": {
				"minimum": 50,
				"maximum": 50000,
				"minimum_increment": 0.01
			}
		}
	]`

	_ = json.Unmarshal([]byte(resultString), &result)

	return nil
}

func TestGetDisbursementChannels(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterGetDisbursementChannelssMock)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		expectedRes []xendit.DisbursementChannel
		expectedErr *xendit.Error
	}{
		{
			desc: "should get available disbursement channels",
			expectedRes: []xendit.DisbursementChannel{
				{
					ChannelCode:     "PH_ONB",
					ChannelCategory: "BANK",
					Currency:        "PHP",
					Name:            "BDO Network Bank",
					AmountLimits: xendit.AmountLimits{
						Minimum:          1,
						Maximum:          50000,
						MinimumIncrement: 0.01,
					},
				},
				{
					ChannelCode:     "PH_PAR",
					ChannelCategory: "BANK",
					Currency:        "PHP",
					Name:            "Partner Rural Bank (Cotabato), Inc.",
					AmountLimits: xendit.AmountLimits{
						Minimum:          50,
						Maximum:          50000,
						MinimumIncrement: 0.01,
					},
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
				xendit.Opt.XenditURL+"/disbursement-channels",
				xendit.Opt.SecretKey,
				nil,
				nil,
				&[]xendit.DisbursementChannel{},
			).Return(nil)

			resp, err := disbursementph.GetDisbursementChannels()

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

type apiRequesterGetBankChannelsByChannelCategorysMock struct {
	mock.Mock
}

func (m *apiRequesterGetBankChannelsByChannelCategorysMock) Call(ctx context.Context, method string, path string, secretKey string, header http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	resultString := `[
		{
			"channel_code": "PH_ONB",
			"channel_category": "BANK",
			"currency": "PHP",
			"name": "BDO Network Bank",
			"amount_limits": {
				"minimum": 1,
				"maximum": 50000,
				"minimum_increment": 0.01
			}
		},
		{
			"channel_code": "PH_PAR",
			"channel_category": "BANK",
			"currency": "PHP",
			"name": "Partner Rural Bank (Cotabato), Inc.",
			"amount_limits": {
				"minimum": 50,
				"maximum": 50000,
				"minimum_increment": 0.01
			}
		}
	]`

	_ = json.Unmarshal([]byte(resultString), &result)

	return nil
}

func TestGetDisbursementChannelsByChannelCategory(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterGetBankChannelsByChannelCategorysMock)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        disbursementph.GetByChannelCategoryParams
		expectedRes []xendit.DisbursementChannel
		expectedErr *xendit.Error
	}{
		{
			desc: "should get available disbursement channels by Channel Category",
			data: disbursementph.GetByChannelCategoryParams{
				ChannelCategory: "BANK",
			},
			expectedRes: []xendit.DisbursementChannel{
				{
					ChannelCode:     "PH_ONB",
					ChannelCategory: "BANK",
					Currency:        "PHP",
					Name:            "BDO Network Bank",
					AmountLimits: xendit.AmountLimits{
						Minimum:          1,
						Maximum:          50000,
						MinimumIncrement: 0.01,
					},
				},
				{
					ChannelCode:     "PH_PAR",
					ChannelCategory: "BANK",
					Currency:        "PHP",
					Name:            "Partner Rural Bank (Cotabato), Inc.",
					AmountLimits: xendit.AmountLimits{
						Minimum:          50,
						Maximum:          50000,
						MinimumIncrement: 0.01,
					},
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
				xendit.Opt.XenditURL+"/disbursement-channels?"+tC.data.QueryString(),
				xendit.Opt.SecretKey,
				nil,
				nil,
				&[]xendit.DisbursementChannel{},
			).Return(nil)

			resp, err := disbursementph.GetDisbursementChannelsByChannelCategory(&tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

type apiRequesterGetBankChannelsByChannelCodeMock struct {
	mock.Mock
}

func (m *apiRequesterGetBankChannelsByChannelCodeMock) Call(ctx context.Context, method string, path string, secretKey string, header http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	resultString := `[
		{
			"channel_code": "PH_ONB",
			"channel_category": "BANK",
			"currency": "PHP",
			"name": "BDO Network Bank",
			"amount_limits": {
				"minimum": 1,
				"maximum": 50000,
				"minimum_increment": 0.01
			}
		},
		{
			"channel_code": "PH_PAR",
			"channel_category": "BANK",
			"currency": "PHP",
			"name": "Partner Rural Bank (Cotabato), Inc.",
			"amount_limits": {
				"minimum": 50,
				"maximum": 50000,
				"minimum_increment": 0.01
			}
		}
	]`

	_ = json.Unmarshal([]byte(resultString), &result)

	return nil
}

func TestGetDisbursementChannelsByChannelCode(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterGetBankChannelsByChannelCodeMock)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        disbursementph.GetByChannelCodeParams
		expectedRes []xendit.DisbursementChannel
		expectedErr *xendit.Error
	}{
		{
			desc: "should get available disbursement channels by Channel Code",
			data: disbursementph.GetByChannelCodeParams{
				ChannelCode: "PH_BDP",
			},
			expectedRes: []xendit.DisbursementChannel{
				{
					ChannelCode:     "PH_ONB",
					ChannelCategory: "BANK",
					Currency:        "PHP",
					Name:            "BDO Network Bank",
					AmountLimits: xendit.AmountLimits{
						Minimum:          1,
						Maximum:          50000,
						MinimumIncrement: 0.01,
					},
				},
				{
					ChannelCode:     "PH_PAR",
					ChannelCategory: "BANK",
					Currency:        "PHP",
					Name:            "Partner Rural Bank (Cotabato), Inc.",
					AmountLimits: xendit.AmountLimits{
						Minimum:          50,
						Maximum:          50000,
						MinimumIncrement: 0.01,
					},
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
				xendit.Opt.XenditURL+"/disbursement-channels?"+tC.data.QueryString(),
				xendit.Opt.SecretKey,
				nil,
				nil,
				&[]xendit.DisbursementChannel{},
			).Return(nil)

			resp, err := disbursementph.GetDisbursementChannelsByChannelCode(&tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
