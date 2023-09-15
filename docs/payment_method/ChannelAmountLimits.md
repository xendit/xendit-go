# ChannelAmountLimits

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **Currency** | Pointer to **string** | Currency supported by the payment channel | [optional]  |
| **MinLimit** | Pointer to **float32** | The minimum allowed transaction amount for the payment channel | [optional]  |
| **MaxLimit** | Pointer to **float32** | The minimum allowed transaction amount for the payment channel | [optional]  |

## Methods

### NewChannelAmountLimits

`func NewChannelAmountLimits() *ChannelAmountLimits`

NewChannelAmountLimits instantiates a new ChannelAmountLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelAmountLimitsWithDefaults

`func NewChannelAmountLimitsWithDefaults() *ChannelAmountLimits`

NewChannelAmountLimitsWithDefaults instantiates a new ChannelAmountLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *ChannelAmountLimits) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ChannelAmountLimits) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ChannelAmountLimits) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ChannelAmountLimits) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetMinLimit

`func (o *ChannelAmountLimits) GetMinLimit() float32`

GetMinLimit returns the MinLimit field if non-nil, zero value otherwise.

### GetMinLimitOk

`func (o *ChannelAmountLimits) GetMinLimitOk() (*float32, bool)`

GetMinLimitOk returns a tuple with the MinLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLimit

`func (o *ChannelAmountLimits) SetMinLimit(v float32)`

SetMinLimit sets MinLimit field to given value.

### HasMinLimit

`func (o *ChannelAmountLimits) HasMinLimit() bool`

HasMinLimit returns a boolean if a field has been set.

### GetMaxLimit

`func (o *ChannelAmountLimits) GetMaxLimit() float32`

GetMaxLimit returns the MaxLimit field if non-nil, zero value otherwise.

### GetMaxLimitOk

`func (o *ChannelAmountLimits) GetMaxLimitOk() (*float32, bool)`

GetMaxLimitOk returns a tuple with the MaxLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLimit

`func (o *ChannelAmountLimits) SetMaxLimit(v float32)`

SetMaxLimit sets MaxLimit field to given value.

### HasMaxLimit

`func (o *ChannelAmountLimits) HasMaxLimit() bool`

HasMaxLimit returns a boolean if a field has been set.


[[Back to README]](../../README.md)


