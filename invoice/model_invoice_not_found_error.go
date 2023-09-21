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

// checks if the InvoiceNotFoundError type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InvoiceNotFoundError{}

// InvoiceNotFoundError Response definition for a 404 Not Found error when creating an invoice.
type InvoiceNotFoundError struct {
	// The error code indicating the type of error that occurred.
	ErrorCode string `json:"error_code"`
	// A human-readable error message that provides additional information about the error.
	Message string `json:"message"`
}

// NewInvoiceNotFoundError instantiates a new InvoiceNotFoundError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceNotFoundError(errorCode string, message string) *InvoiceNotFoundError {
	this := InvoiceNotFoundError{}
	this.ErrorCode = errorCode
	this.Message = message
	return &this
}

// NewInvoiceNotFoundErrorWithDefaults instantiates a new InvoiceNotFoundError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceNotFoundErrorWithDefaults() *InvoiceNotFoundError {
	this := InvoiceNotFoundError{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *InvoiceNotFoundError) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *InvoiceNotFoundError) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *InvoiceNotFoundError) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetMessage returns the Message field value
func (o *InvoiceNotFoundError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *InvoiceNotFoundError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *InvoiceNotFoundError) SetMessage(v string) {
	o.Message = v
}

func (o InvoiceNotFoundError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceNotFoundError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error_code"] = o.ErrorCode
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableInvoiceNotFoundError struct {
	value *InvoiceNotFoundError
	isSet bool
}

func (v NullableInvoiceNotFoundError) Get() *InvoiceNotFoundError {
	return v.value
}

func (v *NullableInvoiceNotFoundError) Set(val *InvoiceNotFoundError) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceNotFoundError) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceNotFoundError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceNotFoundError(val *InvoiceNotFoundError) *NullableInvoiceNotFoundError {
	return &NullableInvoiceNotFoundError{value: val, isSet: true}
}

func (v NullableInvoiceNotFoundError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceNotFoundError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


