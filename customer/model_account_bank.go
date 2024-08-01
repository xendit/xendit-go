/*
XENDIT SDK openapi spec

XENDIT SDK openapi spec

API version: 1.0.0
*/


package customer

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
)

// checks if the AccountBank type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AccountBank{}

// AccountBank struct for AccountBank
type AccountBank struct {
	// Unique account identifier as per the bank records.
	AccountNumber *string `json:"account_number,omitempty"`
	// Name of account holder as per the bank records. Needs to match the registered account name exactly. .
	AccountHolderName NullableString `json:"account_holder_name,omitempty"`
	// The SWIFT code for international payments
	SwiftCode NullableString `json:"swift_code,omitempty"`
	// Free text account type, e.g., Savings, Transaction, Virtual Account.
	AccountType NullableString `json:"account_type,omitempty"`
	// Potentially masked account detail, for display purposes only.
	AccountDetails NullableString `json:"account_details,omitempty"`
	Currency *string `json:"currency,omitempty"`
}

// NewAccountBank instantiates a new AccountBank object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountBank() *AccountBank {
	this := AccountBank{}
	return &this
}

// NewAccountBankWithDefaults instantiates a new AccountBank object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountBankWithDefaults() *AccountBank {
	this := AccountBank{}
	return &this
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *AccountBank) GetAccountNumber() string {
	if o == nil || utils.IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBank) GetAccountNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *AccountBank) HasAccountNumber() bool {
	if o != nil && !utils.IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *AccountBank) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetAccountHolderName returns the AccountHolderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountBank) GetAccountHolderName() string {
	if o == nil || utils.IsNil(o.AccountHolderName.Get()) {
		var ret string
		return ret
	}
	return *o.AccountHolderName.Get()
}

// GetAccountHolderNameOk returns a tuple with the AccountHolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountBank) GetAccountHolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountHolderName.Get(), o.AccountHolderName.IsSet()
}

// HasAccountHolderName returns a boolean if a field has been set.
func (o *AccountBank) HasAccountHolderName() bool {
	if o != nil && o.AccountHolderName.IsSet() {
		return true
	}

	return false
}

// SetAccountHolderName gets a reference to the given NullableString and assigns it to the AccountHolderName field.
func (o *AccountBank) SetAccountHolderName(v string) {
	o.AccountHolderName.Set(&v)
}
// SetAccountHolderNameNil sets the value for AccountHolderName to be an explicit nil
func (o *AccountBank) SetAccountHolderNameNil() {
	o.AccountHolderName.Set(nil)
}

// UnsetAccountHolderName ensures that no value is present for AccountHolderName, not even an explicit nil
func (o *AccountBank) UnsetAccountHolderName() {
	o.AccountHolderName.Unset()
}

// GetSwiftCode returns the SwiftCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountBank) GetSwiftCode() string {
	if o == nil || utils.IsNil(o.SwiftCode.Get()) {
		var ret string
		return ret
	}
	return *o.SwiftCode.Get()
}

// GetSwiftCodeOk returns a tuple with the SwiftCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountBank) GetSwiftCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SwiftCode.Get(), o.SwiftCode.IsSet()
}

// HasSwiftCode returns a boolean if a field has been set.
func (o *AccountBank) HasSwiftCode() bool {
	if o != nil && o.SwiftCode.IsSet() {
		return true
	}

	return false
}

// SetSwiftCode gets a reference to the given NullableString and assigns it to the SwiftCode field.
func (o *AccountBank) SetSwiftCode(v string) {
	o.SwiftCode.Set(&v)
}
// SetSwiftCodeNil sets the value for SwiftCode to be an explicit nil
func (o *AccountBank) SetSwiftCodeNil() {
	o.SwiftCode.Set(nil)
}

// UnsetSwiftCode ensures that no value is present for SwiftCode, not even an explicit nil
func (o *AccountBank) UnsetSwiftCode() {
	o.SwiftCode.Unset()
}

// GetAccountType returns the AccountType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountBank) GetAccountType() string {
	if o == nil || utils.IsNil(o.AccountType.Get()) {
		var ret string
		return ret
	}
	return *o.AccountType.Get()
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountBank) GetAccountTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountType.Get(), o.AccountType.IsSet()
}

// HasAccountType returns a boolean if a field has been set.
func (o *AccountBank) HasAccountType() bool {
	if o != nil && o.AccountType.IsSet() {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given NullableString and assigns it to the AccountType field.
func (o *AccountBank) SetAccountType(v string) {
	o.AccountType.Set(&v)
}
// SetAccountTypeNil sets the value for AccountType to be an explicit nil
func (o *AccountBank) SetAccountTypeNil() {
	o.AccountType.Set(nil)
}

// UnsetAccountType ensures that no value is present for AccountType, not even an explicit nil
func (o *AccountBank) UnsetAccountType() {
	o.AccountType.Unset()
}

// GetAccountDetails returns the AccountDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountBank) GetAccountDetails() string {
	if o == nil || utils.IsNil(o.AccountDetails.Get()) {
		var ret string
		return ret
	}
	return *o.AccountDetails.Get()
}

// GetAccountDetailsOk returns a tuple with the AccountDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountBank) GetAccountDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountDetails.Get(), o.AccountDetails.IsSet()
}

// HasAccountDetails returns a boolean if a field has been set.
func (o *AccountBank) HasAccountDetails() bool {
	if o != nil && o.AccountDetails.IsSet() {
		return true
	}

	return false
}

// SetAccountDetails gets a reference to the given NullableString and assigns it to the AccountDetails field.
func (o *AccountBank) SetAccountDetails(v string) {
	o.AccountDetails.Set(&v)
}
// SetAccountDetailsNil sets the value for AccountDetails to be an explicit nil
func (o *AccountBank) SetAccountDetailsNil() {
	o.AccountDetails.Set(nil)
}

// UnsetAccountDetails ensures that no value is present for AccountDetails, not even an explicit nil
func (o *AccountBank) UnsetAccountDetails() {
	o.AccountDetails.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AccountBank) GetCurrency() string {
	if o == nil || utils.IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBank) GetCurrencyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AccountBank) HasCurrency() bool {
	if o != nil && !utils.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AccountBank) SetCurrency(v string) {
	o.Currency = &v
}

func (o AccountBank) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountBank) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AccountNumber) {
		toSerialize["account_number"] = o.AccountNumber
	}
	if o.AccountHolderName.IsSet() {
		toSerialize["account_holder_name"] = o.AccountHolderName.Get()
    }
	if o.SwiftCode.IsSet() {
		toSerialize["swift_code"] = o.SwiftCode.Get()
    }
	if o.AccountType.IsSet() {
		toSerialize["account_type"] = o.AccountType.Get()
    }
	if o.AccountDetails.IsSet() {
		toSerialize["account_details"] = o.AccountDetails.Get()
    }
	if !utils.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

type NullableAccountBank struct {
	value *AccountBank
	isSet bool
}

func (v NullableAccountBank) Get() *AccountBank {
	return v.value
}

func (v *NullableAccountBank) Set(val *AccountBank) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountBank) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountBank) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountBank(val *AccountBank) *NullableAccountBank {
	return &NullableAccountBank{value: val, isSet: true}
}

func (v NullableAccountBank) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountBank) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


