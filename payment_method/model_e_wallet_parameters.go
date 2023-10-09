/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.89.1
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the EWalletParameters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EWalletParameters{}

// EWalletParameters struct for EWalletParameters
type EWalletParameters struct {
	ChannelCode EWalletChannelCode `json:"channel_code"`
	ChannelProperties *EWalletChannelProperties `json:"channel_properties,omitempty"`
	Account *EWalletAccount `json:"account,omitempty"`
}

// NewEWalletParameters instantiates a new EWalletParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEWalletParameters(channelCode EWalletChannelCode) *EWalletParameters {
	this := EWalletParameters{}
	this.ChannelCode = channelCode
	return &this
}

// NewEWalletParametersWithDefaults instantiates a new EWalletParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEWalletParametersWithDefaults() *EWalletParameters {
	this := EWalletParameters{}
	return &this
}

// GetChannelCode returns the ChannelCode field value
func (o *EWalletParameters) GetChannelCode() EWalletChannelCode {
	if o == nil {
		var ret EWalletChannelCode
		return ret
	}

	return o.ChannelCode
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value
// and a boolean to check if the value has been set.
func (o *EWalletParameters) GetChannelCodeOk() (*EWalletChannelCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelCode, true
}

// SetChannelCode sets field value
func (o *EWalletParameters) SetChannelCode(v EWalletChannelCode) {
	o.ChannelCode = v
}

// GetChannelProperties returns the ChannelProperties field value if set, zero value otherwise.
func (o *EWalletParameters) GetChannelProperties() EWalletChannelProperties {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		var ret EWalletChannelProperties
		return ret
	}
	return *o.ChannelProperties
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EWalletParameters) GetChannelPropertiesOk() (*EWalletChannelProperties, bool) {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		return nil, false
	}
	return o.ChannelProperties, true
}

// HasChannelProperties returns a boolean if a field has been set.
func (o *EWalletParameters) HasChannelProperties() bool {
	if o != nil && !utils.IsNil(o.ChannelProperties) {
		return true
	}

	return false
}

// SetChannelProperties gets a reference to the given EWalletChannelProperties and assigns it to the ChannelProperties field.
func (o *EWalletParameters) SetChannelProperties(v EWalletChannelProperties) {
	o.ChannelProperties = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *EWalletParameters) GetAccount() EWalletAccount {
	if o == nil || utils.IsNil(o.Account) {
		var ret EWalletAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EWalletParameters) GetAccountOk() (*EWalletAccount, bool) {
	if o == nil || utils.IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *EWalletParameters) HasAccount() bool {
	if o != nil && !utils.IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given EWalletAccount and assigns it to the Account field.
func (o *EWalletParameters) SetAccount(v EWalletAccount) {
	o.Account = &v
}

func (o EWalletParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EWalletParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel_code"] = o.ChannelCode
	if !utils.IsNil(o.ChannelProperties) {
		toSerialize["channel_properties"] = o.ChannelProperties
	}
	if !utils.IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	return toSerialize, nil
}

type NullableEWalletParameters struct {
	value *EWalletParameters
	isSet bool
}

func (v NullableEWalletParameters) Get() *EWalletParameters {
	return v.value
}

func (v *NullableEWalletParameters) Set(val *EWalletParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableEWalletParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableEWalletParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEWalletParameters(val *EWalletParameters) *NullableEWalletParameters {
	return &NullableEWalletParameters{value: val, isSet: true}
}

func (v NullableEWalletParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEWalletParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


