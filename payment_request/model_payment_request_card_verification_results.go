/*
Payment Requests

This API is used for Payment Requests

API version: 1.45.2
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v4/utils"
)

// checks if the PaymentRequestCardVerificationResults type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentRequestCardVerificationResults{}

// PaymentRequestCardVerificationResults struct for PaymentRequestCardVerificationResults
type PaymentRequestCardVerificationResults struct {
	ThreeDSecure NullablePaymentRequestCardVerificationResultsThreeDeeSecure `json:"three_d_secure"`
	CvvResult NullableString `json:"cvv_result,omitempty"`
	AddressVerificationResult NullableString `json:"address_verification_result,omitempty"`
}

// NewPaymentRequestCardVerificationResults instantiates a new PaymentRequestCardVerificationResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentRequestCardVerificationResults(threeDSecure NullablePaymentRequestCardVerificationResultsThreeDeeSecure) *PaymentRequestCardVerificationResults {
	this := PaymentRequestCardVerificationResults{}
	this.ThreeDSecure = threeDSecure
	return &this
}

// NewPaymentRequestCardVerificationResultsWithDefaults instantiates a new PaymentRequestCardVerificationResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentRequestCardVerificationResultsWithDefaults() *PaymentRequestCardVerificationResults {
	this := PaymentRequestCardVerificationResults{}
	return &this
}

// GetThreeDSecure returns the ThreeDSecure field value
// If the value is explicit nil, the zero value for PaymentRequestCardVerificationResultsThreeDeeSecure will be returned
func (o *PaymentRequestCardVerificationResults) GetThreeDSecure() PaymentRequestCardVerificationResultsThreeDeeSecure {
	if o == nil || o.ThreeDSecure.Get() == nil {
		var ret PaymentRequestCardVerificationResultsThreeDeeSecure
		return ret
	}

	return *o.ThreeDSecure.Get()
}

// GetThreeDSecureOk returns a tuple with the ThreeDSecure field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestCardVerificationResults) GetThreeDSecureOk() (*PaymentRequestCardVerificationResultsThreeDeeSecure, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThreeDSecure.Get(), o.ThreeDSecure.IsSet()
}

// SetThreeDSecure sets field value
func (o *PaymentRequestCardVerificationResults) SetThreeDSecure(v PaymentRequestCardVerificationResultsThreeDeeSecure) {
	o.ThreeDSecure.Set(&v)
}

// GetCvvResult returns the CvvResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentRequestCardVerificationResults) GetCvvResult() string {
	if o == nil || utils.IsNil(o.CvvResult.Get()) {
		var ret string
		return ret
	}
	return *o.CvvResult.Get()
}

// GetCvvResultOk returns a tuple with the CvvResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestCardVerificationResults) GetCvvResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CvvResult.Get(), o.CvvResult.IsSet()
}

// HasCvvResult returns a boolean if a field has been set.
func (o *PaymentRequestCardVerificationResults) HasCvvResult() bool {
	if o != nil && o.CvvResult.IsSet() {
		return true
	}

	return false
}

// SetCvvResult gets a reference to the given NullableString and assigns it to the CvvResult field.
func (o *PaymentRequestCardVerificationResults) SetCvvResult(v string) {
	o.CvvResult.Set(&v)
}
// SetCvvResultNil sets the value for CvvResult to be an explicit nil
func (o *PaymentRequestCardVerificationResults) SetCvvResultNil() {
	o.CvvResult.Set(nil)
}

// UnsetCvvResult ensures that no value is present for CvvResult, not even an explicit nil
func (o *PaymentRequestCardVerificationResults) UnsetCvvResult() {
	o.CvvResult.Unset()
}

// GetAddressVerificationResult returns the AddressVerificationResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentRequestCardVerificationResults) GetAddressVerificationResult() string {
	if o == nil || utils.IsNil(o.AddressVerificationResult.Get()) {
		var ret string
		return ret
	}
	return *o.AddressVerificationResult.Get()
}

// GetAddressVerificationResultOk returns a tuple with the AddressVerificationResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestCardVerificationResults) GetAddressVerificationResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressVerificationResult.Get(), o.AddressVerificationResult.IsSet()
}

// HasAddressVerificationResult returns a boolean if a field has been set.
func (o *PaymentRequestCardVerificationResults) HasAddressVerificationResult() bool {
	if o != nil && o.AddressVerificationResult.IsSet() {
		return true
	}

	return false
}

// SetAddressVerificationResult gets a reference to the given NullableString and assigns it to the AddressVerificationResult field.
func (o *PaymentRequestCardVerificationResults) SetAddressVerificationResult(v string) {
	o.AddressVerificationResult.Set(&v)
}
// SetAddressVerificationResultNil sets the value for AddressVerificationResult to be an explicit nil
func (o *PaymentRequestCardVerificationResults) SetAddressVerificationResultNil() {
	o.AddressVerificationResult.Set(nil)
}

// UnsetAddressVerificationResult ensures that no value is present for AddressVerificationResult, not even an explicit nil
func (o *PaymentRequestCardVerificationResults) UnsetAddressVerificationResult() {
	o.AddressVerificationResult.Unset()
}

func (o PaymentRequestCardVerificationResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentRequestCardVerificationResults) ToMap() (map[string]interface{}, error) {
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

type NullablePaymentRequestCardVerificationResults struct {
	value *PaymentRequestCardVerificationResults
	isSet bool
}

func (v NullablePaymentRequestCardVerificationResults) Get() *PaymentRequestCardVerificationResults {
	return v.value
}

func (v *NullablePaymentRequestCardVerificationResults) Set(val *PaymentRequestCardVerificationResults) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequestCardVerificationResults) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequestCardVerificationResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequestCardVerificationResults(val *PaymentRequestCardVerificationResults) *NullablePaymentRequestCardVerificationResults {
	return &NullablePaymentRequestCardVerificationResults{value: val, isSet: true}
}

func (v NullablePaymentRequestCardVerificationResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequestCardVerificationResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


