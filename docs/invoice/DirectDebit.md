# DirectDebit
An object representing direct debit details for invoices.

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **DirectDebitType** | [**DirectDebitType**](DirectDebitType.md) | ☑️ |  |  |

## Methods

### NewDirectDebit

`func NewDirectDebit(directDebitType DirectDebitType, ) *DirectDebit`

NewDirectDebit instantiates a new DirectDebit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectDebitWithDefaults

`func NewDirectDebitWithDefaults() *DirectDebit`

NewDirectDebitWithDefaults instantiates a new DirectDebit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectDebitType

`func (o *DirectDebit) GetDirectDebitType() DirectDebitType`

GetDirectDebitType returns the DirectDebitType field if non-nil, zero value otherwise.

### GetDirectDebitTypeOk

`func (o *DirectDebit) GetDirectDebitTypeOk() (*DirectDebitType, bool)`

GetDirectDebitTypeOk returns a tuple with the DirectDebitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebitType

`func (o *DirectDebit) SetDirectDebitType(v DirectDebitType)`

SetDirectDebitType sets DirectDebitType field to given value.



[[Back to README]](../../README.md)


