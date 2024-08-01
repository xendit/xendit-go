# CardInstallmentConfiguration
Card Installment Configuration

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Terms** | Pointer to **int32** |  |  |  |
| **Interval** | Pointer to **string** |  |  |  |
| **Code** | Pointer to **NullableString** |  |  |  |

## Methods

### NewCardInstallmentConfiguration

`func NewCardInstallmentConfiguration() *CardInstallmentConfiguration`

NewCardInstallmentConfiguration instantiates a new CardInstallmentConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardInstallmentConfigurationWithDefaults

`func NewCardInstallmentConfigurationWithDefaults() *CardInstallmentConfiguration`

NewCardInstallmentConfigurationWithDefaults instantiates a new CardInstallmentConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTerms

`func (o *CardInstallmentConfiguration) GetTerms() int32`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *CardInstallmentConfiguration) GetTermsOk() (*int32, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *CardInstallmentConfiguration) SetTerms(v int32)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *CardInstallmentConfiguration) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetInterval

`func (o *CardInstallmentConfiguration) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CardInstallmentConfiguration) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CardInstallmentConfiguration) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CardInstallmentConfiguration) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetCode

`func (o *CardInstallmentConfiguration) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CardInstallmentConfiguration) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CardInstallmentConfiguration) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CardInstallmentConfiguration) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *CardInstallmentConfiguration) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *CardInstallmentConfiguration) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil

[[Back to README]](../../README.md)


