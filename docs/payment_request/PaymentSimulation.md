# PaymentSimulation


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Status** | Pointer to **string** |  |  |  |
| **Message** | Pointer to **string** |  |  |  |

## Methods

### NewPaymentSimulation

`func NewPaymentSimulation() *PaymentSimulation`

NewPaymentSimulation instantiates a new PaymentSimulation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentSimulationWithDefaults

`func NewPaymentSimulationWithDefaults() *PaymentSimulation`

NewPaymentSimulationWithDefaults instantiates a new PaymentSimulation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PaymentSimulation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentSimulation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentSimulation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentSimulation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *PaymentSimulation) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PaymentSimulation) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PaymentSimulation) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PaymentSimulation) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to README]](../../README.md)


