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

// checks if the GetAllPaymentMethods404Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GetAllPaymentMethods404Response{}

// GetAllPaymentMethods404Response struct for GetAllPaymentMethods404Response
type GetAllPaymentMethods404Response struct {
	ErrorCode *string `json:"error_code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewGetAllPaymentMethods404Response instantiates a new GetAllPaymentMethods404Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllPaymentMethods404Response() *GetAllPaymentMethods404Response {
	this := GetAllPaymentMethods404Response{}
	return &this
}

// NewGetAllPaymentMethods404ResponseWithDefaults instantiates a new GetAllPaymentMethods404Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllPaymentMethods404ResponseWithDefaults() *GetAllPaymentMethods404Response {
	this := GetAllPaymentMethods404Response{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *GetAllPaymentMethods404Response) GetErrorCode() string {
	if o == nil || utils.IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllPaymentMethods404Response) GetErrorCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *GetAllPaymentMethods404Response) HasErrorCode() bool {
	if o != nil && !utils.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *GetAllPaymentMethods404Response) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetAllPaymentMethods404Response) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllPaymentMethods404Response) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetAllPaymentMethods404Response) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetAllPaymentMethods404Response) SetMessage(v string) {
	o.Message = &v
}

func (o GetAllPaymentMethods404Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllPaymentMethods404Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !utils.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableGetAllPaymentMethods404Response struct {
	value *GetAllPaymentMethods404Response
	isSet bool
}

func (v NullableGetAllPaymentMethods404Response) Get() *GetAllPaymentMethods404Response {
	return v.value
}

func (v *NullableGetAllPaymentMethods404Response) Set(val *GetAllPaymentMethods404Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllPaymentMethods404Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllPaymentMethods404Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllPaymentMethods404Response(val *GetAllPaymentMethods404Response) *NullableGetAllPaymentMethods404Response {
	return &NullableGetAllPaymentMethods404Response{value: val, isSet: true}
}

func (v NullableGetAllPaymentMethods404Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllPaymentMethods404Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


