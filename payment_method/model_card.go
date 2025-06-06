/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.128.0
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
)

// checks if the Card type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Card{}

// Card Card Payment Method Details
type Card struct {
	ChannelCode *CardChannelCode `json:"channel_code,omitempty"`
	Currency NullableString `json:"currency,omitempty"`
	ChannelProperties NullableCardChannelProperties `json:"channel_properties"`
	CardInformation *TokenizedCardInformation `json:"card_information,omitempty"`
	CardVerificationResults NullableCardVerificationResults `json:"card_verification_results,omitempty"`
}

// NewCard instantiates a new Card object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCard(channelProperties NullableCardChannelProperties) *Card {
	this := Card{}
	this.ChannelProperties = channelProperties
	return &this
}

// NewCardWithDefaults instantiates a new Card object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardWithDefaults() *Card {
	this := Card{}
	return &this
}

// GetChannelCode returns the ChannelCode field value if set, zero value otherwise.
func (o *Card) GetChannelCode() CardChannelCode {
	if o == nil || utils.IsNil(o.ChannelCode) {
		var ret CardChannelCode
		return ret
	}
	return *o.ChannelCode
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetChannelCodeOk() (*CardChannelCode, bool) {
	if o == nil || utils.IsNil(o.ChannelCode) {
		return nil, false
	}
	return o.ChannelCode, true
}

// HasChannelCode returns a boolean if a field has been set.
func (o *Card) HasChannelCode() bool {
	if o != nil && !utils.IsNil(o.ChannelCode) {
		return true
	}

	return false
}

// SetChannelCode gets a reference to the given CardChannelCode and assigns it to the ChannelCode field.
func (o *Card) SetChannelCode(v CardChannelCode) {
	o.ChannelCode = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Card) GetCurrency() string {
	if o == nil || utils.IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Card) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *Card) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *Card) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *Card) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *Card) UnsetCurrency() {
	o.Currency.Unset()
}

// GetChannelProperties returns the ChannelProperties field value
// If the value is explicit nil, the zero value for CardChannelProperties will be returned
func (o *Card) GetChannelProperties() CardChannelProperties {
	if o == nil || o.ChannelProperties.Get() == nil {
		var ret CardChannelProperties
		return ret
	}

	return *o.ChannelProperties.Get()
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Card) GetChannelPropertiesOk() (*CardChannelProperties, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelProperties.Get(), o.ChannelProperties.IsSet()
}

// SetChannelProperties sets field value
func (o *Card) SetChannelProperties(v CardChannelProperties) {
	o.ChannelProperties.Set(&v)
}

// GetCardInformation returns the CardInformation field value if set, zero value otherwise.
func (o *Card) GetCardInformation() TokenizedCardInformation {
	if o == nil || utils.IsNil(o.CardInformation) {
		var ret TokenizedCardInformation
		return ret
	}
	return *o.CardInformation
}

// GetCardInformationOk returns a tuple with the CardInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetCardInformationOk() (*TokenizedCardInformation, bool) {
	if o == nil || utils.IsNil(o.CardInformation) {
		return nil, false
	}
	return o.CardInformation, true
}

// HasCardInformation returns a boolean if a field has been set.
func (o *Card) HasCardInformation() bool {
	if o != nil && !utils.IsNil(o.CardInformation) {
		return true
	}

	return false
}

// SetCardInformation gets a reference to the given TokenizedCardInformation and assigns it to the CardInformation field.
func (o *Card) SetCardInformation(v TokenizedCardInformation) {
	o.CardInformation = &v
}

// GetCardVerificationResults returns the CardVerificationResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Card) GetCardVerificationResults() CardVerificationResults {
	if o == nil || utils.IsNil(o.CardVerificationResults.Get()) {
		var ret CardVerificationResults
		return ret
	}
	return *o.CardVerificationResults.Get()
}

// GetCardVerificationResultsOk returns a tuple with the CardVerificationResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Card) GetCardVerificationResultsOk() (*CardVerificationResults, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardVerificationResults.Get(), o.CardVerificationResults.IsSet()
}

// HasCardVerificationResults returns a boolean if a field has been set.
func (o *Card) HasCardVerificationResults() bool {
	if o != nil && o.CardVerificationResults.IsSet() {
		return true
	}

	return false
}

// SetCardVerificationResults gets a reference to the given NullableCardVerificationResults and assigns it to the CardVerificationResults field.
func (o *Card) SetCardVerificationResults(v CardVerificationResults) {
	o.CardVerificationResults.Set(&v)
}
// SetCardVerificationResultsNil sets the value for CardVerificationResults to be an explicit nil
func (o *Card) SetCardVerificationResultsNil() {
	o.CardVerificationResults.Set(nil)
}

// UnsetCardVerificationResults ensures that no value is present for CardVerificationResults, not even an explicit nil
func (o *Card) UnsetCardVerificationResults() {
	o.CardVerificationResults.Unset()
}

func (o Card) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Card) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ChannelCode) {
		toSerialize["channel_code"] = o.ChannelCode
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
    }
	toSerialize["channel_properties"] = o.ChannelProperties.Get()

	if !utils.IsNil(o.CardInformation) {
		toSerialize["card_information"] = o.CardInformation
	}
	if o.CardVerificationResults.IsSet() {
		toSerialize["card_verification_results"] = o.CardVerificationResults.Get()
    }
	return toSerialize, nil
}

type NullableCard struct {
	value *Card
	isSet bool
}

func (v NullableCard) Get() *Card {
	return v.value
}

func (v *NullableCard) Set(val *Card) {
	v.value = val
	v.isSet = true
}

func (v NullableCard) IsSet() bool {
	return v.isSet
}

func (v *NullableCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCard(val *Card) *NullableCard {
	return &NullableCard{value: val, isSet: true}
}

func (v NullableCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


