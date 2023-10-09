/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.89.1
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the GetAllPaymentMethodsDefaultResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GetAllPaymentMethodsDefaultResponse{}

// GetAllPaymentMethodsDefaultResponse struct for GetAllPaymentMethodsDefaultResponse
type GetAllPaymentMethodsDefaultResponse struct {
	ErrorCode *string `json:"error_code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewGetAllPaymentMethodsDefaultResponse instantiates a new GetAllPaymentMethodsDefaultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllPaymentMethodsDefaultResponse() *GetAllPaymentMethodsDefaultResponse {
	this := GetAllPaymentMethodsDefaultResponse{}
	return &this
}

// NewGetAllPaymentMethodsDefaultResponseWithDefaults instantiates a new GetAllPaymentMethodsDefaultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllPaymentMethodsDefaultResponseWithDefaults() *GetAllPaymentMethodsDefaultResponse {
	this := GetAllPaymentMethodsDefaultResponse{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *GetAllPaymentMethodsDefaultResponse) GetErrorCode() string {
	if o == nil || utils.IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllPaymentMethodsDefaultResponse) GetErrorCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *GetAllPaymentMethodsDefaultResponse) HasErrorCode() bool {
	if o != nil && !utils.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *GetAllPaymentMethodsDefaultResponse) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetAllPaymentMethodsDefaultResponse) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllPaymentMethodsDefaultResponse) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetAllPaymentMethodsDefaultResponse) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetAllPaymentMethodsDefaultResponse) SetMessage(v string) {
	o.Message = &v
}

func (o GetAllPaymentMethodsDefaultResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllPaymentMethodsDefaultResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !utils.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableGetAllPaymentMethodsDefaultResponse struct {
	value *GetAllPaymentMethodsDefaultResponse
	isSet bool
}

func (v NullableGetAllPaymentMethodsDefaultResponse) Get() *GetAllPaymentMethodsDefaultResponse {
	return v.value
}

func (v *NullableGetAllPaymentMethodsDefaultResponse) Set(val *GetAllPaymentMethodsDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllPaymentMethodsDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllPaymentMethodsDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllPaymentMethodsDefaultResponse(val *GetAllPaymentMethodsDefaultResponse) *NullableGetAllPaymentMethodsDefaultResponse {
	return &NullableGetAllPaymentMethodsDefaultResponse{value: val, isSet: true}
}

func (v NullableGetAllPaymentMethodsDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllPaymentMethodsDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


