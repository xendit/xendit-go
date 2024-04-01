/*
Refund Service

This API is used for the unified refund service

API version: 1.3.4
*/


package refund

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v5/utils"
)

// checks if the CreateRefund403Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateRefund403Response{}

// CreateRefund403Response struct for CreateRefund403Response
type CreateRefund403Response struct {
	ErrorCode *string `json:"error_code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewCreateRefund403Response instantiates a new CreateRefund403Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRefund403Response() *CreateRefund403Response {
	this := CreateRefund403Response{}
	return &this
}

// NewCreateRefund403ResponseWithDefaults instantiates a new CreateRefund403Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRefund403ResponseWithDefaults() *CreateRefund403Response {
	this := CreateRefund403Response{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *CreateRefund403Response) GetErrorCode() string {
	if o == nil || utils.IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRefund403Response) GetErrorCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *CreateRefund403Response) HasErrorCode() bool {
	if o != nil && !utils.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *CreateRefund403Response) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateRefund403Response) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRefund403Response) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateRefund403Response) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateRefund403Response) SetMessage(v string) {
	o.Message = &v
}

func (o CreateRefund403Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRefund403Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !utils.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableCreateRefund403Response struct {
	value *CreateRefund403Response
	isSet bool
}

func (v NullableCreateRefund403Response) Get() *CreateRefund403Response {
	return v.value
}

func (v *NullableCreateRefund403Response) Set(val *CreateRefund403Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRefund403Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRefund403Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRefund403Response(val *CreateRefund403Response) *NullableCreateRefund403Response {
	return &NullableCreateRefund403Response{value: val, isSet: true}
}

func (v NullableCreateRefund403Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRefund403Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


