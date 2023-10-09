/*
XENDIT SDK openapi spec

XENDIT SDK openapi spec

API version: 1.0.0
*/


package customer

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the GetCustomerByReferenceID400Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GetCustomerByReferenceID400Response{}

// GetCustomerByReferenceID400Response struct for GetCustomerByReferenceID400Response
type GetCustomerByReferenceID400Response struct {
	ErrorCode string `json:"error_code"`
	Message interface{} `json:"message"`
	Errors []map[string]interface{} `json:"errors,omitempty"`
}

// NewGetCustomerByReferenceID400Response instantiates a new GetCustomerByReferenceID400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCustomerByReferenceID400Response(errorCode string, message interface{}) *GetCustomerByReferenceID400Response {
	this := GetCustomerByReferenceID400Response{}
	this.ErrorCode = errorCode
	this.Message = message
	return &this
}

// NewGetCustomerByReferenceID400ResponseWithDefaults instantiates a new GetCustomerByReferenceID400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCustomerByReferenceID400ResponseWithDefaults() *GetCustomerByReferenceID400Response {
	this := GetCustomerByReferenceID400Response{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *GetCustomerByReferenceID400Response) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *GetCustomerByReferenceID400Response) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *GetCustomerByReferenceID400Response) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GetCustomerByReferenceID400Response) GetMessage() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetCustomerByReferenceID400Response) GetMessageOk() (*interface{}, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *GetCustomerByReferenceID400Response) SetMessage(v interface{}) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetCustomerByReferenceID400Response) GetErrors() []map[string]interface{} {
	if o == nil || utils.IsNil(o.Errors) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCustomerByReferenceID400Response) GetErrorsOk() ([]map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetCustomerByReferenceID400Response) HasErrors() bool {
	if o != nil && !utils.IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []map[string]interface{} and assigns it to the Errors field.
func (o *GetCustomerByReferenceID400Response) SetErrors(v []map[string]interface{}) {
	o.Errors = v
}

func (o GetCustomerByReferenceID400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCustomerByReferenceID400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error_code"] = o.ErrorCode
    if o.ErrorCode != "ENTITY_NOT_FOUND_ERROR" && o.ErrorCode != "CLIENT_NOT_FOUND_ERROR" && o.ErrorCode != "END_CUSTOMER_NOT_FOUND_ERROR" && o.ErrorCode != "API_VALIDATION_ERROR" {
        toSerialize["error_code"] = nil
        return toSerialize, utils.NewError("invalid value for ErrorCode when marshalling to JSON, must be one of ENTITY_NOT_FOUND_ERROR, CLIENT_NOT_FOUND_ERROR, END_CUSTOMER_NOT_FOUND_ERROR, API_VALIDATION_ERROR")
    }
	if o.Message != nil {
		toSerialize["message"] = o.Message
    }
	if !utils.IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetCustomerByReferenceID400Response struct {
	value *GetCustomerByReferenceID400Response
	isSet bool
}

func (v NullableGetCustomerByReferenceID400Response) Get() *GetCustomerByReferenceID400Response {
	return v.value
}

func (v *NullableGetCustomerByReferenceID400Response) Set(val *GetCustomerByReferenceID400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCustomerByReferenceID400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCustomerByReferenceID400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCustomerByReferenceID400Response(val *GetCustomerByReferenceID400Response) *NullableGetCustomerByReferenceID400Response {
	return &NullableGetCustomerByReferenceID400Response{value: val, isSet: true}
}

func (v NullableGetCustomerByReferenceID400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCustomerByReferenceID400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


