# DigitalPayoutChannelProperties

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **AccountHolderName** | Pointer to **NullableString** | Registered account name | [optional]  |
| **AccountNumber** | **string** | Registered account number |  |
| **AccountType** | Pointer to [**ChannelAccountType**](ChannelAccountType.md) |  | [optional]  |

## Methods

### NewDigitalPayoutChannelProperties

`func NewDigitalPayoutChannelProperties(accountNumber string, ) *DigitalPayoutChannelProperties`

NewDigitalPayoutChannelProperties instantiates a new DigitalPayoutChannelProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalPayoutChannelPropertiesWithDefaults

`func NewDigitalPayoutChannelPropertiesWithDefaults() *DigitalPayoutChannelProperties`

NewDigitalPayoutChannelPropertiesWithDefaults instantiates a new DigitalPayoutChannelProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolderName

`func (o *DigitalPayoutChannelProperties) GetAccountHolderName() string`

GetAccountHolderName returns the AccountHolderName field if non-nil, zero value otherwise.

### GetAccountHolderNameOk

`func (o *DigitalPayoutChannelProperties) GetAccountHolderNameOk() (*string, bool)`

GetAccountHolderNameOk returns a tuple with the AccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderName

`func (o *DigitalPayoutChannelProperties) SetAccountHolderName(v string)`

SetAccountHolderName sets AccountHolderName field to given value.

### HasAccountHolderName

`func (o *DigitalPayoutChannelProperties) HasAccountHolderName() bool`

HasAccountHolderName returns a boolean if a field has been set.

### SetAccountHolderNameNil

`func (o *DigitalPayoutChannelProperties) SetAccountHolderNameNil(b bool)`

 SetAccountHolderNameNil sets the value for AccountHolderName to be an explicit nil

### UnsetAccountHolderName
`func (o *DigitalPayoutChannelProperties) UnsetAccountHolderName()`

UnsetAccountHolderName ensures that no value is present for AccountHolderName, not even an explicit nil
### GetAccountNumber

`func (o *DigitalPayoutChannelProperties) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *DigitalPayoutChannelProperties) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *DigitalPayoutChannelProperties) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAccountType

`func (o *DigitalPayoutChannelProperties) GetAccountType() ChannelAccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *DigitalPayoutChannelProperties) GetAccountTypeOk() (*ChannelAccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *DigitalPayoutChannelProperties) SetAccountType(v ChannelAccountType)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *DigitalPayoutChannelProperties) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.


[[Back to README]](../../README.md)


