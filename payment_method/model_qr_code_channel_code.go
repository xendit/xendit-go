/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.128.0
*/


package payment_method

import (
	"encoding/json"
	
	"fmt"
)

// QRCodeChannelCode QR Code Channel Code
type QRCodeChannelCode string

// List of QRCodeChannelCode
const (
	QRCODECHANNELCODE_QRIS QRCodeChannelCode = "QRIS"
	QRCODECHANNELCODE_DANA QRCodeChannelCode = "DANA"
	QRCODECHANNELCODE_RCBC QRCodeChannelCode = "RCBC"
	QRCODECHANNELCODE_PROMPTPAY QRCodeChannelCode = "PROMPTPAY"
	QRCODECHANNELCODE_LINKAJA QRCodeChannelCode = "LINKAJA"
	QRCODECHANNELCODE_XENDIT QRCodeChannelCode = "XENDIT"
	QRCODECHANNELCODE_QRPH QRCodeChannelCode = "QRPH"
    QRCODECHANNELCODE_XENDIT_ENUM_DEFAULT_FALLBACK QRCodeChannelCode = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of QRCodeChannelCode enum
var AllowedQRCodeChannelCodeEnumValues = []QRCodeChannelCode{
	"QRIS",
	"DANA",
	"RCBC",
	"PROMPTPAY",
	"LINKAJA",
	"XENDIT",
	"QRPH",
    "UNKNOWN_ENUM_VALUE",
}

func (v *QRCodeChannelCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QRCodeChannelCode(value)
	for _, existing := range AllowedQRCodeChannelCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = QRCODECHANNELCODE_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewQRCodeChannelCodeFromValue returns a pointer to a valid QRCodeChannelCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQRCodeChannelCodeFromValue(v string) (*QRCodeChannelCode, error) {
	ev := QRCodeChannelCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QRCodeChannelCode: valid values are %v", v, AllowedQRCodeChannelCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QRCodeChannelCode) IsValid() bool {
	for _, existing := range AllowedQRCodeChannelCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v QRCodeChannelCode) String() string {
	return string(v)
}

// Ptr returns reference to QRCodeChannelCode value
func (v QRCodeChannelCode) Ptr() *QRCodeChannelCode {
	return &v
}

type NullableQRCodeChannelCode struct {
	value *QRCodeChannelCode
	isSet bool
}

func (v NullableQRCodeChannelCode) Get() *QRCodeChannelCode {
	return v.value
}

func (v *NullableQRCodeChannelCode) Set(val *QRCodeChannelCode) {
	v.value = val
	v.isSet = true
}

func (v NullableQRCodeChannelCode) IsSet() bool {
	return v.isSet
}

func (v *NullableQRCodeChannelCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQRCodeChannelCode(val *QRCodeChannelCode) *NullableQRCodeChannelCode {
	return &NullableQRCodeChannelCode{value: val, isSet: true}
}

func (v NullableQRCodeChannelCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQRCodeChannelCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
