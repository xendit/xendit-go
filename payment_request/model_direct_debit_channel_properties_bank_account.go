/*
Payment Requests

This API is used for Payment Requests

API version: 1.70.0
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
)

// checks if the DirectDebitChannelPropertiesBankAccount type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DirectDebitChannelPropertiesBankAccount{}

// DirectDebitChannelPropertiesBankAccount Direct Debit Bank Account Channel Properties
type DirectDebitChannelPropertiesBankAccount struct {
	SuccessReturnUrl *string `json:"success_return_url,omitempty"`
	FailureReturnUrl *string `json:"failure_return_url,omitempty"`
	MobileNumber *string `json:"mobile_number,omitempty"`
	IdentityDocumentNumber *string `json:"identity_document_number,omitempty"`
}

// NewDirectDebitChannelPropertiesBankAccount instantiates a new DirectDebitChannelPropertiesBankAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectDebitChannelPropertiesBankAccount() *DirectDebitChannelPropertiesBankAccount {
	this := DirectDebitChannelPropertiesBankAccount{}
	return &this
}

// NewDirectDebitChannelPropertiesBankAccountWithDefaults instantiates a new DirectDebitChannelPropertiesBankAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectDebitChannelPropertiesBankAccountWithDefaults() *DirectDebitChannelPropertiesBankAccount {
	this := DirectDebitChannelPropertiesBankAccount{}
	return &this
}

// GetSuccessReturnUrl returns the SuccessReturnUrl field value if set, zero value otherwise.
func (o *DirectDebitChannelPropertiesBankAccount) GetSuccessReturnUrl() string {
	if o == nil || utils.IsNil(o.SuccessReturnUrl) {
		var ret string
		return ret
	}
	return *o.SuccessReturnUrl
}

// GetSuccessReturnUrlOk returns a tuple with the SuccessReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDebitChannelPropertiesBankAccount) GetSuccessReturnUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SuccessReturnUrl) {
		return nil, false
	}
	return o.SuccessReturnUrl, true
}

// HasSuccessReturnUrl returns a boolean if a field has been set.
func (o *DirectDebitChannelPropertiesBankAccount) HasSuccessReturnUrl() bool {
	if o != nil && !utils.IsNil(o.SuccessReturnUrl) {
		return true
	}

	return false
}

// SetSuccessReturnUrl gets a reference to the given string and assigns it to the SuccessReturnUrl field.
func (o *DirectDebitChannelPropertiesBankAccount) SetSuccessReturnUrl(v string) {
	o.SuccessReturnUrl = &v
}

// GetFailureReturnUrl returns the FailureReturnUrl field value if set, zero value otherwise.
func (o *DirectDebitChannelPropertiesBankAccount) GetFailureReturnUrl() string {
	if o == nil || utils.IsNil(o.FailureReturnUrl) {
		var ret string
		return ret
	}
	return *o.FailureReturnUrl
}

// GetFailureReturnUrlOk returns a tuple with the FailureReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDebitChannelPropertiesBankAccount) GetFailureReturnUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.FailureReturnUrl) {
		return nil, false
	}
	return o.FailureReturnUrl, true
}

// HasFailureReturnUrl returns a boolean if a field has been set.
func (o *DirectDebitChannelPropertiesBankAccount) HasFailureReturnUrl() bool {
	if o != nil && !utils.IsNil(o.FailureReturnUrl) {
		return true
	}

	return false
}

// SetFailureReturnUrl gets a reference to the given string and assigns it to the FailureReturnUrl field.
func (o *DirectDebitChannelPropertiesBankAccount) SetFailureReturnUrl(v string) {
	o.FailureReturnUrl = &v
}

// GetMobileNumber returns the MobileNumber field value if set, zero value otherwise.
func (o *DirectDebitChannelPropertiesBankAccount) GetMobileNumber() string {
	if o == nil || utils.IsNil(o.MobileNumber) {
		var ret string
		return ret
	}
	return *o.MobileNumber
}

// GetMobileNumberOk returns a tuple with the MobileNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDebitChannelPropertiesBankAccount) GetMobileNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MobileNumber) {
		return nil, false
	}
	return o.MobileNumber, true
}

// HasMobileNumber returns a boolean if a field has been set.
func (o *DirectDebitChannelPropertiesBankAccount) HasMobileNumber() bool {
	if o != nil && !utils.IsNil(o.MobileNumber) {
		return true
	}

	return false
}

// SetMobileNumber gets a reference to the given string and assigns it to the MobileNumber field.
func (o *DirectDebitChannelPropertiesBankAccount) SetMobileNumber(v string) {
	o.MobileNumber = &v
}

// GetIdentityDocumentNumber returns the IdentityDocumentNumber field value if set, zero value otherwise.
func (o *DirectDebitChannelPropertiesBankAccount) GetIdentityDocumentNumber() string {
	if o == nil || utils.IsNil(o.IdentityDocumentNumber) {
		var ret string
		return ret
	}
	return *o.IdentityDocumentNumber
}

// GetIdentityDocumentNumberOk returns a tuple with the IdentityDocumentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDebitChannelPropertiesBankAccount) GetIdentityDocumentNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.IdentityDocumentNumber) {
		return nil, false
	}
	return o.IdentityDocumentNumber, true
}

// HasIdentityDocumentNumber returns a boolean if a field has been set.
func (o *DirectDebitChannelPropertiesBankAccount) HasIdentityDocumentNumber() bool {
	if o != nil && !utils.IsNil(o.IdentityDocumentNumber) {
		return true
	}

	return false
}

// SetIdentityDocumentNumber gets a reference to the given string and assigns it to the IdentityDocumentNumber field.
func (o *DirectDebitChannelPropertiesBankAccount) SetIdentityDocumentNumber(v string) {
	o.IdentityDocumentNumber = &v
}

func (o DirectDebitChannelPropertiesBankAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectDebitChannelPropertiesBankAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.SuccessReturnUrl) {
		toSerialize["success_return_url"] = o.SuccessReturnUrl
	}
	if !utils.IsNil(o.FailureReturnUrl) {
		toSerialize["failure_return_url"] = o.FailureReturnUrl
	}
	if !utils.IsNil(o.MobileNumber) {
		toSerialize["mobile_number"] = o.MobileNumber
	}
	if !utils.IsNil(o.IdentityDocumentNumber) {
		toSerialize["identity_document_number"] = o.IdentityDocumentNumber
	}
	return toSerialize, nil
}

type NullableDirectDebitChannelPropertiesBankAccount struct {
	value *DirectDebitChannelPropertiesBankAccount
	isSet bool
}

func (v NullableDirectDebitChannelPropertiesBankAccount) Get() *DirectDebitChannelPropertiesBankAccount {
	return v.value
}

func (v *NullableDirectDebitChannelPropertiesBankAccount) Set(val *DirectDebitChannelPropertiesBankAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectDebitChannelPropertiesBankAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectDebitChannelPropertiesBankAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectDebitChannelPropertiesBankAccount(val *DirectDebitChannelPropertiesBankAccount) *NullableDirectDebitChannelPropertiesBankAccount {
	return &NullableDirectDebitChannelPropertiesBankAccount{value: val, isSet: true}
}

func (v NullableDirectDebitChannelPropertiesBankAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectDebitChannelPropertiesBankAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


