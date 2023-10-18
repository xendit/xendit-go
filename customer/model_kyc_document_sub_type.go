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

// KYCDocumentSubType the model 'KYCDocumentSubType'
type KYCDocumentSubType string

// List of KYCDocumentSubType
const (
	KYCDOCUMENTSUBTYPE_NATIONAL_ID KYCDocumentSubType = "NATIONAL_ID"
	KYCDOCUMENTSUBTYPE_CONSULAR_ID KYCDocumentSubType = "CONSULAR_ID"
	KYCDOCUMENTSUBTYPE_VOTER_ID KYCDocumentSubType = "VOTER_ID"
	KYCDOCUMENTSUBTYPE_POSTAL_ID KYCDocumentSubType = "POSTAL_ID"
	KYCDOCUMENTSUBTYPE_RESIDENCE_PERMIT KYCDocumentSubType = "RESIDENCE_PERMIT"
	KYCDOCUMENTSUBTYPE_TAX_ID KYCDocumentSubType = "TAX_ID"
	KYCDOCUMENTSUBTYPE_STUDENT_ID KYCDocumentSubType = "STUDENT_ID"
	KYCDOCUMENTSUBTYPE_MILITARY_ID KYCDocumentSubType = "MILITARY_ID"
	KYCDOCUMENTSUBTYPE_MEDICAL_ID KYCDocumentSubType = "MEDICAL_ID"
	KYCDOCUMENTSUBTYPE_OTHERS KYCDocumentSubType = "OTHERS"
    KYCDOCUMENTSUBTYPE_XENDIT_ENUM_DEFAULT_FALLBACK KYCDocumentSubType = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of KYCDocumentSubType enum
var AllowedKYCDocumentSubTypeEnumValues = []KYCDocumentSubType{
	"NATIONAL_ID",
	"CONSULAR_ID",
	"VOTER_ID",
	"POSTAL_ID",
	"RESIDENCE_PERMIT",
	"TAX_ID",
	"STUDENT_ID",
	"MILITARY_ID",
	"MEDICAL_ID",
	"OTHERS",
    "UNKNOWN_ENUM_VALUE",
}

func (v *KYCDocumentSubType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KYCDocumentSubType(value)
	for _, existing := range AllowedKYCDocumentSubTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = KYCDOCUMENTSUBTYPE_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewKYCDocumentSubTypeFromValue returns a pointer to a valid KYCDocumentSubType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKYCDocumentSubTypeFromValue(v string) (*KYCDocumentSubType, error) {
	ev := KYCDocumentSubType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KYCDocumentSubType: valid values are %v", v, AllowedKYCDocumentSubTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KYCDocumentSubType) IsValid() bool {
	for _, existing := range AllowedKYCDocumentSubTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v KYCDocumentSubType) String() string {
	return string(v)
}

// Ptr returns reference to KYCDocumentSubType value
func (v KYCDocumentSubType) Ptr() *KYCDocumentSubType {
	return &v
}

type NullableKYCDocumentSubType struct {
	value *KYCDocumentSubType
	isSet bool
}

func (v NullableKYCDocumentSubType) Get() *KYCDocumentSubType {
	return v.value
}

func (v *NullableKYCDocumentSubType) Set(val *KYCDocumentSubType) {
	v.value = val
	v.isSet = true
}

func (v NullableKYCDocumentSubType) IsSet() bool {
	return v.isSet
}

func (v *NullableKYCDocumentSubType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKYCDocumentSubType(val *KYCDocumentSubType) *NullableKYCDocumentSubType {
	return &NullableKYCDocumentSubType{value: val, isSet: true}
}

func (v NullableKYCDocumentSubType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKYCDocumentSubType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
