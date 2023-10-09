/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.89.1
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
	"time"
)

// checks if the OverTheCounterChannelPropertiesUpdate type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &OverTheCounterChannelPropertiesUpdate{}

// OverTheCounterChannelPropertiesUpdate Over The Counter Channel properties that can be updated
type OverTheCounterChannelPropertiesUpdate struct {
	// Name of customer.
	CustomerName *string `json:"customer_name,omitempty"`
	// The time when the payment code will be expired. The minimum is 2 hours and the maximum is 9 days for 7ELEVEN. Default expired date will be 2 days from payment code generated.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// NewOverTheCounterChannelPropertiesUpdate instantiates a new OverTheCounterChannelPropertiesUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverTheCounterChannelPropertiesUpdate() *OverTheCounterChannelPropertiesUpdate {
	this := OverTheCounterChannelPropertiesUpdate{}
	return &this
}

// NewOverTheCounterChannelPropertiesUpdateWithDefaults instantiates a new OverTheCounterChannelPropertiesUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverTheCounterChannelPropertiesUpdateWithDefaults() *OverTheCounterChannelPropertiesUpdate {
	this := OverTheCounterChannelPropertiesUpdate{}
	return &this
}

// GetCustomerName returns the CustomerName field value if set, zero value otherwise.
func (o *OverTheCounterChannelPropertiesUpdate) GetCustomerName() string {
	if o == nil || utils.IsNil(o.CustomerName) {
		var ret string
		return ret
	}
	return *o.CustomerName
}

// GetCustomerNameOk returns a tuple with the CustomerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverTheCounterChannelPropertiesUpdate) GetCustomerNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CustomerName) {
		return nil, false
	}
	return o.CustomerName, true
}

// HasCustomerName returns a boolean if a field has been set.
func (o *OverTheCounterChannelPropertiesUpdate) HasCustomerName() bool {
	if o != nil && !utils.IsNil(o.CustomerName) {
		return true
	}

	return false
}

// SetCustomerName gets a reference to the given string and assigns it to the CustomerName field.
func (o *OverTheCounterChannelPropertiesUpdate) SetCustomerName(v string) {
	o.CustomerName = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *OverTheCounterChannelPropertiesUpdate) GetExpiresAt() time.Time {
	if o == nil || utils.IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverTheCounterChannelPropertiesUpdate) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *OverTheCounterChannelPropertiesUpdate) HasExpiresAt() bool {
	if o != nil && !utils.IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *OverTheCounterChannelPropertiesUpdate) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

func (o OverTheCounterChannelPropertiesUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OverTheCounterChannelPropertiesUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CustomerName) {
		toSerialize["customer_name"] = o.CustomerName
	}
	if !utils.IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	return toSerialize, nil
}

type NullableOverTheCounterChannelPropertiesUpdate struct {
	value *OverTheCounterChannelPropertiesUpdate
	isSet bool
}

func (v NullableOverTheCounterChannelPropertiesUpdate) Get() *OverTheCounterChannelPropertiesUpdate {
	return v.value
}

func (v *NullableOverTheCounterChannelPropertiesUpdate) Set(val *OverTheCounterChannelPropertiesUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableOverTheCounterChannelPropertiesUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableOverTheCounterChannelPropertiesUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverTheCounterChannelPropertiesUpdate(val *OverTheCounterChannelPropertiesUpdate) *NullableOverTheCounterChannelPropertiesUpdate {
	return &NullableOverTheCounterChannelPropertiesUpdate{value: val, isSet: true}
}

func (v NullableOverTheCounterChannelPropertiesUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverTheCounterChannelPropertiesUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


