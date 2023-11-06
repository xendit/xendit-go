# BillingInformation
Billing Information

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Country** | **string** | ☑️ |  |  |
| **StreetLine1** | Pointer to **NullableString** |  |  |  |
| **StreetLine2** | Pointer to **NullableString** |  |  |  |
| **City** | Pointer to **NullableString** |  |  |  |
| **ProvinceState** | Pointer to **NullableString** |  |  |  |
| **PostalCode** | Pointer to **NullableString** |  |  |  |

## Methods

### NewBillingInformation

`func NewBillingInformation(country string, ) *BillingInformation`

NewBillingInformation instantiates a new BillingInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInformationWithDefaults

`func NewBillingInformationWithDefaults() *BillingInformation`

NewBillingInformationWithDefaults instantiates a new BillingInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *BillingInformation) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BillingInformation) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BillingInformation) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetStreetLine1

`func (o *BillingInformation) GetStreetLine1() string`

GetStreetLine1 returns the StreetLine1 field if non-nil, zero value otherwise.

### GetStreetLine1Ok

`func (o *BillingInformation) GetStreetLine1Ok() (*string, bool)`

GetStreetLine1Ok returns a tuple with the StreetLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetLine1

`func (o *BillingInformation) SetStreetLine1(v string)`

SetStreetLine1 sets StreetLine1 field to given value.

### HasStreetLine1

`func (o *BillingInformation) HasStreetLine1() bool`

HasStreetLine1 returns a boolean if a field has been set.

### SetStreetLine1Nil

`func (o *BillingInformation) SetStreetLine1Nil(b bool)`

 SetStreetLine1Nil sets the value for StreetLine1 to be an explicit nil

### UnsetStreetLine1
`func (o *BillingInformation) UnsetStreetLine1()`

UnsetStreetLine1 ensures that no value is present for StreetLine1, not even an explicit nil
### GetStreetLine2

`func (o *BillingInformation) GetStreetLine2() string`

GetStreetLine2 returns the StreetLine2 field if non-nil, zero value otherwise.

### GetStreetLine2Ok

`func (o *BillingInformation) GetStreetLine2Ok() (*string, bool)`

GetStreetLine2Ok returns a tuple with the StreetLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetLine2

`func (o *BillingInformation) SetStreetLine2(v string)`

SetStreetLine2 sets StreetLine2 field to given value.

### HasStreetLine2

`func (o *BillingInformation) HasStreetLine2() bool`

HasStreetLine2 returns a boolean if a field has been set.

### SetStreetLine2Nil

`func (o *BillingInformation) SetStreetLine2Nil(b bool)`

 SetStreetLine2Nil sets the value for StreetLine2 to be an explicit nil

### UnsetStreetLine2
`func (o *BillingInformation) UnsetStreetLine2()`

UnsetStreetLine2 ensures that no value is present for StreetLine2, not even an explicit nil
### GetCity

`func (o *BillingInformation) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *BillingInformation) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *BillingInformation) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *BillingInformation) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *BillingInformation) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *BillingInformation) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetProvinceState

`func (o *BillingInformation) GetProvinceState() string`

GetProvinceState returns the ProvinceState field if non-nil, zero value otherwise.

### GetProvinceStateOk

`func (o *BillingInformation) GetProvinceStateOk() (*string, bool)`

GetProvinceStateOk returns a tuple with the ProvinceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvinceState

`func (o *BillingInformation) SetProvinceState(v string)`

SetProvinceState sets ProvinceState field to given value.

### HasProvinceState

`func (o *BillingInformation) HasProvinceState() bool`

HasProvinceState returns a boolean if a field has been set.

### SetProvinceStateNil

`func (o *BillingInformation) SetProvinceStateNil(b bool)`

 SetProvinceStateNil sets the value for ProvinceState to be an explicit nil

### UnsetProvinceState
`func (o *BillingInformation) UnsetProvinceState()`

UnsetProvinceState ensures that no value is present for ProvinceState, not even an explicit nil
### GetPostalCode

`func (o *BillingInformation) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *BillingInformation) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *BillingInformation) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *BillingInformation) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *BillingInformation) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *BillingInformation) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil

[[Back to README]](../../README.md)


