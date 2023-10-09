# UpdateCustomer400Response

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **ErrorCode** | **string** |  |  |
| **Message** | **interface{}** |  |  |
| **Errors** | Pointer to **map[string]interface{}[]** |  | [optional]  |

## Methods

### NewUpdateCustomer400Response

`func NewUpdateCustomer400Response(errorCode string, message interface{}, ) *UpdateCustomer400Response`

NewUpdateCustomer400Response instantiates a new UpdateCustomer400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomer400ResponseWithDefaults

`func NewUpdateCustomer400ResponseWithDefaults() *UpdateCustomer400Response`

NewUpdateCustomer400ResponseWithDefaults instantiates a new UpdateCustomer400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *UpdateCustomer400Response) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *UpdateCustomer400Response) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *UpdateCustomer400Response) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetMessage

`func (o *UpdateCustomer400Response) GetMessage() interface{}`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateCustomer400Response) GetMessageOk() (*interface{}, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateCustomer400Response) SetMessage(v interface{})`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *UpdateCustomer400Response) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *UpdateCustomer400Response) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetErrors

`func (o *UpdateCustomer400Response) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *UpdateCustomer400Response) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *UpdateCustomer400Response) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *UpdateCustomer400Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to README]](../../README.md)


