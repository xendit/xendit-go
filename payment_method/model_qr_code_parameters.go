/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.128.0
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
)

// checks if the QRCodeParameters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &QRCodeParameters{}

// QRCodeParameters struct for QRCodeParameters
type QRCodeParameters struct {
	Amount NullableFloat64 `json:"amount,omitempty"`
	Currency *string `json:"currency,omitempty"`
	ChannelCode NullableQRCodeChannelCode `json:"channel_code,omitempty"`
	ChannelProperties NullableQRCodeChannelProperties `json:"channel_properties,omitempty"`
}

// NewQRCodeParameters instantiates a new QRCodeParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQRCodeParameters() *QRCodeParameters {
	this := QRCodeParameters{}
	return &this
}

// NewQRCodeParametersWithDefaults instantiates a new QRCodeParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQRCodeParametersWithDefaults() *QRCodeParameters {
	this := QRCodeParameters{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QRCodeParameters) GetAmount() float64 {
	if o == nil || utils.IsNil(o.Amount.Get()) {
		var ret float64
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QRCodeParameters) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *QRCodeParameters) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat64 and assigns it to the Amount field.
func (o *QRCodeParameters) SetAmount(v float64) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *QRCodeParameters) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *QRCodeParameters) UnsetAmount() {
	o.Amount.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *QRCodeParameters) GetCurrency() string {
	if o == nil || utils.IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QRCodeParameters) GetCurrencyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *QRCodeParameters) HasCurrency() bool {
	if o != nil && !utils.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *QRCodeParameters) SetCurrency(v string) {
	o.Currency = &v
}

// GetChannelCode returns the ChannelCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QRCodeParameters) GetChannelCode() QRCodeChannelCode {
	if o == nil || utils.IsNil(o.ChannelCode.Get()) {
		var ret QRCodeChannelCode
		return ret
	}
	return *o.ChannelCode.Get()
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QRCodeParameters) GetChannelCodeOk() (*QRCodeChannelCode, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelCode.Get(), o.ChannelCode.IsSet()
}

// HasChannelCode returns a boolean if a field has been set.
func (o *QRCodeParameters) HasChannelCode() bool {
	if o != nil && o.ChannelCode.IsSet() {
		return true
	}

	return false
}

// SetChannelCode gets a reference to the given NullableQRCodeChannelCode and assigns it to the ChannelCode field.
func (o *QRCodeParameters) SetChannelCode(v QRCodeChannelCode) {
	o.ChannelCode.Set(&v)
}
// SetChannelCodeNil sets the value for ChannelCode to be an explicit nil
func (o *QRCodeParameters) SetChannelCodeNil() {
	o.ChannelCode.Set(nil)
}

// UnsetChannelCode ensures that no value is present for ChannelCode, not even an explicit nil
func (o *QRCodeParameters) UnsetChannelCode() {
	o.ChannelCode.Unset()
}

// GetChannelProperties returns the ChannelProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QRCodeParameters) GetChannelProperties() QRCodeChannelProperties {
	if o == nil || utils.IsNil(o.ChannelProperties.Get()) {
		var ret QRCodeChannelProperties
		return ret
	}
	return *o.ChannelProperties.Get()
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QRCodeParameters) GetChannelPropertiesOk() (*QRCodeChannelProperties, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelProperties.Get(), o.ChannelProperties.IsSet()
}

// HasChannelProperties returns a boolean if a field has been set.
func (o *QRCodeParameters) HasChannelProperties() bool {
	if o != nil && o.ChannelProperties.IsSet() {
		return true
	}

	return false
}

// SetChannelProperties gets a reference to the given NullableQRCodeChannelProperties and assigns it to the ChannelProperties field.
func (o *QRCodeParameters) SetChannelProperties(v QRCodeChannelProperties) {
	o.ChannelProperties.Set(&v)
}
// SetChannelPropertiesNil sets the value for ChannelProperties to be an explicit nil
func (o *QRCodeParameters) SetChannelPropertiesNil() {
	o.ChannelProperties.Set(nil)
}

// UnsetChannelProperties ensures that no value is present for ChannelProperties, not even an explicit nil
func (o *QRCodeParameters) UnsetChannelProperties() {
	o.ChannelProperties.Unset()
}

func (o QRCodeParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QRCodeParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
    }
	if !utils.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if o.ChannelCode.IsSet() {
		toSerialize["channel_code"] = o.ChannelCode.Get()
    }
	if o.ChannelProperties.IsSet() {
		toSerialize["channel_properties"] = o.ChannelProperties.Get()
    }
	return toSerialize, nil
}

type NullableQRCodeParameters struct {
	value *QRCodeParameters
	isSet bool
}

func (v NullableQRCodeParameters) Get() *QRCodeParameters {
	return v.value
}

func (v *NullableQRCodeParameters) Set(val *QRCodeParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableQRCodeParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableQRCodeParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQRCodeParameters(val *QRCodeParameters) *NullableQRCodeParameters {
	return &NullableQRCodeParameters{value: val, isSet: true}
}

func (v NullableQRCodeParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQRCodeParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


