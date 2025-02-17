/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.8.8
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
)

// checks if the ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner{}

// ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner struct for ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner
type ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner struct {
	// The bank code of the installment provider / issuer
	Issuer *string `json:"issuer,omitempty"`
	// An array containing list of installment tenor available to choose
	Terms []float32 `json:"terms,omitempty"`
}

// NewChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner instantiates a new ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner() *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner {
	this := ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner{}
	return &this
}

// NewChannelPropertiesCardsInstallmentConfigurationAllowedTermsInnerWithDefaults instantiates a new ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPropertiesCardsInstallmentConfigurationAllowedTermsInnerWithDefaults() *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner {
	this := ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner{}
	return &this
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) GetIssuer() string {
	if o == nil || utils.IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) GetIssuerOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) HasIssuer() bool {
	if o != nil && !utils.IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) SetIssuer(v string) {
	o.Issuer = &v
}

// GetTerms returns the Terms field value if set, zero value otherwise.
func (o *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) GetTerms() []float32 {
	if o == nil || utils.IsNil(o.Terms) {
		var ret []float32
		return ret
	}
	return o.Terms
}

// GetTermsOk returns a tuple with the Terms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) GetTermsOk() ([]float32, bool) {
	if o == nil || utils.IsNil(o.Terms) {
		return nil, false
	}
	return o.Terms, true
}

// HasTerms returns a boolean if a field has been set.
func (o *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) HasTerms() bool {
	if o != nil && !utils.IsNil(o.Terms) {
		return true
	}

	return false
}

// SetTerms gets a reference to the given []float32 and assigns it to the Terms field.
func (o *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) SetTerms(v []float32) {
	o.Terms = v
}

func (o ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !utils.IsNil(o.Terms) {
		toSerialize["terms"] = o.Terms
	}
	return toSerialize, nil
}

type NullableChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner struct {
	value *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner
	isSet bool
}

func (v NullableChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) Get() *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner {
	return v.value
}

func (v *NullableChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) Set(val *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner(val *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) *NullableChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner {
	return &NullableChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner{value: val, isSet: true}
}

func (v NullableChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


