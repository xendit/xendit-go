# InvoiceError404ResponseDefinition
An error object used to indicate that the requested resource, in this case, an invoice, was not found.

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **ErrorCode** | **string** | ☑️ | The specific error code indicating that the requested invoice was not found. |  |
| **Message** | **string** | ☑️ | A human-readable error message providing additional context about the resource not being found. |  |

## Methods

### NewInvoiceError404ResponseDefinition

`func NewInvoiceError404ResponseDefinition(errorCode string, message string, ) *InvoiceError404ResponseDefinition`

NewInvoiceError404ResponseDefinition instantiates a new InvoiceError404ResponseDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceError404ResponseDefinitionWithDefaults

`func NewInvoiceError404ResponseDefinitionWithDefaults() *InvoiceError404ResponseDefinition`

NewInvoiceError404ResponseDefinitionWithDefaults instantiates a new InvoiceError404ResponseDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *InvoiceError404ResponseDefinition) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *InvoiceError404ResponseDefinition) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *InvoiceError404ResponseDefinition) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetMessage

`func (o *InvoiceError404ResponseDefinition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InvoiceError404ResponseDefinition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InvoiceError404ResponseDefinition) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to README]](../../README.md)


