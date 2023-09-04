# OverTheCounterUpdateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **NullableFloat64** |  | [optional] 
**ChannelProperties** | Pointer to [**OverTheCounterChannelPropertiesUpdate**](OverTheCounterChannelPropertiesUpdate.md) |  | [optional] 

## Methods

### NewOverTheCounterUpdateParameters

`func NewOverTheCounterUpdateParameters() *OverTheCounterUpdateParameters`

NewOverTheCounterUpdateParameters instantiates a new OverTheCounterUpdateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverTheCounterUpdateParametersWithDefaults

`func NewOverTheCounterUpdateParametersWithDefaults() *OverTheCounterUpdateParameters`

NewOverTheCounterUpdateParametersWithDefaults instantiates a new OverTheCounterUpdateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *OverTheCounterUpdateParameters) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OverTheCounterUpdateParameters) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OverTheCounterUpdateParameters) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *OverTheCounterUpdateParameters) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *OverTheCounterUpdateParameters) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *OverTheCounterUpdateParameters) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetChannelProperties

`func (o *OverTheCounterUpdateParameters) GetChannelProperties() OverTheCounterChannelPropertiesUpdate`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *OverTheCounterUpdateParameters) GetChannelPropertiesOk() (*OverTheCounterChannelPropertiesUpdate, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *OverTheCounterUpdateParameters) SetChannelProperties(v OverTheCounterChannelPropertiesUpdate)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *OverTheCounterUpdateParameters) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.


[[Back to README]](../../README.md)


