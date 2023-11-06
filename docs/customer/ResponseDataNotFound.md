# ResponseDataNotFound


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **ErrorCode** | Pointer to **string** |  |  |  |
| **Message** | Pointer to **interface{}** |  |  |  |

## Methods

### NewResponseDataNotFound

`func NewResponseDataNotFound() *ResponseDataNotFound`

NewResponseDataNotFound instantiates a new ResponseDataNotFound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDataNotFoundWithDefaults

`func NewResponseDataNotFoundWithDefaults() *ResponseDataNotFound`

NewResponseDataNotFoundWithDefaults instantiates a new ResponseDataNotFound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *ResponseDataNotFound) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ResponseDataNotFound) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ResponseDataNotFound) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ResponseDataNotFound) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMessage

`func (o *ResponseDataNotFound) GetMessage() interface{}`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResponseDataNotFound) GetMessageOk() (*interface{}, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResponseDataNotFound) SetMessage(v interface{})`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResponseDataNotFound) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ResponseDataNotFound) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ResponseDataNotFound) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to README]](../../README.md)


