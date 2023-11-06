/*
Payment Requests

This API is used for Payment Requests

API version: 1.45.1
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the OverTheCounterParameters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &OverTheCounterParameters{}

// OverTheCounterParameters struct for OverTheCounterParameters
type OverTheCounterParameters struct {
	Amount NullableFloat64 `json:"amount,omitempty"`
	Currency *PaymentRequestCurrency `json:"currency,omitempty"`
	ChannelCode OverTheCounterChannelCode `json:"channel_code"`
	ChannelProperties OverTheCounterChannelProperties `json:"channel_properties"`
}

// NewOverTheCounterParameters instantiates a new OverTheCounterParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverTheCounterParameters(channelCode OverTheCounterChannelCode, channelProperties OverTheCounterChannelProperties) *OverTheCounterParameters {
	this := OverTheCounterParameters{}
	this.ChannelCode = channelCode
	this.ChannelProperties = channelProperties
	return &this
}

// NewOverTheCounterParametersWithDefaults instantiates a new OverTheCounterParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverTheCounterParametersWithDefaults() *OverTheCounterParameters {
	this := OverTheCounterParameters{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OverTheCounterParameters) GetAmount() float64 {
	if o == nil || utils.IsNil(o.Amount.Get()) {
		var ret float64
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OverTheCounterParameters) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *OverTheCounterParameters) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat64 and assigns it to the Amount field.
func (o *OverTheCounterParameters) SetAmount(v float64) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *OverTheCounterParameters) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *OverTheCounterParameters) UnsetAmount() {
	o.Amount.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *OverTheCounterParameters) GetCurrency() PaymentRequestCurrency {
	if o == nil || utils.IsNil(o.Currency) {
		var ret PaymentRequestCurrency
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverTheCounterParameters) GetCurrencyOk() (*PaymentRequestCurrency, bool) {
	if o == nil || utils.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *OverTheCounterParameters) HasCurrency() bool {
	if o != nil && !utils.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given PaymentRequestCurrency and assigns it to the Currency field.
func (o *OverTheCounterParameters) SetCurrency(v PaymentRequestCurrency) {
	o.Currency = &v
}

// GetChannelCode returns the ChannelCode field value
func (o *OverTheCounterParameters) GetChannelCode() OverTheCounterChannelCode {
	if o == nil {
		var ret OverTheCounterChannelCode
		return ret
	}

	return o.ChannelCode
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value
// and a boolean to check if the value has been set.
func (o *OverTheCounterParameters) GetChannelCodeOk() (*OverTheCounterChannelCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelCode, true
}

// SetChannelCode sets field value
func (o *OverTheCounterParameters) SetChannelCode(v OverTheCounterChannelCode) {
	o.ChannelCode = v
}

// GetChannelProperties returns the ChannelProperties field value
func (o *OverTheCounterParameters) GetChannelProperties() OverTheCounterChannelProperties {
	if o == nil {
		var ret OverTheCounterChannelProperties
		return ret
	}

	return o.ChannelProperties
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value
// and a boolean to check if the value has been set.
func (o *OverTheCounterParameters) GetChannelPropertiesOk() (*OverTheCounterChannelProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelProperties, true
}

// SetChannelProperties sets field value
func (o *OverTheCounterParameters) SetChannelProperties(v OverTheCounterChannelProperties) {
	o.ChannelProperties = v
}

func (o OverTheCounterParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OverTheCounterParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
    }
	if !utils.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	toSerialize["channel_code"] = o.ChannelCode
	toSerialize["channel_properties"] = o.ChannelProperties
	return toSerialize, nil
}

type NullableOverTheCounterParameters struct {
	value *OverTheCounterParameters
	isSet bool
}

func (v NullableOverTheCounterParameters) Get() *OverTheCounterParameters {
	return v.value
}

func (v *NullableOverTheCounterParameters) Set(val *OverTheCounterParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableOverTheCounterParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableOverTheCounterParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverTheCounterParameters(val *OverTheCounterParameters) *NullableOverTheCounterParameters {
	return &NullableOverTheCounterParameters{value: val, isSet: true}
}

func (v NullableOverTheCounterParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverTheCounterParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


