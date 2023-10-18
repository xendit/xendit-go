/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.89.2
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the CreatePaymentMethod503Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreatePaymentMethod503Response{}

// CreatePaymentMethod503Response struct for CreatePaymentMethod503Response
type CreatePaymentMethod503Response struct {
	ErrorCode *string `json:"error_code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewCreatePaymentMethod503Response instantiates a new CreatePaymentMethod503Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePaymentMethod503Response() *CreatePaymentMethod503Response {
	this := CreatePaymentMethod503Response{}
	return &this
}

// NewCreatePaymentMethod503ResponseWithDefaults instantiates a new CreatePaymentMethod503Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePaymentMethod503ResponseWithDefaults() *CreatePaymentMethod503Response {
	this := CreatePaymentMethod503Response{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *CreatePaymentMethod503Response) GetErrorCode() string {
	if o == nil || utils.IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentMethod503Response) GetErrorCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *CreatePaymentMethod503Response) HasErrorCode() bool {
	if o != nil && !utils.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *CreatePaymentMethod503Response) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreatePaymentMethod503Response) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentMethod503Response) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreatePaymentMethod503Response) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreatePaymentMethod503Response) SetMessage(v string) {
	o.Message = &v
}

func (o CreatePaymentMethod503Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePaymentMethod503Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !utils.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableCreatePaymentMethod503Response struct {
	value *CreatePaymentMethod503Response
	isSet bool
}

func (v NullableCreatePaymentMethod503Response) Get() *CreatePaymentMethod503Response {
	return v.value
}

func (v *NullableCreatePaymentMethod503Response) Set(val *CreatePaymentMethod503Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePaymentMethod503Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePaymentMethod503Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePaymentMethod503Response(val *CreatePaymentMethod503Response) *NullableCreatePaymentMethod503Response {
	return &NullableCreatePaymentMethod503Response{value: val, isSet: true}
}

func (v NullableCreatePaymentMethod503Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePaymentMethod503Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


