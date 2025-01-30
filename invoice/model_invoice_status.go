/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.8.7
*/


package invoice

import (
	"encoding/json"
	
	"fmt"
)

// InvoiceStatus Representing the status of an invoice.
type InvoiceStatus string

// List of InvoiceStatus
const (
	INVOICESTATUS_PENDING InvoiceStatus = "PENDING"
	INVOICESTATUS_PAID InvoiceStatus = "PAID"
	INVOICESTATUS_SETTLED InvoiceStatus = "SETTLED"
	INVOICESTATUS_EXPIRED InvoiceStatus = "EXPIRED"
    INVOICESTATUS_XENDIT_ENUM_DEFAULT_FALLBACK InvoiceStatus = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of InvoiceStatus enum
var AllowedInvoiceStatusEnumValues = []InvoiceStatus{
	"PENDING",
	"PAID",
	"SETTLED",
	"EXPIRED",
    "UNKNOWN_ENUM_VALUE",
}

func (v *InvoiceStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InvoiceStatus(value)
	for _, existing := range AllowedInvoiceStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = INVOICESTATUS_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewInvoiceStatusFromValue returns a pointer to a valid InvoiceStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInvoiceStatusFromValue(v string) (*InvoiceStatus, error) {
	ev := InvoiceStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InvoiceStatus: valid values are %v", v, AllowedInvoiceStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InvoiceStatus) IsValid() bool {
	for _, existing := range AllowedInvoiceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v InvoiceStatus) String() string {
	return string(v)
}

// Ptr returns reference to InvoiceStatus value
func (v InvoiceStatus) Ptr() *InvoiceStatus {
	return &v
}

type NullableInvoiceStatus struct {
	value *InvoiceStatus
	isSet bool
}

func (v NullableInvoiceStatus) Get() *InvoiceStatus {
	return v.value
}

func (v *NullableInvoiceStatus) Set(val *InvoiceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceStatus(val *InvoiceStatus) *NullableInvoiceStatus {
	return &NullableInvoiceStatus{value: val, isSet: true}
}

func (v NullableInvoiceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
