/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.6.0
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v4/utils"
)

// checks if the ServerError type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ServerError{}

// ServerError struct for ServerError
type ServerError struct {
	ErrorCode string `json:"error_code"`
	Message string `json:"message"`
}

// NewServerError instantiates a new ServerError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerError(errorCode string, message string) *ServerError {
	this := ServerError{}
	this.ErrorCode = errorCode
	this.Message = message
	return &this
}

// NewServerErrorWithDefaults instantiates a new ServerError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerErrorWithDefaults() *ServerError {
	this := ServerError{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *ServerError) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *ServerError) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *ServerError) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetMessage returns the Message field value
func (o *ServerError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ServerError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ServerError) SetMessage(v string) {
	o.Message = v
}

func (o ServerError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error_code"] = o.ErrorCode
    if o.ErrorCode != "SERVER_ERROR" {
        toSerialize["error_code"] = nil
        return toSerialize, utils.NewError("invalid value for ErrorCode when marshalling to JSON, must be one of SERVER_ERROR")
    }
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableServerError struct {
	value *ServerError
	isSet bool
}

func (v NullableServerError) Get() *ServerError {
	return v.value
}

func (v *NullableServerError) Set(val *ServerError) {
	v.value = val
	v.isSet = true
}

func (v NullableServerError) IsSet() bool {
	return v.isSet
}

func (v *NullableServerError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerError(val *ServerError) *NullableServerError {
	return &NullableServerError{value: val, isSet: true}
}

func (v NullableServerError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


