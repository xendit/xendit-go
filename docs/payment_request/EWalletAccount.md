# EWalletAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name of the eWallet account holder. The value is null if unavailableName of the eWallet account holder. The value is null if unavailable | [optional] 
**AccountDetails** | Pointer to **NullableString** | Identifier from eWallet provider e.g. phone number. The value is null if unavailable | [optional] 
**Balance** | Pointer to **NullableFloat64** | The main balance amount on eWallet account provided from eWallet provider. The value is null if unavailable | [optional] 
**PointBalance** | Pointer to **NullableFloat64** | The point balance amount on eWallet account. Applicable only on some eWallet provider that has point system. The value is null if unavailabl | [optional] 

## Methods

### NewEWalletAccount

`func NewEWalletAccount() *EWalletAccount`

NewEWalletAccount instantiates a new EWalletAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEWalletAccountWithDefaults

`func NewEWalletAccountWithDefaults() *EWalletAccount`

NewEWalletAccountWithDefaults instantiates a new EWalletAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EWalletAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EWalletAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EWalletAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EWalletAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EWalletAccount) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EWalletAccount) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAccountDetails

`func (o *EWalletAccount) GetAccountDetails() string`

GetAccountDetails returns the AccountDetails field if non-nil, zero value otherwise.

### GetAccountDetailsOk

`func (o *EWalletAccount) GetAccountDetailsOk() (*string, bool)`

GetAccountDetailsOk returns a tuple with the AccountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDetails

`func (o *EWalletAccount) SetAccountDetails(v string)`

SetAccountDetails sets AccountDetails field to given value.

### HasAccountDetails

`func (o *EWalletAccount) HasAccountDetails() bool`

HasAccountDetails returns a boolean if a field has been set.

### SetAccountDetailsNil

`func (o *EWalletAccount) SetAccountDetailsNil(b bool)`

 SetAccountDetailsNil sets the value for AccountDetails to be an explicit nil

### UnsetAccountDetails
`func (o *EWalletAccount) UnsetAccountDetails()`

UnsetAccountDetails ensures that no value is present for AccountDetails, not even an explicit nil
### GetBalance

`func (o *EWalletAccount) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *EWalletAccount) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *EWalletAccount) SetBalance(v float64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *EWalletAccount) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### SetBalanceNil

`func (o *EWalletAccount) SetBalanceNil(b bool)`

 SetBalanceNil sets the value for Balance to be an explicit nil

### UnsetBalance
`func (o *EWalletAccount) UnsetBalance()`

UnsetBalance ensures that no value is present for Balance, not even an explicit nil
### GetPointBalance

`func (o *EWalletAccount) GetPointBalance() float64`

GetPointBalance returns the PointBalance field if non-nil, zero value otherwise.

### GetPointBalanceOk

`func (o *EWalletAccount) GetPointBalanceOk() (*float64, bool)`

GetPointBalanceOk returns a tuple with the PointBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointBalance

`func (o *EWalletAccount) SetPointBalance(v float64)`

SetPointBalance sets PointBalance field to given value.

### HasPointBalance

`func (o *EWalletAccount) HasPointBalance() bool`

HasPointBalance returns a boolean if a field has been set.

### SetPointBalanceNil

`func (o *EWalletAccount) SetPointBalanceNil(b bool)`

 SetPointBalanceNil sets the value for PointBalance to be an explicit nil

### UnsetPointBalance
`func (o *EWalletAccount) UnsetPointBalance()`

UnsetPointBalance ensures that no value is present for PointBalance, not even an explicit nil

[[Back to README]](../../README.md)


