# EWallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelCode** | Pointer to [**EWalletChannelCode**](EWalletChannelCode.md) |  | [optional] 
**ChannelProperties** | Pointer to [**EWalletChannelProperties**](EWalletChannelProperties.md) |  | [optional] 
**Account** | Pointer to [**EWalletAccount**](EWalletAccount.md) |  | [optional] 

## Methods

### NewEWallet

`func NewEWallet() *EWallet`

NewEWallet instantiates a new EWallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEWalletWithDefaults

`func NewEWalletWithDefaults() *EWallet`

NewEWalletWithDefaults instantiates a new EWallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelCode

`func (o *EWallet) GetChannelCode() EWalletChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *EWallet) GetChannelCodeOk() (*EWalletChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *EWallet) SetChannelCode(v EWalletChannelCode)`

SetChannelCode sets ChannelCode field to given value.

### HasChannelCode

`func (o *EWallet) HasChannelCode() bool`

HasChannelCode returns a boolean if a field has been set.

### GetChannelProperties

`func (o *EWallet) GetChannelProperties() EWalletChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *EWallet) GetChannelPropertiesOk() (*EWalletChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *EWallet) SetChannelProperties(v EWalletChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *EWallet) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.

### GetAccount

`func (o *EWallet) GetAccount() EWalletAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *EWallet) GetAccountOk() (*EWalletAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *EWallet) SetAccount(v EWalletAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *EWallet) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to README]](../../README.md)


