/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.9.0
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
)

// checks if the DirectDebit type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DirectDebit{}

// DirectDebit An object representing direct debit details for invoices.
type DirectDebit struct {
	DirectDebitType DirectDebitType `json:"direct_debit_type"`
}

// NewDirectDebit instantiates a new DirectDebit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectDebit(directDebitType DirectDebitType) *DirectDebit {
	this := DirectDebit{}
	this.DirectDebitType = directDebitType
	return &this
}

// NewDirectDebitWithDefaults instantiates a new DirectDebit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectDebitWithDefaults() *DirectDebit {
	this := DirectDebit{}
	return &this
}

// GetDirectDebitType returns the DirectDebitType field value
func (o *DirectDebit) GetDirectDebitType() DirectDebitType {
	if o == nil {
		var ret DirectDebitType
		return ret
	}

	return o.DirectDebitType
}

// GetDirectDebitTypeOk returns a tuple with the DirectDebitType field value
// and a boolean to check if the value has been set.
func (o *DirectDebit) GetDirectDebitTypeOk() (*DirectDebitType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DirectDebitType, true
}

// SetDirectDebitType sets field value
func (o *DirectDebit) SetDirectDebitType(v DirectDebitType) {
	o.DirectDebitType = v
}

func (o DirectDebit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectDebit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["direct_debit_type"] = o.DirectDebitType
	return toSerialize, nil
}

type NullableDirectDebit struct {
	value *DirectDebit
	isSet bool
}

func (v NullableDirectDebit) Get() *DirectDebit {
	return v.value
}

func (v *NullableDirectDebit) Set(val *DirectDebit) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectDebit) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectDebit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectDebit(val *DirectDebit) *NullableDirectDebit {
	return &NullableDirectDebit{value: val, isSet: true}
}

func (v NullableDirectDebit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectDebit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


