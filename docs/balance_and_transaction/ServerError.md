# ServerError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **float32** |  | 
**Error** | **string** |  | 
**Message** | **string** |  | 

## Methods

### NewServerError

`func NewServerError(statusCode float32, error_ string, message string, ) *ServerError`

NewServerError instantiates a new ServerError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerErrorWithDefaults

`func NewServerErrorWithDefaults() *ServerError`

NewServerErrorWithDefaults instantiates a new ServerError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *ServerError) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ServerError) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ServerError) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.


### GetError

`func (o *ServerError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ServerError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ServerError) SetError(v string)`

SetError sets Error field to given value.


### GetMessage

`func (o *ServerError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServerError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServerError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to README]](../../README.md)


