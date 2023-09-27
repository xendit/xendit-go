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

// DirectDebitType Representing the available Direct Debit channels used for invoice-related transactions.
type DirectDebitType string

// List of DirectDebitType
const (
	DIRECTDEBITTYPE_BA_BRI DirectDebitType = "BA_BRI"
	DIRECTDEBITTYPE_DC_BRI DirectDebitType = "DC_BRI"
	DIRECTDEBITTYPE_DD_BRI DirectDebitType = "DD_BRI"
	DIRECTDEBITTYPE_DD_MANDIRI DirectDebitType = "DD_MANDIRI"
	DIRECTDEBITTYPE_BA_BPI DirectDebitType = "BA_BPI"
	DIRECTDEBITTYPE_DC_BPI DirectDebitType = "DC_BPI"
	DIRECTDEBITTYPE_DD_BPI DirectDebitType = "DD_BPI"
	DIRECTDEBITTYPE_BA_UBP DirectDebitType = "BA_UBP"
	DIRECTDEBITTYPE_DC_UBP DirectDebitType = "DC_UBP"
	DIRECTDEBITTYPE_DD_UBP DirectDebitType = "DD_UBP"
	DIRECTDEBITTYPE_BCA_KLIKPAY DirectDebitType = "BCA_KLIKPAY"
	DIRECTDEBITTYPE_BA_BCA_KLIKPAY DirectDebitType = "BA_BCA_KLIKPAY"
	DIRECTDEBITTYPE_DC_BCA_KLIKPAY DirectDebitType = "DC_BCA_KLIKPAY"
	DIRECTDEBITTYPE_DD_BCA_KLIKPAY DirectDebitType = "DD_BCA_KLIKPAY"
	DIRECTDEBITTYPE_DD_BDO_EPAY DirectDebitType = "DD_BDO_EPAY"
	DIRECTDEBITTYPE_DD_RCBC DirectDebitType = "DD_RCBC"
	DIRECTDEBITTYPE_DD_CHINABANK DirectDebitType = "DD_CHINABANK"
	DIRECTDEBITTYPE_BA_CHINABANK DirectDebitType = "BA_CHINABANK"
	DIRECTDEBITTYPE_DC_CHINABANK DirectDebitType = "DC_CHINABANK"
)

// All allowed values of DirectDebitType enum
var AllowedDirectDebitTypeEnumValues = []DirectDebitType{
	"BA_BRI",
	"DC_BRI",
	"DD_BRI",
	"DD_MANDIRI",
	"BA_BPI",
	"DC_BPI",
	"DD_BPI",
	"BA_UBP",
	"DC_UBP",
	"DD_UBP",
	"BCA_KLIKPAY",
	"BA_BCA_KLIKPAY",
	"DC_BCA_KLIKPAY",
	"DD_BCA_KLIKPAY",
	"DD_BDO_EPAY",
	"DD_RCBC",
	"DD_CHINABANK",
	"BA_CHINABANK",
	"DC_CHINABANK",
}

func (v *DirectDebitType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DirectDebitType(value)
	for _, existing := range AllowedDirectDebitTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DirectDebitType", value)
}

// NewDirectDebitTypeFromValue returns a pointer to a valid DirectDebitType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDirectDebitTypeFromValue(v string) (*DirectDebitType, error) {
	ev := DirectDebitType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DirectDebitType: valid values are %v", v, AllowedDirectDebitTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DirectDebitType) IsValid() bool {
	for _, existing := range AllowedDirectDebitTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v DirectDebitType) String() string {
	return string(v)
}

// Ptr returns reference to DirectDebitType value
func (v DirectDebitType) Ptr() *DirectDebitType {
	return &v
}

type NullableDirectDebitType struct {
	value *DirectDebitType
	isSet bool
}

func (v NullableDirectDebitType) Get() *DirectDebitType {
	return v.value
}

func (v *NullableDirectDebitType) Set(val *DirectDebitType) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectDebitType) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectDebitType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectDebitType(val *DirectDebitType) *NullableDirectDebitType {
	return &NullableDirectDebitType{value: val, isSet: true}
}

func (v NullableDirectDebitType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectDebitType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
