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

// OverTheCounterChannelCode Over The Counter Channel Code
type OverTheCounterChannelCode string

// List of OverTheCounterChannelCode
const (
	OVERTHECOUNTERCHANNELCODE__7_ELEVEN OverTheCounterChannelCode = "7ELEVEN"
	OVERTHECOUNTERCHANNELCODE__7_ELEVEN_CLIQQ OverTheCounterChannelCode = "7ELEVEN_CLIQQ"
	OVERTHECOUNTERCHANNELCODE_CEBUANA OverTheCounterChannelCode = "CEBUANA"
	OVERTHECOUNTERCHANNELCODE_ECPAY OverTheCounterChannelCode = "ECPAY"
	OVERTHECOUNTERCHANNELCODE_PALAWAN OverTheCounterChannelCode = "PALAWAN"
	OVERTHECOUNTERCHANNELCODE_MLHUILLIER OverTheCounterChannelCode = "MLHUILLIER"
	OVERTHECOUNTERCHANNELCODE_ECPAY_DRAGONLOAN OverTheCounterChannelCode = "ECPAY_DRAGONLOAN"
	OVERTHECOUNTERCHANNELCODE_LBC OverTheCounterChannelCode = "LBC"
	OVERTHECOUNTERCHANNELCODE_ECPAY_SCHOOL OverTheCounterChannelCode = "ECPAY_SCHOOL"
	OVERTHECOUNTERCHANNELCODE_RD_PAWNSHOP OverTheCounterChannelCode = "RD_PAWNSHOP"
	OVERTHECOUNTERCHANNELCODE_CVM OverTheCounterChannelCode = "CVM"
	OVERTHECOUNTERCHANNELCODE_USSC OverTheCounterChannelCode = "USSC"
	OVERTHECOUNTERCHANNELCODE_SM_BILLS OverTheCounterChannelCode = "SM_BILLS"
	OVERTHECOUNTERCHANNELCODE_ROBINSONS_BILLS OverTheCounterChannelCode = "ROBINSONS_BILLS"
	OVERTHECOUNTERCHANNELCODE_ALFAMART OverTheCounterChannelCode = "ALFAMART"
	OVERTHECOUNTERCHANNELCODE_INDOMARET OverTheCounterChannelCode = "INDOMARET"
    OVERTHECOUNTERCHANNELCODE_XENDIT_ENUM_DEFAULT_FALLBACK OverTheCounterChannelCode = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of OverTheCounterChannelCode enum
var AllowedOverTheCounterChannelCodeEnumValues = []OverTheCounterChannelCode{
	"7ELEVEN",
	"7ELEVEN_CLIQQ",
	"CEBUANA",
	"ECPAY",
	"PALAWAN",
	"MLHUILLIER",
	"ECPAY_DRAGONLOAN",
	"LBC",
	"ECPAY_SCHOOL",
	"RD_PAWNSHOP",
	"CVM",
	"USSC",
	"SM_BILLS",
	"ROBINSONS_BILLS",
	"ALFAMART",
	"INDOMARET",
    "UNKNOWN_ENUM_VALUE",
}

func (v *OverTheCounterChannelCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OverTheCounterChannelCode(value)
	for _, existing := range AllowedOverTheCounterChannelCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = OVERTHECOUNTERCHANNELCODE_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewOverTheCounterChannelCodeFromValue returns a pointer to a valid OverTheCounterChannelCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOverTheCounterChannelCodeFromValue(v string) (*OverTheCounterChannelCode, error) {
	ev := OverTheCounterChannelCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OverTheCounterChannelCode: valid values are %v", v, AllowedOverTheCounterChannelCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OverTheCounterChannelCode) IsValid() bool {
	for _, existing := range AllowedOverTheCounterChannelCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v OverTheCounterChannelCode) String() string {
	return string(v)
}

// Ptr returns reference to OverTheCounterChannelCode value
func (v OverTheCounterChannelCode) Ptr() *OverTheCounterChannelCode {
	return &v
}

type NullableOverTheCounterChannelCode struct {
	value *OverTheCounterChannelCode
	isSet bool
}

func (v NullableOverTheCounterChannelCode) Get() *OverTheCounterChannelCode {
	return v.value
}

func (v *NullableOverTheCounterChannelCode) Set(val *OverTheCounterChannelCode) {
	v.value = val
	v.isSet = true
}

func (v NullableOverTheCounterChannelCode) IsSet() bool {
	return v.isSet
}

func (v *NullableOverTheCounterChannelCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverTheCounterChannelCode(val *OverTheCounterChannelCode) *NullableOverTheCounterChannelCode {
	return &NullableOverTheCounterChannelCode{value: val, isSet: true}
}

func (v NullableOverTheCounterChannelCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverTheCounterChannelCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
