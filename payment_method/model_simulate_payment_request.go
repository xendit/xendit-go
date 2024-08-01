/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.128.0
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
)

// checks if the SimulatePaymentRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SimulatePaymentRequest{}

// SimulatePaymentRequest struct for SimulatePaymentRequest
type SimulatePaymentRequest struct {
	Amount *float64 `json:"amount,omitempty"`
}

// NewSimulatePaymentRequest instantiates a new SimulatePaymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimulatePaymentRequest() *SimulatePaymentRequest {
	this := SimulatePaymentRequest{}
	return &this
}

// NewSimulatePaymentRequestWithDefaults instantiates a new SimulatePaymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimulatePaymentRequestWithDefaults() *SimulatePaymentRequest {
	this := SimulatePaymentRequest{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *SimulatePaymentRequest) GetAmount() float64 {
	if o == nil || utils.IsNil(o.Amount) {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimulatePaymentRequest) GetAmountOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *SimulatePaymentRequest) HasAmount() bool {
	if o != nil && !utils.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *SimulatePaymentRequest) SetAmount(v float64) {
	o.Amount = &v
}

func (o SimulatePaymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimulatePaymentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullableSimulatePaymentRequest struct {
	value *SimulatePaymentRequest
	isSet bool
}

func (v NullableSimulatePaymentRequest) Get() *SimulatePaymentRequest {
	return v.value
}

func (v *NullableSimulatePaymentRequest) Set(val *SimulatePaymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSimulatePaymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSimulatePaymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimulatePaymentRequest(val *SimulatePaymentRequest) *NullableSimulatePaymentRequest {
	return &NullableSimulatePaymentRequest{value: val, isSet: true}
}

func (v NullableSimulatePaymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimulatePaymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


