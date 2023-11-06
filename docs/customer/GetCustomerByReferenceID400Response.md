# GetCustomerByReferenceID400Response


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **ErrorCode** | **string** | ☑️ |  |  |
| **Message** | **interface{}** | ☑️ |  |  |
| **Errors** | Pointer to **map[string]interface{}[]** |  |  |  |

## Methods

### NewGetCustomerByReferenceID400Response

`func NewGetCustomerByReferenceID400Response(errorCode string, message interface{}, ) *GetCustomerByReferenceID400Response`

NewGetCustomerByReferenceID400Response instantiates a new GetCustomerByReferenceID400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomerByReferenceID400ResponseWithDefaults

`func NewGetCustomerByReferenceID400ResponseWithDefaults() *GetCustomerByReferenceID400Response`

NewGetCustomerByReferenceID400ResponseWithDefaults instantiates a new GetCustomerByReferenceID400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *GetCustomerByReferenceID400Response) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GetCustomerByReferenceID400Response) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GetCustomerByReferenceID400Response) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetMessage

`func (o *GetCustomerByReferenceID400Response) GetMessage() interface{}`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetCustomerByReferenceID400Response) GetMessageOk() (*interface{}, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetCustomerByReferenceID400Response) SetMessage(v interface{})`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *GetCustomerByReferenceID400Response) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *GetCustomerByReferenceID400Response) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetErrors

`func (o *GetCustomerByReferenceID400Response) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetCustomerByReferenceID400Response) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetCustomerByReferenceID400Response) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetCustomerByReferenceID400Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to README]](../../README.md)


