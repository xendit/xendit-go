# ChannelPropertiesCardsInstallmentConfiguration
An object to pre-set cards installment for a specific invoice

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **AllowFullPayment** | Pointer to **bool** |  | Indicate whether full payment (without installment) is allowed |  |
| **AllowedTerms** | Pointer to [**ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner[]**](ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner.md) |  | An object to set what kind (from specific bank / specific tenor) of cards installments will be available on a specific invoice |  |

## Methods

### NewChannelPropertiesCardsInstallmentConfiguration

`func NewChannelPropertiesCardsInstallmentConfiguration() *ChannelPropertiesCardsInstallmentConfiguration`

NewChannelPropertiesCardsInstallmentConfiguration instantiates a new ChannelPropertiesCardsInstallmentConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelPropertiesCardsInstallmentConfigurationWithDefaults

`func NewChannelPropertiesCardsInstallmentConfigurationWithDefaults() *ChannelPropertiesCardsInstallmentConfiguration`

NewChannelPropertiesCardsInstallmentConfigurationWithDefaults instantiates a new ChannelPropertiesCardsInstallmentConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowFullPayment

`func (o *ChannelPropertiesCardsInstallmentConfiguration) GetAllowFullPayment() bool`

GetAllowFullPayment returns the AllowFullPayment field if non-nil, zero value otherwise.

### GetAllowFullPaymentOk

`func (o *ChannelPropertiesCardsInstallmentConfiguration) GetAllowFullPaymentOk() (*bool, bool)`

GetAllowFullPaymentOk returns a tuple with the AllowFullPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFullPayment

`func (o *ChannelPropertiesCardsInstallmentConfiguration) SetAllowFullPayment(v bool)`

SetAllowFullPayment sets AllowFullPayment field to given value.

### HasAllowFullPayment

`func (o *ChannelPropertiesCardsInstallmentConfiguration) HasAllowFullPayment() bool`

HasAllowFullPayment returns a boolean if a field has been set.

### GetAllowedTerms

`func (o *ChannelPropertiesCardsInstallmentConfiguration) GetAllowedTerms() []ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner`

GetAllowedTerms returns the AllowedTerms field if non-nil, zero value otherwise.

### GetAllowedTermsOk

`func (o *ChannelPropertiesCardsInstallmentConfiguration) GetAllowedTermsOk() (*[]ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner, bool)`

GetAllowedTermsOk returns a tuple with the AllowedTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTerms

`func (o *ChannelPropertiesCardsInstallmentConfiguration) SetAllowedTerms(v []ChannelPropertiesCardsInstallmentConfigurationAllowedTermsInner)`

SetAllowedTerms sets AllowedTerms field to given value.

### HasAllowedTerms

`func (o *ChannelPropertiesCardsInstallmentConfiguration) HasAllowedTerms() bool`

HasAllowedTerms returns a boolean if a field has been set.


[[Back to README]](../../README.md)


