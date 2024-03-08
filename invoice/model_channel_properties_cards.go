/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.7.6
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v4/utils"
)

// checks if the ChannelPropertiesCards type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ChannelPropertiesCards{}

// ChannelPropertiesCards An object representing properties specific for credit card payment method.
type ChannelPropertiesCards struct {
	// An array of allowed BINs (6 or 8 digits) for credit card payments.
	AllowedBins []string `json:"allowed_bins,omitempty"`
}

// NewChannelPropertiesCards instantiates a new ChannelPropertiesCards object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPropertiesCards() *ChannelPropertiesCards {
	this := ChannelPropertiesCards{}
	return &this
}

// NewChannelPropertiesCardsWithDefaults instantiates a new ChannelPropertiesCards object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPropertiesCardsWithDefaults() *ChannelPropertiesCards {
	this := ChannelPropertiesCards{}
	return &this
}

// GetAllowedBins returns the AllowedBins field value if set, zero value otherwise.
func (o *ChannelPropertiesCards) GetAllowedBins() []string {
	if o == nil || utils.IsNil(o.AllowedBins) {
		var ret []string
		return ret
	}
	return o.AllowedBins
}

// GetAllowedBinsOk returns a tuple with the AllowedBins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPropertiesCards) GetAllowedBinsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.AllowedBins) {
		return nil, false
	}
	return o.AllowedBins, true
}

// HasAllowedBins returns a boolean if a field has been set.
func (o *ChannelPropertiesCards) HasAllowedBins() bool {
	if o != nil && !utils.IsNil(o.AllowedBins) {
		return true
	}

	return false
}

// SetAllowedBins gets a reference to the given []string and assigns it to the AllowedBins field.
func (o *ChannelPropertiesCards) SetAllowedBins(v []string) {
	o.AllowedBins = v
}

func (o ChannelPropertiesCards) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelPropertiesCards) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AllowedBins) {
		toSerialize["allowed_bins"] = o.AllowedBins
	}
	return toSerialize, nil
}

type NullableChannelPropertiesCards struct {
	value *ChannelPropertiesCards
	isSet bool
}

func (v NullableChannelPropertiesCards) Get() *ChannelPropertiesCards {
	return v.value
}

func (v *NullableChannelPropertiesCards) Set(val *ChannelPropertiesCards) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPropertiesCards) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPropertiesCards) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPropertiesCards(val *ChannelPropertiesCards) *NullableChannelPropertiesCards {
	return &NullableChannelPropertiesCards{value: val, isSet: true}
}

func (v NullableChannelPropertiesCards) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPropertiesCards) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


