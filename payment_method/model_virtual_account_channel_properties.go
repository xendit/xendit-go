/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.128.0
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
	"time"
)

// checks if the VirtualAccountChannelProperties type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &VirtualAccountChannelProperties{}

// VirtualAccountChannelProperties Virtual Account Channel Properties
type VirtualAccountChannelProperties struct {
	// Name of customer.
	CustomerName *string `json:"customer_name,omitempty"`
	// You can assign specific Virtual Account number using this parameter. If you do not send one, one will be picked at random. Make sure the number you specify is within your Virtual Account range.
	VirtualAccountNumber *string `json:"virtual_account_number,omitempty"`
	// The date and time in ISO 8601 UTC+0 when the virtual account number will be expired. Default: The default expiration date will be 31 years from creation date.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// The suggested amount you want to assign. Note: Suggested amounts is the amounts that can see as a suggestion, but user can still put any numbers (only supported for Mandiri and BRI)
	SuggestedAmount *float64 `json:"suggested_amount,omitempty"`
}

// NewVirtualAccountChannelProperties instantiates a new VirtualAccountChannelProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualAccountChannelProperties() *VirtualAccountChannelProperties {
	this := VirtualAccountChannelProperties{}
	return &this
}

// NewVirtualAccountChannelPropertiesWithDefaults instantiates a new VirtualAccountChannelProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualAccountChannelPropertiesWithDefaults() *VirtualAccountChannelProperties {
	this := VirtualAccountChannelProperties{}
	return &this
}

// GetCustomerName returns the CustomerName field value if set, zero value otherwise.
func (o *VirtualAccountChannelProperties) GetCustomerName() string {
	if o == nil || utils.IsNil(o.CustomerName) {
		var ret string
		return ret
	}
	return *o.CustomerName
}

// GetCustomerNameOk returns a tuple with the CustomerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualAccountChannelProperties) GetCustomerNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CustomerName) {
		return nil, false
	}
	return o.CustomerName, true
}

// HasCustomerName returns a boolean if a field has been set.
func (o *VirtualAccountChannelProperties) HasCustomerName() bool {
	if o != nil && !utils.IsNil(o.CustomerName) {
		return true
	}

	return false
}

// SetCustomerName gets a reference to the given string and assigns it to the CustomerName field.
func (o *VirtualAccountChannelProperties) SetCustomerName(v string) {
	o.CustomerName = &v
}

// GetVirtualAccountNumber returns the VirtualAccountNumber field value if set, zero value otherwise.
func (o *VirtualAccountChannelProperties) GetVirtualAccountNumber() string {
	if o == nil || utils.IsNil(o.VirtualAccountNumber) {
		var ret string
		return ret
	}
	return *o.VirtualAccountNumber
}

// GetVirtualAccountNumberOk returns a tuple with the VirtualAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualAccountChannelProperties) GetVirtualAccountNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.VirtualAccountNumber) {
		return nil, false
	}
	return o.VirtualAccountNumber, true
}

// HasVirtualAccountNumber returns a boolean if a field has been set.
func (o *VirtualAccountChannelProperties) HasVirtualAccountNumber() bool {
	if o != nil && !utils.IsNil(o.VirtualAccountNumber) {
		return true
	}

	return false
}

// SetVirtualAccountNumber gets a reference to the given string and assigns it to the VirtualAccountNumber field.
func (o *VirtualAccountChannelProperties) SetVirtualAccountNumber(v string) {
	o.VirtualAccountNumber = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *VirtualAccountChannelProperties) GetExpiresAt() time.Time {
	if o == nil || utils.IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualAccountChannelProperties) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *VirtualAccountChannelProperties) HasExpiresAt() bool {
	if o != nil && !utils.IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *VirtualAccountChannelProperties) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetSuggestedAmount returns the SuggestedAmount field value if set, zero value otherwise.
func (o *VirtualAccountChannelProperties) GetSuggestedAmount() float64 {
	if o == nil || utils.IsNil(o.SuggestedAmount) {
		var ret float64
		return ret
	}
	return *o.SuggestedAmount
}

// GetSuggestedAmountOk returns a tuple with the SuggestedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualAccountChannelProperties) GetSuggestedAmountOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.SuggestedAmount) {
		return nil, false
	}
	return o.SuggestedAmount, true
}

// HasSuggestedAmount returns a boolean if a field has been set.
func (o *VirtualAccountChannelProperties) HasSuggestedAmount() bool {
	if o != nil && !utils.IsNil(o.SuggestedAmount) {
		return true
	}

	return false
}

// SetSuggestedAmount gets a reference to the given float64 and assigns it to the SuggestedAmount field.
func (o *VirtualAccountChannelProperties) SetSuggestedAmount(v float64) {
	o.SuggestedAmount = &v
}

func (o VirtualAccountChannelProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualAccountChannelProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CustomerName) {
		toSerialize["customer_name"] = o.CustomerName
	}
	if !utils.IsNil(o.VirtualAccountNumber) {
		toSerialize["virtual_account_number"] = o.VirtualAccountNumber
	}
	if !utils.IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !utils.IsNil(o.SuggestedAmount) {
		toSerialize["suggested_amount"] = o.SuggestedAmount
	}
	return toSerialize, nil
}

type NullableVirtualAccountChannelProperties struct {
	value *VirtualAccountChannelProperties
	isSet bool
}

func (v NullableVirtualAccountChannelProperties) Get() *VirtualAccountChannelProperties {
	return v.value
}

func (v *NullableVirtualAccountChannelProperties) Set(val *VirtualAccountChannelProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualAccountChannelProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualAccountChannelProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualAccountChannelProperties(val *VirtualAccountChannelProperties) *NullableVirtualAccountChannelProperties {
	return &NullableVirtualAccountChannelProperties{value: val, isSet: true}
}

func (v NullableVirtualAccountChannelProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualAccountChannelProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


