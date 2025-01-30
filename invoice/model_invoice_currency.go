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

// InvoiceCurrency Representing the currency used for an invoice.
type InvoiceCurrency string

// List of InvoiceCurrency
const (
	INVOICECURRENCY_IDR InvoiceCurrency = "IDR"
	INVOICECURRENCY_USD InvoiceCurrency = "USD"
	INVOICECURRENCY_THB InvoiceCurrency = "THB"
	INVOICECURRENCY_VND InvoiceCurrency = "VND"
	INVOICECURRENCY_PHP InvoiceCurrency = "PHP"
	INVOICECURRENCY_MYR InvoiceCurrency = "MYR"
    INVOICECURRENCY_XENDIT_ENUM_DEFAULT_FALLBACK InvoiceCurrency = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of InvoiceCurrency enum
var AllowedInvoiceCurrencyEnumValues = []InvoiceCurrency{
	"IDR",
	"USD",
	"THB",
	"VND",
	"PHP",
	"MYR",
    "UNKNOWN_ENUM_VALUE",
}

func (v *InvoiceCurrency) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InvoiceCurrency(value)
	for _, existing := range AllowedInvoiceCurrencyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = INVOICECURRENCY_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewInvoiceCurrencyFromValue returns a pointer to a valid InvoiceCurrency
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInvoiceCurrencyFromValue(v string) (*InvoiceCurrency, error) {
	ev := InvoiceCurrency(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InvoiceCurrency: valid values are %v", v, AllowedInvoiceCurrencyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InvoiceCurrency) IsValid() bool {
	for _, existing := range AllowedInvoiceCurrencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v InvoiceCurrency) String() string {
	return string(v)
}

// Ptr returns reference to InvoiceCurrency value
func (v InvoiceCurrency) Ptr() *InvoiceCurrency {
	return &v
}

type NullableInvoiceCurrency struct {
	value *InvoiceCurrency
	isSet bool
}

func (v NullableInvoiceCurrency) Get() *InvoiceCurrency {
	return v.value
}

func (v *NullableInvoiceCurrency) Set(val *InvoiceCurrency) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceCurrency) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceCurrency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceCurrency(val *InvoiceCurrency) *NullableInvoiceCurrency {
	return &NullableInvoiceCurrency{value: val, isSet: true}
}

func (v NullableInvoiceCurrency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceCurrency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
