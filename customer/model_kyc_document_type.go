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

// KYCDocumentType the model 'KYCDocumentType'
type KYCDocumentType string

// List of KYCDocumentType
const (
	KYCDOCUMENTTYPE_BIRTH_CERTIFICATE KYCDocumentType = "BIRTH_CERTIFICATE"
	KYCDOCUMENTTYPE_BANK_STATEMENT KYCDocumentType = "BANK_STATEMENT"
	KYCDOCUMENTTYPE_DRIVING_LICENSE KYCDocumentType = "DRIVING_LICENSE"
	KYCDOCUMENTTYPE_IDENTITY_CARD KYCDocumentType = "IDENTITY_CARD"
	KYCDOCUMENTTYPE_PASSPORT KYCDocumentType = "PASSPORT"
	KYCDOCUMENTTYPE_VISA KYCDocumentType = "VISA"
	KYCDOCUMENTTYPE_BUSINESS_REGISTRATION KYCDocumentType = "BUSINESS_REGISTRATION"
	KYCDOCUMENTTYPE_BUSINESS_LICENSE KYCDocumentType = "BUSINESS_LICENSE"
)

// All allowed values of KYCDocumentType enum
var AllowedKYCDocumentTypeEnumValues = []KYCDocumentType{
	"BIRTH_CERTIFICATE",
	"BANK_STATEMENT",
	"DRIVING_LICENSE",
	"IDENTITY_CARD",
	"PASSPORT",
	"VISA",
	"BUSINESS_REGISTRATION",
	"BUSINESS_LICENSE",
}

func (v *KYCDocumentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KYCDocumentType(value)
	for _, existing := range AllowedKYCDocumentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid KYCDocumentType", value)
}

// NewKYCDocumentTypeFromValue returns a pointer to a valid KYCDocumentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKYCDocumentTypeFromValue(v string) (*KYCDocumentType, error) {
	ev := KYCDocumentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KYCDocumentType: valid values are %v", v, AllowedKYCDocumentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KYCDocumentType) IsValid() bool {
	for _, existing := range AllowedKYCDocumentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v KYCDocumentType) String() string {
	return string(v)
}

// Ptr returns reference to KYCDocumentType value
func (v KYCDocumentType) Ptr() *KYCDocumentType {
	return &v
}

type NullableKYCDocumentType struct {
	value *KYCDocumentType
	isSet bool
}

func (v NullableKYCDocumentType) Get() *KYCDocumentType {
	return v.value
}

func (v *NullableKYCDocumentType) Set(val *KYCDocumentType) {
	v.value = val
	v.isSet = true
}

func (v NullableKYCDocumentType) IsSet() bool {
	return v.isSet
}

func (v *NullableKYCDocumentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKYCDocumentType(val *KYCDocumentType) *NullableKYCDocumentType {
	return &NullableKYCDocumentType{value: val, isSet: true}
}

func (v NullableKYCDocumentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKYCDocumentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
