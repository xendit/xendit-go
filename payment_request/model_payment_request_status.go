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

// PaymentRequestStatus the model 'PaymentRequestStatus'
type PaymentRequestStatus string

// List of PaymentRequestStatus
const (
	PAYMENTREQUESTSTATUS_PENDING PaymentRequestStatus = "PENDING"
	PAYMENTREQUESTSTATUS_REQUIRES_ACTION PaymentRequestStatus = "REQUIRES_ACTION"
	PAYMENTREQUESTSTATUS_CANCELED PaymentRequestStatus = "CANCELED"
	PAYMENTREQUESTSTATUS_SUCCEEDED PaymentRequestStatus = "SUCCEEDED"
	PAYMENTREQUESTSTATUS_FAILED PaymentRequestStatus = "FAILED"
	PAYMENTREQUESTSTATUS_VOIDED PaymentRequestStatus = "VOIDED"
	PAYMENTREQUESTSTATUS_UNKNOWN PaymentRequestStatus = "UNKNOWN"
	PAYMENTREQUESTSTATUS_AWAITING_CAPTURE PaymentRequestStatus = "AWAITING_CAPTURE"
    PAYMENTREQUESTSTATUS_XENDIT_ENUM_DEFAULT_FALLBACK PaymentRequestStatus = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of PaymentRequestStatus enum
var AllowedPaymentRequestStatusEnumValues = []PaymentRequestStatus{
	"PENDING",
	"REQUIRES_ACTION",
	"CANCELED",
	"SUCCEEDED",
	"FAILED",
	"VOIDED",
	"UNKNOWN",
	"AWAITING_CAPTURE",
    "UNKNOWN_ENUM_VALUE",
}

func (v *PaymentRequestStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentRequestStatus(value)
	for _, existing := range AllowedPaymentRequestStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = PAYMENTREQUESTSTATUS_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewPaymentRequestStatusFromValue returns a pointer to a valid PaymentRequestStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentRequestStatusFromValue(v string) (*PaymentRequestStatus, error) {
	ev := PaymentRequestStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentRequestStatus: valid values are %v", v, AllowedPaymentRequestStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentRequestStatus) IsValid() bool {
	for _, existing := range AllowedPaymentRequestStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v PaymentRequestStatus) String() string {
	return string(v)
}

// Ptr returns reference to PaymentRequestStatus value
func (v PaymentRequestStatus) Ptr() *PaymentRequestStatus {
	return &v
}

type NullablePaymentRequestStatus struct {
	value *PaymentRequestStatus
	isSet bool
}

func (v NullablePaymentRequestStatus) Get() *PaymentRequestStatus {
	return v.value
}

func (v *NullablePaymentRequestStatus) Set(val *PaymentRequestStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequestStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequestStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequestStatus(val *PaymentRequestStatus) *NullablePaymentRequestStatus {
	return &NullablePaymentRequestStatus{value: val, isSet: true}
}

func (v NullablePaymentRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequestStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
