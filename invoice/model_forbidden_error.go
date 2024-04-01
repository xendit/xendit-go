/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.7.6
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v5/utils"
)

// checks if the ForbiddenError type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ForbiddenError{}

// ForbiddenError An error object used to indicate a 403 Forbidden response related to invoice operations.
type ForbiddenError struct {
	// The specific error code indicating that access to the invoice operation is suspended.
	ErrorCode string `json:"error_code"`
	// A human-readable error message providing additional context about the 403 Forbidden response.
	Message string `json:"message"`
}

// NewForbiddenError instantiates a new ForbiddenError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForbiddenError(errorCode string, message string) *ForbiddenError {
	this := ForbiddenError{}
	this.ErrorCode = errorCode
	this.Message = message
	return &this
}

// NewForbiddenErrorWithDefaults instantiates a new ForbiddenError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForbiddenErrorWithDefaults() *ForbiddenError {
	this := ForbiddenError{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *ForbiddenError) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *ForbiddenError) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *ForbiddenError) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetMessage returns the Message field value
func (o *ForbiddenError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ForbiddenError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ForbiddenError) SetMessage(v string) {
	o.Message = v
}

func (o ForbiddenError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForbiddenError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error_code"] = o.ErrorCode
    if o.ErrorCode != "ACCESS_SUSPENDED" {
        toSerialize["error_code"] = nil
        return toSerialize, utils.NewError("invalid value for ErrorCode when marshalling to JSON, must be one of ACCESS_SUSPENDED")
    }
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableForbiddenError struct {
	value *ForbiddenError
	isSet bool
}

func (v NullableForbiddenError) Get() *ForbiddenError {
	return v.value
}

func (v *NullableForbiddenError) Set(val *ForbiddenError) {
	v.value = val
	v.isSet = true
}

func (v NullableForbiddenError) IsSet() bool {
	return v.isSet
}

func (v *NullableForbiddenError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForbiddenError(val *ForbiddenError) *NullableForbiddenError {
	return &NullableForbiddenError{value: val, isSet: true}
}

func (v NullableForbiddenError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForbiddenError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


