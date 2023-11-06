# QRCodeParameters


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Amount** | Pointer to **NullableFloat64** |  |  |  |
| **Currency** | Pointer to **string** |  |  |  |
| **ChannelCode** | Pointer to [**NullableQRCodeChannelCode**](QRCodeChannelCode.md) |  |  |  |
| **ChannelProperties** | Pointer to [**NullableQRCodeChannelProperties**](QRCodeChannelProperties.md) |  |  |  |

## Methods

### NewQRCodeParameters

`func NewQRCodeParameters() *QRCodeParameters`

NewQRCodeParameters instantiates a new QRCodeParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQRCodeParametersWithDefaults

`func NewQRCodeParametersWithDefaults() *QRCodeParameters`

NewQRCodeParametersWithDefaults instantiates a new QRCodeParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *QRCodeParameters) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QRCodeParameters) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QRCodeParameters) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *QRCodeParameters) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *QRCodeParameters) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *QRCodeParameters) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetCurrency

`func (o *QRCodeParameters) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *QRCodeParameters) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *QRCodeParameters) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *QRCodeParameters) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetChannelCode

`func (o *QRCodeParameters) GetChannelCode() QRCodeChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *QRCodeParameters) GetChannelCodeOk() (*QRCodeChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *QRCodeParameters) SetChannelCode(v QRCodeChannelCode)`

SetChannelCode sets ChannelCode field to given value.

### HasChannelCode

`func (o *QRCodeParameters) HasChannelCode() bool`

HasChannelCode returns a boolean if a field has been set.

### SetChannelCodeNil

`func (o *QRCodeParameters) SetChannelCodeNil(b bool)`

 SetChannelCodeNil sets the value for ChannelCode to be an explicit nil

### UnsetChannelCode
`func (o *QRCodeParameters) UnsetChannelCode()`

UnsetChannelCode ensures that no value is present for ChannelCode, not even an explicit nil
### GetChannelProperties

`func (o *QRCodeParameters) GetChannelProperties() QRCodeChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *QRCodeParameters) GetChannelPropertiesOk() (*QRCodeChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *QRCodeParameters) SetChannelProperties(v QRCodeChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *QRCodeParameters) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.

### SetChannelPropertiesNil

`func (o *QRCodeParameters) SetChannelPropertiesNil(b bool)`

 SetChannelPropertiesNil sets the value for ChannelProperties to be an explicit nil

### UnsetChannelProperties
`func (o *QRCodeParameters) UnsetChannelProperties()`

UnsetChannelProperties ensures that no value is present for ChannelProperties, not even an explicit nil

[[Back to README]](../../README.md)


