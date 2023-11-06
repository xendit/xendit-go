# CreateRefund400Response


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **ErrorCode** | Pointer to **string** |  |  |  |
| **Message** | Pointer to **string** |  |  |  |

## Methods

### NewCreateRefund400Response

`func NewCreateRefund400Response() *CreateRefund400Response`

NewCreateRefund400Response instantiates a new CreateRefund400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRefund400ResponseWithDefaults

`func NewCreateRefund400ResponseWithDefaults() *CreateRefund400Response`

NewCreateRefund400ResponseWithDefaults instantiates a new CreateRefund400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *CreateRefund400Response) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CreateRefund400Response) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CreateRefund400Response) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *CreateRefund400Response) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMessage

`func (o *CreateRefund400Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateRefund400Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateRefund400Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateRefund400Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to README]](../../README.md)


