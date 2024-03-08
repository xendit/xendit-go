# ChannelProperties
An object representing channel-specific properties.

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Cards** | Pointer to [**ChannelPropertiesCards**](ChannelPropertiesCards.md) |  |  |  |

## Methods

### NewChannelProperties

`func NewChannelProperties() *ChannelProperties`

NewChannelProperties instantiates a new ChannelProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelPropertiesWithDefaults

`func NewChannelPropertiesWithDefaults() *ChannelProperties`

NewChannelPropertiesWithDefaults instantiates a new ChannelProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCards

`func (o *ChannelProperties) GetCards() ChannelPropertiesCards`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *ChannelProperties) GetCardsOk() (*ChannelPropertiesCards, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *ChannelProperties) SetCards(v ChannelPropertiesCards)`

SetCards sets Cards field to given value.

### HasCards

`func (o *ChannelProperties) HasCards() bool`

HasCards returns a boolean if a field has been set.


[[Back to README]](../../README.md)


