/*
Payment Requests

This API is used for Payment Requests

API version: 1.45.2
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the DirectDebitChannelPropertiesBankRedirect type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DirectDebitChannelPropertiesBankRedirect{}

// DirectDebitChannelPropertiesBankRedirect Direct Debit Bank Account Channel Properties
type DirectDebitChannelPropertiesBankRedirect struct {
	// Mobile number of the customer that is registered to channel
	MobileNumber *string `json:"mobile_number,omitempty"`
	SuccessReturnUrl *string `json:"success_return_url,omitempty"`
	FailureReturnUrl *string `json:"failure_return_url,omitempty"`
}

// NewDirectDebitChannelPropertiesBankRedirect instantiates a new DirectDebitChannelPropertiesBankRedirect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectDebitChannelPropertiesBankRedirect() *DirectDebitChannelPropertiesBankRedirect {
	this := DirectDebitChannelPropertiesBankRedirect{}
	return &this
}

// NewDirectDebitChannelPropertiesBankRedirectWithDefaults instantiates a new DirectDebitChannelPropertiesBankRedirect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectDebitChannelPropertiesBankRedirectWithDefaults() *DirectDebitChannelPropertiesBankRedirect {
	this := DirectDebitChannelPropertiesBankRedirect{}
	return &this
}

// GetMobileNumber returns the MobileNumber field value if set, zero value otherwise.
func (o *DirectDebitChannelPropertiesBankRedirect) GetMobileNumber() string {
	if o == nil || utils.IsNil(o.MobileNumber) {
		var ret string
		return ret
	}
	return *o.MobileNumber
}

// GetMobileNumberOk returns a tuple with the MobileNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDebitChannelPropertiesBankRedirect) GetMobileNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MobileNumber) {
		return nil, false
	}
	return o.MobileNumber, true
}

// HasMobileNumber returns a boolean if a field has been set.
func (o *DirectDebitChannelPropertiesBankRedirect) HasMobileNumber() bool {
	if o != nil && !utils.IsNil(o.MobileNumber) {
		return true
	}

	return false
}

// SetMobileNumber gets a reference to the given string and assigns it to the MobileNumber field.
func (o *DirectDebitChannelPropertiesBankRedirect) SetMobileNumber(v string) {
	o.MobileNumber = &v
}

// GetSuccessReturnUrl returns the SuccessReturnUrl field value if set, zero value otherwise.
func (o *DirectDebitChannelPropertiesBankRedirect) GetSuccessReturnUrl() string {
	if o == nil || utils.IsNil(o.SuccessReturnUrl) {
		var ret string
		return ret
	}
	return *o.SuccessReturnUrl
}

// GetSuccessReturnUrlOk returns a tuple with the SuccessReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDebitChannelPropertiesBankRedirect) GetSuccessReturnUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SuccessReturnUrl) {
		return nil, false
	}
	return o.SuccessReturnUrl, true
}

// HasSuccessReturnUrl returns a boolean if a field has been set.
func (o *DirectDebitChannelPropertiesBankRedirect) HasSuccessReturnUrl() bool {
	if o != nil && !utils.IsNil(o.SuccessReturnUrl) {
		return true
	}

	return false
}

// SetSuccessReturnUrl gets a reference to the given string and assigns it to the SuccessReturnUrl field.
func (o *DirectDebitChannelPropertiesBankRedirect) SetSuccessReturnUrl(v string) {
	o.SuccessReturnUrl = &v
}

// GetFailureReturnUrl returns the FailureReturnUrl field value if set, zero value otherwise.
func (o *DirectDebitChannelPropertiesBankRedirect) GetFailureReturnUrl() string {
	if o == nil || utils.IsNil(o.FailureReturnUrl) {
		var ret string
		return ret
	}
	return *o.FailureReturnUrl
}

// GetFailureReturnUrlOk returns a tuple with the FailureReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDebitChannelPropertiesBankRedirect) GetFailureReturnUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.FailureReturnUrl) {
		return nil, false
	}
	return o.FailureReturnUrl, true
}

// HasFailureReturnUrl returns a boolean if a field has been set.
func (o *DirectDebitChannelPropertiesBankRedirect) HasFailureReturnUrl() bool {
	if o != nil && !utils.IsNil(o.FailureReturnUrl) {
		return true
	}

	return false
}

// SetFailureReturnUrl gets a reference to the given string and assigns it to the FailureReturnUrl field.
func (o *DirectDebitChannelPropertiesBankRedirect) SetFailureReturnUrl(v string) {
	o.FailureReturnUrl = &v
}

func (o DirectDebitChannelPropertiesBankRedirect) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectDebitChannelPropertiesBankRedirect) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.MobileNumber) {
		toSerialize["mobile_number"] = o.MobileNumber
	}
	if !utils.IsNil(o.SuccessReturnUrl) {
		toSerialize["success_return_url"] = o.SuccessReturnUrl
	}
	if !utils.IsNil(o.FailureReturnUrl) {
		toSerialize["failure_return_url"] = o.FailureReturnUrl
	}
	return toSerialize, nil
}

type NullableDirectDebitChannelPropertiesBankRedirect struct {
	value *DirectDebitChannelPropertiesBankRedirect
	isSet bool
}

func (v NullableDirectDebitChannelPropertiesBankRedirect) Get() *DirectDebitChannelPropertiesBankRedirect {
	return v.value
}

func (v *NullableDirectDebitChannelPropertiesBankRedirect) Set(val *DirectDebitChannelPropertiesBankRedirect) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectDebitChannelPropertiesBankRedirect) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectDebitChannelPropertiesBankRedirect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectDebitChannelPropertiesBankRedirect(val *DirectDebitChannelPropertiesBankRedirect) *NullableDirectDebitChannelPropertiesBankRedirect {
	return &NullableDirectDebitChannelPropertiesBankRedirect{value: val, isSet: true}
}

func (v NullableDirectDebitChannelPropertiesBankRedirect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectDebitChannelPropertiesBankRedirect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


