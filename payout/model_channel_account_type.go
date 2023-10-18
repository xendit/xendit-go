/*
Payout Service

This API allows Xendit to send money from an account to a channel (banks, eWallets, retail outlets) from across regions

API version: 1.0.0
*/


package payout

import (
	"encoding/json"
	
	"fmt"
)

// ChannelAccountType Available account types (applicable for MY_DUITNOW)
type ChannelAccountType string

// List of ChannelAccountType
const (
	CHANNELACCOUNTTYPE_NATIONAL_ID ChannelAccountType = "NATIONAL_ID"
	CHANNELACCOUNTTYPE_MOBILE_NO ChannelAccountType = "MOBILE_NO"
	CHANNELACCOUNTTYPE_PASSPORT ChannelAccountType = "PASSPORT"
	CHANNELACCOUNTTYPE_BUSINESS_REGISTRATION ChannelAccountType = "BUSINESS_REGISTRATION"
	CHANNELACCOUNTTYPE_BANK_ACCOUNT ChannelAccountType = "BANK_ACCOUNT"
    CHANNELACCOUNTTYPE_XENDIT_ENUM_DEFAULT_FALLBACK ChannelAccountType = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of ChannelAccountType enum
var AllowedChannelAccountTypeEnumValues = []ChannelAccountType{
	"NATIONAL_ID",
	"MOBILE_NO",
	"PASSPORT",
	"BUSINESS_REGISTRATION",
	"BANK_ACCOUNT",
    "UNKNOWN_ENUM_VALUE",
}

func (v *ChannelAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChannelAccountType(value)
	for _, existing := range AllowedChannelAccountTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = CHANNELACCOUNTTYPE_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewChannelAccountTypeFromValue returns a pointer to a valid ChannelAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChannelAccountTypeFromValue(v string) (*ChannelAccountType, error) {
	ev := ChannelAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChannelAccountType: valid values are %v", v, AllowedChannelAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChannelAccountType) IsValid() bool {
	for _, existing := range AllowedChannelAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v ChannelAccountType) String() string {
	return string(v)
}

// Ptr returns reference to ChannelAccountType value
func (v ChannelAccountType) Ptr() *ChannelAccountType {
	return &v
}

type NullableChannelAccountType struct {
	value *ChannelAccountType
	isSet bool
}

func (v NullableChannelAccountType) Get() *ChannelAccountType {
	return v.value
}

func (v *NullableChannelAccountType) Set(val *ChannelAccountType) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelAccountType) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelAccountType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelAccountType(val *ChannelAccountType) *NullableChannelAccountType {
	return &NullableChannelAccountType{value: val, isSet: true}
}

func (v NullableChannelAccountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelAccountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
