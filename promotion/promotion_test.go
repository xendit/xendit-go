package promotion_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/promotion"
	"github.com/xendit/xendit-go/utils/validator"
)

func initTesting(apiRequesterMockObj xendit.APIRequester) {
	xendit.Opt.SecretKey = "examplesecretkey"
	xendit.SetAPIRequester(apiRequesterMockObj)
}

type apiRequesterPromotionMock struct {
	mock.Mock
}

func (m *apiRequesterPromotionMock) Call(ctx context.Context, method string, path string, secretKey string, header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	mockTime, _ := time.Parse(time.RFC3339, "2020-02-02T00:00:00.000Z")

	result.(*xendit.Promotion).ID = "36ab1517-208a-4f22-b155-96fb101cb378"
	result.(*xendit.Promotion).BusinessID = "5e61664b3dba955c203d232e"
	result.(*xendit.Promotion).ReferenceID = "BRI_20_JAN"
	result.(*xendit.Promotion).Description = "20% discount applied for all BRI cards"
	result.(*xendit.Promotion).Status = "ACTIVE"
	result.(*xendit.Promotion).BinList = []string{"400000", "460000"}
	result.(*xendit.Promotion).DiscountPercent = 20
	result.(*xendit.Promotion).Currency = "IDR"
	result.(*xendit.Promotion).ChannelCode = "BRI"
	result.(*xendit.Promotion).StartTime = &mockTime
	result.(*xendit.Promotion).EndTime = &mockTime
	result.(*xendit.Promotion).MinOriginalAmount = 25000
	result.(*xendit.Promotion).MaxDiscountAmount = 5000

	return nil
}

func TestCreatePromotion(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterPromotionMock)
	initTesting(apiRequesterMockObj)

	mockTime, _ := time.Parse(time.RFC3339, "2020-02-02T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *promotion.CreatePromotionParams
		expectedRes *xendit.Promotion
		expectedErr *xendit.Error
	}{
		{
			desc: "should create a promotion",
			data: &promotion.CreatePromotionParams{
				ReferenceID:       "BRI_20_JAN",
				Description:       "20% discount applied for all BRI cards",
				BinList:           []string{"400000", "460000"},
				DiscountPercent:   20,
				Currency:          "IDR",
				ChannelCode:       "BRI",
				StartTime:         &mockTime,
				EndTime:           &mockTime,
				MinOriginalAmount: 25000,
				MaxDiscountAmount: 5000,
			},
			expectedRes: &xendit.Promotion{
				ID:                "36ab1517-208a-4f22-b155-96fb101cb378",
				BusinessID:        "5e61664b3dba955c203d232e",
				ReferenceID:       "BRI_20_JAN",
				Description:       "20% discount applied for all BRI cards",
				Status:            "ACTIVE",
				BinList:           []string{"400000", "460000"},
				DiscountPercent:   20,
				Currency:          "IDR",
				ChannelCode:       "BRI",
				StartTime:         &mockTime,
				EndTime:           &mockTime,
				MinOriginalAmount: 25000,
				MaxDiscountAmount: 5000,
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &promotion.CreatePromotionParams{
				Description:       "20% discount applied for all BRI cards",
				BinList:           []string{"400000", "460000"},
				DiscountPercent:   20,
				Currency:          "IDR",
				ChannelCode:       "BRI",
				StartTime:         &mockTime,
				EndTime:           &mockTime,
				MinOriginalAmount: 25000,
				MaxDiscountAmount: 5000,
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
				xendit.Opt.XenditURL+"/promotions",
				xendit.Opt.SecretKey,
				nil,
				tC.data,
				&xendit.Promotion{},
			).Return(nil)

			resp, err := promotion.CreatePromotion(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

type apiRequesterGetPromotionsMock struct {
	mock.Mock
}

func (m *apiRequesterGetPromotionsMock) Call(ctx context.Context, method string, path string, secretKey string, header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	resultString := `[
		{
			"id": "36ab1517-208a-4f22-b155-96fb101cb378",
			"business_id": "5e61664b3dba955c203d232e",
			"reference_id": "BRI_20_JAN",
			"status": "ACTIVE",
			"description": "20% discount applied for all BRI cards",
			"bin_list": [
				"400000",
				"460000"
			],
			"channel_code": "BRI",
			"discount_percent": 20,
			"currency": "IDR",
			"start_time": "2020-01-01T00:00:00.000Z",
			"end_time": "2020-01-01T00:00:00.000Z",
			"min_original_amount": 25000,
			"max_discount_amount": 5000
		},
		{
			"id": "36ab1517-208a-4f22-b155-96fb101cb377",
			"business_id": "5e61664b3dba955c203d232e",
			"reference_id": "BRI_20_FEB",
			"status": "ACTIVE",
			"description": "20% discount applied for all BRI cards",
			"start_time": "2020-01-01T00:00:00.000Z",
			"end_time": "2020-01-01T00:00:00.000Z",
			"bin_list": [
				"400000",
				"460000"
			],
			"discount_percent": 20,
			"channel_code": "BRI",
			"currency": "IDR",
			"min_original_amount": 25000,
			"max_discount_amount": 5000
		}
	]`

	_ = json.Unmarshal([]byte(resultString), &result)

	return nil
}

func TestGetPromotions(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterGetPromotionsMock)
	initTesting(apiRequesterMockObj)

	mockTime, _ := time.Parse(time.RFC3339, "2020-01-01T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *promotion.GetPromotionsParams
		expectedRes []xendit.Promotion
		expectedErr *xendit.Error
	}{
		{
			desc: "should get promotions",
			data: &promotion.GetPromotionsParams{
				Currency:    "IDR",
				ChannelCode: "BRI",
			},
			expectedRes: []xendit.Promotion{
				{
					ID:                "36ab1517-208a-4f22-b155-96fb101cb378",
					BusinessID:        "5e61664b3dba955c203d232e",
					ReferenceID:       "BRI_20_JAN",
					Description:       "20% discount applied for all BRI cards",
					Status:            "ACTIVE",
					BinList:           []string{"400000", "460000"},
					DiscountPercent:   20,
					Currency:          "IDR",
					ChannelCode:       "BRI",
					StartTime:         &mockTime,
					EndTime:           &mockTime,
					MinOriginalAmount: 25000,
					MaxDiscountAmount: 5000,
				},
				{
					ID:                "36ab1517-208a-4f22-b155-96fb101cb377",
					BusinessID:        "5e61664b3dba955c203d232e",
					ReferenceID:       "BRI_20_FEB",
					Description:       "20% discount applied for all BRI cards",
					Status:            "ACTIVE",
					BinList:           []string{"400000", "460000"},
					DiscountPercent:   20,
					Currency:          "IDR",
					ChannelCode:       "BRI",
					StartTime:         &mockTime,
					EndTime:           &mockTime,
					MinOriginalAmount: 25000,
					MaxDiscountAmount: 5000,
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
				fmt.Sprintf("%s/promotions?%s", xendit.Opt.XenditURL, tC.data.QueryString()),
				xendit.Opt.SecretKey,
				nil,
				nil,
				&[]xendit.Promotion{},
			).Return(nil)

			resp, err := promotion.GetPromotions(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
