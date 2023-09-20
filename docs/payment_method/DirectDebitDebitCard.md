# DirectDebitDebitCard

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **MobileNumber** | Pointer to **NullableString** | Mobile number of the customer registered to the partner channel | [optional]  |
| **CardLastFour** | Pointer to **NullableString** | Last four digits of the debit card | [optional]  |
| **CardExpiry** | Pointer to **NullableString** | Expiry month and year of the debit card (in MM/YY format) | [optional]  |
| **Email** | Pointer to **NullableString** | Email address of the customer that is registered to the partner channel | [optional]  |

## Methods

### NewDirectDebitDebitCard

`func NewDirectDebitDebitCard() *DirectDebitDebitCard`

NewDirectDebitDebitCard instantiates a new DirectDebitDebitCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectDebitDebitCardWithDefaults

`func NewDirectDebitDebitCardWithDefaults() *DirectDebitDebitCard`

NewDirectDebitDebitCardWithDefaults instantiates a new DirectDebitDebitCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMobileNumber

`func (o *DirectDebitDebitCard) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *DirectDebitDebitCard) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *DirectDebitDebitCard) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *DirectDebitDebitCard) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

### SetMobileNumberNil

`func (o *DirectDebitDebitCard) SetMobileNumberNil(b bool)`

 SetMobileNumberNil sets the value for MobileNumber to be an explicit nil

### UnsetMobileNumber
`func (o *DirectDebitDebitCard) UnsetMobileNumber()`

UnsetMobileNumber ensures that no value is present for MobileNumber, not even an explicit nil
### GetCardLastFour

`func (o *DirectDebitDebitCard) GetCardLastFour() string`

GetCardLastFour returns the CardLastFour field if non-nil, zero value otherwise.

### GetCardLastFourOk

`func (o *DirectDebitDebitCard) GetCardLastFourOk() (*string, bool)`

GetCardLastFourOk returns a tuple with the CardLastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardLastFour

`func (o *DirectDebitDebitCard) SetCardLastFour(v string)`

SetCardLastFour sets CardLastFour field to given value.

### HasCardLastFour

`func (o *DirectDebitDebitCard) HasCardLastFour() bool`

HasCardLastFour returns a boolean if a field has been set.

### SetCardLastFourNil

`func (o *DirectDebitDebitCard) SetCardLastFourNil(b bool)`

 SetCardLastFourNil sets the value for CardLastFour to be an explicit nil

### UnsetCardLastFour
`func (o *DirectDebitDebitCard) UnsetCardLastFour()`

UnsetCardLastFour ensures that no value is present for CardLastFour, not even an explicit nil
### GetCardExpiry

`func (o *DirectDebitDebitCard) GetCardExpiry() string`

GetCardExpiry returns the CardExpiry field if non-nil, zero value otherwise.

### GetCardExpiryOk

`func (o *DirectDebitDebitCard) GetCardExpiryOk() (*string, bool)`

GetCardExpiryOk returns a tuple with the CardExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpiry

`func (o *DirectDebitDebitCard) SetCardExpiry(v string)`

SetCardExpiry sets CardExpiry field to given value.

### HasCardExpiry

`func (o *DirectDebitDebitCard) HasCardExpiry() bool`

HasCardExpiry returns a boolean if a field has been set.

### SetCardExpiryNil

`func (o *DirectDebitDebitCard) SetCardExpiryNil(b bool)`

 SetCardExpiryNil sets the value for CardExpiry to be an explicit nil

### UnsetCardExpiry
`func (o *DirectDebitDebitCard) UnsetCardExpiry()`

UnsetCardExpiry ensures that no value is present for CardExpiry, not even an explicit nil
### GetEmail

`func (o *DirectDebitDebitCard) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DirectDebitDebitCard) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DirectDebitDebitCard) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DirectDebitDebitCard) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *DirectDebitDebitCard) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *DirectDebitDebitCard) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to README]](../../README.md)


