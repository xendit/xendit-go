package customer

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"
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

	result.(*xendit.Customer).ID = "791ac956-397a-400f-9fda-4958894e61b5"
	result.(*xendit.Customer).ReferenceID = "test-reference-id"
	result.(*xendit.Customer).Email = "tes@tes.com"
	result.(*xendit.Customer).MobileNumber = "+6281234567890"
	result.(*xendit.Customer).GivenNames = "Given Names"
	result.(*xendit.Customer).Nationality = "ID"
	result.(*xendit.Customer).DateOfBirth = "1995-12-30"

	customerAddress := xendit.CustomerAddress{
		Country:		"ID",
		StreetLine1:	"Jl. 123",
		StreetLine2:    "Jl. 456",
		City:			"Jakarta Selatan",
		Province:       "DKI Jakarta",
		State:			"-",
		PostalCode:     "12345",
	}

	result.(*xendit.Customer).Addresses = []xendit.CustomerAddress{customerAddress}
	result.(*xendit.Customer).Metadata = map[string]interface{}{
		"meta": "data",
	}

	return nil
}

func TestCreateCustomer(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMock)
	initTesting(apiRequesterMockObj)

	customerAddress := xendit.CustomerAddress{
		Country:		"ID",
		StreetLine1:	"Jl. 123",
		StreetLine2:    "Jl. 456",
		City:			"Jakarta Selatan",
		Province:       "DKI Jakarta",
		State:			"-",
		PostalCode:     "12345",
	}

	metadata := map[string]interface{}{
		"meta": "data",
	}

	testCases := []struct {
		desc        string
		data        *CreateCustomerParams
		expectedRes *xendit.Customer
		expectedErr *xendit.Error
	}{
		{
			desc: "should create a customer",
			data: &CreateCustomerParams{
				ReferenceID: 	"test-reference-id",
				Email:			"tes@tes.com",
				GivenNames:     "Given Names",
				Nationality: 	"ID",
				DateOfBirth: 	"1995-12-30",
				Addresses:		[]xendit.CustomerAddress{customerAddress},
				Metadata:		metadata,
			},
			expectedRes: &xendit.Customer{
				ID:				"791ac956-397a-400f-9fda-4958894e61b5",
				ReferenceID:  	"test-reference-id",
				GivenNames:     "Given Names",
				Email: 			"tes@tes.com",
				MobileNumber: 	"+6281234567890",
				Nationality: 	"ID",
				DateOfBirth: 	"1995-12-30",
				Addresses: 		[]xendit.CustomerAddress{customerAddress},
				Metadata: 		metadata,
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &CreateCustomerParams{
				ReferenceID: "test-reference-id",
			},
			expectedRes: nil,
			expectedErr: validator.APIValidatorErr(errors.New("Missing required fields: 'GivenNames'")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			apiRequesterMockObj.On(
				"Call",
				context.Background(),
				"POST",
				xendit.Opt.XenditURL+"/customers",
				xendit.Opt.SecretKey,
				&http.Header{},
				tC.data,
				&xendit.Customer{},
			).Return(nil)

			resp, err := CreateCustomer(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}

type apiRequesterMockGet struct {
	mock.Mock
}

func (m *apiRequesterMockGet) Call(ctx context.Context, method string, path string, secretKey string, header *http.Header, params interface{}, result interface{}) *xendit.Error {
	m.Called(ctx, method, path, secretKey, nil, params, result)

	resultString := `[{
        "ID": "791ac956-397a-400f-9fda-4958894e61b5",
        "reference_id": "test-reference-id",
        "given_names": "Given Names",
        "email": "tes@tes.com",
        "mobile_number": "+6281234567890",
        "description": null,
        "middle_name": null,
		"surname": null,
		"phone_number": null,
		"nationality": "ID",
		"date_of_birth": "1995-12-30",
		"metadata": {
			"meta": "data"
		},
		"employment": null,
		"addresses": [
			{
				"category": "",
				"country": "ID",
				"state": "-",
				"province": "DKI Jakarta",
				"city": "Jakarta Selatan",
				"postal_code": "12345",
				"street_line1": "Jl. 123",
				"street_line2": "Jl. 456",
				"is_preferred": false
			}
		],
		"source_of_wealth": null
	}]`

	_ = json.Unmarshal([]byte(resultString), &result)

	return nil
}

func TestGetCustomerByReferenceID(t *testing.T) {
	apiRequesterMockObj := new(apiRequesterMockGet)
	initTesting(apiRequesterMockObj)

	customerAddress := xendit.CustomerAddress{
		Country:		"ID",
		StreetLine1:	"Jl. 123",
		StreetLine2:    "Jl. 456",
		City:			"Jakarta Selatan",
		Province:       "DKI Jakarta",
		State:			"-",
		PostalCode:     "12345",
	}

	metadata := map[string]interface{}{
		"meta": "data",
	}

	testCases := []struct {
		desc        string
		data        *GetCustomerByReferenceIDParams
		expectedRes []xendit.Customer
		expectedErr *xendit.Error
	}{
		{
			desc: "should get customer by reference ID",
			data: &GetCustomerByReferenceIDParams{
				ReferenceID: "test-reference-id",
			},
			expectedRes: []xendit.Customer{
				{
					ID:				"791ac956-397a-400f-9fda-4958894e61b5",
					ReferenceID:  	"test-reference-id",
					GivenNames:     "Given Names",
					Email: 			"tes@tes.com",
					MobileNumber: 	"+6281234567890",
					Nationality: 	"ID",
					DateOfBirth: 	"1995-12-30",
					Addresses: 		[]xendit.CustomerAddress{customerAddress},
					Metadata: 		metadata,
				},
			},
			expectedErr: nil,
		},
		{
			desc: "should report missing required fields",
			data: &GetCustomerByReferenceIDParams{},
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
				xendit.Opt.XenditURL+"/customers?"+tC.data.QueryString(),
				xendit.Opt.SecretKey,
				nil,
				nil,
				&[]xendit.Customer{},
			).Return(nil)

			resp, err := GetCustomerByReferenceID(tC.data)

			assert.Equal(t, tC.expectedRes, resp)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
