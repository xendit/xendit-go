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

// checks if the DirectDebitBankAccount type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DirectDebitBankAccount{}

// DirectDebitBankAccount struct for DirectDebitBankAccount
type DirectDebitBankAccount struct {
	MaskedBankAccountNumber NullableString `json:"masked_bank_account_number,omitempty"`
	BankAccountHash NullableString `json:"bank_account_hash,omitempty"`
}

// NewDirectDebitBankAccount instantiates a new DirectDebitBankAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectDebitBankAccount() *DirectDebitBankAccount {
	this := DirectDebitBankAccount{}
	return &this
}

// NewDirectDebitBankAccountWithDefaults instantiates a new DirectDebitBankAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectDebitBankAccountWithDefaults() *DirectDebitBankAccount {
	this := DirectDebitBankAccount{}
	return &this
}

// GetMaskedBankAccountNumber returns the MaskedBankAccountNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectDebitBankAccount) GetMaskedBankAccountNumber() string {
	if o == nil || utils.IsNil(o.MaskedBankAccountNumber.Get()) {
		var ret string
		return ret
	}
	return *o.MaskedBankAccountNumber.Get()
}

// GetMaskedBankAccountNumberOk returns a tuple with the MaskedBankAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectDebitBankAccount) GetMaskedBankAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaskedBankAccountNumber.Get(), o.MaskedBankAccountNumber.IsSet()
}

// HasMaskedBankAccountNumber returns a boolean if a field has been set.
func (o *DirectDebitBankAccount) HasMaskedBankAccountNumber() bool {
	if o != nil && o.MaskedBankAccountNumber.IsSet() {
		return true
	}

	return false
}

// SetMaskedBankAccountNumber gets a reference to the given NullableString and assigns it to the MaskedBankAccountNumber field.
func (o *DirectDebitBankAccount) SetMaskedBankAccountNumber(v string) {
	o.MaskedBankAccountNumber.Set(&v)
}
// SetMaskedBankAccountNumberNil sets the value for MaskedBankAccountNumber to be an explicit nil
func (o *DirectDebitBankAccount) SetMaskedBankAccountNumberNil() {
	o.MaskedBankAccountNumber.Set(nil)
}

// UnsetMaskedBankAccountNumber ensures that no value is present for MaskedBankAccountNumber, not even an explicit nil
func (o *DirectDebitBankAccount) UnsetMaskedBankAccountNumber() {
	o.MaskedBankAccountNumber.Unset()
}

// GetBankAccountHash returns the BankAccountHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectDebitBankAccount) GetBankAccountHash() string {
	if o == nil || utils.IsNil(o.BankAccountHash.Get()) {
		var ret string
		return ret
	}
	return *o.BankAccountHash.Get()
}

// GetBankAccountHashOk returns a tuple with the BankAccountHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectDebitBankAccount) GetBankAccountHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankAccountHash.Get(), o.BankAccountHash.IsSet()
}

// HasBankAccountHash returns a boolean if a field has been set.
func (o *DirectDebitBankAccount) HasBankAccountHash() bool {
	if o != nil && o.BankAccountHash.IsSet() {
		return true
	}

	return false
}

// SetBankAccountHash gets a reference to the given NullableString and assigns it to the BankAccountHash field.
func (o *DirectDebitBankAccount) SetBankAccountHash(v string) {
	o.BankAccountHash.Set(&v)
}
// SetBankAccountHashNil sets the value for BankAccountHash to be an explicit nil
func (o *DirectDebitBankAccount) SetBankAccountHashNil() {
	o.BankAccountHash.Set(nil)
}

// UnsetBankAccountHash ensures that no value is present for BankAccountHash, not even an explicit nil
func (o *DirectDebitBankAccount) UnsetBankAccountHash() {
	o.BankAccountHash.Unset()
}

func (o DirectDebitBankAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectDebitBankAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MaskedBankAccountNumber.IsSet() {
		toSerialize["masked_bank_account_number"] = o.MaskedBankAccountNumber.Get()
    }
	if o.BankAccountHash.IsSet() {
		toSerialize["bank_account_hash"] = o.BankAccountHash.Get()
    }
	return toSerialize, nil
}

type NullableDirectDebitBankAccount struct {
	value *DirectDebitBankAccount
	isSet bool
}

func (v NullableDirectDebitBankAccount) Get() *DirectDebitBankAccount {
	return v.value
}

func (v *NullableDirectDebitBankAccount) Set(val *DirectDebitBankAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectDebitBankAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectDebitBankAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectDebitBankAccount(val *DirectDebitBankAccount) *NullableDirectDebitBankAccount {
	return &NullableDirectDebitBankAccount{value: val, isSet: true}
}

func (v NullableDirectDebitBankAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectDebitBankAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


