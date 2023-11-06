# PaymentMethodAuthParameters


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **AuthCode** | **string** | ☑️ |  |  |

## Methods

### NewPaymentMethodAuthParameters

`func NewPaymentMethodAuthParameters(authCode string, ) *PaymentMethodAuthParameters`

NewPaymentMethodAuthParameters instantiates a new PaymentMethodAuthParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodAuthParametersWithDefaults

`func NewPaymentMethodAuthParametersWithDefaults() *PaymentMethodAuthParameters`

NewPaymentMethodAuthParametersWithDefaults instantiates a new PaymentMethodAuthParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthCode

`func (o *PaymentMethodAuthParameters) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *PaymentMethodAuthParameters) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *PaymentMethodAuthParameters) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.



[[Back to README]](../../README.md)


