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

// ChannelCategory Category of channel code, as some channels might require more fields during processing
type ChannelCategory string

// List of ChannelCategory
const (
	CHANNELCATEGORY_BANK ChannelCategory = "BANK"
	CHANNELCATEGORY_EWALLET ChannelCategory = "EWALLET"
	CHANNELCATEGORY_OTC ChannelCategory = "OTC"
    CHANNELCATEGORY_XENDIT_ENUM_DEFAULT_FALLBACK ChannelCategory = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of ChannelCategory enum
var AllowedChannelCategoryEnumValues = []ChannelCategory{
	"BANK",
	"EWALLET",
	"OTC",
    "UNKNOWN_ENUM_VALUE",
}

func (v *ChannelCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChannelCategory(value)
	for _, existing := range AllowedChannelCategoryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = CHANNELCATEGORY_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewChannelCategoryFromValue returns a pointer to a valid ChannelCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChannelCategoryFromValue(v string) (*ChannelCategory, error) {
	ev := ChannelCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChannelCategory: valid values are %v", v, AllowedChannelCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChannelCategory) IsValid() bool {
	for _, existing := range AllowedChannelCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v ChannelCategory) String() string {
	return string(v)
}

// Ptr returns reference to ChannelCategory value
func (v ChannelCategory) Ptr() *ChannelCategory {
	return &v
}

type NullableChannelCategory struct {
	value *ChannelCategory
	isSet bool
}

func (v NullableChannelCategory) Get() *ChannelCategory {
	return v.value
}

func (v *NullableChannelCategory) Set(val *ChannelCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelCategory(val *ChannelCategory) *NullableChannelCategory {
	return &NullableChannelCategory{value: val, isSet: true}
}

func (v NullableChannelCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
