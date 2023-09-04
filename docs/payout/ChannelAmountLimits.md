# ChannelAmountLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Minimum** | **float32** | Lowest amount supported for a payout to this channel | 
**Maximum** | **float32** | Highest amount supported for a payout to this channel | 
**MinimumIncrement** | **float32** | Supported increments | 

## Methods

### NewChannelAmountLimits

`func NewChannelAmountLimits(minimum float32, maximum float32, minimumIncrement float32, ) *ChannelAmountLimits`

NewChannelAmountLimits instantiates a new ChannelAmountLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelAmountLimitsWithDefaults

`func NewChannelAmountLimitsWithDefaults() *ChannelAmountLimits`

NewChannelAmountLimitsWithDefaults instantiates a new ChannelAmountLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinimum

`func (o *ChannelAmountLimits) GetMinimum() float32`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *ChannelAmountLimits) GetMinimumOk() (*float32, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *ChannelAmountLimits) SetMinimum(v float32)`

SetMinimum sets Minimum field to given value.


### GetMaximum

`func (o *ChannelAmountLimits) GetMaximum() float32`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *ChannelAmountLimits) GetMaximumOk() (*float32, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *ChannelAmountLimits) SetMaximum(v float32)`

SetMaximum sets Maximum field to given value.


### GetMinimumIncrement

`func (o *ChannelAmountLimits) GetMinimumIncrement() float32`

GetMinimumIncrement returns the MinimumIncrement field if non-nil, zero value otherwise.

### GetMinimumIncrementOk

`func (o *ChannelAmountLimits) GetMinimumIncrementOk() (*float32, bool)`

GetMinimumIncrementOk returns a tuple with the MinimumIncrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumIncrement

`func (o *ChannelAmountLimits) SetMinimumIncrement(v float32)`

SetMinimumIncrement sets MinimumIncrement field to given value.



[[Back to README]](../../README.md)


