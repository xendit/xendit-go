/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.7.6
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
)

// checks if the AlternativeDisplayItem type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AlternativeDisplayItem{}

// AlternativeDisplayItem An object representing alternative display of a VA.
type AlternativeDisplayItem struct {
	// Represent type of alternative display.
	Type *string `json:"type,omitempty"`
	// Represent value of alternative display value.
	Value *string `json:"value,omitempty"`
}

// NewAlternativeDisplayItem instantiates a new AlternativeDisplayItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlternativeDisplayItem() *AlternativeDisplayItem {
	this := AlternativeDisplayItem{}
	return &this
}

// NewAlternativeDisplayItemWithDefaults instantiates a new AlternativeDisplayItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlternativeDisplayItemWithDefaults() *AlternativeDisplayItem {
	this := AlternativeDisplayItem{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AlternativeDisplayItem) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeDisplayItem) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AlternativeDisplayItem) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AlternativeDisplayItem) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AlternativeDisplayItem) GetValue() string {
	if o == nil || utils.IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeDisplayItem) GetValueOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AlternativeDisplayItem) HasValue() bool {
	if o != nil && !utils.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AlternativeDisplayItem) SetValue(v string) {
	o.Value = &v
}

func (o AlternativeDisplayItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlternativeDisplayItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !utils.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableAlternativeDisplayItem struct {
	value *AlternativeDisplayItem
	isSet bool
}

func (v NullableAlternativeDisplayItem) Get() *AlternativeDisplayItem {
	return v.value
}

func (v *NullableAlternativeDisplayItem) Set(val *AlternativeDisplayItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAlternativeDisplayItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAlternativeDisplayItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlternativeDisplayItem(val *AlternativeDisplayItem) *NullableAlternativeDisplayItem {
	return &NullableAlternativeDisplayItem{value: val, isSet: true}
}

func (v NullableAlternativeDisplayItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlternativeDisplayItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


