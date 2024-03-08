/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.7.6
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v4/utils"
)

// checks if the Bank type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Bank{}

// Bank An object representing bank details for invoices.
type Bank struct {
	BankCode BankCode `json:"bank_code"`
	// The collection type for the bank details.
	CollectionType string `json:"collection_type"`
	// The branch of the bank.
	BankBranch *string `json:"bank_branch,omitempty"`
	// The bank account number.
	BankAccountNumber *string `json:"bank_account_number,omitempty"`
	// The name of the account holder.
	AccountHolderName string `json:"account_holder_name"`
	// The transfer amount.
	TransferAmount *float64 `json:"transfer_amount,omitempty"`
	AlternativeDisplays []AlternativeDisplayItem `json:"alternative_displays,omitempty"`
}

// NewBank instantiates a new Bank object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBank(bankCode BankCode, collectionType string, accountHolderName string) *Bank {
	this := Bank{}
	this.BankCode = bankCode
	this.CollectionType = collectionType
	this.AccountHolderName = accountHolderName
	return &this
}

// NewBankWithDefaults instantiates a new Bank object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankWithDefaults() *Bank {
	this := Bank{}
	return &this
}

// GetBankCode returns the BankCode field value
func (o *Bank) GetBankCode() BankCode {
	if o == nil {
		var ret BankCode
		return ret
	}

	return o.BankCode
}

// GetBankCodeOk returns a tuple with the BankCode field value
// and a boolean to check if the value has been set.
func (o *Bank) GetBankCodeOk() (*BankCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankCode, true
}

// SetBankCode sets field value
func (o *Bank) SetBankCode(v BankCode) {
	o.BankCode = v
}

// GetCollectionType returns the CollectionType field value
func (o *Bank) GetCollectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionType
}

// GetCollectionTypeOk returns a tuple with the CollectionType field value
// and a boolean to check if the value has been set.
func (o *Bank) GetCollectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionType, true
}

// SetCollectionType sets field value
func (o *Bank) SetCollectionType(v string) {
	o.CollectionType = v
}

// GetBankBranch returns the BankBranch field value if set, zero value otherwise.
func (o *Bank) GetBankBranch() string {
	if o == nil || utils.IsNil(o.BankBranch) {
		var ret string
		return ret
	}
	return *o.BankBranch
}

// GetBankBranchOk returns a tuple with the BankBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bank) GetBankBranchOk() (*string, bool) {
	if o == nil || utils.IsNil(o.BankBranch) {
		return nil, false
	}
	return o.BankBranch, true
}

// HasBankBranch returns a boolean if a field has been set.
func (o *Bank) HasBankBranch() bool {
	if o != nil && !utils.IsNil(o.BankBranch) {
		return true
	}

	return false
}

// SetBankBranch gets a reference to the given string and assigns it to the BankBranch field.
func (o *Bank) SetBankBranch(v string) {
	o.BankBranch = &v
}

// GetBankAccountNumber returns the BankAccountNumber field value if set, zero value otherwise.
func (o *Bank) GetBankAccountNumber() string {
	if o == nil || utils.IsNil(o.BankAccountNumber) {
		var ret string
		return ret
	}
	return *o.BankAccountNumber
}

// GetBankAccountNumberOk returns a tuple with the BankAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bank) GetBankAccountNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.BankAccountNumber) {
		return nil, false
	}
	return o.BankAccountNumber, true
}

// HasBankAccountNumber returns a boolean if a field has been set.
func (o *Bank) HasBankAccountNumber() bool {
	if o != nil && !utils.IsNil(o.BankAccountNumber) {
		return true
	}

	return false
}

// SetBankAccountNumber gets a reference to the given string and assigns it to the BankAccountNumber field.
func (o *Bank) SetBankAccountNumber(v string) {
	o.BankAccountNumber = &v
}

// GetAccountHolderName returns the AccountHolderName field value
func (o *Bank) GetAccountHolderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountHolderName
}

// GetAccountHolderNameOk returns a tuple with the AccountHolderName field value
// and a boolean to check if the value has been set.
func (o *Bank) GetAccountHolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountHolderName, true
}

// SetAccountHolderName sets field value
func (o *Bank) SetAccountHolderName(v string) {
	o.AccountHolderName = v
}

// GetTransferAmount returns the TransferAmount field value if set, zero value otherwise.
func (o *Bank) GetTransferAmount() float64 {
	if o == nil || utils.IsNil(o.TransferAmount) {
		var ret float64
		return ret
	}
	return *o.TransferAmount
}

// GetTransferAmountOk returns a tuple with the TransferAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bank) GetTransferAmountOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.TransferAmount) {
		return nil, false
	}
	return o.TransferAmount, true
}

// HasTransferAmount returns a boolean if a field has been set.
func (o *Bank) HasTransferAmount() bool {
	if o != nil && !utils.IsNil(o.TransferAmount) {
		return true
	}

	return false
}

// SetTransferAmount gets a reference to the given float64 and assigns it to the TransferAmount field.
func (o *Bank) SetTransferAmount(v float64) {
	o.TransferAmount = &v
}

// GetAlternativeDisplays returns the AlternativeDisplays field value if set, zero value otherwise.
func (o *Bank) GetAlternativeDisplays() []AlternativeDisplayItem {
	if o == nil || utils.IsNil(o.AlternativeDisplays) {
		var ret []AlternativeDisplayItem
		return ret
	}
	return o.AlternativeDisplays
}

// GetAlternativeDisplaysOk returns a tuple with the AlternativeDisplays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bank) GetAlternativeDisplaysOk() ([]AlternativeDisplayItem, bool) {
	if o == nil || utils.IsNil(o.AlternativeDisplays) {
		return nil, false
	}
	return o.AlternativeDisplays, true
}

// HasAlternativeDisplays returns a boolean if a field has been set.
func (o *Bank) HasAlternativeDisplays() bool {
	if o != nil && !utils.IsNil(o.AlternativeDisplays) {
		return true
	}

	return false
}

// SetAlternativeDisplays gets a reference to the given []AlternativeDisplayItem and assigns it to the AlternativeDisplays field.
func (o *Bank) SetAlternativeDisplays(v []AlternativeDisplayItem) {
	o.AlternativeDisplays = v
}

func (o Bank) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Bank) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bank_code"] = o.BankCode
	toSerialize["collection_type"] = o.CollectionType
	if !utils.IsNil(o.BankBranch) {
		toSerialize["bank_branch"] = o.BankBranch
	}
	if !utils.IsNil(o.BankAccountNumber) {
		toSerialize["bank_account_number"] = o.BankAccountNumber
	}
	toSerialize["account_holder_name"] = o.AccountHolderName
	if !utils.IsNil(o.TransferAmount) {
		toSerialize["transfer_amount"] = o.TransferAmount
	}
	if !utils.IsNil(o.AlternativeDisplays) {
		toSerialize["alternative_displays"] = o.AlternativeDisplays
	}
	return toSerialize, nil
}

type NullableBank struct {
	value *Bank
	isSet bool
}

func (v NullableBank) Get() *Bank {
	return v.value
}

func (v *NullableBank) Set(val *Bank) {
	v.value = val
	v.isSet = true
}

func (v NullableBank) IsSet() bool {
	return v.isSet
}

func (v *NullableBank) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBank(val *Bank) *NullableBank {
	return &NullableBank{value: val, isSet: true}
}

func (v NullableBank) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBank) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


