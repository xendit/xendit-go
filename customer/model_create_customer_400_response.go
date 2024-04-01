/*
XENDIT SDK openapi spec

XENDIT SDK openapi spec

API version: 1.0.0
*/


package customer

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v5/utils"
)

// checks if the CreateCustomer400Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateCustomer400Response{}

// CreateCustomer400Response struct for CreateCustomer400Response
type CreateCustomer400Response struct {
	ErrorCode string `json:"error_code"`
	Message interface{} `json:"message"`
	Errors []map[string]interface{} `json:"errors,omitempty"`
}

// NewCreateCustomer400Response instantiates a new CreateCustomer400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCustomer400Response(errorCode string, message interface{}) *CreateCustomer400Response {
	this := CreateCustomer400Response{}
	this.ErrorCode = errorCode
	this.Message = message
	return &this
}

// NewCreateCustomer400ResponseWithDefaults instantiates a new CreateCustomer400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCustomer400ResponseWithDefaults() *CreateCustomer400Response {
	this := CreateCustomer400Response{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *CreateCustomer400Response) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *CreateCustomer400Response) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *CreateCustomer400Response) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CreateCustomer400Response) GetMessage() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCustomer400Response) GetMessageOk() (*interface{}, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *CreateCustomer400Response) SetMessage(v interface{}) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateCustomer400Response) GetErrors() []map[string]interface{} {
	if o == nil || utils.IsNil(o.Errors) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomer400Response) GetErrorsOk() ([]map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreateCustomer400Response) HasErrors() bool {
	if o != nil && !utils.IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []map[string]interface{} and assigns it to the Errors field.
func (o *CreateCustomer400Response) SetErrors(v []map[string]interface{}) {
	o.Errors = v
}

func (o CreateCustomer400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCustomer400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error_code"] = o.ErrorCode
    if o.ErrorCode != "DUPLICATE_END_CUSTOMER_ERROR" && o.ErrorCode != "API_VALIDATION_ERROR" {
        toSerialize["error_code"] = nil
        return toSerialize, utils.NewError("invalid value for ErrorCode when marshalling to JSON, must be one of DUPLICATE_END_CUSTOMER_ERROR, API_VALIDATION_ERROR")
    }
	if o.Message != nil {
		toSerialize["message"] = o.Message
    }
	if !utils.IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableCreateCustomer400Response struct {
	value *CreateCustomer400Response
	isSet bool
}

func (v NullableCreateCustomer400Response) Get() *CreateCustomer400Response {
	return v.value
}

func (v *NullableCreateCustomer400Response) Set(val *CreateCustomer400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCustomer400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCustomer400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCustomer400Response(val *CreateCustomer400Response) *NullableCreateCustomer400Response {
	return &NullableCreateCustomer400Response{value: val, isSet: true}
}

func (v NullableCreateCustomer400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCustomer400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


