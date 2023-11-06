# InvoiceNotFoundError
Response definition for a 404 Not Found error when creating an invoice.

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **ErrorCode** | **string** | ☑️ | The error code indicating the type of error that occurred. |  |
| **Message** | **string** | ☑️ | A human-readable error message that provides additional information about the error. |  |

## Methods

### NewInvoiceNotFoundError

`func NewInvoiceNotFoundError(errorCode string, message string, ) *InvoiceNotFoundError`

NewInvoiceNotFoundError instantiates a new InvoiceNotFoundError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceNotFoundErrorWithDefaults

`func NewInvoiceNotFoundErrorWithDefaults() *InvoiceNotFoundError`

NewInvoiceNotFoundErrorWithDefaults instantiates a new InvoiceNotFoundError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *InvoiceNotFoundError) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *InvoiceNotFoundError) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *InvoiceNotFoundError) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetMessage

`func (o *InvoiceNotFoundError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InvoiceNotFoundError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InvoiceNotFoundError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to README]](../../README.md)


