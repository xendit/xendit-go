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

// checks if the DirectDebitParameters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DirectDebitParameters{}

// DirectDebitParameters struct for DirectDebitParameters
type DirectDebitParameters struct {
	ChannelCode DirectDebitChannelCode `json:"channel_code"`
	ChannelProperties NullableDirectDebitChannelProperties `json:"channel_properties"`
	Type *DirectDebitType `json:"type,omitempty"`
}

// NewDirectDebitParameters instantiates a new DirectDebitParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectDebitParameters(channelCode DirectDebitChannelCode, channelProperties NullableDirectDebitChannelProperties) *DirectDebitParameters {
	this := DirectDebitParameters{}
	this.ChannelCode = channelCode
	this.ChannelProperties = channelProperties
	return &this
}

// NewDirectDebitParametersWithDefaults instantiates a new DirectDebitParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectDebitParametersWithDefaults() *DirectDebitParameters {
	this := DirectDebitParameters{}
	return &this
}

// GetChannelCode returns the ChannelCode field value
func (o *DirectDebitParameters) GetChannelCode() DirectDebitChannelCode {
	if o == nil {
		var ret DirectDebitChannelCode
		return ret
	}

	return o.ChannelCode
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value
// and a boolean to check if the value has been set.
func (o *DirectDebitParameters) GetChannelCodeOk() (*DirectDebitChannelCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelCode, true
}

// SetChannelCode sets field value
func (o *DirectDebitParameters) SetChannelCode(v DirectDebitChannelCode) {
	o.ChannelCode = v
}

// GetChannelProperties returns the ChannelProperties field value
// If the value is explicit nil, the zero value for DirectDebitChannelProperties will be returned
func (o *DirectDebitParameters) GetChannelProperties() DirectDebitChannelProperties {
	if o == nil || o.ChannelProperties.Get() == nil {
		var ret DirectDebitChannelProperties
		return ret
	}

	return *o.ChannelProperties.Get()
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectDebitParameters) GetChannelPropertiesOk() (*DirectDebitChannelProperties, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelProperties.Get(), o.ChannelProperties.IsSet()
}

// SetChannelProperties sets field value
func (o *DirectDebitParameters) SetChannelProperties(v DirectDebitChannelProperties) {
	o.ChannelProperties.Set(&v)
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DirectDebitParameters) GetType() DirectDebitType {
	if o == nil || utils.IsNil(o.Type) {
		var ret DirectDebitType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDebitParameters) GetTypeOk() (*DirectDebitType, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DirectDebitParameters) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DirectDebitType and assigns it to the Type field.
func (o *DirectDebitParameters) SetType(v DirectDebitType) {
	o.Type = &v
}

func (o DirectDebitParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectDebitParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel_code"] = o.ChannelCode
	toSerialize["channel_properties"] = o.ChannelProperties.Get()

	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableDirectDebitParameters struct {
	value *DirectDebitParameters
	isSet bool
}

func (v NullableDirectDebitParameters) Get() *DirectDebitParameters {
	return v.value
}

func (v *NullableDirectDebitParameters) Set(val *DirectDebitParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectDebitParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectDebitParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectDebitParameters(val *DirectDebitParameters) *NullableDirectDebitParameters {
	return &NullableDirectDebitParameters{value: val, isSet: true}
}

func (v NullableDirectDebitParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectDebitParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


