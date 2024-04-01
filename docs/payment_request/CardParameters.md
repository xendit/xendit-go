# CardParameters


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **ChannelProperties** | [**CardChannelProperties**](CardChannelProperties.md) | ☑️ |  |  |
| **CardInformation** | Pointer to [**CardInformation**](CardInformation.md) |  |  |  |

## Methods

### NewCardParameters

`func NewCardParameters(channelProperties CardChannelProperties, ) *CardParameters`

NewCardParameters instantiates a new CardParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardParametersWithDefaults

`func NewCardParametersWithDefaults() *CardParameters`

NewCardParametersWithDefaults instantiates a new CardParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetCardInformation

`func (o *CardParameters) GetCardInformation() CardInformation`

GetCardInformation returns the CardInformation field if non-nil, zero value otherwise.

### GetCardInformationOk

`func (o *CardParameters) GetCardInformationOk() (*CardInformation, bool)`

GetCardInformationOk returns a tuple with the CardInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardInformation

`func (o *CardParameters) SetCardInformation(v CardInformation)`

SetCardInformation sets CardInformation field to given value.

### HasCardInformation

`func (o *CardParameters) HasCardInformation() bool`

HasCardInformation returns a boolean if a field has been set.


[[Back to README]](../../README.md)


