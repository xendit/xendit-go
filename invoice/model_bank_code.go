/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.8.7
*/


package invoice

import (
	"encoding/json"
	
	"fmt"
)

// BankCode Representing the available bank channels used for invoice-related transactions.
type BankCode string

// List of BankCode
const (
	BANKCODE_BCA BankCode = "BCA"
	BANKCODE_BNI BankCode = "BNI"
	BANKCODE_BRI BankCode = "BRI"
	BANKCODE_MANDIRI BankCode = "MANDIRI"
	BANKCODE_PERMATA BankCode = "PERMATA"
	BANKCODE_BSI BankCode = "BSI"
	BANKCODE_BJB BankCode = "BJB"
	BANKCODE_SAHABAT_SAMPOERNA BankCode = "SAHABAT_SAMPOERNA"
	BANKCODE_CIMB BankCode = "CIMB"
	BANKCODE_VIETCAPITAL BankCode = "VIETCAPITAL"
	BANKCODE_WOORI BankCode = "WOORI"
	BANKCODE_PV BankCode = "PV"
	BANKCODE_MSB BankCode = "MSB"
	BANKCODE_VPB BankCode = "VPB"
	BANKCODE_BIDV BankCode = "BIDV"
	BANKCODE_CAKE BankCode = "CAKE"
	BANKCODE_BNC BankCode = "BNC"
	BANKCODE_HANA BankCode = "HANA"
	BANKCODE_MUAMALAT BankCode = "MUAMALAT"
    BANKCODE_XENDIT_ENUM_DEFAULT_FALLBACK BankCode = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of BankCode enum
var AllowedBankCodeEnumValues = []BankCode{
	"BCA",
	"BNI",
	"BRI",
	"MANDIRI",
	"PERMATA",
	"BSI",
	"BJB",
	"SAHABAT_SAMPOERNA",
	"CIMB",
	"VIETCAPITAL",
	"WOORI",
	"PV",
	"MSB",
	"VPB",
	"BIDV",
	"CAKE",
	"BNC",
	"HANA",
	"MUAMALAT",
    "UNKNOWN_ENUM_VALUE",
}

func (v *BankCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BankCode(value)
	for _, existing := range AllowedBankCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = BANKCODE_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewBankCodeFromValue returns a pointer to a valid BankCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBankCodeFromValue(v string) (*BankCode, error) {
	ev := BankCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BankCode: valid values are %v", v, AllowedBankCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BankCode) IsValid() bool {
	for _, existing := range AllowedBankCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v BankCode) String() string {
	return string(v)
}

// Ptr returns reference to BankCode value
func (v BankCode) Ptr() *BankCode {
	return &v
}

type NullableBankCode struct {
	value *BankCode
	isSet bool
}

func (v NullableBankCode) Get() *BankCode {
	return v.value
}

func (v *NullableBankCode) Set(val *BankCode) {
	v.value = val
	v.isSet = true
}

func (v NullableBankCode) IsSet() bool {
	return v.isSet
}

func (v *NullableBankCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankCode(val *BankCode) *NullableBankCode {
	return &NullableBankCode{value: val, isSet: true}
}

func (v NullableBankCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
