/*
Payment Requests

This API is used for Payment Requests

API version: 1.45.1
*/


package payment_request

import (
	"encoding/json"
	
	"fmt"
)

// PaymentRequestCountry the model 'PaymentRequestCountry'
type PaymentRequestCountry string

// List of PaymentRequestCountry
const (
	PAYMENTREQUESTCOUNTRY_ID PaymentRequestCountry = "ID"
	PAYMENTREQUESTCOUNTRY_PH PaymentRequestCountry = "PH"
	PAYMENTREQUESTCOUNTRY_VN PaymentRequestCountry = "VN"
	PAYMENTREQUESTCOUNTRY_TH PaymentRequestCountry = "TH"
	PAYMENTREQUESTCOUNTRY_MY PaymentRequestCountry = "MY"
    PAYMENTREQUESTCOUNTRY_XENDIT_ENUM_DEFAULT_FALLBACK PaymentRequestCountry = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of PaymentRequestCountry enum
var AllowedPaymentRequestCountryEnumValues = []PaymentRequestCountry{
	"ID",
	"PH",
	"VN",
	"TH",
	"MY",
    "UNKNOWN_ENUM_VALUE",
}

func (v *PaymentRequestCountry) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentRequestCountry(value)
	for _, existing := range AllowedPaymentRequestCountryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = PAYMENTREQUESTCOUNTRY_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewPaymentRequestCountryFromValue returns a pointer to a valid PaymentRequestCountry
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentRequestCountryFromValue(v string) (*PaymentRequestCountry, error) {
	ev := PaymentRequestCountry(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentRequestCountry: valid values are %v", v, AllowedPaymentRequestCountryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentRequestCountry) IsValid() bool {
	for _, existing := range AllowedPaymentRequestCountryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v PaymentRequestCountry) String() string {
	return string(v)
}

// Ptr returns reference to PaymentRequestCountry value
func (v PaymentRequestCountry) Ptr() *PaymentRequestCountry {
	return &v
}

type NullablePaymentRequestCountry struct {
	value *PaymentRequestCountry
	isSet bool
}

func (v NullablePaymentRequestCountry) Get() *PaymentRequestCountry {
	return v.value
}

func (v *NullablePaymentRequestCountry) Set(val *PaymentRequestCountry) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequestCountry) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequestCountry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequestCountry(val *PaymentRequestCountry) *NullablePaymentRequestCountry {
	return &NullablePaymentRequestCountry{value: val, isSet: true}
}

func (v NullablePaymentRequestCountry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequestCountry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
