/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.7.6
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v4/utils"
)

// checks if the InvoiceError404ResponseDefinition type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InvoiceError404ResponseDefinition{}

// InvoiceError404ResponseDefinition An error object used to indicate that the requested resource, in this case, an invoice, was not found.
type InvoiceError404ResponseDefinition struct {
	// The specific error code indicating that the requested invoice was not found.
	ErrorCode string `json:"error_code"`
	// A human-readable error message providing additional context about the resource not being found.
	Message string `json:"message"`
}

// NewInvoiceError404ResponseDefinition instantiates a new InvoiceError404ResponseDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceError404ResponseDefinition(errorCode string, message string) *InvoiceError404ResponseDefinition {
	this := InvoiceError404ResponseDefinition{}
	this.ErrorCode = errorCode
	this.Message = message
	return &this
}

// NewInvoiceError404ResponseDefinitionWithDefaults instantiates a new InvoiceError404ResponseDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceError404ResponseDefinitionWithDefaults() *InvoiceError404ResponseDefinition {
	this := InvoiceError404ResponseDefinition{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *InvoiceError404ResponseDefinition) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *InvoiceError404ResponseDefinition) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *InvoiceError404ResponseDefinition) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetMessage returns the Message field value
func (o *InvoiceError404ResponseDefinition) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *InvoiceError404ResponseDefinition) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *InvoiceError404ResponseDefinition) SetMessage(v string) {
	o.Message = v
}

func (o InvoiceError404ResponseDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceError404ResponseDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error_code"] = o.ErrorCode
    if o.ErrorCode != "INVOICE_NOT_FOUND_ERROR" {
        toSerialize["error_code"] = nil
        return toSerialize, utils.NewError("invalid value for ErrorCode when marshalling to JSON, must be one of INVOICE_NOT_FOUND_ERROR")
    }
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableInvoiceError404ResponseDefinition struct {
	value *InvoiceError404ResponseDefinition
	isSet bool
}

func (v NullableInvoiceError404ResponseDefinition) Get() *InvoiceError404ResponseDefinition {
	return v.value
}

func (v *NullableInvoiceError404ResponseDefinition) Set(val *InvoiceError404ResponseDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceError404ResponseDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceError404ResponseDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceError404ResponseDefinition(val *InvoiceError404ResponseDefinition) *NullableInvoiceError404ResponseDefinition {
	return &NullableInvoiceError404ResponseDefinition{value: val, isSet: true}
}

func (v NullableInvoiceError404ResponseDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceError404ResponseDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


