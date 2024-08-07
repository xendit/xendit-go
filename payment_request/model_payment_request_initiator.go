/*
Payment Requests

This API is used for Payment Requests

API version: 1.70.0
*/


package payment_request

import (
	"encoding/json"
	
	"fmt"
)

// PaymentRequestInitiator the model 'PaymentRequestInitiator'
type PaymentRequestInitiator string

// List of PaymentRequestInitiator
const (
	PAYMENTREQUESTINITIATOR_CUSTOMER PaymentRequestInitiator = "CUSTOMER"
	PAYMENTREQUESTINITIATOR_MERCHANT PaymentRequestInitiator = "MERCHANT"
    PAYMENTREQUESTINITIATOR_XENDIT_ENUM_DEFAULT_FALLBACK PaymentRequestInitiator = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of PaymentRequestInitiator enum
var AllowedPaymentRequestInitiatorEnumValues = []PaymentRequestInitiator{
	"CUSTOMER",
	"MERCHANT",
    "UNKNOWN_ENUM_VALUE",
}

func (v *PaymentRequestInitiator) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentRequestInitiator(value)
	for _, existing := range AllowedPaymentRequestInitiatorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = PAYMENTREQUESTINITIATOR_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewPaymentRequestInitiatorFromValue returns a pointer to a valid PaymentRequestInitiator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentRequestInitiatorFromValue(v string) (*PaymentRequestInitiator, error) {
	ev := PaymentRequestInitiator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentRequestInitiator: valid values are %v", v, AllowedPaymentRequestInitiatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentRequestInitiator) IsValid() bool {
	for _, existing := range AllowedPaymentRequestInitiatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v PaymentRequestInitiator) String() string {
	return string(v)
}

// Ptr returns reference to PaymentRequestInitiator value
func (v PaymentRequestInitiator) Ptr() *PaymentRequestInitiator {
	return &v
}

type NullablePaymentRequestInitiator struct {
	value *PaymentRequestInitiator
	isSet bool
}

func (v NullablePaymentRequestInitiator) Get() *PaymentRequestInitiator {
	return v.value
}

func (v *NullablePaymentRequestInitiator) Set(val *PaymentRequestInitiator) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequestInitiator) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequestInitiator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequestInitiator(val *PaymentRequestInitiator) *NullablePaymentRequestInitiator {
	return &NullablePaymentRequestInitiator{value: val, isSet: true}
}

func (v NullablePaymentRequestInitiator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequestInitiator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
