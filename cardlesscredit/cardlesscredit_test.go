package cardlesscredit_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/cardlesscredit"
	"github.com/xendit/xendit-go/utils/validator"
	"net/http"
	"testing"
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

	result.(*xendit.CardlessCredit).RedirectURL = "https://google.com"
	result.(*xendit.CardlessCredit).TransactionID = "trx1"
	result.(*xendit.CardlessCredit).OrderID = "ord1"
	result.(*xendit.CardlessCredit).ExternalID = "cardless-credit-1"
	result.(*xendit.CardlessCredit).CardlessCreditType = xendit.CardlessCreditTypeEnumKREDIVO

	return nil
}

func TestCreatePayment(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	testCases := []struct {
		desc        string
		data        *cardlesscredit.CreatePaymentParams
		expectedRes *xendit.CardlessCredit
		expectedErr *xendit.Error
	}{
		{
			desc: "should create a payment",
			data: &cardlesscredit.CreatePaymentParams{
				CardlessCreditType: xendit.CardlessCreditTypeEnumKREDIVO,
				ExternalID:         "cardless-credit-1",
				Amount:             200000,
				PaymentType:        xendit.PaymentTypeEnum3Months,
				Items: []cardlesscredit.Item{
					{
						ID:       "123",
						Name:     "Laptop Asus Ila",
						Price:    200000,
						Type:     "Laptop",
						URL:      "http://asus-ila.com",
						Quantity: 1,
					},
				},
				CustomerDetails: cardlesscredit.CustomerDetails{
					FirstName: "Michael",
					LastName:  "Belajarrock",
					Email:     "michaelbelajarrock@mail.com",
					Phone:     "08123123123",
				},
				ShippingAddress: cardlesscredit.ShippingAddress{
					FirstName:   "Michael",
					LastName:    "Belajarjazz",
					Address:     "Jalan Teknologi No. 12",
					City:        "Jakarta",
					PostalCode:  "40000",
					Phone:       "08123123123",
					CountryCode: "IDN",
				},
				RedirectURL: "https://google.com",
				CallbackURL: "https://google.com",
			},
			expectedRes: &xendit.CardlessCredit{
				RedirectURL:        "https://google.com",
				TransactionID:      "trx1",
				OrderID:            "ord1",
				ExternalID:         "cardless-credit-1",
				CardlessCreditType: xendit.CardlessCreditTypeEnumKREDIVO,
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &cardlesscredit.CreatePaymentParams{
				CardlessCreditType: xendit.CardlessCreditTypeEnumKREDIVO,
				ExternalID:         "cardless-credit-1",
				PaymentType:        xendit.PaymentTypeEnum3Months,
				Items: []cardlesscredit.Item{
					{
						ID:       "123",
						Name:     "Laptop Asus Ila",
						Price:    200000,
						Type:     "Laptop",
						URL:      "http://asus-ila.com",
						Quantity: 1,
					},
				},
				CustomerDetails: cardlesscredit.CustomerDetails{
					FirstName: "Michael",
					LastName:  "Belajarrock",
					Email:     "michaelbelajarrock@mail.com",
					Phone:     "08123123123",
				},
				ShippingAddress: cardlesscredit.ShippingAddress{
					FirstName:   "Michael",
					LastName:    "Belajarjazz",
					Address:     "Jalan Teknologi No. 12",
					City:        "Jakarta",
					PostalCode:  "40000",
					Phone:       "08123123123",
					CountryCode: "IDN",
				},
				RedirectURL: "https://google.com",
				CallbackURL: "https://google.com",
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
				xendit.Opt.XenditURL+"/cardless-credit",
				xendit.Opt.SecretKey,
				&http.Header{},
				tC.data,
				&xendit.CardlessCredit{},
			).Return(nil)

			resp, err := cardlesscredit.CreatePayment(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
