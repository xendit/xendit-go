/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.128.0
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
)

// checks if the GetAllPaymentMethods403Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GetAllPaymentMethods403Response{}

// GetAllPaymentMethods403Response struct for GetAllPaymentMethods403Response
type GetAllPaymentMethods403Response struct {
	ErrorCode *string `json:"error_code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewGetAllPaymentMethods403Response instantiates a new GetAllPaymentMethods403Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllPaymentMethods403Response() *GetAllPaymentMethods403Response {
	this := GetAllPaymentMethods403Response{}
	return &this
}

// NewGetAllPaymentMethods403ResponseWithDefaults instantiates a new GetAllPaymentMethods403Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllPaymentMethods403ResponseWithDefaults() *GetAllPaymentMethods403Response {
	this := GetAllPaymentMethods403Response{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *GetAllPaymentMethods403Response) GetErrorCode() string {
	if o == nil || utils.IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllPaymentMethods403Response) GetErrorCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *GetAllPaymentMethods403Response) HasErrorCode() bool {
	if o != nil && !utils.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *GetAllPaymentMethods403Response) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetAllPaymentMethods403Response) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllPaymentMethods403Response) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetAllPaymentMethods403Response) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetAllPaymentMethods403Response) SetMessage(v string) {
	o.Message = &v
}

func (o GetAllPaymentMethods403Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllPaymentMethods403Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !utils.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableGetAllPaymentMethods403Response struct {
	value *GetAllPaymentMethods403Response
	isSet bool
}

func (v NullableGetAllPaymentMethods403Response) Get() *GetAllPaymentMethods403Response {
	return v.value
}

func (v *NullableGetAllPaymentMethods403Response) Set(val *GetAllPaymentMethods403Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllPaymentMethods403Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllPaymentMethods403Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllPaymentMethods403Response(val *GetAllPaymentMethods403Response) *NullableGetAllPaymentMethods403Response {
	return &NullableGetAllPaymentMethods403Response{value: val, isSet: true}
}

func (v NullableGetAllPaymentMethods403Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllPaymentMethods403Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


