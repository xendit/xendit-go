# KYCDocumentRequest


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Country** | Pointer to **NullableString** |  | ISO3166-2 country code |  |
| **Type** | Pointer to [**KYCDocumentType**](KYCDocumentType.md) |  |  |  |
| **SubType** | Pointer to [**NullableKYCDocumentSubType**](KYCDocumentSubType.md) |  |  |  |
| **DocumentName** | Pointer to **string** |  |  |  |
| **DocumentNumber** | Pointer to **string** |  |  |  |
| **ExpiresAt** | Pointer to **string** |  |  |  |
| **HolderName** | Pointer to **string** |  |  |  |
| **DocumentImages** | Pointer to **string[]** |  |  |  |

## Methods

### NewKYCDocumentRequest

`func NewKYCDocumentRequest() *KYCDocumentRequest`

NewKYCDocumentRequest instantiates a new KYCDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKYCDocumentRequestWithDefaults

`func NewKYCDocumentRequestWithDefaults() *KYCDocumentRequest`

NewKYCDocumentRequestWithDefaults instantiates a new KYCDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *KYCDocumentRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *KYCDocumentRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *KYCDocumentRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *KYCDocumentRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *KYCDocumentRequest) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *KYCDocumentRequest) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetType

`func (o *KYCDocumentRequest) GetType() KYCDocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KYCDocumentRequest) GetTypeOk() (*KYCDocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KYCDocumentRequest) SetType(v KYCDocumentType)`

SetType sets Type field to given value.

### HasType

`func (o *KYCDocumentRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubType

`func (o *KYCDocumentRequest) GetSubType() KYCDocumentSubType`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *KYCDocumentRequest) GetSubTypeOk() (*KYCDocumentSubType, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *KYCDocumentRequest) SetSubType(v KYCDocumentSubType)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *KYCDocumentRequest) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### SetSubTypeNil

`func (o *KYCDocumentRequest) SetSubTypeNil(b bool)`

 SetSubTypeNil sets the value for SubType to be an explicit nil

### UnsetSubType
`func (o *KYCDocumentRequest) UnsetSubType()`

UnsetSubType ensures that no value is present for SubType, not even an explicit nil
### GetDocumentName

`func (o *KYCDocumentRequest) GetDocumentName() string`

GetDocumentName returns the DocumentName field if non-nil, zero value otherwise.

### GetDocumentNameOk

`func (o *KYCDocumentRequest) GetDocumentNameOk() (*string, bool)`

GetDocumentNameOk returns a tuple with the DocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentName

`func (o *KYCDocumentRequest) SetDocumentName(v string)`

SetDocumentName sets DocumentName field to given value.

### HasDocumentName

`func (o *KYCDocumentRequest) HasDocumentName() bool`

HasDocumentName returns a boolean if a field has been set.

### GetDocumentNumber

`func (o *KYCDocumentRequest) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *KYCDocumentRequest) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *KYCDocumentRequest) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *KYCDocumentRequest) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### GetExpiresAt

`func (o *KYCDocumentRequest) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *KYCDocumentRequest) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *KYCDocumentRequest) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *KYCDocumentRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetHolderName

`func (o *KYCDocumentRequest) GetHolderName() string`

GetHolderName returns the HolderName field if non-nil, zero value otherwise.

### GetHolderNameOk

`func (o *KYCDocumentRequest) GetHolderNameOk() (*string, bool)`

GetHolderNameOk returns a tuple with the HolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderName

`func (o *KYCDocumentRequest) SetHolderName(v string)`

SetHolderName sets HolderName field to given value.

### HasHolderName

`func (o *KYCDocumentRequest) HasHolderName() bool`

HasHolderName returns a boolean if a field has been set.

### GetDocumentImages

`func (o *KYCDocumentRequest) GetDocumentImages() []string`

GetDocumentImages returns the DocumentImages field if non-nil, zero value otherwise.

### GetDocumentImagesOk

`func (o *KYCDocumentRequest) GetDocumentImagesOk() (*[]string, bool)`

GetDocumentImagesOk returns a tuple with the DocumentImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentImages

`func (o *KYCDocumentRequest) SetDocumentImages(v []string)`

SetDocumentImages sets DocumentImages field to given value.

### HasDocumentImages

`func (o *KYCDocumentRequest) HasDocumentImages() bool`

HasDocumentImages returns a boolean if a field has been set.


[[Back to README]](../../README.md)


