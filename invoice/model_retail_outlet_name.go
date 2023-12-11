/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.6.0
*/


package invoice

import (
	"encoding/json"
	
	"fmt"
)

// RetailOutletName Representing the available retail outlet channels used for invoice-related transactions.
type RetailOutletName string

// List of RetailOutletName
const (
	RETAILOUTLETNAME_ALFAMART RetailOutletName = "ALFAMART"
	RETAILOUTLETNAME_INDOMARET RetailOutletName = "INDOMARET"
	RETAILOUTLETNAME__7_ELEVEN RetailOutletName = "7ELEVEN"
	RETAILOUTLETNAME_CEBUANA RetailOutletName = "CEBUANA"
	RETAILOUTLETNAME_DP_ECPAY_LOAN RetailOutletName = "DP_ECPAY_LOAN"
	RETAILOUTLETNAME_DP_MLHUILLIER RetailOutletName = "DP_MLHUILLIER"
	RETAILOUTLETNAME_DP_PALAWAN RetailOutletName = "DP_PALAWAN"
	RETAILOUTLETNAME_DP_ECPAY_SCHOOL RetailOutletName = "DP_ECPAY_SCHOOL"
	RETAILOUTLETNAME_LBC RetailOutletName = "LBC"
    RETAILOUTLETNAME_XENDIT_ENUM_DEFAULT_FALLBACK RetailOutletName = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of RetailOutletName enum
var AllowedRetailOutletNameEnumValues = []RetailOutletName{
	"ALFAMART",
	"INDOMARET",
	"7ELEVEN",
	"CEBUANA",
	"DP_ECPAY_LOAN",
	"DP_MLHUILLIER",
	"DP_PALAWAN",
	"DP_ECPAY_SCHOOL",
	"LBC",
    "UNKNOWN_ENUM_VALUE",
}

func (v *RetailOutletName) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RetailOutletName(value)
	for _, existing := range AllowedRetailOutletNameEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = RETAILOUTLETNAME_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewRetailOutletNameFromValue returns a pointer to a valid RetailOutletName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRetailOutletNameFromValue(v string) (*RetailOutletName, error) {
	ev := RetailOutletName(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RetailOutletName: valid values are %v", v, AllowedRetailOutletNameEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RetailOutletName) IsValid() bool {
	for _, existing := range AllowedRetailOutletNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v RetailOutletName) String() string {
	return string(v)
}

// Ptr returns reference to RetailOutletName value
func (v RetailOutletName) Ptr() *RetailOutletName {
	return &v
}

type NullableRetailOutletName struct {
	value *RetailOutletName
	isSet bool
}

func (v NullableRetailOutletName) Get() *RetailOutletName {
	return v.value
}

func (v *NullableRetailOutletName) Set(val *RetailOutletName) {
	v.value = val
	v.isSet = true
}

func (v NullableRetailOutletName) IsSet() bool {
	return v.isSet
}

func (v *NullableRetailOutletName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetailOutletName(val *RetailOutletName) *NullableRetailOutletName {
	return &NullableRetailOutletName{value: val, isSet: true}
}

func (v NullableRetailOutletName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetailOutletName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
