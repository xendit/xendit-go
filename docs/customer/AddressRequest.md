# AddressRequest


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Category** | Pointer to **string** |  | Home, work or provincial |  |
| **CountryCode** | Pointer to **NullableString** |  | ISO3166-2 country code |  |
| **ProvinceState** | Pointer to **string** |  |  |  |
| **City** | Pointer to **string** |  |  |  |
| **Suburb** | Pointer to **string** |  |  |  |
| **PostalCode** | Pointer to **string** |  |  |  |
| **Line1** | Pointer to **string** |  |  |  |
| **Line2** | Pointer to **string** |  |  |  |
| **Status** | Pointer to [**NullableAddressStatus**](AddressStatus.md) |  |  |  |
| **IsPrimary** | Pointer to **bool** |  |  | [false] |

## Methods

### NewAddressRequest

`func NewAddressRequest() *AddressRequest`

NewAddressRequest instantiates a new AddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressRequestWithDefaults

`func NewAddressRequestWithDefaults() *AddressRequest`

NewAddressRequestWithDefaults instantiates a new AddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *AddressRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AddressRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AddressRequest) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AddressRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCountryCode

`func (o *AddressRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AddressRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AddressRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *AddressRequest) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *AddressRequest) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *AddressRequest) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetProvinceState

`func (o *AddressRequest) GetProvinceState() string`

GetProvinceState returns the ProvinceState field if non-nil, zero value otherwise.

### GetProvinceStateOk

`func (o *AddressRequest) GetProvinceStateOk() (*string, bool)`

GetProvinceStateOk returns a tuple with the ProvinceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvinceState

`func (o *AddressRequest) SetProvinceState(v string)`

SetProvinceState sets ProvinceState field to given value.

### HasProvinceState

`func (o *AddressRequest) HasProvinceState() bool`

HasProvinceState returns a boolean if a field has been set.

### GetCity

`func (o *AddressRequest) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressRequest) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressRequest) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AddressRequest) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetSuburb

`func (o *AddressRequest) GetSuburb() string`

GetSuburb returns the Suburb field if non-nil, zero value otherwise.

### GetSuburbOk

`func (o *AddressRequest) GetSuburbOk() (*string, bool)`

GetSuburbOk returns a tuple with the Suburb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuburb

`func (o *AddressRequest) SetSuburb(v string)`

SetSuburb sets Suburb field to given value.

### HasSuburb

`func (o *AddressRequest) HasSuburb() bool`

HasSuburb returns a boolean if a field has been set.

### GetPostalCode

`func (o *AddressRequest) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AddressRequest) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AddressRequest) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *AddressRequest) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetLine1

`func (o *AddressRequest) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *AddressRequest) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *AddressRequest) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *AddressRequest) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### GetLine2

`func (o *AddressRequest) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *AddressRequest) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *AddressRequest) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *AddressRequest) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### GetStatus

`func (o *AddressRequest) GetStatus() AddressStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddressRequest) GetStatusOk() (*AddressStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddressRequest) SetStatus(v AddressStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddressRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AddressRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AddressRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetIsPrimary

`func (o *AddressRequest) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *AddressRequest) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *AddressRequest) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *AddressRequest) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.


[[Back to README]](../../README.md)


