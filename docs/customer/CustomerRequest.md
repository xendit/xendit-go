# CustomerRequest


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **ClientName** | Pointer to **string** |  | Entity&#39;s name for this client |  |
| **ReferenceId** | **string** | ☑️ | Merchant&#39;s reference of this end customer, eg Merchant&#39;s user&#39;s id. Must be unique. |  |
| **Type** | Pointer to **string** |  |  | ["INDIVIDUAL"] |
| **IndividualDetail** | Pointer to [**NullableIndividualDetail**](IndividualDetail.md) |  |  |  |
| **BusinessDetail** | Pointer to [**NullableBusinessDetail**](BusinessDetail.md) |  |  |  |
| **Description** | Pointer to **NullableString** |  |  |  |
| **Email** | Pointer to **string** |  |  |  |
| **MobileNumber** | Pointer to **string** |  |  |  |
| **PhoneNumber** | Pointer to **string** |  |  |  |
| **Addresses** | Pointer to [**AddressRequest[]**](AddressRequest.md) |  |  |  |
| **IdentityAccounts** | Pointer to [**IdentityAccountRequest[]**](IdentityAccountRequest.md) |  |  |  |
| **KycDocuments** | Pointer to [**KYCDocumentRequest[]**](KYCDocumentRequest.md) |  |  |  |
| **Metadata** | Pointer to **map[string]interface{}** |  |  |  |

## Methods

### NewCustomerRequest

`func NewCustomerRequest(referenceId string, ) *CustomerRequest`

NewCustomerRequest instantiates a new CustomerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerRequestWithDefaults

`func NewCustomerRequestWithDefaults() *CustomerRequest`

NewCustomerRequestWithDefaults instantiates a new CustomerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientName

`func (o *CustomerRequest) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *CustomerRequest) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *CustomerRequest) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *CustomerRequest) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetReferenceId

`func (o *CustomerRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *CustomerRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *CustomerRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetType

`func (o *CustomerRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomerRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIndividualDetail

`func (o *CustomerRequest) GetIndividualDetail() IndividualDetail`

GetIndividualDetail returns the IndividualDetail field if non-nil, zero value otherwise.

### GetIndividualDetailOk

`func (o *CustomerRequest) GetIndividualDetailOk() (*IndividualDetail, bool)`

GetIndividualDetailOk returns a tuple with the IndividualDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualDetail

`func (o *CustomerRequest) SetIndividualDetail(v IndividualDetail)`

SetIndividualDetail sets IndividualDetail field to given value.

### HasIndividualDetail

`func (o *CustomerRequest) HasIndividualDetail() bool`

HasIndividualDetail returns a boolean if a field has been set.

### SetIndividualDetailNil

`func (o *CustomerRequest) SetIndividualDetailNil(b bool)`

 SetIndividualDetailNil sets the value for IndividualDetail to be an explicit nil

### UnsetIndividualDetail
`func (o *CustomerRequest) UnsetIndividualDetail()`

UnsetIndividualDetail ensures that no value is present for IndividualDetail, not even an explicit nil
### GetBusinessDetail

`func (o *CustomerRequest) GetBusinessDetail() BusinessDetail`

GetBusinessDetail returns the BusinessDetail field if non-nil, zero value otherwise.

### GetBusinessDetailOk

`func (o *CustomerRequest) GetBusinessDetailOk() (*BusinessDetail, bool)`

GetBusinessDetailOk returns a tuple with the BusinessDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessDetail

`func (o *CustomerRequest) SetBusinessDetail(v BusinessDetail)`

SetBusinessDetail sets BusinessDetail field to given value.

### HasBusinessDetail

`func (o *CustomerRequest) HasBusinessDetail() bool`

HasBusinessDetail returns a boolean if a field has been set.

### SetBusinessDetailNil

`func (o *CustomerRequest) SetBusinessDetailNil(b bool)`

 SetBusinessDetailNil sets the value for BusinessDetail to be an explicit nil

### UnsetBusinessDetail
`func (o *CustomerRequest) UnsetBusinessDetail()`

UnsetBusinessDetail ensures that no value is present for BusinessDetail, not even an explicit nil
### GetDescription

`func (o *CustomerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CustomerRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CustomerRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEmail

`func (o *CustomerRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMobileNumber

`func (o *CustomerRequest) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *CustomerRequest) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *CustomerRequest) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *CustomerRequest) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *CustomerRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CustomerRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CustomerRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *CustomerRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetAddresses

`func (o *CustomerRequest) GetAddresses() []AddressRequest`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *CustomerRequest) GetAddressesOk() (*[]AddressRequest, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *CustomerRequest) SetAddresses(v []AddressRequest)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *CustomerRequest) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetIdentityAccounts

`func (o *CustomerRequest) GetIdentityAccounts() []IdentityAccountRequest`

GetIdentityAccounts returns the IdentityAccounts field if non-nil, zero value otherwise.

### GetIdentityAccountsOk

`func (o *CustomerRequest) GetIdentityAccountsOk() (*[]IdentityAccountRequest, bool)`

GetIdentityAccountsOk returns a tuple with the IdentityAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAccounts

`func (o *CustomerRequest) SetIdentityAccounts(v []IdentityAccountRequest)`

SetIdentityAccounts sets IdentityAccounts field to given value.

### HasIdentityAccounts

`func (o *CustomerRequest) HasIdentityAccounts() bool`

HasIdentityAccounts returns a boolean if a field has been set.

### GetKycDocuments

`func (o *CustomerRequest) GetKycDocuments() []KYCDocumentRequest`

GetKycDocuments returns the KycDocuments field if non-nil, zero value otherwise.

### GetKycDocumentsOk

`func (o *CustomerRequest) GetKycDocumentsOk() (*[]KYCDocumentRequest, bool)`

GetKycDocumentsOk returns a tuple with the KycDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycDocuments

`func (o *CustomerRequest) SetKycDocuments(v []KYCDocumentRequest)`

SetKycDocuments sets KycDocuments field to given value.

### HasKycDocuments

`func (o *CustomerRequest) HasKycDocuments() bool`

HasKycDocuments returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomerRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomerRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomerRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomerRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to README]](../../README.md)


