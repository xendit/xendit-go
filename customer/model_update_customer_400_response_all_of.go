/*
XENDIT SDK openapi spec

XENDIT SDK openapi spec

API version: 1.0.0
*/


package customer

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v4/utils"
)

// checks if the UpdateCustomer400ResponseAllOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &UpdateCustomer400ResponseAllOf{}

// UpdateCustomer400ResponseAllOf struct for UpdateCustomer400ResponseAllOf
type UpdateCustomer400ResponseAllOf struct {
	ErrorCode *string `json:"error_code,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

// NewUpdateCustomer400ResponseAllOf instantiates a new UpdateCustomer400ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCustomer400ResponseAllOf() *UpdateCustomer400ResponseAllOf {
	this := UpdateCustomer400ResponseAllOf{}
	return &this
}

// NewUpdateCustomer400ResponseAllOfWithDefaults instantiates a new UpdateCustomer400ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCustomer400ResponseAllOfWithDefaults() *UpdateCustomer400ResponseAllOf {
	this := UpdateCustomer400ResponseAllOf{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *UpdateCustomer400ResponseAllOf) GetErrorCode() string {
	if o == nil || utils.IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomer400ResponseAllOf) GetErrorCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *UpdateCustomer400ResponseAllOf) HasErrorCode() bool {
	if o != nil && !utils.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *UpdateCustomer400ResponseAllOf) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCustomer400ResponseAllOf) GetMessage() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCustomer400ResponseAllOf) GetMessageOk() (*interface{}, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return &o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UpdateCustomer400ResponseAllOf) HasMessage() bool {
	if o != nil && utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given interface{} and assigns it to the Message field.
func (o *UpdateCustomer400ResponseAllOf) SetMessage(v interface{}) {
	o.Message = v
}

func (o UpdateCustomer400ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCustomer400ResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
    }
	return toSerialize, nil
}

type NullableUpdateCustomer400ResponseAllOf struct {
	value *UpdateCustomer400ResponseAllOf
	isSet bool
}

func (v NullableUpdateCustomer400ResponseAllOf) Get() *UpdateCustomer400ResponseAllOf {
	return v.value
}

func (v *NullableUpdateCustomer400ResponseAllOf) Set(val *UpdateCustomer400ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCustomer400ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCustomer400ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCustomer400ResponseAllOf(val *UpdateCustomer400ResponseAllOf) *NullableUpdateCustomer400ResponseAllOf {
	return &NullableUpdateCustomer400ResponseAllOf{value: val, isSet: true}
}

func (v NullableUpdateCustomer400ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCustomer400ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


