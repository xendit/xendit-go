/*
Refund Service

This API is used for the unified refund service

API version: 1.2.3
*/


package refund

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the CreateRefundDefaultResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateRefundDefaultResponse{}

// CreateRefundDefaultResponse struct for CreateRefundDefaultResponse
type CreateRefundDefaultResponse struct {
	ErrorCode *string `json:"error_code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewCreateRefundDefaultResponse instantiates a new CreateRefundDefaultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRefundDefaultResponse() *CreateRefundDefaultResponse {
	this := CreateRefundDefaultResponse{}
	return &this
}

// NewCreateRefundDefaultResponseWithDefaults instantiates a new CreateRefundDefaultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRefundDefaultResponseWithDefaults() *CreateRefundDefaultResponse {
	this := CreateRefundDefaultResponse{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *CreateRefundDefaultResponse) GetErrorCode() string {
	if o == nil || utils.IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRefundDefaultResponse) GetErrorCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *CreateRefundDefaultResponse) HasErrorCode() bool {
	if o != nil && !utils.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *CreateRefundDefaultResponse) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateRefundDefaultResponse) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRefundDefaultResponse) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateRefundDefaultResponse) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateRefundDefaultResponse) SetMessage(v string) {
	o.Message = &v
}

func (o CreateRefundDefaultResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRefundDefaultResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !utils.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableCreateRefundDefaultResponse struct {
	value *CreateRefundDefaultResponse
	isSet bool
}

func (v NullableCreateRefundDefaultResponse) Get() *CreateRefundDefaultResponse {
	return v.value
}

func (v *NullableCreateRefundDefaultResponse) Set(val *CreateRefundDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRefundDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRefundDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRefundDefaultResponse(val *CreateRefundDefaultResponse) *NullableCreateRefundDefaultResponse {
	return &NullableCreateRefundDefaultResponse{value: val, isSet: true}
}

func (v NullableCreateRefundDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRefundDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


