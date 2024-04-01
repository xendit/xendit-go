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

// checks if the AccountEwallet type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AccountEwallet{}

// AccountEwallet struct for AccountEwallet
type AccountEwallet struct {
	// Unique account identifier as per the bank records.
	AccountNumber *string `json:"account_number,omitempty"`
	// Name of account holder as per the bank records. Needs to match the registered account name exactly.
	AccountHolderName NullableString `json:"account_holder_name,omitempty"`
	Currency *string `json:"currency,omitempty"`
}

// NewAccountEwallet instantiates a new AccountEwallet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountEwallet() *AccountEwallet {
	this := AccountEwallet{}
	return &this
}

// NewAccountEwalletWithDefaults instantiates a new AccountEwallet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountEwalletWithDefaults() *AccountEwallet {
	this := AccountEwallet{}
	return &this
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *AccountEwallet) GetAccountNumber() string {
	if o == nil || utils.IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountEwallet) GetAccountNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *AccountEwallet) HasAccountNumber() bool {
	if o != nil && !utils.IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *AccountEwallet) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetAccountHolderName returns the AccountHolderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountEwallet) GetAccountHolderName() string {
	if o == nil || utils.IsNil(o.AccountHolderName.Get()) {
		var ret string
		return ret
	}
	return *o.AccountHolderName.Get()
}

// GetAccountHolderNameOk returns a tuple with the AccountHolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountEwallet) GetAccountHolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountHolderName.Get(), o.AccountHolderName.IsSet()
}

// HasAccountHolderName returns a boolean if a field has been set.
func (o *AccountEwallet) HasAccountHolderName() bool {
	if o != nil && o.AccountHolderName.IsSet() {
		return true
	}

	return false
}

// SetAccountHolderName gets a reference to the given NullableString and assigns it to the AccountHolderName field.
func (o *AccountEwallet) SetAccountHolderName(v string) {
	o.AccountHolderName.Set(&v)
}
// SetAccountHolderNameNil sets the value for AccountHolderName to be an explicit nil
func (o *AccountEwallet) SetAccountHolderNameNil() {
	o.AccountHolderName.Set(nil)
}

// UnsetAccountHolderName ensures that no value is present for AccountHolderName, not even an explicit nil
func (o *AccountEwallet) UnsetAccountHolderName() {
	o.AccountHolderName.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AccountEwallet) GetCurrency() string {
	if o == nil || utils.IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountEwallet) GetCurrencyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AccountEwallet) HasCurrency() bool {
	if o != nil && !utils.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AccountEwallet) SetCurrency(v string) {
	o.Currency = &v
}

func (o AccountEwallet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountEwallet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AccountNumber) {
		toSerialize["account_number"] = o.AccountNumber
	}
	if o.AccountHolderName.IsSet() {
		toSerialize["account_holder_name"] = o.AccountHolderName.Get()
    }
	if !utils.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

type NullableAccountEwallet struct {
	value *AccountEwallet
	isSet bool
}

func (v NullableAccountEwallet) Get() *AccountEwallet {
	return v.value
}

func (v *NullableAccountEwallet) Set(val *AccountEwallet) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountEwallet) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountEwallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountEwallet(val *AccountEwallet) *NullableAccountEwallet {
	return &NullableAccountEwallet{value: val, isSet: true}
}

func (v NullableAccountEwallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountEwallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


