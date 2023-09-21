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

// checks if the PaymentChannelList type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentChannelList{}

// PaymentChannelList struct for PaymentChannelList
type PaymentChannelList struct {
	// Array of resources that match the provided filters
	Data []PaymentChannel `json:"data,omitempty"`
	// Array of objects that can be used to assist on navigating through the data
	Links []PaymentChannelListLinksInner `json:"links,omitempty"`
	// Indicates whether there are more items in the list
	HasMore *bool `json:"has_more,omitempty"`
}

// NewPaymentChannelList instantiates a new PaymentChannelList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentChannelList() *PaymentChannelList {
	this := PaymentChannelList{}
	return &this
}

// NewPaymentChannelListWithDefaults instantiates a new PaymentChannelList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentChannelListWithDefaults() *PaymentChannelList {
	this := PaymentChannelList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PaymentChannelList) GetData() []PaymentChannel {
	if o == nil || utils.IsNil(o.Data) {
		var ret []PaymentChannel
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelList) GetDataOk() ([]PaymentChannel, bool) {
	if o == nil || utils.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PaymentChannelList) HasData() bool {
	if o != nil && !utils.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PaymentChannel and assigns it to the Data field.
func (o *PaymentChannelList) SetData(v []PaymentChannel) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaymentChannelList) GetLinks() []PaymentChannelListLinksInner {
	if o == nil || utils.IsNil(o.Links) {
		var ret []PaymentChannelListLinksInner
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelList) GetLinksOk() ([]PaymentChannelListLinksInner, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaymentChannelList) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []PaymentChannelListLinksInner and assigns it to the Links field.
func (o *PaymentChannelList) SetLinks(v []PaymentChannelListLinksInner) {
	o.Links = v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *PaymentChannelList) GetHasMore() bool {
	if o == nil || utils.IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelList) GetHasMoreOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *PaymentChannelList) HasHasMore() bool {
	if o != nil && !utils.IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *PaymentChannelList) SetHasMore(v bool) {
	o.HasMore = &v
}

func (o PaymentChannelList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentChannelList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !utils.IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !utils.IsNil(o.HasMore) {
		toSerialize["has_more"] = o.HasMore
	}
	return toSerialize, nil
}

type NullablePaymentChannelList struct {
	value *PaymentChannelList
	isSet bool
}

func (v NullablePaymentChannelList) Get() *PaymentChannelList {
	return v.value
}

func (v *NullablePaymentChannelList) Set(val *PaymentChannelList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentChannelList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentChannelList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentChannelList(val *PaymentChannelList) *NullablePaymentChannelList {
	return &NullablePaymentChannelList{value: val, isSet: true}
}

func (v NullablePaymentChannelList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentChannelList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


