/*
Payment Requests

This API is used for Payment Requests

API version: 1.45.2
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the EWallet type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EWallet{}

// EWallet Ewallet Payment Method Details
type EWallet struct {
	ChannelCode *EWalletChannelCode `json:"channel_code,omitempty"`
	ChannelProperties *EWalletChannelProperties `json:"channel_properties,omitempty"`
	Account *EWalletAccount `json:"account,omitempty"`
}

// NewEWallet instantiates a new EWallet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEWallet() *EWallet {
	this := EWallet{}
	return &this
}

// NewEWalletWithDefaults instantiates a new EWallet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEWalletWithDefaults() *EWallet {
	this := EWallet{}
	return &this
}

// GetChannelCode returns the ChannelCode field value if set, zero value otherwise.
func (o *EWallet) GetChannelCode() EWalletChannelCode {
	if o == nil || utils.IsNil(o.ChannelCode) {
		var ret EWalletChannelCode
		return ret
	}
	return *o.ChannelCode
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EWallet) GetChannelCodeOk() (*EWalletChannelCode, bool) {
	if o == nil || utils.IsNil(o.ChannelCode) {
		return nil, false
	}
	return o.ChannelCode, true
}

// HasChannelCode returns a boolean if a field has been set.
func (o *EWallet) HasChannelCode() bool {
	if o != nil && !utils.IsNil(o.ChannelCode) {
		return true
	}

	return false
}

// SetChannelCode gets a reference to the given EWalletChannelCode and assigns it to the ChannelCode field.
func (o *EWallet) SetChannelCode(v EWalletChannelCode) {
	o.ChannelCode = &v
}

// GetChannelProperties returns the ChannelProperties field value if set, zero value otherwise.
func (o *EWallet) GetChannelProperties() EWalletChannelProperties {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		var ret EWalletChannelProperties
		return ret
	}
	return *o.ChannelProperties
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EWallet) GetChannelPropertiesOk() (*EWalletChannelProperties, bool) {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		return nil, false
	}
	return o.ChannelProperties, true
}

// HasChannelProperties returns a boolean if a field has been set.
func (o *EWallet) HasChannelProperties() bool {
	if o != nil && !utils.IsNil(o.ChannelProperties) {
		return true
	}

	return false
}

// SetChannelProperties gets a reference to the given EWalletChannelProperties and assigns it to the ChannelProperties field.
func (o *EWallet) SetChannelProperties(v EWalletChannelProperties) {
	o.ChannelProperties = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *EWallet) GetAccount() EWalletAccount {
	if o == nil || utils.IsNil(o.Account) {
		var ret EWalletAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EWallet) GetAccountOk() (*EWalletAccount, bool) {
	if o == nil || utils.IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *EWallet) HasAccount() bool {
	if o != nil && !utils.IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given EWalletAccount and assigns it to the Account field.
func (o *EWallet) SetAccount(v EWalletAccount) {
	o.Account = &v
}

func (o EWallet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EWallet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ChannelCode) {
		toSerialize["channel_code"] = o.ChannelCode
	}
	if !utils.IsNil(o.ChannelProperties) {
		toSerialize["channel_properties"] = o.ChannelProperties
	}
	if !utils.IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	return toSerialize, nil
}

type NullableEWallet struct {
	value *EWallet
	isSet bool
}

func (v NullableEWallet) Get() *EWallet {
	return v.value
}

func (v *NullableEWallet) Set(val *EWallet) {
	v.value = val
	v.isSet = true
}

func (v NullableEWallet) IsSet() bool {
	return v.isSet
}

func (v *NullableEWallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEWallet(val *EWallet) *NullableEWallet {
	return &NullableEWallet{value: val, isSet: true}
}

func (v NullableEWallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEWallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


