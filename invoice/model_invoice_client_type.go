/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.8.8
*/


package invoice

import (
	"encoding/json"
	
	"fmt"
)

// InvoiceClientType Representing the client type or source of an invoice.
type InvoiceClientType string

// List of InvoiceClientType
const (
	INVOICECLIENTTYPE_DASHBOARD InvoiceClientType = "DASHBOARD"
	INVOICECLIENTTYPE_API_GATEWAY InvoiceClientType = "API_GATEWAY"
	INVOICECLIENTTYPE_INTEGRATION InvoiceClientType = "INTEGRATION"
	INVOICECLIENTTYPE_ON_DEMAND InvoiceClientType = "ON_DEMAND"
	INVOICECLIENTTYPE_RECURRING InvoiceClientType = "RECURRING"
	INVOICECLIENTTYPE_MOBILE InvoiceClientType = "MOBILE"
    INVOICECLIENTTYPE_XENDIT_ENUM_DEFAULT_FALLBACK InvoiceClientType = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of InvoiceClientType enum
var AllowedInvoiceClientTypeEnumValues = []InvoiceClientType{
	"DASHBOARD",
	"API_GATEWAY",
	"INTEGRATION",
	"ON_DEMAND",
	"RECURRING",
	"MOBILE",
    "UNKNOWN_ENUM_VALUE",
}

func (v *InvoiceClientType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InvoiceClientType(value)
	for _, existing := range AllowedInvoiceClientTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = INVOICECLIENTTYPE_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewInvoiceClientTypeFromValue returns a pointer to a valid InvoiceClientType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInvoiceClientTypeFromValue(v string) (*InvoiceClientType, error) {
	ev := InvoiceClientType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InvoiceClientType: valid values are %v", v, AllowedInvoiceClientTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InvoiceClientType) IsValid() bool {
	for _, existing := range AllowedInvoiceClientTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v InvoiceClientType) String() string {
	return string(v)
}

// Ptr returns reference to InvoiceClientType value
func (v InvoiceClientType) Ptr() *InvoiceClientType {
	return &v
}

type NullableInvoiceClientType struct {
	value *InvoiceClientType
	isSet bool
}

func (v NullableInvoiceClientType) Get() *InvoiceClientType {
	return v.value
}

func (v *NullableInvoiceClientType) Set(val *InvoiceClientType) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceClientType) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceClientType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceClientType(val *InvoiceClientType) *NullableInvoiceClientType {
	return &NullableInvoiceClientType{value: val, isSet: true}
}

func (v NullableInvoiceClientType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceClientType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
