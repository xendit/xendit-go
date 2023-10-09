/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.89.1
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the TokenizedCardInformation type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TokenizedCardInformation{}

// TokenizedCardInformation Tokenized Card Information
type TokenizedCardInformation struct {
	TokenId string `json:"token_id"`
	// 1st 6 and last 4 digits of the card
	MaskedCardNumber string `json:"masked_card_number"`
	// Cardholder name is optional but recommended for 3DS 2 / AVS verification
	CardholderName NullableString `json:"cardholder_name,omitempty"`
	// Card expiry month in MM format
	ExpiryMonth string `json:"expiry_month"`
	// Card expiry month in YY format
	ExpiryYear string `json:"expiry_year"`
	// Xendit-generated identifier for the unique card number. Multiple payment method objects can be created for the same account - e.g. if the user first creates a one-time payment request, and then later on creates a multiple-use payment method using the same account.   The fingerprint helps to identify the unique account being used.
	Fingerprint string `json:"fingerprint"`
	// Whether the card is a credit or debit card
	Type string `json:"type"`
	// Card network - VISA, MASTERCARD, JCB, AMEX, DISCOVER, BCA
	Network string `json:"network"`
	// Country where the card was issued ISO 3166-1 Alpha-2
	Country string `json:"country"`
	// Issuer of the card, most often an issuing bank For example, “BCA”, “MANDIRI”
	Issuer string `json:"issuer"`
}

// NewTokenizedCardInformation instantiates a new TokenizedCardInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenizedCardInformation(tokenId string, maskedCardNumber string, expiryMonth string, expiryYear string, fingerprint string, type_ string, network string, country string, issuer string) *TokenizedCardInformation {
	this := TokenizedCardInformation{}
	this.TokenId = tokenId
	this.MaskedCardNumber = maskedCardNumber
	this.ExpiryMonth = expiryMonth
	this.ExpiryYear = expiryYear
	this.Fingerprint = fingerprint
	this.Type = type_
	this.Network = network
	this.Country = country
	this.Issuer = issuer
	return &this
}

// NewTokenizedCardInformationWithDefaults instantiates a new TokenizedCardInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenizedCardInformationWithDefaults() *TokenizedCardInformation {
	this := TokenizedCardInformation{}
	return &this
}

// GetTokenId returns the TokenId field value
func (o *TokenizedCardInformation) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *TokenizedCardInformation) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *TokenizedCardInformation) SetTokenId(v string) {
	o.TokenId = v
}

// GetMaskedCardNumber returns the MaskedCardNumber field value
func (o *TokenizedCardInformation) GetMaskedCardNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaskedCardNumber
}

// GetMaskedCardNumberOk returns a tuple with the MaskedCardNumber field value
// and a boolean to check if the value has been set.
func (o *TokenizedCardInformation) GetMaskedCardNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaskedCardNumber, true
}

// SetMaskedCardNumber sets field value
func (o *TokenizedCardInformation) SetMaskedCardNumber(v string) {
	o.MaskedCardNumber = v
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

// GetExpiryMonth returns the ExpiryMonth field value
func (o *TokenizedCardInformation) GetExpiryMonth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpiryMonth
}

// GetExpiryMonthOk returns a tuple with the ExpiryMonth field value
// and a boolean to check if the value has been set.
func (o *TokenizedCardInformation) GetExpiryMonthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiryMonth, true
}

// SetExpiryMonth sets field value
func (o *TokenizedCardInformation) SetExpiryMonth(v string) {
	o.ExpiryMonth = v
}

// GetExpiryYear returns the ExpiryYear field value
func (o *TokenizedCardInformation) GetExpiryYear() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpiryYear
}

// GetExpiryYearOk returns a tuple with the ExpiryYear field value
// and a boolean to check if the value has been set.
func (o *TokenizedCardInformation) GetExpiryYearOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiryYear, true
}

// SetExpiryYear sets field value
func (o *TokenizedCardInformation) SetExpiryYear(v string) {
	o.ExpiryYear = v
}

// GetFingerprint returns the Fingerprint field value
func (o *TokenizedCardInformation) GetFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
func (o *TokenizedCardInformation) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fingerprint, true
}

// SetFingerprint sets field value
func (o *TokenizedCardInformation) SetFingerprint(v string) {
	o.Fingerprint = v
}

// GetType returns the Type field value
func (o *TokenizedCardInformation) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenizedCardInformation) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenizedCardInformation) SetType(v string) {
	o.Type = v
}

// GetNetwork returns the Network field value
func (o *TokenizedCardInformation) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *TokenizedCardInformation) GetNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *TokenizedCardInformation) SetNetwork(v string) {
	o.Network = v
}

// GetCountry returns the Country field value
func (o *TokenizedCardInformation) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *TokenizedCardInformation) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *TokenizedCardInformation) SetCountry(v string) {
	o.Country = v
}

// GetIssuer returns the Issuer field value
func (o *TokenizedCardInformation) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *TokenizedCardInformation) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *TokenizedCardInformation) SetIssuer(v string) {
	o.Issuer = v
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
	toSerialize["token_id"] = o.TokenId
	toSerialize["masked_card_number"] = o.MaskedCardNumber
	if o.CardholderName.IsSet() {
		toSerialize["cardholder_name"] = o.CardholderName.Get()
    }
	toSerialize["expiry_month"] = o.ExpiryMonth
	toSerialize["expiry_year"] = o.ExpiryYear
	toSerialize["fingerprint"] = o.Fingerprint
	toSerialize["type"] = o.Type
	toSerialize["network"] = o.Network
	toSerialize["country"] = o.Country
	toSerialize["issuer"] = o.Issuer
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


