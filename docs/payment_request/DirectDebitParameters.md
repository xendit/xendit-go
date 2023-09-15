# DirectDebitParameters

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **ChannelCode** | [**DirectDebitChannelCode**](DirectDebitChannelCode.md) |  |  |
| **ChannelProperties** | [**NullableDirectDebitChannelProperties**](DirectDebitChannelProperties.md) |  |  |
| **Type** | Pointer to [**DirectDebitType**](DirectDebitType.md) |  | [optional]  |

## Methods

### NewDirectDebitParameters

`func NewDirectDebitParameters(channelCode DirectDebitChannelCode, channelProperties NullableDirectDebitChannelProperties, ) *DirectDebitParameters`

NewDirectDebitParameters instantiates a new DirectDebitParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectDebitParametersWithDefaults

`func NewDirectDebitParametersWithDefaults() *DirectDebitParameters`

NewDirectDebitParametersWithDefaults instantiates a new DirectDebitParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelCode

`func (o *DirectDebitParameters) GetChannelCode() DirectDebitChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *DirectDebitParameters) GetChannelCodeOk() (*DirectDebitChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *DirectDebitParameters) SetChannelCode(v DirectDebitChannelCode)`

SetChannelCode sets ChannelCode field to given value.


### GetChannelProperties

`func (o *DirectDebitParameters) GetChannelProperties() DirectDebitChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *DirectDebitParameters) GetChannelPropertiesOk() (*DirectDebitChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *DirectDebitParameters) SetChannelProperties(v DirectDebitChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.


### SetChannelPropertiesNil

`func (o *DirectDebitParameters) SetChannelPropertiesNil(b bool)`

 SetChannelPropertiesNil sets the value for ChannelProperties to be an explicit nil

### UnsetChannelProperties
`func (o *DirectDebitParameters) UnsetChannelProperties()`

UnsetChannelProperties ensures that no value is present for ChannelProperties, not even an explicit nil
### GetType

`func (o *DirectDebitParameters) GetType() DirectDebitType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DirectDebitParameters) GetTypeOk() (*DirectDebitType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DirectDebitParameters) SetType(v DirectDebitType)`

SetType sets Type field to given value.

### HasType

`func (o *DirectDebitParameters) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to README]](../../README.md)


