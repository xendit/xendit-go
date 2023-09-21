/*
Payment Requests

This API is used for Payment Requests

API version: 1.44.0
*/


package payment_request

import (
	"encoding/json"
	
	"fmt"
)

// DirectDebitType the model 'DirectDebitType'
type DirectDebitType string

// List of DirectDebitType
const (
	DIRECTDEBITTYPE_DEBIT_CARD DirectDebitType = "DEBIT_CARD"
	DIRECTDEBITTYPE_BANK_ACCOUNT DirectDebitType = "BANK_ACCOUNT"
	DIRECTDEBITTYPE_BANK_REDIRECT DirectDebitType = "BANK_REDIRECT"
)

// All allowed values of DirectDebitType enum
var AllowedDirectDebitTypeEnumValues = []DirectDebitType{
	"DEBIT_CARD",
	"BANK_ACCOUNT",
	"BANK_REDIRECT",
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
