# VirtualAccountUpdateParameters

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **Amount** | Pointer to **NullableFloat64** |  | [optional]  |
| **MinAmount** | Pointer to **NullableFloat64** |  | [optional]  |
| **MaxAmount** | Pointer to **NullableFloat64** |  | [optional]  |
| **ChannelProperties** | Pointer to [**VirtualAccountChannelPropertiesPatch**](VirtualAccountChannelPropertiesPatch.md) |  | [optional]  |
| **AlternativeDisplayTypes** | Pointer to **string[]** | For payments in Vietnam only, alternative display requested for the virtual account | [optional]  |

## Methods

### NewVirtualAccountUpdateParameters

`func NewVirtualAccountUpdateParameters() *VirtualAccountUpdateParameters`

NewVirtualAccountUpdateParameters instantiates a new VirtualAccountUpdateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualAccountUpdateParametersWithDefaults

`func NewVirtualAccountUpdateParametersWithDefaults() *VirtualAccountUpdateParameters`

NewVirtualAccountUpdateParametersWithDefaults instantiates a new VirtualAccountUpdateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *VirtualAccountUpdateParameters) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *VirtualAccountUpdateParameters) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *VirtualAccountUpdateParameters) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *VirtualAccountUpdateParameters) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *VirtualAccountUpdateParameters) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *VirtualAccountUpdateParameters) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetMinAmount

`func (o *VirtualAccountUpdateParameters) GetMinAmount() float64`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *VirtualAccountUpdateParameters) GetMinAmountOk() (*float64, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *VirtualAccountUpdateParameters) SetMinAmount(v float64)`

SetMinAmount sets MinAmount field to given value.

### HasMinAmount

`func (o *VirtualAccountUpdateParameters) HasMinAmount() bool`

HasMinAmount returns a boolean if a field has been set.

### SetMinAmountNil

`func (o *VirtualAccountUpdateParameters) SetMinAmountNil(b bool)`

 SetMinAmountNil sets the value for MinAmount to be an explicit nil

### UnsetMinAmount
`func (o *VirtualAccountUpdateParameters) UnsetMinAmount()`

UnsetMinAmount ensures that no value is present for MinAmount, not even an explicit nil
### GetMaxAmount

`func (o *VirtualAccountUpdateParameters) GetMaxAmount() float64`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *VirtualAccountUpdateParameters) GetMaxAmountOk() (*float64, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *VirtualAccountUpdateParameters) SetMaxAmount(v float64)`

SetMaxAmount sets MaxAmount field to given value.

### HasMaxAmount

`func (o *VirtualAccountUpdateParameters) HasMaxAmount() bool`

HasMaxAmount returns a boolean if a field has been set.

### SetMaxAmountNil

`func (o *VirtualAccountUpdateParameters) SetMaxAmountNil(b bool)`

 SetMaxAmountNil sets the value for MaxAmount to be an explicit nil

### UnsetMaxAmount
`func (o *VirtualAccountUpdateParameters) UnsetMaxAmount()`

UnsetMaxAmount ensures that no value is present for MaxAmount, not even an explicit nil
### GetChannelProperties

`func (o *VirtualAccountUpdateParameters) GetChannelProperties() VirtualAccountChannelPropertiesPatch`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *VirtualAccountUpdateParameters) GetChannelPropertiesOk() (*VirtualAccountChannelPropertiesPatch, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *VirtualAccountUpdateParameters) SetChannelProperties(v VirtualAccountChannelPropertiesPatch)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *VirtualAccountUpdateParameters) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.

### GetAlternativeDisplayTypes

`func (o *VirtualAccountUpdateParameters) GetAlternativeDisplayTypes() []string`

GetAlternativeDisplayTypes returns the AlternativeDisplayTypes field if non-nil, zero value otherwise.

### GetAlternativeDisplayTypesOk

`func (o *VirtualAccountUpdateParameters) GetAlternativeDisplayTypesOk() (*[]string, bool)`

GetAlternativeDisplayTypesOk returns a tuple with the AlternativeDisplayTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeDisplayTypes

`func (o *VirtualAccountUpdateParameters) SetAlternativeDisplayTypes(v []string)`

SetAlternativeDisplayTypes sets AlternativeDisplayTypes field to given value.

### HasAlternativeDisplayTypes

`func (o *VirtualAccountUpdateParameters) HasAlternativeDisplayTypes() bool`

HasAlternativeDisplayTypes returns a boolean if a field has been set.


[[Back to README]](../../README.md)


