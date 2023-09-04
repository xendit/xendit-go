# ValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **float32** |  | 
**Error** | **string** |  | 
**Message** | **string** |  | 
**Validation** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewValidationError

`func NewValidationError(statusCode float32, error_ string, message string, ) *ValidationError`

NewValidationError instantiates a new ValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorWithDefaults

`func NewValidationErrorWithDefaults() *ValidationError`

NewValidationErrorWithDefaults instantiates a new ValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *ValidationError) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ValidationError) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ValidationError) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.


### GetError

`func (o *ValidationError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ValidationError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ValidationError) SetError(v string)`

SetError sets Error field to given value.


### GetMessage

`func (o *ValidationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidationError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidationError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetValidation

`func (o *ValidationError) GetValidation() map[string]interface{}`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *ValidationError) GetValidationOk() (*map[string]interface{}, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *ValidationError) SetValidation(v map[string]interface{})`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *ValidationError) HasValidation() bool`

HasValidation returns a boolean if a field has been set.

### SetValidationNil

`func (o *ValidationError) SetValidationNil(b bool)`

 SetValidationNil sets the value for Validation to be an explicit nil

### UnsetValidation
`func (o *ValidationError) UnsetValidation()`

UnsetValidation ensures that no value is present for Validation, not even an explicit nil

[[Back to README]](../../README.md)


