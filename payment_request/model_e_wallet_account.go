/*
Payment Requests

This API is used for Payment Requests

API version: 1.44.0
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the EWalletAccount type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EWalletAccount{}

// EWalletAccount EWallet Account Properties
type EWalletAccount struct {
	// Name of the eWallet account holder. The value is null if unavailableName of the eWallet account holder. The value is null if unavailable
	Name NullableString `json:"name,omitempty"`
	// Identifier from eWallet provider e.g. phone number. The value is null if unavailable
	AccountDetails NullableString `json:"account_details,omitempty"`
	// The main balance amount on eWallet account provided from eWallet provider. The value is null if unavailable
	Balance NullableFloat64 `json:"balance,omitempty"`
	// The point balance amount on eWallet account. Applicable only on some eWallet provider that has point system. The value is null if unavailabl
	PointBalance NullableFloat64 `json:"point_balance,omitempty"`
}

// NewEWalletAccount instantiates a new EWalletAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEWalletAccount() *EWalletAccount {
	this := EWalletAccount{}
	return &this
}

// NewEWalletAccountWithDefaults instantiates a new EWalletAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEWalletAccountWithDefaults() *EWalletAccount {
	this := EWalletAccount{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EWalletAccount) GetName() string {
	if o == nil || utils.IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EWalletAccount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *EWalletAccount) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *EWalletAccount) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *EWalletAccount) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *EWalletAccount) UnsetName() {
	o.Name.Unset()
}

// GetAccountDetails returns the AccountDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EWalletAccount) GetAccountDetails() string {
	if o == nil || utils.IsNil(o.AccountDetails.Get()) {
		var ret string
		return ret
	}
	return *o.AccountDetails.Get()
}

// GetAccountDetailsOk returns a tuple with the AccountDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EWalletAccount) GetAccountDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountDetails.Get(), o.AccountDetails.IsSet()
}

// HasAccountDetails returns a boolean if a field has been set.
func (o *EWalletAccount) HasAccountDetails() bool {
	if o != nil && o.AccountDetails.IsSet() {
		return true
	}

	return false
}

// SetAccountDetails gets a reference to the given NullableString and assigns it to the AccountDetails field.
func (o *EWalletAccount) SetAccountDetails(v string) {
	o.AccountDetails.Set(&v)
}
// SetAccountDetailsNil sets the value for AccountDetails to be an explicit nil
func (o *EWalletAccount) SetAccountDetailsNil() {
	o.AccountDetails.Set(nil)
}

// UnsetAccountDetails ensures that no value is present for AccountDetails, not even an explicit nil
func (o *EWalletAccount) UnsetAccountDetails() {
	o.AccountDetails.Unset()
}

// GetBalance returns the Balance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EWalletAccount) GetBalance() float64 {
	if o == nil || utils.IsNil(o.Balance.Get()) {
		var ret float64
		return ret
	}
	return *o.Balance.Get()
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EWalletAccount) GetBalanceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Balance.Get(), o.Balance.IsSet()
}

// HasBalance returns a boolean if a field has been set.
func (o *EWalletAccount) HasBalance() bool {
	if o != nil && o.Balance.IsSet() {
		return true
	}

	return false
}

// SetBalance gets a reference to the given NullableFloat64 and assigns it to the Balance field.
func (o *EWalletAccount) SetBalance(v float64) {
	o.Balance.Set(&v)
}
// SetBalanceNil sets the value for Balance to be an explicit nil
func (o *EWalletAccount) SetBalanceNil() {
	o.Balance.Set(nil)
}

// UnsetBalance ensures that no value is present for Balance, not even an explicit nil
func (o *EWalletAccount) UnsetBalance() {
	o.Balance.Unset()
}

// GetPointBalance returns the PointBalance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EWalletAccount) GetPointBalance() float64 {
	if o == nil || utils.IsNil(o.PointBalance.Get()) {
		var ret float64
		return ret
	}
	return *o.PointBalance.Get()
}

// GetPointBalanceOk returns a tuple with the PointBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EWalletAccount) GetPointBalanceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PointBalance.Get(), o.PointBalance.IsSet()
}

// HasPointBalance returns a boolean if a field has been set.
func (o *EWalletAccount) HasPointBalance() bool {
	if o != nil && o.PointBalance.IsSet() {
		return true
	}

	return false
}

// SetPointBalance gets a reference to the given NullableFloat64 and assigns it to the PointBalance field.
func (o *EWalletAccount) SetPointBalance(v float64) {
	o.PointBalance.Set(&v)
}
// SetPointBalanceNil sets the value for PointBalance to be an explicit nil
func (o *EWalletAccount) SetPointBalanceNil() {
	o.PointBalance.Set(nil)
}

// UnsetPointBalance ensures that no value is present for PointBalance, not even an explicit nil
func (o *EWalletAccount) UnsetPointBalance() {
	o.PointBalance.Unset()
}

func (o EWalletAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EWalletAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.AccountDetails.IsSet() {
		toSerialize["account_details"] = o.AccountDetails.Get()
	}
	if o.Balance.IsSet() {
		toSerialize["balance"] = o.Balance.Get()
	}
	if o.PointBalance.IsSet() {
		toSerialize["point_balance"] = o.PointBalance.Get()
	}
	return toSerialize, nil
}

type NullableEWalletAccount struct {
	value *EWalletAccount
	isSet bool
}

func (v NullableEWalletAccount) Get() *EWalletAccount {
	return v.value
}

func (v *NullableEWalletAccount) Set(val *EWalletAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableEWalletAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableEWalletAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEWalletAccount(val *EWalletAccount) *NullableEWalletAccount {
	return &NullableEWalletAccount{value: val, isSet: true}
}

func (v NullableEWalletAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEWalletAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


