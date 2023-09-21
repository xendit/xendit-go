/*
Payment Requests

This API is used for Payment Requests

API version: 1.44.0
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the CaptureListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CaptureListResponse{}

// CaptureListResponse struct for CaptureListResponse
type CaptureListResponse struct {
	Data []Capture `json:"data"`
	HasMore bool `json:"has_more"`
}

// NewCaptureListResponse instantiates a new CaptureListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptureListResponse(data []Capture, hasMore bool) *CaptureListResponse {
	this := CaptureListResponse{}
	this.Data = data
	this.HasMore = hasMore
	return &this
}

// NewCaptureListResponseWithDefaults instantiates a new CaptureListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptureListResponseWithDefaults() *CaptureListResponse {
	this := CaptureListResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CaptureListResponse) GetData() []Capture {
	if o == nil {
		var ret []Capture
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CaptureListResponse) GetDataOk() ([]Capture, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *CaptureListResponse) SetData(v []Capture) {
	o.Data = v
}

// GetHasMore returns the HasMore field value
func (o *CaptureListResponse) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *CaptureListResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *CaptureListResponse) SetHasMore(v bool) {
	o.HasMore = v
}

func (o CaptureListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaptureListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["has_more"] = o.HasMore
	return toSerialize, nil
}

type NullableCaptureListResponse struct {
	value *CaptureListResponse
	isSet bool
}

func (v NullableCaptureListResponse) Get() *CaptureListResponse {
	return v.value
}

func (v *NullableCaptureListResponse) Set(val *CaptureListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureListResponse(val *CaptureListResponse) *NullableCaptureListResponse {
	return &NullableCaptureListResponse{value: val, isSet: true}
}

func (v NullableCaptureListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


