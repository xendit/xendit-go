/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.99.0
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v5/utils"
)

// checks if the VirtualAccountAlternativeDisplay type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &VirtualAccountAlternativeDisplay{}

// VirtualAccountAlternativeDisplay Alternative Display Object
type VirtualAccountAlternativeDisplay struct {
	// Type of the alternative display
	Type *string `json:"type,omitempty"`
	// Data payload of the given alternative display
	Data *string `json:"data,omitempty"`
}

// NewVirtualAccountAlternativeDisplay instantiates a new VirtualAccountAlternativeDisplay object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualAccountAlternativeDisplay() *VirtualAccountAlternativeDisplay {
	this := VirtualAccountAlternativeDisplay{}
	return &this
}

// NewVirtualAccountAlternativeDisplayWithDefaults instantiates a new VirtualAccountAlternativeDisplay object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualAccountAlternativeDisplayWithDefaults() *VirtualAccountAlternativeDisplay {
	this := VirtualAccountAlternativeDisplay{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VirtualAccountAlternativeDisplay) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualAccountAlternativeDisplay) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VirtualAccountAlternativeDisplay) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VirtualAccountAlternativeDisplay) SetType(v string) {
	o.Type = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *VirtualAccountAlternativeDisplay) GetData() string {
	if o == nil || utils.IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualAccountAlternativeDisplay) GetDataOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *VirtualAccountAlternativeDisplay) HasData() bool {
	if o != nil && !utils.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *VirtualAccountAlternativeDisplay) SetData(v string) {
	o.Data = &v
}

func (o VirtualAccountAlternativeDisplay) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualAccountAlternativeDisplay) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !utils.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableVirtualAccountAlternativeDisplay struct {
	value *VirtualAccountAlternativeDisplay
	isSet bool
}

func (v NullableVirtualAccountAlternativeDisplay) Get() *VirtualAccountAlternativeDisplay {
	return v.value
}

func (v *NullableVirtualAccountAlternativeDisplay) Set(val *VirtualAccountAlternativeDisplay) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualAccountAlternativeDisplay) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualAccountAlternativeDisplay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualAccountAlternativeDisplay(val *VirtualAccountAlternativeDisplay) *NullableVirtualAccountAlternativeDisplay {
	return &NullableVirtualAccountAlternativeDisplay{value: val, isSet: true}
}

func (v NullableVirtualAccountAlternativeDisplay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualAccountAlternativeDisplay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


