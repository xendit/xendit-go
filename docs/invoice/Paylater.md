# Paylater
An object representing paylater details for invoices.

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **PaylaterType** | [**PaylaterType**](PaylaterType.md) | ☑️ |  |  |
| **ShouldExclude** | Pointer to **bool** |  | Indicates whether this paylater option should be excluded. |  |

## Methods

### NewPaylater

`func NewPaylater(paylaterType PaylaterType, ) *Paylater`

NewPaylater instantiates a new Paylater object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylaterWithDefaults

`func NewPaylaterWithDefaults() *Paylater`

NewPaylaterWithDefaults instantiates a new Paylater object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaylaterType

`func (o *Paylater) GetPaylaterType() PaylaterType`

GetPaylaterType returns the PaylaterType field if non-nil, zero value otherwise.

### GetPaylaterTypeOk

`func (o *Paylater) GetPaylaterTypeOk() (*PaylaterType, bool)`

GetPaylaterTypeOk returns a tuple with the PaylaterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaylaterType

`func (o *Paylater) SetPaylaterType(v PaylaterType)`

SetPaylaterType sets PaylaterType field to given value.


### GetShouldExclude

`func (o *Paylater) GetShouldExclude() bool`

GetShouldExclude returns the ShouldExclude field if non-nil, zero value otherwise.

### GetShouldExcludeOk

`func (o *Paylater) GetShouldExcludeOk() (*bool, bool)`

GetShouldExcludeOk returns a tuple with the ShouldExclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldExclude

`func (o *Paylater) SetShouldExclude(v bool)`

SetShouldExclude sets ShouldExclude field to given value.

### HasShouldExclude

`func (o *Paylater) HasShouldExclude() bool`

HasShouldExclude returns a boolean if a field has been set.


[[Back to README]](../../README.md)


