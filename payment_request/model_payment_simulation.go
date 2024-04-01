/*
Payment Requests

This API is used for Payment Requests

API version: 1.59.0
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v5/utils"
)

// checks if the PaymentSimulation type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentSimulation{}

// PaymentSimulation struct for PaymentSimulation
type PaymentSimulation struct {
	Status *string `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewPaymentSimulation instantiates a new PaymentSimulation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentSimulation() *PaymentSimulation {
	this := PaymentSimulation{}
	return &this
}

// NewPaymentSimulationWithDefaults instantiates a new PaymentSimulation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentSimulationWithDefaults() *PaymentSimulation {
	this := PaymentSimulation{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PaymentSimulation) GetStatus() string {
	if o == nil || utils.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentSimulation) GetStatusOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PaymentSimulation) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PaymentSimulation) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PaymentSimulation) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentSimulation) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PaymentSimulation) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PaymentSimulation) SetMessage(v string) {
	o.Message = &v
}

func (o PaymentSimulation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentSimulation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !utils.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullablePaymentSimulation struct {
	value *PaymentSimulation
	isSet bool
}

func (v NullablePaymentSimulation) Get() *PaymentSimulation {
	return v.value
}

func (v *NullablePaymentSimulation) Set(val *PaymentSimulation) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentSimulation) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentSimulation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentSimulation(val *PaymentSimulation) *NullablePaymentSimulation {
	return &NullablePaymentSimulation{value: val, isSet: true}
}

func (v NullablePaymentSimulation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentSimulation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


