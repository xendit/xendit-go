/*
XENDIT SDK openapi spec

XENDIT SDK openapi spec

API version: 1.0.0
*/


package customer

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v4/utils"
)

// checks if the GetCustomerByReferenceID200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GetCustomerByReferenceID200Response{}

// GetCustomerByReferenceID200Response struct for GetCustomerByReferenceID200Response
type GetCustomerByReferenceID200Response struct {
	HasMore *bool `json:"has_more,omitempty"`
	Data []Customer `json:"data,omitempty"`
}

// NewGetCustomerByReferenceID200Response instantiates a new GetCustomerByReferenceID200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCustomerByReferenceID200Response() *GetCustomerByReferenceID200Response {
	this := GetCustomerByReferenceID200Response{}
	return &this
}

// NewGetCustomerByReferenceID200ResponseWithDefaults instantiates a new GetCustomerByReferenceID200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCustomerByReferenceID200ResponseWithDefaults() *GetCustomerByReferenceID200Response {
	this := GetCustomerByReferenceID200Response{}
	return &this
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *GetCustomerByReferenceID200Response) GetHasMore() bool {
	if o == nil || utils.IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCustomerByReferenceID200Response) GetHasMoreOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *GetCustomerByReferenceID200Response) HasHasMore() bool {
	if o != nil && !utils.IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *GetCustomerByReferenceID200Response) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetCustomerByReferenceID200Response) GetData() []Customer {
	if o == nil || utils.IsNil(o.Data) {
		var ret []Customer
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCustomerByReferenceID200Response) GetDataOk() ([]Customer, bool) {
	if o == nil || utils.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetCustomerByReferenceID200Response) HasData() bool {
	if o != nil && !utils.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Customer and assigns it to the Data field.
func (o *GetCustomerByReferenceID200Response) SetData(v []Customer) {
	o.Data = v
}

func (o GetCustomerByReferenceID200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCustomerByReferenceID200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.HasMore) {
		toSerialize["has_more"] = o.HasMore
	}
	if !utils.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetCustomerByReferenceID200Response struct {
	value *GetCustomerByReferenceID200Response
	isSet bool
}

func (v NullableGetCustomerByReferenceID200Response) Get() *GetCustomerByReferenceID200Response {
	return v.value
}

func (v *NullableGetCustomerByReferenceID200Response) Set(val *GetCustomerByReferenceID200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCustomerByReferenceID200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCustomerByReferenceID200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCustomerByReferenceID200Response(val *GetCustomerByReferenceID200Response) *NullableGetCustomerByReferenceID200Response {
	return &NullableGetCustomerByReferenceID200Response{value: val, isSet: true}
}

func (v NullableGetCustomerByReferenceID200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCustomerByReferenceID200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


