# Card

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **Currency** | **NullableString** |  |  |
| **ChannelProperties** | [**NullableCardChannelProperties**](CardChannelProperties.md) |  |  |
| **CardInformation** | Pointer to [**TokenizedCardInformation**](TokenizedCardInformation.md) |  | [optional]  |
| **CardVerificationResults** | Pointer to [**NullableCardVerificationResults**](CardVerificationResults.md) |  | [optional]  |

## Methods

### NewCard

`func NewCard(currency NullableString, channelProperties NullableCardChannelProperties, ) *Card`

NewCard instantiates a new Card object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardWithDefaults

`func NewCardWithDefaults() *Card`

NewCardWithDefaults instantiates a new Card object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *Card) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Card) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Card) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *Card) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *Card) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetChannelProperties

`func (o *Card) GetChannelProperties() CardChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *Card) GetChannelPropertiesOk() (*CardChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *Card) SetChannelProperties(v CardChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.


### SetChannelPropertiesNil

`func (o *Card) SetChannelPropertiesNil(b bool)`

 SetChannelPropertiesNil sets the value for ChannelProperties to be an explicit nil

### UnsetChannelProperties
`func (o *Card) UnsetChannelProperties()`

UnsetChannelProperties ensures that no value is present for ChannelProperties, not even an explicit nil
### GetCardInformation

`func (o *Card) GetCardInformation() TokenizedCardInformation`

GetCardInformation returns the CardInformation field if non-nil, zero value otherwise.

### GetCardInformationOk

`func (o *Card) GetCardInformationOk() (*TokenizedCardInformation, bool)`

GetCardInformationOk returns a tuple with the CardInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardInformation

`func (o *Card) SetCardInformation(v TokenizedCardInformation)`

SetCardInformation sets CardInformation field to given value.

### HasCardInformation

`func (o *Card) HasCardInformation() bool`

HasCardInformation returns a boolean if a field has been set.

### GetCardVerificationResults

`func (o *Card) GetCardVerificationResults() CardVerificationResults`

GetCardVerificationResults returns the CardVerificationResults field if non-nil, zero value otherwise.

### GetCardVerificationResultsOk

`func (o *Card) GetCardVerificationResultsOk() (*CardVerificationResults, bool)`

GetCardVerificationResultsOk returns a tuple with the CardVerificationResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardVerificationResults

`func (o *Card) SetCardVerificationResults(v CardVerificationResults)`

SetCardVerificationResults sets CardVerificationResults field to given value.

### HasCardVerificationResults

`func (o *Card) HasCardVerificationResults() bool`

HasCardVerificationResults returns a boolean if a field has been set.

### SetCardVerificationResultsNil

`func (o *Card) SetCardVerificationResultsNil(b bool)`

 SetCardVerificationResultsNil sets the value for CardVerificationResults to be an explicit nil

### UnsetCardVerificationResults
`func (o *Card) UnsetCardVerificationResults()`

UnsetCardVerificationResults ensures that no value is present for CardVerificationResults, not even an explicit nil

[[Back to README]](../../README.md)


