/*
Payout Service

This API allows Xendit to send money from an account to a channel (banks, eWallets, retail outlets) from across regions

API version: 1.0.0
*/


package payout

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v5/utils"
)

// checks if the DigitalPayoutChannelProperties type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DigitalPayoutChannelProperties{}

// DigitalPayoutChannelProperties Channel information for digital destinations (banks, e-wallets)
type DigitalPayoutChannelProperties struct {
	// Registered account name
	AccountHolderName NullableString `json:"account_holder_name,omitempty"`
	// Registered account number
	AccountNumber string `json:"account_number"`
	AccountType *ChannelAccountType `json:"account_type,omitempty"`
}

// NewDigitalPayoutChannelProperties instantiates a new DigitalPayoutChannelProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigitalPayoutChannelProperties(accountNumber string) *DigitalPayoutChannelProperties {
	this := DigitalPayoutChannelProperties{}
	this.AccountNumber = accountNumber
	return &this
}

// NewDigitalPayoutChannelPropertiesWithDefaults instantiates a new DigitalPayoutChannelProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigitalPayoutChannelPropertiesWithDefaults() *DigitalPayoutChannelProperties {
	this := DigitalPayoutChannelProperties{}
	return &this
}

// GetAccountHolderName returns the AccountHolderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigitalPayoutChannelProperties) GetAccountHolderName() string {
	if o == nil || utils.IsNil(o.AccountHolderName.Get()) {
		var ret string
		return ret
	}
	return *o.AccountHolderName.Get()
}

// GetAccountHolderNameOk returns a tuple with the AccountHolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigitalPayoutChannelProperties) GetAccountHolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountHolderName.Get(), o.AccountHolderName.IsSet()
}

// HasAccountHolderName returns a boolean if a field has been set.
func (o *DigitalPayoutChannelProperties) HasAccountHolderName() bool {
	if o != nil && o.AccountHolderName.IsSet() {
		return true
	}

	return false
}

// SetAccountHolderName gets a reference to the given NullableString and assigns it to the AccountHolderName field.
func (o *DigitalPayoutChannelProperties) SetAccountHolderName(v string) {
	o.AccountHolderName.Set(&v)
}
// SetAccountHolderNameNil sets the value for AccountHolderName to be an explicit nil
func (o *DigitalPayoutChannelProperties) SetAccountHolderNameNil() {
	o.AccountHolderName.Set(nil)
}

// UnsetAccountHolderName ensures that no value is present for AccountHolderName, not even an explicit nil
func (o *DigitalPayoutChannelProperties) UnsetAccountHolderName() {
	o.AccountHolderName.Unset()
}

// GetAccountNumber returns the AccountNumber field value
func (o *DigitalPayoutChannelProperties) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *DigitalPayoutChannelProperties) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *DigitalPayoutChannelProperties) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *DigitalPayoutChannelProperties) GetAccountType() ChannelAccountType {
	if o == nil || utils.IsNil(o.AccountType) {
		var ret ChannelAccountType
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalPayoutChannelProperties) GetAccountTypeOk() (*ChannelAccountType, bool) {
	if o == nil || utils.IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *DigitalPayoutChannelProperties) HasAccountType() bool {
	if o != nil && !utils.IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given ChannelAccountType and assigns it to the AccountType field.
func (o *DigitalPayoutChannelProperties) SetAccountType(v ChannelAccountType) {
	o.AccountType = &v
}

func (o DigitalPayoutChannelProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DigitalPayoutChannelProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountHolderName.IsSet() {
		toSerialize["account_holder_name"] = o.AccountHolderName.Get()
    }
	toSerialize["account_number"] = o.AccountNumber
	if !utils.IsNil(o.AccountType) {
		toSerialize["account_type"] = o.AccountType
	}
	return toSerialize, nil
}

type NullableDigitalPayoutChannelProperties struct {
	value *DigitalPayoutChannelProperties
	isSet bool
}

func (v NullableDigitalPayoutChannelProperties) Get() *DigitalPayoutChannelProperties {
	return v.value
}

func (v *NullableDigitalPayoutChannelProperties) Set(val *DigitalPayoutChannelProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableDigitalPayoutChannelProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableDigitalPayoutChannelProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigitalPayoutChannelProperties(val *DigitalPayoutChannelProperties) *NullableDigitalPayoutChannelProperties {
	return &NullableDigitalPayoutChannelProperties{value: val, isSet: true}
}

func (v NullableDigitalPayoutChannelProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigitalPayoutChannelProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


