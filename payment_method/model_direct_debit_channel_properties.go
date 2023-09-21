/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.87.2
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the DirectDebitChannelProperties type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DirectDebitChannelProperties{}

// DirectDebitChannelProperties Direct Debit Channel Properties
type DirectDebitChannelProperties struct {
	SuccessReturnUrl *string `json:"success_return_url,omitempty"`
	FailureReturnUrl NullableString `json:"failure_return_url,omitempty"`
	// Mobile number of the customer registered to the partner channel
	MobileNumber NullableString `json:"mobile_number,omitempty"`
	// Last four digits of the debit card
	CardLastFour NullableString `json:"card_last_four,omitempty"`
	// Expiry month and year of the debit card (in MM/YY format)
	CardExpiry NullableString `json:"card_expiry,omitempty"`
	// Email address of the customer that is registered to the partner channel
	Email NullableString `json:"email,omitempty"`
	// Identity number of the customer registered to the partner channel
	IdentityDocumentNumber NullableString `json:"identity_document_number,omitempty"`
	RequireAuth NullableBool `json:"require_auth,omitempty"`
}

// NewDirectDebitChannelProperties instantiates a new DirectDebitChannelProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectDebitChannelProperties() *DirectDebitChannelProperties {
	this := DirectDebitChannelProperties{}
	return &this
}

// NewDirectDebitChannelPropertiesWithDefaults instantiates a new DirectDebitChannelProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectDebitChannelPropertiesWithDefaults() *DirectDebitChannelProperties {
	this := DirectDebitChannelProperties{}
	return &this
}

// GetSuccessReturnUrl returns the SuccessReturnUrl field value if set, zero value otherwise.
func (o *DirectDebitChannelProperties) GetSuccessReturnUrl() string {
	if o == nil || utils.IsNil(o.SuccessReturnUrl) {
		var ret string
		return ret
	}
	return *o.SuccessReturnUrl
}

// GetSuccessReturnUrlOk returns a tuple with the SuccessReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDebitChannelProperties) GetSuccessReturnUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SuccessReturnUrl) {
		return nil, false
	}
	return o.SuccessReturnUrl, true
}

// HasSuccessReturnUrl returns a boolean if a field has been set.
func (o *DirectDebitChannelProperties) HasSuccessReturnUrl() bool {
	if o != nil && !utils.IsNil(o.SuccessReturnUrl) {
		return true
	}

	return false
}

// SetSuccessReturnUrl gets a reference to the given string and assigns it to the SuccessReturnUrl field.
func (o *DirectDebitChannelProperties) SetSuccessReturnUrl(v string) {
	o.SuccessReturnUrl = &v
}

// GetFailureReturnUrl returns the FailureReturnUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectDebitChannelProperties) GetFailureReturnUrl() string {
	if o == nil || utils.IsNil(o.FailureReturnUrl.Get()) {
		var ret string
		return ret
	}
	return *o.FailureReturnUrl.Get()
}

// GetFailureReturnUrlOk returns a tuple with the FailureReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectDebitChannelProperties) GetFailureReturnUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureReturnUrl.Get(), o.FailureReturnUrl.IsSet()
}

// HasFailureReturnUrl returns a boolean if a field has been set.
func (o *DirectDebitChannelProperties) HasFailureReturnUrl() bool {
	if o != nil && o.FailureReturnUrl.IsSet() {
		return true
	}

	return false
}

// SetFailureReturnUrl gets a reference to the given NullableString and assigns it to the FailureReturnUrl field.
func (o *DirectDebitChannelProperties) SetFailureReturnUrl(v string) {
	o.FailureReturnUrl.Set(&v)
}
// SetFailureReturnUrlNil sets the value for FailureReturnUrl to be an explicit nil
func (o *DirectDebitChannelProperties) SetFailureReturnUrlNil() {
	o.FailureReturnUrl.Set(nil)
}

// UnsetFailureReturnUrl ensures that no value is present for FailureReturnUrl, not even an explicit nil
func (o *DirectDebitChannelProperties) UnsetFailureReturnUrl() {
	o.FailureReturnUrl.Unset()
}

// GetMobileNumber returns the MobileNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectDebitChannelProperties) GetMobileNumber() string {
	if o == nil || utils.IsNil(o.MobileNumber.Get()) {
		var ret string
		return ret
	}
	return *o.MobileNumber.Get()
}

// GetMobileNumberOk returns a tuple with the MobileNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectDebitChannelProperties) GetMobileNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileNumber.Get(), o.MobileNumber.IsSet()
}

// HasMobileNumber returns a boolean if a field has been set.
func (o *DirectDebitChannelProperties) HasMobileNumber() bool {
	if o != nil && o.MobileNumber.IsSet() {
		return true
	}

	return false
}

// SetMobileNumber gets a reference to the given NullableString and assigns it to the MobileNumber field.
func (o *DirectDebitChannelProperties) SetMobileNumber(v string) {
	o.MobileNumber.Set(&v)
}
// SetMobileNumberNil sets the value for MobileNumber to be an explicit nil
func (o *DirectDebitChannelProperties) SetMobileNumberNil() {
	o.MobileNumber.Set(nil)
}

// UnsetMobileNumber ensures that no value is present for MobileNumber, not even an explicit nil
func (o *DirectDebitChannelProperties) UnsetMobileNumber() {
	o.MobileNumber.Unset()
}

// GetCardLastFour returns the CardLastFour field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectDebitChannelProperties) GetCardLastFour() string {
	if o == nil || utils.IsNil(o.CardLastFour.Get()) {
		var ret string
		return ret
	}
	return *o.CardLastFour.Get()
}

// GetCardLastFourOk returns a tuple with the CardLastFour field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectDebitChannelProperties) GetCardLastFourOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardLastFour.Get(), o.CardLastFour.IsSet()
}

// HasCardLastFour returns a boolean if a field has been set.
func (o *DirectDebitChannelProperties) HasCardLastFour() bool {
	if o != nil && o.CardLastFour.IsSet() {
		return true
	}

	return false
}

// SetCardLastFour gets a reference to the given NullableString and assigns it to the CardLastFour field.
func (o *DirectDebitChannelProperties) SetCardLastFour(v string) {
	o.CardLastFour.Set(&v)
}
// SetCardLastFourNil sets the value for CardLastFour to be an explicit nil
func (o *DirectDebitChannelProperties) SetCardLastFourNil() {
	o.CardLastFour.Set(nil)
}

// UnsetCardLastFour ensures that no value is present for CardLastFour, not even an explicit nil
func (o *DirectDebitChannelProperties) UnsetCardLastFour() {
	o.CardLastFour.Unset()
}

// GetCardExpiry returns the CardExpiry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectDebitChannelProperties) GetCardExpiry() string {
	if o == nil || utils.IsNil(o.CardExpiry.Get()) {
		var ret string
		return ret
	}
	return *o.CardExpiry.Get()
}

// GetCardExpiryOk returns a tuple with the CardExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectDebitChannelProperties) GetCardExpiryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardExpiry.Get(), o.CardExpiry.IsSet()
}

// HasCardExpiry returns a boolean if a field has been set.
func (o *DirectDebitChannelProperties) HasCardExpiry() bool {
	if o != nil && o.CardExpiry.IsSet() {
		return true
	}

	return false
}

// SetCardExpiry gets a reference to the given NullableString and assigns it to the CardExpiry field.
func (o *DirectDebitChannelProperties) SetCardExpiry(v string) {
	o.CardExpiry.Set(&v)
}
// SetCardExpiryNil sets the value for CardExpiry to be an explicit nil
func (o *DirectDebitChannelProperties) SetCardExpiryNil() {
	o.CardExpiry.Set(nil)
}

// UnsetCardExpiry ensures that no value is present for CardExpiry, not even an explicit nil
func (o *DirectDebitChannelProperties) UnsetCardExpiry() {
	o.CardExpiry.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectDebitChannelProperties) GetEmail() string {
	if o == nil || utils.IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectDebitChannelProperties) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *DirectDebitChannelProperties) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *DirectDebitChannelProperties) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *DirectDebitChannelProperties) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *DirectDebitChannelProperties) UnsetEmail() {
	o.Email.Unset()
}

// GetIdentityDocumentNumber returns the IdentityDocumentNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectDebitChannelProperties) GetIdentityDocumentNumber() string {
	if o == nil || utils.IsNil(o.IdentityDocumentNumber.Get()) {
		var ret string
		return ret
	}
	return *o.IdentityDocumentNumber.Get()
}

// GetIdentityDocumentNumberOk returns a tuple with the IdentityDocumentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectDebitChannelProperties) GetIdentityDocumentNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdentityDocumentNumber.Get(), o.IdentityDocumentNumber.IsSet()
}

// HasIdentityDocumentNumber returns a boolean if a field has been set.
func (o *DirectDebitChannelProperties) HasIdentityDocumentNumber() bool {
	if o != nil && o.IdentityDocumentNumber.IsSet() {
		return true
	}

	return false
}

// SetIdentityDocumentNumber gets a reference to the given NullableString and assigns it to the IdentityDocumentNumber field.
func (o *DirectDebitChannelProperties) SetIdentityDocumentNumber(v string) {
	o.IdentityDocumentNumber.Set(&v)
}
// SetIdentityDocumentNumberNil sets the value for IdentityDocumentNumber to be an explicit nil
func (o *DirectDebitChannelProperties) SetIdentityDocumentNumberNil() {
	o.IdentityDocumentNumber.Set(nil)
}

// UnsetIdentityDocumentNumber ensures that no value is present for IdentityDocumentNumber, not even an explicit nil
func (o *DirectDebitChannelProperties) UnsetIdentityDocumentNumber() {
	o.IdentityDocumentNumber.Unset()
}

// GetRequireAuth returns the RequireAuth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectDebitChannelProperties) GetRequireAuth() bool {
	if o == nil || utils.IsNil(o.RequireAuth.Get()) {
		var ret bool
		return ret
	}
	return *o.RequireAuth.Get()
}

// GetRequireAuthOk returns a tuple with the RequireAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectDebitChannelProperties) GetRequireAuthOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequireAuth.Get(), o.RequireAuth.IsSet()
}

// HasRequireAuth returns a boolean if a field has been set.
func (o *DirectDebitChannelProperties) HasRequireAuth() bool {
	if o != nil && o.RequireAuth.IsSet() {
		return true
	}

	return false
}

// SetRequireAuth gets a reference to the given NullableBool and assigns it to the RequireAuth field.
func (o *DirectDebitChannelProperties) SetRequireAuth(v bool) {
	o.RequireAuth.Set(&v)
}
// SetRequireAuthNil sets the value for RequireAuth to be an explicit nil
func (o *DirectDebitChannelProperties) SetRequireAuthNil() {
	o.RequireAuth.Set(nil)
}

// UnsetRequireAuth ensures that no value is present for RequireAuth, not even an explicit nil
func (o *DirectDebitChannelProperties) UnsetRequireAuth() {
	o.RequireAuth.Unset()
}

func (o DirectDebitChannelProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectDebitChannelProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.SuccessReturnUrl) {
		toSerialize["success_return_url"] = o.SuccessReturnUrl
	}
	if o.FailureReturnUrl.IsSet() {
		toSerialize["failure_return_url"] = o.FailureReturnUrl.Get()
	}
	if o.MobileNumber.IsSet() {
		toSerialize["mobile_number"] = o.MobileNumber.Get()
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
	if o.IdentityDocumentNumber.IsSet() {
		toSerialize["identity_document_number"] = o.IdentityDocumentNumber.Get()
	}
	if o.RequireAuth.IsSet() {
		toSerialize["require_auth"] = o.RequireAuth.Get()
	}
	return toSerialize, nil
}

type NullableDirectDebitChannelProperties struct {
	value *DirectDebitChannelProperties
	isSet bool
}

func (v NullableDirectDebitChannelProperties) Get() *DirectDebitChannelProperties {
	return v.value
}

func (v *NullableDirectDebitChannelProperties) Set(val *DirectDebitChannelProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectDebitChannelProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectDebitChannelProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectDebitChannelProperties(val *DirectDebitChannelProperties) *NullableDirectDebitChannelProperties {
	return &NullableDirectDebitChannelProperties{value: val, isSet: true}
}

func (v NullableDirectDebitChannelProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectDebitChannelProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


