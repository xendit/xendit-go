/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.99.0
*/


package payment_method

import (
	"encoding/json"
	
	"fmt"
)

// CardChannelCode Card Channel Code
type CardChannelCode string

// List of CardChannelCode
const (
	CARDCHANNELCODE_GPN CardChannelCode = "GPN"
    CARDCHANNELCODE_XENDIT_ENUM_DEFAULT_FALLBACK CardChannelCode = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of CardChannelCode enum
var AllowedCardChannelCodeEnumValues = []CardChannelCode{
	"GPN",
    "UNKNOWN_ENUM_VALUE",
}

func (v *CardChannelCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CardChannelCode(value)
	for _, existing := range AllowedCardChannelCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = CARDCHANNELCODE_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewCardChannelCodeFromValue returns a pointer to a valid CardChannelCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCardChannelCodeFromValue(v string) (*CardChannelCode, error) {
	ev := CardChannelCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CardChannelCode: valid values are %v", v, AllowedCardChannelCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CardChannelCode) IsValid() bool {
	for _, existing := range AllowedCardChannelCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v CardChannelCode) String() string {
	return string(v)
}

// Ptr returns reference to CardChannelCode value
func (v CardChannelCode) Ptr() *CardChannelCode {
	return &v
}

type NullableCardChannelCode struct {
	value *CardChannelCode
	isSet bool
}

func (v NullableCardChannelCode) Get() *CardChannelCode {
	return v.value
}

func (v *NullableCardChannelCode) Set(val *CardChannelCode) {
	v.value = val
	v.isSet = true
}

func (v NullableCardChannelCode) IsSet() bool {
	return v.isSet
}

func (v *NullableCardChannelCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardChannelCode(val *CardChannelCode) *NullableCardChannelCode {
	return &NullableCardChannelCode{value: val, isSet: true}
}

func (v NullableCardChannelCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardChannelCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
