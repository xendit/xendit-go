# IdentityAccountRequestProperties

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **AccountNumber** | Pointer to **string** | Unique account identifier as per the bank records. | [optional]  |
| **AccountHolderName** | Pointer to **NullableString** | Name of account holder as per the cardless credit account. | [optional]  |
| **SwiftCode** | Pointer to **NullableString** | The SWIFT code for international payments | [optional]  |
| **AccountType** | Pointer to **NullableString** | Free text account type, e.g., Savings, Transaction, Virtual Account. | [optional]  |
| **AccountDetails** | Pointer to **NullableString** | Potentially masked account detail, for display purposes only. | [optional]  |
| **Currency** | Pointer to **string** |  | [optional]  |
| **TokenId** | Pointer to **string** | The token id returned in tokenisation | [optional]  |
| **AccountId** | Pointer to **string** | Alphanumeric string identifying this account. Usually an email address or phone number. | [optional]  |
| **PaymentCode** | Pointer to **string** | Complete fixed payment code (including prefix) | [optional]  |
| **ExpiresAt** | Pointer to **NullableString** | YYYY-MM-DD string with expiry date for the payment code | [optional]  |
| **QrString** | Pointer to **string** | String representation of the QR Code image | [optional]  |

## Methods

### NewIdentityAccountRequestProperties

`func NewIdentityAccountRequestProperties() *IdentityAccountRequestProperties`

NewIdentityAccountRequestProperties instantiates a new IdentityAccountRequestProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityAccountRequestPropertiesWithDefaults

`func NewIdentityAccountRequestPropertiesWithDefaults() *IdentityAccountRequestProperties`

NewIdentityAccountRequestPropertiesWithDefaults instantiates a new IdentityAccountRequestProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *IdentityAccountRequestProperties) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *IdentityAccountRequestProperties) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *IdentityAccountRequestProperties) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *IdentityAccountRequestProperties) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountHolderName

`func (o *IdentityAccountRequestProperties) GetAccountHolderName() string`

GetAccountHolderName returns the AccountHolderName field if non-nil, zero value otherwise.

### GetAccountHolderNameOk

`func (o *IdentityAccountRequestProperties) GetAccountHolderNameOk() (*string, bool)`

GetAccountHolderNameOk returns a tuple with the AccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderName

`func (o *IdentityAccountRequestProperties) SetAccountHolderName(v string)`

SetAccountHolderName sets AccountHolderName field to given value.

### HasAccountHolderName

`func (o *IdentityAccountRequestProperties) HasAccountHolderName() bool`

HasAccountHolderName returns a boolean if a field has been set.

### SetAccountHolderNameNil

`func (o *IdentityAccountRequestProperties) SetAccountHolderNameNil(b bool)`

 SetAccountHolderNameNil sets the value for AccountHolderName to be an explicit nil

### UnsetAccountHolderName
`func (o *IdentityAccountRequestProperties) UnsetAccountHolderName()`

UnsetAccountHolderName ensures that no value is present for AccountHolderName, not even an explicit nil
### GetSwiftCode

`func (o *IdentityAccountRequestProperties) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *IdentityAccountRequestProperties) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *IdentityAccountRequestProperties) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.

### HasSwiftCode

`func (o *IdentityAccountRequestProperties) HasSwiftCode() bool`

HasSwiftCode returns a boolean if a field has been set.

### SetSwiftCodeNil

`func (o *IdentityAccountRequestProperties) SetSwiftCodeNil(b bool)`

 SetSwiftCodeNil sets the value for SwiftCode to be an explicit nil

### UnsetSwiftCode
`func (o *IdentityAccountRequestProperties) UnsetSwiftCode()`

UnsetSwiftCode ensures that no value is present for SwiftCode, not even an explicit nil
### GetAccountType

`func (o *IdentityAccountRequestProperties) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *IdentityAccountRequestProperties) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *IdentityAccountRequestProperties) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *IdentityAccountRequestProperties) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### SetAccountTypeNil

`func (o *IdentityAccountRequestProperties) SetAccountTypeNil(b bool)`

 SetAccountTypeNil sets the value for AccountType to be an explicit nil

### UnsetAccountType
`func (o *IdentityAccountRequestProperties) UnsetAccountType()`

UnsetAccountType ensures that no value is present for AccountType, not even an explicit nil
### GetAccountDetails

`func (o *IdentityAccountRequestProperties) GetAccountDetails() string`

GetAccountDetails returns the AccountDetails field if non-nil, zero value otherwise.

### GetAccountDetailsOk

`func (o *IdentityAccountRequestProperties) GetAccountDetailsOk() (*string, bool)`

GetAccountDetailsOk returns a tuple with the AccountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDetails

`func (o *IdentityAccountRequestProperties) SetAccountDetails(v string)`

SetAccountDetails sets AccountDetails field to given value.

### HasAccountDetails

`func (o *IdentityAccountRequestProperties) HasAccountDetails() bool`

HasAccountDetails returns a boolean if a field has been set.

### SetAccountDetailsNil

`func (o *IdentityAccountRequestProperties) SetAccountDetailsNil(b bool)`

 SetAccountDetailsNil sets the value for AccountDetails to be an explicit nil

### UnsetAccountDetails
`func (o *IdentityAccountRequestProperties) UnsetAccountDetails()`

UnsetAccountDetails ensures that no value is present for AccountDetails, not even an explicit nil
### GetCurrency

`func (o *IdentityAccountRequestProperties) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *IdentityAccountRequestProperties) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *IdentityAccountRequestProperties) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *IdentityAccountRequestProperties) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTokenId

`func (o *IdentityAccountRequestProperties) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *IdentityAccountRequestProperties) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *IdentityAccountRequestProperties) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *IdentityAccountRequestProperties) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetAccountId

`func (o *IdentityAccountRequestProperties) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *IdentityAccountRequestProperties) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *IdentityAccountRequestProperties) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *IdentityAccountRequestProperties) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPaymentCode

`func (o *IdentityAccountRequestProperties) GetPaymentCode() string`

GetPaymentCode returns the PaymentCode field if non-nil, zero value otherwise.

### GetPaymentCodeOk

`func (o *IdentityAccountRequestProperties) GetPaymentCodeOk() (*string, bool)`

GetPaymentCodeOk returns a tuple with the PaymentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCode

`func (o *IdentityAccountRequestProperties) SetPaymentCode(v string)`

SetPaymentCode sets PaymentCode field to given value.

### HasPaymentCode

`func (o *IdentityAccountRequestProperties) HasPaymentCode() bool`

HasPaymentCode returns a boolean if a field has been set.

### GetExpiresAt

`func (o *IdentityAccountRequestProperties) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *IdentityAccountRequestProperties) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *IdentityAccountRequestProperties) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *IdentityAccountRequestProperties) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *IdentityAccountRequestProperties) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *IdentityAccountRequestProperties) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetQrString

`func (o *IdentityAccountRequestProperties) GetQrString() string`

GetQrString returns the QrString field if non-nil, zero value otherwise.

### GetQrStringOk

`func (o *IdentityAccountRequestProperties) GetQrStringOk() (*string, bool)`

GetQrStringOk returns a tuple with the QrString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrString

`func (o *IdentityAccountRequestProperties) SetQrString(v string)`

SetQrString sets QrString field to given value.

### HasQrString

`func (o *IdentityAccountRequestProperties) HasQrString() bool`

HasQrString returns a boolean if a field has been set.


[[Back to README]](../../README.md)


