# AccountPayLater


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **AccountId** | Pointer to **string** |  | Alphanumeric string identifying this account. Usually an email address or phone number. |  |
| **AccountHolderName** | Pointer to **NullableString** |  | Name of account holder as per the cardless credit account. |  |
| **Currency** | Pointer to **string** |  |  |  |

## Methods

### NewAccountPayLater

`func NewAccountPayLater() *AccountPayLater`

NewAccountPayLater instantiates a new AccountPayLater object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountPayLaterWithDefaults

`func NewAccountPayLaterWithDefaults() *AccountPayLater`

NewAccountPayLaterWithDefaults instantiates a new AccountPayLater object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountPayLater) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountPayLater) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountPayLater) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AccountPayLater) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountHolderName

`func (o *AccountPayLater) GetAccountHolderName() string`

GetAccountHolderName returns the AccountHolderName field if non-nil, zero value otherwise.

### GetAccountHolderNameOk

`func (o *AccountPayLater) GetAccountHolderNameOk() (*string, bool)`

GetAccountHolderNameOk returns a tuple with the AccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderName

`func (o *AccountPayLater) SetAccountHolderName(v string)`

SetAccountHolderName sets AccountHolderName field to given value.

### HasAccountHolderName

`func (o *AccountPayLater) HasAccountHolderName() bool`

HasAccountHolderName returns a boolean if a field has been set.

### SetAccountHolderNameNil

`func (o *AccountPayLater) SetAccountHolderNameNil(b bool)`

 SetAccountHolderNameNil sets the value for AccountHolderName to be an explicit nil

### UnsetAccountHolderName
`func (o *AccountPayLater) UnsetAccountHolderName()`

UnsetAccountHolderName ensures that no value is present for AccountHolderName, not even an explicit nil
### GetCurrency

`func (o *AccountPayLater) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AccountPayLater) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AccountPayLater) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AccountPayLater) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to README]](../../README.md)


