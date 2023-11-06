# DirectDebitChannelProperties
Direct Debit Channel Properties

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **SuccessReturnUrl** | Pointer to **string** |  |  |  |
| **FailureReturnUrl** | Pointer to **NullableString** |  |  |  |
| **MobileNumber** | Pointer to **NullableString** |  | Mobile number of the customer registered to the partner channel |  |
| **CardLastFour** | Pointer to **NullableString** |  | Last four digits of the debit card |  |
| **CardExpiry** | Pointer to **NullableString** |  | Expiry month and year of the debit card (in MM/YY format) |  |
| **Email** | Pointer to **NullableString** |  | Email address of the customer that is registered to the partner channel |  |
| **IdentityDocumentNumber** | Pointer to **NullableString** |  | Identity number of the customer registered to the partner channel |  |
| **RequireAuth** | Pointer to **NullableBool** |  |  |  |
| **AccountNumber** | Pointer to **NullableString** |  | Account number of the customer |  |

## Methods

### NewDirectDebitChannelProperties

`func NewDirectDebitChannelProperties() *DirectDebitChannelProperties`

NewDirectDebitChannelProperties instantiates a new DirectDebitChannelProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectDebitChannelPropertiesWithDefaults

`func NewDirectDebitChannelPropertiesWithDefaults() *DirectDebitChannelProperties`

NewDirectDebitChannelPropertiesWithDefaults instantiates a new DirectDebitChannelProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessReturnUrl

`func (o *DirectDebitChannelProperties) GetSuccessReturnUrl() string`

GetSuccessReturnUrl returns the SuccessReturnUrl field if non-nil, zero value otherwise.

### GetSuccessReturnUrlOk

`func (o *DirectDebitChannelProperties) GetSuccessReturnUrlOk() (*string, bool)`

GetSuccessReturnUrlOk returns a tuple with the SuccessReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessReturnUrl

`func (o *DirectDebitChannelProperties) SetSuccessReturnUrl(v string)`

SetSuccessReturnUrl sets SuccessReturnUrl field to given value.

### HasSuccessReturnUrl

`func (o *DirectDebitChannelProperties) HasSuccessReturnUrl() bool`

HasSuccessReturnUrl returns a boolean if a field has been set.

### GetFailureReturnUrl

`func (o *DirectDebitChannelProperties) GetFailureReturnUrl() string`

GetFailureReturnUrl returns the FailureReturnUrl field if non-nil, zero value otherwise.

### GetFailureReturnUrlOk

`func (o *DirectDebitChannelProperties) GetFailureReturnUrlOk() (*string, bool)`

GetFailureReturnUrlOk returns a tuple with the FailureReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReturnUrl

`func (o *DirectDebitChannelProperties) SetFailureReturnUrl(v string)`

SetFailureReturnUrl sets FailureReturnUrl field to given value.

### HasFailureReturnUrl

`func (o *DirectDebitChannelProperties) HasFailureReturnUrl() bool`

HasFailureReturnUrl returns a boolean if a field has been set.

### SetFailureReturnUrlNil

`func (o *DirectDebitChannelProperties) SetFailureReturnUrlNil(b bool)`

 SetFailureReturnUrlNil sets the value for FailureReturnUrl to be an explicit nil

### UnsetFailureReturnUrl
`func (o *DirectDebitChannelProperties) UnsetFailureReturnUrl()`

UnsetFailureReturnUrl ensures that no value is present for FailureReturnUrl, not even an explicit nil
### GetMobileNumber

`func (o *DirectDebitChannelProperties) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *DirectDebitChannelProperties) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *DirectDebitChannelProperties) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *DirectDebitChannelProperties) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

### SetMobileNumberNil

`func (o *DirectDebitChannelProperties) SetMobileNumberNil(b bool)`

 SetMobileNumberNil sets the value for MobileNumber to be an explicit nil

### UnsetMobileNumber
`func (o *DirectDebitChannelProperties) UnsetMobileNumber()`

UnsetMobileNumber ensures that no value is present for MobileNumber, not even an explicit nil
### GetCardLastFour

`func (o *DirectDebitChannelProperties) GetCardLastFour() string`

GetCardLastFour returns the CardLastFour field if non-nil, zero value otherwise.

### GetCardLastFourOk

`func (o *DirectDebitChannelProperties) GetCardLastFourOk() (*string, bool)`

GetCardLastFourOk returns a tuple with the CardLastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardLastFour

`func (o *DirectDebitChannelProperties) SetCardLastFour(v string)`

SetCardLastFour sets CardLastFour field to given value.

### HasCardLastFour

`func (o *DirectDebitChannelProperties) HasCardLastFour() bool`

HasCardLastFour returns a boolean if a field has been set.

### SetCardLastFourNil

`func (o *DirectDebitChannelProperties) SetCardLastFourNil(b bool)`

 SetCardLastFourNil sets the value for CardLastFour to be an explicit nil

### UnsetCardLastFour
`func (o *DirectDebitChannelProperties) UnsetCardLastFour()`

UnsetCardLastFour ensures that no value is present for CardLastFour, not even an explicit nil
### GetCardExpiry

`func (o *DirectDebitChannelProperties) GetCardExpiry() string`

GetCardExpiry returns the CardExpiry field if non-nil, zero value otherwise.

### GetCardExpiryOk

`func (o *DirectDebitChannelProperties) GetCardExpiryOk() (*string, bool)`

GetCardExpiryOk returns a tuple with the CardExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpiry

`func (o *DirectDebitChannelProperties) SetCardExpiry(v string)`

SetCardExpiry sets CardExpiry field to given value.

### HasCardExpiry

`func (o *DirectDebitChannelProperties) HasCardExpiry() bool`

HasCardExpiry returns a boolean if a field has been set.

### SetCardExpiryNil

`func (o *DirectDebitChannelProperties) SetCardExpiryNil(b bool)`

 SetCardExpiryNil sets the value for CardExpiry to be an explicit nil

### UnsetCardExpiry
`func (o *DirectDebitChannelProperties) UnsetCardExpiry()`

UnsetCardExpiry ensures that no value is present for CardExpiry, not even an explicit nil
### GetEmail

`func (o *DirectDebitChannelProperties) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DirectDebitChannelProperties) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DirectDebitChannelProperties) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DirectDebitChannelProperties) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *DirectDebitChannelProperties) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *DirectDebitChannelProperties) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetIdentityDocumentNumber

`func (o *DirectDebitChannelProperties) GetIdentityDocumentNumber() string`

GetIdentityDocumentNumber returns the IdentityDocumentNumber field if non-nil, zero value otherwise.

### GetIdentityDocumentNumberOk

`func (o *DirectDebitChannelProperties) GetIdentityDocumentNumberOk() (*string, bool)`

GetIdentityDocumentNumberOk returns a tuple with the IdentityDocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityDocumentNumber

`func (o *DirectDebitChannelProperties) SetIdentityDocumentNumber(v string)`

SetIdentityDocumentNumber sets IdentityDocumentNumber field to given value.

### HasIdentityDocumentNumber

`func (o *DirectDebitChannelProperties) HasIdentityDocumentNumber() bool`

HasIdentityDocumentNumber returns a boolean if a field has been set.

### SetIdentityDocumentNumberNil

`func (o *DirectDebitChannelProperties) SetIdentityDocumentNumberNil(b bool)`

 SetIdentityDocumentNumberNil sets the value for IdentityDocumentNumber to be an explicit nil

### UnsetIdentityDocumentNumber
`func (o *DirectDebitChannelProperties) UnsetIdentityDocumentNumber()`

UnsetIdentityDocumentNumber ensures that no value is present for IdentityDocumentNumber, not even an explicit nil
### GetRequireAuth

`func (o *DirectDebitChannelProperties) GetRequireAuth() bool`

GetRequireAuth returns the RequireAuth field if non-nil, zero value otherwise.

### GetRequireAuthOk

`func (o *DirectDebitChannelProperties) GetRequireAuthOk() (*bool, bool)`

GetRequireAuthOk returns a tuple with the RequireAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuth

`func (o *DirectDebitChannelProperties) SetRequireAuth(v bool)`

SetRequireAuth sets RequireAuth field to given value.

### HasRequireAuth

`func (o *DirectDebitChannelProperties) HasRequireAuth() bool`

HasRequireAuth returns a boolean if a field has been set.

### SetRequireAuthNil

`func (o *DirectDebitChannelProperties) SetRequireAuthNil(b bool)`

 SetRequireAuthNil sets the value for RequireAuth to be an explicit nil

### UnsetRequireAuth
`func (o *DirectDebitChannelProperties) UnsetRequireAuth()`

UnsetRequireAuth ensures that no value is present for RequireAuth, not even an explicit nil
### GetAccountNumber

`func (o *DirectDebitChannelProperties) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *DirectDebitChannelProperties) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *DirectDebitChannelProperties) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *DirectDebitChannelProperties) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### SetAccountNumberNil

`func (o *DirectDebitChannelProperties) SetAccountNumberNil(b bool)`

 SetAccountNumberNil sets the value for AccountNumber to be an explicit nil

### UnsetAccountNumber
`func (o *DirectDebitChannelProperties) UnsetAccountNumber()`

UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil

[[Back to README]](../../README.md)


