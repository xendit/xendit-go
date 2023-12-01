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

// checks if the PaymentRequestCardVerificationResultsThreeDeeSecure type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentRequestCardVerificationResultsThreeDeeSecure{}

// PaymentRequestCardVerificationResultsThreeDeeSecure struct for PaymentRequestCardVerificationResultsThreeDeeSecure
type PaymentRequestCardVerificationResultsThreeDeeSecure struct {
	ThreeDSecureFlow *string `json:"three_d_secure_flow,omitempty"`
	EciCode *string `json:"eci_code,omitempty"`
	ThreeDSecureResult *string `json:"three_d_secure_result,omitempty"`
	ThreeDSecureResultReason NullableString `json:"three_d_secure_result_reason,omitempty"`
	ThreeDSecureVersion *string `json:"three_d_secure_version,omitempty"`
}

// NewPaymentRequestCardVerificationResultsThreeDeeSecure instantiates a new PaymentRequestCardVerificationResultsThreeDeeSecure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentRequestCardVerificationResultsThreeDeeSecure() *PaymentRequestCardVerificationResultsThreeDeeSecure {
	this := PaymentRequestCardVerificationResultsThreeDeeSecure{}
	return &this
}

// NewPaymentRequestCardVerificationResultsThreeDeeSecureWithDefaults instantiates a new PaymentRequestCardVerificationResultsThreeDeeSecure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentRequestCardVerificationResultsThreeDeeSecureWithDefaults() *PaymentRequestCardVerificationResultsThreeDeeSecure {
	this := PaymentRequestCardVerificationResultsThreeDeeSecure{}
	return &this
}

// GetThreeDSecureFlow returns the ThreeDSecureFlow field value if set, zero value otherwise.
func (o *PaymentRequestCardVerificationResultsThreeDeeSecure) GetThreeDSecureFlow() string {
	if o == nil || utils.IsNil(o.ThreeDSecureFlow) {
		var ret string
		return ret
	}
	return *o.ThreeDSecureFlow
}

// GetThreeDSecureFlowOk returns a tuple with the ThreeDSecureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestCardVerificationResultsThreeDeeSecure) GetThreeDSecureFlowOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ThreeDSecureFlow) {
		return nil, false
	}
	return o.ThreeDSecureFlow, true
}

// HasThreeDSecureFlow returns a boolean if a field has been set.
func (o *PaymentRequestCardVerificationResultsThreeDeeSecure) HasThreeDSecureFlow() bool {
	if o != nil && !utils.IsNil(o.ThreeDSecureFlow) {
		return true
	}

	return false
}

// SetThreeDSecureFlow gets a reference to the given string and assigns it to the ThreeDSecureFlow field.
func (o *PaymentRequestCardVerificationResultsThreeDeeSecure) SetThreeDSecureFlow(v string) {
	o.ThreeDSecureFlow = &v
}

// GetEciCode returns the EciCode field value if set, zero value otherwise.
func (o *PaymentRequestCardVerificationResultsThreeDeeSecure) GetEciCode() string {
	if o == nil || utils.IsNil(o.EciCode) {
		var ret string
		return ret
	}
	return *o.EciCode
}

// GetEciCodeOk returns a tuple with the EciCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestCardVerificationResultsThreeDeeSecure) GetEciCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.EciCode) {
		return nil, false
	}
	return o.EciCode, true
}

// HasEciCode returns a boolean if a field has been set.
func (o *PaymentRequestCardVerificationResultsThreeDeeSecure) HasEciCode() bool {
	if o != nil && !utils.IsNil(o.EciCode) {
		return true
	}

	return false
}

// SetEciCode gets a reference to the given string and assigns it to the EciCode field.
func (o *PaymentRequestCardVerificationResultsThreeDeeSecure) SetEciCode(v string) {
	o.EciCode = &v
}

// GetThreeDSecureResult returns the ThreeDSecureResult field value if set, zero value otherwise.
func (o *PaymentRequestCardVerificationResultsThreeDeeSecure) GetThreeDSecureResult() string {
	if o == nil || utils.IsNil(o.ThreeDSecureResult) {
		var ret string
		return ret
	}
	return *o.ThreeDSecureResult
}

// GetThreeDSecureResultOk returns a tuple with the ThreeDSecureResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestCardVerificationResultsThreeDeeSecure) GetThreeDSecureResultOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ThreeDSecureResult) {
		return nil, false
	}
	return o.ThreeDSecureResult, true
}

// HasThreeDSecureResult returns a boolean if a field has been set.
func (o *PaymentRequestCardVerificationResultsThreeDeeSecure) HasThreeDSecureResult() bool {
	if o != nil && !utils.IsNil(o.ThreeDSecureResult) {
		return true
	}

	return false
}

// SetThreeDSecureResult gets a reference to the given string and assigns it to the ThreeDSecureResult field.
func (o *PaymentRequestCardVerificationResultsThreeDeeSecure) SetThreeDSecureResult(v string) {
	o.ThreeDSecureResult = &v
}

// GetThreeDSecureResultReason returns the ThreeDSecureResultReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentRequestCardVerificationResultsThreeDeeSecure) GetThreeDSecureResultReason() string {
	if o == nil || utils.IsNil(o.ThreeDSecureResultReason.Get()) {
		var ret string
		return ret
	}
	return *o.ThreeDSecureResultReason.Get()
}

// GetThreeDSecureResultReasonOk returns a tuple with the ThreeDSecureResultReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestCardVerificationResultsThreeDeeSecure) GetThreeDSecureResultReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThreeDSecureResultReason.Get(), o.ThreeDSecureResultReason.IsSet()
}

// HasThreeDSecureResultReason returns a boolean if a field has been set.
func (o *PaymentRequestCardVerificationResultsThreeDeeSecure) HasThreeDSecureResultReason() bool {
	if o != nil && o.ThreeDSecureResultReason.IsSet() {
		return true
	}

	return false
}

// SetThreeDSecureResultReason gets a reference to the given NullableString and assigns it to the ThreeDSecureResultReason field.
func (o *PaymentRequestCardVerificationResultsThreeDeeSecure) SetThreeDSecureResultReason(v string) {
	o.ThreeDSecureResultReason.Set(&v)
}
// SetThreeDSecureResultReasonNil sets the value for ThreeDSecureResultReason to be an explicit nil
func (o *PaymentRequestCardVerificationResultsThreeDeeSecure) SetThreeDSecureResultReasonNil() {
	o.ThreeDSecureResultReason.Set(nil)
}

// UnsetThreeDSecureResultReason ensures that no value is present for ThreeDSecureResultReason, not even an explicit nil
func (o *PaymentRequestCardVerificationResultsThreeDeeSecure) UnsetThreeDSecureResultReason() {
	o.ThreeDSecureResultReason.Unset()
}

// GetThreeDSecureVersion returns the ThreeDSecureVersion field value if set, zero value otherwise.
func (o *PaymentRequestCardVerificationResultsThreeDeeSecure) GetThreeDSecureVersion() string {
	if o == nil || utils.IsNil(o.ThreeDSecureVersion) {
		var ret string
		return ret
	}
	return *o.ThreeDSecureVersion
}

// GetThreeDSecureVersionOk returns a tuple with the ThreeDSecureVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestCardVerificationResultsThreeDeeSecure) GetThreeDSecureVersionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ThreeDSecureVersion) {
		return nil, false
	}
	return o.ThreeDSecureVersion, true
}

// HasThreeDSecureVersion returns a boolean if a field has been set.
func (o *PaymentRequestCardVerificationResultsThreeDeeSecure) HasThreeDSecureVersion() bool {
	if o != nil && !utils.IsNil(o.ThreeDSecureVersion) {
		return true
	}

	return false
}

// SetThreeDSecureVersion gets a reference to the given string and assigns it to the ThreeDSecureVersion field.
func (o *PaymentRequestCardVerificationResultsThreeDeeSecure) SetThreeDSecureVersion(v string) {
	o.ThreeDSecureVersion = &v
}

func (o PaymentRequestCardVerificationResultsThreeDeeSecure) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentRequestCardVerificationResultsThreeDeeSecure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ThreeDSecureFlow) {
		toSerialize["three_d_secure_flow"] = o.ThreeDSecureFlow
	}
	if !utils.IsNil(o.EciCode) {
		toSerialize["eci_code"] = o.EciCode
	}
	if !utils.IsNil(o.ThreeDSecureResult) {
		toSerialize["three_d_secure_result"] = o.ThreeDSecureResult
	}
	if o.ThreeDSecureResultReason.IsSet() {
		toSerialize["three_d_secure_result_reason"] = o.ThreeDSecureResultReason.Get()
    }
	if !utils.IsNil(o.ThreeDSecureVersion) {
		toSerialize["three_d_secure_version"] = o.ThreeDSecureVersion
	}
	return toSerialize, nil
}

type NullablePaymentRequestCardVerificationResultsThreeDeeSecure struct {
	value *PaymentRequestCardVerificationResultsThreeDeeSecure
	isSet bool
}

func (v NullablePaymentRequestCardVerificationResultsThreeDeeSecure) Get() *PaymentRequestCardVerificationResultsThreeDeeSecure {
	return v.value
}

func (v *NullablePaymentRequestCardVerificationResultsThreeDeeSecure) Set(val *PaymentRequestCardVerificationResultsThreeDeeSecure) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequestCardVerificationResultsThreeDeeSecure) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequestCardVerificationResultsThreeDeeSecure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequestCardVerificationResultsThreeDeeSecure(val *PaymentRequestCardVerificationResultsThreeDeeSecure) *NullablePaymentRequestCardVerificationResultsThreeDeeSecure {
	return &NullablePaymentRequestCardVerificationResultsThreeDeeSecure{value: val, isSet: true}
}

func (v NullablePaymentRequestCardVerificationResultsThreeDeeSecure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequestCardVerificationResultsThreeDeeSecure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


