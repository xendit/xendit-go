/*
Payment Requests

This API is used for Payment Requests

API version: 1.70.0
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
)

// checks if the CardParameters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CardParameters{}

// CardParameters struct for CardParameters
type CardParameters struct {
	ChannelProperties CardChannelProperties `json:"channel_properties"`
	CardInformation *CardInformation `json:"card_information,omitempty"`
}

// NewCardParameters instantiates a new CardParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardParameters(channelProperties CardChannelProperties) *CardParameters {
	this := CardParameters{}
	this.ChannelProperties = channelProperties
	return &this
}

// NewCardParametersWithDefaults instantiates a new CardParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardParametersWithDefaults() *CardParameters {
	this := CardParameters{}
	return &this
}

// GetChannelProperties returns the ChannelProperties field value
func (o *CardParameters) GetChannelProperties() CardChannelProperties {
	if o == nil {
		var ret CardChannelProperties
		return ret
	}

	return o.ChannelProperties
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value
// and a boolean to check if the value has been set.
func (o *CardParameters) GetChannelPropertiesOk() (*CardChannelProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelProperties, true
}

// SetChannelProperties sets field value
func (o *CardParameters) SetChannelProperties(v CardChannelProperties) {
	o.ChannelProperties = v
}

// GetCardInformation returns the CardInformation field value if set, zero value otherwise.
func (o *CardParameters) GetCardInformation() CardInformation {
	if o == nil || utils.IsNil(o.CardInformation) {
		var ret CardInformation
		return ret
	}
	return *o.CardInformation
}

// GetCardInformationOk returns a tuple with the CardInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardParameters) GetCardInformationOk() (*CardInformation, bool) {
	if o == nil || utils.IsNil(o.CardInformation) {
		return nil, false
	}
	return o.CardInformation, true
}

// HasCardInformation returns a boolean if a field has been set.
func (o *CardParameters) HasCardInformation() bool {
	if o != nil && !utils.IsNil(o.CardInformation) {
		return true
	}

	return false
}

// SetCardInformation gets a reference to the given CardInformation and assigns it to the CardInformation field.
func (o *CardParameters) SetCardInformation(v CardInformation) {
	o.CardInformation = &v
}

func (o CardParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel_properties"] = o.ChannelProperties
	if !utils.IsNil(o.CardInformation) {
		toSerialize["card_information"] = o.CardInformation
	}
	return toSerialize, nil
}

type NullableCardParameters struct {
	value *CardParameters
	isSet bool
}

func (v NullableCardParameters) Get() *CardParameters {
	return v.value
}

func (v *NullableCardParameters) Set(val *CardParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableCardParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableCardParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardParameters(val *CardParameters) *NullableCardParameters {
	return &NullableCardParameters{value: val, isSet: true}
}

func (v NullableCardParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


