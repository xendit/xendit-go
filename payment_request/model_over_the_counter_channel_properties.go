/*
Payment Requests

This API is used for Payment Requests

API version: 1.70.0
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
	"time"
)

// checks if the OverTheCounterChannelProperties type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &OverTheCounterChannelProperties{}

// OverTheCounterChannelProperties Over The Counter Channel Properties
type OverTheCounterChannelProperties struct {
	// The payment code that you want to assign, e.g 12345. If you do not send one, one will be picked at random.
	PaymentCode *string `json:"payment_code,omitempty"`
	// Name of customer.
	CustomerName string `json:"customer_name"`
	// The time when the payment code will be expired. The minimum is 2 hours and the maximum is 9 days for 7ELEVEN. Default expired date will be 2 days from payment code generated.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// NewOverTheCounterChannelProperties instantiates a new OverTheCounterChannelProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverTheCounterChannelProperties(customerName string) *OverTheCounterChannelProperties {
	this := OverTheCounterChannelProperties{}
	this.CustomerName = customerName
	return &this
}

// NewOverTheCounterChannelPropertiesWithDefaults instantiates a new OverTheCounterChannelProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverTheCounterChannelPropertiesWithDefaults() *OverTheCounterChannelProperties {
	this := OverTheCounterChannelProperties{}
	return &this
}

// GetPaymentCode returns the PaymentCode field value if set, zero value otherwise.
func (o *OverTheCounterChannelProperties) GetPaymentCode() string {
	if o == nil || utils.IsNil(o.PaymentCode) {
		var ret string
		return ret
	}
	return *o.PaymentCode
}

// GetPaymentCodeOk returns a tuple with the PaymentCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverTheCounterChannelProperties) GetPaymentCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PaymentCode) {
		return nil, false
	}
	return o.PaymentCode, true
}

// HasPaymentCode returns a boolean if a field has been set.
func (o *OverTheCounterChannelProperties) HasPaymentCode() bool {
	if o != nil && !utils.IsNil(o.PaymentCode) {
		return true
	}

	return false
}

// SetPaymentCode gets a reference to the given string and assigns it to the PaymentCode field.
func (o *OverTheCounterChannelProperties) SetPaymentCode(v string) {
	o.PaymentCode = &v
}

// GetCustomerName returns the CustomerName field value
func (o *OverTheCounterChannelProperties) GetCustomerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerName
}

// GetCustomerNameOk returns a tuple with the CustomerName field value
// and a boolean to check if the value has been set.
func (o *OverTheCounterChannelProperties) GetCustomerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerName, true
}

// SetCustomerName sets field value
func (o *OverTheCounterChannelProperties) SetCustomerName(v string) {
	o.CustomerName = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *OverTheCounterChannelProperties) GetExpiresAt() time.Time {
	if o == nil || utils.IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverTheCounterChannelProperties) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *OverTheCounterChannelProperties) HasExpiresAt() bool {
	if o != nil && !utils.IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *OverTheCounterChannelProperties) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

func (o OverTheCounterChannelProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OverTheCounterChannelProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.PaymentCode) {
		toSerialize["payment_code"] = o.PaymentCode
	}
	toSerialize["customer_name"] = o.CustomerName
	if !utils.IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	return toSerialize, nil
}

type NullableOverTheCounterChannelProperties struct {
	value *OverTheCounterChannelProperties
	isSet bool
}

func (v NullableOverTheCounterChannelProperties) Get() *OverTheCounterChannelProperties {
	return v.value
}

func (v *NullableOverTheCounterChannelProperties) Set(val *OverTheCounterChannelProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableOverTheCounterChannelProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableOverTheCounterChannelProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverTheCounterChannelProperties(val *OverTheCounterChannelProperties) *NullableOverTheCounterChannelProperties {
	return &NullableOverTheCounterChannelProperties{value: val, isSet: true}
}

func (v NullableOverTheCounterChannelProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverTheCounterChannelProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


