/*
Payment Requests

This API is used for Payment Requests

API version: 1.59.0
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v5/utils"
)

// checks if the CardInformation type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CardInformation{}

// CardInformation Card Information
type CardInformation struct {
	TokenId *string `json:"token_id,omitempty"`
	// 1st 6 and last 4 digits of the card
	MaskedCardNumber *string `json:"masked_card_number,omitempty"`
	// Card expiry month in MM format
	ExpiryMonth *string `json:"expiry_month,omitempty"`
	// Card expiry month in YY format
	ExpiryYear *string `json:"expiry_year,omitempty"`
	// Cardholder name
	CardholderName NullableString `json:"cardholder_name,omitempty"`
	Fingerprint *string `json:"fingerprint,omitempty"`
	Type *string `json:"type,omitempty"`
	Network *string `json:"network,omitempty"`
	Country *string `json:"country,omitempty"`
	Issuer *string `json:"issuer,omitempty"`
	CardNumber *string `json:"card_number,omitempty"`
	OneTimeToken *string `json:"one_time_token,omitempty"`
}

// NewCardInformation instantiates a new CardInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardInformation() *CardInformation {
	this := CardInformation{}
	return &this
}

// NewCardInformationWithDefaults instantiates a new CardInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardInformationWithDefaults() *CardInformation {
	this := CardInformation{}
	return &this
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *CardInformation) GetTokenId() string {
	if o == nil || utils.IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInformation) GetTokenIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *CardInformation) HasTokenId() bool {
	if o != nil && !utils.IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *CardInformation) SetTokenId(v string) {
	o.TokenId = &v
}

// GetMaskedCardNumber returns the MaskedCardNumber field value if set, zero value otherwise.
func (o *CardInformation) GetMaskedCardNumber() string {
	if o == nil || utils.IsNil(o.MaskedCardNumber) {
		var ret string
		return ret
	}
	return *o.MaskedCardNumber
}

// GetMaskedCardNumberOk returns a tuple with the MaskedCardNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInformation) GetMaskedCardNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MaskedCardNumber) {
		return nil, false
	}
	return o.MaskedCardNumber, true
}

// HasMaskedCardNumber returns a boolean if a field has been set.
func (o *CardInformation) HasMaskedCardNumber() bool {
	if o != nil && !utils.IsNil(o.MaskedCardNumber) {
		return true
	}

	return false
}

// SetMaskedCardNumber gets a reference to the given string and assigns it to the MaskedCardNumber field.
func (o *CardInformation) SetMaskedCardNumber(v string) {
	o.MaskedCardNumber = &v
}

// GetExpiryMonth returns the ExpiryMonth field value if set, zero value otherwise.
func (o *CardInformation) GetExpiryMonth() string {
	if o == nil || utils.IsNil(o.ExpiryMonth) {
		var ret string
		return ret
	}
	return *o.ExpiryMonth
}

// GetExpiryMonthOk returns a tuple with the ExpiryMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInformation) GetExpiryMonthOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExpiryMonth) {
		return nil, false
	}
	return o.ExpiryMonth, true
}

// HasExpiryMonth returns a boolean if a field has been set.
func (o *CardInformation) HasExpiryMonth() bool {
	if o != nil && !utils.IsNil(o.ExpiryMonth) {
		return true
	}

	return false
}

// SetExpiryMonth gets a reference to the given string and assigns it to the ExpiryMonth field.
func (o *CardInformation) SetExpiryMonth(v string) {
	o.ExpiryMonth = &v
}

// GetExpiryYear returns the ExpiryYear field value if set, zero value otherwise.
func (o *CardInformation) GetExpiryYear() string {
	if o == nil || utils.IsNil(o.ExpiryYear) {
		var ret string
		return ret
	}
	return *o.ExpiryYear
}

// GetExpiryYearOk returns a tuple with the ExpiryYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInformation) GetExpiryYearOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExpiryYear) {
		return nil, false
	}
	return o.ExpiryYear, true
}

// HasExpiryYear returns a boolean if a field has been set.
func (o *CardInformation) HasExpiryYear() bool {
	if o != nil && !utils.IsNil(o.ExpiryYear) {
		return true
	}

	return false
}

// SetExpiryYear gets a reference to the given string and assigns it to the ExpiryYear field.
func (o *CardInformation) SetExpiryYear(v string) {
	o.ExpiryYear = &v
}

// GetCardholderName returns the CardholderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardInformation) GetCardholderName() string {
	if o == nil || utils.IsNil(o.CardholderName.Get()) {
		var ret string
		return ret
	}
	return *o.CardholderName.Get()
}

// GetCardholderNameOk returns a tuple with the CardholderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardInformation) GetCardholderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardholderName.Get(), o.CardholderName.IsSet()
}

// HasCardholderName returns a boolean if a field has been set.
func (o *CardInformation) HasCardholderName() bool {
	if o != nil && o.CardholderName.IsSet() {
		return true
	}

	return false
}

// SetCardholderName gets a reference to the given NullableString and assigns it to the CardholderName field.
func (o *CardInformation) SetCardholderName(v string) {
	o.CardholderName.Set(&v)
}
// SetCardholderNameNil sets the value for CardholderName to be an explicit nil
func (o *CardInformation) SetCardholderNameNil() {
	o.CardholderName.Set(nil)
}

// UnsetCardholderName ensures that no value is present for CardholderName, not even an explicit nil
func (o *CardInformation) UnsetCardholderName() {
	o.CardholderName.Unset()
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *CardInformation) GetFingerprint() string {
	if o == nil || utils.IsNil(o.Fingerprint) {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInformation) GetFingerprintOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Fingerprint) {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *CardInformation) HasFingerprint() bool {
	if o != nil && !utils.IsNil(o.Fingerprint) {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *CardInformation) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CardInformation) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInformation) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CardInformation) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CardInformation) SetType(v string) {
	o.Type = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *CardInformation) GetNetwork() string {
	if o == nil || utils.IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInformation) GetNetworkOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *CardInformation) HasNetwork() bool {
	if o != nil && !utils.IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *CardInformation) SetNetwork(v string) {
	o.Network = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CardInformation) GetCountry() string {
	if o == nil || utils.IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInformation) GetCountryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CardInformation) HasCountry() bool {
	if o != nil && !utils.IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CardInformation) SetCountry(v string) {
	o.Country = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *CardInformation) GetIssuer() string {
	if o == nil || utils.IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInformation) GetIssuerOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *CardInformation) HasIssuer() bool {
	if o != nil && !utils.IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *CardInformation) SetIssuer(v string) {
	o.Issuer = &v
}

// GetCardNumber returns the CardNumber field value if set, zero value otherwise.
func (o *CardInformation) GetCardNumber() string {
	if o == nil || utils.IsNil(o.CardNumber) {
		var ret string
		return ret
	}
	return *o.CardNumber
}

// GetCardNumberOk returns a tuple with the CardNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInformation) GetCardNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CardNumber) {
		return nil, false
	}
	return o.CardNumber, true
}

// HasCardNumber returns a boolean if a field has been set.
func (o *CardInformation) HasCardNumber() bool {
	if o != nil && !utils.IsNil(o.CardNumber) {
		return true
	}

	return false
}

// SetCardNumber gets a reference to the given string and assigns it to the CardNumber field.
func (o *CardInformation) SetCardNumber(v string) {
	o.CardNumber = &v
}

// GetOneTimeToken returns the OneTimeToken field value if set, zero value otherwise.
func (o *CardInformation) GetOneTimeToken() string {
	if o == nil || utils.IsNil(o.OneTimeToken) {
		var ret string
		return ret
	}
	return *o.OneTimeToken
}

// GetOneTimeTokenOk returns a tuple with the OneTimeToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInformation) GetOneTimeTokenOk() (*string, bool) {
	if o == nil || utils.IsNil(o.OneTimeToken) {
		return nil, false
	}
	return o.OneTimeToken, true
}

// HasOneTimeToken returns a boolean if a field has been set.
func (o *CardInformation) HasOneTimeToken() bool {
	if o != nil && !utils.IsNil(o.OneTimeToken) {
		return true
	}

	return false
}

// SetOneTimeToken gets a reference to the given string and assigns it to the OneTimeToken field.
func (o *CardInformation) SetOneTimeToken(v string) {
	o.OneTimeToken = &v
}

func (o CardInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
	}
	if !utils.IsNil(o.MaskedCardNumber) {
		toSerialize["masked_card_number"] = o.MaskedCardNumber
	}
	if !utils.IsNil(o.ExpiryMonth) {
		toSerialize["expiry_month"] = o.ExpiryMonth
	}
	if !utils.IsNil(o.ExpiryYear) {
		toSerialize["expiry_year"] = o.ExpiryYear
	}
	if o.CardholderName.IsSet() {
		toSerialize["cardholder_name"] = o.CardholderName.Get()
    }
	if !utils.IsNil(o.Fingerprint) {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !utils.IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !utils.IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !utils.IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !utils.IsNil(o.CardNumber) {
		toSerialize["card_number"] = o.CardNumber
	}
	if !utils.IsNil(o.OneTimeToken) {
		toSerialize["one_time_token"] = o.OneTimeToken
	}
	return toSerialize, nil
}

type NullableCardInformation struct {
	value *CardInformation
	isSet bool
}

func (v NullableCardInformation) Get() *CardInformation {
	return v.value
}

func (v *NullableCardInformation) Set(val *CardInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableCardInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableCardInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardInformation(val *CardInformation) *NullableCardInformation {
	return &NullableCardInformation{value: val, isSet: true}
}

func (v NullableCardInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


