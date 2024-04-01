/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.7.6
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v5/utils"
)

// checks if the ChannelProperties type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ChannelProperties{}

// ChannelProperties An object representing channel-specific properties.
type ChannelProperties struct {
	Cards *ChannelPropertiesCards `json:"cards,omitempty"`
}

// NewChannelProperties instantiates a new ChannelProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelProperties() *ChannelProperties {
	this := ChannelProperties{}
	return &this
}

// NewChannelPropertiesWithDefaults instantiates a new ChannelProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPropertiesWithDefaults() *ChannelProperties {
	this := ChannelProperties{}
	return &this
}

// GetCards returns the Cards field value if set, zero value otherwise.
func (o *ChannelProperties) GetCards() ChannelPropertiesCards {
	if o == nil || utils.IsNil(o.Cards) {
		var ret ChannelPropertiesCards
		return ret
	}
	return *o.Cards
}

// GetCardsOk returns a tuple with the Cards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelProperties) GetCardsOk() (*ChannelPropertiesCards, bool) {
	if o == nil || utils.IsNil(o.Cards) {
		return nil, false
	}
	return o.Cards, true
}

// HasCards returns a boolean if a field has been set.
func (o *ChannelProperties) HasCards() bool {
	if o != nil && !utils.IsNil(o.Cards) {
		return true
	}

	return false
}

// SetCards gets a reference to the given ChannelPropertiesCards and assigns it to the Cards field.
func (o *ChannelProperties) SetCards(v ChannelPropertiesCards) {
	o.Cards = &v
}

func (o ChannelProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Cards) {
		toSerialize["cards"] = o.Cards
	}
	return toSerialize, nil
}

type NullableChannelProperties struct {
	value *ChannelProperties
	isSet bool
}

func (v NullableChannelProperties) Get() *ChannelProperties {
	return v.value
}

func (v *NullableChannelProperties) Set(val *ChannelProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelProperties(val *ChannelProperties) *NullableChannelProperties {
	return &NullableChannelProperties{value: val, isSet: true}
}

func (v NullableChannelProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


