/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.6.0
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v4/utils"
)

// checks if the Paylater type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Paylater{}

// Paylater An object representing paylater details for invoices.
type Paylater struct {
	PaylaterType PaylaterType `json:"paylater_type"`
	// Indicates whether this paylater option should be excluded.
	ShouldExclude *bool `json:"should_exclude,omitempty"`
}

// NewPaylater instantiates a new Paylater object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaylater(paylaterType PaylaterType) *Paylater {
	this := Paylater{}
	this.PaylaterType = paylaterType
	return &this
}

// NewPaylaterWithDefaults instantiates a new Paylater object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaylaterWithDefaults() *Paylater {
	this := Paylater{}
	return &this
}

// GetPaylaterType returns the PaylaterType field value
func (o *Paylater) GetPaylaterType() PaylaterType {
	if o == nil {
		var ret PaylaterType
		return ret
	}

	return o.PaylaterType
}

// GetPaylaterTypeOk returns a tuple with the PaylaterType field value
// and a boolean to check if the value has been set.
func (o *Paylater) GetPaylaterTypeOk() (*PaylaterType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaylaterType, true
}

// SetPaylaterType sets field value
func (o *Paylater) SetPaylaterType(v PaylaterType) {
	o.PaylaterType = v
}

// GetShouldExclude returns the ShouldExclude field value if set, zero value otherwise.
func (o *Paylater) GetShouldExclude() bool {
	if o == nil || utils.IsNil(o.ShouldExclude) {
		var ret bool
		return ret
	}
	return *o.ShouldExclude
}

// GetShouldExcludeOk returns a tuple with the ShouldExclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Paylater) GetShouldExcludeOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.ShouldExclude) {
		return nil, false
	}
	return o.ShouldExclude, true
}

// HasShouldExclude returns a boolean if a field has been set.
func (o *Paylater) HasShouldExclude() bool {
	if o != nil && !utils.IsNil(o.ShouldExclude) {
		return true
	}

	return false
}

// SetShouldExclude gets a reference to the given bool and assigns it to the ShouldExclude field.
func (o *Paylater) SetShouldExclude(v bool) {
	o.ShouldExclude = &v
}

func (o Paylater) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Paylater) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["paylater_type"] = o.PaylaterType
	if !utils.IsNil(o.ShouldExclude) {
		toSerialize["should_exclude"] = o.ShouldExclude
	}
	return toSerialize, nil
}

type NullablePaylater struct {
	value *Paylater
	isSet bool
}

func (v NullablePaylater) Get() *Paylater {
	return v.value
}

func (v *NullablePaylater) Set(val *Paylater) {
	v.value = val
	v.isSet = true
}

func (v NullablePaylater) IsSet() bool {
	return v.isSet
}

func (v *NullablePaylater) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaylater(val *Paylater) *NullablePaylater {
	return &NullablePaylater{value: val, isSet: true}
}

func (v NullablePaylater) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaylater) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


