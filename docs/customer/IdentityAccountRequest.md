# IdentityAccountRequest


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Type** | Pointer to [**IdentityAccountType**](IdentityAccountType.md) |  |  |  |
| **Company** | Pointer to **string** |  | The issuing institution associated with the account (e.g., OCBC, GOPAY, 7-11). If adding financial accounts that Xendit supports, we recommend you use the channel_name found at https://xendit.github.io/apireference/#payment-channels for this field |  |
| **Description** | Pointer to **string** |  | Free text description of this account |  |
| **Country** | Pointer to **NullableString** |  | ISO3166-2 country code |  |
| **Properties** | Pointer to [**IdentityAccountRequestProperties**](IdentityAccountRequestProperties.md) |  |  |  |

## Methods

### NewIdentityAccountRequest

`func NewIdentityAccountRequest() *IdentityAccountRequest`

NewIdentityAccountRequest instantiates a new IdentityAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityAccountRequestWithDefaults

`func NewIdentityAccountRequestWithDefaults() *IdentityAccountRequest`

NewIdentityAccountRequestWithDefaults instantiates a new IdentityAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IdentityAccountRequest) GetType() IdentityAccountType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentityAccountRequest) GetTypeOk() (*IdentityAccountType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentityAccountRequest) SetType(v IdentityAccountType)`

SetType sets Type field to given value.

### HasType

`func (o *IdentityAccountRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCompany

`func (o *IdentityAccountRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *IdentityAccountRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *IdentityAccountRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *IdentityAccountRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetDescription

`func (o *IdentityAccountRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentityAccountRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentityAccountRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdentityAccountRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCountry

`func (o *IdentityAccountRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *IdentityAccountRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *IdentityAccountRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *IdentityAccountRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *IdentityAccountRequest) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *IdentityAccountRequest) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetProperties

`func (o *IdentityAccountRequest) GetProperties() IdentityAccountRequestProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IdentityAccountRequest) GetPropertiesOk() (*IdentityAccountRequestProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IdentityAccountRequest) SetProperties(v IdentityAccountRequestProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *IdentityAccountRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to README]](../../README.md)


