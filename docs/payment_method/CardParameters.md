# CardParameters

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **Currency** | **string** |  |  |
| **ChannelProperties** | Pointer to [**NullableCardChannelProperties**](CardChannelProperties.md) |  | [optional]  |
| **CardInformation** | Pointer to [**CardParametersCardInformation**](CardParametersCardInformation.md) |  | [optional]  |

## Methods

### NewCardParameters

`func NewCardParameters(currency string, ) *CardParameters`

NewCardParameters instantiates a new CardParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardParametersWithDefaults

`func NewCardParametersWithDefaults() *CardParameters`

NewCardParametersWithDefaults instantiates a new CardParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *CardParameters) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CardParameters) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CardParameters) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetChannelProperties

`func (o *CardParameters) GetChannelProperties() CardChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *CardParameters) GetChannelPropertiesOk() (*CardChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *CardParameters) SetChannelProperties(v CardChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *CardParameters) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.

### SetChannelPropertiesNil

`func (o *CardParameters) SetChannelPropertiesNil(b bool)`

 SetChannelPropertiesNil sets the value for ChannelProperties to be an explicit nil

### UnsetChannelProperties
`func (o *CardParameters) UnsetChannelProperties()`

UnsetChannelProperties ensures that no value is present for ChannelProperties, not even an explicit nil
### GetCardInformation

`func (o *CardParameters) GetCardInformation() CardParametersCardInformation`

GetCardInformation returns the CardInformation field if non-nil, zero value otherwise.

### GetCardInformationOk

`func (o *CardParameters) GetCardInformationOk() (*CardParametersCardInformation, bool)`

GetCardInformationOk returns a tuple with the CardInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardInformation

`func (o *CardParameters) SetCardInformation(v CardParametersCardInformation)`

SetCardInformation sets CardInformation field to given value.

### HasCardInformation

`func (o *CardParameters) HasCardInformation() bool`

HasCardInformation returns a boolean if a field has been set.


[[Back to README]](../../README.md)


