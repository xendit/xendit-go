/*
XENDIT SDK openapi spec

XENDIT SDK openapi spec

API version: 1.0.0
*/


package customer

import (
	"encoding/json"
	
	"fmt"
)

// IdentityAccountType the model 'IdentityAccountType'
type IdentityAccountType string

// List of IdentityAccountType
const (
	IDENTITYACCOUNTTYPE_BANK_ACCOUNT IdentityAccountType = "BANK_ACCOUNT"
	IDENTITYACCOUNTTYPE_EWALLET IdentityAccountType = "EWALLET"
	IDENTITYACCOUNTTYPE_CREDIT_CARD IdentityAccountType = "CREDIT_CARD"
	IDENTITYACCOUNTTYPE_PAY_LATER IdentityAccountType = "PAY_LATER"
	IDENTITYACCOUNTTYPE_OTC IdentityAccountType = "OTC"
	IDENTITYACCOUNTTYPE_QR_CODE IdentityAccountType = "QR_CODE"
    IDENTITYACCOUNTTYPE_XENDIT_ENUM_DEFAULT_FALLBACK IdentityAccountType = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of IdentityAccountType enum
var AllowedIdentityAccountTypeEnumValues = []IdentityAccountType{
	"BANK_ACCOUNT",
	"EWALLET",
	"CREDIT_CARD",
	"PAY_LATER",
	"OTC",
	"QR_CODE",
    "UNKNOWN_ENUM_VALUE",
}

func (v *IdentityAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IdentityAccountType(value)
	for _, existing := range AllowedIdentityAccountTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = IDENTITYACCOUNTTYPE_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewIdentityAccountTypeFromValue returns a pointer to a valid IdentityAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIdentityAccountTypeFromValue(v string) (*IdentityAccountType, error) {
	ev := IdentityAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IdentityAccountType: valid values are %v", v, AllowedIdentityAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IdentityAccountType) IsValid() bool {
	for _, existing := range AllowedIdentityAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v IdentityAccountType) String() string {
	return string(v)
}

// Ptr returns reference to IdentityAccountType value
func (v IdentityAccountType) Ptr() *IdentityAccountType {
	return &v
}

type NullableIdentityAccountType struct {
	value *IdentityAccountType
	isSet bool
}

func (v NullableIdentityAccountType) Get() *IdentityAccountType {
	return v.value
}

func (v *NullableIdentityAccountType) Set(val *IdentityAccountType) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityAccountType) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityAccountType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityAccountType(val *IdentityAccountType) *NullableIdentityAccountType {
	return &NullableIdentityAccountType{value: val, isSet: true}
}

func (v NullableIdentityAccountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityAccountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
