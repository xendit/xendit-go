# PaymentMethodExpireParameters

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **SuccessReturnUrl** | Pointer to **NullableString** | URL where the end customer is redirected if the unlinking authorization is successful. | [optional]  |
| **FailureReturnUrl** | Pointer to **NullableString** | URL where the end customer is redirected if the unlinking authorization is failed. | [optional]  |

## Methods

### NewPaymentMethodExpireParameters

`func NewPaymentMethodExpireParameters() *PaymentMethodExpireParameters`

NewPaymentMethodExpireParameters instantiates a new PaymentMethodExpireParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodExpireParametersWithDefaults

`func NewPaymentMethodExpireParametersWithDefaults() *PaymentMethodExpireParameters`

NewPaymentMethodExpireParametersWithDefaults instantiates a new PaymentMethodExpireParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessReturnUrl

`func (o *PaymentMethodExpireParameters) GetSuccessReturnUrl() string`

GetSuccessReturnUrl returns the SuccessReturnUrl field if non-nil, zero value otherwise.

### GetSuccessReturnUrlOk

`func (o *PaymentMethodExpireParameters) GetSuccessReturnUrlOk() (*string, bool)`

GetSuccessReturnUrlOk returns a tuple with the SuccessReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessReturnUrl

`func (o *PaymentMethodExpireParameters) SetSuccessReturnUrl(v string)`

SetSuccessReturnUrl sets SuccessReturnUrl field to given value.

### HasSuccessReturnUrl

`func (o *PaymentMethodExpireParameters) HasSuccessReturnUrl() bool`

HasSuccessReturnUrl returns a boolean if a field has been set.

### SetSuccessReturnUrlNil

`func (o *PaymentMethodExpireParameters) SetSuccessReturnUrlNil(b bool)`

 SetSuccessReturnUrlNil sets the value for SuccessReturnUrl to be an explicit nil

### UnsetSuccessReturnUrl
`func (o *PaymentMethodExpireParameters) UnsetSuccessReturnUrl()`

UnsetSuccessReturnUrl ensures that no value is present for SuccessReturnUrl, not even an explicit nil
### GetFailureReturnUrl

`func (o *PaymentMethodExpireParameters) GetFailureReturnUrl() string`

GetFailureReturnUrl returns the FailureReturnUrl field if non-nil, zero value otherwise.

### GetFailureReturnUrlOk

`func (o *PaymentMethodExpireParameters) GetFailureReturnUrlOk() (*string, bool)`

GetFailureReturnUrlOk returns a tuple with the FailureReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReturnUrl

`func (o *PaymentMethodExpireParameters) SetFailureReturnUrl(v string)`

SetFailureReturnUrl sets FailureReturnUrl field to given value.

### HasFailureReturnUrl

`func (o *PaymentMethodExpireParameters) HasFailureReturnUrl() bool`

HasFailureReturnUrl returns a boolean if a field has been set.

### SetFailureReturnUrlNil

`func (o *PaymentMethodExpireParameters) SetFailureReturnUrlNil(b bool)`

 SetFailureReturnUrlNil sets the value for FailureReturnUrl to be an explicit nil

### UnsetFailureReturnUrl
`func (o *PaymentMethodExpireParameters) UnsetFailureReturnUrl()`

UnsetFailureReturnUrl ensures that no value is present for FailureReturnUrl, not even an explicit nil

[[Back to README]](../../README.md)


