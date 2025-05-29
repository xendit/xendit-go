/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.128.0
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
)

// checks if the TokenizedCardInformation type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TokenizedCardInformation{}

// TokenizedCardInformation Tokenized Card Information
type TokenizedCardInformation struct {
	TokenId *string `json:"token_id,omitempty"`
	// 1st 6 and last 4 digits of the card
	MaskedCardNumber *string `json:"masked_card_number,omitempty"`
	// Cardholder name is optional but recommended for 3DS 2 / AVS verification
	CardholderName NullableString `json:"cardholder_name,omitempty"`
	// Card expiry month in MM format
	ExpiryMonth *string `json:"expiry_month,omitempty"`
	// Card expiry month in YY format
	ExpiryYear *string `json:"expiry_year,omitempty"`
	// Xendit-generated identifier for the unique card number. Multiple payment method objects can be created for the same account - e.g. if the user first creates a one-time payment request, and then later on creates a multiple-use payment method using the same account.   The fingerprint helps to identify the unique account being used.
	Fingerprint *string `json:"fingerprint,omitempty"`
	// Whether the card is a credit or debit card
	Type *string `json:"type,omitempty"`
	// Card network - VISA, MASTERCARD, JCB, AMEX, DISCOVER, BCA
	Network *string `json:"network,omitempty"`
	// Country where the card was issued ISO 3166-1 Alpha-2
	Country *string `json:"country,omitempty"`
	// Issuer of the card, most often an issuing bank For example, “BCA”, “MANDIRI”
	Issuer *string `json:"issuer,omitempty"`
	CardNumber *string `json:"card_number,omitempty"`
	OneTimeToken *string `json:"one_time_token,omitempty"`
}

// NewTokenizedCardInformation instantiates a new TokenizedCardInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenizedCardInformation() *TokenizedCardInformation {
	this := TokenizedCardInformation{}
	return &this
}

// NewTokenizedCardInformationWithDefaults instantiates a new TokenizedCardInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenizedCardInformationWithDefaults() *TokenizedCardInformation {
	this := TokenizedCardInformation{}
	return &this
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *TokenizedCardInformation) GetTokenId() string {
	if o == nil || utils.IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedCardInformation) GetTokenIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *TokenizedCardInformation) HasTokenId() bool {
	if o != nil && !utils.IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *TokenizedCardInformation) SetTokenId(v string) {
	o.TokenId = &v
}

// GetMaskedCardNumber returns the MaskedCardNumber field value if set, zero value otherwise.
func (o *TokenizedCardInformation) GetMaskedCardNumber() string {
	if o == nil || utils.IsNil(o.MaskedCardNumber) {
		var ret string
		return ret
	}
	return *o.MaskedCardNumber
}

// GetMaskedCardNumberOk returns a tuple with the MaskedCardNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedCardInformation) GetMaskedCardNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MaskedCardNumber) {
		return nil, false
	}
	return o.MaskedCardNumber, true
}

// HasMaskedCardNumber returns a boolean if a field has been set.
func (o *TokenizedCardInformation) HasMaskedCardNumber() bool {
	if o != nil && !utils.IsNil(o.MaskedCardNumber) {
		return true
	}

	return false
}

// SetMaskedCardNumber gets a reference to the given string and assigns it to the MaskedCardNumber field.
func (o *TokenizedCardInformation) SetMaskedCardNumber(v string) {
	o.MaskedCardNumber = &v
}

// GetCardholderName returns the CardholderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TokenizedCardInformation) GetCardholderName() string {
	if o == nil || utils.IsNil(o.CardholderName.Get()) {
		var ret string
		return ret
	}
	return *o.CardholderName.Get()
}

// GetCardholderNameOk returns a tuple with the CardholderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenizedCardInformation) GetCardholderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardholderName.Get(), o.CardholderName.IsSet()
}

// HasCardholderName returns a boolean if a field has been set.
func (o *TokenizedCardInformation) HasCardholderName() bool {
	if o != nil && o.CardholderName.IsSet() {
		return true
	}

	return false
}

// SetCardholderName gets a reference to the given NullableString and assigns it to the CardholderName field.
func (o *TokenizedCardInformation) SetCardholderName(v string) {
	o.CardholderName.Set(&v)
}
// SetCardholderNameNil sets the value for CardholderName to be an explicit nil
func (o *TokenizedCardInformation) SetCardholderNameNil() {
	o.CardholderName.Set(nil)
}

// UnsetCardholderName ensures that no value is present for CardholderName, not even an explicit nil
func (o *TokenizedCardInformation) UnsetCardholderName() {
	o.CardholderName.Unset()
}

// GetExpiryMonth returns the ExpiryMonth field value if set, zero value otherwise.
func (o *TokenizedCardInformation) GetExpiryMonth() string {
	if o == nil || utils.IsNil(o.ExpiryMonth) {
		var ret string
		return ret
	}
	return *o.ExpiryMonth
}

// GetExpiryMonthOk returns a tuple with the ExpiryMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedCardInformation) GetExpiryMonthOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExpiryMonth) {
		return nil, false
	}
	return o.ExpiryMonth, true
}

// HasExpiryMonth returns a boolean if a field has been set.
func (o *TokenizedCardInformation) HasExpiryMonth() bool {
	if o != nil && !utils.IsNil(o.ExpiryMonth) {
		return true
	}

	return false
}

// SetExpiryMonth gets a reference to the given string and assigns it to the ExpiryMonth field.
func (o *TokenizedCardInformation) SetExpiryMonth(v string) {
	o.ExpiryMonth = &v
}

// GetExpiryYear returns the ExpiryYear field value if set, zero value otherwise.
func (o *TokenizedCardInformation) GetExpiryYear() string {
	if o == nil || utils.IsNil(o.ExpiryYear) {
		var ret string
		return ret
	}
	return *o.ExpiryYear
}

// GetExpiryYearOk returns a tuple with the ExpiryYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedCardInformation) GetExpiryYearOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExpiryYear) {
		return nil, false
	}
	return o.ExpiryYear, true
}

// HasExpiryYear returns a boolean if a field has been set.
func (o *TokenizedCardInformation) HasExpiryYear() bool {
	if o != nil && !utils.IsNil(o.ExpiryYear) {
		return true
	}

	return false
}

// SetExpiryYear gets a reference to the given string and assigns it to the ExpiryYear field.
func (o *TokenizedCardInformation) SetExpiryYear(v string) {
	o.ExpiryYear = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *TokenizedCardInformation) GetFingerprint() string {
	if o == nil || utils.IsNil(o.Fingerprint) {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedCardInformation) GetFingerprintOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Fingerprint) {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *TokenizedCardInformation) HasFingerprint() bool {
	if o != nil && !utils.IsNil(o.Fingerprint) {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *TokenizedCardInformation) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TokenizedCardInformation) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedCardInformation) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TokenizedCardInformation) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TokenizedCardInformation) SetType(v string) {
	o.Type = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *TokenizedCardInformation) GetNetwork() string {
	if o == nil || utils.IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedCardInformation) GetNetworkOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *TokenizedCardInformation) HasNetwork() bool {
	if o != nil && !utils.IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *TokenizedCardInformation) SetNetwork(v string) {
	o.Network = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *TokenizedCardInformation) GetCountry() string {
	if o == nil || utils.IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedCardInformation) GetCountryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *TokenizedCardInformation) HasCountry() bool {
	if o != nil && !utils.IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *TokenizedCardInformation) SetCountry(v string) {
	o.Country = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *TokenizedCardInformation) GetIssuer() string {
	if o == nil || utils.IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedCardInformation) GetIssuerOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *TokenizedCardInformation) HasIssuer() bool {
	if o != nil && !utils.IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *TokenizedCardInformation) SetIssuer(v string) {
	o.Issuer = &v
}

// GetCardNumber returns the CardNumber field value if set, zero value otherwise.
func (o *TokenizedCardInformation) GetCardNumber() string {
	if o == nil || utils.IsNil(o.CardNumber) {
		var ret string
		return ret
	}
	return *o.CardNumber
}

// GetCardNumberOk returns a tuple with the CardNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedCardInformation) GetCardNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CardNumber) {
		return nil, false
	}
	return o.CardNumber, true
}

// HasCardNumber returns a boolean if a field has been set.
func (o *TokenizedCardInformation) HasCardNumber() bool {
	if o != nil && !utils.IsNil(o.CardNumber) {
		return true
	}

	return false
}

// SetCardNumber gets a reference to the given string and assigns it to the CardNumber field.
func (o *TokenizedCardInformation) SetCardNumber(v string) {
	o.CardNumber = &v
}

// GetOneTimeToken returns the OneTimeToken field value if set, zero value otherwise.
func (o *TokenizedCardInformation) GetOneTimeToken() string {
	if o == nil || utils.IsNil(o.OneTimeToken) {
		var ret string
		return ret
	}
	return *o.OneTimeToken
}

// GetOneTimeTokenOk returns a tuple with the OneTimeToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizedCardInformation) GetOneTimeTokenOk() (*string, bool) {
	if o == nil || utils.IsNil(o.OneTimeToken) {
		return nil, false
	}
	return o.OneTimeToken, true
}

// HasOneTimeToken returns a boolean if a field has been set.
func (o *TokenizedCardInformation) HasOneTimeToken() bool {
	if o != nil && !utils.IsNil(o.OneTimeToken) {
		return true
	}

	return false
}

// SetOneTimeToken gets a reference to the given string and assigns it to the OneTimeToken field.
func (o *TokenizedCardInformation) SetOneTimeToken(v string) {
	o.OneTimeToken = &v
}

func (o TokenizedCardInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenizedCardInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
	}
	if !utils.IsNil(o.MaskedCardNumber) {
		toSerialize["masked_card_number"] = o.MaskedCardNumber
	}
	if o.CardholderName.IsSet() {
		toSerialize["cardholder_name"] = o.CardholderName.Get()
    }
	if !utils.IsNil(o.ExpiryMonth) {
		toSerialize["expiry_month"] = o.ExpiryMonth
	}
	if !utils.IsNil(o.ExpiryYear) {
		toSerialize["expiry_year"] = o.ExpiryYear
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

type NullableTokenizedCardInformation struct {
	value *TokenizedCardInformation
	isSet bool
}

func (v NullableTokenizedCardInformation) Get() *TokenizedCardInformation {
	return v.value
}

func (v *NullableTokenizedCardInformation) Set(val *TokenizedCardInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenizedCardInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenizedCardInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenizedCardInformation(val *TokenizedCardInformation) *NullableTokenizedCardInformation {
	return &NullableTokenizedCardInformation{value: val, isSet: true}
}

func (v NullableTokenizedCardInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenizedCardInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


