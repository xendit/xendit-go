/*
XENDIT SDK openapi spec

XENDIT SDK openapi spec

API version: 1.0.0
*/


package customer

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
)

// checks if the AccountOTC type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AccountOTC{}

// AccountOTC struct for AccountOTC
type AccountOTC struct {
	// Complete fixed payment code (including prefix)
	PaymentCode *string `json:"payment_code,omitempty"`
	// YYYY-MM-DD string with expiry date for the payment code
	ExpiresAt NullableString `json:"expires_at,omitempty"`
}

// NewAccountOTC instantiates a new AccountOTC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountOTC() *AccountOTC {
	this := AccountOTC{}
	return &this
}

// NewAccountOTCWithDefaults instantiates a new AccountOTC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountOTCWithDefaults() *AccountOTC {
	this := AccountOTC{}
	return &this
}

// GetPaymentCode returns the PaymentCode field value if set, zero value otherwise.
func (o *AccountOTC) GetPaymentCode() string {
	if o == nil || utils.IsNil(o.PaymentCode) {
		var ret string
		return ret
	}
	return *o.PaymentCode
}

// GetPaymentCodeOk returns a tuple with the PaymentCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOTC) GetPaymentCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PaymentCode) {
		return nil, false
	}
	return o.PaymentCode, true
}

// HasPaymentCode returns a boolean if a field has been set.
func (o *AccountOTC) HasPaymentCode() bool {
	if o != nil && !utils.IsNil(o.PaymentCode) {
		return true
	}

	return false
}

// SetPaymentCode gets a reference to the given string and assigns it to the PaymentCode field.
func (o *AccountOTC) SetPaymentCode(v string) {
	o.PaymentCode = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountOTC) GetExpiresAt() string {
	if o == nil || utils.IsNil(o.ExpiresAt.Get()) {
		var ret string
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountOTC) GetExpiresAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *AccountOTC) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableString and assigns it to the ExpiresAt field.
func (o *AccountOTC) SetExpiresAt(v string) {
	o.ExpiresAt.Set(&v)
}
// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *AccountOTC) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *AccountOTC) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

func (o AccountOTC) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountOTC) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.PaymentCode) {
		toSerialize["payment_code"] = o.PaymentCode
	}
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
    }
	return toSerialize, nil
}

type NullableAccountOTC struct {
	value *AccountOTC
	isSet bool
}

func (v NullableAccountOTC) Get() *AccountOTC {
	return v.value
}

func (v *NullableAccountOTC) Set(val *AccountOTC) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountOTC) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountOTC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountOTC(val *AccountOTC) *NullableAccountOTC {
	return &NullableAccountOTC{value: val, isSet: true}
}

func (v NullableAccountOTC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountOTC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


