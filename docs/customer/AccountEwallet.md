# AccountEwallet

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **AccountNumber** | Pointer to **string** | Unique account identifier as per the bank records. | [optional]  |
| **AccountHolderName** | Pointer to **NullableString** | Name of account holder as per the bank records. Needs to match the registered account name exactly. | [optional]  |
| **Currency** | Pointer to **string** |  | [optional]  |

## Methods

### NewAccountEwallet

`func NewAccountEwallet() *AccountEwallet`

NewAccountEwallet instantiates a new AccountEwallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountEwalletWithDefaults

`func NewAccountEwalletWithDefaults() *AccountEwallet`

NewAccountEwalletWithDefaults instantiates a new AccountEwallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *AccountEwallet) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AccountEwallet) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AccountEwallet) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *AccountEwallet) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountHolderName

`func (o *AccountEwallet) GetAccountHolderName() string`

GetAccountHolderName returns the AccountHolderName field if non-nil, zero value otherwise.

### GetAccountHolderNameOk

`func (o *AccountEwallet) GetAccountHolderNameOk() (*string, bool)`

GetAccountHolderNameOk returns a tuple with the AccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderName

`func (o *AccountEwallet) SetAccountHolderName(v string)`

SetAccountHolderName sets AccountHolderName field to given value.

### HasAccountHolderName

`func (o *AccountEwallet) HasAccountHolderName() bool`

HasAccountHolderName returns a boolean if a field has been set.

### SetAccountHolderNameNil

`func (o *AccountEwallet) SetAccountHolderNameNil(b bool)`

 SetAccountHolderNameNil sets the value for AccountHolderName to be an explicit nil

### UnsetAccountHolderName
`func (o *AccountEwallet) UnsetAccountHolderName()`

UnsetAccountHolderName ensures that no value is present for AccountHolderName, not even an explicit nil
### GetCurrency

`func (o *AccountEwallet) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AccountEwallet) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AccountEwallet) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AccountEwallet) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to README]](../../README.md)


