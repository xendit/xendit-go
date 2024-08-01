/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.128.0
*/


package payment_method

import (
	"encoding/json"
	
	"fmt"
)

// PaymentMethodCountry the model 'PaymentMethodCountry'
type PaymentMethodCountry string

// List of PaymentMethodCountry
const (
	PAYMENTMETHODCOUNTRY_PH PaymentMethodCountry = "PH"
	PAYMENTMETHODCOUNTRY_ID PaymentMethodCountry = "ID"
	PAYMENTMETHODCOUNTRY_VN PaymentMethodCountry = "VN"
	PAYMENTMETHODCOUNTRY_TH PaymentMethodCountry = "TH"
	PAYMENTMETHODCOUNTRY_MY PaymentMethodCountry = "MY"
	PAYMENTMETHODCOUNTRY_US PaymentMethodCountry = "US"
    PAYMENTMETHODCOUNTRY_XENDIT_ENUM_DEFAULT_FALLBACK PaymentMethodCountry = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of PaymentMethodCountry enum
var AllowedPaymentMethodCountryEnumValues = []PaymentMethodCountry{
	"PH",
	"ID",
	"VN",
	"TH",
	"MY",
	"US",
    "UNKNOWN_ENUM_VALUE",
}

func (v *PaymentMethodCountry) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentMethodCountry(value)
	for _, existing := range AllowedPaymentMethodCountryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = PAYMENTMETHODCOUNTRY_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewPaymentMethodCountryFromValue returns a pointer to a valid PaymentMethodCountry
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentMethodCountryFromValue(v string) (*PaymentMethodCountry, error) {
	ev := PaymentMethodCountry(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentMethodCountry: valid values are %v", v, AllowedPaymentMethodCountryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentMethodCountry) IsValid() bool {
	for _, existing := range AllowedPaymentMethodCountryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v PaymentMethodCountry) String() string {
	return string(v)
}

// Ptr returns reference to PaymentMethodCountry value
func (v PaymentMethodCountry) Ptr() *PaymentMethodCountry {
	return &v
}

type NullablePaymentMethodCountry struct {
	value *PaymentMethodCountry
	isSet bool
}

func (v NullablePaymentMethodCountry) Get() *PaymentMethodCountry {
	return v.value
}

func (v *NullablePaymentMethodCountry) Set(val *PaymentMethodCountry) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodCountry) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodCountry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodCountry(val *PaymentMethodCountry) *NullablePaymentMethodCountry {
	return &NullablePaymentMethodCountry{value: val, isSet: true}
}

func (v NullablePaymentMethodCountry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodCountry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
