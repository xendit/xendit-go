package retailoutlet_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/retailoutlet"
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

	expirationDate, _ := time.Parse(time.RFC3339, "2050-01-01T00:00:00.000Z")

	result.(*xendit.RetailOutlet).IsSingleUse = false
	result.(*xendit.RetailOutlet).Status = "ACTIVE"
	result.(*xendit.RetailOutlet).OwnerID = "someone-owner-id"
	result.(*xendit.RetailOutlet).ExternalID = "retailoutlet-external-id"
	result.(*xendit.RetailOutlet).RetailOutletName = xendit.RetailOutletNameAlfamart
	result.(*xendit.RetailOutlet).Prefix = "TEST"
	result.(*xendit.RetailOutlet).Name = "Michael Jackson"
	result.(*xendit.RetailOutlet).PaymentCode = "TEST123"
	result.(*xendit.RetailOutlet).Type = "USER"
	result.(*xendit.RetailOutlet).ExpectedAmount = 200000
	result.(*xendit.RetailOutlet).ExpirationDate = &expirationDate
	result.(*xendit.RetailOutlet).ID = "123"

	return nil
}

func TestCreateFixedPaymentCode(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	expirationDate, _ := time.Parse(time.RFC3339, "2050-01-01T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *retailoutlet.CreateFixedPaymentCodeParams
		expectedRes *xendit.RetailOutlet
		expectedErr *xendit.Error
	}{
		{
			desc: "should create a retail outlet fixed payment code",
			data: &retailoutlet.CreateFixedPaymentCodeParams{
				ExternalID:       "retailoutlet-external-id",
				RetailOutletName: xendit.RetailOutletNameAlfamart,
				Name:             "Michael Jackson",
				ExpectedAmount:   200000,
			},
			expectedRes: &xendit.RetailOutlet{
				IsSingleUse:      false,
				Status:           "ACTIVE",
				OwnerID:          "someone-owner-id",
				ExternalID:       "retailoutlet-external-id",
				RetailOutletName: xendit.RetailOutletNameAlfamart,
				Prefix:           "TEST",
				Name:             "Michael Jackson",
				PaymentCode:      "TEST123",
				Type:             "USER",
				ExpectedAmount:   200000,
				ExpirationDate:   &expirationDate,
				ID:               "123",
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &retailoutlet.CreateFixedPaymentCodeParams{
				ExternalID:       "retailoutlet-external-id",
				RetailOutletName: xendit.RetailOutletNameAlfamart,
				ExpectedAmount:   200000,
			},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'Name'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"POST",
				"https://api.xendit.co/fixed_payment_code",
				xendit.Opt.SecretKey,
				nil,
				tC.data,
				&xendit.RetailOutlet{},
			).Return(nil)

			resp, err := retailoutlet.CreateFixedPaymentCode(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestGetFixedPaymentCode(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	expirationDate, _ := time.Parse(time.RFC3339, "2050-01-01T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *retailoutlet.GetFixedPaymentCodeParams
		expectedRes *xendit.RetailOutlet
		expectedErr *xendit.Error
	}{
		{
			desc: "should gets a retail outlet fixed payment code",
			data: &retailoutlet.GetFixedPaymentCodeParams{
				FixedPaymentCodeID: "123",
			},
			expectedRes: &xendit.RetailOutlet{
				IsSingleUse:      false,
				Status:           "ACTIVE",
				OwnerID:          "someone-owner-id",
				ExternalID:       "retailoutlet-external-id",
				RetailOutletName: xendit.RetailOutletNameAlfamart,
				Prefix:           "TEST",
				Name:             "Michael Jackson",
				PaymentCode:      "TEST123",
				Type:             "USER",
				ExpectedAmount:   200000,
				ExpirationDate:   &expirationDate,
				ID:               "123",
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &retailoutlet.GetFixedPaymentCodeParams{},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'FixedPaymentCodeID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"GET",
				"https://api.xendit.co/fixed_payment_code/"+tC.data.FixedPaymentCodeID,
				xendit.Opt.SecretKey,
				nil,
				nil,
				&xendit.RetailOutlet{},
			).Return(nil)

			resp, err := retailoutlet.GetFixedPaymentCode(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestUpdateFixedPaymentCode(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	expirationDate, _ := time.Parse(time.RFC3339, "2050-01-01T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *retailoutlet.UpdateFixedPaymentCodeParams
		expectedRes *xendit.RetailOutlet
		expectedErr *xendit.Error
	}{
		{
			desc: "should update a retail outlet fixed payment code",
			data: &retailoutlet.UpdateFixedPaymentCodeParams{
				FixedPaymentCodeID: "123",
				ExpirationDate:     &expirationDate,
				Name:               "Michael Jackson",
				ExpectedAmount:     200000,
			},
			expectedRes: &xendit.RetailOutlet{
				IsSingleUse:      false,
				Status:           "ACTIVE",
				OwnerID:          "someone-owner-id",
				ExternalID:       "retailoutlet-external-id",
				RetailOutletName: xendit.RetailOutletNameAlfamart,
				Prefix:           "TEST",
				Name:             "Michael Jackson",
				PaymentCode:      "TEST123",
				Type:             "USER",
				ExpectedAmount:   200000,
				ExpirationDate:   &expirationDate,
				ID:               "123",
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &retailoutlet.UpdateFixedPaymentCodeParams{
				ExpectedAmount: 200000,
			},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'FixedPaymentCodeID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"PATCH",
				"https://api.xendit.co/fixed_payment_code/"+tC.data.FixedPaymentCodeID,
				xendit.Opt.SecretKey,
				nil,
				tC.data,
				&xendit.RetailOutlet{},
			).Return(nil)

			resp, err := retailoutlet.UpdateFixedPaymentCode(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
