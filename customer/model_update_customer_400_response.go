/*
XENDIT SDK openapi spec

XENDIT SDK openapi spec

API version: 1.0.0
*/


package customer

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
)

// checks if the UpdateCustomer400Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &UpdateCustomer400Response{}

// UpdateCustomer400Response struct for UpdateCustomer400Response
type UpdateCustomer400Response struct {
	ErrorCode string `json:"error_code"`
	Message interface{} `json:"message"`
	Errors []map[string]interface{} `json:"errors,omitempty"`
}

// NewUpdateCustomer400Response instantiates a new UpdateCustomer400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCustomer400Response(errorCode string, message interface{}) *UpdateCustomer400Response {
	this := UpdateCustomer400Response{}
	this.ErrorCode = errorCode
	this.Message = message
	return &this
}

// NewUpdateCustomer400ResponseWithDefaults instantiates a new UpdateCustomer400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCustomer400ResponseWithDefaults() *UpdateCustomer400Response {
	this := UpdateCustomer400Response{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *UpdateCustomer400Response) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *UpdateCustomer400Response) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *UpdateCustomer400Response) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *UpdateCustomer400Response) GetMessage() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCustomer400Response) GetMessageOk() (*interface{}, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *UpdateCustomer400Response) SetMessage(v interface{}) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *UpdateCustomer400Response) GetErrors() []map[string]interface{} {
	if o == nil || utils.IsNil(o.Errors) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomer400Response) GetErrorsOk() ([]map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *UpdateCustomer400Response) HasErrors() bool {
	if o != nil && !utils.IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []map[string]interface{} and assigns it to the Errors field.
func (o *UpdateCustomer400Response) SetErrors(v []map[string]interface{}) {
	o.Errors = v
}

func (o UpdateCustomer400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCustomer400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error_code"] = o.ErrorCode
    if o.ErrorCode != "ENTITY_NOT_FOUND_ERROR" && o.ErrorCode != "CLIENT_NOT_FOUND_ERROR" && o.ErrorCode != "END_CUSTOMER_NOT_FOUND_ERROR" && o.ErrorCode != "DUPLICATE_END_CUSTOMER_ERROR" && o.ErrorCode != "API_VALIDATION_ERROR" {
        toSerialize["error_code"] = nil
        return toSerialize, utils.NewError("invalid value for ErrorCode when marshalling to JSON, must be one of ENTITY_NOT_FOUND_ERROR, CLIENT_NOT_FOUND_ERROR, END_CUSTOMER_NOT_FOUND_ERROR, DUPLICATE_END_CUSTOMER_ERROR, API_VALIDATION_ERROR")
    }
	if o.Message != nil {
		toSerialize["message"] = o.Message
    }
	if !utils.IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableUpdateCustomer400Response struct {
	value *UpdateCustomer400Response
	isSet bool
}

func (v NullableUpdateCustomer400Response) Get() *UpdateCustomer400Response {
	return v.value
}

func (v *NullableUpdateCustomer400Response) Set(val *UpdateCustomer400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCustomer400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCustomer400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCustomer400Response(val *UpdateCustomer400Response) *NullableUpdateCustomer400Response {
	return &NullableUpdateCustomer400Response{value: val, isSet: true}
}

func (v NullableUpdateCustomer400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCustomer400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


