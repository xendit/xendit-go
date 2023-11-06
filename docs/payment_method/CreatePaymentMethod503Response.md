# CreatePaymentMethod503Response


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **ErrorCode** | Pointer to **string** |  |  |  |
| **Message** | Pointer to **string** |  |  |  |

## Methods

### NewCreatePaymentMethod503Response

`func NewCreatePaymentMethod503Response() *CreatePaymentMethod503Response`

NewCreatePaymentMethod503Response instantiates a new CreatePaymentMethod503Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentMethod503ResponseWithDefaults

`func NewCreatePaymentMethod503ResponseWithDefaults() *CreatePaymentMethod503Response`

NewCreatePaymentMethod503ResponseWithDefaults instantiates a new CreatePaymentMethod503Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *CreatePaymentMethod503Response) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CreatePaymentMethod503Response) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CreatePaymentMethod503Response) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *CreatePaymentMethod503Response) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMessage

`func (o *CreatePaymentMethod503Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreatePaymentMethod503Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreatePaymentMethod503Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreatePaymentMethod503Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to README]](../../README.md)


