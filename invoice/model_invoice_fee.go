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

// checks if the InvoiceFee type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InvoiceFee{}

// InvoiceFee An object representing internal details for a fee associated with an invoice.
type InvoiceFee struct {
	// The type of fee.
	Type string `json:"type"`
	// The value or amount of the fee.
	Value float32 `json:"value"`
}

// NewInvoiceFee instantiates a new InvoiceFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceFee(type_ string, value float32) *InvoiceFee {
	this := InvoiceFee{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewInvoiceFeeWithDefaults instantiates a new InvoiceFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceFeeWithDefaults() *InvoiceFee {
	this := InvoiceFee{}
	return &this
}

// GetType returns the Type field value
func (o *InvoiceFee) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InvoiceFee) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InvoiceFee) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *InvoiceFee) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *InvoiceFee) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *InvoiceFee) SetValue(v float32) {
	o.Value = v
}

func (o InvoiceFee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceFee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableInvoiceFee struct {
	value *InvoiceFee
	isSet bool
}

func (v NullableInvoiceFee) Get() *InvoiceFee {
	return v.value
}

func (v *NullableInvoiceFee) Set(val *InvoiceFee) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceFee) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceFee(val *InvoiceFee) *NullableInvoiceFee {
	return &NullableInvoiceFee{value: val, isSet: true}
}

func (v NullableInvoiceFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


