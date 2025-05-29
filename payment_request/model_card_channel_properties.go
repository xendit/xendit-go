/*
Payment Requests

This API is used for Payment Requests

API version: 1.70.0
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
	"time"
)

// checks if the CardChannelProperties type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CardChannelProperties{}

// CardChannelProperties Card Channel Properties
type CardChannelProperties struct {
	// To indicate whether to perform 3DS during the linking phase
	SkipThreeDSecure NullableBool `json:"skip_three_d_secure,omitempty"`
	// URL where the end-customer is redirected if the authorization is successful
	SuccessReturnUrl NullableString `json:"success_return_url,omitempty"`
	// URL where the end-customer is redirected if the authorization failed
	FailureReturnUrl NullableString `json:"failure_return_url,omitempty"`
	// Type of “credential-on-file” / “card-on-file” payment being made. Indicate that this payment uses a previously linked Payment Method for charging.
	CardonfileType NullableString `json:"cardonfile_type,omitempty"`
	// Tag for a Merchant ID that you want to associate this payment with. For merchants using their own MIDs to specify which MID they want to use
	MerchantIdTag *string `json:"merchant_id_tag,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	InstallmentConfiguration NullableCardInstallmentConfiguration `json:"installment_configuration,omitempty"`
	// To indicate whether to skip the authorization phase
	SkipAuthorization *bool `json:"skip_authorization,omitempty"`
}

// NewCardChannelProperties instantiates a new CardChannelProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardChannelProperties() *CardChannelProperties {
	this := CardChannelProperties{}
	return &this
}

// NewCardChannelPropertiesWithDefaults instantiates a new CardChannelProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardChannelPropertiesWithDefaults() *CardChannelProperties {
	this := CardChannelProperties{}
	return &this
}

// GetSkipThreeDSecure returns the SkipThreeDSecure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardChannelProperties) GetSkipThreeDSecure() bool {
	if o == nil || utils.IsNil(o.SkipThreeDSecure.Get()) {
		var ret bool
		return ret
	}
	return *o.SkipThreeDSecure.Get()
}

// GetSkipThreeDSecureOk returns a tuple with the SkipThreeDSecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardChannelProperties) GetSkipThreeDSecureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipThreeDSecure.Get(), o.SkipThreeDSecure.IsSet()
}

// HasSkipThreeDSecure returns a boolean if a field has been set.
func (o *CardChannelProperties) HasSkipThreeDSecure() bool {
	if o != nil && o.SkipThreeDSecure.IsSet() {
		return true
	}

	return false
}

// SetSkipThreeDSecure gets a reference to the given NullableBool and assigns it to the SkipThreeDSecure field.
func (o *CardChannelProperties) SetSkipThreeDSecure(v bool) {
	o.SkipThreeDSecure.Set(&v)
}
// SetSkipThreeDSecureNil sets the value for SkipThreeDSecure to be an explicit nil
func (o *CardChannelProperties) SetSkipThreeDSecureNil() {
	o.SkipThreeDSecure.Set(nil)
}

// UnsetSkipThreeDSecure ensures that no value is present for SkipThreeDSecure, not even an explicit nil
func (o *CardChannelProperties) UnsetSkipThreeDSecure() {
	o.SkipThreeDSecure.Unset()
}

// GetSuccessReturnUrl returns the SuccessReturnUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardChannelProperties) GetSuccessReturnUrl() string {
	if o == nil || utils.IsNil(o.SuccessReturnUrl.Get()) {
		var ret string
		return ret
	}
	return *o.SuccessReturnUrl.Get()
}

// GetSuccessReturnUrlOk returns a tuple with the SuccessReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardChannelProperties) GetSuccessReturnUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SuccessReturnUrl.Get(), o.SuccessReturnUrl.IsSet()
}

// HasSuccessReturnUrl returns a boolean if a field has been set.
func (o *CardChannelProperties) HasSuccessReturnUrl() bool {
	if o != nil && o.SuccessReturnUrl.IsSet() {
		return true
	}

	return false
}

// SetSuccessReturnUrl gets a reference to the given NullableString and assigns it to the SuccessReturnUrl field.
func (o *CardChannelProperties) SetSuccessReturnUrl(v string) {
	o.SuccessReturnUrl.Set(&v)
}
// SetSuccessReturnUrlNil sets the value for SuccessReturnUrl to be an explicit nil
func (o *CardChannelProperties) SetSuccessReturnUrlNil() {
	o.SuccessReturnUrl.Set(nil)
}

// UnsetSuccessReturnUrl ensures that no value is present for SuccessReturnUrl, not even an explicit nil
func (o *CardChannelProperties) UnsetSuccessReturnUrl() {
	o.SuccessReturnUrl.Unset()
}

// GetFailureReturnUrl returns the FailureReturnUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardChannelProperties) GetFailureReturnUrl() string {
	if o == nil || utils.IsNil(o.FailureReturnUrl.Get()) {
		var ret string
		return ret
	}
	return *o.FailureReturnUrl.Get()
}

// GetFailureReturnUrlOk returns a tuple with the FailureReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardChannelProperties) GetFailureReturnUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureReturnUrl.Get(), o.FailureReturnUrl.IsSet()
}

// HasFailureReturnUrl returns a boolean if a field has been set.
func (o *CardChannelProperties) HasFailureReturnUrl() bool {
	if o != nil && o.FailureReturnUrl.IsSet() {
		return true
	}

	return false
}

// SetFailureReturnUrl gets a reference to the given NullableString and assigns it to the FailureReturnUrl field.
func (o *CardChannelProperties) SetFailureReturnUrl(v string) {
	o.FailureReturnUrl.Set(&v)
}
// SetFailureReturnUrlNil sets the value for FailureReturnUrl to be an explicit nil
func (o *CardChannelProperties) SetFailureReturnUrlNil() {
	o.FailureReturnUrl.Set(nil)
}

// UnsetFailureReturnUrl ensures that no value is present for FailureReturnUrl, not even an explicit nil
func (o *CardChannelProperties) UnsetFailureReturnUrl() {
	o.FailureReturnUrl.Unset()
}

// GetCardonfileType returns the CardonfileType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardChannelProperties) GetCardonfileType() string {
	if o == nil || utils.IsNil(o.CardonfileType.Get()) {
		var ret string
		return ret
	}
	return *o.CardonfileType.Get()
}

// GetCardonfileTypeOk returns a tuple with the CardonfileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardChannelProperties) GetCardonfileTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardonfileType.Get(), o.CardonfileType.IsSet()
}

// HasCardonfileType returns a boolean if a field has been set.
func (o *CardChannelProperties) HasCardonfileType() bool {
	if o != nil && o.CardonfileType.IsSet() {
		return true
	}

	return false
}

// SetCardonfileType gets a reference to the given NullableString and assigns it to the CardonfileType field.
func (o *CardChannelProperties) SetCardonfileType(v string) {
	o.CardonfileType.Set(&v)
}
// SetCardonfileTypeNil sets the value for CardonfileType to be an explicit nil
func (o *CardChannelProperties) SetCardonfileTypeNil() {
	o.CardonfileType.Set(nil)
}

// UnsetCardonfileType ensures that no value is present for CardonfileType, not even an explicit nil
func (o *CardChannelProperties) UnsetCardonfileType() {
	o.CardonfileType.Unset()
}

// GetMerchantIdTag returns the MerchantIdTag field value if set, zero value otherwise.
func (o *CardChannelProperties) GetMerchantIdTag() string {
	if o == nil || utils.IsNil(o.MerchantIdTag) {
		var ret string
		return ret
	}
	return *o.MerchantIdTag
}

// GetMerchantIdTagOk returns a tuple with the MerchantIdTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardChannelProperties) GetMerchantIdTagOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MerchantIdTag) {
		return nil, false
	}
	return o.MerchantIdTag, true
}

// HasMerchantIdTag returns a boolean if a field has been set.
func (o *CardChannelProperties) HasMerchantIdTag() bool {
	if o != nil && !utils.IsNil(o.MerchantIdTag) {
		return true
	}

	return false
}

// SetMerchantIdTag gets a reference to the given string and assigns it to the MerchantIdTag field.
func (o *CardChannelProperties) SetMerchantIdTag(v string) {
	o.MerchantIdTag = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *CardChannelProperties) GetExpiresAt() time.Time {
	if o == nil || utils.IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardChannelProperties) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CardChannelProperties) HasExpiresAt() bool {
	if o != nil && !utils.IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *CardChannelProperties) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetInstallmentConfiguration returns the InstallmentConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardChannelProperties) GetInstallmentConfiguration() CardInstallmentConfiguration {
	if o == nil || utils.IsNil(o.InstallmentConfiguration.Get()) {
		var ret CardInstallmentConfiguration
		return ret
	}
	return *o.InstallmentConfiguration.Get()
}

// GetInstallmentConfigurationOk returns a tuple with the InstallmentConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardChannelProperties) GetInstallmentConfigurationOk() (*CardInstallmentConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstallmentConfiguration.Get(), o.InstallmentConfiguration.IsSet()
}

// HasInstallmentConfiguration returns a boolean if a field has been set.
func (o *CardChannelProperties) HasInstallmentConfiguration() bool {
	if o != nil && o.InstallmentConfiguration.IsSet() {
		return true
	}

	return false
}

// SetInstallmentConfiguration gets a reference to the given NullableCardInstallmentConfiguration and assigns it to the InstallmentConfiguration field.
func (o *CardChannelProperties) SetInstallmentConfiguration(v CardInstallmentConfiguration) {
	o.InstallmentConfiguration.Set(&v)
}
// SetInstallmentConfigurationNil sets the value for InstallmentConfiguration to be an explicit nil
func (o *CardChannelProperties) SetInstallmentConfigurationNil() {
	o.InstallmentConfiguration.Set(nil)
}

// UnsetInstallmentConfiguration ensures that no value is present for InstallmentConfiguration, not even an explicit nil
func (o *CardChannelProperties) UnsetInstallmentConfiguration() {
	o.InstallmentConfiguration.Unset()
}

// GetSkipAuthorization returns the SkipAuthorization field value if set, zero value otherwise.
func (o *CardChannelProperties) GetSkipAuthorization() bool {
	if o == nil || utils.IsNil(o.SkipAuthorization) {
		var ret bool
		return ret
	}
	return *o.SkipAuthorization
}

// GetSkipAuthorizationOk returns a tuple with the SkipAuthorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardChannelProperties) GetSkipAuthorizationOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.SkipAuthorization) {
		return nil, false
	}
	return o.SkipAuthorization, true
}

// HasSkipAuthorization returns a boolean if a field has been set.
func (o *CardChannelProperties) HasSkipAuthorization() bool {
	if o != nil && !utils.IsNil(o.SkipAuthorization) {
		return true
	}

	return false
}

// SetSkipAuthorization gets a reference to the given bool and assigns it to the SkipAuthorization field.
func (o *CardChannelProperties) SetSkipAuthorization(v bool) {
	o.SkipAuthorization = &v
}

func (o CardChannelProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardChannelProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SkipThreeDSecure.IsSet() {
		toSerialize["skip_three_d_secure"] = o.SkipThreeDSecure.Get()
    }
	if o.SuccessReturnUrl.IsSet() {
		toSerialize["success_return_url"] = o.SuccessReturnUrl.Get()
    }
	if o.FailureReturnUrl.IsSet() {
		toSerialize["failure_return_url"] = o.FailureReturnUrl.Get()
    }
	if o.CardonfileType.IsSet() {
		toSerialize["cardonfile_type"] = o.CardonfileType.Get()
    }
	if !utils.IsNil(o.MerchantIdTag) {
		toSerialize["merchant_id_tag"] = o.MerchantIdTag
	}
	if !utils.IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.InstallmentConfiguration.IsSet() {
		toSerialize["installment_configuration"] = o.InstallmentConfiguration.Get()
    }
	if !utils.IsNil(o.SkipAuthorization) {
		toSerialize["skip_authorization"] = o.SkipAuthorization
	}
	return toSerialize, nil
}

type NullableCardChannelProperties struct {
	value *CardChannelProperties
	isSet bool
}

func (v NullableCardChannelProperties) Get() *CardChannelProperties {
	return v.value
}

func (v *NullableCardChannelProperties) Set(val *CardChannelProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableCardChannelProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableCardChannelProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardChannelProperties(val *CardChannelProperties) *NullableCardChannelProperties {
	return &NullableCardChannelProperties{value: val, isSet: true}
}

func (v NullableCardChannelProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardChannelProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


