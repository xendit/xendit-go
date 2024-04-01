# TokenizedCardInformation
Tokenized Card Information

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **TokenId** | Pointer to **string** |  |  |  |
| **MaskedCardNumber** | Pointer to **string** |  | 1st 6 and last 4 digits of the card |  |
| **CardholderName** | Pointer to **NullableString** |  | Cardholder name is optional but recommended for 3DS 2 / AVS verification |  |
| **ExpiryMonth** | Pointer to **string** |  | Card expiry month in MM format |  |
| **ExpiryYear** | Pointer to **string** |  | Card expiry month in YY format |  |
| **Fingerprint** | Pointer to **string** |  | Xendit-generated identifier for the unique card number. Multiple payment method objects can be created for the same account - e.g. if the user first creates a one-time payment request, and then later on creates a multiple-use payment method using the same account.   The fingerprint helps to identify the unique account being used. |  |
| **Type** | Pointer to **string** |  | Whether the card is a credit or debit card |  |
| **Network** | Pointer to **string** |  | Card network - VISA, MASTERCARD, JCB, AMEX, DISCOVER, BCA |  |
| **Country** | Pointer to **string** |  | Country where the card was issued ISO 3166-1 Alpha-2 |  |
| **Issuer** | Pointer to **string** |  | Issuer of the card, most often an issuing bank For example, “BCA”, “MANDIRI” |  |
| **CardNumber** | Pointer to **string** |  |  |  |
| **OneTimeToken** | Pointer to **string** |  |  |  |

## Methods

### NewTokenizedCardInformation

`func NewTokenizedCardInformation() *TokenizedCardInformation`

NewTokenizedCardInformation instantiates a new TokenizedCardInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizedCardInformationWithDefaults

`func NewTokenizedCardInformationWithDefaults() *TokenizedCardInformation`

NewTokenizedCardInformationWithDefaults instantiates a new TokenizedCardInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *TokenizedCardInformation) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenizedCardInformation) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenizedCardInformation) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *TokenizedCardInformation) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetMaskedCardNumber

`func (o *TokenizedCardInformation) GetMaskedCardNumber() string`

GetMaskedCardNumber returns the MaskedCardNumber field if non-nil, zero value otherwise.

### GetMaskedCardNumberOk

`func (o *TokenizedCardInformation) GetMaskedCardNumberOk() (*string, bool)`

GetMaskedCardNumberOk returns a tuple with the MaskedCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedCardNumber

`func (o *TokenizedCardInformation) SetMaskedCardNumber(v string)`

SetMaskedCardNumber sets MaskedCardNumber field to given value.

### HasMaskedCardNumber

`func (o *TokenizedCardInformation) HasMaskedCardNumber() bool`

HasMaskedCardNumber returns a boolean if a field has been set.

### GetCardholderName

`func (o *TokenizedCardInformation) GetCardholderName() string`

GetCardholderName returns the CardholderName field if non-nil, zero value otherwise.

### GetCardholderNameOk

`func (o *TokenizedCardInformation) GetCardholderNameOk() (*string, bool)`

GetCardholderNameOk returns a tuple with the CardholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderName

`func (o *TokenizedCardInformation) SetCardholderName(v string)`

SetCardholderName sets CardholderName field to given value.

### HasCardholderName

`func (o *TokenizedCardInformation) HasCardholderName() bool`

HasCardholderName returns a boolean if a field has been set.

### SetCardholderNameNil

`func (o *TokenizedCardInformation) SetCardholderNameNil(b bool)`

 SetCardholderNameNil sets the value for CardholderName to be an explicit nil

### UnsetCardholderName
`func (o *TokenizedCardInformation) UnsetCardholderName()`

UnsetCardholderName ensures that no value is present for CardholderName, not even an explicit nil
### GetExpiryMonth

`func (o *TokenizedCardInformation) GetExpiryMonth() string`

GetExpiryMonth returns the ExpiryMonth field if non-nil, zero value otherwise.

### GetExpiryMonthOk

`func (o *TokenizedCardInformation) GetExpiryMonthOk() (*string, bool)`

GetExpiryMonthOk returns a tuple with the ExpiryMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryMonth

`func (o *TokenizedCardInformation) SetExpiryMonth(v string)`

SetExpiryMonth sets ExpiryMonth field to given value.

### HasExpiryMonth

`func (o *TokenizedCardInformation) HasExpiryMonth() bool`

HasExpiryMonth returns a boolean if a field has been set.

### GetExpiryYear

`func (o *TokenizedCardInformation) GetExpiryYear() string`

GetExpiryYear returns the ExpiryYear field if non-nil, zero value otherwise.

### GetExpiryYearOk

`func (o *TokenizedCardInformation) GetExpiryYearOk() (*string, bool)`

GetExpiryYearOk returns a tuple with the ExpiryYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryYear

`func (o *TokenizedCardInformation) SetExpiryYear(v string)`

SetExpiryYear sets ExpiryYear field to given value.

### HasExpiryYear

`func (o *TokenizedCardInformation) HasExpiryYear() bool`

HasExpiryYear returns a boolean if a field has been set.

### GetFingerprint

`func (o *TokenizedCardInformation) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *TokenizedCardInformation) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *TokenizedCardInformation) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *TokenizedCardInformation) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetType

`func (o *TokenizedCardInformation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenizedCardInformation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenizedCardInformation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TokenizedCardInformation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNetwork

`func (o *TokenizedCardInformation) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *TokenizedCardInformation) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *TokenizedCardInformation) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *TokenizedCardInformation) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCountry

`func (o *TokenizedCardInformation) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *TokenizedCardInformation) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *TokenizedCardInformation) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *TokenizedCardInformation) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetIssuer

`func (o *TokenizedCardInformation) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *TokenizedCardInformation) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *TokenizedCardInformation) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *TokenizedCardInformation) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetCardNumber

`func (o *TokenizedCardInformation) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *TokenizedCardInformation) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *TokenizedCardInformation) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *TokenizedCardInformation) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetOneTimeToken

`func (o *TokenizedCardInformation) GetOneTimeToken() string`

GetOneTimeToken returns the OneTimeToken field if non-nil, zero value otherwise.

### GetOneTimeTokenOk

`func (o *TokenizedCardInformation) GetOneTimeTokenOk() (*string, bool)`

GetOneTimeTokenOk returns a tuple with the OneTimeToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTimeToken

`func (o *TokenizedCardInformation) SetOneTimeToken(v string)`

SetOneTimeToken sets OneTimeToken field to given value.

### HasOneTimeToken

`func (o *TokenizedCardInformation) HasOneTimeToken() bool`

HasOneTimeToken returns a boolean if a field has been set.


[[Back to README]](../../README.md)


