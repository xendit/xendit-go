# BusinessDetail

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **BusinessName** | Pointer to **string** |  | [optional]  |
| **BusinessType** | Pointer to **NullableString** |  | [optional]  |
| **DateOfRegistration** | Pointer to **NullableString** |  | [optional]  |
| **NatureOfBusiness** | Pointer to **NullableString** |  | [optional]  |
| **BusinessDomicile** | Pointer to **NullableString** | ISO3166-2 country code | [optional]  |

## Methods

### NewBusinessDetail

`func NewBusinessDetail() *BusinessDetail`

NewBusinessDetail instantiates a new BusinessDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessDetailWithDefaults

`func NewBusinessDetailWithDefaults() *BusinessDetail`

NewBusinessDetailWithDefaults instantiates a new BusinessDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessName

`func (o *BusinessDetail) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *BusinessDetail) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *BusinessDetail) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.

### HasBusinessName

`func (o *BusinessDetail) HasBusinessName() bool`

HasBusinessName returns a boolean if a field has been set.

### GetBusinessType

`func (o *BusinessDetail) GetBusinessType() string`

GetBusinessType returns the BusinessType field if non-nil, zero value otherwise.

### GetBusinessTypeOk

`func (o *BusinessDetail) GetBusinessTypeOk() (*string, bool)`

GetBusinessTypeOk returns a tuple with the BusinessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessType

`func (o *BusinessDetail) SetBusinessType(v string)`

SetBusinessType sets BusinessType field to given value.

### HasBusinessType

`func (o *BusinessDetail) HasBusinessType() bool`

HasBusinessType returns a boolean if a field has been set.

### SetBusinessTypeNil

`func (o *BusinessDetail) SetBusinessTypeNil(b bool)`

 SetBusinessTypeNil sets the value for BusinessType to be an explicit nil

### UnsetBusinessType
`func (o *BusinessDetail) UnsetBusinessType()`

UnsetBusinessType ensures that no value is present for BusinessType, not even an explicit nil
### GetDateOfRegistration

`func (o *BusinessDetail) GetDateOfRegistration() string`

GetDateOfRegistration returns the DateOfRegistration field if non-nil, zero value otherwise.

### GetDateOfRegistrationOk

`func (o *BusinessDetail) GetDateOfRegistrationOk() (*string, bool)`

GetDateOfRegistrationOk returns a tuple with the DateOfRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfRegistration

`func (o *BusinessDetail) SetDateOfRegistration(v string)`

SetDateOfRegistration sets DateOfRegistration field to given value.

### HasDateOfRegistration

`func (o *BusinessDetail) HasDateOfRegistration() bool`

HasDateOfRegistration returns a boolean if a field has been set.

### SetDateOfRegistrationNil

`func (o *BusinessDetail) SetDateOfRegistrationNil(b bool)`

 SetDateOfRegistrationNil sets the value for DateOfRegistration to be an explicit nil

### UnsetDateOfRegistration
`func (o *BusinessDetail) UnsetDateOfRegistration()`

UnsetDateOfRegistration ensures that no value is present for DateOfRegistration, not even an explicit nil
### GetNatureOfBusiness

`func (o *BusinessDetail) GetNatureOfBusiness() string`

GetNatureOfBusiness returns the NatureOfBusiness field if non-nil, zero value otherwise.

### GetNatureOfBusinessOk

`func (o *BusinessDetail) GetNatureOfBusinessOk() (*string, bool)`

GetNatureOfBusinessOk returns a tuple with the NatureOfBusiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatureOfBusiness

`func (o *BusinessDetail) SetNatureOfBusiness(v string)`

SetNatureOfBusiness sets NatureOfBusiness field to given value.

### HasNatureOfBusiness

`func (o *BusinessDetail) HasNatureOfBusiness() bool`

HasNatureOfBusiness returns a boolean if a field has been set.

### SetNatureOfBusinessNil

`func (o *BusinessDetail) SetNatureOfBusinessNil(b bool)`

 SetNatureOfBusinessNil sets the value for NatureOfBusiness to be an explicit nil

### UnsetNatureOfBusiness
`func (o *BusinessDetail) UnsetNatureOfBusiness()`

UnsetNatureOfBusiness ensures that no value is present for NatureOfBusiness, not even an explicit nil
### GetBusinessDomicile

`func (o *BusinessDetail) GetBusinessDomicile() string`

GetBusinessDomicile returns the BusinessDomicile field if non-nil, zero value otherwise.

### GetBusinessDomicileOk

`func (o *BusinessDetail) GetBusinessDomicileOk() (*string, bool)`

GetBusinessDomicileOk returns a tuple with the BusinessDomicile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessDomicile

`func (o *BusinessDetail) SetBusinessDomicile(v string)`

SetBusinessDomicile sets BusinessDomicile field to given value.

### HasBusinessDomicile

`func (o *BusinessDetail) HasBusinessDomicile() bool`

HasBusinessDomicile returns a boolean if a field has been set.

### SetBusinessDomicileNil

`func (o *BusinessDetail) SetBusinessDomicileNil(b bool)`

 SetBusinessDomicileNil sets the value for BusinessDomicile to be an explicit nil

### UnsetBusinessDomicile
`func (o *BusinessDetail) UnsetBusinessDomicile()`

UnsetBusinessDomicile ensures that no value is present for BusinessDomicile, not even an explicit nil

[[Back to README]](../../README.md)


