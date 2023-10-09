/*
Payment Requests

This API is used for Payment Requests

API version: 1.44.1
*/


package payment_request

import (
	"encoding/json"
	
	"fmt"
)

// PaymentRequestCurrency the model 'PaymentRequestCurrency'
type PaymentRequestCurrency string

// List of PaymentRequestCurrency
const (
	PAYMENTREQUESTCURRENCY_IDR PaymentRequestCurrency = "IDR"
	PAYMENTREQUESTCURRENCY_PHP PaymentRequestCurrency = "PHP"
	PAYMENTREQUESTCURRENCY_VND PaymentRequestCurrency = "VND"
	PAYMENTREQUESTCURRENCY_THB PaymentRequestCurrency = "THB"
	PAYMENTREQUESTCURRENCY_MYR PaymentRequestCurrency = "MYR"
)

// All allowed values of PaymentRequestCurrency enum
var AllowedPaymentRequestCurrencyEnumValues = []PaymentRequestCurrency{
	"IDR",
	"PHP",
	"VND",
	"THB",
	"MYR",
}

func (v *PaymentRequestCurrency) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentRequestCurrency(value)
	for _, existing := range AllowedPaymentRequestCurrencyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentRequestCurrency", value)
}

// NewPaymentRequestCurrencyFromValue returns a pointer to a valid PaymentRequestCurrency
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentRequestCurrencyFromValue(v string) (*PaymentRequestCurrency, error) {
	ev := PaymentRequestCurrency(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentRequestCurrency: valid values are %v", v, AllowedPaymentRequestCurrencyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentRequestCurrency) IsValid() bool {
	for _, existing := range AllowedPaymentRequestCurrencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v PaymentRequestCurrency) String() string {
	return string(v)
}

// Ptr returns reference to PaymentRequestCurrency value
func (v PaymentRequestCurrency) Ptr() *PaymentRequestCurrency {
	return &v
}

type NullablePaymentRequestCurrency struct {
	value *PaymentRequestCurrency
	isSet bool
}

func (v NullablePaymentRequestCurrency) Get() *PaymentRequestCurrency {
	return v.value
}

func (v *NullablePaymentRequestCurrency) Set(val *PaymentRequestCurrency) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequestCurrency) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequestCurrency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequestCurrency(val *PaymentRequestCurrency) *NullablePaymentRequestCurrency {
	return &NullablePaymentRequestCurrency{value: val, isSet: true}
}

func (v NullablePaymentRequestCurrency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequestCurrency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
