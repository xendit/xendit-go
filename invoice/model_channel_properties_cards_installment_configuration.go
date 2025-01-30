/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.8.7
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
)

// checks if the ChannelPropertiesCardsInstallmentConfiguration type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ChannelPropertiesCardsInstallmentConfiguration{}

// ChannelPropertiesCardsInstallmentConfiguration An object to pre-set cards installment for a specific invoice
type ChannelPropertiesCardsInstallmentConfiguration struct {
	// Indicate whether full payment (without installment) is allowed
	AllowFullPayment *bool `json:"allow_full_payment,omitempty"`
	// An object to set what kind (from specific bank / specific tenor) of cards installments will be available on a specific invoice
	AllowedTerms []ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner `json:"allowed_terms,omitempty"`
}

// NewChannelPropertiesCardsInstallmentConfiguration instantiates a new ChannelPropertiesCardsInstallmentConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPropertiesCardsInstallmentConfiguration() *ChannelPropertiesCardsInstallmentConfiguration {
	this := ChannelPropertiesCardsInstallmentConfiguration{}
	return &this
}

// NewChannelPropertiesCardsInstallmentConfigurationWithDefaults instantiates a new ChannelPropertiesCardsInstallmentConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPropertiesCardsInstallmentConfigurationWithDefaults() *ChannelPropertiesCardsInstallmentConfiguration {
	this := ChannelPropertiesCardsInstallmentConfiguration{}
	return &this
}

// GetAllowFullPayment returns the AllowFullPayment field value if set, zero value otherwise.
func (o *ChannelPropertiesCardsInstallmentConfiguration) GetAllowFullPayment() bool {
	if o == nil || utils.IsNil(o.AllowFullPayment) {
		var ret bool
		return ret
	}
	return *o.AllowFullPayment
}

// GetAllowFullPaymentOk returns a tuple with the AllowFullPayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPropertiesCardsInstallmentConfiguration) GetAllowFullPaymentOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.AllowFullPayment) {
		return nil, false
	}
	return o.AllowFullPayment, true
}

// HasAllowFullPayment returns a boolean if a field has been set.
func (o *ChannelPropertiesCardsInstallmentConfiguration) HasAllowFullPayment() bool {
	if o != nil && !utils.IsNil(o.AllowFullPayment) {
		return true
	}

	return false
}

// SetAllowFullPayment gets a reference to the given bool and assigns it to the AllowFullPayment field.
func (o *ChannelPropertiesCardsInstallmentConfiguration) SetAllowFullPayment(v bool) {
	o.AllowFullPayment = &v
}

// GetAllowedTerms returns the AllowedTerms field value if set, zero value otherwise.
func (o *ChannelPropertiesCardsInstallmentConfiguration) GetAllowedTerms() []ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner {
	if o == nil || utils.IsNil(o.AllowedTerms) {
		var ret []ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner
		return ret
	}
	return o.AllowedTerms
}

// GetAllowedTermsOk returns a tuple with the AllowedTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPropertiesCardsInstallmentConfiguration) GetAllowedTermsOk() ([]ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner, bool) {
	if o == nil || utils.IsNil(o.AllowedTerms) {
		return nil, false
	}
	return o.AllowedTerms, true
}

// HasAllowedTerms returns a boolean if a field has been set.
func (o *ChannelPropertiesCardsInstallmentConfiguration) HasAllowedTerms() bool {
	if o != nil && !utils.IsNil(o.AllowedTerms) {
		return true
	}

	return false
}

// SetAllowedTerms gets a reference to the given []ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner and assigns it to the AllowedTerms field.
func (o *ChannelPropertiesCardsInstallmentConfiguration) SetAllowedTerms(v []ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) {
	o.AllowedTerms = v
}

func (o ChannelPropertiesCardsInstallmentConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelPropertiesCardsInstallmentConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AllowFullPayment) {
		toSerialize["allow_full_payment"] = o.AllowFullPayment
	}
	if !utils.IsNil(o.AllowedTerms) {
		toSerialize["allowed_terms"] = o.AllowedTerms
	}
	return toSerialize, nil
}

type NullableChannelPropertiesCardsInstallmentConfiguration struct {
	value *ChannelPropertiesCardsInstallmentConfiguration
	isSet bool
}

func (v NullableChannelPropertiesCardsInstallmentConfiguration) Get() *ChannelPropertiesCardsInstallmentConfiguration {
	return v.value
}

func (v *NullableChannelPropertiesCardsInstallmentConfiguration) Set(val *ChannelPropertiesCardsInstallmentConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPropertiesCardsInstallmentConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPropertiesCardsInstallmentConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPropertiesCardsInstallmentConfiguration(val *ChannelPropertiesCardsInstallmentConfiguration) *NullableChannelPropertiesCardsInstallmentConfiguration {
	return &NullableChannelPropertiesCardsInstallmentConfiguration{value: val, isSet: true}
}

func (v NullableChannelPropertiesCardsInstallmentConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPropertiesCardsInstallmentConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


