/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.89.2
*/


package payment_method

import (
	"encoding/json"
	
	"fmt"
)

// VirtualAccountChannelCode Virtual Account Channel Code
type VirtualAccountChannelCode string

// List of VirtualAccountChannelCode
const (
	VIRTUALACCOUNTCHANNELCODE_BCA VirtualAccountChannelCode = "BCA"
	VIRTUALACCOUNTCHANNELCODE_BJB VirtualAccountChannelCode = "BJB"
	VIRTUALACCOUNTCHANNELCODE_BNI VirtualAccountChannelCode = "BNI"
	VIRTUALACCOUNTCHANNELCODE_BRI VirtualAccountChannelCode = "BRI"
	VIRTUALACCOUNTCHANNELCODE_MANDIRI VirtualAccountChannelCode = "MANDIRI"
	VIRTUALACCOUNTCHANNELCODE_PERMATA VirtualAccountChannelCode = "PERMATA"
	VIRTUALACCOUNTCHANNELCODE_BSI VirtualAccountChannelCode = "BSI"
	VIRTUALACCOUNTCHANNELCODE_CIMB VirtualAccountChannelCode = "CIMB"
	VIRTUALACCOUNTCHANNELCODE_SAHABAT_SAMPOERNA VirtualAccountChannelCode = "SAHABAT_SAMPOERNA"
	VIRTUALACCOUNTCHANNELCODE_ARTAJASA VirtualAccountChannelCode = "ARTAJASA"
	VIRTUALACCOUNTCHANNELCODE_PV VirtualAccountChannelCode = "PV"
	VIRTUALACCOUNTCHANNELCODE_VIETCAPITAL VirtualAccountChannelCode = "VIETCAPITAL"
	VIRTUALACCOUNTCHANNELCODE_WOORI VirtualAccountChannelCode = "WOORI"
	VIRTUALACCOUNTCHANNELCODE_MSB VirtualAccountChannelCode = "MSB"
	VIRTUALACCOUNTCHANNELCODE_STANDARD_CHARTERED VirtualAccountChannelCode = "STANDARD_CHARTERED"
    VIRTUALACCOUNTCHANNELCODE_XENDIT_ENUM_DEFAULT_FALLBACK VirtualAccountChannelCode = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of VirtualAccountChannelCode enum
var AllowedVirtualAccountChannelCodeEnumValues = []VirtualAccountChannelCode{
	"BCA",
	"BJB",
	"BNI",
	"BRI",
	"MANDIRI",
	"PERMATA",
	"BSI",
	"CIMB",
	"SAHABAT_SAMPOERNA",
	"ARTAJASA",
	"PV",
	"VIETCAPITAL",
	"WOORI",
	"MSB",
	"STANDARD_CHARTERED",
    "UNKNOWN_ENUM_VALUE",
}

func (v *VirtualAccountChannelCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VirtualAccountChannelCode(value)
	for _, existing := range AllowedVirtualAccountChannelCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = VIRTUALACCOUNTCHANNELCODE_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewVirtualAccountChannelCodeFromValue returns a pointer to a valid VirtualAccountChannelCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVirtualAccountChannelCodeFromValue(v string) (*VirtualAccountChannelCode, error) {
	ev := VirtualAccountChannelCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VirtualAccountChannelCode: valid values are %v", v, AllowedVirtualAccountChannelCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VirtualAccountChannelCode) IsValid() bool {
	for _, existing := range AllowedVirtualAccountChannelCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v VirtualAccountChannelCode) String() string {
	return string(v)
}

// Ptr returns reference to VirtualAccountChannelCode value
func (v VirtualAccountChannelCode) Ptr() *VirtualAccountChannelCode {
	return &v
}

type NullableVirtualAccountChannelCode struct {
	value *VirtualAccountChannelCode
	isSet bool
}

func (v NullableVirtualAccountChannelCode) Get() *VirtualAccountChannelCode {
	return v.value
}

func (v *NullableVirtualAccountChannelCode) Set(val *VirtualAccountChannelCode) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualAccountChannelCode) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualAccountChannelCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualAccountChannelCode(val *VirtualAccountChannelCode) *NullableVirtualAccountChannelCode {
	return &NullableVirtualAccountChannelCode{value: val, isSet: true}
}

func (v NullableVirtualAccountChannelCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualAccountChannelCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
