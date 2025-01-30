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

// PaylaterType Representing the available paylater channels used for invoice-related transactions.
type PaylaterType string

// List of PaylaterType
const (
	PAYLATERTYPE_KREDIVO PaylaterType = "KREDIVO"
	PAYLATERTYPE_AKULAKU PaylaterType = "AKULAKU"
	PAYLATERTYPE_UANGME PaylaterType = "UANGME"
	PAYLATERTYPE_BILLEASE PaylaterType = "BILLEASE"
	PAYLATERTYPE_CASHALO PaylaterType = "CASHALO"
	PAYLATERTYPE_ATOME PaylaterType = "ATOME"
    PAYLATERTYPE_XENDIT_ENUM_DEFAULT_FALLBACK PaylaterType = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of PaylaterType enum
var AllowedPaylaterTypeEnumValues = []PaylaterType{
	"KREDIVO",
	"AKULAKU",
	"UANGME",
	"BILLEASE",
	"CASHALO",
	"ATOME",
    "UNKNOWN_ENUM_VALUE",
}

func (v *PaylaterType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaylaterType(value)
	for _, existing := range AllowedPaylaterTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = PAYLATERTYPE_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewPaylaterTypeFromValue returns a pointer to a valid PaylaterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaylaterTypeFromValue(v string) (*PaylaterType, error) {
	ev := PaylaterType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaylaterType: valid values are %v", v, AllowedPaylaterTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaylaterType) IsValid() bool {
	for _, existing := range AllowedPaylaterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v PaylaterType) String() string {
	return string(v)
}

// Ptr returns reference to PaylaterType value
func (v PaylaterType) Ptr() *PaylaterType {
	return &v
}

type NullablePaylaterType struct {
	value *PaylaterType
	isSet bool
}

func (v NullablePaylaterType) Get() *PaylaterType {
	return v.value
}

func (v *NullablePaylaterType) Set(val *PaylaterType) {
	v.value = val
	v.isSet = true
}

func (v NullablePaylaterType) IsSet() bool {
	return v.isSet
}

func (v *NullablePaylaterType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaylaterType(val *PaylaterType) *NullablePaylaterType {
	return &NullablePaylaterType{value: val, isSet: true}
}

func (v NullablePaylaterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaylaterType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
