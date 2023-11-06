# Address


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Id** | Pointer to **string** |  |  |  |
| **Category** | **NullableString** | ☑️ |  |  |
| **Country** | **string** | ☑️ |  |  |
| **ProvinceState** | **NullableString** | ☑️ |  |  |
| **City** | **NullableString** | ☑️ |  |  |
| **PostalCode** | **NullableString** | ☑️ |  |  |
| **StreetLine1** | **NullableString** | ☑️ |  |  |
| **StreetLine2** | **NullableString** | ☑️ |  |  |
| **Status** | Pointer to [**NullableAddressStatus**](AddressStatus.md) |  |  |  |
| **IsPrimary** | **NullableBool** | ☑️ |  |  |
| **Meta** | Pointer to **map[string]interface{}** |  |  |  |
| **Created** | Pointer to **time.Time** |  |  |  |
| **Updated** | Pointer to **time.Time** |  |  |  |

## Methods

### NewAddress

`func NewAddress(category NullableString, country string, provinceState NullableString, city NullableString, postalCode NullableString, streetLine1 NullableString, streetLine2 NullableString, isPrimary NullableBool, ) *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Address) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Address) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Address) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Address) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCategory

`func (o *Address) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Address) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Address) SetCategory(v string)`

SetCategory sets Category field to given value.


### SetCategoryNil

`func (o *Address) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *Address) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetCountry

`func (o *Address) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Address) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Address) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetProvinceState

`func (o *Address) GetProvinceState() string`

GetProvinceState returns the ProvinceState field if non-nil, zero value otherwise.

### GetProvinceStateOk

`func (o *Address) GetProvinceStateOk() (*string, bool)`

GetProvinceStateOk returns a tuple with the ProvinceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvinceState

`func (o *Address) SetProvinceState(v string)`

SetProvinceState sets ProvinceState field to given value.


### SetProvinceStateNil

`func (o *Address) SetProvinceStateNil(b bool)`

 SetProvinceStateNil sets the value for ProvinceState to be an explicit nil

### UnsetProvinceState
`func (o *Address) UnsetProvinceState()`

UnsetProvinceState ensures that no value is present for ProvinceState, not even an explicit nil
### GetCity

`func (o *Address) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Address) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Address) SetCity(v string)`

SetCity sets City field to given value.


### SetCityNil

`func (o *Address) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *Address) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetPostalCode

`func (o *Address) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Address) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Address) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### SetPostalCodeNil

`func (o *Address) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *Address) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetStreetLine1

`func (o *Address) GetStreetLine1() string`

GetStreetLine1 returns the StreetLine1 field if non-nil, zero value otherwise.

### GetStreetLine1Ok

`func (o *Address) GetStreetLine1Ok() (*string, bool)`

GetStreetLine1Ok returns a tuple with the StreetLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetLine1

`func (o *Address) SetStreetLine1(v string)`

SetStreetLine1 sets StreetLine1 field to given value.


### SetStreetLine1Nil

`func (o *Address) SetStreetLine1Nil(b bool)`

 SetStreetLine1Nil sets the value for StreetLine1 to be an explicit nil

### UnsetStreetLine1
`func (o *Address) UnsetStreetLine1()`

UnsetStreetLine1 ensures that no value is present for StreetLine1, not even an explicit nil
### GetStreetLine2

`func (o *Address) GetStreetLine2() string`

GetStreetLine2 returns the StreetLine2 field if non-nil, zero value otherwise.

### GetStreetLine2Ok

`func (o *Address) GetStreetLine2Ok() (*string, bool)`

GetStreetLine2Ok returns a tuple with the StreetLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetLine2

`func (o *Address) SetStreetLine2(v string)`

SetStreetLine2 sets StreetLine2 field to given value.


### SetStreetLine2Nil

`func (o *Address) SetStreetLine2Nil(b bool)`

 SetStreetLine2Nil sets the value for StreetLine2 to be an explicit nil

### UnsetStreetLine2
`func (o *Address) UnsetStreetLine2()`

UnsetStreetLine2 ensures that no value is present for StreetLine2, not even an explicit nil
### GetStatus

`func (o *Address) GetStatus() AddressStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Address) GetStatusOk() (*AddressStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Address) SetStatus(v AddressStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Address) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Address) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Address) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetIsPrimary

`func (o *Address) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *Address) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *Address) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.


### SetIsPrimaryNil

`func (o *Address) SetIsPrimaryNil(b bool)`

 SetIsPrimaryNil sets the value for IsPrimary to be an explicit nil

### UnsetIsPrimary
`func (o *Address) UnsetIsPrimary()`

UnsetIsPrimary ensures that no value is present for IsPrimary, not even an explicit nil
### GetMeta

`func (o *Address) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Address) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Address) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Address) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *Address) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *Address) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetCreated

`func (o *Address) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Address) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Address) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Address) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *Address) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Address) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Address) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Address) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to README]](../../README.md)


