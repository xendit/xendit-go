/*
Payment Requests

This API is used for Payment Requests

API version: 1.45.2
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the Card type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Card{}

// Card struct for Card
type Card struct {
	Currency PaymentRequestCurrency `json:"currency"`
	ChannelProperties CardChannelProperties `json:"channel_properties"`
	CardInformation CardInformation `json:"card_information"`
	CardVerificationResults NullableCardVerificationResults `json:"card_verification_results,omitempty"`
}

// NewCard instantiates a new Card object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCard(currency PaymentRequestCurrency, channelProperties CardChannelProperties, cardInformation CardInformation) *Card {
	this := Card{}
	this.Currency = currency
	this.ChannelProperties = channelProperties
	this.CardInformation = cardInformation
	return &this
}

// NewCardWithDefaults instantiates a new Card object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardWithDefaults() *Card {
	this := Card{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *Card) GetCurrency() PaymentRequestCurrency {
	if o == nil {
		var ret PaymentRequestCurrency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Card) GetCurrencyOk() (*PaymentRequestCurrency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Card) SetCurrency(v PaymentRequestCurrency) {
	o.Currency = v
}

// GetChannelProperties returns the ChannelProperties field value
func (o *Card) GetChannelProperties() CardChannelProperties {
	if o == nil {
		var ret CardChannelProperties
		return ret
	}

	return o.ChannelProperties
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value
// and a boolean to check if the value has been set.
func (o *Card) GetChannelPropertiesOk() (*CardChannelProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelProperties, true
}

// SetChannelProperties sets field value
func (o *Card) SetChannelProperties(v CardChannelProperties) {
	o.ChannelProperties = v
}

// GetCardInformation returns the CardInformation field value
func (o *Card) GetCardInformation() CardInformation {
	if o == nil {
		var ret CardInformation
		return ret
	}

	return o.CardInformation
}

// GetCardInformationOk returns a tuple with the CardInformation field value
// and a boolean to check if the value has been set.
func (o *Card) GetCardInformationOk() (*CardInformation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardInformation, true
}

// SetCardInformation sets field value
func (o *Card) SetCardInformation(v CardInformation) {
	o.CardInformation = v
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
	toSerialize["currency"] = o.Currency
	toSerialize["channel_properties"] = o.ChannelProperties
	toSerialize["card_information"] = o.CardInformation
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


