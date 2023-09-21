/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.5.0
*/


package invoice

import (
	"encoding/json"
	
	"fmt"
)

// InvoicePaymentMethod Representing the payment method used for an invoice.
type InvoicePaymentMethod string

// List of InvoicePaymentMethod
const (
	INVOICEPAYMENTMETHOD_POOL InvoicePaymentMethod = "POOL"
	INVOICEPAYMENTMETHOD_CALLBACK_VIRTUAL_ACCOUNT InvoicePaymentMethod = "CALLBACK_VIRTUAL_ACCOUNT"
	INVOICEPAYMENTMETHOD_CREDIT_CARD InvoicePaymentMethod = "CREDIT_CARD"
	INVOICEPAYMENTMETHOD_RETAIL_OUTLET InvoicePaymentMethod = "RETAIL_OUTLET"
	INVOICEPAYMENTMETHOD_QR_CODE InvoicePaymentMethod = "QR_CODE"
	INVOICEPAYMENTMETHOD_QRIS InvoicePaymentMethod = "QRIS"
	INVOICEPAYMENTMETHOD_EWALLET InvoicePaymentMethod = "EWALLET"
	INVOICEPAYMENTMETHOD_DIRECT_DEBIT InvoicePaymentMethod = "DIRECT_DEBIT"
	INVOICEPAYMENTMETHOD_BANK_TRANSFER InvoicePaymentMethod = "BANK_TRANSFER"
	INVOICEPAYMENTMETHOD_PAYLATER InvoicePaymentMethod = "PAYLATER"
)

// All allowed values of InvoicePaymentMethod enum
var AllowedInvoicePaymentMethodEnumValues = []InvoicePaymentMethod{
	"POOL",
	"CALLBACK_VIRTUAL_ACCOUNT",
	"CREDIT_CARD",
	"RETAIL_OUTLET",
	"QR_CODE",
	"QRIS",
	"EWALLET",
	"DIRECT_DEBIT",
	"BANK_TRANSFER",
	"PAYLATER",
}

func (v *InvoicePaymentMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InvoicePaymentMethod(value)
	for _, existing := range AllowedInvoicePaymentMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InvoicePaymentMethod", value)
}

// NewInvoicePaymentMethodFromValue returns a pointer to a valid InvoicePaymentMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInvoicePaymentMethodFromValue(v string) (*InvoicePaymentMethod, error) {
	ev := InvoicePaymentMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InvoicePaymentMethod: valid values are %v", v, AllowedInvoicePaymentMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InvoicePaymentMethod) IsValid() bool {
	for _, existing := range AllowedInvoicePaymentMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v InvoicePaymentMethod) String() string {
	return string(v)
}

// Ptr returns reference to InvoicePaymentMethod value
func (v InvoicePaymentMethod) Ptr() *InvoicePaymentMethod {
	return &v
}

type NullableInvoicePaymentMethod struct {
	value *InvoicePaymentMethod
	isSet bool
}

func (v NullableInvoicePaymentMethod) Get() *InvoicePaymentMethod {
	return v.value
}

func (v *NullableInvoicePaymentMethod) Set(val *InvoicePaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoicePaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoicePaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoicePaymentMethod(val *InvoicePaymentMethod) *NullableInvoicePaymentMethod {
	return &NullableInvoicePaymentMethod{value: val, isSet: true}
}

func (v NullableInvoicePaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoicePaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
