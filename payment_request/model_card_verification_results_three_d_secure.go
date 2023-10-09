/*
Payment Requests

This API is used for Payment Requests

API version: 1.44.1
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the CardVerificationResultsThreeDSecure type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CardVerificationResultsThreeDSecure{}

// CardVerificationResultsThreeDSecure struct for CardVerificationResultsThreeDSecure
type CardVerificationResultsThreeDSecure struct {
	ThreeDSecureFlow NullableString `json:"three_d_secure_flow,omitempty"`
	EciCode NullableString `json:"eci_code,omitempty"`
	ThreeDSecureResult NullableString `json:"three_d_secure_result,omitempty"`
	ThreeDSecureResultReason NullableString `json:"three_d_secure_result_reason,omitempty"`
	ThreeDSecureVersion NullableString `json:"three_d_secure_version,omitempty"`
}

// NewCardVerificationResultsThreeDSecure instantiates a new CardVerificationResultsThreeDSecure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardVerificationResultsThreeDSecure() *CardVerificationResultsThreeDSecure {
	this := CardVerificationResultsThreeDSecure{}
	return &this
}

// NewCardVerificationResultsThreeDSecureWithDefaults instantiates a new CardVerificationResultsThreeDSecure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardVerificationResultsThreeDSecureWithDefaults() *CardVerificationResultsThreeDSecure {
	this := CardVerificationResultsThreeDSecure{}
	return &this
}

// GetThreeDSecureFlow returns the ThreeDSecureFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardVerificationResultsThreeDSecure) GetThreeDSecureFlow() string {
	if o == nil || utils.IsNil(o.ThreeDSecureFlow.Get()) {
		var ret string
		return ret
	}
	return *o.ThreeDSecureFlow.Get()
}

// GetThreeDSecureFlowOk returns a tuple with the ThreeDSecureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardVerificationResultsThreeDSecure) GetThreeDSecureFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThreeDSecureFlow.Get(), o.ThreeDSecureFlow.IsSet()
}

// HasThreeDSecureFlow returns a boolean if a field has been set.
func (o *CardVerificationResultsThreeDSecure) HasThreeDSecureFlow() bool {
	if o != nil && o.ThreeDSecureFlow.IsSet() {
		return true
	}

	return false
}

// SetThreeDSecureFlow gets a reference to the given NullableString and assigns it to the ThreeDSecureFlow field.
func (o *CardVerificationResultsThreeDSecure) SetThreeDSecureFlow(v string) {
	o.ThreeDSecureFlow.Set(&v)
}
// SetThreeDSecureFlowNil sets the value for ThreeDSecureFlow to be an explicit nil
func (o *CardVerificationResultsThreeDSecure) SetThreeDSecureFlowNil() {
	o.ThreeDSecureFlow.Set(nil)
}

// UnsetThreeDSecureFlow ensures that no value is present for ThreeDSecureFlow, not even an explicit nil
func (o *CardVerificationResultsThreeDSecure) UnsetThreeDSecureFlow() {
	o.ThreeDSecureFlow.Unset()
}

// GetEciCode returns the EciCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardVerificationResultsThreeDSecure) GetEciCode() string {
	if o == nil || utils.IsNil(o.EciCode.Get()) {
		var ret string
		return ret
	}
	return *o.EciCode.Get()
}

// GetEciCodeOk returns a tuple with the EciCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardVerificationResultsThreeDSecure) GetEciCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EciCode.Get(), o.EciCode.IsSet()
}

// HasEciCode returns a boolean if a field has been set.
func (o *CardVerificationResultsThreeDSecure) HasEciCode() bool {
	if o != nil && o.EciCode.IsSet() {
		return true
	}

	return false
}

// SetEciCode gets a reference to the given NullableString and assigns it to the EciCode field.
func (o *CardVerificationResultsThreeDSecure) SetEciCode(v string) {
	o.EciCode.Set(&v)
}
// SetEciCodeNil sets the value for EciCode to be an explicit nil
func (o *CardVerificationResultsThreeDSecure) SetEciCodeNil() {
	o.EciCode.Set(nil)
}

// UnsetEciCode ensures that no value is present for EciCode, not even an explicit nil
func (o *CardVerificationResultsThreeDSecure) UnsetEciCode() {
	o.EciCode.Unset()
}

// GetThreeDSecureResult returns the ThreeDSecureResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardVerificationResultsThreeDSecure) GetThreeDSecureResult() string {
	if o == nil || utils.IsNil(o.ThreeDSecureResult.Get()) {
		var ret string
		return ret
	}
	return *o.ThreeDSecureResult.Get()
}

// GetThreeDSecureResultOk returns a tuple with the ThreeDSecureResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardVerificationResultsThreeDSecure) GetThreeDSecureResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThreeDSecureResult.Get(), o.ThreeDSecureResult.IsSet()
}

// HasThreeDSecureResult returns a boolean if a field has been set.
func (o *CardVerificationResultsThreeDSecure) HasThreeDSecureResult() bool {
	if o != nil && o.ThreeDSecureResult.IsSet() {
		return true
	}

	return false
}

// SetThreeDSecureResult gets a reference to the given NullableString and assigns it to the ThreeDSecureResult field.
func (o *CardVerificationResultsThreeDSecure) SetThreeDSecureResult(v string) {
	o.ThreeDSecureResult.Set(&v)
}
// SetThreeDSecureResultNil sets the value for ThreeDSecureResult to be an explicit nil
func (o *CardVerificationResultsThreeDSecure) SetThreeDSecureResultNil() {
	o.ThreeDSecureResult.Set(nil)
}

// UnsetThreeDSecureResult ensures that no value is present for ThreeDSecureResult, not even an explicit nil
func (o *CardVerificationResultsThreeDSecure) UnsetThreeDSecureResult() {
	o.ThreeDSecureResult.Unset()
}

// GetThreeDSecureResultReason returns the ThreeDSecureResultReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardVerificationResultsThreeDSecure) GetThreeDSecureResultReason() string {
	if o == nil || utils.IsNil(o.ThreeDSecureResultReason.Get()) {
		var ret string
		return ret
	}
	return *o.ThreeDSecureResultReason.Get()
}

// GetThreeDSecureResultReasonOk returns a tuple with the ThreeDSecureResultReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardVerificationResultsThreeDSecure) GetThreeDSecureResultReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThreeDSecureResultReason.Get(), o.ThreeDSecureResultReason.IsSet()
}

// HasThreeDSecureResultReason returns a boolean if a field has been set.
func (o *CardVerificationResultsThreeDSecure) HasThreeDSecureResultReason() bool {
	if o != nil && o.ThreeDSecureResultReason.IsSet() {
		return true
	}

	return false
}

// SetThreeDSecureResultReason gets a reference to the given NullableString and assigns it to the ThreeDSecureResultReason field.
func (o *CardVerificationResultsThreeDSecure) SetThreeDSecureResultReason(v string) {
	o.ThreeDSecureResultReason.Set(&v)
}
// SetThreeDSecureResultReasonNil sets the value for ThreeDSecureResultReason to be an explicit nil
func (o *CardVerificationResultsThreeDSecure) SetThreeDSecureResultReasonNil() {
	o.ThreeDSecureResultReason.Set(nil)
}

// UnsetThreeDSecureResultReason ensures that no value is present for ThreeDSecureResultReason, not even an explicit nil
func (o *CardVerificationResultsThreeDSecure) UnsetThreeDSecureResultReason() {
	o.ThreeDSecureResultReason.Unset()
}

// GetThreeDSecureVersion returns the ThreeDSecureVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardVerificationResultsThreeDSecure) GetThreeDSecureVersion() string {
	if o == nil || utils.IsNil(o.ThreeDSecureVersion.Get()) {
		var ret string
		return ret
	}
	return *o.ThreeDSecureVersion.Get()
}

// GetThreeDSecureVersionOk returns a tuple with the ThreeDSecureVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardVerificationResultsThreeDSecure) GetThreeDSecureVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThreeDSecureVersion.Get(), o.ThreeDSecureVersion.IsSet()
}

// HasThreeDSecureVersion returns a boolean if a field has been set.
func (o *CardVerificationResultsThreeDSecure) HasThreeDSecureVersion() bool {
	if o != nil && o.ThreeDSecureVersion.IsSet() {
		return true
	}

	return false
}

// SetThreeDSecureVersion gets a reference to the given NullableString and assigns it to the ThreeDSecureVersion field.
func (o *CardVerificationResultsThreeDSecure) SetThreeDSecureVersion(v string) {
	o.ThreeDSecureVersion.Set(&v)
}
// SetThreeDSecureVersionNil sets the value for ThreeDSecureVersion to be an explicit nil
func (o *CardVerificationResultsThreeDSecure) SetThreeDSecureVersionNil() {
	o.ThreeDSecureVersion.Set(nil)
}

// UnsetThreeDSecureVersion ensures that no value is present for ThreeDSecureVersion, not even an explicit nil
func (o *CardVerificationResultsThreeDSecure) UnsetThreeDSecureVersion() {
	o.ThreeDSecureVersion.Unset()
}

func (o CardVerificationResultsThreeDSecure) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardVerificationResultsThreeDSecure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ThreeDSecureFlow.IsSet() {
		toSerialize["three_d_secure_flow"] = o.ThreeDSecureFlow.Get()
        if o.ThreeDSecureFlow.Get() != nil && (*o.ThreeDSecureFlow.Get() != "CHALLENGE" && *o.ThreeDSecureFlow.Get() != "FRICTIONLESS") {
            toSerialize["three_d_secure_flow"] = nil
            return toSerialize, utils.NewError("invalid value for ThreeDSecureFlow when marshalling to JSON, must be one of CHALLENGE, FRICTIONLESS")
        }
    }
	if o.EciCode.IsSet() {
		toSerialize["eci_code"] = o.EciCode.Get()
    }
	if o.ThreeDSecureResult.IsSet() {
		toSerialize["three_d_secure_result"] = o.ThreeDSecureResult.Get()
    }
	if o.ThreeDSecureResultReason.IsSet() {
		toSerialize["three_d_secure_result_reason"] = o.ThreeDSecureResultReason.Get()
    }
	if o.ThreeDSecureVersion.IsSet() {
		toSerialize["three_d_secure_version"] = o.ThreeDSecureVersion.Get()
    }
	return toSerialize, nil
}

type NullableCardVerificationResultsThreeDSecure struct {
	value *CardVerificationResultsThreeDSecure
	isSet bool
}

func (v NullableCardVerificationResultsThreeDSecure) Get() *CardVerificationResultsThreeDSecure {
	return v.value
}

func (v *NullableCardVerificationResultsThreeDSecure) Set(val *CardVerificationResultsThreeDSecure) {
	v.value = val
	v.isSet = true
}

func (v NullableCardVerificationResultsThreeDSecure) IsSet() bool {
	return v.isSet
}

func (v *NullableCardVerificationResultsThreeDSecure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardVerificationResultsThreeDSecure(val *CardVerificationResultsThreeDSecure) *NullableCardVerificationResultsThreeDSecure {
	return &NullableCardVerificationResultsThreeDSecure{value: val, isSet: true}
}

func (v NullableCardVerificationResultsThreeDSecure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardVerificationResultsThreeDSecure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


