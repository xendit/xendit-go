/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.86.1
*/

// Code generated by OpenAPI Generator; DO NOT EDIT.

package payment_method

import (
	"encoding/json"
	
	"fmt"
)

// PaymentMethodReusability the model 'PaymentMethodReusability'
type PaymentMethodReusability string

// List of PaymentMethodReusability
const (
	PAYMENTMETHODREUSABILITY_MULTIPLE_USE PaymentMethodReusability = "MULTIPLE_USE"
	PAYMENTMETHODREUSABILITY_ONE_TIME_USE PaymentMethodReusability = "ONE_TIME_USE"
)

// All allowed values of PaymentMethodReusability enum
var AllowedPaymentMethodReusabilityEnumValues = []PaymentMethodReusability{
	"MULTIPLE_USE",
	"ONE_TIME_USE",
}

func (v *PaymentMethodReusability) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentMethodReusability(value)
	for _, existing := range AllowedPaymentMethodReusabilityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentMethodReusability", value)
}

// NewPaymentMethodReusabilityFromValue returns a pointer to a valid PaymentMethodReusability
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentMethodReusabilityFromValue(v string) (*PaymentMethodReusability, error) {
	ev := PaymentMethodReusability(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentMethodReusability: valid values are %v", v, AllowedPaymentMethodReusabilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentMethodReusability) IsValid() bool {
	for _, existing := range AllowedPaymentMethodReusabilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v PaymentMethodReusability) String() string {
	return string(v)
}

// Ptr returns reference to PaymentMethodReusability value
func (v PaymentMethodReusability) Ptr() *PaymentMethodReusability {
	return &v
}

type NullablePaymentMethodReusability struct {
	value *PaymentMethodReusability
	isSet bool
}

func (v NullablePaymentMethodReusability) Get() *PaymentMethodReusability {
	return v.value
}

func (v *NullablePaymentMethodReusability) Set(val *PaymentMethodReusability) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodReusability) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodReusability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodReusability(val *PaymentMethodReusability) *NullablePaymentMethodReusability {
	return &NullablePaymentMethodReusability{value: val, isSet: true}
}

func (v NullablePaymentMethodReusability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodReusability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
