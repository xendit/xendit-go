# EmploymentDetail


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **EmployerName** | Pointer to **NullableString** |  | Name of employer |  |
| **NatureOfBusiness** | Pointer to **NullableString** |  | Industry or nature of business |  |
| **RoleDescription** | Pointer to **NullableString** |  | Occupation or title |  |

## Methods

### NewEmploymentDetail

`func NewEmploymentDetail() *EmploymentDetail`

NewEmploymentDetail instantiates a new EmploymentDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmploymentDetailWithDefaults

`func NewEmploymentDetailWithDefaults() *EmploymentDetail`

NewEmploymentDetailWithDefaults instantiates a new EmploymentDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployerName

`func (o *EmploymentDetail) GetEmployerName() string`

GetEmployerName returns the EmployerName field if non-nil, zero value otherwise.

### GetEmployerNameOk

`func (o *EmploymentDetail) GetEmployerNameOk() (*string, bool)`

GetEmployerNameOk returns a tuple with the EmployerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerName

`func (o *EmploymentDetail) SetEmployerName(v string)`

SetEmployerName sets EmployerName field to given value.

### HasEmployerName

`func (o *EmploymentDetail) HasEmployerName() bool`

HasEmployerName returns a boolean if a field has been set.

### SetEmployerNameNil

`func (o *EmploymentDetail) SetEmployerNameNil(b bool)`

 SetEmployerNameNil sets the value for EmployerName to be an explicit nil

### UnsetEmployerName
`func (o *EmploymentDetail) UnsetEmployerName()`

UnsetEmployerName ensures that no value is present for EmployerName, not even an explicit nil
### GetNatureOfBusiness

`func (o *EmploymentDetail) GetNatureOfBusiness() string`

GetNatureOfBusiness returns the NatureOfBusiness field if non-nil, zero value otherwise.

### GetNatureOfBusinessOk

`func (o *EmploymentDetail) GetNatureOfBusinessOk() (*string, bool)`

GetNatureOfBusinessOk returns a tuple with the NatureOfBusiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatureOfBusiness

`func (o *EmploymentDetail) SetNatureOfBusiness(v string)`

SetNatureOfBusiness sets NatureOfBusiness field to given value.

### HasNatureOfBusiness

`func (o *EmploymentDetail) HasNatureOfBusiness() bool`

HasNatureOfBusiness returns a boolean if a field has been set.

### SetNatureOfBusinessNil

`func (o *EmploymentDetail) SetNatureOfBusinessNil(b bool)`

 SetNatureOfBusinessNil sets the value for NatureOfBusiness to be an explicit nil

### UnsetNatureOfBusiness
`func (o *EmploymentDetail) UnsetNatureOfBusiness()`

UnsetNatureOfBusiness ensures that no value is present for NatureOfBusiness, not even an explicit nil
### GetRoleDescription

`func (o *EmploymentDetail) GetRoleDescription() string`

GetRoleDescription returns the RoleDescription field if non-nil, zero value otherwise.

### GetRoleDescriptionOk

`func (o *EmploymentDetail) GetRoleDescriptionOk() (*string, bool)`

GetRoleDescriptionOk returns a tuple with the RoleDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleDescription

`func (o *EmploymentDetail) SetRoleDescription(v string)`

SetRoleDescription sets RoleDescription field to given value.

### HasRoleDescription

`func (o *EmploymentDetail) HasRoleDescription() bool`

HasRoleDescription returns a boolean if a field has been set.

### SetRoleDescriptionNil

`func (o *EmploymentDetail) SetRoleDescriptionNil(b bool)`

 SetRoleDescriptionNil sets the value for RoleDescription to be an explicit nil

### UnsetRoleDescription
`func (o *EmploymentDetail) UnsetRoleDescription()`

UnsetRoleDescription ensures that no value is present for RoleDescription, not even an explicit nil

[[Back to README]](../../README.md)


