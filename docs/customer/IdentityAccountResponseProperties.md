# IdentityAccountResponseProperties


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **AccountNumber** | Pointer to **string** |  | Unique account identifier as per the bank records. |  |
| **AccountHolderName** | Pointer to **NullableString** |  | Name of account holder as per the cardless credit account. |  |
| **SwiftCode** | Pointer to **NullableString** |  | The SWIFT code for international payments |  |
| **AccountType** | Pointer to **NullableString** |  | Free text account type, e.g., Savings, Transaction, Virtual Account. |  |
| **AccountDetails** | Pointer to **NullableString** |  | Potentially masked account detail, for display purposes only. |  |
| **Currency** | Pointer to **string** |  |  |  |
| **TokenId** | Pointer to **string** |  | The token id returned in tokenisation |  |
| **PaymentCode** | Pointer to **string** |  | Complete fixed payment code (including prefix) |  |
| **ExpiresAt** | Pointer to **NullableString** |  | YYYY-MM-DD string with expiry date for the payment code |  |
| **QrString** | Pointer to **string** |  | String representation of the QR Code image |  |
| **AccountId** | Pointer to **string** |  | Alphanumeric string identifying this account. Usually an email address or phone number. |  |

## Methods

### NewIdentityAccountResponseProperties

`func NewIdentityAccountResponseProperties() *IdentityAccountResponseProperties`

NewIdentityAccountResponseProperties instantiates a new IdentityAccountResponseProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityAccountResponsePropertiesWithDefaults

`func NewIdentityAccountResponsePropertiesWithDefaults() *IdentityAccountResponseProperties`

NewIdentityAccountResponsePropertiesWithDefaults instantiates a new IdentityAccountResponseProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *IdentityAccountResponseProperties) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *IdentityAccountResponseProperties) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *IdentityAccountResponseProperties) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *IdentityAccountResponseProperties) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountHolderName

`func (o *IdentityAccountResponseProperties) GetAccountHolderName() string`

GetAccountHolderName returns the AccountHolderName field if non-nil, zero value otherwise.

### GetAccountHolderNameOk

`func (o *IdentityAccountResponseProperties) GetAccountHolderNameOk() (*string, bool)`

GetAccountHolderNameOk returns a tuple with the AccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderName

`func (o *IdentityAccountResponseProperties) SetAccountHolderName(v string)`

SetAccountHolderName sets AccountHolderName field to given value.

### HasAccountHolderName

`func (o *IdentityAccountResponseProperties) HasAccountHolderName() bool`

HasAccountHolderName returns a boolean if a field has been set.

### SetAccountHolderNameNil

`func (o *IdentityAccountResponseProperties) SetAccountHolderNameNil(b bool)`

 SetAccountHolderNameNil sets the value for AccountHolderName to be an explicit nil

### UnsetAccountHolderName
`func (o *IdentityAccountResponseProperties) UnsetAccountHolderName()`

UnsetAccountHolderName ensures that no value is present for AccountHolderName, not even an explicit nil
### GetSwiftCode

`func (o *IdentityAccountResponseProperties) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *IdentityAccountResponseProperties) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *IdentityAccountResponseProperties) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.

### HasSwiftCode

`func (o *IdentityAccountResponseProperties) HasSwiftCode() bool`

HasSwiftCode returns a boolean if a field has been set.

### SetSwiftCodeNil

`func (o *IdentityAccountResponseProperties) SetSwiftCodeNil(b bool)`

 SetSwiftCodeNil sets the value for SwiftCode to be an explicit nil

### UnsetSwiftCode
`func (o *IdentityAccountResponseProperties) UnsetSwiftCode()`

UnsetSwiftCode ensures that no value is present for SwiftCode, not even an explicit nil
### GetAccountType

`func (o *IdentityAccountResponseProperties) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *IdentityAccountResponseProperties) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *IdentityAccountResponseProperties) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *IdentityAccountResponseProperties) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### SetAccountTypeNil

`func (o *IdentityAccountResponseProperties) SetAccountTypeNil(b bool)`

 SetAccountTypeNil sets the value for AccountType to be an explicit nil

### UnsetAccountType
`func (o *IdentityAccountResponseProperties) UnsetAccountType()`

UnsetAccountType ensures that no value is present for AccountType, not even an explicit nil
### GetAccountDetails

`func (o *IdentityAccountResponseProperties) GetAccountDetails() string`

GetAccountDetails returns the AccountDetails field if non-nil, zero value otherwise.

### GetAccountDetailsOk

`func (o *IdentityAccountResponseProperties) GetAccountDetailsOk() (*string, bool)`

GetAccountDetailsOk returns a tuple with the AccountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDetails

`func (o *IdentityAccountResponseProperties) SetAccountDetails(v string)`

SetAccountDetails sets AccountDetails field to given value.

### HasAccountDetails

`func (o *IdentityAccountResponseProperties) HasAccountDetails() bool`

HasAccountDetails returns a boolean if a field has been set.

### SetAccountDetailsNil

`func (o *IdentityAccountResponseProperties) SetAccountDetailsNil(b bool)`

 SetAccountDetailsNil sets the value for AccountDetails to be an explicit nil

### UnsetAccountDetails
`func (o *IdentityAccountResponseProperties) UnsetAccountDetails()`

UnsetAccountDetails ensures that no value is present for AccountDetails, not even an explicit nil
### GetCurrency

`func (o *IdentityAccountResponseProperties) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *IdentityAccountResponseProperties) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *IdentityAccountResponseProperties) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *IdentityAccountResponseProperties) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTokenId

`func (o *IdentityAccountResponseProperties) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *IdentityAccountResponseProperties) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *IdentityAccountResponseProperties) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *IdentityAccountResponseProperties) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetPaymentCode

`func (o *IdentityAccountResponseProperties) GetPaymentCode() string`

GetPaymentCode returns the PaymentCode field if non-nil, zero value otherwise.

### GetPaymentCodeOk

`func (o *IdentityAccountResponseProperties) GetPaymentCodeOk() (*string, bool)`

GetPaymentCodeOk returns a tuple with the PaymentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCode

`func (o *IdentityAccountResponseProperties) SetPaymentCode(v string)`

SetPaymentCode sets PaymentCode field to given value.

### HasPaymentCode

`func (o *IdentityAccountResponseProperties) HasPaymentCode() bool`

HasPaymentCode returns a boolean if a field has been set.

### GetExpiresAt

`func (o *IdentityAccountResponseProperties) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *IdentityAccountResponseProperties) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *IdentityAccountResponseProperties) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *IdentityAccountResponseProperties) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *IdentityAccountResponseProperties) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *IdentityAccountResponseProperties) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetQrString

`func (o *IdentityAccountResponseProperties) GetQrString() string`

GetQrString returns the QrString field if non-nil, zero value otherwise.

### GetQrStringOk

`func (o *IdentityAccountResponseProperties) GetQrStringOk() (*string, bool)`

GetQrStringOk returns a tuple with the QrString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrString

`func (o *IdentityAccountResponseProperties) SetQrString(v string)`

SetQrString sets QrString field to given value.

### HasQrString

`func (o *IdentityAccountResponseProperties) HasQrString() bool`

HasQrString returns a boolean if a field has been set.

### GetAccountId

`func (o *IdentityAccountResponseProperties) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *IdentityAccountResponseProperties) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *IdentityAccountResponseProperties) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *IdentityAccountResponseProperties) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.


[[Back to README]](../../README.md)


