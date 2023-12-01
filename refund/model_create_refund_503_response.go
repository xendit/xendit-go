/*
Refund Service

This API is used for the unified refund service

API version: 1.3.4
*/


package refund

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v4/utils"
)

// checks if the CreateRefund503Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateRefund503Response{}

// CreateRefund503Response struct for CreateRefund503Response
type CreateRefund503Response struct {
	ErrorCode *string `json:"error_code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewCreateRefund503Response instantiates a new CreateRefund503Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRefund503Response() *CreateRefund503Response {
	this := CreateRefund503Response{}
	return &this
}

// NewCreateRefund503ResponseWithDefaults instantiates a new CreateRefund503Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRefund503ResponseWithDefaults() *CreateRefund503Response {
	this := CreateRefund503Response{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *CreateRefund503Response) GetErrorCode() string {
	if o == nil || utils.IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRefund503Response) GetErrorCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *CreateRefund503Response) HasErrorCode() bool {
	if o != nil && !utils.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *CreateRefund503Response) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateRefund503Response) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRefund503Response) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateRefund503Response) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateRefund503Response) SetMessage(v string) {
	o.Message = &v
}

func (o CreateRefund503Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRefund503Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !utils.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableCreateRefund503Response struct {
	value *CreateRefund503Response
	isSet bool
}

func (v NullableCreateRefund503Response) Get() *CreateRefund503Response {
	return v.value
}

func (v *NullableCreateRefund503Response) Set(val *CreateRefund503Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRefund503Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRefund503Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRefund503Response(val *CreateRefund503Response) *NullableCreateRefund503Response {
	return &NullableCreateRefund503Response{value: val, isSet: true}
}

func (v NullableCreateRefund503Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRefund503Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


