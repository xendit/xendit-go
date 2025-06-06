/*
Payout Service

This API allows Xendit to send money from an account to a channel (banks, eWallets, retail outlets) from across regions

API version: 1.0.0
*/


package payout

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
)

// checks if the GetPayouts200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GetPayouts200Response{}

// GetPayouts200Response struct for GetPayouts200Response
type GetPayouts200Response struct {
	Data []GetPayouts200ResponseDataInner `json:"data,omitempty"`
	HasMore *bool `json:"has_more,omitempty"`
	Links *GetPayouts200ResponseLinks `json:"links,omitempty"`
}

// NewGetPayouts200Response instantiates a new GetPayouts200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPayouts200Response() *GetPayouts200Response {
	this := GetPayouts200Response{}
	return &this
}

// NewGetPayouts200ResponseWithDefaults instantiates a new GetPayouts200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPayouts200ResponseWithDefaults() *GetPayouts200Response {
	this := GetPayouts200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetPayouts200Response) GetData() []GetPayouts200ResponseDataInner {
	if o == nil || utils.IsNil(o.Data) {
		var ret []GetPayouts200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayouts200Response) GetDataOk() ([]GetPayouts200ResponseDataInner, bool) {
	if o == nil || utils.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetPayouts200Response) HasData() bool {
	if o != nil && !utils.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetPayouts200ResponseDataInner and assigns it to the Data field.
func (o *GetPayouts200Response) SetData(v []GetPayouts200ResponseDataInner) {
	o.Data = v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *GetPayouts200Response) GetHasMore() bool {
	if o == nil || utils.IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayouts200Response) GetHasMoreOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *GetPayouts200Response) HasHasMore() bool {
	if o != nil && !utils.IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *GetPayouts200Response) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetPayouts200Response) GetLinks() GetPayouts200ResponseLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret GetPayouts200ResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayouts200Response) GetLinksOk() (*GetPayouts200ResponseLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetPayouts200Response) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GetPayouts200ResponseLinks and assigns it to the Links field.
func (o *GetPayouts200Response) SetLinks(v GetPayouts200ResponseLinks) {
	o.Links = &v
}

func (o GetPayouts200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPayouts200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !utils.IsNil(o.HasMore) {
		toSerialize["has_more"] = o.HasMore
	}
	if !utils.IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableGetPayouts200Response struct {
	value *GetPayouts200Response
	isSet bool
}

func (v NullableGetPayouts200Response) Get() *GetPayouts200Response {
	return v.value
}

func (v *NullableGetPayouts200Response) Set(val *GetPayouts200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPayouts200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPayouts200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPayouts200Response(val *GetPayouts200Response) *NullableGetPayouts200Response {
	return &NullableGetPayouts200Response{value: val, isSet: true}
}

func (v NullableGetPayouts200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPayouts200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


