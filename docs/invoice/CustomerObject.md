# CustomerObject

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **Id** | Pointer to **NullableString** | The unique identifier for the customer. | [optional]  |
| **PhoneNumber** | Pointer to **NullableString** | The customer&#39;s phone number. | [optional]  |
| **GivenNames** | Pointer to **NullableString** | The customer&#39;s given names or first names. | [optional]  |
| **Surname** | Pointer to **NullableString** | The customer&#39;s surname or last name. | [optional]  |
| **Email** | Pointer to **NullableString** | The customer&#39;s email address. | [optional]  |
| **MobileNumber** | Pointer to **NullableString** | The customer&#39;s mobile phone number. | [optional]  |
| **CustomerId** | Pointer to **NullableString** | An additional identifier for the customer. | [optional]  |
| **Addresses** | Pointer to [**AddressObject[]**](AddressObject.md) | An array of addresses associated with the customer. | [optional]  |

## Methods

### NewCustomerObject

`func NewCustomerObject() *CustomerObject`

NewCustomerObject instantiates a new CustomerObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerObjectWithDefaults

`func NewCustomerObjectWithDefaults() *CustomerObject`

NewCustomerObjectWithDefaults instantiates a new CustomerObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerObject) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CustomerObject) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CustomerObject) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPhoneNumber

`func (o *CustomerObject) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CustomerObject) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CustomerObject) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *CustomerObject) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *CustomerObject) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *CustomerObject) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetGivenNames

`func (o *CustomerObject) GetGivenNames() string`

GetGivenNames returns the GivenNames field if non-nil, zero value otherwise.

### GetGivenNamesOk

`func (o *CustomerObject) GetGivenNamesOk() (*string, bool)`

GetGivenNamesOk returns a tuple with the GivenNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenNames

`func (o *CustomerObject) SetGivenNames(v string)`

SetGivenNames sets GivenNames field to given value.

### HasGivenNames

`func (o *CustomerObject) HasGivenNames() bool`

HasGivenNames returns a boolean if a field has been set.

### SetGivenNamesNil

`func (o *CustomerObject) SetGivenNamesNil(b bool)`

 SetGivenNamesNil sets the value for GivenNames to be an explicit nil

### UnsetGivenNames
`func (o *CustomerObject) UnsetGivenNames()`

UnsetGivenNames ensures that no value is present for GivenNames, not even an explicit nil
### GetSurname

`func (o *CustomerObject) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *CustomerObject) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *CustomerObject) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *CustomerObject) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### SetSurnameNil

`func (o *CustomerObject) SetSurnameNil(b bool)`

 SetSurnameNil sets the value for Surname to be an explicit nil

### UnsetSurname
`func (o *CustomerObject) UnsetSurname()`

UnsetSurname ensures that no value is present for Surname, not even an explicit nil
### GetEmail

`func (o *CustomerObject) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerObject) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerObject) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerObject) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CustomerObject) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CustomerObject) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetMobileNumber

`func (o *CustomerObject) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *CustomerObject) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *CustomerObject) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *CustomerObject) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

### SetMobileNumberNil

`func (o *CustomerObject) SetMobileNumberNil(b bool)`

 SetMobileNumberNil sets the value for MobileNumber to be an explicit nil

### UnsetMobileNumber
`func (o *CustomerObject) UnsetMobileNumber()`

UnsetMobileNumber ensures that no value is present for MobileNumber, not even an explicit nil
### GetCustomerId

`func (o *CustomerObject) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomerObject) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CustomerObject) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CustomerObject) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *CustomerObject) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *CustomerObject) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetAddresses

`func (o *CustomerObject) GetAddresses() []AddressObject`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *CustomerObject) GetAddressesOk() (*[]AddressObject, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *CustomerObject) SetAddresses(v []AddressObject)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *CustomerObject) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### SetAddressesNil

`func (o *CustomerObject) SetAddressesNil(b bool)`

 SetAddressesNil sets the value for Addresses to be an explicit nil

### UnsetAddresses
`func (o *CustomerObject) UnsetAddresses()`

UnsetAddresses ensures that no value is present for Addresses, not even an explicit nil

[[Back to README]](../../README.md)


