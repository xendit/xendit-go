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

// checks if the CardParameters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CardParameters{}

// CardParameters struct for CardParameters
type CardParameters struct {
	Currency string `json:"currency"`
	ChannelProperties NullableCardChannelProperties `json:"channel_properties,omitempty"`
	CardInformation *CardParametersCardInformation `json:"card_information,omitempty"`
}

// NewCardParameters instantiates a new CardParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardParameters(currency string) *CardParameters {
	this := CardParameters{}
	this.Currency = currency
	return &this
}

// NewCardParametersWithDefaults instantiates a new CardParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardParametersWithDefaults() *CardParameters {
	this := CardParameters{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *CardParameters) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CardParameters) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CardParameters) SetCurrency(v string) {
	o.Currency = v
}

// GetChannelProperties returns the ChannelProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardParameters) GetChannelProperties() CardChannelProperties {
	if o == nil || utils.IsNil(o.ChannelProperties.Get()) {
		var ret CardChannelProperties
		return ret
	}
	return *o.ChannelProperties.Get()
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardParameters) GetChannelPropertiesOk() (*CardChannelProperties, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelProperties.Get(), o.ChannelProperties.IsSet()
}

// HasChannelProperties returns a boolean if a field has been set.
func (o *CardParameters) HasChannelProperties() bool {
	if o != nil && o.ChannelProperties.IsSet() {
		return true
	}

	return false
}

// SetChannelProperties gets a reference to the given NullableCardChannelProperties and assigns it to the ChannelProperties field.
func (o *CardParameters) SetChannelProperties(v CardChannelProperties) {
	o.ChannelProperties.Set(&v)
}
// SetChannelPropertiesNil sets the value for ChannelProperties to be an explicit nil
func (o *CardParameters) SetChannelPropertiesNil() {
	o.ChannelProperties.Set(nil)
}

// UnsetChannelProperties ensures that no value is present for ChannelProperties, not even an explicit nil
func (o *CardParameters) UnsetChannelProperties() {
	o.ChannelProperties.Unset()
}

// GetCardInformation returns the CardInformation field value if set, zero value otherwise.
func (o *CardParameters) GetCardInformation() CardParametersCardInformation {
	if o == nil || utils.IsNil(o.CardInformation) {
		var ret CardParametersCardInformation
		return ret
	}
	return *o.CardInformation
}

// GetCardInformationOk returns a tuple with the CardInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardParameters) GetCardInformationOk() (*CardParametersCardInformation, bool) {
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

// SetCardInformation gets a reference to the given CardParametersCardInformation and assigns it to the CardInformation field.
func (o *CardParameters) SetCardInformation(v CardParametersCardInformation) {
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
	toSerialize["currency"] = o.Currency
	if o.ChannelProperties.IsSet() {
		toSerialize["channel_properties"] = o.ChannelProperties.Get()
	}
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


