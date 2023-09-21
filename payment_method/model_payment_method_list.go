/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.87.2
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the PaymentMethodList type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentMethodList{}

// PaymentMethodList struct for PaymentMethodList
type PaymentMethodList struct {
	Data []PaymentMethod `json:"data"`
	HasMore *bool `json:"has_more,omitempty"`
}

// NewPaymentMethodList instantiates a new PaymentMethodList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodList(data []PaymentMethod) *PaymentMethodList {
	this := PaymentMethodList{}
	this.Data = data
	return &this
}

// NewPaymentMethodListWithDefaults instantiates a new PaymentMethodList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodListWithDefaults() *PaymentMethodList {
	this := PaymentMethodList{}
	return &this
}

// GetData returns the Data field value
func (o *PaymentMethodList) GetData() []PaymentMethod {
	if o == nil {
		var ret []PaymentMethod
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodList) GetDataOk() ([]PaymentMethod, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PaymentMethodList) SetData(v []PaymentMethod) {
	o.Data = v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *PaymentMethodList) GetHasMore() bool {
	if o == nil || utils.IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodList) GetHasMoreOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *PaymentMethodList) HasHasMore() bool {
	if o != nil && !utils.IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *PaymentMethodList) SetHasMore(v bool) {
	o.HasMore = &v
}

func (o PaymentMethodList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !utils.IsNil(o.HasMore) {
		toSerialize["has_more"] = o.HasMore
	}
	return toSerialize, nil
}

type NullablePaymentMethodList struct {
	value *PaymentMethodList
	isSet bool
}

func (v NullablePaymentMethodList) Get() *PaymentMethodList {
	return v.value
}

func (v *NullablePaymentMethodList) Set(val *PaymentMethodList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodList(val *PaymentMethodList) *NullablePaymentMethodList {
	return &NullablePaymentMethodList{value: val, isSet: true}
}

func (v NullablePaymentMethodList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


