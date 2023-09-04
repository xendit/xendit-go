# CaptureParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceId** | Pointer to **NullableString** |  | [optional] 
**CaptureAmount** | **float64** |  | 

## Methods

### NewCaptureParameters

`func NewCaptureParameters(captureAmount float64, ) *CaptureParameters`

NewCaptureParameters instantiates a new CaptureParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureParametersWithDefaults

`func NewCaptureParametersWithDefaults() *CaptureParameters`

NewCaptureParametersWithDefaults instantiates a new CaptureParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceId

`func (o *CaptureParameters) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *CaptureParameters) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *CaptureParameters) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *CaptureParameters) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### SetReferenceIdNil

`func (o *CaptureParameters) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *CaptureParameters) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetCaptureAmount

`func (o *CaptureParameters) GetCaptureAmount() float64`

GetCaptureAmount returns the CaptureAmount field if non-nil, zero value otherwise.

### GetCaptureAmountOk

`func (o *CaptureParameters) GetCaptureAmountOk() (*float64, bool)`

GetCaptureAmountOk returns a tuple with the CaptureAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureAmount

`func (o *CaptureParameters) SetCaptureAmount(v float64)`

SetCaptureAmount sets CaptureAmount field to given value.



[[Back to README]](../../README.md)


