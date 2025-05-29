/*
Payment Requests

This API is used for Payment Requests

API version: 1.70.0
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
)

// checks if the DirectDebitChannelPropertiesDebitCard type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DirectDebitChannelPropertiesDebitCard{}

// DirectDebitChannelPropertiesDebitCard Direct Debit Debit Card Channel Properties
type DirectDebitChannelPropertiesDebitCard struct {
	// Mobile number of the customer registered to the partner channel
	MobileNumber *string `json:"mobile_number,omitempty"`
	AccountNumber *string `json:"account_number,omitempty"`
	// Last four digits of the debit card
	CardLastFour *string `json:"card_last_four,omitempty"`
	// Expiry month and year of the debit card (in MM/YY format)
	CardExpiry *string `json:"card_expiry,omitempty"`
	// Email address of the customer that is registered to the partner channel
	Email *string `json:"email,omitempty"`
}

// NewDirectDebitChannelPropertiesDebitCard instantiates a new DirectDebitChannelPropertiesDebitCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectDebitChannelPropertiesDebitCard() *DirectDebitChannelPropertiesDebitCard {
	this := DirectDebitChannelPropertiesDebitCard{}
	return &this
}

// NewDirectDebitChannelPropertiesDebitCardWithDefaults instantiates a new DirectDebitChannelPropertiesDebitCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectDebitChannelPropertiesDebitCardWithDefaults() *DirectDebitChannelPropertiesDebitCard {
	this := DirectDebitChannelPropertiesDebitCard{}
	return &this
}

// GetMobileNumber returns the MobileNumber field value if set, zero value otherwise.
func (o *DirectDebitChannelPropertiesDebitCard) GetMobileNumber() string {
	if o == nil || utils.IsNil(o.MobileNumber) {
		var ret string
		return ret
	}
	return *o.MobileNumber
}

// GetMobileNumberOk returns a tuple with the MobileNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDebitChannelPropertiesDebitCard) GetMobileNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MobileNumber) {
		return nil, false
	}
	return o.MobileNumber, true
}

// HasMobileNumber returns a boolean if a field has been set.
func (o *DirectDebitChannelPropertiesDebitCard) HasMobileNumber() bool {
	if o != nil && !utils.IsNil(o.MobileNumber) {
		return true
	}

	return false
}

// SetMobileNumber gets a reference to the given string and assigns it to the MobileNumber field.
func (o *DirectDebitChannelPropertiesDebitCard) SetMobileNumber(v string) {
	o.MobileNumber = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *DirectDebitChannelPropertiesDebitCard) GetAccountNumber() string {
	if o == nil || utils.IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDebitChannelPropertiesDebitCard) GetAccountNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *DirectDebitChannelPropertiesDebitCard) HasAccountNumber() bool {
	if o != nil && !utils.IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *DirectDebitChannelPropertiesDebitCard) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetCardLastFour returns the CardLastFour field value if set, zero value otherwise.
func (o *DirectDebitChannelPropertiesDebitCard) GetCardLastFour() string {
	if o == nil || utils.IsNil(o.CardLastFour) {
		var ret string
		return ret
	}
	return *o.CardLastFour
}

// GetCardLastFourOk returns a tuple with the CardLastFour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDebitChannelPropertiesDebitCard) GetCardLastFourOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CardLastFour) {
		return nil, false
	}
	return o.CardLastFour, true
}

// HasCardLastFour returns a boolean if a field has been set.
func (o *DirectDebitChannelPropertiesDebitCard) HasCardLastFour() bool {
	if o != nil && !utils.IsNil(o.CardLastFour) {
		return true
	}

	return false
}

// SetCardLastFour gets a reference to the given string and assigns it to the CardLastFour field.
func (o *DirectDebitChannelPropertiesDebitCard) SetCardLastFour(v string) {
	o.CardLastFour = &v
}

// GetCardExpiry returns the CardExpiry field value if set, zero value otherwise.
func (o *DirectDebitChannelPropertiesDebitCard) GetCardExpiry() string {
	if o == nil || utils.IsNil(o.CardExpiry) {
		var ret string
		return ret
	}
	return *o.CardExpiry
}

// GetCardExpiryOk returns a tuple with the CardExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDebitChannelPropertiesDebitCard) GetCardExpiryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CardExpiry) {
		return nil, false
	}
	return o.CardExpiry, true
}

// HasCardExpiry returns a boolean if a field has been set.
func (o *DirectDebitChannelPropertiesDebitCard) HasCardExpiry() bool {
	if o != nil && !utils.IsNil(o.CardExpiry) {
		return true
	}

	return false
}

// SetCardExpiry gets a reference to the given string and assigns it to the CardExpiry field.
func (o *DirectDebitChannelPropertiesDebitCard) SetCardExpiry(v string) {
	o.CardExpiry = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *DirectDebitChannelPropertiesDebitCard) GetEmail() string {
	if o == nil || utils.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDebitChannelPropertiesDebitCard) GetEmailOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *DirectDebitChannelPropertiesDebitCard) HasEmail() bool {
	if o != nil && !utils.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *DirectDebitChannelPropertiesDebitCard) SetEmail(v string) {
	o.Email = &v
}

func (o DirectDebitChannelPropertiesDebitCard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectDebitChannelPropertiesDebitCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.MobileNumber) {
		toSerialize["mobile_number"] = o.MobileNumber
	}
	if !utils.IsNil(o.AccountNumber) {
		toSerialize["account_number"] = o.AccountNumber
	}
	if !utils.IsNil(o.CardLastFour) {
		toSerialize["card_last_four"] = o.CardLastFour
	}
	if !utils.IsNil(o.CardExpiry) {
		toSerialize["card_expiry"] = o.CardExpiry
	}
	if !utils.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return toSerialize, nil
}

type NullableDirectDebitChannelPropertiesDebitCard struct {
	value *DirectDebitChannelPropertiesDebitCard
	isSet bool
}

func (v NullableDirectDebitChannelPropertiesDebitCard) Get() *DirectDebitChannelPropertiesDebitCard {
	return v.value
}

func (v *NullableDirectDebitChannelPropertiesDebitCard) Set(val *DirectDebitChannelPropertiesDebitCard) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectDebitChannelPropertiesDebitCard) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectDebitChannelPropertiesDebitCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectDebitChannelPropertiesDebitCard(val *DirectDebitChannelPropertiesDebitCard) *NullableDirectDebitChannelPropertiesDebitCard {
	return &NullableDirectDebitChannelPropertiesDebitCard{value: val, isSet: true}
}

func (v NullableDirectDebitChannelPropertiesDebitCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectDebitChannelPropertiesDebitCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


