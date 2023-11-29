/*
Payment Requests

This API is used for Payment Requests

API version: 1.45.2
*/


package payment_request

import (
	"encoding/json"
	"fmt"
)

// DirectDebitChannelProperties struct for DirectDebitChannelProperties
type DirectDebitChannelProperties struct {
	DirectDebitChannelPropertiesBankAccount *DirectDebitChannelPropertiesBankAccount
	DirectDebitChannelPropertiesBankRedirect *DirectDebitChannelPropertiesBankRedirect
	DirectDebitChannelPropertiesDebitCard *DirectDebitChannelPropertiesDebitCard
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DirectDebitChannelProperties) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into DirectDebitChannelPropertiesBankAccount
	err = json.Unmarshal(data, &dst.DirectDebitChannelPropertiesBankAccount);
	if err == nil {
		jsonDirectDebitChannelPropertiesBankAccount, _ := json.Marshal(dst.DirectDebitChannelPropertiesBankAccount)
		if string(jsonDirectDebitChannelPropertiesBankAccount) == "{}" { // empty struct
			dst.DirectDebitChannelPropertiesBankAccount = nil
		} else {
			return nil // data stored in dst.DirectDebitChannelPropertiesBankAccount, return on the first match
		}
	} else {
		dst.DirectDebitChannelPropertiesBankAccount = nil
	}

	// try to unmarshal JSON data into DirectDebitChannelPropertiesBankRedirect
	err = json.Unmarshal(data, &dst.DirectDebitChannelPropertiesBankRedirect);
	if err == nil {
		jsonDirectDebitChannelPropertiesBankRedirect, _ := json.Marshal(dst.DirectDebitChannelPropertiesBankRedirect)
		if string(jsonDirectDebitChannelPropertiesBankRedirect) == "{}" { // empty struct
			dst.DirectDebitChannelPropertiesBankRedirect = nil
		} else {
			return nil // data stored in dst.DirectDebitChannelPropertiesBankRedirect, return on the first match
		}
	} else {
		dst.DirectDebitChannelPropertiesBankRedirect = nil
	}

	// try to unmarshal JSON data into DirectDebitChannelPropertiesDebitCard
	err = json.Unmarshal(data, &dst.DirectDebitChannelPropertiesDebitCard);
	if err == nil {
		jsonDirectDebitChannelPropertiesDebitCard, _ := json.Marshal(dst.DirectDebitChannelPropertiesDebitCard)
		if string(jsonDirectDebitChannelPropertiesDebitCard) == "{}" { // empty struct
			dst.DirectDebitChannelPropertiesDebitCard = nil
		} else {
			return nil // data stored in dst.DirectDebitChannelPropertiesDebitCard, return on the first match
		}
	} else {
		dst.DirectDebitChannelPropertiesDebitCard = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(DirectDebitChannelProperties)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DirectDebitChannelProperties) MarshalJSON() ([]byte, error) {
	if src.DirectDebitChannelPropertiesBankAccount != nil {
		return json.Marshal(&src.DirectDebitChannelPropertiesBankAccount)
	}

	if src.DirectDebitChannelPropertiesBankRedirect != nil {
		return json.Marshal(&src.DirectDebitChannelPropertiesBankRedirect)
	}

	if src.DirectDebitChannelPropertiesDebitCard != nil {
		return json.Marshal(&src.DirectDebitChannelPropertiesDebitCard)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDirectDebitChannelProperties struct {
	value *DirectDebitChannelProperties
	isSet bool
}

func (v NullableDirectDebitChannelProperties) Get() *DirectDebitChannelProperties {
	return v.value
}

func (v *NullableDirectDebitChannelProperties) Set(val *DirectDebitChannelProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectDebitChannelProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectDebitChannelProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectDebitChannelProperties(val *DirectDebitChannelProperties) *NullableDirectDebitChannelProperties {
	return &NullableDirectDebitChannelProperties{value: val, isSet: true}
}

func (v NullableDirectDebitChannelProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectDebitChannelProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


