/*
XENDIT SDK openapi spec

XENDIT SDK openapi spec

API version: 1.0.0
*/


package customer

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v5/utils"
)

// checks if the PatchCustomer type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PatchCustomer{}

// PatchCustomer struct for PatchCustomer
type PatchCustomer struct {
	// Entity's name for this client
	ClientName NullableString `json:"client_name,omitempty"`
	// Merchant's reference of this end customer, eg Merchant's user's id. Must be unique.
	ReferenceId NullableString `json:"reference_id,omitempty"`
	IndividualDetail NullableIndividualDetail `json:"individual_detail,omitempty"`
	BusinessDetail NullableBusinessDetail `json:"business_detail,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Email NullableString `json:"email,omitempty"`
	MobileNumber NullableString `json:"mobile_number,omitempty"`
	PhoneNumber NullableString `json:"phone_number,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Addresses []AddressRequest `json:"addresses,omitempty"`
	IdentityAccounts []IdentityAccountRequest `json:"identity_accounts,omitempty"`
	KycDocuments []KYCDocumentRequest `json:"kyc_documents,omitempty"`
	Status NullableEndCustomerStatus `json:"status,omitempty"`
}

// NewPatchCustomer instantiates a new PatchCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchCustomer() *PatchCustomer {
	this := PatchCustomer{}
	return &this
}

// NewPatchCustomerWithDefaults instantiates a new PatchCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchCustomerWithDefaults() *PatchCustomer {
	this := PatchCustomer{}
	return &this
}

// GetClientName returns the ClientName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchCustomer) GetClientName() string {
	if o == nil || utils.IsNil(o.ClientName.Get()) {
		var ret string
		return ret
	}
	return *o.ClientName.Get()
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchCustomer) GetClientNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientName.Get(), o.ClientName.IsSet()
}

// HasClientName returns a boolean if a field has been set.
func (o *PatchCustomer) HasClientName() bool {
	if o != nil && o.ClientName.IsSet() {
		return true
	}

	return false
}

// SetClientName gets a reference to the given NullableString and assigns it to the ClientName field.
func (o *PatchCustomer) SetClientName(v string) {
	o.ClientName.Set(&v)
}
// SetClientNameNil sets the value for ClientName to be an explicit nil
func (o *PatchCustomer) SetClientNameNil() {
	o.ClientName.Set(nil)
}

// UnsetClientName ensures that no value is present for ClientName, not even an explicit nil
func (o *PatchCustomer) UnsetClientName() {
	o.ClientName.Unset()
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchCustomer) GetReferenceId() string {
	if o == nil || utils.IsNil(o.ReferenceId.Get()) {
		var ret string
		return ret
	}
	return *o.ReferenceId.Get()
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchCustomer) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceId.Get(), o.ReferenceId.IsSet()
}

// HasReferenceId returns a boolean if a field has been set.
func (o *PatchCustomer) HasReferenceId() bool {
	if o != nil && o.ReferenceId.IsSet() {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given NullableString and assigns it to the ReferenceId field.
func (o *PatchCustomer) SetReferenceId(v string) {
	o.ReferenceId.Set(&v)
}
// SetReferenceIdNil sets the value for ReferenceId to be an explicit nil
func (o *PatchCustomer) SetReferenceIdNil() {
	o.ReferenceId.Set(nil)
}

// UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
func (o *PatchCustomer) UnsetReferenceId() {
	o.ReferenceId.Unset()
}

// GetIndividualDetail returns the IndividualDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchCustomer) GetIndividualDetail() IndividualDetail {
	if o == nil || utils.IsNil(o.IndividualDetail.Get()) {
		var ret IndividualDetail
		return ret
	}
	return *o.IndividualDetail.Get()
}

// GetIndividualDetailOk returns a tuple with the IndividualDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchCustomer) GetIndividualDetailOk() (*IndividualDetail, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndividualDetail.Get(), o.IndividualDetail.IsSet()
}

// HasIndividualDetail returns a boolean if a field has been set.
func (o *PatchCustomer) HasIndividualDetail() bool {
	if o != nil && o.IndividualDetail.IsSet() {
		return true
	}

	return false
}

// SetIndividualDetail gets a reference to the given NullableIndividualDetail and assigns it to the IndividualDetail field.
func (o *PatchCustomer) SetIndividualDetail(v IndividualDetail) {
	o.IndividualDetail.Set(&v)
}
// SetIndividualDetailNil sets the value for IndividualDetail to be an explicit nil
func (o *PatchCustomer) SetIndividualDetailNil() {
	o.IndividualDetail.Set(nil)
}

// UnsetIndividualDetail ensures that no value is present for IndividualDetail, not even an explicit nil
func (o *PatchCustomer) UnsetIndividualDetail() {
	o.IndividualDetail.Unset()
}

// GetBusinessDetail returns the BusinessDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchCustomer) GetBusinessDetail() BusinessDetail {
	if o == nil || utils.IsNil(o.BusinessDetail.Get()) {
		var ret BusinessDetail
		return ret
	}
	return *o.BusinessDetail.Get()
}

// GetBusinessDetailOk returns a tuple with the BusinessDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchCustomer) GetBusinessDetailOk() (*BusinessDetail, bool) {
	if o == nil {
		return nil, false
	}
	return o.BusinessDetail.Get(), o.BusinessDetail.IsSet()
}

// HasBusinessDetail returns a boolean if a field has been set.
func (o *PatchCustomer) HasBusinessDetail() bool {
	if o != nil && o.BusinessDetail.IsSet() {
		return true
	}

	return false
}

// SetBusinessDetail gets a reference to the given NullableBusinessDetail and assigns it to the BusinessDetail field.
func (o *PatchCustomer) SetBusinessDetail(v BusinessDetail) {
	o.BusinessDetail.Set(&v)
}
// SetBusinessDetailNil sets the value for BusinessDetail to be an explicit nil
func (o *PatchCustomer) SetBusinessDetailNil() {
	o.BusinessDetail.Set(nil)
}

// UnsetBusinessDetail ensures that no value is present for BusinessDetail, not even an explicit nil
func (o *PatchCustomer) UnsetBusinessDetail() {
	o.BusinessDetail.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchCustomer) GetDescription() string {
	if o == nil || utils.IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchCustomer) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchCustomer) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PatchCustomer) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PatchCustomer) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PatchCustomer) UnsetDescription() {
	o.Description.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchCustomer) GetEmail() string {
	if o == nil || utils.IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchCustomer) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *PatchCustomer) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *PatchCustomer) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *PatchCustomer) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *PatchCustomer) UnsetEmail() {
	o.Email.Unset()
}

// GetMobileNumber returns the MobileNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchCustomer) GetMobileNumber() string {
	if o == nil || utils.IsNil(o.MobileNumber.Get()) {
		var ret string
		return ret
	}
	return *o.MobileNumber.Get()
}

// GetMobileNumberOk returns a tuple with the MobileNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchCustomer) GetMobileNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileNumber.Get(), o.MobileNumber.IsSet()
}

// HasMobileNumber returns a boolean if a field has been set.
func (o *PatchCustomer) HasMobileNumber() bool {
	if o != nil && o.MobileNumber.IsSet() {
		return true
	}

	return false
}

// SetMobileNumber gets a reference to the given NullableString and assigns it to the MobileNumber field.
func (o *PatchCustomer) SetMobileNumber(v string) {
	o.MobileNumber.Set(&v)
}
// SetMobileNumberNil sets the value for MobileNumber to be an explicit nil
func (o *PatchCustomer) SetMobileNumberNil() {
	o.MobileNumber.Set(nil)
}

// UnsetMobileNumber ensures that no value is present for MobileNumber, not even an explicit nil
func (o *PatchCustomer) UnsetMobileNumber() {
	o.MobileNumber.Unset()
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchCustomer) GetPhoneNumber() string {
	if o == nil || utils.IsNil(o.PhoneNumber.Get()) {
		var ret string
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchCustomer) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *PatchCustomer) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullableString and assigns it to the PhoneNumber field.
func (o *PatchCustomer) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *PatchCustomer) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *PatchCustomer) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchCustomer) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchCustomer) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PatchCustomer) HasMetadata() bool {
	if o != nil && utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PatchCustomer) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetAddresses returns the Addresses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchCustomer) GetAddresses() []AddressRequest {
	if o == nil {
		var ret []AddressRequest
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchCustomer) GetAddressesOk() ([]AddressRequest, bool) {
	if o == nil || utils.IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *PatchCustomer) HasAddresses() bool {
	if o != nil && utils.IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []AddressRequest and assigns it to the Addresses field.
func (o *PatchCustomer) SetAddresses(v []AddressRequest) {
	o.Addresses = v
}

// GetIdentityAccounts returns the IdentityAccounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchCustomer) GetIdentityAccounts() []IdentityAccountRequest {
	if o == nil {
		var ret []IdentityAccountRequest
		return ret
	}
	return o.IdentityAccounts
}

// GetIdentityAccountsOk returns a tuple with the IdentityAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchCustomer) GetIdentityAccountsOk() ([]IdentityAccountRequest, bool) {
	if o == nil || utils.IsNil(o.IdentityAccounts) {
		return nil, false
	}
	return o.IdentityAccounts, true
}

// HasIdentityAccounts returns a boolean if a field has been set.
func (o *PatchCustomer) HasIdentityAccounts() bool {
	if o != nil && utils.IsNil(o.IdentityAccounts) {
		return true
	}

	return false
}

// SetIdentityAccounts gets a reference to the given []IdentityAccountRequest and assigns it to the IdentityAccounts field.
func (o *PatchCustomer) SetIdentityAccounts(v []IdentityAccountRequest) {
	o.IdentityAccounts = v
}

// GetKycDocuments returns the KycDocuments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchCustomer) GetKycDocuments() []KYCDocumentRequest {
	if o == nil {
		var ret []KYCDocumentRequest
		return ret
	}
	return o.KycDocuments
}

// GetKycDocumentsOk returns a tuple with the KycDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchCustomer) GetKycDocumentsOk() ([]KYCDocumentRequest, bool) {
	if o == nil || utils.IsNil(o.KycDocuments) {
		return nil, false
	}
	return o.KycDocuments, true
}

// HasKycDocuments returns a boolean if a field has been set.
func (o *PatchCustomer) HasKycDocuments() bool {
	if o != nil && utils.IsNil(o.KycDocuments) {
		return true
	}

	return false
}

// SetKycDocuments gets a reference to the given []KYCDocumentRequest and assigns it to the KycDocuments field.
func (o *PatchCustomer) SetKycDocuments(v []KYCDocumentRequest) {
	o.KycDocuments = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchCustomer) GetStatus() EndCustomerStatus {
	if o == nil || utils.IsNil(o.Status.Get()) {
		var ret EndCustomerStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchCustomer) GetStatusOk() (*EndCustomerStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchCustomer) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableEndCustomerStatus and assigns it to the Status field.
func (o *PatchCustomer) SetStatus(v EndCustomerStatus) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *PatchCustomer) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *PatchCustomer) UnsetStatus() {
	o.Status.Unset()
}

func (o PatchCustomer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchCustomer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientName.IsSet() {
		toSerialize["client_name"] = o.ClientName.Get()
    }
	if o.ReferenceId.IsSet() {
		toSerialize["reference_id"] = o.ReferenceId.Get()
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
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
    }
	if o.MobileNumber.IsSet() {
		toSerialize["mobile_number"] = o.MobileNumber.Get()
    }
	if o.PhoneNumber.IsSet() {
		toSerialize["phone_number"] = o.PhoneNumber.Get()
    }
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
    }
	if o.Addresses != nil {
		toSerialize["addresses"] = o.Addresses
    }
	if o.IdentityAccounts != nil {
		toSerialize["identity_accounts"] = o.IdentityAccounts
    }
	if o.KycDocuments != nil {
		toSerialize["kyc_documents"] = o.KycDocuments
    }
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
    }
	return toSerialize, nil
}

type NullablePatchCustomer struct {
	value *PatchCustomer
	isSet bool
}

func (v NullablePatchCustomer) Get() *PatchCustomer {
	return v.value
}

func (v *NullablePatchCustomer) Set(val *PatchCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchCustomer(val *PatchCustomer) *NullablePatchCustomer {
	return &NullablePatchCustomer{value: val, isSet: true}
}

func (v NullablePatchCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


