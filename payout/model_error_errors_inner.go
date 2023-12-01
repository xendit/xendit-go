/*
Payout Service

This API allows Xendit to send money from an account to a channel (banks, eWallets, retail outlets) from across regions

API version: 1.0.0
*/


package payout

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v4/utils"
)

// checks if the ErrorErrorsInner type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ErrorErrorsInner{}

// ErrorErrorsInner struct for ErrorErrorsInner
type ErrorErrorsInner struct {
	// Precise location of the error
	Path string `json:"path"`
	// Specific description of the error
	Message string `json:"message"`
}

// NewErrorErrorsInner instantiates a new ErrorErrorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorErrorsInner(path string, message string) *ErrorErrorsInner {
	this := ErrorErrorsInner{}
	this.Path = path
	this.Message = message
	return &this
}

// NewErrorErrorsInnerWithDefaults instantiates a new ErrorErrorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorErrorsInnerWithDefaults() *ErrorErrorsInner {
	this := ErrorErrorsInner{}
	return &this
}

// GetPath returns the Path field value
func (o *ErrorErrorsInner) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ErrorErrorsInner) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *ErrorErrorsInner) SetPath(v string) {
	o.Path = v
}

// GetMessage returns the Message field value
func (o *ErrorErrorsInner) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorErrorsInner) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ErrorErrorsInner) SetMessage(v string) {
	o.Message = v
}

func (o ErrorErrorsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorErrorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableErrorErrorsInner struct {
	value *ErrorErrorsInner
	isSet bool
}

func (v NullableErrorErrorsInner) Get() *ErrorErrorsInner {
	return v.value
}

func (v *NullableErrorErrorsInner) Set(val *ErrorErrorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorErrorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorErrorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorErrorsInner(val *ErrorErrorsInner) *NullableErrorErrorsInner {
	return &NullableErrorErrorsInner{value: val, isSet: true}
}

func (v NullableErrorErrorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorErrorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


