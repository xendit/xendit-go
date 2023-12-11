/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.6.0
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v4/utils"
)

// checks if the Ewallet type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Ewallet{}

// Ewallet An object representing e-wallet details for invoices.
type Ewallet struct {
	EwalletType EwalletType `json:"ewallet_type"`
}

// NewEwallet instantiates a new Ewallet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEwallet(ewalletType EwalletType) *Ewallet {
	this := Ewallet{}
	this.EwalletType = ewalletType
	return &this
}

// NewEwalletWithDefaults instantiates a new Ewallet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEwalletWithDefaults() *Ewallet {
	this := Ewallet{}
	return &this
}

// GetEwalletType returns the EwalletType field value
func (o *Ewallet) GetEwalletType() EwalletType {
	if o == nil {
		var ret EwalletType
		return ret
	}

	return o.EwalletType
}

// GetEwalletTypeOk returns a tuple with the EwalletType field value
// and a boolean to check if the value has been set.
func (o *Ewallet) GetEwalletTypeOk() (*EwalletType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EwalletType, true
}

// SetEwalletType sets field value
func (o *Ewallet) SetEwalletType(v EwalletType) {
	o.EwalletType = v
}

func (o Ewallet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ewallet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ewallet_type"] = o.EwalletType
	return toSerialize, nil
}

type NullableEwallet struct {
	value *Ewallet
	isSet bool
}

func (v NullableEwallet) Get() *Ewallet {
	return v.value
}

func (v *NullableEwallet) Set(val *Ewallet) {
	v.value = val
	v.isSet = true
}

func (v NullableEwallet) IsSet() bool {
	return v.isSet
}

func (v *NullableEwallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEwallet(val *Ewallet) *NullableEwallet {
	return &NullableEwallet{value: val, isSet: true}
}

func (v NullableEwallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEwallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


