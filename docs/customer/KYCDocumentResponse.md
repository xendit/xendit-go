# KYCDocumentResponse


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Country** | **string** | ☑️ |  |  |
| **Type** | [**KYCDocumentType**](KYCDocumentType.md) | ☑️ |  |  |
| **SubType** | [**KYCDocumentSubType**](KYCDocumentSubType.md) | ☑️ |  |  |
| **DocumentName** | **NullableString** | ☑️ |  |  |
| **DocumentNumber** | **NullableString** | ☑️ |  |  |
| **ExpiresAt** | **NullableString** | ☑️ |  |  |
| **HolderName** | **NullableString** | ☑️ |  |  |
| **DocumentImages** | **string[]** | ☑️ |  |  |

## Methods

### NewKYCDocumentResponse

`func NewKYCDocumentResponse(country string, type_ KYCDocumentType, subType KYCDocumentSubType, documentName NullableString, documentNumber NullableString, expiresAt NullableString, holderName NullableString, documentImages []string, ) *KYCDocumentResponse`

NewKYCDocumentResponse instantiates a new KYCDocumentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKYCDocumentResponseWithDefaults

`func NewKYCDocumentResponseWithDefaults() *KYCDocumentResponse`

NewKYCDocumentResponseWithDefaults instantiates a new KYCDocumentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *KYCDocumentResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *KYCDocumentResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *KYCDocumentResponse) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetType

`func (o *KYCDocumentResponse) GetType() KYCDocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KYCDocumentResponse) GetTypeOk() (*KYCDocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KYCDocumentResponse) SetType(v KYCDocumentType)`

SetType sets Type field to given value.


### GetSubType

`func (o *KYCDocumentResponse) GetSubType() KYCDocumentSubType`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *KYCDocumentResponse) GetSubTypeOk() (*KYCDocumentSubType, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *KYCDocumentResponse) SetSubType(v KYCDocumentSubType)`

SetSubType sets SubType field to given value.


### GetDocumentName

`func (o *KYCDocumentResponse) GetDocumentName() string`

GetDocumentName returns the DocumentName field if non-nil, zero value otherwise.

### GetDocumentNameOk

`func (o *KYCDocumentResponse) GetDocumentNameOk() (*string, bool)`

GetDocumentNameOk returns a tuple with the DocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentName

`func (o *KYCDocumentResponse) SetDocumentName(v string)`

SetDocumentName sets DocumentName field to given value.


### SetDocumentNameNil

`func (o *KYCDocumentResponse) SetDocumentNameNil(b bool)`

 SetDocumentNameNil sets the value for DocumentName to be an explicit nil

### UnsetDocumentName
`func (o *KYCDocumentResponse) UnsetDocumentName()`

UnsetDocumentName ensures that no value is present for DocumentName, not even an explicit nil
### GetDocumentNumber

`func (o *KYCDocumentResponse) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *KYCDocumentResponse) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *KYCDocumentResponse) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.


### SetDocumentNumberNil

`func (o *KYCDocumentResponse) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *KYCDocumentResponse) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetExpiresAt

`func (o *KYCDocumentResponse) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *KYCDocumentResponse) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *KYCDocumentResponse) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *KYCDocumentResponse) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *KYCDocumentResponse) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetHolderName

`func (o *KYCDocumentResponse) GetHolderName() string`

GetHolderName returns the HolderName field if non-nil, zero value otherwise.

### GetHolderNameOk

`func (o *KYCDocumentResponse) GetHolderNameOk() (*string, bool)`

GetHolderNameOk returns a tuple with the HolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderName

`func (o *KYCDocumentResponse) SetHolderName(v string)`

SetHolderName sets HolderName field to given value.


### SetHolderNameNil

`func (o *KYCDocumentResponse) SetHolderNameNil(b bool)`

 SetHolderNameNil sets the value for HolderName to be an explicit nil

### UnsetHolderName
`func (o *KYCDocumentResponse) UnsetHolderName()`

UnsetHolderName ensures that no value is present for HolderName, not even an explicit nil
### GetDocumentImages

`func (o *KYCDocumentResponse) GetDocumentImages() []string`

GetDocumentImages returns the DocumentImages field if non-nil, zero value otherwise.

### GetDocumentImagesOk

`func (o *KYCDocumentResponse) GetDocumentImagesOk() (*[]string, bool)`

GetDocumentImagesOk returns a tuple with the DocumentImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentImages

`func (o *KYCDocumentResponse) SetDocumentImages(v []string)`

SetDocumentImages sets DocumentImages field to given value.


### SetDocumentImagesNil

`func (o *KYCDocumentResponse) SetDocumentImagesNil(b bool)`

 SetDocumentImagesNil sets the value for DocumentImages to be an explicit nil

### UnsetDocumentImages
`func (o *KYCDocumentResponse) UnsetDocumentImages()`

UnsetDocumentImages ensures that no value is present for DocumentImages, not even an explicit nil

[[Back to README]](../../README.md)


