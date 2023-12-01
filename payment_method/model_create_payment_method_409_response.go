/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.91.2
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v4/utils"
)

// checks if the CreatePaymentMethod409Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreatePaymentMethod409Response{}

// CreatePaymentMethod409Response struct for CreatePaymentMethod409Response
type CreatePaymentMethod409Response struct {
	ErrorCode *string `json:"error_code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewCreatePaymentMethod409Response instantiates a new CreatePaymentMethod409Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePaymentMethod409Response() *CreatePaymentMethod409Response {
	this := CreatePaymentMethod409Response{}
	return &this
}

// NewCreatePaymentMethod409ResponseWithDefaults instantiates a new CreatePaymentMethod409Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePaymentMethod409ResponseWithDefaults() *CreatePaymentMethod409Response {
	this := CreatePaymentMethod409Response{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *CreatePaymentMethod409Response) GetErrorCode() string {
	if o == nil || utils.IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentMethod409Response) GetErrorCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *CreatePaymentMethod409Response) HasErrorCode() bool {
	if o != nil && !utils.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *CreatePaymentMethod409Response) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreatePaymentMethod409Response) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentMethod409Response) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreatePaymentMethod409Response) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreatePaymentMethod409Response) SetMessage(v string) {
	o.Message = &v
}

func (o CreatePaymentMethod409Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePaymentMethod409Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !utils.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableCreatePaymentMethod409Response struct {
	value *CreatePaymentMethod409Response
	isSet bool
}

func (v NullableCreatePaymentMethod409Response) Get() *CreatePaymentMethod409Response {
	return v.value
}

func (v *NullableCreatePaymentMethod409Response) Set(val *CreatePaymentMethod409Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePaymentMethod409Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePaymentMethod409Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePaymentMethod409Response(val *CreatePaymentMethod409Response) *NullableCreatePaymentMethod409Response {
	return &NullableCreatePaymentMethod409Response{value: val, isSet: true}
}

func (v NullableCreatePaymentMethod409Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePaymentMethod409Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


