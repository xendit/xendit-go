/*
Payment Requests

This API is used for Payment Requests

API version: 1.59.0
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v5/utils"
)

// checks if the EWalletParameters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EWalletParameters{}

// EWalletParameters struct for EWalletParameters
type EWalletParameters struct {
	ChannelCode *EWalletChannelCode `json:"channel_code,omitempty"`
	ChannelProperties *EWalletChannelProperties `json:"channel_properties,omitempty"`
}

// NewEWalletParameters instantiates a new EWalletParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEWalletParameters() *EWalletParameters {
	this := EWalletParameters{}
	return &this
}

// NewEWalletParametersWithDefaults instantiates a new EWalletParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEWalletParametersWithDefaults() *EWalletParameters {
	this := EWalletParameters{}
	return &this
}

// GetChannelCode returns the ChannelCode field value if set, zero value otherwise.
func (o *EWalletParameters) GetChannelCode() EWalletChannelCode {
	if o == nil || utils.IsNil(o.ChannelCode) {
		var ret EWalletChannelCode
		return ret
	}
	return *o.ChannelCode
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EWalletParameters) GetChannelCodeOk() (*EWalletChannelCode, bool) {
	if o == nil || utils.IsNil(o.ChannelCode) {
		return nil, false
	}
	return o.ChannelCode, true
}

// HasChannelCode returns a boolean if a field has been set.
func (o *EWalletParameters) HasChannelCode() bool {
	if o != nil && !utils.IsNil(o.ChannelCode) {
		return true
	}

	return false
}

// SetChannelCode gets a reference to the given EWalletChannelCode and assigns it to the ChannelCode field.
func (o *EWalletParameters) SetChannelCode(v EWalletChannelCode) {
	o.ChannelCode = &v
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

func (o EWalletParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EWalletParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ChannelCode) {
		toSerialize["channel_code"] = o.ChannelCode
	}
	if !utils.IsNil(o.ChannelProperties) {
		toSerialize["channel_properties"] = o.ChannelProperties
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


