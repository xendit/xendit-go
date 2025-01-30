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

// DirectDebitType Representing the available Direct Debit channels used for invoice-related transactions.
type DirectDebitType string

// List of DirectDebitType
const (
	DIRECTDEBITTYPE_BA_BRI DirectDebitType = "BA_BRI"
	DIRECTDEBITTYPE_DC_BRI DirectDebitType = "DC_BRI"
	DIRECTDEBITTYPE_DD_BRI DirectDebitType = "DD_BRI"
	DIRECTDEBITTYPE_DD_MANDIRI DirectDebitType = "DD_MANDIRI"
	DIRECTDEBITTYPE_BA_BPI DirectDebitType = "BA_BPI"
	DIRECTDEBITTYPE_DC_BPI DirectDebitType = "DC_BPI"
	DIRECTDEBITTYPE_DD_BPI DirectDebitType = "DD_BPI"
	DIRECTDEBITTYPE_BA_UBP DirectDebitType = "BA_UBP"
	DIRECTDEBITTYPE_DC_UBP DirectDebitType = "DC_UBP"
	DIRECTDEBITTYPE_DD_UBP DirectDebitType = "DD_UBP"
	DIRECTDEBITTYPE_BCA_KLIKPAY DirectDebitType = "BCA_KLIKPAY"
	DIRECTDEBITTYPE_BA_BCA_KLIKPAY DirectDebitType = "BA_BCA_KLIKPAY"
	DIRECTDEBITTYPE_DC_BCA_KLIKPAY DirectDebitType = "DC_BCA_KLIKPAY"
	DIRECTDEBITTYPE_DD_BCA_KLIKPAY DirectDebitType = "DD_BCA_KLIKPAY"
	DIRECTDEBITTYPE_DD_BDO_EPAY DirectDebitType = "DD_BDO_EPAY"
	DIRECTDEBITTYPE_DD_RCBC DirectDebitType = "DD_RCBC"
	DIRECTDEBITTYPE_DD_CHINABANK DirectDebitType = "DD_CHINABANK"
	DIRECTDEBITTYPE_BA_CHINABANK DirectDebitType = "BA_CHINABANK"
	DIRECTDEBITTYPE_DC_CHINABANK DirectDebitType = "DC_CHINABANK"
	DIRECTDEBITTYPE_DD_PUBLIC_FPX DirectDebitType = "DD_PUBLIC_FPX"
	DIRECTDEBITTYPE_DD_AMBANK_FPX DirectDebitType = "DD_AMBANK_FPX"
	DIRECTDEBITTYPE_DD_KFH_FPX DirectDebitType = "DD_KFH_FPX"
	DIRECTDEBITTYPE_DD_AGRO_FPX DirectDebitType = "DD_AGRO_FPX"
	DIRECTDEBITTYPE_DD_AFFIN_FPX DirectDebitType = "DD_AFFIN_FPX"
	DIRECTDEBITTYPE_DD_ALLIANCE_FPX DirectDebitType = "DD_ALLIANCE_FPX"
	DIRECTDEBITTYPE_DD_MUAMALAT_FPX DirectDebitType = "DD_MUAMALAT_FPX"
	DIRECTDEBITTYPE_DD_HLB_FPX DirectDebitType = "DD_HLB_FPX"
	DIRECTDEBITTYPE_DD_ISLAM_FPX DirectDebitType = "DD_ISLAM_FPX"
	DIRECTDEBITTYPE_DD_RAKYAT_FPX DirectDebitType = "DD_RAKYAT_FPX"
	DIRECTDEBITTYPE_DD_CIMB_FPX DirectDebitType = "DD_CIMB_FPX"
	DIRECTDEBITTYPE_DD_UOB_FPX DirectDebitType = "DD_UOB_FPX"
	DIRECTDEBITTYPE_DD_BOC_FPX DirectDebitType = "DD_BOC_FPX"
	DIRECTDEBITTYPE_DD_BSN_FPX DirectDebitType = "DD_BSN_FPX"
	DIRECTDEBITTYPE_DD_OCBC_FPX DirectDebitType = "DD_OCBC_FPX"
	DIRECTDEBITTYPE_DD_HSBC_FPX DirectDebitType = "DD_HSBC_FPX"
	DIRECTDEBITTYPE_DD_SCH_FPX DirectDebitType = "DD_SCH_FPX"
	DIRECTDEBITTYPE_DD_MAYB2_U_FPX DirectDebitType = "DD_MAYB2U_FPX"
	DIRECTDEBITTYPE_DD_RHB_FPX DirectDebitType = "DD_RHB_FPX"
	DIRECTDEBITTYPE_DD_UOB_FPX_BUSINESS DirectDebitType = "DD_UOB_FPX_BUSINESS"
	DIRECTDEBITTYPE_DD_AGRO_FPX_BUSINESS DirectDebitType = "DD_AGRO_FPX_BUSINESS"
	DIRECTDEBITTYPE_DD_ALLIANCE_FPX_BUSINESS DirectDebitType = "DD_ALLIANCE_FPX_BUSINESS"
	DIRECTDEBITTYPE_DD_AMBANK_FPX_BUSINESS DirectDebitType = "DD_AMBANK_FPX_BUSINESS"
	DIRECTDEBITTYPE_DD_ISLAM_FPX_BUSINESS DirectDebitType = "DD_ISLAM_FPX_BUSINESS"
	DIRECTDEBITTYPE_DD_MUAMALAT_FPX_BUSINESS DirectDebitType = "DD_MUAMALAT_FPX_BUSINESS"
	DIRECTDEBITTYPE_DD_HLB_FPX_BUSINESS DirectDebitType = "DD_HLB_FPX_BUSINESS"
	DIRECTDEBITTYPE_DD_HSBC_FPX_BUSINESS DirectDebitType = "DD_HSBC_FPX_BUSINESS"
	DIRECTDEBITTYPE_DD_RAKYAT_FPX_BUSINESS DirectDebitType = "DD_RAKYAT_FPX_BUSINESS"
	DIRECTDEBITTYPE_DD_KFH_FPX_BUSINESS DirectDebitType = "DD_KFH_FPX_BUSINESS"
	DIRECTDEBITTYPE_DD_OCBC_FPX_BUSINESS DirectDebitType = "DD_OCBC_FPX_BUSINESS"
	DIRECTDEBITTYPE_DD_PUBLIC_FPX_BUSINESS DirectDebitType = "DD_PUBLIC_FPX_BUSINESS"
	DIRECTDEBITTYPE_DD_RHB_FPX_BUSINESS DirectDebitType = "DD_RHB_FPX_BUSINESS"
	DIRECTDEBITTYPE_DD_SCH_FPX_BUSINESS DirectDebitType = "DD_SCH_FPX_BUSINESS"
	DIRECTDEBITTYPE_DD_CITIBANK_FPX_BUSINESS DirectDebitType = "DD_CITIBANK_FPX_BUSINESS"
	DIRECTDEBITTYPE_DD_BNP_FPX_BUSINESS DirectDebitType = "DD_BNP_FPX_BUSINESS"
	DIRECTDEBITTYPE_DD_DEUTSCHE_FPX_BUSINESS DirectDebitType = "DD_DEUTSCHE_FPX_BUSINESS"
	DIRECTDEBITTYPE_DD_MAYB2_E_FPX_BUSINESS DirectDebitType = "DD_MAYB2E_FPX_BUSINESS"
	DIRECTDEBITTYPE_DD_CIMB_FPX_BUSINESS DirectDebitType = "DD_CIMB_FPX_BUSINESS"
	DIRECTDEBITTYPE_DD_AFFIN_FPX_BUSINESS DirectDebitType = "DD_AFFIN_FPX_BUSINESS"
    DIRECTDEBITTYPE_XENDIT_ENUM_DEFAULT_FALLBACK DirectDebitType = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of DirectDebitType enum
var AllowedDirectDebitTypeEnumValues = []DirectDebitType{
	"BA_BRI",
	"DC_BRI",
	"DD_BRI",
	"DD_MANDIRI",
	"BA_BPI",
	"DC_BPI",
	"DD_BPI",
	"BA_UBP",
	"DC_UBP",
	"DD_UBP",
	"BCA_KLIKPAY",
	"BA_BCA_KLIKPAY",
	"DC_BCA_KLIKPAY",
	"DD_BCA_KLIKPAY",
	"DD_BDO_EPAY",
	"DD_RCBC",
	"DD_CHINABANK",
	"BA_CHINABANK",
	"DC_CHINABANK",
	"DD_PUBLIC_FPX",
	"DD_AMBANK_FPX",
	"DD_KFH_FPX",
	"DD_AGRO_FPX",
	"DD_AFFIN_FPX",
	"DD_ALLIANCE_FPX",
	"DD_MUAMALAT_FPX",
	"DD_HLB_FPX",
	"DD_ISLAM_FPX",
	"DD_RAKYAT_FPX",
	"DD_CIMB_FPX",
	"DD_UOB_FPX",
	"DD_BOC_FPX",
	"DD_BSN_FPX",
	"DD_OCBC_FPX",
	"DD_HSBC_FPX",
	"DD_SCH_FPX",
	"DD_MAYB2U_FPX",
	"DD_RHB_FPX",
	"DD_UOB_FPX_BUSINESS",
	"DD_AGRO_FPX_BUSINESS",
	"DD_ALLIANCE_FPX_BUSINESS",
	"DD_AMBANK_FPX_BUSINESS",
	"DD_ISLAM_FPX_BUSINESS",
	"DD_MUAMALAT_FPX_BUSINESS",
	"DD_HLB_FPX_BUSINESS",
	"DD_HSBC_FPX_BUSINESS",
	"DD_RAKYAT_FPX_BUSINESS",
	"DD_KFH_FPX_BUSINESS",
	"DD_OCBC_FPX_BUSINESS",
	"DD_PUBLIC_FPX_BUSINESS",
	"DD_RHB_FPX_BUSINESS",
	"DD_SCH_FPX_BUSINESS",
	"DD_CITIBANK_FPX_BUSINESS",
	"DD_BNP_FPX_BUSINESS",
	"DD_DEUTSCHE_FPX_BUSINESS",
	"DD_MAYB2E_FPX_BUSINESS",
	"DD_CIMB_FPX_BUSINESS",
	"DD_AFFIN_FPX_BUSINESS",
    "UNKNOWN_ENUM_VALUE",
}

func (v *DirectDebitType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DirectDebitType(value)
	for _, existing := range AllowedDirectDebitTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = DIRECTDEBITTYPE_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewDirectDebitTypeFromValue returns a pointer to a valid DirectDebitType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDirectDebitTypeFromValue(v string) (*DirectDebitType, error) {
	ev := DirectDebitType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DirectDebitType: valid values are %v", v, AllowedDirectDebitTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DirectDebitType) IsValid() bool {
	for _, existing := range AllowedDirectDebitTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v DirectDebitType) String() string {
	return string(v)
}

// Ptr returns reference to DirectDebitType value
func (v DirectDebitType) Ptr() *DirectDebitType {
	return &v
}

type NullableDirectDebitType struct {
	value *DirectDebitType
	isSet bool
}

func (v NullableDirectDebitType) Get() *DirectDebitType {
	return v.value
}

func (v *NullableDirectDebitType) Set(val *DirectDebitType) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectDebitType) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectDebitType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectDebitType(val *DirectDebitType) *NullableDirectDebitType {
	return &NullableDirectDebitType{value: val, isSet: true}
}

func (v NullableDirectDebitType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectDebitType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
