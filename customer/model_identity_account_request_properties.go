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

// IdentityAccountRequestProperties struct for IdentityAccountRequestProperties
type IdentityAccountRequestProperties struct {
	AccountBank *AccountBank
	AccountCard *AccountCard
	AccountEwallet *AccountEwallet
	AccountOTC *AccountOTC
	AccountPayLater *AccountPayLater
	AccountQRCode *AccountQRCode
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *IdentityAccountRequestProperties) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AccountBank
	err = json.Unmarshal(data, &dst.AccountBank);
	if err == nil {
		jsonAccountBank, _ := json.Marshal(dst.AccountBank)
		if string(jsonAccountBank) == "{}" { // empty struct
			dst.AccountBank = nil
		} else {
			return nil // data stored in dst.AccountBank, return on the first match
		}
	} else {
		dst.AccountBank = nil
	}

	// try to unmarshal JSON data into AccountCard
	err = json.Unmarshal(data, &dst.AccountCard);
	if err == nil {
		jsonAccountCard, _ := json.Marshal(dst.AccountCard)
		if string(jsonAccountCard) == "{}" { // empty struct
			dst.AccountCard = nil
		} else {
			return nil // data stored in dst.AccountCard, return on the first match
		}
	} else {
		dst.AccountCard = nil
	}

	// try to unmarshal JSON data into AccountEwallet
	err = json.Unmarshal(data, &dst.AccountEwallet);
	if err == nil {
		jsonAccountEwallet, _ := json.Marshal(dst.AccountEwallet)
		if string(jsonAccountEwallet) == "{}" { // empty struct
			dst.AccountEwallet = nil
		} else {
			return nil // data stored in dst.AccountEwallet, return on the first match
		}
	} else {
		dst.AccountEwallet = nil
	}

	// try to unmarshal JSON data into AccountOTC
	err = json.Unmarshal(data, &dst.AccountOTC);
	if err == nil {
		jsonAccountOTC, _ := json.Marshal(dst.AccountOTC)
		if string(jsonAccountOTC) == "{}" { // empty struct
			dst.AccountOTC = nil
		} else {
			return nil // data stored in dst.AccountOTC, return on the first match
		}
	} else {
		dst.AccountOTC = nil
	}

	// try to unmarshal JSON data into AccountPayLater
	err = json.Unmarshal(data, &dst.AccountPayLater);
	if err == nil {
		jsonAccountPayLater, _ := json.Marshal(dst.AccountPayLater)
		if string(jsonAccountPayLater) == "{}" { // empty struct
			dst.AccountPayLater = nil
		} else {
			return nil // data stored in dst.AccountPayLater, return on the first match
		}
	} else {
		dst.AccountPayLater = nil
	}

	// try to unmarshal JSON data into AccountQRCode
	err = json.Unmarshal(data, &dst.AccountQRCode);
	if err == nil {
		jsonAccountQRCode, _ := json.Marshal(dst.AccountQRCode)
		if string(jsonAccountQRCode) == "{}" { // empty struct
			dst.AccountQRCode = nil
		} else {
			return nil // data stored in dst.AccountQRCode, return on the first match
		}
	} else {
		dst.AccountQRCode = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(IdentityAccountRequestProperties)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *IdentityAccountRequestProperties) MarshalJSON() ([]byte, error) {
	if src.AccountBank != nil {
		return json.Marshal(&src.AccountBank)
	}

	if src.AccountCard != nil {
		return json.Marshal(&src.AccountCard)
	}

	if src.AccountEwallet != nil {
		return json.Marshal(&src.AccountEwallet)
	}

	if src.AccountOTC != nil {
		return json.Marshal(&src.AccountOTC)
	}

	if src.AccountPayLater != nil {
		return json.Marshal(&src.AccountPayLater)
	}

	if src.AccountQRCode != nil {
		return json.Marshal(&src.AccountQRCode)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableIdentityAccountRequestProperties struct {
	value *IdentityAccountRequestProperties
	isSet bool
}

func (v NullableIdentityAccountRequestProperties) Get() *IdentityAccountRequestProperties {
	return v.value
}

func (v *NullableIdentityAccountRequestProperties) Set(val *IdentityAccountRequestProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityAccountRequestProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityAccountRequestProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityAccountRequestProperties(val *IdentityAccountRequestProperties) *NullableIdentityAccountRequestProperties {
	return &NullableIdentityAccountRequestProperties{value: val, isSet: true}
}

func (v NullableIdentityAccountRequestProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityAccountRequestProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


