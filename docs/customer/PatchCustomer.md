# PatchCustomer


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **ClientName** | Pointer to **NullableString** |  | Entity&#39;s name for this client |  |
| **ReferenceId** | Pointer to **NullableString** |  | Merchant&#39;s reference of this end customer, eg Merchant&#39;s user&#39;s id. Must be unique. |  |
| **IndividualDetail** | Pointer to [**NullableIndividualDetail**](IndividualDetail.md) |  |  |  |
| **BusinessDetail** | Pointer to [**NullableBusinessDetail**](BusinessDetail.md) |  |  |  |
| **Description** | Pointer to **NullableString** |  |  |  |
| **Email** | Pointer to **NullableString** |  |  |  |
| **MobileNumber** | Pointer to **NullableString** |  |  |  |
| **PhoneNumber** | Pointer to **NullableString** |  |  |  |
| **Metadata** | Pointer to **map[string]interface{}** |  |  |  |
| **Addresses** | Pointer to [**AddressRequest[]**](AddressRequest.md) |  |  |  |
| **IdentityAccounts** | Pointer to [**IdentityAccountRequest[]**](IdentityAccountRequest.md) |  |  |  |
| **KycDocuments** | Pointer to [**KYCDocumentRequest[]**](KYCDocumentRequest.md) |  |  |  |
| **Status** | Pointer to [**NullableEndCustomerStatus**](EndCustomerStatus.md) |  |  |  |

## Methods

### NewPatchCustomer

`func NewPatchCustomer() *PatchCustomer`

NewPatchCustomer instantiates a new PatchCustomer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchCustomerWithDefaults

`func NewPatchCustomerWithDefaults() *PatchCustomer`

NewPatchCustomerWithDefaults instantiates a new PatchCustomer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientName

`func (o *PatchCustomer) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *PatchCustomer) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *PatchCustomer) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *PatchCustomer) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### SetClientNameNil

`func (o *PatchCustomer) SetClientNameNil(b bool)`

 SetClientNameNil sets the value for ClientName to be an explicit nil

### UnsetClientName
`func (o *PatchCustomer) UnsetClientName()`

UnsetClientName ensures that no value is present for ClientName, not even an explicit nil
### GetReferenceId

`func (o *PatchCustomer) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *PatchCustomer) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *PatchCustomer) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *PatchCustomer) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### SetReferenceIdNil

`func (o *PatchCustomer) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *PatchCustomer) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetIndividualDetail

`func (o *PatchCustomer) GetIndividualDetail() IndividualDetail`

GetIndividualDetail returns the IndividualDetail field if non-nil, zero value otherwise.

### GetIndividualDetailOk

`func (o *PatchCustomer) GetIndividualDetailOk() (*IndividualDetail, bool)`

GetIndividualDetailOk returns a tuple with the IndividualDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualDetail

`func (o *PatchCustomer) SetIndividualDetail(v IndividualDetail)`

SetIndividualDetail sets IndividualDetail field to given value.

### HasIndividualDetail

`func (o *PatchCustomer) HasIndividualDetail() bool`

HasIndividualDetail returns a boolean if a field has been set.

### SetIndividualDetailNil

`func (o *PatchCustomer) SetIndividualDetailNil(b bool)`

 SetIndividualDetailNil sets the value for IndividualDetail to be an explicit nil

### UnsetIndividualDetail
`func (o *PatchCustomer) UnsetIndividualDetail()`

UnsetIndividualDetail ensures that no value is present for IndividualDetail, not even an explicit nil
### GetBusinessDetail

`func (o *PatchCustomer) GetBusinessDetail() BusinessDetail`

GetBusinessDetail returns the BusinessDetail field if non-nil, zero value otherwise.

### GetBusinessDetailOk

`func (o *PatchCustomer) GetBusinessDetailOk() (*BusinessDetail, bool)`

GetBusinessDetailOk returns a tuple with the BusinessDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessDetail

`func (o *PatchCustomer) SetBusinessDetail(v BusinessDetail)`

SetBusinessDetail sets BusinessDetail field to given value.

### HasBusinessDetail

`func (o *PatchCustomer) HasBusinessDetail() bool`

HasBusinessDetail returns a boolean if a field has been set.

### SetBusinessDetailNil

`func (o *PatchCustomer) SetBusinessDetailNil(b bool)`

 SetBusinessDetailNil sets the value for BusinessDetail to be an explicit nil

### UnsetBusinessDetail
`func (o *PatchCustomer) UnsetBusinessDetail()`

UnsetBusinessDetail ensures that no value is present for BusinessDetail, not even an explicit nil
### GetDescription

`func (o *PatchCustomer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchCustomer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchCustomer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchCustomer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchCustomer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchCustomer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEmail

`func (o *PatchCustomer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PatchCustomer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PatchCustomer) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PatchCustomer) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PatchCustomer) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PatchCustomer) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetMobileNumber

`func (o *PatchCustomer) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *PatchCustomer) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *PatchCustomer) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *PatchCustomer) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

### SetMobileNumberNil

`func (o *PatchCustomer) SetMobileNumberNil(b bool)`

 SetMobileNumberNil sets the value for MobileNumber to be an explicit nil

### UnsetMobileNumber
`func (o *PatchCustomer) UnsetMobileNumber()`

UnsetMobileNumber ensures that no value is present for MobileNumber, not even an explicit nil
### GetPhoneNumber

`func (o *PatchCustomer) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PatchCustomer) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PatchCustomer) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PatchCustomer) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *PatchCustomer) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *PatchCustomer) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetMetadata

`func (o *PatchCustomer) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PatchCustomer) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PatchCustomer) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PatchCustomer) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PatchCustomer) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PatchCustomer) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAddresses

`func (o *PatchCustomer) GetAddresses() []AddressRequest`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *PatchCustomer) GetAddressesOk() (*[]AddressRequest, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *PatchCustomer) SetAddresses(v []AddressRequest)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *PatchCustomer) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### SetAddressesNil

`func (o *PatchCustomer) SetAddressesNil(b bool)`

 SetAddressesNil sets the value for Addresses to be an explicit nil

### UnsetAddresses
`func (o *PatchCustomer) UnsetAddresses()`

UnsetAddresses ensures that no value is present for Addresses, not even an explicit nil
### GetIdentityAccounts

`func (o *PatchCustomer) GetIdentityAccounts() []IdentityAccountRequest`

GetIdentityAccounts returns the IdentityAccounts field if non-nil, zero value otherwise.

### GetIdentityAccountsOk

`func (o *PatchCustomer) GetIdentityAccountsOk() (*[]IdentityAccountRequest, bool)`

GetIdentityAccountsOk returns a tuple with the IdentityAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAccounts

`func (o *PatchCustomer) SetIdentityAccounts(v []IdentityAccountRequest)`

SetIdentityAccounts sets IdentityAccounts field to given value.

### HasIdentityAccounts

`func (o *PatchCustomer) HasIdentityAccounts() bool`

HasIdentityAccounts returns a boolean if a field has been set.

### SetIdentityAccountsNil

`func (o *PatchCustomer) SetIdentityAccountsNil(b bool)`

 SetIdentityAccountsNil sets the value for IdentityAccounts to be an explicit nil

### UnsetIdentityAccounts
`func (o *PatchCustomer) UnsetIdentityAccounts()`

UnsetIdentityAccounts ensures that no value is present for IdentityAccounts, not even an explicit nil
### GetKycDocuments

`func (o *PatchCustomer) GetKycDocuments() []KYCDocumentRequest`

GetKycDocuments returns the KycDocuments field if non-nil, zero value otherwise.

### GetKycDocumentsOk

`func (o *PatchCustomer) GetKycDocumentsOk() (*[]KYCDocumentRequest, bool)`

GetKycDocumentsOk returns a tuple with the KycDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycDocuments

`func (o *PatchCustomer) SetKycDocuments(v []KYCDocumentRequest)`

SetKycDocuments sets KycDocuments field to given value.

### HasKycDocuments

`func (o *PatchCustomer) HasKycDocuments() bool`

HasKycDocuments returns a boolean if a field has been set.

### SetKycDocumentsNil

`func (o *PatchCustomer) SetKycDocumentsNil(b bool)`

 SetKycDocumentsNil sets the value for KycDocuments to be an explicit nil

### UnsetKycDocuments
`func (o *PatchCustomer) UnsetKycDocuments()`

UnsetKycDocuments ensures that no value is present for KycDocuments, not even an explicit nil
### GetStatus

`func (o *PatchCustomer) GetStatus() EndCustomerStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchCustomer) GetStatusOk() (*EndCustomerStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchCustomer) SetStatus(v EndCustomerStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchCustomer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PatchCustomer) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PatchCustomer) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to README]](../../README.md)


