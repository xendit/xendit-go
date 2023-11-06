# InvoiceFee
An object representing internal details for a fee associated with an invoice.

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Type** | **string** | ☑️ | The type of fee. |  |
| **Value** | **float32** | ☑️ | The value or amount of the fee. |  |

## Methods

### NewInvoiceFee

`func NewInvoiceFee(type_ string, value float32, ) *InvoiceFee`

NewInvoiceFee instantiates a new InvoiceFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceFeeWithDefaults

`func NewInvoiceFeeWithDefaults() *InvoiceFee`

NewInvoiceFeeWithDefaults instantiates a new InvoiceFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InvoiceFee) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoiceFee) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoiceFee) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *InvoiceFee) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InvoiceFee) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InvoiceFee) SetValue(v float32)`

SetValue sets Value field to given value.



[[Back to README]](../../README.md)


