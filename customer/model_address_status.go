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

// AddressStatus the model 'AddressStatus'
type AddressStatus string

// List of AddressStatus
const (
	ADDRESSSTATUS_ACTIVE AddressStatus = "ACTIVE"
	ADDRESSSTATUS_DELETED AddressStatus = "DELETED"
)

// All allowed values of AddressStatus enum
var AllowedAddressStatusEnumValues = []AddressStatus{
	"ACTIVE",
	"DELETED",
}

func (v *AddressStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AddressStatus(value)
	for _, existing := range AllowedAddressStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AddressStatus", value)
}

// NewAddressStatusFromValue returns a pointer to a valid AddressStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAddressStatusFromValue(v string) (*AddressStatus, error) {
	ev := AddressStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AddressStatus: valid values are %v", v, AllowedAddressStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AddressStatus) IsValid() bool {
	for _, existing := range AllowedAddressStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v AddressStatus) String() string {
	return string(v)
}

// Ptr returns reference to AddressStatus value
func (v AddressStatus) Ptr() *AddressStatus {
	return &v
}

type NullableAddressStatus struct {
	value *AddressStatus
	isSet bool
}

func (v NullableAddressStatus) Get() *AddressStatus {
	return v.value
}

func (v *NullableAddressStatus) Set(val *AddressStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressStatus(val *AddressStatus) *NullableAddressStatus {
	return &NullableAddressStatus{value: val, isSet: true}
}

func (v NullableAddressStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
