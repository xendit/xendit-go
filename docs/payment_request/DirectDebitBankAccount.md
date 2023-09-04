# DirectDebitBankAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaskedBankAccountNumber** | Pointer to **NullableString** |  | [optional] 
**BankAccountHash** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDirectDebitBankAccount

`func NewDirectDebitBankAccount() *DirectDebitBankAccount`

NewDirectDebitBankAccount instantiates a new DirectDebitBankAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectDebitBankAccountWithDefaults

`func NewDirectDebitBankAccountWithDefaults() *DirectDebitBankAccount`

NewDirectDebitBankAccountWithDefaults instantiates a new DirectDebitBankAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaskedBankAccountNumber

`func (o *DirectDebitBankAccount) GetMaskedBankAccountNumber() string`

GetMaskedBankAccountNumber returns the MaskedBankAccountNumber field if non-nil, zero value otherwise.

### GetMaskedBankAccountNumberOk

`func (o *DirectDebitBankAccount) GetMaskedBankAccountNumberOk() (*string, bool)`

GetMaskedBankAccountNumberOk returns a tuple with the MaskedBankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedBankAccountNumber

`func (o *DirectDebitBankAccount) SetMaskedBankAccountNumber(v string)`

SetMaskedBankAccountNumber sets MaskedBankAccountNumber field to given value.

### HasMaskedBankAccountNumber

`func (o *DirectDebitBankAccount) HasMaskedBankAccountNumber() bool`

HasMaskedBankAccountNumber returns a boolean if a field has been set.

### SetMaskedBankAccountNumberNil

`func (o *DirectDebitBankAccount) SetMaskedBankAccountNumberNil(b bool)`

 SetMaskedBankAccountNumberNil sets the value for MaskedBankAccountNumber to be an explicit nil

### UnsetMaskedBankAccountNumber
`func (o *DirectDebitBankAccount) UnsetMaskedBankAccountNumber()`

UnsetMaskedBankAccountNumber ensures that no value is present for MaskedBankAccountNumber, not even an explicit nil
### GetBankAccountHash

`func (o *DirectDebitBankAccount) GetBankAccountHash() string`

GetBankAccountHash returns the BankAccountHash field if non-nil, zero value otherwise.

### GetBankAccountHashOk

`func (o *DirectDebitBankAccount) GetBankAccountHashOk() (*string, bool)`

GetBankAccountHashOk returns a tuple with the BankAccountHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountHash

`func (o *DirectDebitBankAccount) SetBankAccountHash(v string)`

SetBankAccountHash sets BankAccountHash field to given value.

### HasBankAccountHash

`func (o *DirectDebitBankAccount) HasBankAccountHash() bool`

HasBankAccountHash returns a boolean if a field has been set.

### SetBankAccountHashNil

`func (o *DirectDebitBankAccount) SetBankAccountHashNil(b bool)`

 SetBankAccountHashNil sets the value for BankAccountHash to be an explicit nil

### UnsetBankAccountHash
`func (o *DirectDebitBankAccount) UnsetBankAccountHash()`

UnsetBankAccountHash ensures that no value is present for BankAccountHash, not even an explicit nil

[[Back to README]](../../README.md)


