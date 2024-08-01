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

// checks if the PaymentMethodAuthParameters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentMethodAuthParameters{}

// PaymentMethodAuthParameters struct for PaymentMethodAuthParameters
type PaymentMethodAuthParameters struct {
	AuthCode string `json:"auth_code"`
}

// NewPaymentMethodAuthParameters instantiates a new PaymentMethodAuthParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodAuthParameters(authCode string) *PaymentMethodAuthParameters {
	this := PaymentMethodAuthParameters{}
	this.AuthCode = authCode
	return &this
}

// NewPaymentMethodAuthParametersWithDefaults instantiates a new PaymentMethodAuthParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodAuthParametersWithDefaults() *PaymentMethodAuthParameters {
	this := PaymentMethodAuthParameters{}
	return &this
}

// GetAuthCode returns the AuthCode field value
func (o *PaymentMethodAuthParameters) GetAuthCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthCode
}

// GetAuthCodeOk returns a tuple with the AuthCode field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodAuthParameters) GetAuthCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthCode, true
}

// SetAuthCode sets field value
func (o *PaymentMethodAuthParameters) SetAuthCode(v string) {
	o.AuthCode = v
}

func (o PaymentMethodAuthParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodAuthParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auth_code"] = o.AuthCode
	return toSerialize, nil
}

type NullablePaymentMethodAuthParameters struct {
	value *PaymentMethodAuthParameters
	isSet bool
}

func (v NullablePaymentMethodAuthParameters) Get() *PaymentMethodAuthParameters {
	return v.value
}

func (v *NullablePaymentMethodAuthParameters) Set(val *PaymentMethodAuthParameters) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodAuthParameters) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodAuthParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodAuthParameters(val *PaymentMethodAuthParameters) *NullablePaymentMethodAuthParameters {
	return &NullablePaymentMethodAuthParameters{value: val, isSet: true}
}

func (v NullablePaymentMethodAuthParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodAuthParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


