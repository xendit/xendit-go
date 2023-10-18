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

// PaymentMethodType the model 'PaymentMethodType'
type PaymentMethodType string

// List of PaymentMethodType
const (
	PAYMENTMETHODTYPE_CARD PaymentMethodType = "CARD"
	PAYMENTMETHODTYPE_DIRECT_DEBIT PaymentMethodType = "DIRECT_DEBIT"
	PAYMENTMETHODTYPE_EWALLET PaymentMethodType = "EWALLET"
	PAYMENTMETHODTYPE_OVER_THE_COUNTER PaymentMethodType = "OVER_THE_COUNTER"
	PAYMENTMETHODTYPE_QR_CODE PaymentMethodType = "QR_CODE"
	PAYMENTMETHODTYPE_VIRTUAL_ACCOUNT PaymentMethodType = "VIRTUAL_ACCOUNT"
    PAYMENTMETHODTYPE_XENDIT_ENUM_DEFAULT_FALLBACK PaymentMethodType = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of PaymentMethodType enum
var AllowedPaymentMethodTypeEnumValues = []PaymentMethodType{
	"CARD",
	"DIRECT_DEBIT",
	"EWALLET",
	"OVER_THE_COUNTER",
	"QR_CODE",
	"VIRTUAL_ACCOUNT",
    "UNKNOWN_ENUM_VALUE",
}

func (v *PaymentMethodType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentMethodType(value)
	for _, existing := range AllowedPaymentMethodTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = PAYMENTMETHODTYPE_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewPaymentMethodTypeFromValue returns a pointer to a valid PaymentMethodType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentMethodTypeFromValue(v string) (*PaymentMethodType, error) {
	ev := PaymentMethodType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentMethodType: valid values are %v", v, AllowedPaymentMethodTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentMethodType) IsValid() bool {
	for _, existing := range AllowedPaymentMethodTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v PaymentMethodType) String() string {
	return string(v)
}

// Ptr returns reference to PaymentMethodType value
func (v PaymentMethodType) Ptr() *PaymentMethodType {
	return &v
}

type NullablePaymentMethodType struct {
	value *PaymentMethodType
	isSet bool
}

func (v NullablePaymentMethodType) Get() *PaymentMethodType {
	return v.value
}

func (v *NullablePaymentMethodType) Set(val *PaymentMethodType) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodType) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodType(val *PaymentMethodType) *NullablePaymentMethodType {
	return &NullablePaymentMethodType{value: val, isSet: true}
}

func (v NullablePaymentMethodType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
