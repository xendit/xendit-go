/*
XENDIT SDK openapi spec

XENDIT SDK openapi spec

API version: 1.0.0
*/


package customer

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
)

// checks if the CustomerRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CustomerRequest{}

// CustomerRequest struct for CustomerRequest
type CustomerRequest struct {
	// Entity's name for this client
	ClientName *string `json:"client_name,omitempty"`
	// Merchant's reference of this end customer, eg Merchant's user's id. Must be unique.
	ReferenceId string `json:"reference_id"`
	Type *string `json:"type,omitempty"`
	IndividualDetail NullableIndividualDetail `json:"individual_detail,omitempty"`
	BusinessDetail NullableBusinessDetail `json:"business_detail,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Email *string `json:"email,omitempty"`
	MobileNumber *string `json:"mobile_number,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	Addresses []AddressRequest `json:"addresses,omitempty"`
	IdentityAccounts []IdentityAccountRequest `json:"identity_accounts,omitempty"`
	KycDocuments []KYCDocumentRequest `json:"kyc_documents,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewCustomerRequest instantiates a new CustomerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerRequest(referenceId string) *CustomerRequest {
	this := CustomerRequest{}
	this.ReferenceId = referenceId
	var type_ string = "INDIVIDUAL"
	this.Type = &type_
	return &this
}

// NewCustomerRequestWithDefaults instantiates a new CustomerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerRequestWithDefaults() *CustomerRequest {
	this := CustomerRequest{}
	var type_ string = "INDIVIDUAL"
	this.Type = &type_
	return &this
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *CustomerRequest) GetClientName() string {
	if o == nil || utils.IsNil(o.ClientName) {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerRequest) GetClientNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ClientName) {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *CustomerRequest) HasClientName() bool {
	if o != nil && !utils.IsNil(o.ClientName) {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *CustomerRequest) SetClientName(v string) {
	o.ClientName = &v
}

// GetReferenceId returns the ReferenceId field value
func (o *CustomerRequest) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *CustomerRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *CustomerRequest) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CustomerRequest) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerRequest) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CustomerRequest) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CustomerRequest) SetType(v string) {
	o.Type = &v
}

// GetIndividualDetail returns the IndividualDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerRequest) GetIndividualDetail() IndividualDetail {
	if o == nil || utils.IsNil(o.IndividualDetail.Get()) {
		var ret IndividualDetail
		return ret
	}
	return *o.IndividualDetail.Get()
}

// GetIndividualDetailOk returns a tuple with the IndividualDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerRequest) GetIndividualDetailOk() (*IndividualDetail, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndividualDetail.Get(), o.IndividualDetail.IsSet()
}

// HasIndividualDetail returns a boolean if a field has been set.
func (o *CustomerRequest) HasIndividualDetail() bool {
	if o != nil && o.IndividualDetail.IsSet() {
		return true
	}

	return false
}

// SetIndividualDetail gets a reference to the given NullableIndividualDetail and assigns it to the IndividualDetail field.
func (o *CustomerRequest) SetIndividualDetail(v IndividualDetail) {
	o.IndividualDetail.Set(&v)
}
// SetIndividualDetailNil sets the value for IndividualDetail to be an explicit nil
func (o *CustomerRequest) SetIndividualDetailNil() {
	o.IndividualDetail.Set(nil)
}

// UnsetIndividualDetail ensures that no value is present for IndividualDetail, not even an explicit nil
func (o *CustomerRequest) UnsetIndividualDetail() {
	o.IndividualDetail.Unset()
}

// GetBusinessDetail returns the BusinessDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerRequest) GetBusinessDetail() BusinessDetail {
	if o == nil || utils.IsNil(o.BusinessDetail.Get()) {
		var ret BusinessDetail
		return ret
	}
	return *o.BusinessDetail.Get()
}

// GetBusinessDetailOk returns a tuple with the BusinessDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerRequest) GetBusinessDetailOk() (*BusinessDetail, bool) {
	if o == nil {
		return nil, false
	}
	return o.BusinessDetail.Get(), o.BusinessDetail.IsSet()
}

// HasBusinessDetail returns a boolean if a field has been set.
func (o *CustomerRequest) HasBusinessDetail() bool {
	if o != nil && o.BusinessDetail.IsSet() {
		return true
	}

	return false
}

// SetBusinessDetail gets a reference to the given NullableBusinessDetail and assigns it to the BusinessDetail field.
func (o *CustomerRequest) SetBusinessDetail(v BusinessDetail) {
	o.BusinessDetail.Set(&v)
}
// SetBusinessDetailNil sets the value for BusinessDetail to be an explicit nil
func (o *CustomerRequest) SetBusinessDetailNil() {
	o.BusinessDetail.Set(nil)
}

// UnsetBusinessDetail ensures that no value is present for BusinessDetail, not even an explicit nil
func (o *CustomerRequest) UnsetBusinessDetail() {
	o.BusinessDetail.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerRequest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomerRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CustomerRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CustomerRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CustomerRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CustomerRequest) GetEmail() string {
	if o == nil || utils.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerRequest) GetEmailOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CustomerRequest) HasEmail() bool {
	if o != nil && !utils.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CustomerRequest) SetEmail(v string) {
	o.Email = &v
}

// GetMobileNumber returns the MobileNumber field value if set, zero value otherwise.
func (o *CustomerRequest) GetMobileNumber() string {
	if o == nil || utils.IsNil(o.MobileNumber) {
		var ret string
		return ret
	}
	return *o.MobileNumber
}

// GetMobileNumberOk returns a tuple with the MobileNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerRequest) GetMobileNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MobileNumber) {
		return nil, false
	}
	return o.MobileNumber, true
}

// HasMobileNumber returns a boolean if a field has been set.
func (o *CustomerRequest) HasMobileNumber() bool {
	if o != nil && !utils.IsNil(o.MobileNumber) {
		return true
	}

	return false
}

// SetMobileNumber gets a reference to the given string and assigns it to the MobileNumber field.
func (o *CustomerRequest) SetMobileNumber(v string) {
	o.MobileNumber = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *CustomerRequest) GetPhoneNumber() string {
	if o == nil || utils.IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerRequest) GetPhoneNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *CustomerRequest) HasPhoneNumber() bool {
	if o != nil && !utils.IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *CustomerRequest) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *CustomerRequest) GetAddresses() []AddressRequest {
	if o == nil || utils.IsNil(o.Addresses) {
		var ret []AddressRequest
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerRequest) GetAddressesOk() ([]AddressRequest, bool) {
	if o == nil || utils.IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *CustomerRequest) HasAddresses() bool {
	if o != nil && !utils.IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []AddressRequest and assigns it to the Addresses field.
func (o *CustomerRequest) SetAddresses(v []AddressRequest) {
	o.Addresses = v
}

// GetIdentityAccounts returns the IdentityAccounts field value if set, zero value otherwise.
func (o *CustomerRequest) GetIdentityAccounts() []IdentityAccountRequest {
	if o == nil || utils.IsNil(o.IdentityAccounts) {
		var ret []IdentityAccountRequest
		return ret
	}
	return o.IdentityAccounts
}

// GetIdentityAccountsOk returns a tuple with the IdentityAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerRequest) GetIdentityAccountsOk() ([]IdentityAccountRequest, bool) {
	if o == nil || utils.IsNil(o.IdentityAccounts) {
		return nil, false
	}
	return o.IdentityAccounts, true
}

// HasIdentityAccounts returns a boolean if a field has been set.
func (o *CustomerRequest) HasIdentityAccounts() bool {
	if o != nil && !utils.IsNil(o.IdentityAccounts) {
		return true
	}

	return false
}

// SetIdentityAccounts gets a reference to the given []IdentityAccountRequest and assigns it to the IdentityAccounts field.
func (o *CustomerRequest) SetIdentityAccounts(v []IdentityAccountRequest) {
	o.IdentityAccounts = v
}

// GetKycDocuments returns the KycDocuments field value if set, zero value otherwise.
func (o *CustomerRequest) GetKycDocuments() []KYCDocumentRequest {
	if o == nil || utils.IsNil(o.KycDocuments) {
		var ret []KYCDocumentRequest
		return ret
	}
	return o.KycDocuments
}

// GetKycDocumentsOk returns a tuple with the KycDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerRequest) GetKycDocumentsOk() ([]KYCDocumentRequest, bool) {
	if o == nil || utils.IsNil(o.KycDocuments) {
		return nil, false
	}
	return o.KycDocuments, true
}

// HasKycDocuments returns a boolean if a field has been set.
func (o *CustomerRequest) HasKycDocuments() bool {
	if o != nil && !utils.IsNil(o.KycDocuments) {
		return true
	}

	return false
}

// SetKycDocuments gets a reference to the given []KYCDocumentRequest and assigns it to the KycDocuments field.
func (o *CustomerRequest) SetKycDocuments(v []KYCDocumentRequest) {
	o.KycDocuments = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CustomerRequest) GetMetadata() map[string]interface{} {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CustomerRequest) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *CustomerRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o CustomerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ClientName) {
		toSerialize["client_name"] = o.ClientName
	}
	toSerialize["reference_id"] = o.ReferenceId
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.IndividualDetail.IsSet() {
		toSerialize["individual_detail"] = o.IndividualDetail.Get()
    }
	if o.BusinessDetail.IsSet() {
		toSerialize["business_detail"] = o.BusinessDetail.Get()
    }
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
    }
	if !utils.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !utils.IsNil(o.MobileNumber) {
		toSerialize["mobile_number"] = o.MobileNumber
	}
	if !utils.IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if !utils.IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !utils.IsNil(o.IdentityAccounts) {
		toSerialize["identity_accounts"] = o.IdentityAccounts
	}
	if !utils.IsNil(o.KycDocuments) {
		toSerialize["kyc_documents"] = o.KycDocuments
	}
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableCustomerRequest struct {
	value *CustomerRequest
	isSet bool
}

func (v NullableCustomerRequest) Get() *CustomerRequest {
	return v.value
}

func (v *NullableCustomerRequest) Set(val *CustomerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerRequest(val *CustomerRequest) *NullableCustomerRequest {
	return &NullableCustomerRequest{value: val, isSet: true}
}

func (v NullableCustomerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


