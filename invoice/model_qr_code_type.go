/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.6.0
*/


package invoice

import (
	"encoding/json"
	
	"fmt"
)

// QrCodeType Representing the available QR Code channels used for invoice-related transactions.
type QrCodeType string

// List of QrCodeType
const (
	QRCODETYPE_QRIS QrCodeType = "QRIS"
	QRCODETYPE_PROMPTPAY QrCodeType = "PROMPTPAY"
    QRCODETYPE_XENDIT_ENUM_DEFAULT_FALLBACK QrCodeType = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of QrCodeType enum
var AllowedQrCodeTypeEnumValues = []QrCodeType{
	"QRIS",
	"PROMPTPAY",
    "UNKNOWN_ENUM_VALUE",
}

func (v *QrCodeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QrCodeType(value)
	for _, existing := range AllowedQrCodeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = QRCODETYPE_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewQrCodeTypeFromValue returns a pointer to a valid QrCodeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQrCodeTypeFromValue(v string) (*QrCodeType, error) {
	ev := QrCodeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QrCodeType: valid values are %v", v, AllowedQrCodeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QrCodeType) IsValid() bool {
	for _, existing := range AllowedQrCodeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v QrCodeType) String() string {
	return string(v)
}

// Ptr returns reference to QrCodeType value
func (v QrCodeType) Ptr() *QrCodeType {
	return &v
}

type NullableQrCodeType struct {
	value *QrCodeType
	isSet bool
}

func (v NullableQrCodeType) Get() *QrCodeType {
	return v.value
}

func (v *NullableQrCodeType) Set(val *QrCodeType) {
	v.value = val
	v.isSet = true
}

func (v NullableQrCodeType) IsSet() bool {
	return v.isSet
}

func (v *NullableQrCodeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQrCodeType(val *QrCodeType) *NullableQrCodeType {
	return &NullableQrCodeType{value: val, isSet: true}
}

func (v NullableQrCodeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQrCodeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
