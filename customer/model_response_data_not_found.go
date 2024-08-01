/*
XENDIT SDK openapi spec

XENDIT SDK openapi spec

API version: 1.0.0
*/


package customer

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
)

// checks if the ResponseDataNotFound type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ResponseDataNotFound{}

// ResponseDataNotFound struct for ResponseDataNotFound
type ResponseDataNotFound struct {
	ErrorCode *string `json:"error_code,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

// NewResponseDataNotFound instantiates a new ResponseDataNotFound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseDataNotFound() *ResponseDataNotFound {
	this := ResponseDataNotFound{}
	return &this
}

// NewResponseDataNotFoundWithDefaults instantiates a new ResponseDataNotFound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseDataNotFoundWithDefaults() *ResponseDataNotFound {
	this := ResponseDataNotFound{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *ResponseDataNotFound) GetErrorCode() string {
	if o == nil || utils.IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDataNotFound) GetErrorCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ResponseDataNotFound) HasErrorCode() bool {
	if o != nil && !utils.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *ResponseDataNotFound) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseDataNotFound) GetMessage() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseDataNotFound) GetMessageOk() (*interface{}, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return &o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResponseDataNotFound) HasMessage() bool {
	if o != nil && utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given interface{} and assigns it to the Message field.
func (o *ResponseDataNotFound) SetMessage(v interface{}) {
	o.Message = v
}

func (o ResponseDataNotFound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseDataNotFound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
    }
	return toSerialize, nil
}

type NullableResponseDataNotFound struct {
	value *ResponseDataNotFound
	isSet bool
}

func (v NullableResponseDataNotFound) Get() *ResponseDataNotFound {
	return v.value
}

func (v *NullableResponseDataNotFound) Set(val *ResponseDataNotFound) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseDataNotFound) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseDataNotFound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseDataNotFound(val *ResponseDataNotFound) *NullableResponseDataNotFound {
	return &NullableResponseDataNotFound{value: val, isSet: true}
}

func (v NullableResponseDataNotFound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseDataNotFound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


