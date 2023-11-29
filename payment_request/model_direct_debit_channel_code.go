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

// DirectDebitChannelCode Direct Debit Channel Code
type DirectDebitChannelCode string

// List of DirectDebitChannelCode
const (
	DIRECTDEBITCHANNELCODE_BCA_KLIKPAY DirectDebitChannelCode = "BCA_KLIKPAY"
	DIRECTDEBITCHANNELCODE_BCA_ONEKLIK DirectDebitChannelCode = "BCA_ONEKLIK"
	DIRECTDEBITCHANNELCODE_BRI DirectDebitChannelCode = "BRI"
	DIRECTDEBITCHANNELCODE_BNI DirectDebitChannelCode = "BNI"
	DIRECTDEBITCHANNELCODE_MANDIRI DirectDebitChannelCode = "MANDIRI"
	DIRECTDEBITCHANNELCODE_BPI DirectDebitChannelCode = "BPI"
	DIRECTDEBITCHANNELCODE_BDO DirectDebitChannelCode = "BDO"
	DIRECTDEBITCHANNELCODE_CIMBNIAGA DirectDebitChannelCode = "CIMBNIAGA"
	DIRECTDEBITCHANNELCODE_MTB DirectDebitChannelCode = "MTB"
	DIRECTDEBITCHANNELCODE_RCBC DirectDebitChannelCode = "RCBC"
	DIRECTDEBITCHANNELCODE_UBP DirectDebitChannelCode = "UBP"
	DIRECTDEBITCHANNELCODE_AUTODEBIT_UBP DirectDebitChannelCode = "AUTODEBIT_UBP"
	DIRECTDEBITCHANNELCODE_CHINABANK DirectDebitChannelCode = "CHINABANK"
	DIRECTDEBITCHANNELCODE_BAY DirectDebitChannelCode = "BAY"
	DIRECTDEBITCHANNELCODE_KTB DirectDebitChannelCode = "KTB"
	DIRECTDEBITCHANNELCODE_BBL DirectDebitChannelCode = "BBL"
	DIRECTDEBITCHANNELCODE_SCB DirectDebitChannelCode = "SCB"
	DIRECTDEBITCHANNELCODE_KBANK_MB DirectDebitChannelCode = "KBANK_MB"
	DIRECTDEBITCHANNELCODE_BAY_MB DirectDebitChannelCode = "BAY_MB"
	DIRECTDEBITCHANNELCODE_KTB_MB DirectDebitChannelCode = "KTB_MB"
	DIRECTDEBITCHANNELCODE_BBL_MB DirectDebitChannelCode = "BBL_MB"
	DIRECTDEBITCHANNELCODE_SCB_MB DirectDebitChannelCode = "SCB_MB"
	DIRECTDEBITCHANNELCODE_BDO_EPAY DirectDebitChannelCode = "BDO_EPAY"
	DIRECTDEBITCHANNELCODE_AFFIN_FPX DirectDebitChannelCode = "AFFIN_FPX"
	DIRECTDEBITCHANNELCODE_AGRO_FPX DirectDebitChannelCode = "AGRO_FPX"
	DIRECTDEBITCHANNELCODE_ALLIANCE_FPX DirectDebitChannelCode = "ALLIANCE_FPX"
	DIRECTDEBITCHANNELCODE_AMBANK_FPX DirectDebitChannelCode = "AMBANK_FPX"
	DIRECTDEBITCHANNELCODE_ISLAM_FPX DirectDebitChannelCode = "ISLAM_FPX"
	DIRECTDEBITCHANNELCODE_MUAMALAT_FPX DirectDebitChannelCode = "MUAMALAT_FPX"
	DIRECTDEBITCHANNELCODE_BOC_FPX DirectDebitChannelCode = "BOC_FPX"
	DIRECTDEBITCHANNELCODE_RAKYAT_FPX DirectDebitChannelCode = "RAKYAT_FPX"
	DIRECTDEBITCHANNELCODE_BSN_FPX DirectDebitChannelCode = "BSN_FPX"
	DIRECTDEBITCHANNELCODE_CIMB_FPX DirectDebitChannelCode = "CIMB_FPX"
	DIRECTDEBITCHANNELCODE_HLB_FPX DirectDebitChannelCode = "HLB_FPX"
	DIRECTDEBITCHANNELCODE_HSBC_FPX DirectDebitChannelCode = "HSBC_FPX"
	DIRECTDEBITCHANNELCODE_KFH_FPX DirectDebitChannelCode = "KFH_FPX"
	DIRECTDEBITCHANNELCODE_MAYB2_E_FPX DirectDebitChannelCode = "MAYB2E_FPX"
	DIRECTDEBITCHANNELCODE_MAYB2_U_FPX DirectDebitChannelCode = "MAYB2U_FPX"
	DIRECTDEBITCHANNELCODE_OCBC_FPX DirectDebitChannelCode = "OCBC_FPX"
	DIRECTDEBITCHANNELCODE_PUBLIC_FPX DirectDebitChannelCode = "PUBLIC_FPX"
	DIRECTDEBITCHANNELCODE_RHB_FPX DirectDebitChannelCode = "RHB_FPX"
	DIRECTDEBITCHANNELCODE_SCH_FPX DirectDebitChannelCode = "SCH_FPX"
	DIRECTDEBITCHANNELCODE_UOB_FPX DirectDebitChannelCode = "UOB_FPX"
	DIRECTDEBITCHANNELCODE_AFFIN_FPX_BUSINESS DirectDebitChannelCode = "AFFIN_FPX_BUSINESS"
	DIRECTDEBITCHANNELCODE_AGRO_FPX_BUSINESS DirectDebitChannelCode = "AGRO_FPX_BUSINESS"
	DIRECTDEBITCHANNELCODE_ALLIANCE_FPX_BUSINESS DirectDebitChannelCode = "ALLIANCE_FPX_BUSINESS"
	DIRECTDEBITCHANNELCODE_AMBANK_FPX_BUSINESS DirectDebitChannelCode = "AMBANK_FPX_BUSINESS"
	DIRECTDEBITCHANNELCODE_ISLAM_FPX_BUSINESS DirectDebitChannelCode = "ISLAM_FPX_BUSINESS"
	DIRECTDEBITCHANNELCODE_MUAMALAT_FPX_BUSINESS DirectDebitChannelCode = "MUAMALAT_FPX_BUSINESS"
	DIRECTDEBITCHANNELCODE_BNP_FPX_BUSINESS DirectDebitChannelCode = "BNP_FPX_BUSINESS"
	DIRECTDEBITCHANNELCODE_CIMB_FPX_BUSINESS DirectDebitChannelCode = "CIMB_FPX_BUSINESS"
	DIRECTDEBITCHANNELCODE_CITIBANK_FPX_BUSINESS DirectDebitChannelCode = "CITIBANK_FPX_BUSINESS"
	DIRECTDEBITCHANNELCODE_DEUTSCHE_FPX_BUSINESS DirectDebitChannelCode = "DEUTSCHE_FPX_BUSINESS"
	DIRECTDEBITCHANNELCODE_HLB_FPX_BUSINESS DirectDebitChannelCode = "HLB_FPX_BUSINESS"
	DIRECTDEBITCHANNELCODE_HSBC_FPX_BUSINESS DirectDebitChannelCode = "HSBC_FPX_BUSINESS"
	DIRECTDEBITCHANNELCODE_RAKYAT_FPX_BUSINESS DirectDebitChannelCode = "RAKYAT_FPX_BUSINESS"
	DIRECTDEBITCHANNELCODE_KFH_FPX_BUSINESS DirectDebitChannelCode = "KFH_FPX_BUSINESS"
	DIRECTDEBITCHANNELCODE_MAYB2_E_FPX_BUSINESS DirectDebitChannelCode = "MAYB2E_FPX_BUSINESS"
	DIRECTDEBITCHANNELCODE_OCBC_FPX_BUSINESS DirectDebitChannelCode = "OCBC_FPX_BUSINESS"
	DIRECTDEBITCHANNELCODE_PUBLIC_FPX_BUSINESS DirectDebitChannelCode = "PUBLIC_FPX_BUSINESS"
	DIRECTDEBITCHANNELCODE_RHB_FPX_BUSINESS DirectDebitChannelCode = "RHB_FPX_BUSINESS"
	DIRECTDEBITCHANNELCODE_SCH_FPX_BUSINESS DirectDebitChannelCode = "SCH_FPX_BUSINESS"
	DIRECTDEBITCHANNELCODE_UOB_FPX_BUSINESS DirectDebitChannelCode = "UOB_FPX_BUSINESS"
    DIRECTDEBITCHANNELCODE_XENDIT_ENUM_DEFAULT_FALLBACK DirectDebitChannelCode = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of DirectDebitChannelCode enum
var AllowedDirectDebitChannelCodeEnumValues = []DirectDebitChannelCode{
	"BCA_KLIKPAY",
	"BCA_ONEKLIK",
	"BRI",
	"BNI",
	"MANDIRI",
	"BPI",
	"BDO",
	"CIMBNIAGA",
	"MTB",
	"RCBC",
	"UBP",
	"AUTODEBIT_UBP",
	"CHINABANK",
	"BAY",
	"KTB",
	"BBL",
	"SCB",
	"KBANK_MB",
	"BAY_MB",
	"KTB_MB",
	"BBL_MB",
	"SCB_MB",
	"BDO_EPAY",
	"AFFIN_FPX",
	"AGRO_FPX",
	"ALLIANCE_FPX",
	"AMBANK_FPX",
	"ISLAM_FPX",
	"MUAMALAT_FPX",
	"BOC_FPX",
	"RAKYAT_FPX",
	"BSN_FPX",
	"CIMB_FPX",
	"HLB_FPX",
	"HSBC_FPX",
	"KFH_FPX",
	"MAYB2E_FPX",
	"MAYB2U_FPX",
	"OCBC_FPX",
	"PUBLIC_FPX",
	"RHB_FPX",
	"SCH_FPX",
	"UOB_FPX",
	"AFFIN_FPX_BUSINESS",
	"AGRO_FPX_BUSINESS",
	"ALLIANCE_FPX_BUSINESS",
	"AMBANK_FPX_BUSINESS",
	"ISLAM_FPX_BUSINESS",
	"MUAMALAT_FPX_BUSINESS",
	"BNP_FPX_BUSINESS",
	"CIMB_FPX_BUSINESS",
	"CITIBANK_FPX_BUSINESS",
	"DEUTSCHE_FPX_BUSINESS",
	"HLB_FPX_BUSINESS",
	"HSBC_FPX_BUSINESS",
	"RAKYAT_FPX_BUSINESS",
	"KFH_FPX_BUSINESS",
	"MAYB2E_FPX_BUSINESS",
	"OCBC_FPX_BUSINESS",
	"PUBLIC_FPX_BUSINESS",
	"RHB_FPX_BUSINESS",
	"SCH_FPX_BUSINESS",
	"UOB_FPX_BUSINESS",
    "UNKNOWN_ENUM_VALUE",
}

func (v *DirectDebitChannelCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DirectDebitChannelCode(value)
	for _, existing := range AllowedDirectDebitChannelCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = DIRECTDEBITCHANNELCODE_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewDirectDebitChannelCodeFromValue returns a pointer to a valid DirectDebitChannelCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDirectDebitChannelCodeFromValue(v string) (*DirectDebitChannelCode, error) {
	ev := DirectDebitChannelCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DirectDebitChannelCode: valid values are %v", v, AllowedDirectDebitChannelCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DirectDebitChannelCode) IsValid() bool {
	for _, existing := range AllowedDirectDebitChannelCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v DirectDebitChannelCode) String() string {
	return string(v)
}

// Ptr returns reference to DirectDebitChannelCode value
func (v DirectDebitChannelCode) Ptr() *DirectDebitChannelCode {
	return &v
}

type NullableDirectDebitChannelCode struct {
	value *DirectDebitChannelCode
	isSet bool
}

func (v NullableDirectDebitChannelCode) Get() *DirectDebitChannelCode {
	return v.value
}

func (v *NullableDirectDebitChannelCode) Set(val *DirectDebitChannelCode) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectDebitChannelCode) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectDebitChannelCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectDebitChannelCode(val *DirectDebitChannelCode) *NullableDirectDebitChannelCode {
	return &NullableDirectDebitChannelCode{value: val, isSet: true}
}

func (v NullableDirectDebitChannelCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectDebitChannelCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
