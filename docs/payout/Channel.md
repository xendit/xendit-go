# Channel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelCode** | **string** | Destination channel to send the money to, prefixed by ISO-3166 country code | 
**ChannelCategory** | [**ChannelCategory**](ChannelCategory.md) |  | 
**Currency** | **string** | Currency of the destination channel using ISO-4217 currency code | 
**ChannelName** | **string** | Name of the destination channel | 
**AmountLimits** | [**ChannelAmountLimits**](ChannelAmountLimits.md) |  | 

## Methods

### NewChannel

`func NewChannel(channelCode string, channelCategory ChannelCategory, currency string, channelName string, amountLimits ChannelAmountLimits, ) *Channel`

NewChannel instantiates a new Channel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelWithDefaults

`func NewChannelWithDefaults() *Channel`

NewChannelWithDefaults instantiates a new Channel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelCode

`func (o *Channel) GetChannelCode() string`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *Channel) GetChannelCodeOk() (*string, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *Channel) SetChannelCode(v string)`

SetChannelCode sets ChannelCode field to given value.


### GetChannelCategory

`func (o *Channel) GetChannelCategory() ChannelCategory`

GetChannelCategory returns the ChannelCategory field if non-nil, zero value otherwise.

### GetChannelCategoryOk

`func (o *Channel) GetChannelCategoryOk() (*ChannelCategory, bool)`

GetChannelCategoryOk returns a tuple with the ChannelCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCategory

`func (o *Channel) SetChannelCategory(v ChannelCategory)`

SetChannelCategory sets ChannelCategory field to given value.


### GetCurrency

`func (o *Channel) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Channel) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Channel) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetChannelName

`func (o *Channel) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *Channel) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *Channel) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.


### GetAmountLimits

`func (o *Channel) GetAmountLimits() ChannelAmountLimits`

GetAmountLimits returns the AmountLimits field if non-nil, zero value otherwise.

### GetAmountLimitsOk

`func (o *Channel) GetAmountLimitsOk() (*ChannelAmountLimits, bool)`

GetAmountLimitsOk returns a tuple with the AmountLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountLimits

`func (o *Channel) SetAmountLimits(v ChannelAmountLimits)`

SetAmountLimits sets AmountLimits field to given value.



[[Back to README]](../../README.md)


