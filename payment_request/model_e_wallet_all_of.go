/*
Payment Requests

This API is used for Payment Requests

API version: 1.70.0
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
)

// checks if the EWalletAllOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EWalletAllOf{}

// EWalletAllOf struct for EWalletAllOf
type EWalletAllOf struct {
	Account *EWalletAccount `json:"account,omitempty"`
}

// NewEWalletAllOf instantiates a new EWalletAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEWalletAllOf() *EWalletAllOf {
	this := EWalletAllOf{}
	return &this
}

// NewEWalletAllOfWithDefaults instantiates a new EWalletAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEWalletAllOfWithDefaults() *EWalletAllOf {
	this := EWalletAllOf{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *EWalletAllOf) GetAccount() EWalletAccount {
	if o == nil || utils.IsNil(o.Account) {
		var ret EWalletAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EWalletAllOf) GetAccountOk() (*EWalletAccount, bool) {
	if o == nil || utils.IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *EWalletAllOf) HasAccount() bool {
	if o != nil && !utils.IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given EWalletAccount and assigns it to the Account field.
func (o *EWalletAllOf) SetAccount(v EWalletAccount) {
	o.Account = &v
}

func (o EWalletAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EWalletAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	return toSerialize, nil
}

type NullableEWalletAllOf struct {
	value *EWalletAllOf
	isSet bool
}

func (v NullableEWalletAllOf) Get() *EWalletAllOf {
	return v.value
}

func (v *NullableEWalletAllOf) Set(val *EWalletAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEWalletAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEWalletAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEWalletAllOf(val *EWalletAllOf) *NullableEWalletAllOf {
	return &NullableEWalletAllOf{value: val, isSet: true}
}

func (v NullableEWalletAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEWalletAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


