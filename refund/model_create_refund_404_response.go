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

// checks if the CreateRefund404Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateRefund404Response{}

// CreateRefund404Response struct for CreateRefund404Response
type CreateRefund404Response struct {
	ErrorCode *string `json:"error_code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewCreateRefund404Response instantiates a new CreateRefund404Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRefund404Response() *CreateRefund404Response {
	this := CreateRefund404Response{}
	return &this
}

// NewCreateRefund404ResponseWithDefaults instantiates a new CreateRefund404Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRefund404ResponseWithDefaults() *CreateRefund404Response {
	this := CreateRefund404Response{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *CreateRefund404Response) GetErrorCode() string {
	if o == nil || utils.IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRefund404Response) GetErrorCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *CreateRefund404Response) HasErrorCode() bool {
	if o != nil && !utils.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *CreateRefund404Response) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateRefund404Response) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRefund404Response) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateRefund404Response) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateRefund404Response) SetMessage(v string) {
	o.Message = &v
}

func (o CreateRefund404Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRefund404Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !utils.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableCreateRefund404Response struct {
	value *CreateRefund404Response
	isSet bool
}

func (v NullableCreateRefund404Response) Get() *CreateRefund404Response {
	return v.value
}

func (v *NullableCreateRefund404Response) Set(val *CreateRefund404Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRefund404Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRefund404Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRefund404Response(val *CreateRefund404Response) *NullableCreateRefund404Response {
	return &NullableCreateRefund404Response{value: val, isSet: true}
}

func (v NullableCreateRefund404Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRefund404Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


