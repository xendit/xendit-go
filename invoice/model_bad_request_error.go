/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.9.0
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
)

// checks if the BadRequestError type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BadRequestError{}

// BadRequestError Response definition for a 400 Bad Request error when creating an invoice.
type BadRequestError struct {
	// The error code indicating the type of error that occurred.
	ErrorCode string `json:"error_code"`
	// A human-readable error message that provides additional information about the error.
	Message string `json:"message"`
}

// NewBadRequestError instantiates a new BadRequestError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBadRequestError(errorCode string, message string) *BadRequestError {
	this := BadRequestError{}
	this.ErrorCode = errorCode
	this.Message = message
	return &this
}

// NewBadRequestErrorWithDefaults instantiates a new BadRequestError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBadRequestErrorWithDefaults() *BadRequestError {
	this := BadRequestError{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *BadRequestError) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *BadRequestError) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *BadRequestError) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetMessage returns the Message field value
func (o *BadRequestError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *BadRequestError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *BadRequestError) SetMessage(v string) {
	o.Message = v
}

func (o BadRequestError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BadRequestError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error_code"] = o.ErrorCode
    if o.ErrorCode != "MAXIMUM_TRANSFER_AMOUNT_ERROR" && o.ErrorCode != "NO_COLLECTION_METHODS_ERROR" && o.ErrorCode != "EMAIL_FORMAT_ERROR" && o.ErrorCode != "UNAVAILABLE_PAYMENT_METHOD_ERROR" && o.ErrorCode != "UNSUPPORTED_CURRENCY" && o.ErrorCode != "MISMATCH_CURRENCY_ERROR" && o.ErrorCode != "INVALID_REMINDER_TIME" {
        toSerialize["error_code"] = nil
        return toSerialize, utils.NewError("invalid value for ErrorCode when marshalling to JSON, must be one of MAXIMUM_TRANSFER_AMOUNT_ERROR, NO_COLLECTION_METHODS_ERROR, EMAIL_FORMAT_ERROR, UNAVAILABLE_PAYMENT_METHOD_ERROR, UNSUPPORTED_CURRENCY, MISMATCH_CURRENCY_ERROR, INVALID_REMINDER_TIME")
    }
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableBadRequestError struct {
	value *BadRequestError
	isSet bool
}

func (v NullableBadRequestError) Get() *BadRequestError {
	return v.value
}

func (v *NullableBadRequestError) Set(val *BadRequestError) {
	v.value = val
	v.isSet = true
}

func (v NullableBadRequestError) IsSet() bool {
	return v.isSet
}

func (v *NullableBadRequestError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBadRequestError(val *BadRequestError) *NullableBadRequestError {
	return &NullableBadRequestError{value: val, isSet: true}
}

func (v NullableBadRequestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBadRequestError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


