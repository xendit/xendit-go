# IdentityAccountResponse

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **Id** | Pointer to **string** |  | [optional]  |
| **Code** | Pointer to **NullableString** |  | [optional]  |
| **Company** | **NullableString** |  |  |
| **Description** | **NullableString** |  |  |
| **Country** | **NullableString** | ISO3166-2 country code |  |
| **HolderName** | Pointer to **NullableString** |  | [optional]  |
| **Type** | **NullableString** |  |  |
| **Properties** | [**IdentityAccountResponseProperties**](IdentityAccountResponseProperties.md) |  |  |
| **Created** | Pointer to **time.Time** |  | [optional]  |

## Methods

### NewIdentityAccountResponse

`func NewIdentityAccountResponse(company NullableString, description NullableString, country NullableString, type_ NullableString, properties IdentityAccountResponseProperties, ) *IdentityAccountResponse`

NewIdentityAccountResponse instantiates a new IdentityAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityAccountResponseWithDefaults

`func NewIdentityAccountResponseWithDefaults() *IdentityAccountResponse`

NewIdentityAccountResponseWithDefaults instantiates a new IdentityAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityAccountResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityAccountResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityAccountResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityAccountResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *IdentityAccountResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *IdentityAccountResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *IdentityAccountResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *IdentityAccountResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *IdentityAccountResponse) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *IdentityAccountResponse) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCompany

`func (o *IdentityAccountResponse) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *IdentityAccountResponse) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *IdentityAccountResponse) SetCompany(v string)`

SetCompany sets Company field to given value.


### SetCompanyNil

`func (o *IdentityAccountResponse) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *IdentityAccountResponse) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetDescription

`func (o *IdentityAccountResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentityAccountResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentityAccountResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *IdentityAccountResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IdentityAccountResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCountry

`func (o *IdentityAccountResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *IdentityAccountResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *IdentityAccountResponse) SetCountry(v string)`

SetCountry sets Country field to given value.


### SetCountryNil

`func (o *IdentityAccountResponse) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *IdentityAccountResponse) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetHolderName

`func (o *IdentityAccountResponse) GetHolderName() string`

GetHolderName returns the HolderName field if non-nil, zero value otherwise.

### GetHolderNameOk

`func (o *IdentityAccountResponse) GetHolderNameOk() (*string, bool)`

GetHolderNameOk returns a tuple with the HolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderName

`func (o *IdentityAccountResponse) SetHolderName(v string)`

SetHolderName sets HolderName field to given value.

### HasHolderName

`func (o *IdentityAccountResponse) HasHolderName() bool`

HasHolderName returns a boolean if a field has been set.

### SetHolderNameNil

`func (o *IdentityAccountResponse) SetHolderNameNil(b bool)`

 SetHolderNameNil sets the value for HolderName to be an explicit nil

### UnsetHolderName
`func (o *IdentityAccountResponse) UnsetHolderName()`

UnsetHolderName ensures that no value is present for HolderName, not even an explicit nil
### GetType

`func (o *IdentityAccountResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentityAccountResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentityAccountResponse) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *IdentityAccountResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *IdentityAccountResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetProperties

`func (o *IdentityAccountResponse) GetProperties() IdentityAccountResponseProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IdentityAccountResponse) GetPropertiesOk() (*IdentityAccountResponseProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IdentityAccountResponse) SetProperties(v IdentityAccountResponseProperties)`

SetProperties sets Properties field to given value.


### GetCreated

`func (o *IdentityAccountResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IdentityAccountResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IdentityAccountResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IdentityAccountResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to README]](../../README.md)


