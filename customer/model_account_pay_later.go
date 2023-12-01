/*
XENDIT SDK openapi spec

XENDIT SDK openapi spec

API version: 1.0.0
*/


package customer

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v4/utils"
)

// checks if the AccountPayLater type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AccountPayLater{}

// AccountPayLater struct for AccountPayLater
type AccountPayLater struct {
	// Alphanumeric string identifying this account. Usually an email address or phone number.
	AccountId *string `json:"account_id,omitempty"`
	// Name of account holder as per the cardless credit account.
	AccountHolderName NullableString `json:"account_holder_name,omitempty"`
	Currency *string `json:"currency,omitempty"`
}

// NewAccountPayLater instantiates a new AccountPayLater object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountPayLater() *AccountPayLater {
	this := AccountPayLater{}
	return &this
}

// NewAccountPayLaterWithDefaults instantiates a new AccountPayLater object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountPayLaterWithDefaults() *AccountPayLater {
	this := AccountPayLater{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AccountPayLater) GetAccountId() string {
	if o == nil || utils.IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPayLater) GetAccountIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AccountPayLater) HasAccountId() bool {
	if o != nil && !utils.IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AccountPayLater) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAccountHolderName returns the AccountHolderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountPayLater) GetAccountHolderName() string {
	if o == nil || utils.IsNil(o.AccountHolderName.Get()) {
		var ret string
		return ret
	}
	return *o.AccountHolderName.Get()
}

// GetAccountHolderNameOk returns a tuple with the AccountHolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountPayLater) GetAccountHolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountHolderName.Get(), o.AccountHolderName.IsSet()
}

// HasAccountHolderName returns a boolean if a field has been set.
func (o *AccountPayLater) HasAccountHolderName() bool {
	if o != nil && o.AccountHolderName.IsSet() {
		return true
	}

	return false
}

// SetAccountHolderName gets a reference to the given NullableString and assigns it to the AccountHolderName field.
func (o *AccountPayLater) SetAccountHolderName(v string) {
	o.AccountHolderName.Set(&v)
}
// SetAccountHolderNameNil sets the value for AccountHolderName to be an explicit nil
func (o *AccountPayLater) SetAccountHolderNameNil() {
	o.AccountHolderName.Set(nil)
}

// UnsetAccountHolderName ensures that no value is present for AccountHolderName, not even an explicit nil
func (o *AccountPayLater) UnsetAccountHolderName() {
	o.AccountHolderName.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AccountPayLater) GetCurrency() string {
	if o == nil || utils.IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPayLater) GetCurrencyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AccountPayLater) HasCurrency() bool {
	if o != nil && !utils.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AccountPayLater) SetCurrency(v string) {
	o.Currency = &v
}

func (o AccountPayLater) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountPayLater) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if o.AccountHolderName.IsSet() {
		toSerialize["account_holder_name"] = o.AccountHolderName.Get()
    }
	if !utils.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

type NullableAccountPayLater struct {
	value *AccountPayLater
	isSet bool
}

func (v NullableAccountPayLater) Get() *AccountPayLater {
	return v.value
}

func (v *NullableAccountPayLater) Set(val *AccountPayLater) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountPayLater) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountPayLater) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountPayLater(val *AccountPayLater) *NullableAccountPayLater {
	return &NullableAccountPayLater{value: val, isSet: true}
}

func (v NullableAccountPayLater) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountPayLater) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


