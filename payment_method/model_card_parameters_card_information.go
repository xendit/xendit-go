/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.89.2
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the CardParametersCardInformation type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CardParametersCardInformation{}

// CardParametersCardInformation Card Information
type CardParametersCardInformation struct {
	CardNumber string `json:"card_number"`
	// Card expiry month in MM format
	ExpiryMonth string `json:"expiry_month"`
	// Card expiry month in YY format
	ExpiryYear string `json:"expiry_year"`
	// Cardholder name
	CardholderName NullableString `json:"cardholder_name,omitempty"`
	Cvv NullableString `json:"cvv,omitempty"`
}

// NewCardParametersCardInformation instantiates a new CardParametersCardInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardParametersCardInformation(cardNumber string, expiryMonth string, expiryYear string) *CardParametersCardInformation {
	this := CardParametersCardInformation{}
	this.CardNumber = cardNumber
	this.ExpiryMonth = expiryMonth
	this.ExpiryYear = expiryYear
	return &this
}

// NewCardParametersCardInformationWithDefaults instantiates a new CardParametersCardInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardParametersCardInformationWithDefaults() *CardParametersCardInformation {
	this := CardParametersCardInformation{}
	return &this
}

// GetCardNumber returns the CardNumber field value
func (o *CardParametersCardInformation) GetCardNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardNumber
}

// GetCardNumberOk returns a tuple with the CardNumber field value
// and a boolean to check if the value has been set.
func (o *CardParametersCardInformation) GetCardNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardNumber, true
}

// SetCardNumber sets field value
func (o *CardParametersCardInformation) SetCardNumber(v string) {
	o.CardNumber = v
}

// GetExpiryMonth returns the ExpiryMonth field value
func (o *CardParametersCardInformation) GetExpiryMonth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpiryMonth
}

// GetExpiryMonthOk returns a tuple with the ExpiryMonth field value
// and a boolean to check if the value has been set.
func (o *CardParametersCardInformation) GetExpiryMonthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiryMonth, true
}

// SetExpiryMonth sets field value
func (o *CardParametersCardInformation) SetExpiryMonth(v string) {
	o.ExpiryMonth = v
}

// GetExpiryYear returns the ExpiryYear field value
func (o *CardParametersCardInformation) GetExpiryYear() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpiryYear
}

// GetExpiryYearOk returns a tuple with the ExpiryYear field value
// and a boolean to check if the value has been set.
func (o *CardParametersCardInformation) GetExpiryYearOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiryYear, true
}

// SetExpiryYear sets field value
func (o *CardParametersCardInformation) SetExpiryYear(v string) {
	o.ExpiryYear = v
}

// GetCardholderName returns the CardholderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardParametersCardInformation) GetCardholderName() string {
	if o == nil || utils.IsNil(o.CardholderName.Get()) {
		var ret string
		return ret
	}
	return *o.CardholderName.Get()
}

// GetCardholderNameOk returns a tuple with the CardholderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardParametersCardInformation) GetCardholderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardholderName.Get(), o.CardholderName.IsSet()
}

// HasCardholderName returns a boolean if a field has been set.
func (o *CardParametersCardInformation) HasCardholderName() bool {
	if o != nil && o.CardholderName.IsSet() {
		return true
	}

	return false
}

// SetCardholderName gets a reference to the given NullableString and assigns it to the CardholderName field.
func (o *CardParametersCardInformation) SetCardholderName(v string) {
	o.CardholderName.Set(&v)
}
// SetCardholderNameNil sets the value for CardholderName to be an explicit nil
func (o *CardParametersCardInformation) SetCardholderNameNil() {
	o.CardholderName.Set(nil)
}

// UnsetCardholderName ensures that no value is present for CardholderName, not even an explicit nil
func (o *CardParametersCardInformation) UnsetCardholderName() {
	o.CardholderName.Unset()
}

// GetCvv returns the Cvv field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardParametersCardInformation) GetCvv() string {
	if o == nil || utils.IsNil(o.Cvv.Get()) {
		var ret string
		return ret
	}
	return *o.Cvv.Get()
}

// GetCvvOk returns a tuple with the Cvv field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardParametersCardInformation) GetCvvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cvv.Get(), o.Cvv.IsSet()
}

// HasCvv returns a boolean if a field has been set.
func (o *CardParametersCardInformation) HasCvv() bool {
	if o != nil && o.Cvv.IsSet() {
		return true
	}

	return false
}

// SetCvv gets a reference to the given NullableString and assigns it to the Cvv field.
func (o *CardParametersCardInformation) SetCvv(v string) {
	o.Cvv.Set(&v)
}
// SetCvvNil sets the value for Cvv to be an explicit nil
func (o *CardParametersCardInformation) SetCvvNil() {
	o.Cvv.Set(nil)
}

// UnsetCvv ensures that no value is present for Cvv, not even an explicit nil
func (o *CardParametersCardInformation) UnsetCvv() {
	o.Cvv.Unset()
}

func (o CardParametersCardInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardParametersCardInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["card_number"] = o.CardNumber
	toSerialize["expiry_month"] = o.ExpiryMonth
	toSerialize["expiry_year"] = o.ExpiryYear
	if o.CardholderName.IsSet() {
		toSerialize["cardholder_name"] = o.CardholderName.Get()
    }
	if o.Cvv.IsSet() {
		toSerialize["cvv"] = o.Cvv.Get()
    }
	return toSerialize, nil
}

type NullableCardParametersCardInformation struct {
	value *CardParametersCardInformation
	isSet bool
}

func (v NullableCardParametersCardInformation) Get() *CardParametersCardInformation {
	return v.value
}

func (v *NullableCardParametersCardInformation) Set(val *CardParametersCardInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableCardParametersCardInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableCardParametersCardInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardParametersCardInformation(val *CardParametersCardInformation) *NullableCardParametersCardInformation {
	return &NullableCardParametersCardInformation{value: val, isSet: true}
}

func (v NullableCardParametersCardInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardParametersCardInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


