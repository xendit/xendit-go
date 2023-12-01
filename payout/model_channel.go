/*
Payout Service

This API allows Xendit to send money from an account to a channel (banks, eWallets, retail outlets) from across regions

API version: 1.0.0
*/


package payout

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v4/utils"
)

// checks if the Channel type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Channel{}

// Channel Channel information where you can send the money to
type Channel struct {
	// Destination channel to send the money to, prefixed by ISO-3166 country code
	ChannelCode string `json:"channel_code"`
	ChannelCategory ChannelCategory `json:"channel_category"`
	// Currency of the destination channel using ISO-4217 currency code
	Currency string `json:"currency"`
	// Name of the destination channel
	ChannelName string `json:"channel_name"`
	AmountLimits ChannelAmountLimits `json:"amount_limits"`
}

// NewChannel instantiates a new Channel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannel(channelCode string, channelCategory ChannelCategory, currency string, channelName string, amountLimits ChannelAmountLimits) *Channel {
	this := Channel{}
	this.ChannelCode = channelCode
	this.ChannelCategory = channelCategory
	this.Currency = currency
	this.ChannelName = channelName
	this.AmountLimits = amountLimits
	return &this
}

// NewChannelWithDefaults instantiates a new Channel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelWithDefaults() *Channel {
	this := Channel{}
	return &this
}

// GetChannelCode returns the ChannelCode field value
func (o *Channel) GetChannelCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelCode
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value
// and a boolean to check if the value has been set.
func (o *Channel) GetChannelCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelCode, true
}

// SetChannelCode sets field value
func (o *Channel) SetChannelCode(v string) {
	o.ChannelCode = v
}

// GetChannelCategory returns the ChannelCategory field value
func (o *Channel) GetChannelCategory() ChannelCategory {
	if o == nil {
		var ret ChannelCategory
		return ret
	}

	return o.ChannelCategory
}

// GetChannelCategoryOk returns a tuple with the ChannelCategory field value
// and a boolean to check if the value has been set.
func (o *Channel) GetChannelCategoryOk() (*ChannelCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelCategory, true
}

// SetChannelCategory sets field value
func (o *Channel) SetChannelCategory(v ChannelCategory) {
	o.ChannelCategory = v
}

// GetCurrency returns the Currency field value
func (o *Channel) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Channel) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Channel) SetCurrency(v string) {
	o.Currency = v
}

// GetChannelName returns the ChannelName field value
func (o *Channel) GetChannelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value
// and a boolean to check if the value has been set.
func (o *Channel) GetChannelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelName, true
}

// SetChannelName sets field value
func (o *Channel) SetChannelName(v string) {
	o.ChannelName = v
}

// GetAmountLimits returns the AmountLimits field value
func (o *Channel) GetAmountLimits() ChannelAmountLimits {
	if o == nil {
		var ret ChannelAmountLimits
		return ret
	}

	return o.AmountLimits
}

// GetAmountLimitsOk returns a tuple with the AmountLimits field value
// and a boolean to check if the value has been set.
func (o *Channel) GetAmountLimitsOk() (*ChannelAmountLimits, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountLimits, true
}

// SetAmountLimits sets field value
func (o *Channel) SetAmountLimits(v ChannelAmountLimits) {
	o.AmountLimits = v
}

func (o Channel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Channel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel_code"] = o.ChannelCode
	toSerialize["channel_category"] = o.ChannelCategory
	toSerialize["currency"] = o.Currency
	toSerialize["channel_name"] = o.ChannelName
	toSerialize["amount_limits"] = o.AmountLimits
	return toSerialize, nil
}

type NullableChannel struct {
	value *Channel
	isSet bool
}

func (v NullableChannel) Get() *Channel {
	return v.value
}

func (v *NullableChannel) Set(val *Channel) {
	v.value = val
	v.isSet = true
}

func (v NullableChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannel(val *Channel) *NullableChannel {
	return &NullableChannel{value: val, isSet: true}
}

func (v NullableChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


