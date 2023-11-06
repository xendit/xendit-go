# OverTheCounterChannelProperties
Over The Counter Channel Properties

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **PaymentCode** | Pointer to **string** |  | The payment code that you want to assign, e.g 12345. If you do not send one, one will be picked at random. |  |
| **CustomerName** | **string** | ☑️ | Name of customer. |  |
| **ExpiresAt** | Pointer to **time.Time** |  | The time when the payment code will be expired. The minimum is 2 hours and the maximum is 9 days for 7ELEVEN. Default expired date will be 2 days from payment code generated. |  |

## Methods

### NewOverTheCounterChannelProperties

`func NewOverTheCounterChannelProperties(customerName string, ) *OverTheCounterChannelProperties`

NewOverTheCounterChannelProperties instantiates a new OverTheCounterChannelProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverTheCounterChannelPropertiesWithDefaults

`func NewOverTheCounterChannelPropertiesWithDefaults() *OverTheCounterChannelProperties`

NewOverTheCounterChannelPropertiesWithDefaults instantiates a new OverTheCounterChannelProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentCode

`func (o *OverTheCounterChannelProperties) GetPaymentCode() string`

GetPaymentCode returns the PaymentCode field if non-nil, zero value otherwise.

### GetPaymentCodeOk

`func (o *OverTheCounterChannelProperties) GetPaymentCodeOk() (*string, bool)`

GetPaymentCodeOk returns a tuple with the PaymentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCode

`func (o *OverTheCounterChannelProperties) SetPaymentCode(v string)`

SetPaymentCode sets PaymentCode field to given value.

### HasPaymentCode

`func (o *OverTheCounterChannelProperties) HasPaymentCode() bool`

HasPaymentCode returns a boolean if a field has been set.

### GetCustomerName

`func (o *OverTheCounterChannelProperties) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *OverTheCounterChannelProperties) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *OverTheCounterChannelProperties) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.


### GetExpiresAt

`func (o *OverTheCounterChannelProperties) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OverTheCounterChannelProperties) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OverTheCounterChannelProperties) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OverTheCounterChannelProperties) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to README]](../../README.md)


