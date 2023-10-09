/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.5.0
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the UnauthorizedError type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &UnauthorizedError{}

// UnauthorizedError An error object used to indicate unauthorized access to an invoice-related resource.
type UnauthorizedError struct {
	// The specific error code associated with the unauthorized access.
	ErrorCode string `json:"error_code"`
	// A human-readable error message providing additional context about the unauthorized access.
	Message string `json:"message"`
}

// NewUnauthorizedError instantiates a new UnauthorizedError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnauthorizedError(errorCode string, message string) *UnauthorizedError {
	this := UnauthorizedError{}
	this.ErrorCode = errorCode
	this.Message = message
	return &this
}

// NewUnauthorizedErrorWithDefaults instantiates a new UnauthorizedError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnauthorizedErrorWithDefaults() *UnauthorizedError {
	this := UnauthorizedError{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *UnauthorizedError) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *UnauthorizedError) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *UnauthorizedError) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetMessage returns the Message field value
func (o *UnauthorizedError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *UnauthorizedError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *UnauthorizedError) SetMessage(v string) {
	o.Message = v
}

func (o UnauthorizedError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnauthorizedError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error_code"] = o.ErrorCode
    if o.ErrorCode != "INVALID_API_KEY" {
        toSerialize["error_code"] = nil
        return toSerialize, utils.NewError("invalid value for ErrorCode when marshalling to JSON, must be one of INVALID_API_KEY")
    }
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableUnauthorizedError struct {
	value *UnauthorizedError
	isSet bool
}

func (v NullableUnauthorizedError) Get() *UnauthorizedError {
	return v.value
}

func (v *NullableUnauthorizedError) Set(val *UnauthorizedError) {
	v.value = val
	v.isSet = true
}

func (v NullableUnauthorizedError) IsSet() bool {
	return v.isSet
}

func (v *NullableUnauthorizedError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnauthorizedError(val *UnauthorizedError) *NullableUnauthorizedError {
	return &NullableUnauthorizedError{value: val, isSet: true}
}

func (v NullableUnauthorizedError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnauthorizedError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


