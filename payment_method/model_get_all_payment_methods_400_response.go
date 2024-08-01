/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.128.0
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
)

// checks if the GetAllPaymentMethods400Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GetAllPaymentMethods400Response{}

// GetAllPaymentMethods400Response struct for GetAllPaymentMethods400Response
type GetAllPaymentMethods400Response struct {
	ErrorCode *string `json:"error_code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewGetAllPaymentMethods400Response instantiates a new GetAllPaymentMethods400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllPaymentMethods400Response() *GetAllPaymentMethods400Response {
	this := GetAllPaymentMethods400Response{}
	return &this
}

// NewGetAllPaymentMethods400ResponseWithDefaults instantiates a new GetAllPaymentMethods400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllPaymentMethods400ResponseWithDefaults() *GetAllPaymentMethods400Response {
	this := GetAllPaymentMethods400Response{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *GetAllPaymentMethods400Response) GetErrorCode() string {
	if o == nil || utils.IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllPaymentMethods400Response) GetErrorCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *GetAllPaymentMethods400Response) HasErrorCode() bool {
	if o != nil && !utils.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *GetAllPaymentMethods400Response) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetAllPaymentMethods400Response) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllPaymentMethods400Response) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetAllPaymentMethods400Response) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetAllPaymentMethods400Response) SetMessage(v string) {
	o.Message = &v
}

func (o GetAllPaymentMethods400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllPaymentMethods400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !utils.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableGetAllPaymentMethods400Response struct {
	value *GetAllPaymentMethods400Response
	isSet bool
}

func (v NullableGetAllPaymentMethods400Response) Get() *GetAllPaymentMethods400Response {
	return v.value
}

func (v *NullableGetAllPaymentMethods400Response) Set(val *GetAllPaymentMethods400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllPaymentMethods400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllPaymentMethods400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllPaymentMethods400Response(val *GetAllPaymentMethods400Response) *NullableGetAllPaymentMethods400Response {
	return &NullableGetAllPaymentMethods400Response{value: val, isSet: true}
}

func (v NullableGetAllPaymentMethods400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllPaymentMethods400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


