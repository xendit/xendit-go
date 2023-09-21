/*
Payment Requests

This API is used for Payment Requests

API version: 1.44.0
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the QRCode type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &QRCode{}

// QRCode QRCode Payment Method Details
type QRCode struct {
	ChannelCode NullableQRCodeChannelCode `json:"channel_code,omitempty"`
	ChannelProperties *QRCodeChannelProperties `json:"channel_properties,omitempty"`
}

// NewQRCode instantiates a new QRCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQRCode() *QRCode {
	this := QRCode{}
	return &this
}

// NewQRCodeWithDefaults instantiates a new QRCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQRCodeWithDefaults() *QRCode {
	this := QRCode{}
	return &this
}

// GetChannelCode returns the ChannelCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QRCode) GetChannelCode() QRCodeChannelCode {
	if o == nil || utils.IsNil(o.ChannelCode.Get()) {
		var ret QRCodeChannelCode
		return ret
	}
	return *o.ChannelCode.Get()
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QRCode) GetChannelCodeOk() (*QRCodeChannelCode, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelCode.Get(), o.ChannelCode.IsSet()
}

// HasChannelCode returns a boolean if a field has been set.
func (o *QRCode) HasChannelCode() bool {
	if o != nil && o.ChannelCode.IsSet() {
		return true
	}

	return false
}

// SetChannelCode gets a reference to the given NullableQRCodeChannelCode and assigns it to the ChannelCode field.
func (o *QRCode) SetChannelCode(v QRCodeChannelCode) {
	o.ChannelCode.Set(&v)
}
// SetChannelCodeNil sets the value for ChannelCode to be an explicit nil
func (o *QRCode) SetChannelCodeNil() {
	o.ChannelCode.Set(nil)
}

// UnsetChannelCode ensures that no value is present for ChannelCode, not even an explicit nil
func (o *QRCode) UnsetChannelCode() {
	o.ChannelCode.Unset()
}

// GetChannelProperties returns the ChannelProperties field value if set, zero value otherwise.
func (o *QRCode) GetChannelProperties() QRCodeChannelProperties {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		var ret QRCodeChannelProperties
		return ret
	}
	return *o.ChannelProperties
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QRCode) GetChannelPropertiesOk() (*QRCodeChannelProperties, bool) {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		return nil, false
	}
	return o.ChannelProperties, true
}

// HasChannelProperties returns a boolean if a field has been set.
func (o *QRCode) HasChannelProperties() bool {
	if o != nil && !utils.IsNil(o.ChannelProperties) {
		return true
	}

	return false
}

// SetChannelProperties gets a reference to the given QRCodeChannelProperties and assigns it to the ChannelProperties field.
func (o *QRCode) SetChannelProperties(v QRCodeChannelProperties) {
	o.ChannelProperties = &v
}

func (o QRCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QRCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ChannelCode.IsSet() {
		toSerialize["channel_code"] = o.ChannelCode.Get()
	}
	if !utils.IsNil(o.ChannelProperties) {
		toSerialize["channel_properties"] = o.ChannelProperties
	}
	return toSerialize, nil
}

type NullableQRCode struct {
	value *QRCode
	isSet bool
}

func (v NullableQRCode) Get() *QRCode {
	return v.value
}

func (v *NullableQRCode) Set(val *QRCode) {
	v.value = val
	v.isSet = true
}

func (v NullableQRCode) IsSet() bool {
	return v.isSet
}

func (v *NullableQRCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQRCode(val *QRCode) *NullableQRCode {
	return &NullableQRCode{value: val, isSet: true}
}

func (v NullableQRCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQRCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


