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

// PaymentMethodStatus the model 'PaymentMethodStatus'
type PaymentMethodStatus string

// List of PaymentMethodStatus
const (
	PAYMENTMETHODSTATUS_ACTIVE PaymentMethodStatus = "ACTIVE"
	PAYMENTMETHODSTATUS_EXPIRED PaymentMethodStatus = "EXPIRED"
	PAYMENTMETHODSTATUS_INACTIVE PaymentMethodStatus = "INACTIVE"
	PAYMENTMETHODSTATUS_PENDING PaymentMethodStatus = "PENDING"
	PAYMENTMETHODSTATUS_REQUIRES_ACTION PaymentMethodStatus = "REQUIRES_ACTION"
	PAYMENTMETHODSTATUS_FAILED PaymentMethodStatus = "FAILED"
    PAYMENTMETHODSTATUS_XENDIT_ENUM_DEFAULT_FALLBACK PaymentMethodStatus = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of PaymentMethodStatus enum
var AllowedPaymentMethodStatusEnumValues = []PaymentMethodStatus{
	"ACTIVE",
	"EXPIRED",
	"INACTIVE",
	"PENDING",
	"REQUIRES_ACTION",
	"FAILED",
    "UNKNOWN_ENUM_VALUE",
}

func (v *PaymentMethodStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentMethodStatus(value)
	for _, existing := range AllowedPaymentMethodStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = PAYMENTMETHODSTATUS_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewPaymentMethodStatusFromValue returns a pointer to a valid PaymentMethodStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentMethodStatusFromValue(v string) (*PaymentMethodStatus, error) {
	ev := PaymentMethodStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentMethodStatus: valid values are %v", v, AllowedPaymentMethodStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentMethodStatus) IsValid() bool {
	for _, existing := range AllowedPaymentMethodStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v PaymentMethodStatus) String() string {
	return string(v)
}

// Ptr returns reference to PaymentMethodStatus value
func (v PaymentMethodStatus) Ptr() *PaymentMethodStatus {
	return &v
}

type NullablePaymentMethodStatus struct {
	value *PaymentMethodStatus
	isSet bool
}

func (v NullablePaymentMethodStatus) Get() *PaymentMethodStatus {
	return v.value
}

func (v *NullablePaymentMethodStatus) Set(val *PaymentMethodStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodStatus(val *PaymentMethodStatus) *NullablePaymentMethodStatus {
	return &NullablePaymentMethodStatus{value: val, isSet: true}
}

func (v NullablePaymentMethodStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
