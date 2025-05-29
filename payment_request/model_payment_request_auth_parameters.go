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

// checks if the PaymentRequestAuthParameters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentRequestAuthParameters{}

// PaymentRequestAuthParameters struct for PaymentRequestAuthParameters
type PaymentRequestAuthParameters struct {
	AuthCode string `json:"auth_code"`
}

// NewPaymentRequestAuthParameters instantiates a new PaymentRequestAuthParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentRequestAuthParameters(authCode string) *PaymentRequestAuthParameters {
	this := PaymentRequestAuthParameters{}
	this.AuthCode = authCode
	return &this
}

// NewPaymentRequestAuthParametersWithDefaults instantiates a new PaymentRequestAuthParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentRequestAuthParametersWithDefaults() *PaymentRequestAuthParameters {
	this := PaymentRequestAuthParameters{}
	return &this
}

// GetAuthCode returns the AuthCode field value
func (o *PaymentRequestAuthParameters) GetAuthCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthCode
}

// GetAuthCodeOk returns a tuple with the AuthCode field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestAuthParameters) GetAuthCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthCode, true
}

// SetAuthCode sets field value
func (o *PaymentRequestAuthParameters) SetAuthCode(v string) {
	o.AuthCode = v
}

func (o PaymentRequestAuthParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentRequestAuthParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auth_code"] = o.AuthCode
	return toSerialize, nil
}

type NullablePaymentRequestAuthParameters struct {
	value *PaymentRequestAuthParameters
	isSet bool
}

func (v NullablePaymentRequestAuthParameters) Get() *PaymentRequestAuthParameters {
	return v.value
}

func (v *NullablePaymentRequestAuthParameters) Set(val *PaymentRequestAuthParameters) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequestAuthParameters) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequestAuthParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequestAuthParameters(val *PaymentRequestAuthParameters) *NullablePaymentRequestAuthParameters {
	return &NullablePaymentRequestAuthParameters{value: val, isSet: true}
}

func (v NullablePaymentRequestAuthParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequestAuthParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


