/*
Refund Service

This API is used for the unified refund service

API version: 1.3.4
*/


package refund

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
)

// checks if the CreateRefund400Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateRefund400Response{}

// CreateRefund400Response struct for CreateRefund400Response
type CreateRefund400Response struct {
	ErrorCode *string `json:"error_code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewCreateRefund400Response instantiates a new CreateRefund400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRefund400Response() *CreateRefund400Response {
	this := CreateRefund400Response{}
	return &this
}

// NewCreateRefund400ResponseWithDefaults instantiates a new CreateRefund400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRefund400ResponseWithDefaults() *CreateRefund400Response {
	this := CreateRefund400Response{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *CreateRefund400Response) GetErrorCode() string {
	if o == nil || utils.IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRefund400Response) GetErrorCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *CreateRefund400Response) HasErrorCode() bool {
	if o != nil && !utils.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *CreateRefund400Response) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateRefund400Response) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRefund400Response) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateRefund400Response) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateRefund400Response) SetMessage(v string) {
	o.Message = &v
}

func (o CreateRefund400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRefund400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !utils.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableCreateRefund400Response struct {
	value *CreateRefund400Response
	isSet bool
}

func (v NullableCreateRefund400Response) Get() *CreateRefund400Response {
	return v.value
}

func (v *NullableCreateRefund400Response) Set(val *CreateRefund400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRefund400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRefund400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRefund400Response(val *CreateRefund400Response) *NullableCreateRefund400Response {
	return &NullableCreateRefund400Response{value: val, isSet: true}
}

func (v NullableCreateRefund400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRefund400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


