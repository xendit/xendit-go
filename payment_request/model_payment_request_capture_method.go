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

// PaymentRequestCaptureMethod the model 'PaymentRequestCaptureMethod'
type PaymentRequestCaptureMethod string

// List of PaymentRequestCaptureMethod
const (
	PAYMENTREQUESTCAPTUREMETHOD_AUTOMATIC PaymentRequestCaptureMethod = "AUTOMATIC"
	PAYMENTREQUESTCAPTUREMETHOD_MANUAL PaymentRequestCaptureMethod = "MANUAL"
)

// All allowed values of PaymentRequestCaptureMethod enum
var AllowedPaymentRequestCaptureMethodEnumValues = []PaymentRequestCaptureMethod{
	"AUTOMATIC",
	"MANUAL",
}

func (v *PaymentRequestCaptureMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentRequestCaptureMethod(value)
	for _, existing := range AllowedPaymentRequestCaptureMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentRequestCaptureMethod", value)
}

// NewPaymentRequestCaptureMethodFromValue returns a pointer to a valid PaymentRequestCaptureMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentRequestCaptureMethodFromValue(v string) (*PaymentRequestCaptureMethod, error) {
	ev := PaymentRequestCaptureMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentRequestCaptureMethod: valid values are %v", v, AllowedPaymentRequestCaptureMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentRequestCaptureMethod) IsValid() bool {
	for _, existing := range AllowedPaymentRequestCaptureMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v PaymentRequestCaptureMethod) String() string {
	return string(v)
}

// Ptr returns reference to PaymentRequestCaptureMethod value
func (v PaymentRequestCaptureMethod) Ptr() *PaymentRequestCaptureMethod {
	return &v
}

type NullablePaymentRequestCaptureMethod struct {
	value *PaymentRequestCaptureMethod
	isSet bool
}

func (v NullablePaymentRequestCaptureMethod) Get() *PaymentRequestCaptureMethod {
	return v.value
}

func (v *NullablePaymentRequestCaptureMethod) Set(val *PaymentRequestCaptureMethod) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequestCaptureMethod) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequestCaptureMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequestCaptureMethod(val *PaymentRequestCaptureMethod) *NullablePaymentRequestCaptureMethod {
	return &NullablePaymentRequestCaptureMethod{value: val, isSet: true}
}

func (v NullablePaymentRequestCaptureMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequestCaptureMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
