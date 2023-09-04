# VirtualAccountParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinAmount** | Pointer to **NullableFloat64** |  | [optional] 
**MaxAmount** | Pointer to **NullableFloat64** |  | [optional] 
**Amount** | Pointer to **NullableFloat64** |  | [optional] 
**Currency** | Pointer to [**PaymentRequestCurrency**](PaymentRequestCurrency.md) |  | [optional] 
**ChannelCode** | [**VirtualAccountChannelCode**](VirtualAccountChannelCode.md) |  | 
**ChannelProperties** | [**VirtualAccountChannelProperties**](VirtualAccountChannelProperties.md) |  | 
**AlternativeDisplayTypes** | Pointer to **[]string** | Alternative display requested for the virtual account | [optional] 

## Methods

### NewVirtualAccountParameters

`func NewVirtualAccountParameters(channelCode VirtualAccountChannelCode, channelProperties VirtualAccountChannelProperties, ) *VirtualAccountParameters`

NewVirtualAccountParameters instantiates a new VirtualAccountParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualAccountParametersWithDefaults

`func NewVirtualAccountParametersWithDefaults() *VirtualAccountParameters`

NewVirtualAccountParametersWithDefaults instantiates a new VirtualAccountParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinAmount

`func (o *VirtualAccountParameters) GetMinAmount() float64`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *VirtualAccountParameters) GetMinAmountOk() (*float64, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *VirtualAccountParameters) SetMinAmount(v float64)`

SetMinAmount sets MinAmount field to given value.

### HasMinAmount

`func (o *VirtualAccountParameters) HasMinAmount() bool`

HasMinAmount returns a boolean if a field has been set.

### SetMinAmountNil

`func (o *VirtualAccountParameters) SetMinAmountNil(b bool)`

 SetMinAmountNil sets the value for MinAmount to be an explicit nil

### UnsetMinAmount
`func (o *VirtualAccountParameters) UnsetMinAmount()`

UnsetMinAmount ensures that no value is present for MinAmount, not even an explicit nil
### GetMaxAmount

`func (o *VirtualAccountParameters) GetMaxAmount() float64`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *VirtualAccountParameters) GetMaxAmountOk() (*float64, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *VirtualAccountParameters) SetMaxAmount(v float64)`

SetMaxAmount sets MaxAmount field to given value.

### HasMaxAmount

`func (o *VirtualAccountParameters) HasMaxAmount() bool`

HasMaxAmount returns a boolean if a field has been set.

### SetMaxAmountNil

`func (o *VirtualAccountParameters) SetMaxAmountNil(b bool)`

 SetMaxAmountNil sets the value for MaxAmount to be an explicit nil

### UnsetMaxAmount
`func (o *VirtualAccountParameters) UnsetMaxAmount()`

UnsetMaxAmount ensures that no value is present for MaxAmount, not even an explicit nil
### GetAmount

`func (o *VirtualAccountParameters) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *VirtualAccountParameters) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *VirtualAccountParameters) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *VirtualAccountParameters) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *VirtualAccountParameters) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *VirtualAccountParameters) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetCurrency

`func (o *VirtualAccountParameters) GetCurrency() PaymentRequestCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *VirtualAccountParameters) GetCurrencyOk() (*PaymentRequestCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *VirtualAccountParameters) SetCurrency(v PaymentRequestCurrency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *VirtualAccountParameters) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetChannelCode

`func (o *VirtualAccountParameters) GetChannelCode() VirtualAccountChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *VirtualAccountParameters) GetChannelCodeOk() (*VirtualAccountChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *VirtualAccountParameters) SetChannelCode(v VirtualAccountChannelCode)`

SetChannelCode sets ChannelCode field to given value.


### GetChannelProperties

`func (o *VirtualAccountParameters) GetChannelProperties() VirtualAccountChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *VirtualAccountParameters) GetChannelPropertiesOk() (*VirtualAccountChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *VirtualAccountParameters) SetChannelProperties(v VirtualAccountChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.


### GetAlternativeDisplayTypes

`func (o *VirtualAccountParameters) GetAlternativeDisplayTypes() []string`

GetAlternativeDisplayTypes returns the AlternativeDisplayTypes field if non-nil, zero value otherwise.

### GetAlternativeDisplayTypesOk

`func (o *VirtualAccountParameters) GetAlternativeDisplayTypesOk() (*[]string, bool)`

GetAlternativeDisplayTypesOk returns a tuple with the AlternativeDisplayTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeDisplayTypes

`func (o *VirtualAccountParameters) SetAlternativeDisplayTypes(v []string)`

SetAlternativeDisplayTypes sets AlternativeDisplayTypes field to given value.

### HasAlternativeDisplayTypes

`func (o *VirtualAccountParameters) HasAlternativeDisplayTypes() bool`

HasAlternativeDisplayTypes returns a boolean if a field has been set.


[[Back to README]](../../README.md)


