# ErrorErrorsInner


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Path** | **string** | ☑️ | Precise location of the error |  |
| **Message** | **string** | ☑️ | Specific description of the error |  |

## Methods

### NewErrorErrorsInner

`func NewErrorErrorsInner(path string, message string, ) *ErrorErrorsInner`

NewErrorErrorsInner instantiates a new ErrorErrorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorErrorsInnerWithDefaults

`func NewErrorErrorsInnerWithDefaults() *ErrorErrorsInner`

NewErrorErrorsInnerWithDefaults instantiates a new ErrorErrorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ErrorErrorsInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ErrorErrorsInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ErrorErrorsInner) SetPath(v string)`

SetPath sets Path field to given value.


### GetMessage

`func (o *ErrorErrorsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorErrorsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorErrorsInner) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to README]](../../README.md)


