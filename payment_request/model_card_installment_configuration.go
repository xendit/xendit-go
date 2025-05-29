/*
Payment Requests

This API is used for Payment Requests

API version: 1.70.0
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
)

// checks if the CardInstallmentConfiguration type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CardInstallmentConfiguration{}

// CardInstallmentConfiguration Card Installment Configuration
type CardInstallmentConfiguration struct {
	Terms *int32 `json:"terms,omitempty"`
	Interval *string `json:"interval,omitempty"`
	Code NullableString `json:"code,omitempty"`
}

// NewCardInstallmentConfiguration instantiates a new CardInstallmentConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardInstallmentConfiguration() *CardInstallmentConfiguration {
	this := CardInstallmentConfiguration{}
	return &this
}

// NewCardInstallmentConfigurationWithDefaults instantiates a new CardInstallmentConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardInstallmentConfigurationWithDefaults() *CardInstallmentConfiguration {
	this := CardInstallmentConfiguration{}
	return &this
}

// GetTerms returns the Terms field value if set, zero value otherwise.
func (o *CardInstallmentConfiguration) GetTerms() int32 {
	if o == nil || utils.IsNil(o.Terms) {
		var ret int32
		return ret
	}
	return *o.Terms
}

// GetTermsOk returns a tuple with the Terms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInstallmentConfiguration) GetTermsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Terms) {
		return nil, false
	}
	return o.Terms, true
}

// HasTerms returns a boolean if a field has been set.
func (o *CardInstallmentConfiguration) HasTerms() bool {
	if o != nil && !utils.IsNil(o.Terms) {
		return true
	}

	return false
}

// SetTerms gets a reference to the given int32 and assigns it to the Terms field.
func (o *CardInstallmentConfiguration) SetTerms(v int32) {
	o.Terms = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *CardInstallmentConfiguration) GetInterval() string {
	if o == nil || utils.IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInstallmentConfiguration) GetIntervalOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *CardInstallmentConfiguration) HasInterval() bool {
	if o != nil && !utils.IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *CardInstallmentConfiguration) SetInterval(v string) {
	o.Interval = &v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardInstallmentConfiguration) GetCode() string {
	if o == nil || utils.IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardInstallmentConfiguration) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *CardInstallmentConfiguration) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *CardInstallmentConfiguration) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *CardInstallmentConfiguration) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *CardInstallmentConfiguration) UnsetCode() {
	o.Code.Unset()
}

func (o CardInstallmentConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardInstallmentConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Terms) {
		toSerialize["terms"] = o.Terms
	}
	if !utils.IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
    }
	return toSerialize, nil
}

type NullableCardInstallmentConfiguration struct {
	value *CardInstallmentConfiguration
	isSet bool
}

func (v NullableCardInstallmentConfiguration) Get() *CardInstallmentConfiguration {
	return v.value
}

func (v *NullableCardInstallmentConfiguration) Set(val *CardInstallmentConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableCardInstallmentConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableCardInstallmentConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardInstallmentConfiguration(val *CardInstallmentConfiguration) *NullableCardInstallmentConfiguration {
	return &NullableCardInstallmentConfiguration{value: val, isSet: true}
}

func (v NullableCardInstallmentConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardInstallmentConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


