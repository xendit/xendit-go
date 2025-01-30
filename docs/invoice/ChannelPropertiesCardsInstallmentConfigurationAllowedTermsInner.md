# ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Issuer** | Pointer to **string** |  | The bank code of the installment provider / issuer |  |
| **AllowedTerms** | Pointer to **float32[]** |  | An array containing list of installment tenor available to choose |  |

## Methods

### NewChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner

`func NewChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner() *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner`

NewChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner instantiates a new ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelPropertiesCardsInstallmentConfigurationAllowedTermsInnerWithDefaults

`func NewChannelPropertiesCardsInstallmentConfigurationAllowedTermsInnerWithDefaults() *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner`

NewChannelPropertiesCardsInstallmentConfigurationAllowedTermsInnerWithDefaults instantiates a new ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetAllowedTerms

`func (o *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) GetAllowedTerms() []float32`

GetAllowedTerms returns the AllowedTerms field if non-nil, zero value otherwise.

### GetAllowedTermsOk

`func (o *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) GetAllowedTermsOk() (*[]float32, bool)`

GetAllowedTermsOk returns a tuple with the AllowedTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTerms

`func (o *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) SetAllowedTerms(v []float32)`

SetAllowedTerms sets AllowedTerms field to given value.

### HasAllowedTerms

`func (o *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) HasAllowedTerms() bool`

HasAllowedTerms returns a boolean if a field has been set.


[[Back to README]](../../README.md)


