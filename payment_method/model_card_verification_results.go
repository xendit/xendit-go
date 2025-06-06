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

// checks if the CardVerificationResults type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CardVerificationResults{}

// CardVerificationResults struct for CardVerificationResults
type CardVerificationResults struct {
	ThreeDSecure NullableCardVerificationResultsThreeDSecure `json:"three_d_secure"`
	CvvResult NullableString `json:"cvv_result,omitempty"`
	AddressVerificationResult NullableString `json:"address_verification_result,omitempty"`
}

// NewCardVerificationResults instantiates a new CardVerificationResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardVerificationResults(threeDSecure NullableCardVerificationResultsThreeDSecure) *CardVerificationResults {
	this := CardVerificationResults{}
	this.ThreeDSecure = threeDSecure
	return &this
}

// NewCardVerificationResultsWithDefaults instantiates a new CardVerificationResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardVerificationResultsWithDefaults() *CardVerificationResults {
	this := CardVerificationResults{}
	return &this
}

// GetThreeDSecure returns the ThreeDSecure field value
// If the value is explicit nil, the zero value for CardVerificationResultsThreeDSecure will be returned
func (o *CardVerificationResults) GetThreeDSecure() CardVerificationResultsThreeDSecure {
	if o == nil || o.ThreeDSecure.Get() == nil {
		var ret CardVerificationResultsThreeDSecure
		return ret
	}

	return *o.ThreeDSecure.Get()
}

// GetThreeDSecureOk returns a tuple with the ThreeDSecure field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardVerificationResults) GetThreeDSecureOk() (*CardVerificationResultsThreeDSecure, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThreeDSecure.Get(), o.ThreeDSecure.IsSet()
}

// SetThreeDSecure sets field value
func (o *CardVerificationResults) SetThreeDSecure(v CardVerificationResultsThreeDSecure) {
	o.ThreeDSecure.Set(&v)
}

// GetCvvResult returns the CvvResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardVerificationResults) GetCvvResult() string {
	if o == nil || utils.IsNil(o.CvvResult.Get()) {
		var ret string
		return ret
	}
	return *o.CvvResult.Get()
}

// GetCvvResultOk returns a tuple with the CvvResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardVerificationResults) GetCvvResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CvvResult.Get(), o.CvvResult.IsSet()
}

// HasCvvResult returns a boolean if a field has been set.
func (o *CardVerificationResults) HasCvvResult() bool {
	if o != nil && o.CvvResult.IsSet() {
		return true
	}

	return false
}

// SetCvvResult gets a reference to the given NullableString and assigns it to the CvvResult field.
func (o *CardVerificationResults) SetCvvResult(v string) {
	o.CvvResult.Set(&v)
}
// SetCvvResultNil sets the value for CvvResult to be an explicit nil
func (o *CardVerificationResults) SetCvvResultNil() {
	o.CvvResult.Set(nil)
}

// UnsetCvvResult ensures that no value is present for CvvResult, not even an explicit nil
func (o *CardVerificationResults) UnsetCvvResult() {
	o.CvvResult.Unset()
}

// GetAddressVerificationResult returns the AddressVerificationResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardVerificationResults) GetAddressVerificationResult() string {
	if o == nil || utils.IsNil(o.AddressVerificationResult.Get()) {
		var ret string
		return ret
	}
	return *o.AddressVerificationResult.Get()
}

// GetAddressVerificationResultOk returns a tuple with the AddressVerificationResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardVerificationResults) GetAddressVerificationResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressVerificationResult.Get(), o.AddressVerificationResult.IsSet()
}

// HasAddressVerificationResult returns a boolean if a field has been set.
func (o *CardVerificationResults) HasAddressVerificationResult() bool {
	if o != nil && o.AddressVerificationResult.IsSet() {
		return true
	}

	return false
}

// SetAddressVerificationResult gets a reference to the given NullableString and assigns it to the AddressVerificationResult field.
func (o *CardVerificationResults) SetAddressVerificationResult(v string) {
	o.AddressVerificationResult.Set(&v)
}
// SetAddressVerificationResultNil sets the value for AddressVerificationResult to be an explicit nil
func (o *CardVerificationResults) SetAddressVerificationResultNil() {
	o.AddressVerificationResult.Set(nil)
}

// UnsetAddressVerificationResult ensures that no value is present for AddressVerificationResult, not even an explicit nil
func (o *CardVerificationResults) UnsetAddressVerificationResult() {
	o.AddressVerificationResult.Unset()
}

func (o CardVerificationResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardVerificationResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["three_d_secure"] = o.ThreeDSecure.Get()

	if o.CvvResult.IsSet() {
		toSerialize["cvv_result"] = o.CvvResult.Get()
    }
	if o.AddressVerificationResult.IsSet() {
		toSerialize["address_verification_result"] = o.AddressVerificationResult.Get()
    }
	return toSerialize, nil
}

type NullableCardVerificationResults struct {
	value *CardVerificationResults
	isSet bool
}

func (v NullableCardVerificationResults) Get() *CardVerificationResults {
	return v.value
}

func (v *NullableCardVerificationResults) Set(val *CardVerificationResults) {
	v.value = val
	v.isSet = true
}

func (v NullableCardVerificationResults) IsSet() bool {
	return v.isSet
}

func (v *NullableCardVerificationResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardVerificationResults(val *CardVerificationResults) *NullableCardVerificationResults {
	return &NullableCardVerificationResults{value: val, isSet: true}
}

func (v NullableCardVerificationResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardVerificationResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


