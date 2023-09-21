/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.87.2
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the VirtualAccountUpdateParameters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &VirtualAccountUpdateParameters{}

// VirtualAccountUpdateParameters struct for VirtualAccountUpdateParameters
type VirtualAccountUpdateParameters struct {
	Amount NullableFloat64 `json:"amount,omitempty"`
	MinAmount NullableFloat64 `json:"min_amount,omitempty"`
	MaxAmount NullableFloat64 `json:"max_amount,omitempty"`
	ChannelProperties *VirtualAccountChannelPropertiesPatch `json:"channel_properties,omitempty"`
	// For payments in Vietnam only, alternative display requested for the virtual account
	AlternativeDisplayTypes []string `json:"alternative_display_types,omitempty"`
}

// NewVirtualAccountUpdateParameters instantiates a new VirtualAccountUpdateParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualAccountUpdateParameters() *VirtualAccountUpdateParameters {
	this := VirtualAccountUpdateParameters{}
	return &this
}

// NewVirtualAccountUpdateParametersWithDefaults instantiates a new VirtualAccountUpdateParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualAccountUpdateParametersWithDefaults() *VirtualAccountUpdateParameters {
	this := VirtualAccountUpdateParameters{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualAccountUpdateParameters) GetAmount() float64 {
	if o == nil || utils.IsNil(o.Amount.Get()) {
		var ret float64
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualAccountUpdateParameters) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *VirtualAccountUpdateParameters) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat64 and assigns it to the Amount field.
func (o *VirtualAccountUpdateParameters) SetAmount(v float64) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *VirtualAccountUpdateParameters) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *VirtualAccountUpdateParameters) UnsetAmount() {
	o.Amount.Unset()
}

// GetMinAmount returns the MinAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualAccountUpdateParameters) GetMinAmount() float64 {
	if o == nil || utils.IsNil(o.MinAmount.Get()) {
		var ret float64
		return ret
	}
	return *o.MinAmount.Get()
}

// GetMinAmountOk returns a tuple with the MinAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualAccountUpdateParameters) GetMinAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinAmount.Get(), o.MinAmount.IsSet()
}

// HasMinAmount returns a boolean if a field has been set.
func (o *VirtualAccountUpdateParameters) HasMinAmount() bool {
	if o != nil && o.MinAmount.IsSet() {
		return true
	}

	return false
}

// SetMinAmount gets a reference to the given NullableFloat64 and assigns it to the MinAmount field.
func (o *VirtualAccountUpdateParameters) SetMinAmount(v float64) {
	o.MinAmount.Set(&v)
}
// SetMinAmountNil sets the value for MinAmount to be an explicit nil
func (o *VirtualAccountUpdateParameters) SetMinAmountNil() {
	o.MinAmount.Set(nil)
}

// UnsetMinAmount ensures that no value is present for MinAmount, not even an explicit nil
func (o *VirtualAccountUpdateParameters) UnsetMinAmount() {
	o.MinAmount.Unset()
}

// GetMaxAmount returns the MaxAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualAccountUpdateParameters) GetMaxAmount() float64 {
	if o == nil || utils.IsNil(o.MaxAmount.Get()) {
		var ret float64
		return ret
	}
	return *o.MaxAmount.Get()
}

// GetMaxAmountOk returns a tuple with the MaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualAccountUpdateParameters) GetMaxAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxAmount.Get(), o.MaxAmount.IsSet()
}

// HasMaxAmount returns a boolean if a field has been set.
func (o *VirtualAccountUpdateParameters) HasMaxAmount() bool {
	if o != nil && o.MaxAmount.IsSet() {
		return true
	}

	return false
}

// SetMaxAmount gets a reference to the given NullableFloat64 and assigns it to the MaxAmount field.
func (o *VirtualAccountUpdateParameters) SetMaxAmount(v float64) {
	o.MaxAmount.Set(&v)
}
// SetMaxAmountNil sets the value for MaxAmount to be an explicit nil
func (o *VirtualAccountUpdateParameters) SetMaxAmountNil() {
	o.MaxAmount.Set(nil)
}

// UnsetMaxAmount ensures that no value is present for MaxAmount, not even an explicit nil
func (o *VirtualAccountUpdateParameters) UnsetMaxAmount() {
	o.MaxAmount.Unset()
}

// GetChannelProperties returns the ChannelProperties field value if set, zero value otherwise.
func (o *VirtualAccountUpdateParameters) GetChannelProperties() VirtualAccountChannelPropertiesPatch {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		var ret VirtualAccountChannelPropertiesPatch
		return ret
	}
	return *o.ChannelProperties
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualAccountUpdateParameters) GetChannelPropertiesOk() (*VirtualAccountChannelPropertiesPatch, bool) {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		return nil, false
	}
	return o.ChannelProperties, true
}

// HasChannelProperties returns a boolean if a field has been set.
func (o *VirtualAccountUpdateParameters) HasChannelProperties() bool {
	if o != nil && !utils.IsNil(o.ChannelProperties) {
		return true
	}

	return false
}

// SetChannelProperties gets a reference to the given VirtualAccountChannelPropertiesPatch and assigns it to the ChannelProperties field.
func (o *VirtualAccountUpdateParameters) SetChannelProperties(v VirtualAccountChannelPropertiesPatch) {
	o.ChannelProperties = &v
}

// GetAlternativeDisplayTypes returns the AlternativeDisplayTypes field value if set, zero value otherwise.
func (o *VirtualAccountUpdateParameters) GetAlternativeDisplayTypes() []string {
	if o == nil || utils.IsNil(o.AlternativeDisplayTypes) {
		var ret []string
		return ret
	}
	return o.AlternativeDisplayTypes
}

// GetAlternativeDisplayTypesOk returns a tuple with the AlternativeDisplayTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualAccountUpdateParameters) GetAlternativeDisplayTypesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.AlternativeDisplayTypes) {
		return nil, false
	}
	return o.AlternativeDisplayTypes, true
}

// HasAlternativeDisplayTypes returns a boolean if a field has been set.
func (o *VirtualAccountUpdateParameters) HasAlternativeDisplayTypes() bool {
	if o != nil && !utils.IsNil(o.AlternativeDisplayTypes) {
		return true
	}

	return false
}

// SetAlternativeDisplayTypes gets a reference to the given []string and assigns it to the AlternativeDisplayTypes field.
func (o *VirtualAccountUpdateParameters) SetAlternativeDisplayTypes(v []string) {
	o.AlternativeDisplayTypes = v
}

func (o VirtualAccountUpdateParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualAccountUpdateParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if o.MinAmount.IsSet() {
		toSerialize["min_amount"] = o.MinAmount.Get()
	}
	if o.MaxAmount.IsSet() {
		toSerialize["max_amount"] = o.MaxAmount.Get()
	}
	if !utils.IsNil(o.ChannelProperties) {
		toSerialize["channel_properties"] = o.ChannelProperties
	}
	if !utils.IsNil(o.AlternativeDisplayTypes) {
		toSerialize["alternative_display_types"] = o.AlternativeDisplayTypes
	}
	return toSerialize, nil
}

type NullableVirtualAccountUpdateParameters struct {
	value *VirtualAccountUpdateParameters
	isSet bool
}

func (v NullableVirtualAccountUpdateParameters) Get() *VirtualAccountUpdateParameters {
	return v.value
}

func (v *NullableVirtualAccountUpdateParameters) Set(val *VirtualAccountUpdateParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualAccountUpdateParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualAccountUpdateParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualAccountUpdateParameters(val *VirtualAccountUpdateParameters) *NullableVirtualAccountUpdateParameters {
	return &NullableVirtualAccountUpdateParameters{value: val, isSet: true}
}

func (v NullableVirtualAccountUpdateParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualAccountUpdateParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


