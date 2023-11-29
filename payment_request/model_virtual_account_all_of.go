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

// checks if the VirtualAccountAllOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &VirtualAccountAllOf{}

// VirtualAccountAllOf struct for VirtualAccountAllOf
type VirtualAccountAllOf struct {
	AlternativeDisplays []VirtualAccountAlternativeDisplay `json:"alternative_displays,omitempty"`
}

// NewVirtualAccountAllOf instantiates a new VirtualAccountAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualAccountAllOf() *VirtualAccountAllOf {
	this := VirtualAccountAllOf{}
	return &this
}

// NewVirtualAccountAllOfWithDefaults instantiates a new VirtualAccountAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualAccountAllOfWithDefaults() *VirtualAccountAllOf {
	this := VirtualAccountAllOf{}
	return &this
}

// GetAlternativeDisplays returns the AlternativeDisplays field value if set, zero value otherwise.
func (o *VirtualAccountAllOf) GetAlternativeDisplays() []VirtualAccountAlternativeDisplay {
	if o == nil || utils.IsNil(o.AlternativeDisplays) {
		var ret []VirtualAccountAlternativeDisplay
		return ret
	}
	return o.AlternativeDisplays
}

// GetAlternativeDisplaysOk returns a tuple with the AlternativeDisplays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualAccountAllOf) GetAlternativeDisplaysOk() ([]VirtualAccountAlternativeDisplay, bool) {
	if o == nil || utils.IsNil(o.AlternativeDisplays) {
		return nil, false
	}
	return o.AlternativeDisplays, true
}

// HasAlternativeDisplays returns a boolean if a field has been set.
func (o *VirtualAccountAllOf) HasAlternativeDisplays() bool {
	if o != nil && !utils.IsNil(o.AlternativeDisplays) {
		return true
	}

	return false
}

// SetAlternativeDisplays gets a reference to the given []VirtualAccountAlternativeDisplay and assigns it to the AlternativeDisplays field.
func (o *VirtualAccountAllOf) SetAlternativeDisplays(v []VirtualAccountAlternativeDisplay) {
	o.AlternativeDisplays = v
}

func (o VirtualAccountAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualAccountAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AlternativeDisplays) {
		toSerialize["alternative_displays"] = o.AlternativeDisplays
	}
	return toSerialize, nil
}

type NullableVirtualAccountAllOf struct {
	value *VirtualAccountAllOf
	isSet bool
}

func (v NullableVirtualAccountAllOf) Get() *VirtualAccountAllOf {
	return v.value
}

func (v *NullableVirtualAccountAllOf) Set(val *VirtualAccountAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualAccountAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualAccountAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualAccountAllOf(val *VirtualAccountAllOf) *NullableVirtualAccountAllOf {
	return &NullableVirtualAccountAllOf{value: val, isSet: true}
}

func (v NullableVirtualAccountAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualAccountAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


