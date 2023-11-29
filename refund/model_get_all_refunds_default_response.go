/*
Refund Service

This API is used for the unified refund service

API version: 1.3.4
*/


package refund

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the GetAllRefundsDefaultResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GetAllRefundsDefaultResponse{}

// GetAllRefundsDefaultResponse struct for GetAllRefundsDefaultResponse
type GetAllRefundsDefaultResponse struct {
	ErrorCode *string `json:"error_code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewGetAllRefundsDefaultResponse instantiates a new GetAllRefundsDefaultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllRefundsDefaultResponse() *GetAllRefundsDefaultResponse {
	this := GetAllRefundsDefaultResponse{}
	return &this
}

// NewGetAllRefundsDefaultResponseWithDefaults instantiates a new GetAllRefundsDefaultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllRefundsDefaultResponseWithDefaults() *GetAllRefundsDefaultResponse {
	this := GetAllRefundsDefaultResponse{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *GetAllRefundsDefaultResponse) GetErrorCode() string {
	if o == nil || utils.IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllRefundsDefaultResponse) GetErrorCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *GetAllRefundsDefaultResponse) HasErrorCode() bool {
	if o != nil && !utils.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *GetAllRefundsDefaultResponse) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetAllRefundsDefaultResponse) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllRefundsDefaultResponse) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetAllRefundsDefaultResponse) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetAllRefundsDefaultResponse) SetMessage(v string) {
	o.Message = &v
}

func (o GetAllRefundsDefaultResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllRefundsDefaultResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !utils.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableGetAllRefundsDefaultResponse struct {
	value *GetAllRefundsDefaultResponse
	isSet bool
}

func (v NullableGetAllRefundsDefaultResponse) Get() *GetAllRefundsDefaultResponse {
	return v.value
}

func (v *NullableGetAllRefundsDefaultResponse) Set(val *GetAllRefundsDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllRefundsDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllRefundsDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllRefundsDefaultResponse(val *GetAllRefundsDefaultResponse) *NullableGetAllRefundsDefaultResponse {
	return &NullableGetAllRefundsDefaultResponse{value: val, isSet: true}
}

func (v NullableGetAllRefundsDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllRefundsDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


