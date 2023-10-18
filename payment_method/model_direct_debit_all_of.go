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

// checks if the DirectDebitAllOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DirectDebitAllOf{}

// DirectDebitAllOf struct for DirectDebitAllOf
type DirectDebitAllOf struct {
	Type DirectDebitType `json:"type"`
	BankAccount NullableDirectDebitBankAccount `json:"bank_account,omitempty"`
	DebitCard NullableDirectDebitDebitCard `json:"debit_card,omitempty"`
}

// NewDirectDebitAllOf instantiates a new DirectDebitAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectDebitAllOf(type_ DirectDebitType) *DirectDebitAllOf {
	this := DirectDebitAllOf{}
	this.Type = type_
	return &this
}

// NewDirectDebitAllOfWithDefaults instantiates a new DirectDebitAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectDebitAllOfWithDefaults() *DirectDebitAllOf {
	this := DirectDebitAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *DirectDebitAllOf) GetType() DirectDebitType {
	if o == nil {
		var ret DirectDebitType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DirectDebitAllOf) GetTypeOk() (*DirectDebitType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DirectDebitAllOf) SetType(v DirectDebitType) {
	o.Type = v
}

// GetBankAccount returns the BankAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectDebitAllOf) GetBankAccount() DirectDebitBankAccount {
	if o == nil || utils.IsNil(o.BankAccount.Get()) {
		var ret DirectDebitBankAccount
		return ret
	}
	return *o.BankAccount.Get()
}

// GetBankAccountOk returns a tuple with the BankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectDebitAllOf) GetBankAccountOk() (*DirectDebitBankAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankAccount.Get(), o.BankAccount.IsSet()
}

// HasBankAccount returns a boolean if a field has been set.
func (o *DirectDebitAllOf) HasBankAccount() bool {
	if o != nil && o.BankAccount.IsSet() {
		return true
	}

	return false
}

// SetBankAccount gets a reference to the given NullableDirectDebitBankAccount and assigns it to the BankAccount field.
func (o *DirectDebitAllOf) SetBankAccount(v DirectDebitBankAccount) {
	o.BankAccount.Set(&v)
}
// SetBankAccountNil sets the value for BankAccount to be an explicit nil
func (o *DirectDebitAllOf) SetBankAccountNil() {
	o.BankAccount.Set(nil)
}

// UnsetBankAccount ensures that no value is present for BankAccount, not even an explicit nil
func (o *DirectDebitAllOf) UnsetBankAccount() {
	o.BankAccount.Unset()
}

// GetDebitCard returns the DebitCard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectDebitAllOf) GetDebitCard() DirectDebitDebitCard {
	if o == nil || utils.IsNil(o.DebitCard.Get()) {
		var ret DirectDebitDebitCard
		return ret
	}
	return *o.DebitCard.Get()
}

// GetDebitCardOk returns a tuple with the DebitCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectDebitAllOf) GetDebitCardOk() (*DirectDebitDebitCard, bool) {
	if o == nil {
		return nil, false
	}
	return o.DebitCard.Get(), o.DebitCard.IsSet()
}

// HasDebitCard returns a boolean if a field has been set.
func (o *DirectDebitAllOf) HasDebitCard() bool {
	if o != nil && o.DebitCard.IsSet() {
		return true
	}

	return false
}

// SetDebitCard gets a reference to the given NullableDirectDebitDebitCard and assigns it to the DebitCard field.
func (o *DirectDebitAllOf) SetDebitCard(v DirectDebitDebitCard) {
	o.DebitCard.Set(&v)
}
// SetDebitCardNil sets the value for DebitCard to be an explicit nil
func (o *DirectDebitAllOf) SetDebitCardNil() {
	o.DebitCard.Set(nil)
}

// UnsetDebitCard ensures that no value is present for DebitCard, not even an explicit nil
func (o *DirectDebitAllOf) UnsetDebitCard() {
	o.DebitCard.Unset()
}

func (o DirectDebitAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectDebitAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.BankAccount.IsSet() {
		toSerialize["bank_account"] = o.BankAccount.Get()
    }
	if o.DebitCard.IsSet() {
		toSerialize["debit_card"] = o.DebitCard.Get()
    }
	return toSerialize, nil
}

type NullableDirectDebitAllOf struct {
	value *DirectDebitAllOf
	isSet bool
}

func (v NullableDirectDebitAllOf) Get() *DirectDebitAllOf {
	return v.value
}

func (v *NullableDirectDebitAllOf) Set(val *DirectDebitAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectDebitAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectDebitAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectDebitAllOf(val *DirectDebitAllOf) *NullableDirectDebitAllOf {
	return &NullableDirectDebitAllOf{value: val, isSet: true}
}

func (v NullableDirectDebitAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectDebitAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


