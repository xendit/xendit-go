# ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Issuer** | Pointer to **string** |  | The bank code of the installment provider / issuer |  |
| **Terms** | Pointer to **float32[]** |  | An array containing list of installment tenor available to choose |  |

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

### GetTerms

`func (o *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) GetTerms() []float32`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) GetTermsOk() (*[]float32, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) SetTerms(v []float32)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner) HasTerms() bool`

HasTerms returns a boolean if a field has been set.


[[Back to README]](../../README.md)


