/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.99.0
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v5/utils"
)

// checks if the VirtualAccount type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &VirtualAccount{}

// VirtualAccount Virtual Account Payment Method Details
type VirtualAccount struct {
	Amount NullableFloat64 `json:"amount,omitempty"`
	MinAmount NullableFloat64 `json:"min_amount,omitempty"`
	MaxAmount NullableFloat64 `json:"max_amount,omitempty"`
	Currency *string `json:"currency,omitempty"`
	ChannelCode VirtualAccountChannelCode `json:"channel_code"`
	ChannelProperties VirtualAccountChannelProperties `json:"channel_properties"`
	// For payments in Vietnam only, alternative display requested for the virtual account
	AlternativeDisplayTypes []string `json:"alternative_display_types,omitempty"`
	AlternativeDisplays []VirtualAccountAlternativeDisplay `json:"alternative_displays,omitempty"`
}

// NewVirtualAccount instantiates a new VirtualAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualAccount(channelCode VirtualAccountChannelCode, channelProperties VirtualAccountChannelProperties) *VirtualAccount {
	this := VirtualAccount{}
	this.ChannelCode = channelCode
	this.ChannelProperties = channelProperties
	return &this
}

// NewVirtualAccountWithDefaults instantiates a new VirtualAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualAccountWithDefaults() *VirtualAccount {
	this := VirtualAccount{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualAccount) GetAmount() float64 {
	if o == nil || utils.IsNil(o.Amount.Get()) {
		var ret float64
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualAccount) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *VirtualAccount) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat64 and assigns it to the Amount field.
func (o *VirtualAccount) SetAmount(v float64) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *VirtualAccount) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *VirtualAccount) UnsetAmount() {
	o.Amount.Unset()
}

// GetMinAmount returns the MinAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualAccount) GetMinAmount() float64 {
	if o == nil || utils.IsNil(o.MinAmount.Get()) {
		var ret float64
		return ret
	}
	return *o.MinAmount.Get()
}

// GetMinAmountOk returns a tuple with the MinAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualAccount) GetMinAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinAmount.Get(), o.MinAmount.IsSet()
}

// HasMinAmount returns a boolean if a field has been set.
func (o *VirtualAccount) HasMinAmount() bool {
	if o != nil && o.MinAmount.IsSet() {
		return true
	}

	return false
}

// SetMinAmount gets a reference to the given NullableFloat64 and assigns it to the MinAmount field.
func (o *VirtualAccount) SetMinAmount(v float64) {
	o.MinAmount.Set(&v)
}
// SetMinAmountNil sets the value for MinAmount to be an explicit nil
func (o *VirtualAccount) SetMinAmountNil() {
	o.MinAmount.Set(nil)
}

// UnsetMinAmount ensures that no value is present for MinAmount, not even an explicit nil
func (o *VirtualAccount) UnsetMinAmount() {
	o.MinAmount.Unset()
}

// GetMaxAmount returns the MaxAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualAccount) GetMaxAmount() float64 {
	if o == nil || utils.IsNil(o.MaxAmount.Get()) {
		var ret float64
		return ret
	}
	return *o.MaxAmount.Get()
}

// GetMaxAmountOk returns a tuple with the MaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualAccount) GetMaxAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxAmount.Get(), o.MaxAmount.IsSet()
}

// HasMaxAmount returns a boolean if a field has been set.
func (o *VirtualAccount) HasMaxAmount() bool {
	if o != nil && o.MaxAmount.IsSet() {
		return true
	}

	return false
}

// SetMaxAmount gets a reference to the given NullableFloat64 and assigns it to the MaxAmount field.
func (o *VirtualAccount) SetMaxAmount(v float64) {
	o.MaxAmount.Set(&v)
}
// SetMaxAmountNil sets the value for MaxAmount to be an explicit nil
func (o *VirtualAccount) SetMaxAmountNil() {
	o.MaxAmount.Set(nil)
}

// UnsetMaxAmount ensures that no value is present for MaxAmount, not even an explicit nil
func (o *VirtualAccount) UnsetMaxAmount() {
	o.MaxAmount.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *VirtualAccount) GetCurrency() string {
	if o == nil || utils.IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualAccount) GetCurrencyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *VirtualAccount) HasCurrency() bool {
	if o != nil && !utils.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *VirtualAccount) SetCurrency(v string) {
	o.Currency = &v
}

// GetChannelCode returns the ChannelCode field value
func (o *VirtualAccount) GetChannelCode() VirtualAccountChannelCode {
	if o == nil {
		var ret VirtualAccountChannelCode
		return ret
	}

	return o.ChannelCode
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value
// and a boolean to check if the value has been set.
func (o *VirtualAccount) GetChannelCodeOk() (*VirtualAccountChannelCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelCode, true
}

// SetChannelCode sets field value
func (o *VirtualAccount) SetChannelCode(v VirtualAccountChannelCode) {
	o.ChannelCode = v
}

// GetChannelProperties returns the ChannelProperties field value
func (o *VirtualAccount) GetChannelProperties() VirtualAccountChannelProperties {
	if o == nil {
		var ret VirtualAccountChannelProperties
		return ret
	}

	return o.ChannelProperties
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value
// and a boolean to check if the value has been set.
func (o *VirtualAccount) GetChannelPropertiesOk() (*VirtualAccountChannelProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelProperties, true
}

// SetChannelProperties sets field value
func (o *VirtualAccount) SetChannelProperties(v VirtualAccountChannelProperties) {
	o.ChannelProperties = v
}

// GetAlternativeDisplayTypes returns the AlternativeDisplayTypes field value if set, zero value otherwise.
func (o *VirtualAccount) GetAlternativeDisplayTypes() []string {
	if o == nil || utils.IsNil(o.AlternativeDisplayTypes) {
		var ret []string
		return ret
	}
	return o.AlternativeDisplayTypes
}

// GetAlternativeDisplayTypesOk returns a tuple with the AlternativeDisplayTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualAccount) GetAlternativeDisplayTypesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.AlternativeDisplayTypes) {
		return nil, false
	}
	return o.AlternativeDisplayTypes, true
}

// HasAlternativeDisplayTypes returns a boolean if a field has been set.
func (o *VirtualAccount) HasAlternativeDisplayTypes() bool {
	if o != nil && !utils.IsNil(o.AlternativeDisplayTypes) {
		return true
	}

	return false
}

// SetAlternativeDisplayTypes gets a reference to the given []string and assigns it to the AlternativeDisplayTypes field.
func (o *VirtualAccount) SetAlternativeDisplayTypes(v []string) {
	o.AlternativeDisplayTypes = v
}

// GetAlternativeDisplays returns the AlternativeDisplays field value if set, zero value otherwise.
func (o *VirtualAccount) GetAlternativeDisplays() []VirtualAccountAlternativeDisplay {
	if o == nil || utils.IsNil(o.AlternativeDisplays) {
		var ret []VirtualAccountAlternativeDisplay
		return ret
	}
	return o.AlternativeDisplays
}

// GetAlternativeDisplaysOk returns a tuple with the AlternativeDisplays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualAccount) GetAlternativeDisplaysOk() ([]VirtualAccountAlternativeDisplay, bool) {
	if o == nil || utils.IsNil(o.AlternativeDisplays) {
		return nil, false
	}
	return o.AlternativeDisplays, true
}

// HasAlternativeDisplays returns a boolean if a field has been set.
func (o *VirtualAccount) HasAlternativeDisplays() bool {
	if o != nil && !utils.IsNil(o.AlternativeDisplays) {
		return true
	}

	return false
}

// SetAlternativeDisplays gets a reference to the given []VirtualAccountAlternativeDisplay and assigns it to the AlternativeDisplays field.
func (o *VirtualAccount) SetAlternativeDisplays(v []VirtualAccountAlternativeDisplay) {
	o.AlternativeDisplays = v
}

func (o VirtualAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
    }
	if o.MinAmount.IsSet() {
		toSerialize["min_amount"] = o.MinAmount.Get()
    }
	if o.MaxAmount.IsSet() {
		toSerialize["max_amount"] = o.MaxAmount.Get()
    }
	if !utils.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	toSerialize["channel_code"] = o.ChannelCode
	toSerialize["channel_properties"] = o.ChannelProperties
	if !utils.IsNil(o.AlternativeDisplayTypes) {
		toSerialize["alternative_display_types"] = o.AlternativeDisplayTypes
	}
	if !utils.IsNil(o.AlternativeDisplays) {
		toSerialize["alternative_displays"] = o.AlternativeDisplays
	}
	return toSerialize, nil
}

type NullableVirtualAccount struct {
	value *VirtualAccount
	isSet bool
}

func (v NullableVirtualAccount) Get() *VirtualAccount {
	return v.value
}

func (v *NullableVirtualAccount) Set(val *VirtualAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualAccount(val *VirtualAccount) *NullableVirtualAccount {
	return &NullableVirtualAccount{value: val, isSet: true}
}

func (v NullableVirtualAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


