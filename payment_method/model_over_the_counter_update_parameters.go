/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.91.2
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v4/utils"
)

// checks if the OverTheCounterUpdateParameters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &OverTheCounterUpdateParameters{}

// OverTheCounterUpdateParameters struct for OverTheCounterUpdateParameters
type OverTheCounterUpdateParameters struct {
	Amount NullableFloat64 `json:"amount,omitempty"`
	ChannelProperties *OverTheCounterChannelPropertiesUpdate `json:"channel_properties,omitempty"`
}

// NewOverTheCounterUpdateParameters instantiates a new OverTheCounterUpdateParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverTheCounterUpdateParameters() *OverTheCounterUpdateParameters {
	this := OverTheCounterUpdateParameters{}
	return &this
}

// NewOverTheCounterUpdateParametersWithDefaults instantiates a new OverTheCounterUpdateParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverTheCounterUpdateParametersWithDefaults() *OverTheCounterUpdateParameters {
	this := OverTheCounterUpdateParameters{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OverTheCounterUpdateParameters) GetAmount() float64 {
	if o == nil || utils.IsNil(o.Amount.Get()) {
		var ret float64
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OverTheCounterUpdateParameters) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *OverTheCounterUpdateParameters) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat64 and assigns it to the Amount field.
func (o *OverTheCounterUpdateParameters) SetAmount(v float64) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *OverTheCounterUpdateParameters) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *OverTheCounterUpdateParameters) UnsetAmount() {
	o.Amount.Unset()
}

// GetChannelProperties returns the ChannelProperties field value if set, zero value otherwise.
func (o *OverTheCounterUpdateParameters) GetChannelProperties() OverTheCounterChannelPropertiesUpdate {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		var ret OverTheCounterChannelPropertiesUpdate
		return ret
	}
	return *o.ChannelProperties
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverTheCounterUpdateParameters) GetChannelPropertiesOk() (*OverTheCounterChannelPropertiesUpdate, bool) {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		return nil, false
	}
	return o.ChannelProperties, true
}

// HasChannelProperties returns a boolean if a field has been set.
func (o *OverTheCounterUpdateParameters) HasChannelProperties() bool {
	if o != nil && !utils.IsNil(o.ChannelProperties) {
		return true
	}

	return false
}

// SetChannelProperties gets a reference to the given OverTheCounterChannelPropertiesUpdate and assigns it to the ChannelProperties field.
func (o *OverTheCounterUpdateParameters) SetChannelProperties(v OverTheCounterChannelPropertiesUpdate) {
	o.ChannelProperties = &v
}

func (o OverTheCounterUpdateParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OverTheCounterUpdateParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
    }
	if !utils.IsNil(o.ChannelProperties) {
		toSerialize["channel_properties"] = o.ChannelProperties
	}
	return toSerialize, nil
}

type NullableOverTheCounterUpdateParameters struct {
	value *OverTheCounterUpdateParameters
	isSet bool
}

func (v NullableOverTheCounterUpdateParameters) Get() *OverTheCounterUpdateParameters {
	return v.value
}

func (v *NullableOverTheCounterUpdateParameters) Set(val *OverTheCounterUpdateParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableOverTheCounterUpdateParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableOverTheCounterUpdateParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverTheCounterUpdateParameters(val *OverTheCounterUpdateParameters) *NullableOverTheCounterUpdateParameters {
	return &NullableOverTheCounterUpdateParameters{value: val, isSet: true}
}

func (v NullableOverTheCounterUpdateParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverTheCounterUpdateParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


