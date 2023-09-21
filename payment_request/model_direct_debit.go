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

// checks if the DirectDebit type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DirectDebit{}

// DirectDebit Direct Debit Payment Method Details
type DirectDebit struct {
	ChannelCode DirectDebitChannelCode `json:"channel_code"`
	ChannelProperties NullableDirectDebitChannelProperties `json:"channel_properties"`
	Type DirectDebitType `json:"type"`
	BankAccount NullableDirectDebitBankAccount `json:"bank_account,omitempty"`
	DebitCard NullableDirectDebitDebitCard `json:"debit_card,omitempty"`
}

// NewDirectDebit instantiates a new DirectDebit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectDebit(channelCode DirectDebitChannelCode, channelProperties NullableDirectDebitChannelProperties, type_ DirectDebitType) *DirectDebit {
	this := DirectDebit{}
	this.ChannelCode = channelCode
	this.ChannelProperties = channelProperties
	this.Type = type_
	return &this
}

// NewDirectDebitWithDefaults instantiates a new DirectDebit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectDebitWithDefaults() *DirectDebit {
	this := DirectDebit{}
	return &this
}

// GetChannelCode returns the ChannelCode field value
func (o *DirectDebit) GetChannelCode() DirectDebitChannelCode {
	if o == nil {
		var ret DirectDebitChannelCode
		return ret
	}

	return o.ChannelCode
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value
// and a boolean to check if the value has been set.
func (o *DirectDebit) GetChannelCodeOk() (*DirectDebitChannelCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelCode, true
}

// SetChannelCode sets field value
func (o *DirectDebit) SetChannelCode(v DirectDebitChannelCode) {
	o.ChannelCode = v
}

// GetChannelProperties returns the ChannelProperties field value
// If the value is explicit nil, the zero value for DirectDebitChannelProperties will be returned
func (o *DirectDebit) GetChannelProperties() DirectDebitChannelProperties {
	if o == nil || o.ChannelProperties.Get() == nil {
		var ret DirectDebitChannelProperties
		return ret
	}

	return *o.ChannelProperties.Get()
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectDebit) GetChannelPropertiesOk() (*DirectDebitChannelProperties, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelProperties.Get(), o.ChannelProperties.IsSet()
}

// SetChannelProperties sets field value
func (o *DirectDebit) SetChannelProperties(v DirectDebitChannelProperties) {
	o.ChannelProperties.Set(&v)
}

// GetType returns the Type field value
func (o *DirectDebit) GetType() DirectDebitType {
	if o == nil {
		var ret DirectDebitType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DirectDebit) GetTypeOk() (*DirectDebitType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DirectDebit) SetType(v DirectDebitType) {
	o.Type = v
}

// GetBankAccount returns the BankAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectDebit) GetBankAccount() DirectDebitBankAccount {
	if o == nil || utils.IsNil(o.BankAccount.Get()) {
		var ret DirectDebitBankAccount
		return ret
	}
	return *o.BankAccount.Get()
}

// GetBankAccountOk returns a tuple with the BankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectDebit) GetBankAccountOk() (*DirectDebitBankAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankAccount.Get(), o.BankAccount.IsSet()
}

// HasBankAccount returns a boolean if a field has been set.
func (o *DirectDebit) HasBankAccount() bool {
	if o != nil && o.BankAccount.IsSet() {
		return true
	}

	return false
}

// SetBankAccount gets a reference to the given NullableDirectDebitBankAccount and assigns it to the BankAccount field.
func (o *DirectDebit) SetBankAccount(v DirectDebitBankAccount) {
	o.BankAccount.Set(&v)
}
// SetBankAccountNil sets the value for BankAccount to be an explicit nil
func (o *DirectDebit) SetBankAccountNil() {
	o.BankAccount.Set(nil)
}

// UnsetBankAccount ensures that no value is present for BankAccount, not even an explicit nil
func (o *DirectDebit) UnsetBankAccount() {
	o.BankAccount.Unset()
}

// GetDebitCard returns the DebitCard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectDebit) GetDebitCard() DirectDebitDebitCard {
	if o == nil || utils.IsNil(o.DebitCard.Get()) {
		var ret DirectDebitDebitCard
		return ret
	}
	return *o.DebitCard.Get()
}

// GetDebitCardOk returns a tuple with the DebitCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectDebit) GetDebitCardOk() (*DirectDebitDebitCard, bool) {
	if o == nil {
		return nil, false
	}
	return o.DebitCard.Get(), o.DebitCard.IsSet()
}

// HasDebitCard returns a boolean if a field has been set.
func (o *DirectDebit) HasDebitCard() bool {
	if o != nil && o.DebitCard.IsSet() {
		return true
	}

	return false
}

// SetDebitCard gets a reference to the given NullableDirectDebitDebitCard and assigns it to the DebitCard field.
func (o *DirectDebit) SetDebitCard(v DirectDebitDebitCard) {
	o.DebitCard.Set(&v)
}
// SetDebitCardNil sets the value for DebitCard to be an explicit nil
func (o *DirectDebit) SetDebitCardNil() {
	o.DebitCard.Set(nil)
}

// UnsetDebitCard ensures that no value is present for DebitCard, not even an explicit nil
func (o *DirectDebit) UnsetDebitCard() {
	o.DebitCard.Unset()
}

func (o DirectDebit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectDebit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel_code"] = o.ChannelCode
	toSerialize["channel_properties"] = o.ChannelProperties.Get()
	toSerialize["type"] = o.Type
	if o.BankAccount.IsSet() {
		toSerialize["bank_account"] = o.BankAccount.Get()
	}
	if o.DebitCard.IsSet() {
		toSerialize["debit_card"] = o.DebitCard.Get()
	}
	return toSerialize, nil
}

type NullableDirectDebit struct {
	value *DirectDebit
	isSet bool
}

func (v NullableDirectDebit) Get() *DirectDebit {
	return v.value
}

func (v *NullableDirectDebit) Set(val *DirectDebit) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectDebit) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectDebit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectDebit(val *DirectDebit) *NullableDirectDebit {
	return &NullableDirectDebit{value: val, isSet: true}
}

func (v NullableDirectDebit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectDebit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


