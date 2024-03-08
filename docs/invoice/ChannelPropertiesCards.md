# ChannelPropertiesCards
An object representing properties specific for credit card payment method.

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **AllowedBins** | Pointer to **string[]** |  | An array of allowed BINs (6 or 8 digits) for credit card payments. |  |

## Methods

### NewChannelPropertiesCards

`func NewChannelPropertiesCards() *ChannelPropertiesCards`

NewChannelPropertiesCards instantiates a new ChannelPropertiesCards object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelPropertiesCardsWithDefaults

`func NewChannelPropertiesCardsWithDefaults() *ChannelPropertiesCards`

NewChannelPropertiesCardsWithDefaults instantiates a new ChannelPropertiesCards object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedBins

`func (o *ChannelPropertiesCards) GetAllowedBins() []string`

GetAllowedBins returns the AllowedBins field if non-nil, zero value otherwise.

### GetAllowedBinsOk

`func (o *ChannelPropertiesCards) GetAllowedBinsOk() (*[]string, bool)`

GetAllowedBinsOk returns a tuple with the AllowedBins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedBins

`func (o *ChannelPropertiesCards) SetAllowedBins(v []string)`

SetAllowedBins sets AllowedBins field to given value.

### HasAllowedBins

`func (o *ChannelPropertiesCards) HasAllowedBins() bool`

HasAllowedBins returns a boolean if a field has been set.


[[Back to README]](../../README.md)


