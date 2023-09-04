# ChannelProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The corresponding parameter name | [optional] 
**Type** | Pointer to **string** | Data type of the parameter | [optional] 
**IsRequired** | Pointer to **bool** | Indicates whether or not the parameter is required | [optional] 

## Methods

### NewChannelProperty

`func NewChannelProperty() *ChannelProperty`

NewChannelProperty instantiates a new ChannelProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelPropertyWithDefaults

`func NewChannelPropertyWithDefaults() *ChannelProperty`

NewChannelPropertyWithDefaults instantiates a new ChannelProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ChannelProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChannelProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChannelProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChannelProperty) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ChannelProperty) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChannelProperty) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChannelProperty) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ChannelProperty) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsRequired

`func (o *ChannelProperty) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *ChannelProperty) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *ChannelProperty) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *ChannelProperty) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.


[[Back to README]](../../README.md)


