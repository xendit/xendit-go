# Card


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **ChannelCode** | Pointer to [**CardChannelCode**](CardChannelCode.md) |  |  |  |
| **Currency** | Pointer to [**PaymentRequestCurrency**](PaymentRequestCurrency.md) |  |  |  |
| **ChannelProperties** | [**CardChannelProperties**](CardChannelProperties.md) | ☑️ |  |  |
| **CardInformation** | Pointer to [**CardInformation**](CardInformation.md) |  |  |  |
| **CardVerificationResults** | Pointer to [**NullableCardVerificationResults**](CardVerificationResults.md) |  |  |  |

## Methods

### NewCard

`func NewCard(channelProperties CardChannelProperties, ) *Card`

NewCard instantiates a new Card object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardWithDefaults

`func NewCardWithDefaults() *Card`

NewCardWithDefaults instantiates a new Card object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelCode

`func (o *Card) GetChannelCode() CardChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *Card) GetChannelCodeOk() (*CardChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *Card) SetChannelCode(v CardChannelCode)`

SetChannelCode sets ChannelCode field to given value.

### HasChannelCode

`func (o *Card) HasChannelCode() bool`

HasChannelCode returns a boolean if a field has been set.

### GetCurrency

`func (o *Card) GetCurrency() PaymentRequestCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Card) GetCurrencyOk() (*PaymentRequestCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Card) SetCurrency(v PaymentRequestCurrency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Card) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

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


### GetCardInformation

`func (o *Card) GetCardInformation() CardInformation`

GetCardInformation returns the CardInformation field if non-nil, zero value otherwise.

### GetCardInformationOk

`func (o *Card) GetCardInformationOk() (*CardInformation, bool)`

GetCardInformationOk returns a tuple with the CardInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardInformation

`func (o *Card) SetCardInformation(v CardInformation)`

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


