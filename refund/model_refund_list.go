/*
Refund Service

This API is used for the unified refund service

API version: 1.3.4
*/


package refund

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
)

// checks if the RefundList type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RefundList{}

// RefundList struct for RefundList
type RefundList struct {
	Data []Refund `json:"data"`
	HasMore *bool `json:"has_more,omitempty"`
}

// NewRefundList instantiates a new RefundList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefundList(data []Refund) *RefundList {
	this := RefundList{}
	this.Data = data
	return &this
}

// NewRefundListWithDefaults instantiates a new RefundList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefundListWithDefaults() *RefundList {
	this := RefundList{}
	return &this
}

// GetData returns the Data field value
func (o *RefundList) GetData() []Refund {
	if o == nil {
		var ret []Refund
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RefundList) GetDataOk() ([]Refund, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *RefundList) SetData(v []Refund) {
	o.Data = v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *RefundList) GetHasMore() bool {
	if o == nil || utils.IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundList) GetHasMoreOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *RefundList) HasHasMore() bool {
	if o != nil && !utils.IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *RefundList) SetHasMore(v bool) {
	o.HasMore = &v
}

func (o RefundList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefundList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !utils.IsNil(o.HasMore) {
		toSerialize["has_more"] = o.HasMore
	}
	return toSerialize, nil
}

type NullableRefundList struct {
	value *RefundList
	isSet bool
}

func (v NullableRefundList) Get() *RefundList {
	return v.value
}

func (v *NullableRefundList) Set(val *RefundList) {
	v.value = val
	v.isSet = true
}

func (v NullableRefundList) IsSet() bool {
	return v.isSet
}

func (v *NullableRefundList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefundList(val *RefundList) *NullableRefundList {
	return &NullableRefundList{value: val, isSet: true}
}

func (v NullableRefundList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefundList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


