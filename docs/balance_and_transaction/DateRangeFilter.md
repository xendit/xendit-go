# DateRangeFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gte** | Pointer to **time.Time** | Start time of transaction. If not specified will list all dates. | [optional] 
**Lte** | Pointer to **time.Time** | End time of transaction. If not specified will list all dates. | [optional] 

## Methods

### NewDateRangeFilter

`func NewDateRangeFilter() *DateRangeFilter`

NewDateRangeFilter instantiates a new DateRangeFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateRangeFilterWithDefaults

`func NewDateRangeFilterWithDefaults() *DateRangeFilter`

NewDateRangeFilterWithDefaults instantiates a new DateRangeFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGte

`func (o *DateRangeFilter) GetGte() time.Time`

GetGte returns the Gte field if non-nil, zero value otherwise.

### GetGteOk

`func (o *DateRangeFilter) GetGteOk() (*time.Time, bool)`

GetGteOk returns a tuple with the Gte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGte

`func (o *DateRangeFilter) SetGte(v time.Time)`

SetGte sets Gte field to given value.

### HasGte

`func (o *DateRangeFilter) HasGte() bool`

HasGte returns a boolean if a field has been set.

### GetLte

`func (o *DateRangeFilter) GetLte() time.Time`

GetLte returns the Lte field if non-nil, zero value otherwise.

### GetLteOk

`func (o *DateRangeFilter) GetLteOk() (*time.Time, bool)`

GetLteOk returns a tuple with the Lte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLte

`func (o *DateRangeFilter) SetLte(v time.Time)`

SetLte sets Lte field to given value.

### HasLte

`func (o *DateRangeFilter) HasLte() bool`

HasLte returns a boolean if a field has been set.


[[Back to README]](../../README.md)


