# SimulatePaymentRequest


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Amount** | Pointer to **float64** |  |  |  |

## Methods

### NewSimulatePaymentRequest

`func NewSimulatePaymentRequest() *SimulatePaymentRequest`

NewSimulatePaymentRequest instantiates a new SimulatePaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulatePaymentRequestWithDefaults

`func NewSimulatePaymentRequestWithDefaults() *SimulatePaymentRequest`

NewSimulatePaymentRequestWithDefaults instantiates a new SimulatePaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *SimulatePaymentRequest) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SimulatePaymentRequest) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SimulatePaymentRequest) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *SimulatePaymentRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to README]](../../README.md)


