package qrcode_test

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
	"github.com/xendit/xendit-go/qrcode"
	"github.com/xendit/xendit-go/utils/validator"
)

func initTesting(apiRequesterMockObj xendit.APIRequester) {
	xendit.Opt.SecretKey = "examplesecretkey"
	xendit.SetAPIRequester(apiRequesterMockObj)
}

type apiRequesterQRCodeMock struct {
	mock.Mock
}

func (m *apiRequesterQRCodeMock) Call(ctx context.Context, method string, path string, secretKey string, header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	mockTime, _ := time.Parse(time.RFC3339, "2020-02-02T00:00:00.000Z")

	result.(*xendit.QRCode).ID = "qr_8182837te-87st-49ing-8696-1239bd4d759c"
	result.(*xendit.QRCode).ExternalID = "testing_id_123"
	result.(*xendit.QRCode).CallbackURL = "https://webhook-site"
	result.(*xendit.QRCode).Amount = 50000
	result.(*xendit.QRCode).Type = xendit.StaticQRCode
	result.(*xendit.QRCode).QRString = "0002010102##########CO.XENDIT.WWW011893600#######14220002152#####414220010303TTT####015CO.XENDIT.WWW02180000000000000000000TTT52  045######ID5911XenditQRIS6007Jakarta6105121606##########3k1mOnF73h11111111#3k1mOnF73h6v53033605401163040BDB"
	result.(*xendit.QRCode).Status = "ACTIVE"
	result.(*xendit.QRCode).Created = &mockTime
	result.(*xendit.QRCode).Updated = &mockTime

	return nil
}

func TestCreateQRCode(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterQRCodeMock)
	initTesting(apiRequesterMockObj)

	mockTime, _ := time.Parse(time.RFC3339, "2020-02-02T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *qrcode.CreateQRCodeParams
		expectedRes *xendit.QRCode
		expectedErr *xendit.Error
	}{
		{
			desc: "should create a QR Code",
			data: &qrcode.CreateQRCodeParams{
				ExternalID:  "testing_id_123",
				CallbackURL: "https://webhook-site",
				Type:        xendit.StaticQRCode,
				Amount:      50000,
			},
			expectedRes: &xendit.QRCode{
				ID:          "qr_8182837te-87st-49ing-8696-1239bd4d759c",
				CallbackURL: "https://webhook-site",
				Amount:      50000,
				Type:        xendit.StaticQRCode,
				ExternalID:  "testing_id_123",
				QRString:    "0002010102##########CO.XENDIT.WWW011893600#######14220002152#####414220010303TTT####015CO.XENDIT.WWW02180000000000000000000TTT52  045######ID5911XenditQRIS6007Jakarta6105121606##########3k1mOnF73h11111111#3k1mOnF73h6v53033605401163040BDB",
				Status:      "ACTIVE",
				Created:     &mockTime,
				Updated:     &mockTime,
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &qrcode.CreateQRCodeParams{
				Amount:      10000,
				CallbackURL: "https://webhook-site",
				Type:        xendit.DynamicQRCode,
			},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'ExternalID'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"POST",
				xendit.Opt.XenditURL+"/qr_codes",
				xendit.Opt.SecretKey,
				nil,
				tC.data,
				&xendit.QRCode{},
			).Return(nil)

			resp, err := qrcode.CreateQRCode(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

func TestGetQRCode(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterQRCodeMock)
	initTesting(apiRequesterMockObj)

	mockTime, _ := time.Parse(time.RFC3339, "2020-02-02T00:00:00.000Z")

	testCases := []struct {
		desc        string
		data        *qrcode.GetQRCodeParams
		expectedRes *xendit.QRCode
		expectedErr *xendit.Error
	}{
		{
			desc: "should get a QR Code",
			data: &qrcode.GetQRCodeParams{
				ExternalID: "testing_id_123",
			},
			expectedRes: &xendit.QRCode{
				ID:          "qr_8182837te-87st-49ing-8696-1239bd4d759c",
				CallbackURL: "https://webhook-site",
				Amount:      50000,
				Type:        xendit.StaticQRCode,
				ExternalID:  "testing_id_123",
				QRString:    "0002010102##########CO.XENDIT.WWW011893600#######14220002152#####414220010303TTT####015CO.XENDIT.WWW02180000000000000000000TTT52  045######ID5911XenditQRIS6007Jakarta6105121606##########3k1mOnF73h11111111#3k1mOnF73h6v53033605401163040BDB",
				Status:      "ACTIVE",
				Created:     &mockTime,
				Updated:     &mockTime,
			},
			expectedErr: nil,
		},
		{
			desc:        "should report missing required fields",
			data:        &qrcode.GetQRCodeParams{},
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
				fmt.Sprintf("%s/qr_codes/%s", xendit.Opt.XenditURL, tC.data.ExternalID),
				xendit.Opt.SecretKey,
				nil,
				nil,
				&xendit.QRCode{},
			).Return(nil)

			resp, err := qrcode.GetQRCode(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

type apiRequesterGetQRCodePaymentsMock struct {
	mock.Mock
}

func (m *apiRequesterGetQRCodePaymentsMock) Call(ctx context.Context, method string, path string, secretKey string, headers *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	resultString := `[
		{
			"id": "qrpy_8182837te-87st-49ing-8696-1239bd4d759c",
			"amount": 1500,
			"created": "2020-01-08T18:18:18.857Z",
			"qr_code": {
				"id": "qr_8182837te-87st-49ing-8696-1239bd4d759c",
				"external_id": "testing_id_123",
				"qr_string": "0002010102##########CO.XENDIT.WWW011893600#######14220002152#####414220010303TTT####015CO.XENDIT.WWW02180000000000000000000TTT52045######ID5911XenditQRIS6007Jakarta6105121606##########3k1mOnF73h11111111#3k1mOnF73h6v53033605401163040BDB",
				"type": "DYNAMIC"
			},
			"status": "COMPLETED"
		}
	]`

	_ = json.Unmarshal([]byte(resultString), &result)
	return nil
}

func TestGetQRPayments(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterGetQRCodePaymentsMock)
	initTesting(apiRequesterMockObj)

	mockTime, _ := time.Parse(time.RFC3339, "2020-01-08T18:18:18.857Z")

	testCases := []struct {
		desc        string
		data        *qrcode.GetQRCodePaymentsParams
		expectedRes []xendit.QRCodePayments
		expectedErr *xendit.Error
	}{
		{
			desc: "should get QR code payments",
			data: &qrcode.GetQRCodePaymentsParams{
				ExternalID: "testing_id_123",
			},
			expectedRes: []xendit.QRCodePayments{
				{
					ID:      "qrpy_8182837te-87st-49ing-8696-1239bd4d759c",
					Amount:  1500,
					Created: &mockTime,
					Status:  "COMPLETED",
					QRCode: xendit.QRDetail{
						ID:         "qr_8182837te-87st-49ing-8696-1239bd4d759c",
						ExternalID: "testing_id_123",
						QRString:   "0002010102##########CO.XENDIT.WWW011893600#######14220002152#####414220010303TTT####015CO.XENDIT.WWW02180000000000000000000TTT52045######ID5911XenditQRIS6007Jakarta6105121606##########3k1mOnF73h11111111#3k1mOnF73h6v53033605401163040BDB",
						Type:       xendit.DynamicQRCode,
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
				fmt.Sprintf("%s/qr_codes/payments?%s", xendit.Opt.XenditURL, tC.data.QueryString()),
				xendit.Opt.SecretKey,
				nil,
				nil,
				&[]xendit.QRCodePayments{},
			).Return(nil)

			resp, err := qrcode.GetQRCodePayments(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
