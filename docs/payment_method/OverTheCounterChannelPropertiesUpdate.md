# OverTheCounterChannelPropertiesUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerName** | Pointer to **string** | Name of customer. | [optional] 
**ExpiresAt** | Pointer to **time.Time** | The time when the payment code will be expired. The minimum is 2 hours and the maximum is 9 days for 7ELEVEN. Default expired date will be 2 days from payment code generated. | [optional] 

## Methods

### NewOverTheCounterChannelPropertiesUpdate

`func NewOverTheCounterChannelPropertiesUpdate() *OverTheCounterChannelPropertiesUpdate`

NewOverTheCounterChannelPropertiesUpdate instantiates a new OverTheCounterChannelPropertiesUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverTheCounterChannelPropertiesUpdateWithDefaults

`func NewOverTheCounterChannelPropertiesUpdateWithDefaults() *OverTheCounterChannelPropertiesUpdate`

NewOverTheCounterChannelPropertiesUpdateWithDefaults instantiates a new OverTheCounterChannelPropertiesUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerName

`func (o *OverTheCounterChannelPropertiesUpdate) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *OverTheCounterChannelPropertiesUpdate) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *OverTheCounterChannelPropertiesUpdate) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *OverTheCounterChannelPropertiesUpdate) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetExpiresAt

`func (o *OverTheCounterChannelPropertiesUpdate) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OverTheCounterChannelPropertiesUpdate) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OverTheCounterChannelPropertiesUpdate) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OverTheCounterChannelPropertiesUpdate) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to README]](../../README.md)


