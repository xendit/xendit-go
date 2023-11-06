# VirtualAccount
Virtual Account Payment Method Details

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Amount** | Pointer to **NullableFloat64** |  |  |  |
| **MinAmount** | Pointer to **NullableFloat64** |  |  |  |
| **MaxAmount** | Pointer to **NullableFloat64** |  |  |  |
| **Currency** | Pointer to **string** |  |  |  |
| **ChannelCode** | [**VirtualAccountChannelCode**](VirtualAccountChannelCode.md) | ☑️ |  |  |
| **ChannelProperties** | [**VirtualAccountChannelProperties**](VirtualAccountChannelProperties.md) | ☑️ |  |  |
| **AlternativeDisplayTypes** | Pointer to **string[]** |  | For payments in Vietnam only, alternative display requested for the virtual account |  |
| **AlternativeDisplays** | Pointer to [**VirtualAccountAlternativeDisplay[]**](VirtualAccountAlternativeDisplay.md) |  |  |  |

## Methods

### NewVirtualAccount

`func NewVirtualAccount(channelCode VirtualAccountChannelCode, channelProperties VirtualAccountChannelProperties, ) *VirtualAccount`

NewVirtualAccount instantiates a new VirtualAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualAccountWithDefaults

`func NewVirtualAccountWithDefaults() *VirtualAccount`

NewVirtualAccountWithDefaults instantiates a new VirtualAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *VirtualAccount) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *VirtualAccount) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *VirtualAccount) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *VirtualAccount) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *VirtualAccount) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *VirtualAccount) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetMinAmount

`func (o *VirtualAccount) GetMinAmount() float64`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *VirtualAccount) GetMinAmountOk() (*float64, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *VirtualAccount) SetMinAmount(v float64)`

SetMinAmount sets MinAmount field to given value.

### HasMinAmount

`func (o *VirtualAccount) HasMinAmount() bool`

HasMinAmount returns a boolean if a field has been set.

### SetMinAmountNil

`func (o *VirtualAccount) SetMinAmountNil(b bool)`

 SetMinAmountNil sets the value for MinAmount to be an explicit nil

### UnsetMinAmount
`func (o *VirtualAccount) UnsetMinAmount()`

UnsetMinAmount ensures that no value is present for MinAmount, not even an explicit nil
### GetMaxAmount

`func (o *VirtualAccount) GetMaxAmount() float64`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *VirtualAccount) GetMaxAmountOk() (*float64, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *VirtualAccount) SetMaxAmount(v float64)`

SetMaxAmount sets MaxAmount field to given value.

### HasMaxAmount

`func (o *VirtualAccount) HasMaxAmount() bool`

HasMaxAmount returns a boolean if a field has been set.

### SetMaxAmountNil

`func (o *VirtualAccount) SetMaxAmountNil(b bool)`

 SetMaxAmountNil sets the value for MaxAmount to be an explicit nil

### UnsetMaxAmount
`func (o *VirtualAccount) UnsetMaxAmount()`

UnsetMaxAmount ensures that no value is present for MaxAmount, not even an explicit nil
### GetCurrency

`func (o *VirtualAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *VirtualAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *VirtualAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *VirtualAccount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetChannelCode

`func (o *VirtualAccount) GetChannelCode() VirtualAccountChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *VirtualAccount) GetChannelCodeOk() (*VirtualAccountChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *VirtualAccount) SetChannelCode(v VirtualAccountChannelCode)`

SetChannelCode sets ChannelCode field to given value.


### GetChannelProperties

`func (o *VirtualAccount) GetChannelProperties() VirtualAccountChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *VirtualAccount) GetChannelPropertiesOk() (*VirtualAccountChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *VirtualAccount) SetChannelProperties(v VirtualAccountChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.


### GetAlternativeDisplayTypes

`func (o *VirtualAccount) GetAlternativeDisplayTypes() []string`

GetAlternativeDisplayTypes returns the AlternativeDisplayTypes field if non-nil, zero value otherwise.

### GetAlternativeDisplayTypesOk

`func (o *VirtualAccount) GetAlternativeDisplayTypesOk() (*[]string, bool)`

GetAlternativeDisplayTypesOk returns a tuple with the AlternativeDisplayTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeDisplayTypes

`func (o *VirtualAccount) SetAlternativeDisplayTypes(v []string)`

SetAlternativeDisplayTypes sets AlternativeDisplayTypes field to given value.

### HasAlternativeDisplayTypes

`func (o *VirtualAccount) HasAlternativeDisplayTypes() bool`

HasAlternativeDisplayTypes returns a boolean if a field has been set.

### GetAlternativeDisplays

`func (o *VirtualAccount) GetAlternativeDisplays() []VirtualAccountAlternativeDisplay`

GetAlternativeDisplays returns the AlternativeDisplays field if non-nil, zero value otherwise.

### GetAlternativeDisplaysOk

`func (o *VirtualAccount) GetAlternativeDisplaysOk() (*[]VirtualAccountAlternativeDisplay, bool)`

GetAlternativeDisplaysOk returns a tuple with the AlternativeDisplays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeDisplays

`func (o *VirtualAccount) SetAlternativeDisplays(v []VirtualAccountAlternativeDisplay)`

SetAlternativeDisplays sets AlternativeDisplays field to given value.

### HasAlternativeDisplays

`func (o *VirtualAccount) HasAlternativeDisplays() bool`

HasAlternativeDisplays returns a boolean if a field has been set.


[[Back to README]](../../README.md)


