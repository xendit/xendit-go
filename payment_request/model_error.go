/*
Payment Requests

This API is used for Payment Requests

API version: 1.44.1
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the Error type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Error{}

// Error struct for Error
type Error struct {
	ErrorCode NullableString `json:"error_code,omitempty"`
	Message NullableString `json:"message,omitempty"`
}

// NewError instantiates a new Error object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewError() *Error {
	this := Error{}
	return &this
}

// NewErrorWithDefaults instantiates a new Error object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorWithDefaults() *Error {
	this := Error{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Error) GetErrorCode() string {
	if o == nil || utils.IsNil(o.ErrorCode.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorCode.Get()
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Error) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorCode.Get(), o.ErrorCode.IsSet()
}

// HasErrorCode returns a boolean if a field has been set.
func (o *Error) HasErrorCode() bool {
	if o != nil && o.ErrorCode.IsSet() {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableString and assigns it to the ErrorCode field.
func (o *Error) SetErrorCode(v string) {
	o.ErrorCode.Set(&v)
}
// SetErrorCodeNil sets the value for ErrorCode to be an explicit nil
func (o *Error) SetErrorCodeNil() {
	o.ErrorCode.Set(nil)
}

// UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
func (o *Error) UnsetErrorCode() {
	o.ErrorCode.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Error) GetMessage() string {
	if o == nil || utils.IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Error) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *Error) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *Error) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *Error) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *Error) UnsetMessage() {
	o.Message.Unset()
}

func (o Error) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Error) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorCode.IsSet() {
		toSerialize["error_code"] = o.ErrorCode.Get()
        if o.ErrorCode.Get() != nil && (*o.ErrorCode.Get() != "ACCOUNT_ACCESS_BLOCKED" && *o.ErrorCode.Get() != "ADDRESS_VALIDATION_FAILED" && *o.ErrorCode.Get() != "AMOUNT_MISMATCHED" && *o.ErrorCode.Get() != "API_VALIDATION_ERROR" && *o.ErrorCode.Get() != "AUTHENTICATION_FAILED" && *o.ErrorCode.Get() != "AUTHENTICATION_REQUIRED" && *o.ErrorCode.Get() != "CARD_DECLINED" && *o.ErrorCode.Get() != "CHANNEL_CODE_NOT_SUPPORTED_ERROR" && *o.ErrorCode.Get() != "CHANNEL_NOT_ACTIVATED" && *o.ErrorCode.Get() != "CHANNEL_UNAVAILABLE" && *o.ErrorCode.Get() != "COF_COMBINATION_NOT_ALLOWED_ERROR" && *o.ErrorCode.Get() != "CURRENCY_MISMATCHED" && *o.ErrorCode.Get() != "CUSTOMER_NOT_FOUND_ERROR" && *o.ErrorCode.Get() != "CUSTOMER_PAYMENT_METHOD_MISMATCHED" && *o.ErrorCode.Get() != "DATA_NOT_FOUND" && *o.ErrorCode.Get() != "DATA_NOT_FOUND_ERROR" && *o.ErrorCode.Get() != "DECLINED_BY_ISSUER" && *o.ErrorCode.Get() != "DECLINED_BY_PROCESSOR" && *o.ErrorCode.Get() != "DENIED_PERSON_LIST_MATCHED" && *o.ErrorCode.Get() != "DUPLICATE_ERROR" && *o.ErrorCode.Get() != "DUPLICATE_REFERENCE" && *o.ErrorCode.Get() != "EXCEEDS_CAPTURABLE_AMOUNT" && *o.ErrorCode.Get() != "EXPIRED_CARD" && *o.ErrorCode.Get() != "EXPIRED_OTP_ERROR" && *o.ErrorCode.Get() != "FEATURE_NOT_ACTIVATED" && *o.ErrorCode.Get() != "IDEMPOTENCY_ERROR" && *o.ErrorCode.Get() != "INACTIVE_OR_UNAUTHORIZED_CARD" && *o.ErrorCode.Get() != "INSUFFICIENT_BALANCE" && *o.ErrorCode.Get() != "INVALID_ACCOUNT_DETAILS" && *o.ErrorCode.Get() != "INVALID_CVV" && *o.ErrorCode.Get() != "INVALID_OTP_ERROR" && *o.ErrorCode.Get() != "INVALID_PAYMENT_METHOD" && *o.ErrorCode.Get() != "ISSUER_UNAVAILABLE" && *o.ErrorCode.Get() != "MANUAL_CAPTURE_NOT_SUPPORTED" && *o.ErrorCode.Get() != "MAX_ACCOUNT_LINKING" && *o.ErrorCode.Get() != "MAX_AMOUNT_LIMIT_ERROR" && *o.ErrorCode.Get() != "MAX_OTP_ATTEMPTS_ERROR" && *o.ErrorCode.Get() != "OPERATION_NOT_ALLOWED" && *o.ErrorCode.Get() != "OTP_DELIVERY_ERROR" && *o.ErrorCode.Get() != "PAYMENT_METHOD_NOT_FOUND_ERROR" && *o.ErrorCode.Get() != "PAYMENT_REQUEST_ALREADY_COMPLETED" && *o.ErrorCode.Get() != "PAYMENT_REQUEST_ALREADY_FAILED" && *o.ErrorCode.Get() != "PAYMENT_REQUEST_ALREADY_FULLY_CAPTURED" && *o.ErrorCode.Get() != "PAYMENT_STATUS_FAILED" && *o.ErrorCode.Get() != "PROCESSOR_CONFIGURATION_ERROR" && *o.ErrorCode.Get() != "PROCESSOR_ERROR" && *o.ErrorCode.Get() != "PROCESSOR_TEMPORARILY_UNAVAILABLE" && *o.ErrorCode.Get() != "PROCESSOR_TIMEOUT" && *o.ErrorCode.Get() != "REJECTED_BY_ACQUIRER" && *o.ErrorCode.Get() != "SERVER_ERROR" && *o.ErrorCode.Get() != "STOLEN_CARD" && *o.ErrorCode.Get() != "STRONG_CUSTOMER_AUTHENTICATION_REQUIRED" && *o.ErrorCode.Get() != "SUSPECTED_FRAUDULENT" && *o.ErrorCode.Get() != "UNAUTHORIZED" && *o.ErrorCode.Get() != "DUPLICATED_FIXED_PAYMENT_INSTRUMENT") {
            toSerialize["error_code"] = nil
            return toSerialize, utils.NewError("invalid value for ErrorCode when marshalling to JSON, must be one of ACCOUNT_ACCESS_BLOCKED, ADDRESS_VALIDATION_FAILED, AMOUNT_MISMATCHED, API_VALIDATION_ERROR, AUTHENTICATION_FAILED, AUTHENTICATION_REQUIRED, CARD_DECLINED, CHANNEL_CODE_NOT_SUPPORTED_ERROR, CHANNEL_NOT_ACTIVATED, CHANNEL_UNAVAILABLE, COF_COMBINATION_NOT_ALLOWED_ERROR, CURRENCY_MISMATCHED, CUSTOMER_NOT_FOUND_ERROR, CUSTOMER_PAYMENT_METHOD_MISMATCHED, DATA_NOT_FOUND, DATA_NOT_FOUND_ERROR, DECLINED_BY_ISSUER, DECLINED_BY_PROCESSOR, DENIED_PERSON_LIST_MATCHED, DUPLICATE_ERROR, DUPLICATE_REFERENCE, EXCEEDS_CAPTURABLE_AMOUNT, EXPIRED_CARD, EXPIRED_OTP_ERROR, FEATURE_NOT_ACTIVATED, IDEMPOTENCY_ERROR, INACTIVE_OR_UNAUTHORIZED_CARD, INSUFFICIENT_BALANCE, INVALID_ACCOUNT_DETAILS, INVALID_CVV, INVALID_OTP_ERROR, INVALID_PAYMENT_METHOD, ISSUER_UNAVAILABLE, MANUAL_CAPTURE_NOT_SUPPORTED, MAX_ACCOUNT_LINKING, MAX_AMOUNT_LIMIT_ERROR, MAX_OTP_ATTEMPTS_ERROR, OPERATION_NOT_ALLOWED, OTP_DELIVERY_ERROR, PAYMENT_METHOD_NOT_FOUND_ERROR, PAYMENT_REQUEST_ALREADY_COMPLETED, PAYMENT_REQUEST_ALREADY_FAILED, PAYMENT_REQUEST_ALREADY_FULLY_CAPTURED, PAYMENT_STATUS_FAILED, PROCESSOR_CONFIGURATION_ERROR, PROCESSOR_ERROR, PROCESSOR_TEMPORARILY_UNAVAILABLE, PROCESSOR_TIMEOUT, REJECTED_BY_ACQUIRER, SERVER_ERROR, STOLEN_CARD, STRONG_CUSTOMER_AUTHENTICATION_REQUIRED, SUSPECTED_FRAUDULENT, UNAUTHORIZED, DUPLICATED_FIXED_PAYMENT_INSTRUMENT")
        }
    }
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
    }
	return toSerialize, nil
}

type NullableError struct {
	value *Error
	isSet bool
}

func (v NullableError) Get() *Error {
	return v.value
}

func (v *NullableError) Set(val *Error) {
	v.value = val
	v.isSet = true
}

func (v NullableError) IsSet() bool {
	return v.isSet
}

func (v *NullableError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableError(val *Error) *NullableError {
	return &NullableError{value: val, isSet: true}
}

func (v NullableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


