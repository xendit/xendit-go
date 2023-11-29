/*
Payment Requests

This API is used for Payment Requests

API version: 1.45.2
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the PaymentRequestParametersChannelPropertiesAllOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentRequestParametersChannelPropertiesAllOf{}

// PaymentRequestParametersChannelPropertiesAllOf struct for PaymentRequestParametersChannelPropertiesAllOf
type PaymentRequestParametersChannelPropertiesAllOf struct {
	// Three digit code written on the back of the card (usually called CVV/CVN).
	Cvv *string `json:"cvv,omitempty"`
}

// NewPaymentRequestParametersChannelPropertiesAllOf instantiates a new PaymentRequestParametersChannelPropertiesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentRequestParametersChannelPropertiesAllOf() *PaymentRequestParametersChannelPropertiesAllOf {
	this := PaymentRequestParametersChannelPropertiesAllOf{}
	return &this
}

// NewPaymentRequestParametersChannelPropertiesAllOfWithDefaults instantiates a new PaymentRequestParametersChannelPropertiesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentRequestParametersChannelPropertiesAllOfWithDefaults() *PaymentRequestParametersChannelPropertiesAllOf {
	this := PaymentRequestParametersChannelPropertiesAllOf{}
	return &this
}

// GetCvv returns the Cvv field value if set, zero value otherwise.
func (o *PaymentRequestParametersChannelPropertiesAllOf) GetCvv() string {
	if o == nil || utils.IsNil(o.Cvv) {
		var ret string
		return ret
	}
	return *o.Cvv
}

// GetCvvOk returns a tuple with the Cvv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestParametersChannelPropertiesAllOf) GetCvvOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Cvv) {
		return nil, false
	}
	return o.Cvv, true
}

// HasCvv returns a boolean if a field has been set.
func (o *PaymentRequestParametersChannelPropertiesAllOf) HasCvv() bool {
	if o != nil && !utils.IsNil(o.Cvv) {
		return true
	}

	return false
}

// SetCvv gets a reference to the given string and assigns it to the Cvv field.
func (o *PaymentRequestParametersChannelPropertiesAllOf) SetCvv(v string) {
	o.Cvv = &v
}

func (o PaymentRequestParametersChannelPropertiesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentRequestParametersChannelPropertiesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Cvv) {
		toSerialize["cvv"] = o.Cvv
	}
	return toSerialize, nil
}

type NullablePaymentRequestParametersChannelPropertiesAllOf struct {
	value *PaymentRequestParametersChannelPropertiesAllOf
	isSet bool
}

func (v NullablePaymentRequestParametersChannelPropertiesAllOf) Get() *PaymentRequestParametersChannelPropertiesAllOf {
	return v.value
}

func (v *NullablePaymentRequestParametersChannelPropertiesAllOf) Set(val *PaymentRequestParametersChannelPropertiesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequestParametersChannelPropertiesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequestParametersChannelPropertiesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequestParametersChannelPropertiesAllOf(val *PaymentRequestParametersChannelPropertiesAllOf) *NullablePaymentRequestParametersChannelPropertiesAllOf {
	return &NullablePaymentRequestParametersChannelPropertiesAllOf{value: val, isSet: true}
}

func (v NullablePaymentRequestParametersChannelPropertiesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequestParametersChannelPropertiesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


