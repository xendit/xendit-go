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

// GetPayouts200ResponseDataInner struct for GetPayouts200ResponseDataInner
type GetPayouts200ResponseDataInner struct {
	Payout *Payout
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *GetPayouts200ResponseDataInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into Payout
	err = json.Unmarshal(data, &dst.Payout);
	if err == nil {
		jsonPayout, _ := json.Marshal(dst.Payout)
		if string(jsonPayout) == "{}" { // empty struct
			dst.Payout = nil
		} else {
			return nil // data stored in dst.Payout, return on the first match
		}
	} else {
		dst.Payout = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(GetPayouts200ResponseDataInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *GetPayouts200ResponseDataInner) MarshalJSON() ([]byte, error) {
	if src.Payout != nil {
		return json.Marshal(&src.Payout)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableGetPayouts200ResponseDataInner struct {
	value *GetPayouts200ResponseDataInner
	isSet bool
}

func (v NullableGetPayouts200ResponseDataInner) Get() *GetPayouts200ResponseDataInner {
	return v.value
}

func (v *NullableGetPayouts200ResponseDataInner) Set(val *GetPayouts200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPayouts200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPayouts200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPayouts200ResponseDataInner(val *GetPayouts200ResponseDataInner) *NullableGetPayouts200ResponseDataInner {
	return &NullableGetPayouts200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetPayouts200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPayouts200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


