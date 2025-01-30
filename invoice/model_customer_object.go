/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.8.7
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
)

// checks if the CustomerObject type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CustomerObject{}

// CustomerObject An object representing a customer with various properties, including addresses.
type CustomerObject struct {
	// The unique identifier for the customer.
	Id NullableString `json:"id,omitempty"`
	// The customer's phone number.
	PhoneNumber NullableString `json:"phone_number,omitempty"`
	// The customer's given names or first names.
	GivenNames NullableString `json:"given_names,omitempty"`
	// The customer's surname or last name.
	Surname NullableString `json:"surname,omitempty"`
	// The customer's email address.
	Email NullableString `json:"email,omitempty"`
	// The customer's mobile phone number.
	MobileNumber NullableString `json:"mobile_number,omitempty"`
	// An additional identifier for the customer.
	CustomerId NullableString `json:"customer_id,omitempty"`
	// An array of addresses associated with the customer.
	Addresses []AddressObject `json:"addresses,omitempty"`
}

// NewCustomerObject instantiates a new CustomerObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerObject() *CustomerObject {
	this := CustomerObject{}
	return &this
}

// NewCustomerObjectWithDefaults instantiates a new CustomerObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerObjectWithDefaults() *CustomerObject {
	this := CustomerObject{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerObject) GetId() string {
	if o == nil || utils.IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CustomerObject) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *CustomerObject) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *CustomerObject) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CustomerObject) UnsetId() {
	o.Id.Unset()
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerObject) GetPhoneNumber() string {
	if o == nil || utils.IsNil(o.PhoneNumber.Get()) {
		var ret string
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerObject) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *CustomerObject) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullableString and assigns it to the PhoneNumber field.
func (o *CustomerObject) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *CustomerObject) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *CustomerObject) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetGivenNames returns the GivenNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerObject) GetGivenNames() string {
	if o == nil || utils.IsNil(o.GivenNames.Get()) {
		var ret string
		return ret
	}
	return *o.GivenNames.Get()
}

// GetGivenNamesOk returns a tuple with the GivenNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerObject) GetGivenNamesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GivenNames.Get(), o.GivenNames.IsSet()
}

// HasGivenNames returns a boolean if a field has been set.
func (o *CustomerObject) HasGivenNames() bool {
	if o != nil && o.GivenNames.IsSet() {
		return true
	}

	return false
}

// SetGivenNames gets a reference to the given NullableString and assigns it to the GivenNames field.
func (o *CustomerObject) SetGivenNames(v string) {
	o.GivenNames.Set(&v)
}
// SetGivenNamesNil sets the value for GivenNames to be an explicit nil
func (o *CustomerObject) SetGivenNamesNil() {
	o.GivenNames.Set(nil)
}

// UnsetGivenNames ensures that no value is present for GivenNames, not even an explicit nil
func (o *CustomerObject) UnsetGivenNames() {
	o.GivenNames.Unset()
}

// GetSurname returns the Surname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerObject) GetSurname() string {
	if o == nil || utils.IsNil(o.Surname.Get()) {
		var ret string
		return ret
	}
	return *o.Surname.Get()
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerObject) GetSurnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Surname.Get(), o.Surname.IsSet()
}

// HasSurname returns a boolean if a field has been set.
func (o *CustomerObject) HasSurname() bool {
	if o != nil && o.Surname.IsSet() {
		return true
	}

	return false
}

// SetSurname gets a reference to the given NullableString and assigns it to the Surname field.
func (o *CustomerObject) SetSurname(v string) {
	o.Surname.Set(&v)
}
// SetSurnameNil sets the value for Surname to be an explicit nil
func (o *CustomerObject) SetSurnameNil() {
	o.Surname.Set(nil)
}

// UnsetSurname ensures that no value is present for Surname, not even an explicit nil
func (o *CustomerObject) UnsetSurname() {
	o.Surname.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerObject) GetEmail() string {
	if o == nil || utils.IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerObject) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *CustomerObject) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *CustomerObject) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *CustomerObject) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *CustomerObject) UnsetEmail() {
	o.Email.Unset()
}

// GetMobileNumber returns the MobileNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerObject) GetMobileNumber() string {
	if o == nil || utils.IsNil(o.MobileNumber.Get()) {
		var ret string
		return ret
	}
	return *o.MobileNumber.Get()
}

// GetMobileNumberOk returns a tuple with the MobileNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerObject) GetMobileNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileNumber.Get(), o.MobileNumber.IsSet()
}

// HasMobileNumber returns a boolean if a field has been set.
func (o *CustomerObject) HasMobileNumber() bool {
	if o != nil && o.MobileNumber.IsSet() {
		return true
	}

	return false
}

// SetMobileNumber gets a reference to the given NullableString and assigns it to the MobileNumber field.
func (o *CustomerObject) SetMobileNumber(v string) {
	o.MobileNumber.Set(&v)
}
// SetMobileNumberNil sets the value for MobileNumber to be an explicit nil
func (o *CustomerObject) SetMobileNumberNil() {
	o.MobileNumber.Set(nil)
}

// UnsetMobileNumber ensures that no value is present for MobileNumber, not even an explicit nil
func (o *CustomerObject) UnsetMobileNumber() {
	o.MobileNumber.Unset()
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerObject) GetCustomerId() string {
	if o == nil || utils.IsNil(o.CustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerObject) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *CustomerObject) HasCustomerId() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *CustomerObject) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}
// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *CustomerObject) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *CustomerObject) UnsetCustomerId() {
	o.CustomerId.Unset()
}

// GetAddresses returns the Addresses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerObject) GetAddresses() []AddressObject {
	if o == nil {
		var ret []AddressObject
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerObject) GetAddressesOk() ([]AddressObject, bool) {
	if o == nil || utils.IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *CustomerObject) HasAddresses() bool {
	if o != nil && utils.IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []AddressObject and assigns it to the Addresses field.
func (o *CustomerObject) SetAddresses(v []AddressObject) {
	o.Addresses = v
}

func (o CustomerObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
    }
	if o.PhoneNumber.IsSet() {
		toSerialize["phone_number"] = o.PhoneNumber.Get()
    }
	if o.GivenNames.IsSet() {
		toSerialize["given_names"] = o.GivenNames.Get()
    }
	if o.Surname.IsSet() {
		toSerialize["surname"] = o.Surname.Get()
    }
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
    }
	if o.MobileNumber.IsSet() {
		toSerialize["mobile_number"] = o.MobileNumber.Get()
    }
	if o.CustomerId.IsSet() {
		toSerialize["customer_id"] = o.CustomerId.Get()
    }
	if o.Addresses != nil {
		toSerialize["addresses"] = o.Addresses
    }
	return toSerialize, nil
}

type NullableCustomerObject struct {
	value *CustomerObject
	isSet bool
}

func (v NullableCustomerObject) Get() *CustomerObject {
	return v.value
}

func (v *NullableCustomerObject) Set(val *CustomerObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerObject(val *CustomerObject) *NullableCustomerObject {
	return &NullableCustomerObject{value: val, isSet: true}
}

func (v NullableCustomerObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


