# EWalletParameters


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **ChannelCode** | [**EWalletChannelCode**](EWalletChannelCode.md) | ☑️ |  |  |
| **ChannelProperties** | Pointer to [**EWalletChannelProperties**](EWalletChannelProperties.md) |  |  |  |
| **Account** | Pointer to [**EWalletAccount**](EWalletAccount.md) |  |  |  |

## Methods

### NewEWalletParameters

`func NewEWalletParameters(channelCode EWalletChannelCode, ) *EWalletParameters`

NewEWalletParameters instantiates a new EWalletParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEWalletParametersWithDefaults

`func NewEWalletParametersWithDefaults() *EWalletParameters`

NewEWalletParametersWithDefaults instantiates a new EWalletParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelCode

`func (o *EWalletParameters) GetChannelCode() EWalletChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *EWalletParameters) GetChannelCodeOk() (*EWalletChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *EWalletParameters) SetChannelCode(v EWalletChannelCode)`

SetChannelCode sets ChannelCode field to given value.


### GetChannelProperties

`func (o *EWalletParameters) GetChannelProperties() EWalletChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *EWalletParameters) GetChannelPropertiesOk() (*EWalletChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *EWalletParameters) SetChannelProperties(v EWalletChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *EWalletParameters) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.

### GetAccount

`func (o *EWalletParameters) GetAccount() EWalletAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *EWalletParameters) GetAccountOk() (*EWalletAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *EWalletParameters) SetAccount(v EWalletAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *EWalletParameters) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to README]](../../README.md)


