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

// checks if the PaymentRequestListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentRequestListResponse{}

// PaymentRequestListResponse struct for PaymentRequestListResponse
type PaymentRequestListResponse struct {
	Data []PaymentRequest `json:"data"`
	HasMore bool `json:"has_more"`
}

// NewPaymentRequestListResponse instantiates a new PaymentRequestListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentRequestListResponse(data []PaymentRequest, hasMore bool) *PaymentRequestListResponse {
	this := PaymentRequestListResponse{}
	this.Data = data
	this.HasMore = hasMore
	return &this
}

// NewPaymentRequestListResponseWithDefaults instantiates a new PaymentRequestListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentRequestListResponseWithDefaults() *PaymentRequestListResponse {
	this := PaymentRequestListResponse{}
	return &this
}

// GetData returns the Data field value
func (o *PaymentRequestListResponse) GetData() []PaymentRequest {
	if o == nil {
		var ret []PaymentRequest
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestListResponse) GetDataOk() ([]PaymentRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PaymentRequestListResponse) SetData(v []PaymentRequest) {
	o.Data = v
}

// GetHasMore returns the HasMore field value
func (o *PaymentRequestListResponse) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestListResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *PaymentRequestListResponse) SetHasMore(v bool) {
	o.HasMore = v
}

func (o PaymentRequestListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentRequestListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["has_more"] = o.HasMore
	return toSerialize, nil
}

type NullablePaymentRequestListResponse struct {
	value *PaymentRequestListResponse
	isSet bool
}

func (v NullablePaymentRequestListResponse) Get() *PaymentRequestListResponse {
	return v.value
}

func (v *NullablePaymentRequestListResponse) Set(val *PaymentRequestListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequestListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequestListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequestListResponse(val *PaymentRequestListResponse) *NullablePaymentRequestListResponse {
	return &NullablePaymentRequestListResponse{value: val, isSet: true}
}

func (v NullablePaymentRequestListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequestListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


