package recurringpayment_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/recurringpayment"
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
	m.Called(ctx, method, path, secretKey, nil, params, result)

	date, _ := time.Parse(time.RFC3339, "2050-01-01T00:00:00.000Z")

	result.(*xendit.RecurringPayment).ID = "123"
	result.(*xendit.RecurringPayment).UserID = "someone-id"
	result.(*xendit.RecurringPayment).ExternalID = "recurring-external-id"
	result.(*xendit.RecurringPayment).Status = "ACTIVE"
	result.(*xendit.RecurringPayment).Amount = 200000
	result.(*xendit.RecurringPayment).PayerEmail = "customer@customer.com"
	result.(*xendit.RecurringPayment).Description = "recurring test"
	result.(*xendit.RecurringPayment).Interval = "DAY"
	result.(*xendit.RecurringPayment).IntervalCount = 3
	result.(*xendit.RecurringPayment).RecurrenceProgress = 1
	result.(*xendit.RecurringPayment).ShouldSendEmail = false
	result.(*xendit.RecurringPayment).MissedPaymentAction = "IGNORE"
	result.(*xendit.RecurringPayment).LastCreatedInvoiceURL = "https://invoice-url.com"
	result.(*xendit.RecurringPayment).Created = &date
	result.(*xendit.RecurringPayment).Updated = &date
	result.(*xendit.RecurringPayment).StartDate = &date
	result.(*xendit.RecurringPayment).Recharge = false

	return nil
}

func TestCreateRecurringPayment(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	date, _ := time.Parse(time.RFC3339, "2050-01-01T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *recurringpayment.CreateParams
		expectedRes *xendit.RecurringPayment
		expectedErr *xendit.Error
	}{
		{
			desc: "should create a recurring payment",
			data: &recurringpayment.CreateParams{
				ExternalID:    "recurring-external-id",
				Amount:        200000,
				PayerEmail:    "customer@customer.com",
				Description:   "recurring test",
				Interval:      xendit.RecurringPaymentIntervalDay,
				IntervalCount: 3,
				Recharge:      new(bool),
			},
			expectedRes: &xendit.RecurringPayment{
				ID:                    "123",
				UserID:                "someone-id",
				ExternalID:            "recurring-external-id",
				Status:                "ACTIVE",
				Amount:                200000,
				PayerEmail:            "customer@customer.com",
				Description:           "recurring test",
				Interval:              "DAY",
				IntervalCount:         3,
				RecurrenceProgress:    1,
				ShouldSendEmail:       false,
				MissedPaymentAction:   "IGNORE",
				LastCreatedInvoiceURL: "https://invoice-url.com",
				Created:               &date,
				Updated:               &date,
				StartDate:             &date,
				Recharge:              false,
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &recurringpayment.CreateParams{
				ExternalID:  "recurring-external-id",
				Amount:      200000,
				PayerEmail:  "customer@customer.com",
				Description: "recurring test",
				Interval:    xendit.RecurringPaymentIntervalDay,
			},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'IntervalCount'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"POST",
				"https://api.xendit.co/recurring_payments",
				xendit.Opt.SecretKey,
				nil,
				tC.data,
				&xendit.RecurringPayment{},
			).Return(nil)

			resp, err := recurringpayment.Create(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestGetRecurringPayment(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	date, _ := time.Parse(time.RFC3339, "2050-01-01T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *recurringpayment.GetParams
		expectedRes *xendit.RecurringPayment
		expectedErr *xendit.Error
	}{
		{
			desc: "should get a recurring payment",
			data: &recurringpayment.GetParams{
				ID: "123",
			},
			expectedRes: &xendit.RecurringPayment{
				ID:                    "123",
				UserID:                "someone-id",
				ExternalID:            "recurring-external-id",
				Status:                "ACTIVE",
				Amount:                200000,
				PayerEmail:            "customer@customer.com",
				Description:           "recurring test",
				Interval:              "DAY",
				IntervalCount:         3,
				RecurrenceProgress:    1,
				ShouldSendEmail:       false,
				MissedPaymentAction:   "IGNORE",
				LastCreatedInvoiceURL: "https://invoice-url.com",
				Created:               &date,
				Updated:               &date,
				StartDate:             &date,
				Recharge:              false,
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &recurringpayment.GetParams{},
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
				"https://api.xendit.co/recurring_payments/"+tC.data.ID,
				xendit.Opt.SecretKey,
				nil,
				nil,
				&xendit.RecurringPayment{},
			).Return(nil)

			resp, err := recurringpayment.Get(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestEditRecurringPayment(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	date, _ := time.Parse(time.RFC3339, "2050-01-01T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *recurringpayment.EditParams
		expectedRes *xendit.RecurringPayment
		expectedErr *xendit.Error
	}{
		{
			desc: "should edit a recurring payment",
			data: &recurringpayment.EditParams{
				ID:            "123",
				Amount:        200000,
				Interval:      xendit.RecurringPaymentIntervalDay,
				IntervalCount: 3,
			},
			expectedRes: &xendit.RecurringPayment{
				ID:                    "123",
				UserID:                "someone-id",
				ExternalID:            "recurring-external-id",
				Status:                "ACTIVE",
				Amount:                200000,
				PayerEmail:            "customer@customer.com",
				Description:           "recurring test",
				Interval:              "DAY",
				IntervalCount:         3,
				RecurrenceProgress:    1,
				ShouldSendEmail:       false,
				MissedPaymentAction:   "IGNORE",
				LastCreatedInvoiceURL: "https://invoice-url.com",
				Created:               &date,
				Updated:               &date,
				StartDate:             &date,
				Recharge:              false,
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &recurringpayment.EditParams{},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'ID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"PATCH",
				"https://api.xendit.co/recurring_payments/"+tC.data.ID,
				xendit.Opt.SecretKey,
				nil,
				tC.data,
				&xendit.RecurringPayment{},
			).Return(nil)

			resp, err := recurringpayment.Edit(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestStopRecurringPayment(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	date, _ := time.Parse(time.RFC3339, "2050-01-01T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *recurringpayment.StopParams
		expectedRes *xendit.RecurringPayment
		expectedErr *xendit.Error
	}{
		{
			desc: "should stop a recurring payment",
			data: &recurringpayment.StopParams{
				ID: "123",
			},
			expectedRes: &xendit.RecurringPayment{
				ID:                    "123",
				UserID:                "someone-id",
				ExternalID:            "recurring-external-id",
				Status:                "ACTIVE",
				Amount:                200000,
				PayerEmail:            "customer@customer.com",
				Description:           "recurring test",
				Interval:              "DAY",
				IntervalCount:         3,
				RecurrenceProgress:    1,
				ShouldSendEmail:       false,
				MissedPaymentAction:   "IGNORE",
				LastCreatedInvoiceURL: "https://invoice-url.com",
				Created:               &date,
				Updated:               &date,
				StartDate:             &date,
				Recharge:              false,
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &recurringpayment.StopParams{},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'ID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"POST",
				"https://api.xendit.co/recurring_payments/"+tC.data.ID+"/stop!",
				xendit.Opt.SecretKey,
				nil,
				nil,
				&xendit.RecurringPayment{},
			).Return(nil)

			resp, err := recurringpayment.Stop(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestPauseRecurringPayment(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	date, _ := time.Parse(time.RFC3339, "2050-01-01T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *recurringpayment.PauseParams
		expectedRes *xendit.RecurringPayment
		expectedErr *xendit.Error
	}{
		{
			desc: "should pause a recurring payment",
			data: &recurringpayment.PauseParams{
				ID: "123",
			},
			expectedRes: &xendit.RecurringPayment{
				ID:                    "123",
				UserID:                "someone-id",
				ExternalID:            "recurring-external-id",
				Status:                "ACTIVE",
				Amount:                200000,
				PayerEmail:            "customer@customer.com",
				Description:           "recurring test",
				Interval:              "DAY",
				IntervalCount:         3,
				RecurrenceProgress:    1,
				ShouldSendEmail:       false,
				MissedPaymentAction:   "IGNORE",
				LastCreatedInvoiceURL: "https://invoice-url.com",
				Created:               &date,
				Updated:               &date,
				StartDate:             &date,
				Recharge:              false,
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &recurringpayment.PauseParams{},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'ID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"POST",
				"https://api.xendit.co/recurring_payments/"+tC.data.ID+"/pause!",
				xendit.Opt.SecretKey,
				nil,
				nil,
				&xendit.RecurringPayment{},
			).Return(nil)

			resp, err := recurringpayment.Pause(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestResumeRecurringPayment(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	date, _ := time.Parse(time.RFC3339, "2050-01-01T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *recurringpayment.ResumeParams
		expectedRes *xendit.RecurringPayment
		expectedErr *xendit.Error
	}{
		{
			desc: "should resume a recurring payment",
			data: &recurringpayment.ResumeParams{
				ID: "123",
			},
			expectedRes: &xendit.RecurringPayment{
				ID:                    "123",
				UserID:                "someone-id",
				ExternalID:            "recurring-external-id",
				Status:                "ACTIVE",
				Amount:                200000,
				PayerEmail:            "customer@customer.com",
				Description:           "recurring test",
				Interval:              "DAY",
				IntervalCount:         3,
				RecurrenceProgress:    1,
				ShouldSendEmail:       false,
				MissedPaymentAction:   "IGNORE",
				LastCreatedInvoiceURL: "https://invoice-url.com",
				Created:               &date,
				Updated:               &date,
				StartDate:             &date,
				Recharge:              false,
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &recurringpayment.ResumeParams{},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'ID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"POST",
				"https://api.xendit.co/recurring_payments/"+tC.data.ID+"/resume!",
				xendit.Opt.SecretKey,
				nil,
				nil,
				&xendit.RecurringPayment{},
			).Return(nil)

			resp, err := recurringpayment.Resume(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
