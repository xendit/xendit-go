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

// EndCustomerStatus the model 'EndCustomerStatus'
type EndCustomerStatus string

// List of EndCustomerStatus
const (
	ENDCUSTOMERSTATUS_ACTIVE EndCustomerStatus = "ACTIVE"
	ENDCUSTOMERSTATUS_INACTIVE EndCustomerStatus = "INACTIVE"
	ENDCUSTOMERSTATUS_PENDING EndCustomerStatus = "PENDING"
	ENDCUSTOMERSTATUS_BLOCKED EndCustomerStatus = "BLOCKED"
	ENDCUSTOMERSTATUS_DELETED EndCustomerStatus = "DELETED"
    ENDCUSTOMERSTATUS_XENDIT_ENUM_DEFAULT_FALLBACK EndCustomerStatus = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of EndCustomerStatus enum
var AllowedEndCustomerStatusEnumValues = []EndCustomerStatus{
	"ACTIVE",
	"INACTIVE",
	"PENDING",
	"BLOCKED",
	"DELETED",
    "UNKNOWN_ENUM_VALUE",
}

func (v *EndCustomerStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EndCustomerStatus(value)
	for _, existing := range AllowedEndCustomerStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = ENDCUSTOMERSTATUS_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewEndCustomerStatusFromValue returns a pointer to a valid EndCustomerStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEndCustomerStatusFromValue(v string) (*EndCustomerStatus, error) {
	ev := EndCustomerStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EndCustomerStatus: valid values are %v", v, AllowedEndCustomerStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EndCustomerStatus) IsValid() bool {
	for _, existing := range AllowedEndCustomerStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v EndCustomerStatus) String() string {
	return string(v)
}

// Ptr returns reference to EndCustomerStatus value
func (v EndCustomerStatus) Ptr() *EndCustomerStatus {
	return &v
}

type NullableEndCustomerStatus struct {
	value *EndCustomerStatus
	isSet bool
}

func (v NullableEndCustomerStatus) Get() *EndCustomerStatus {
	return v.value
}

func (v *NullableEndCustomerStatus) Set(val *EndCustomerStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEndCustomerStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEndCustomerStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndCustomerStatus(val *EndCustomerStatus) *NullableEndCustomerStatus {
	return &NullableEndCustomerStatus{value: val, isSet: true}
}

func (v NullableEndCustomerStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndCustomerStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
