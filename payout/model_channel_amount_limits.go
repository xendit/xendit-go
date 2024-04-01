/*
Payout Service

This API allows Xendit to send money from an account to a channel (banks, eWallets, retail outlets) from across regions

API version: 1.0.0
*/


package payout

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v5/utils"
)

// checks if the ChannelAmountLimits type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ChannelAmountLimits{}

// ChannelAmountLimits Supported amount ranges for payouts to this channel
type ChannelAmountLimits struct {
	// Lowest amount supported for a payout to this channel
	Minimum float32 `json:"minimum"`
	// Highest amount supported for a payout to this channel
	Maximum float32 `json:"maximum"`
	// Supported increments
	MinimumIncrement float32 `json:"minimum_increment"`
}

// NewChannelAmountLimits instantiates a new ChannelAmountLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelAmountLimits(minimum float32, maximum float32, minimumIncrement float32) *ChannelAmountLimits {
	this := ChannelAmountLimits{}
	this.Minimum = minimum
	this.Maximum = maximum
	this.MinimumIncrement = minimumIncrement
	return &this
}

// NewChannelAmountLimitsWithDefaults instantiates a new ChannelAmountLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelAmountLimitsWithDefaults() *ChannelAmountLimits {
	this := ChannelAmountLimits{}
	return &this
}

// GetMinimum returns the Minimum field value
func (o *ChannelAmountLimits) GetMinimum() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value
// and a boolean to check if the value has been set.
func (o *ChannelAmountLimits) GetMinimumOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minimum, true
}

// SetMinimum sets field value
func (o *ChannelAmountLimits) SetMinimum(v float32) {
	o.Minimum = v
}

// GetMaximum returns the Maximum field value
func (o *ChannelAmountLimits) GetMaximum() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Maximum
}

// GetMaximumOk returns a tuple with the Maximum field value
// and a boolean to check if the value has been set.
func (o *ChannelAmountLimits) GetMaximumOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Maximum, true
}

// SetMaximum sets field value
func (o *ChannelAmountLimits) SetMaximum(v float32) {
	o.Maximum = v
}

// GetMinimumIncrement returns the MinimumIncrement field value
func (o *ChannelAmountLimits) GetMinimumIncrement() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MinimumIncrement
}

// GetMinimumIncrementOk returns a tuple with the MinimumIncrement field value
// and a boolean to check if the value has been set.
func (o *ChannelAmountLimits) GetMinimumIncrementOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinimumIncrement, true
}

// SetMinimumIncrement sets field value
func (o *ChannelAmountLimits) SetMinimumIncrement(v float32) {
	o.MinimumIncrement = v
}

func (o ChannelAmountLimits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelAmountLimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["minimum"] = o.Minimum
	toSerialize["maximum"] = o.Maximum
	toSerialize["minimum_increment"] = o.MinimumIncrement
	return toSerialize, nil
}

type NullableChannelAmountLimits struct {
	value *ChannelAmountLimits
	isSet bool
}

func (v NullableChannelAmountLimits) Get() *ChannelAmountLimits {
	return v.value
}

func (v *NullableChannelAmountLimits) Set(val *ChannelAmountLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelAmountLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelAmountLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelAmountLimits(val *ChannelAmountLimits) *NullableChannelAmountLimits {
	return &NullableChannelAmountLimits{value: val, isSet: true}
}

func (v NullableChannelAmountLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelAmountLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


