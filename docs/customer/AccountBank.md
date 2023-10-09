# AccountBank

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **AccountNumber** | Pointer to **string** | Unique account identifier as per the bank records. | [optional]  |
| **AccountHolderName** | Pointer to **NullableString** | Name of account holder as per the bank records. Needs to match the registered account name exactly. . | [optional]  |
| **SwiftCode** | Pointer to **NullableString** | The SWIFT code for international payments | [optional]  |
| **AccountType** | Pointer to **NullableString** | Free text account type, e.g., Savings, Transaction, Virtual Account. | [optional]  |
| **AccountDetails** | Pointer to **NullableString** | Potentially masked account detail, for display purposes only. | [optional]  |
| **Currency** | Pointer to **string** |  | [optional]  |

## Methods

### NewAccountBank

`func NewAccountBank() *AccountBank`

NewAccountBank instantiates a new AccountBank object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountBankWithDefaults

`func NewAccountBankWithDefaults() *AccountBank`

NewAccountBankWithDefaults instantiates a new AccountBank object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *AccountBank) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AccountBank) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AccountBank) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *AccountBank) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountHolderName

`func (o *AccountBank) GetAccountHolderName() string`

GetAccountHolderName returns the AccountHolderName field if non-nil, zero value otherwise.

### GetAccountHolderNameOk

`func (o *AccountBank) GetAccountHolderNameOk() (*string, bool)`

GetAccountHolderNameOk returns a tuple with the AccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderName

`func (o *AccountBank) SetAccountHolderName(v string)`

SetAccountHolderName sets AccountHolderName field to given value.

### HasAccountHolderName

`func (o *AccountBank) HasAccountHolderName() bool`

HasAccountHolderName returns a boolean if a field has been set.

### SetAccountHolderNameNil

`func (o *AccountBank) SetAccountHolderNameNil(b bool)`

 SetAccountHolderNameNil sets the value for AccountHolderName to be an explicit nil

### UnsetAccountHolderName
`func (o *AccountBank) UnsetAccountHolderName()`

UnsetAccountHolderName ensures that no value is present for AccountHolderName, not even an explicit nil
### GetSwiftCode

`func (o *AccountBank) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *AccountBank) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *AccountBank) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.

### HasSwiftCode

`func (o *AccountBank) HasSwiftCode() bool`

HasSwiftCode returns a boolean if a field has been set.

### SetSwiftCodeNil

`func (o *AccountBank) SetSwiftCodeNil(b bool)`

 SetSwiftCodeNil sets the value for SwiftCode to be an explicit nil

### UnsetSwiftCode
`func (o *AccountBank) UnsetSwiftCode()`

UnsetSwiftCode ensures that no value is present for SwiftCode, not even an explicit nil
### GetAccountType

`func (o *AccountBank) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountBank) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountBank) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AccountBank) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### SetAccountTypeNil

`func (o *AccountBank) SetAccountTypeNil(b bool)`

 SetAccountTypeNil sets the value for AccountType to be an explicit nil

### UnsetAccountType
`func (o *AccountBank) UnsetAccountType()`

UnsetAccountType ensures that no value is present for AccountType, not even an explicit nil
### GetAccountDetails

`func (o *AccountBank) GetAccountDetails() string`

GetAccountDetails returns the AccountDetails field if non-nil, zero value otherwise.

### GetAccountDetailsOk

`func (o *AccountBank) GetAccountDetailsOk() (*string, bool)`

GetAccountDetailsOk returns a tuple with the AccountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDetails

`func (o *AccountBank) SetAccountDetails(v string)`

SetAccountDetails sets AccountDetails field to given value.

### HasAccountDetails

`func (o *AccountBank) HasAccountDetails() bool`

HasAccountDetails returns a boolean if a field has been set.

### SetAccountDetailsNil

`func (o *AccountBank) SetAccountDetailsNil(b bool)`

 SetAccountDetailsNil sets the value for AccountDetails to be an explicit nil

### UnsetAccountDetails
`func (o *AccountBank) UnsetAccountDetails()`

UnsetAccountDetails ensures that no value is present for AccountDetails, not even an explicit nil
### GetCurrency

`func (o *AccountBank) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AccountBank) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AccountBank) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AccountBank) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to README]](../../README.md)


