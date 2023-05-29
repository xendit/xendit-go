package report

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/utils/validator"
	"net/http"
	"testing"
	"time"
)

func initTesting(apiRequesterMockObj xendit.APIRequester) {
	xendit.Opt.SecretKey = "examplesecretkey"
	xendit.SetAPIRequester(apiRequesterMockObj)
}

type apiRequesterMock struct {
	mock.Mock
}

func (m *apiRequesterMock) Call(ctx context.Context, method string, url string, secretKey string, header http.Header, body interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, url, secretKey, header, body, result)

	resultString := `{
		"id": "report_5c1b34a2-6ceb-4c24-aba9-c836bac82b28",
		"type": "BALANCE_HISTORY",
		"status": "COMPLETED",
		"filter": {
			"from": "2021-06-23T04:01:55.574Z",
			"to": "2021-06-24T04:01:55.574Z"
		},
		"format": "CSV",
		"url": "https://transaction-report-files.s3-us-west-2.amazonaws.com/{report_name}",
		"currency": "IDR",
		"business_id": "5f34f60535ba7c1c0eed846a",
		"created": "2021-06-24T04:01:55.570Z",
		"updated": "2021-06-24T04:01:55.570Z"
	}`

	_ = json.Unmarshal([]byte(resultString), &result)

	return nil
}

func TestGenerateReport(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	created := time.Date(2021, 6, 24, 04, 01, 55, 570000000, time.UTC)
	updated := time.Date(2021, 6, 24, 04, 01, 55, 570000000, time.UTC)

	testCases := []struct {
		desc        string
		data        *GenerateReportParams
		expectedRes *xendit.Report
		expectedErr *xendit.Error
	}{
		{
			desc: "should generate report",
			data: &GenerateReportParams{
				Type: "BALANCE_HISTORY", // "BALANCE_HISTORY", "TRANSACTIONS", "UPCOMING_TRANSACTIONS"
				Filter: Filter{
					From: "2021-06-23T04:01:55.574Z",
					To:   "2021-06-24T04:01:55.574Z",
				},
			},
			expectedRes: &xendit.Report{
				ID:     "report_5c1b34a2-6ceb-4c24-aba9-c836bac82b28",
				Type:   "BALANCE_HISTORY",
				Status: "COMPLETED",
				Filter: xendit.Filter{
					From: "2021-06-23T04:01:55.574Z",
					To:   "2021-06-24T04:01:55.574Z",
				},
				Format:     "CSV",
				Url:        "https://transaction-report-files.s3-us-west-2.amazonaws.com/{report_name}",
				Currency:   "IDR",
				BusinessID: "5f34f60535ba7c1c0eed846a",
				Created:    created,
				Updated:    updated,
			},
		},
		{
			desc:        "should report error",
			data:        &GenerateReportParams{},
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'Type'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On("Call",
				context.Background(),
				"POST",
				xendit.Opt.XenditURL+"/reports",
				xendit.Opt.SecretKey,
				mock.Anything,
				mock.Anything,
				&xendit.Report{},
			).Return(nil)

			resp, err := GenerateReport(tC.data)
			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestGetReport(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	created := time.Date(2021, 6, 24, 04, 01, 55, 570000000, time.UTC)
	updated := time.Date(2021, 6, 24, 04, 01, 55, 570000000, time.UTC)

	testCases := []struct {
		desc        string
		data        *GetReportParams
		expectedRes *xendit.Report
		expectedErr *xendit.Error
	}{
		{
			desc: "should generate report",
			data: &GetReportParams{
				ReportID: "report_5c1b34a2-6ceb-4c24-aba9-c836bac82b28",
			},
			expectedRes: &xendit.Report{
				ID:     "report_5c1b34a2-6ceb-4c24-aba9-c836bac82b28",
				Type:   "BALANCE_HISTORY",
				Status: "COMPLETED",
				Filter: xendit.Filter{
					From: "2021-06-23T04:01:55.574Z",
					To:   "2021-06-24T04:01:55.574Z",
				},
				Format:     "CSV",
				Url:        "https://transaction-report-files.s3-us-west-2.amazonaws.com/{report_name}",
				Currency:   "IDR",
				BusinessID: "5f34f60535ba7c1c0eed846a",
				Created:    created,
				Updated:    updated,
			},
		},
		{
			desc:        "should report error",
			data:        &GetReportParams{},
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'ReportID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On("Call",
				context.Background(),
				"GET",
				xendit.Opt.XenditURL+"/reports/"+tC.data.ReportID,
				xendit.Opt.SecretKey,
				mock.Anything,
				mock.Anything,
				&xendit.Report{},
			).Return(nil)

			resp, err := GetReport(tC.data)
			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
