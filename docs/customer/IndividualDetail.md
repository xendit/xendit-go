# IndividualDetail


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **GivenNames** | Pointer to **string** |  |  |  |
| **GivenNamesNonRoman** | Pointer to **NullableString** |  |  |  |
| **MiddleName** | Pointer to **NullableString** |  |  |  |
| **Surname** | Pointer to **NullableString** |  |  |  |
| **SurnameNonRoman** | Pointer to **NullableString** |  |  |  |
| **MotherMaidenName** | Pointer to **NullableString** |  |  |  |
| **Gender** | Pointer to **NullableString** |  |  |  |
| **DateOfBirth** | Pointer to **NullableString** |  |  |  |
| **Nationality** | Pointer to **NullableString** |  | ISO3166-2 country code |  |
| **PlaceOfBirth** | Pointer to **NullableString** |  |  |  |
| **Employment** | Pointer to [**NullableEmploymentDetail**](EmploymentDetail.md) |  |  |  |

## Methods

### NewIndividualDetail

`func NewIndividualDetail() *IndividualDetail`

NewIndividualDetail instantiates a new IndividualDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndividualDetailWithDefaults

`func NewIndividualDetailWithDefaults() *IndividualDetail`

NewIndividualDetailWithDefaults instantiates a new IndividualDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGivenNames

`func (o *IndividualDetail) GetGivenNames() string`

GetGivenNames returns the GivenNames field if non-nil, zero value otherwise.

### GetGivenNamesOk

`func (o *IndividualDetail) GetGivenNamesOk() (*string, bool)`

GetGivenNamesOk returns a tuple with the GivenNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenNames

`func (o *IndividualDetail) SetGivenNames(v string)`

SetGivenNames sets GivenNames field to given value.

### HasGivenNames

`func (o *IndividualDetail) HasGivenNames() bool`

HasGivenNames returns a boolean if a field has been set.

### GetGivenNamesNonRoman

`func (o *IndividualDetail) GetGivenNamesNonRoman() string`

GetGivenNamesNonRoman returns the GivenNamesNonRoman field if non-nil, zero value otherwise.

### GetGivenNamesNonRomanOk

`func (o *IndividualDetail) GetGivenNamesNonRomanOk() (*string, bool)`

GetGivenNamesNonRomanOk returns a tuple with the GivenNamesNonRoman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenNamesNonRoman

`func (o *IndividualDetail) SetGivenNamesNonRoman(v string)`

SetGivenNamesNonRoman sets GivenNamesNonRoman field to given value.

### HasGivenNamesNonRoman

`func (o *IndividualDetail) HasGivenNamesNonRoman() bool`

HasGivenNamesNonRoman returns a boolean if a field has been set.

### SetGivenNamesNonRomanNil

`func (o *IndividualDetail) SetGivenNamesNonRomanNil(b bool)`

 SetGivenNamesNonRomanNil sets the value for GivenNamesNonRoman to be an explicit nil

### UnsetGivenNamesNonRoman
`func (o *IndividualDetail) UnsetGivenNamesNonRoman()`

UnsetGivenNamesNonRoman ensures that no value is present for GivenNamesNonRoman, not even an explicit nil
### GetMiddleName

`func (o *IndividualDetail) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *IndividualDetail) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *IndividualDetail) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *IndividualDetail) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *IndividualDetail) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *IndividualDetail) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetSurname

`func (o *IndividualDetail) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *IndividualDetail) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *IndividualDetail) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *IndividualDetail) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### SetSurnameNil

`func (o *IndividualDetail) SetSurnameNil(b bool)`

 SetSurnameNil sets the value for Surname to be an explicit nil

### UnsetSurname
`func (o *IndividualDetail) UnsetSurname()`

UnsetSurname ensures that no value is present for Surname, not even an explicit nil
### GetSurnameNonRoman

`func (o *IndividualDetail) GetSurnameNonRoman() string`

GetSurnameNonRoman returns the SurnameNonRoman field if non-nil, zero value otherwise.

### GetSurnameNonRomanOk

`func (o *IndividualDetail) GetSurnameNonRomanOk() (*string, bool)`

GetSurnameNonRomanOk returns a tuple with the SurnameNonRoman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurnameNonRoman

`func (o *IndividualDetail) SetSurnameNonRoman(v string)`

SetSurnameNonRoman sets SurnameNonRoman field to given value.

### HasSurnameNonRoman

`func (o *IndividualDetail) HasSurnameNonRoman() bool`

HasSurnameNonRoman returns a boolean if a field has been set.

### SetSurnameNonRomanNil

`func (o *IndividualDetail) SetSurnameNonRomanNil(b bool)`

 SetSurnameNonRomanNil sets the value for SurnameNonRoman to be an explicit nil

### UnsetSurnameNonRoman
`func (o *IndividualDetail) UnsetSurnameNonRoman()`

UnsetSurnameNonRoman ensures that no value is present for SurnameNonRoman, not even an explicit nil
### GetMotherMaidenName

`func (o *IndividualDetail) GetMotherMaidenName() string`

GetMotherMaidenName returns the MotherMaidenName field if non-nil, zero value otherwise.

### GetMotherMaidenNameOk

`func (o *IndividualDetail) GetMotherMaidenNameOk() (*string, bool)`

GetMotherMaidenNameOk returns a tuple with the MotherMaidenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotherMaidenName

`func (o *IndividualDetail) SetMotherMaidenName(v string)`

SetMotherMaidenName sets MotherMaidenName field to given value.

### HasMotherMaidenName

`func (o *IndividualDetail) HasMotherMaidenName() bool`

HasMotherMaidenName returns a boolean if a field has been set.

### SetMotherMaidenNameNil

`func (o *IndividualDetail) SetMotherMaidenNameNil(b bool)`

 SetMotherMaidenNameNil sets the value for MotherMaidenName to be an explicit nil

### UnsetMotherMaidenName
`func (o *IndividualDetail) UnsetMotherMaidenName()`

UnsetMotherMaidenName ensures that no value is present for MotherMaidenName, not even an explicit nil
### GetGender

`func (o *IndividualDetail) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *IndividualDetail) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *IndividualDetail) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *IndividualDetail) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *IndividualDetail) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *IndividualDetail) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetDateOfBirth

`func (o *IndividualDetail) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *IndividualDetail) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *IndividualDetail) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *IndividualDetail) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *IndividualDetail) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *IndividualDetail) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
### GetNationality

`func (o *IndividualDetail) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *IndividualDetail) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *IndividualDetail) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *IndividualDetail) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *IndividualDetail) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *IndividualDetail) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetPlaceOfBirth

`func (o *IndividualDetail) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *IndividualDetail) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *IndividualDetail) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *IndividualDetail) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### SetPlaceOfBirthNil

`func (o *IndividualDetail) SetPlaceOfBirthNil(b bool)`

 SetPlaceOfBirthNil sets the value for PlaceOfBirth to be an explicit nil

### UnsetPlaceOfBirth
`func (o *IndividualDetail) UnsetPlaceOfBirth()`

UnsetPlaceOfBirth ensures that no value is present for PlaceOfBirth, not even an explicit nil
### GetEmployment

`func (o *IndividualDetail) GetEmployment() EmploymentDetail`

GetEmployment returns the Employment field if non-nil, zero value otherwise.

### GetEmploymentOk

`func (o *IndividualDetail) GetEmploymentOk() (*EmploymentDetail, bool)`

GetEmploymentOk returns a tuple with the Employment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployment

`func (o *IndividualDetail) SetEmployment(v EmploymentDetail)`

SetEmployment sets Employment field to given value.

### HasEmployment

`func (o *IndividualDetail) HasEmployment() bool`

HasEmployment returns a boolean if a field has been set.

### SetEmploymentNil

`func (o *IndividualDetail) SetEmploymentNil(b bool)`

 SetEmploymentNil sets the value for Employment to be an explicit nil

### UnsetEmployment
`func (o *IndividualDetail) UnsetEmployment()`

UnsetEmployment ensures that no value is present for Employment, not even an explicit nil

[[Back to README]](../../README.md)


