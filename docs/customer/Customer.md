# Customer

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **Type** | **string** |  | [default to "INDIVIDUAL"] |
| **ReferenceId** | **string** | Merchant&#39;s reference of this end customer, eg Merchant&#39;s user&#39;s id. Must be unique. |  |
| **IndividualDetail** | [**NullableIndividualDetail**](IndividualDetail.md) |  |  |
| **BusinessDetail** | [**NullableBusinessDetail**](BusinessDetail.md) |  |  |
| **Description** | **NullableString** |  |  |
| **Email** | **NullableString** |  |  |
| **MobileNumber** | **NullableString** |  |  |
| **PhoneNumber** | **NullableString** |  |  |
| **Addresses** | [**Address[]**](Address.md) |  |  |
| **IdentityAccounts** | [**IdentityAccountResponse[]**](IdentityAccountResponse.md) |  |  |
| **KycDocuments** | [**KYCDocumentResponse[]**](KYCDocumentResponse.md) |  |  |
| **Metadata** | **map[string]interface{}** |  |  |
| **Status** | Pointer to [**NullableEndCustomerStatus**](EndCustomerStatus.md) |  | [optional]  |
| **Id** | **string** |  |  |
| **Created** | **time.Time** |  |  |
| **Updated** | **time.Time** |  |  |

## Methods

### NewCustomer

`func NewCustomer(type_ string, referenceId string, individualDetail NullableIndividualDetail, businessDetail NullableBusinessDetail, description NullableString, email NullableString, mobileNumber NullableString, phoneNumber NullableString, addresses []Address, identityAccounts []IdentityAccountResponse, kycDocuments []KYCDocumentResponse, metadata map[string]interface{}, id string, created time.Time, updated time.Time, ) *Customer`

NewCustomer instantiates a new Customer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerWithDefaults

`func NewCustomerWithDefaults() *Customer`

NewCustomerWithDefaults instantiates a new Customer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Customer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Customer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Customer) SetType(v string)`

SetType sets Type field to given value.


### GetReferenceId

`func (o *Customer) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *Customer) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *Customer) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetIndividualDetail

`func (o *Customer) GetIndividualDetail() IndividualDetail`

GetIndividualDetail returns the IndividualDetail field if non-nil, zero value otherwise.

### GetIndividualDetailOk

`func (o *Customer) GetIndividualDetailOk() (*IndividualDetail, bool)`

GetIndividualDetailOk returns a tuple with the IndividualDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualDetail

`func (o *Customer) SetIndividualDetail(v IndividualDetail)`

SetIndividualDetail sets IndividualDetail field to given value.


### SetIndividualDetailNil

`func (o *Customer) SetIndividualDetailNil(b bool)`

 SetIndividualDetailNil sets the value for IndividualDetail to be an explicit nil

### UnsetIndividualDetail
`func (o *Customer) UnsetIndividualDetail()`

UnsetIndividualDetail ensures that no value is present for IndividualDetail, not even an explicit nil
### GetBusinessDetail

`func (o *Customer) GetBusinessDetail() BusinessDetail`

GetBusinessDetail returns the BusinessDetail field if non-nil, zero value otherwise.

### GetBusinessDetailOk

`func (o *Customer) GetBusinessDetailOk() (*BusinessDetail, bool)`

GetBusinessDetailOk returns a tuple with the BusinessDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessDetail

`func (o *Customer) SetBusinessDetail(v BusinessDetail)`

SetBusinessDetail sets BusinessDetail field to given value.


### SetBusinessDetailNil

`func (o *Customer) SetBusinessDetailNil(b bool)`

 SetBusinessDetailNil sets the value for BusinessDetail to be an explicit nil

### UnsetBusinessDetail
`func (o *Customer) UnsetBusinessDetail()`

UnsetBusinessDetail ensures that no value is present for BusinessDetail, not even an explicit nil
### GetDescription

`func (o *Customer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Customer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Customer) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Customer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Customer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEmail

`func (o *Customer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Customer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Customer) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *Customer) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *Customer) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetMobileNumber

`func (o *Customer) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *Customer) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *Customer) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.


### SetMobileNumberNil

`func (o *Customer) SetMobileNumberNil(b bool)`

 SetMobileNumberNil sets the value for MobileNumber to be an explicit nil

### UnsetMobileNumber
`func (o *Customer) UnsetMobileNumber()`

UnsetMobileNumber ensures that no value is present for MobileNumber, not even an explicit nil
### GetPhoneNumber

`func (o *Customer) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Customer) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Customer) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### SetPhoneNumberNil

`func (o *Customer) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *Customer) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetAddresses

`func (o *Customer) GetAddresses() []Address`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *Customer) GetAddressesOk() (*[]Address, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *Customer) SetAddresses(v []Address)`

SetAddresses sets Addresses field to given value.


### SetAddressesNil

`func (o *Customer) SetAddressesNil(b bool)`

 SetAddressesNil sets the value for Addresses to be an explicit nil

### UnsetAddresses
`func (o *Customer) UnsetAddresses()`

UnsetAddresses ensures that no value is present for Addresses, not even an explicit nil
### GetIdentityAccounts

`func (o *Customer) GetIdentityAccounts() []IdentityAccountResponse`

GetIdentityAccounts returns the IdentityAccounts field if non-nil, zero value otherwise.

### GetIdentityAccountsOk

`func (o *Customer) GetIdentityAccountsOk() (*[]IdentityAccountResponse, bool)`

GetIdentityAccountsOk returns a tuple with the IdentityAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAccounts

`func (o *Customer) SetIdentityAccounts(v []IdentityAccountResponse)`

SetIdentityAccounts sets IdentityAccounts field to given value.


### SetIdentityAccountsNil

`func (o *Customer) SetIdentityAccountsNil(b bool)`

 SetIdentityAccountsNil sets the value for IdentityAccounts to be an explicit nil

### UnsetIdentityAccounts
`func (o *Customer) UnsetIdentityAccounts()`

UnsetIdentityAccounts ensures that no value is present for IdentityAccounts, not even an explicit nil
### GetKycDocuments

`func (o *Customer) GetKycDocuments() []KYCDocumentResponse`

GetKycDocuments returns the KycDocuments field if non-nil, zero value otherwise.

### GetKycDocumentsOk

`func (o *Customer) GetKycDocumentsOk() (*[]KYCDocumentResponse, bool)`

GetKycDocumentsOk returns a tuple with the KycDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycDocuments

`func (o *Customer) SetKycDocuments(v []KYCDocumentResponse)`

SetKycDocuments sets KycDocuments field to given value.


### SetKycDocumentsNil

`func (o *Customer) SetKycDocumentsNil(b bool)`

 SetKycDocumentsNil sets the value for KycDocuments to be an explicit nil

### UnsetKycDocuments
`func (o *Customer) UnsetKycDocuments()`

UnsetKycDocuments ensures that no value is present for KycDocuments, not even an explicit nil
### GetMetadata

`func (o *Customer) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Customer) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Customer) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *Customer) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *Customer) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetStatus

`func (o *Customer) GetStatus() EndCustomerStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Customer) GetStatusOk() (*EndCustomerStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Customer) SetStatus(v EndCustomerStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Customer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Customer) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Customer) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetId

`func (o *Customer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Customer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Customer) SetId(v string)`

SetId sets Id field to given value.


### GetCreated

`func (o *Customer) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Customer) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Customer) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *Customer) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Customer) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Customer) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.



[[Back to README]](../../README.md)


