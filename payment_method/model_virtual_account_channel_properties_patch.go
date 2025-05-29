/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.128.0
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
	"time"
)

// checks if the VirtualAccountChannelPropertiesPatch type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &VirtualAccountChannelPropertiesPatch{}

// VirtualAccountChannelPropertiesPatch Virtual Account Channel Properties
type VirtualAccountChannelPropertiesPatch struct {
	// The date and time in ISO 8601 UTC+0 when the virtual account number will be expired. Default: The default expiration date will be 31 years from creation date.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// The suggested amount you want to assign. Note: Suggested amounts is the amounts that can see as a suggestion, but user can still put any numbers (only supported for Mandiri and BRI)
	SuggestedAmount *float64 `json:"suggested_amount,omitempty"`
}

// NewVirtualAccountChannelPropertiesPatch instantiates a new VirtualAccountChannelPropertiesPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualAccountChannelPropertiesPatch() *VirtualAccountChannelPropertiesPatch {
	this := VirtualAccountChannelPropertiesPatch{}
	return &this
}

// NewVirtualAccountChannelPropertiesPatchWithDefaults instantiates a new VirtualAccountChannelPropertiesPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualAccountChannelPropertiesPatchWithDefaults() *VirtualAccountChannelPropertiesPatch {
	this := VirtualAccountChannelPropertiesPatch{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *VirtualAccountChannelPropertiesPatch) GetExpiresAt() time.Time {
	if o == nil || utils.IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualAccountChannelPropertiesPatch) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *VirtualAccountChannelPropertiesPatch) HasExpiresAt() bool {
	if o != nil && !utils.IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *VirtualAccountChannelPropertiesPatch) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetSuggestedAmount returns the SuggestedAmount field value if set, zero value otherwise.
func (o *VirtualAccountChannelPropertiesPatch) GetSuggestedAmount() float64 {
	if o == nil || utils.IsNil(o.SuggestedAmount) {
		var ret float64
		return ret
	}
	return *o.SuggestedAmount
}

// GetSuggestedAmountOk returns a tuple with the SuggestedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualAccountChannelPropertiesPatch) GetSuggestedAmountOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.SuggestedAmount) {
		return nil, false
	}
	return o.SuggestedAmount, true
}

// HasSuggestedAmount returns a boolean if a field has been set.
func (o *VirtualAccountChannelPropertiesPatch) HasSuggestedAmount() bool {
	if o != nil && !utils.IsNil(o.SuggestedAmount) {
		return true
	}

	return false
}

// SetSuggestedAmount gets a reference to the given float64 and assigns it to the SuggestedAmount field.
func (o *VirtualAccountChannelPropertiesPatch) SetSuggestedAmount(v float64) {
	o.SuggestedAmount = &v
}

func (o VirtualAccountChannelPropertiesPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualAccountChannelPropertiesPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !utils.IsNil(o.SuggestedAmount) {
		toSerialize["suggested_amount"] = o.SuggestedAmount
	}
	return toSerialize, nil
}

type NullableVirtualAccountChannelPropertiesPatch struct {
	value *VirtualAccountChannelPropertiesPatch
	isSet bool
}

func (v NullableVirtualAccountChannelPropertiesPatch) Get() *VirtualAccountChannelPropertiesPatch {
	return v.value
}

func (v *NullableVirtualAccountChannelPropertiesPatch) Set(val *VirtualAccountChannelPropertiesPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualAccountChannelPropertiesPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualAccountChannelPropertiesPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualAccountChannelPropertiesPatch(val *VirtualAccountChannelPropertiesPatch) *NullableVirtualAccountChannelPropertiesPatch {
	return &NullableVirtualAccountChannelPropertiesPatch{value: val, isSet: true}
}

func (v NullableVirtualAccountChannelPropertiesPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualAccountChannelPropertiesPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


