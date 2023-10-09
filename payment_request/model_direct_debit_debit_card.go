/*
Payment Requests

This API is used for Payment Requests

API version: 1.44.1
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the DirectDebitDebitCard type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DirectDebitDebitCard{}

// DirectDebitDebitCard struct for DirectDebitDebitCard
type DirectDebitDebitCard struct {
	// Mobile number of the customer registered to the partner channel
	MobileNumber NullableString `json:"mobile_number,omitempty"`
	AccountNumber *string `json:"account_number,omitempty"`
	// Last four digits of the debit card
	CardLastFour NullableString `json:"card_last_four,omitempty"`
	// Expiry month and year of the debit card (in MM/YY format)
	CardExpiry NullableString `json:"card_expiry,omitempty"`
	// Email address of the customer that is registered to the partner channel
	Email NullableString `json:"email,omitempty"`
}

// NewDirectDebitDebitCard instantiates a new DirectDebitDebitCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectDebitDebitCard() *DirectDebitDebitCard {
	this := DirectDebitDebitCard{}
	return &this
}

// NewDirectDebitDebitCardWithDefaults instantiates a new DirectDebitDebitCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectDebitDebitCardWithDefaults() *DirectDebitDebitCard {
	this := DirectDebitDebitCard{}
	return &this
}

// GetMobileNumber returns the MobileNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectDebitDebitCard) GetMobileNumber() string {
	if o == nil || utils.IsNil(o.MobileNumber.Get()) {
		var ret string
		return ret
	}
	return *o.MobileNumber.Get()
}

// GetMobileNumberOk returns a tuple with the MobileNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectDebitDebitCard) GetMobileNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileNumber.Get(), o.MobileNumber.IsSet()
}

// HasMobileNumber returns a boolean if a field has been set.
func (o *DirectDebitDebitCard) HasMobileNumber() bool {
	if o != nil && o.MobileNumber.IsSet() {
		return true
	}

	return false
}

// SetMobileNumber gets a reference to the given NullableString and assigns it to the MobileNumber field.
func (o *DirectDebitDebitCard) SetMobileNumber(v string) {
	o.MobileNumber.Set(&v)
}
// SetMobileNumberNil sets the value for MobileNumber to be an explicit nil
func (o *DirectDebitDebitCard) SetMobileNumberNil() {
	o.MobileNumber.Set(nil)
}

// UnsetMobileNumber ensures that no value is present for MobileNumber, not even an explicit nil
func (o *DirectDebitDebitCard) UnsetMobileNumber() {
	o.MobileNumber.Unset()
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *DirectDebitDebitCard) GetAccountNumber() string {
	if o == nil || utils.IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDebitDebitCard) GetAccountNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *DirectDebitDebitCard) HasAccountNumber() bool {
	if o != nil && !utils.IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *DirectDebitDebitCard) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetCardLastFour returns the CardLastFour field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectDebitDebitCard) GetCardLastFour() string {
	if o == nil || utils.IsNil(o.CardLastFour.Get()) {
		var ret string
		return ret
	}
	return *o.CardLastFour.Get()
}

// GetCardLastFourOk returns a tuple with the CardLastFour field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectDebitDebitCard) GetCardLastFourOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardLastFour.Get(), o.CardLastFour.IsSet()
}

// HasCardLastFour returns a boolean if a field has been set.
func (o *DirectDebitDebitCard) HasCardLastFour() bool {
	if o != nil && o.CardLastFour.IsSet() {
		return true
	}

	return false
}

// SetCardLastFour gets a reference to the given NullableString and assigns it to the CardLastFour field.
func (o *DirectDebitDebitCard) SetCardLastFour(v string) {
	o.CardLastFour.Set(&v)
}
// SetCardLastFourNil sets the value for CardLastFour to be an explicit nil
func (o *DirectDebitDebitCard) SetCardLastFourNil() {
	o.CardLastFour.Set(nil)
}

// UnsetCardLastFour ensures that no value is present for CardLastFour, not even an explicit nil
func (o *DirectDebitDebitCard) UnsetCardLastFour() {
	o.CardLastFour.Unset()
}

// GetCardExpiry returns the CardExpiry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectDebitDebitCard) GetCardExpiry() string {
	if o == nil || utils.IsNil(o.CardExpiry.Get()) {
		var ret string
		return ret
	}
	return *o.CardExpiry.Get()
}

// GetCardExpiryOk returns a tuple with the CardExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectDebitDebitCard) GetCardExpiryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardExpiry.Get(), o.CardExpiry.IsSet()
}

// HasCardExpiry returns a boolean if a field has been set.
func (o *DirectDebitDebitCard) HasCardExpiry() bool {
	if o != nil && o.CardExpiry.IsSet() {
		return true
	}

	return false
}

// SetCardExpiry gets a reference to the given NullableString and assigns it to the CardExpiry field.
func (o *DirectDebitDebitCard) SetCardExpiry(v string) {
	o.CardExpiry.Set(&v)
}
// SetCardExpiryNil sets the value for CardExpiry to be an explicit nil
func (o *DirectDebitDebitCard) SetCardExpiryNil() {
	o.CardExpiry.Set(nil)
}

// UnsetCardExpiry ensures that no value is present for CardExpiry, not even an explicit nil
func (o *DirectDebitDebitCard) UnsetCardExpiry() {
	o.CardExpiry.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectDebitDebitCard) GetEmail() string {
	if o == nil || utils.IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectDebitDebitCard) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *DirectDebitDebitCard) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *DirectDebitDebitCard) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *DirectDebitDebitCard) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *DirectDebitDebitCard) UnsetEmail() {
	o.Email.Unset()
}

func (o DirectDebitDebitCard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectDebitDebitCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MobileNumber.IsSet() {
		toSerialize["mobile_number"] = o.MobileNumber.Get()
    }
	if !utils.IsNil(o.AccountNumber) {
		toSerialize["account_number"] = o.AccountNumber
	}
	if o.CardLastFour.IsSet() {
		toSerialize["card_last_four"] = o.CardLastFour.Get()
    }
	if o.CardExpiry.IsSet() {
		toSerialize["card_expiry"] = o.CardExpiry.Get()
    }
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
    }
	return toSerialize, nil
}

type NullableDirectDebitDebitCard struct {
	value *DirectDebitDebitCard
	isSet bool
}

func (v NullableDirectDebitDebitCard) Get() *DirectDebitDebitCard {
	return v.value
}

func (v *NullableDirectDebitDebitCard) Set(val *DirectDebitDebitCard) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectDebitDebitCard) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectDebitDebitCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectDebitDebitCard(val *DirectDebitDebitCard) *NullableDirectDebitDebitCard {
	return &NullableDirectDebitDebitCard{value: val, isSet: true}
}

func (v NullableDirectDebitDebitCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectDebitDebitCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


