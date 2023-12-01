/*
XENDIT SDK openapi spec

XENDIT SDK openapi spec

API version: 1.0.0
*/


package customer

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v4/utils"
	"time"
)

// checks if the Customer type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Customer{}

// Customer struct for Customer
type Customer struct {
	Type string `json:"type"`
	// Merchant's reference of this end customer, eg Merchant's user's id. Must be unique.
	ReferenceId string `json:"reference_id"`
	IndividualDetail NullableIndividualDetail `json:"individual_detail"`
	BusinessDetail NullableBusinessDetail `json:"business_detail"`
	Description NullableString `json:"description"`
	Email NullableString `json:"email"`
	MobileNumber NullableString `json:"mobile_number"`
	PhoneNumber NullableString `json:"phone_number"`
	Addresses []Address `json:"addresses"`
	IdentityAccounts []IdentityAccountResponse `json:"identity_accounts"`
	KycDocuments []KYCDocumentResponse `json:"kyc_documents"`
	Metadata map[string]interface{} `json:"metadata"`
	Status NullableEndCustomerStatus `json:"status,omitempty"`
	Id string `json:"id"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// NewCustomer instantiates a new Customer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomer(type_ string, referenceId string, individualDetail NullableIndividualDetail, businessDetail NullableBusinessDetail, description NullableString, email NullableString, mobileNumber NullableString, phoneNumber NullableString, addresses []Address, identityAccounts []IdentityAccountResponse, kycDocuments []KYCDocumentResponse, metadata map[string]interface{}, id string, created time.Time, updated time.Time) *Customer {
	this := Customer{}
	this.Type = type_
	this.ReferenceId = referenceId
	this.IndividualDetail = individualDetail
	this.BusinessDetail = businessDetail
	this.Description = description
	this.Email = email
	this.MobileNumber = mobileNumber
	this.PhoneNumber = phoneNumber
	this.Addresses = addresses
	this.IdentityAccounts = identityAccounts
	this.KycDocuments = kycDocuments
	this.Metadata = metadata
	this.Id = id
	this.Created = created
	this.Updated = updated
	return &this
}

// NewCustomerWithDefaults instantiates a new Customer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerWithDefaults() *Customer {
	this := Customer{}
	var type_ string = "INDIVIDUAL"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *Customer) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Customer) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Customer) SetType(v string) {
	o.Type = v
}

// GetReferenceId returns the ReferenceId field value
func (o *Customer) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *Customer) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *Customer) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetIndividualDetail returns the IndividualDetail field value
// If the value is explicit nil, the zero value for IndividualDetail will be returned
func (o *Customer) GetIndividualDetail() IndividualDetail {
	if o == nil || o.IndividualDetail.Get() == nil {
		var ret IndividualDetail
		return ret
	}

	return *o.IndividualDetail.Get()
}

// GetIndividualDetailOk returns a tuple with the IndividualDetail field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetIndividualDetailOk() (*IndividualDetail, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndividualDetail.Get(), o.IndividualDetail.IsSet()
}

// SetIndividualDetail sets field value
func (o *Customer) SetIndividualDetail(v IndividualDetail) {
	o.IndividualDetail.Set(&v)
}

// GetBusinessDetail returns the BusinessDetail field value
// If the value is explicit nil, the zero value for BusinessDetail will be returned
func (o *Customer) GetBusinessDetail() BusinessDetail {
	if o == nil || o.BusinessDetail.Get() == nil {
		var ret BusinessDetail
		return ret
	}

	return *o.BusinessDetail.Get()
}

// GetBusinessDetailOk returns a tuple with the BusinessDetail field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetBusinessDetailOk() (*BusinessDetail, bool) {
	if o == nil {
		return nil, false
	}
	return o.BusinessDetail.Get(), o.BusinessDetail.IsSet()
}

// SetBusinessDetail sets field value
func (o *Customer) SetBusinessDetail(v BusinessDetail) {
	o.BusinessDetail.Set(&v)
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Customer) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *Customer) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetEmail returns the Email field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Customer) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}

	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// SetEmail sets field value
func (o *Customer) SetEmail(v string) {
	o.Email.Set(&v)
}

// GetMobileNumber returns the MobileNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Customer) GetMobileNumber() string {
	if o == nil || o.MobileNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.MobileNumber.Get()
}

// GetMobileNumberOk returns a tuple with the MobileNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetMobileNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileNumber.Get(), o.MobileNumber.IsSet()
}

// SetMobileNumber sets field value
func (o *Customer) SetMobileNumber(v string) {
	o.MobileNumber.Set(&v)
}

// GetPhoneNumber returns the PhoneNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Customer) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// SetPhoneNumber sets field value
func (o *Customer) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}

// GetAddresses returns the Addresses field value
// If the value is explicit nil, the zero value for []Address will be returned
func (o *Customer) GetAddresses() []Address {
	if o == nil {
		var ret []Address
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetAddressesOk() ([]Address, bool) {
	if o == nil || utils.IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *Customer) SetAddresses(v []Address) {
	o.Addresses = v
}

// GetIdentityAccounts returns the IdentityAccounts field value
// If the value is explicit nil, the zero value for []IdentityAccountResponse will be returned
func (o *Customer) GetIdentityAccounts() []IdentityAccountResponse {
	if o == nil {
		var ret []IdentityAccountResponse
		return ret
	}

	return o.IdentityAccounts
}

// GetIdentityAccountsOk returns a tuple with the IdentityAccounts field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetIdentityAccountsOk() ([]IdentityAccountResponse, bool) {
	if o == nil || utils.IsNil(o.IdentityAccounts) {
		return nil, false
	}
	return o.IdentityAccounts, true
}

// SetIdentityAccounts sets field value
func (o *Customer) SetIdentityAccounts(v []IdentityAccountResponse) {
	o.IdentityAccounts = v
}

// GetKycDocuments returns the KycDocuments field value
// If the value is explicit nil, the zero value for []KYCDocumentResponse will be returned
func (o *Customer) GetKycDocuments() []KYCDocumentResponse {
	if o == nil {
		var ret []KYCDocumentResponse
		return ret
	}

	return o.KycDocuments
}

// GetKycDocumentsOk returns a tuple with the KycDocuments field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetKycDocumentsOk() ([]KYCDocumentResponse, bool) {
	if o == nil || utils.IsNil(o.KycDocuments) {
		return nil, false
	}
	return o.KycDocuments, true
}

// SetKycDocuments sets field value
func (o *Customer) SetKycDocuments(v []KYCDocumentResponse) {
	o.KycDocuments = v
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *Customer) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *Customer) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Customer) GetStatus() EndCustomerStatus {
	if o == nil || utils.IsNil(o.Status.Get()) {
		var ret EndCustomerStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetStatusOk() (*EndCustomerStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *Customer) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableEndCustomerStatus and assigns it to the Status field.
func (o *Customer) SetStatus(v EndCustomerStatus) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *Customer) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *Customer) UnsetStatus() {
	o.Status.Unset()
}

// GetId returns the Id field value
func (o *Customer) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Customer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Customer) SetId(v string) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *Customer) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Customer) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Customer) SetCreated(v time.Time) {
	o.Created = v
}

// GetUpdated returns the Updated field value
func (o *Customer) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *Customer) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *Customer) SetUpdated(v time.Time) {
	o.Updated = v
}

func (o Customer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Customer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
    if o.Type != "INDIVIDUAL" && o.Type != "BUSINESS" {
        toSerialize["type"] = nil
        return toSerialize, utils.NewError("invalid value for Type when marshalling to JSON, must be one of INDIVIDUAL, BUSINESS")
    }
	toSerialize["reference_id"] = o.ReferenceId
	toSerialize["individual_detail"] = o.IndividualDetail.Get()

	toSerialize["business_detail"] = o.BusinessDetail.Get()

	toSerialize["description"] = o.Description.Get()

	toSerialize["email"] = o.Email.Get()

	toSerialize["mobile_number"] = o.MobileNumber.Get()

	toSerialize["phone_number"] = o.PhoneNumber.Get()

	if o.Addresses != nil {
		toSerialize["addresses"] = o.Addresses
    }
	if o.IdentityAccounts != nil {
		toSerialize["identity_accounts"] = o.IdentityAccounts
    }
	if o.KycDocuments != nil {
		toSerialize["kyc_documents"] = o.KycDocuments
    }
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
    }
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
    }
	toSerialize["id"] = o.Id
	toSerialize["created"] = o.Created
	toSerialize["updated"] = o.Updated
	return toSerialize, nil
}

type NullableCustomer struct {
	value *Customer
	isSet bool
}

func (v NullableCustomer) Get() *Customer {
	return v.value
}

func (v *NullableCustomer) Set(val *Customer) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomer(val *Customer) *NullableCustomer {
	return &NullableCustomer{value: val, isSet: true}
}

func (v NullableCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


