# AccountOTC

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **PaymentCode** | Pointer to **string** | Complete fixed payment code (including prefix) | [optional]  |
| **ExpiresAt** | Pointer to **NullableString** | YYYY-MM-DD string with expiry date for the payment code | [optional]  |

## Methods

### NewAccountOTC

`func NewAccountOTC() *AccountOTC`

NewAccountOTC instantiates a new AccountOTC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountOTCWithDefaults

`func NewAccountOTCWithDefaults() *AccountOTC`

NewAccountOTCWithDefaults instantiates a new AccountOTC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentCode

`func (o *AccountOTC) GetPaymentCode() string`

GetPaymentCode returns the PaymentCode field if non-nil, zero value otherwise.

### GetPaymentCodeOk

`func (o *AccountOTC) GetPaymentCodeOk() (*string, bool)`

GetPaymentCodeOk returns a tuple with the PaymentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCode

`func (o *AccountOTC) SetPaymentCode(v string)`

SetPaymentCode sets PaymentCode field to given value.

### HasPaymentCode

`func (o *AccountOTC) HasPaymentCode() bool`

HasPaymentCode returns a boolean if a field has been set.

### GetExpiresAt

`func (o *AccountOTC) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AccountOTC) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AccountOTC) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AccountOTC) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *AccountOTC) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *AccountOTC) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to README]](../../README.md)


