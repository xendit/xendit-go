# Bank

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankCode** | [**BankCode**](BankCode.md) |  | 
**CollectionType** | **string** | The collection type for the bank details. | 
**BankBranch** | Pointer to **string** | The branch of the bank. | [optional] 
**BankAccountNumber** | Pointer to **string** | The bank account number. | [optional] 
**AccountHolderName** | **string** | The name of the account holder. | 
**TransferAmount** | Pointer to **float32** | The transfer amount. | [optional] 

## Methods

### NewBank

`func NewBank(bankCode BankCode, collectionType string, accountHolderName string, ) *Bank`

NewBank instantiates a new Bank object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankWithDefaults

`func NewBankWithDefaults() *Bank`

NewBankWithDefaults instantiates a new Bank object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankCode

`func (o *Bank) GetBankCode() BankCode`

GetBankCode returns the BankCode field if non-nil, zero value otherwise.

### GetBankCodeOk

`func (o *Bank) GetBankCodeOk() (*BankCode, bool)`

GetBankCodeOk returns a tuple with the BankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCode

`func (o *Bank) SetBankCode(v BankCode)`

SetBankCode sets BankCode field to given value.


### GetCollectionType

`func (o *Bank) GetCollectionType() string`

GetCollectionType returns the CollectionType field if non-nil, zero value otherwise.

### GetCollectionTypeOk

`func (o *Bank) GetCollectionTypeOk() (*string, bool)`

GetCollectionTypeOk returns a tuple with the CollectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionType

`func (o *Bank) SetCollectionType(v string)`

SetCollectionType sets CollectionType field to given value.


### GetBankBranch

`func (o *Bank) GetBankBranch() string`

GetBankBranch returns the BankBranch field if non-nil, zero value otherwise.

### GetBankBranchOk

`func (o *Bank) GetBankBranchOk() (*string, bool)`

GetBankBranchOk returns a tuple with the BankBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankBranch

`func (o *Bank) SetBankBranch(v string)`

SetBankBranch sets BankBranch field to given value.

### HasBankBranch

`func (o *Bank) HasBankBranch() bool`

HasBankBranch returns a boolean if a field has been set.

### GetBankAccountNumber

`func (o *Bank) GetBankAccountNumber() string`

GetBankAccountNumber returns the BankAccountNumber field if non-nil, zero value otherwise.

### GetBankAccountNumberOk

`func (o *Bank) GetBankAccountNumberOk() (*string, bool)`

GetBankAccountNumberOk returns a tuple with the BankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountNumber

`func (o *Bank) SetBankAccountNumber(v string)`

SetBankAccountNumber sets BankAccountNumber field to given value.

### HasBankAccountNumber

`func (o *Bank) HasBankAccountNumber() bool`

HasBankAccountNumber returns a boolean if a field has been set.

### GetAccountHolderName

`func (o *Bank) GetAccountHolderName() string`

GetAccountHolderName returns the AccountHolderName field if non-nil, zero value otherwise.

### GetAccountHolderNameOk

`func (o *Bank) GetAccountHolderNameOk() (*string, bool)`

GetAccountHolderNameOk returns a tuple with the AccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderName

`func (o *Bank) SetAccountHolderName(v string)`

SetAccountHolderName sets AccountHolderName field to given value.


### GetTransferAmount

`func (o *Bank) GetTransferAmount() float32`

GetTransferAmount returns the TransferAmount field if non-nil, zero value otherwise.

### GetTransferAmountOk

`func (o *Bank) GetTransferAmountOk() (*float32, bool)`

GetTransferAmountOk returns a tuple with the TransferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAmount

`func (o *Bank) SetTransferAmount(v float32)`

SetTransferAmount sets TransferAmount field to given value.

### HasTransferAmount

`func (o *Bank) HasTransferAmount() bool`

HasTransferAmount returns a boolean if a field has been set.


[[Back to README]](../../README.md)


