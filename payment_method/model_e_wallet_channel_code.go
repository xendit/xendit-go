/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.128.0
*/


package payment_method

import (
	"encoding/json"
	
	"fmt"
)

// EWalletChannelCode EWallet Channel Code
type EWalletChannelCode string

// List of EWalletChannelCode
const (
	EWALLETCHANNELCODE_GCASH EWalletChannelCode = "GCASH"
	EWALLETCHANNELCODE_GRABPAY EWalletChannelCode = "GRABPAY"
	EWALLETCHANNELCODE_PAYMAYA EWalletChannelCode = "PAYMAYA"
	EWALLETCHANNELCODE_OVO EWalletChannelCode = "OVO"
	EWALLETCHANNELCODE_DANA EWalletChannelCode = "DANA"
	EWALLETCHANNELCODE_LINKAJA EWalletChannelCode = "LINKAJA"
	EWALLETCHANNELCODE_SHOPEEPAY EWalletChannelCode = "SHOPEEPAY"
	EWALLETCHANNELCODE_SAKUKU EWalletChannelCode = "SAKUKU"
	EWALLETCHANNELCODE_NEXCASH EWalletChannelCode = "NEXCASH"
	EWALLETCHANNELCODE_ASTRAPAY EWalletChannelCode = "ASTRAPAY"
	EWALLETCHANNELCODE_JENIUSPAY EWalletChannelCode = "JENIUSPAY"
	EWALLETCHANNELCODE_APPOTA EWalletChannelCode = "APPOTA"
	EWALLETCHANNELCODE_MOMO EWalletChannelCode = "MOMO"
	EWALLETCHANNELCODE_VNPTWALLET EWalletChannelCode = "VNPTWALLET"
	EWALLETCHANNELCODE_VIETTELPAY EWalletChannelCode = "VIETTELPAY"
	EWALLETCHANNELCODE_ZALOPAY EWalletChannelCode = "ZALOPAY"
	EWALLETCHANNELCODE_WECHATPAY EWalletChannelCode = "WECHATPAY"
	EWALLETCHANNELCODE_LINEPAY EWalletChannelCode = "LINEPAY"
	EWALLETCHANNELCODE_TRUEMONEY EWalletChannelCode = "TRUEMONEY"
	EWALLETCHANNELCODE_ALIPAY EWalletChannelCode = "ALIPAY"
	EWALLETCHANNELCODE_TOUCHNGO EWalletChannelCode = "TOUCHNGO"
    EWALLETCHANNELCODE_XENDIT_ENUM_DEFAULT_FALLBACK EWalletChannelCode = "UNKNOWN_ENUM_VALUE"
)

// All allowed values of EWalletChannelCode enum
var AllowedEWalletChannelCodeEnumValues = []EWalletChannelCode{
	"GCASH",
	"GRABPAY",
	"PAYMAYA",
	"OVO",
	"DANA",
	"LINKAJA",
	"SHOPEEPAY",
	"SAKUKU",
	"NEXCASH",
	"ASTRAPAY",
	"JENIUSPAY",
	"APPOTA",
	"MOMO",
	"VNPTWALLET",
	"VIETTELPAY",
	"ZALOPAY",
	"WECHATPAY",
	"LINEPAY",
	"TRUEMONEY",
	"ALIPAY",
	"TOUCHNGO",
    "UNKNOWN_ENUM_VALUE",
}

func (v *EWalletChannelCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EWalletChannelCode(value)
	for _, existing := range AllowedEWalletChannelCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

    *v = EWALLETCHANNELCODE_XENDIT_ENUM_DEFAULT_FALLBACK
    return nil
}

// NewEWalletChannelCodeFromValue returns a pointer to a valid EWalletChannelCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEWalletChannelCodeFromValue(v string) (*EWalletChannelCode, error) {
	ev := EWalletChannelCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EWalletChannelCode: valid values are %v", v, AllowedEWalletChannelCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EWalletChannelCode) IsValid() bool {
	for _, existing := range AllowedEWalletChannelCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v EWalletChannelCode) String() string {
	return string(v)
}

// Ptr returns reference to EWalletChannelCode value
func (v EWalletChannelCode) Ptr() *EWalletChannelCode {
	return &v
}

type NullableEWalletChannelCode struct {
	value *EWalletChannelCode
	isSet bool
}

func (v NullableEWalletChannelCode) Get() *EWalletChannelCode {
	return v.value
}

func (v *NullableEWalletChannelCode) Set(val *EWalletChannelCode) {
	v.value = val
	v.isSet = true
}

func (v NullableEWalletChannelCode) IsSet() bool {
	return v.isSet
}

func (v *NullableEWalletChannelCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEWalletChannelCode(val *EWalletChannelCode) *NullableEWalletChannelCode {
	return &NullableEWalletChannelCode{value: val, isSet: true}
}

func (v NullableEWalletChannelCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEWalletChannelCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
