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

// checks if the CaptureParameters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CaptureParameters{}

// CaptureParameters struct for CaptureParameters
type CaptureParameters struct {
	ReferenceId NullableString `json:"reference_id,omitempty"`
	CaptureAmount float64 `json:"capture_amount"`
}

// NewCaptureParameters instantiates a new CaptureParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptureParameters(captureAmount float64) *CaptureParameters {
	this := CaptureParameters{}
	this.CaptureAmount = captureAmount
	return &this
}

// NewCaptureParametersWithDefaults instantiates a new CaptureParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptureParametersWithDefaults() *CaptureParameters {
	this := CaptureParameters{}
	return &this
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CaptureParameters) GetReferenceId() string {
	if o == nil || utils.IsNil(o.ReferenceId.Get()) {
		var ret string
		return ret
	}
	return *o.ReferenceId.Get()
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CaptureParameters) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceId.Get(), o.ReferenceId.IsSet()
}

// HasReferenceId returns a boolean if a field has been set.
func (o *CaptureParameters) HasReferenceId() bool {
	if o != nil && o.ReferenceId.IsSet() {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given NullableString and assigns it to the ReferenceId field.
func (o *CaptureParameters) SetReferenceId(v string) {
	o.ReferenceId.Set(&v)
}
// SetReferenceIdNil sets the value for ReferenceId to be an explicit nil
func (o *CaptureParameters) SetReferenceIdNil() {
	o.ReferenceId.Set(nil)
}

// UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
func (o *CaptureParameters) UnsetReferenceId() {
	o.ReferenceId.Unset()
}

// GetCaptureAmount returns the CaptureAmount field value
func (o *CaptureParameters) GetCaptureAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.CaptureAmount
}

// GetCaptureAmountOk returns a tuple with the CaptureAmount field value
// and a boolean to check if the value has been set.
func (o *CaptureParameters) GetCaptureAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaptureAmount, true
}

// SetCaptureAmount sets field value
func (o *CaptureParameters) SetCaptureAmount(v float64) {
	o.CaptureAmount = v
}

func (o CaptureParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaptureParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ReferenceId.IsSet() {
		toSerialize["reference_id"] = o.ReferenceId.Get()
    }
	toSerialize["capture_amount"] = o.CaptureAmount
	return toSerialize, nil
}

type NullableCaptureParameters struct {
	value *CaptureParameters
	isSet bool
}

func (v NullableCaptureParameters) Get() *CaptureParameters {
	return v.value
}

func (v *NullableCaptureParameters) Set(val *CaptureParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureParameters(val *CaptureParameters) *NullableCaptureParameters {
	return &NullableCaptureParameters{value: val, isSet: true}
}

func (v NullableCaptureParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


