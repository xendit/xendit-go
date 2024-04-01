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

// checks if the IndividualDetail type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &IndividualDetail{}

// IndividualDetail struct for IndividualDetail
type IndividualDetail struct {
	GivenNames *string `json:"given_names,omitempty"`
	GivenNamesNonRoman NullableString `json:"given_names_non_roman,omitempty"`
	MiddleName NullableString `json:"middle_name,omitempty"`
	Surname NullableString `json:"surname,omitempty"`
	SurnameNonRoman NullableString `json:"surname_non_roman,omitempty"`
	MotherMaidenName NullableString `json:"mother_maiden_name,omitempty"`
	Gender NullableString `json:"gender,omitempty"`
	DateOfBirth NullableString `json:"date_of_birth,omitempty"`
	// ISO3166-2 country code
	Nationality NullableString `json:"nationality,omitempty"`
	PlaceOfBirth NullableString `json:"place_of_birth,omitempty"`
	Employment NullableEmploymentDetail `json:"employment,omitempty"`
}

// NewIndividualDetail instantiates a new IndividualDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndividualDetail() *IndividualDetail {
	this := IndividualDetail{}
	return &this
}

// NewIndividualDetailWithDefaults instantiates a new IndividualDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndividualDetailWithDefaults() *IndividualDetail {
	this := IndividualDetail{}
	return &this
}

// GetGivenNames returns the GivenNames field value if set, zero value otherwise.
func (o *IndividualDetail) GetGivenNames() string {
	if o == nil || utils.IsNil(o.GivenNames) {
		var ret string
		return ret
	}
	return *o.GivenNames
}

// GetGivenNamesOk returns a tuple with the GivenNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDetail) GetGivenNamesOk() (*string, bool) {
	if o == nil || utils.IsNil(o.GivenNames) {
		return nil, false
	}
	return o.GivenNames, true
}

// HasGivenNames returns a boolean if a field has been set.
func (o *IndividualDetail) HasGivenNames() bool {
	if o != nil && !utils.IsNil(o.GivenNames) {
		return true
	}

	return false
}

// SetGivenNames gets a reference to the given string and assigns it to the GivenNames field.
func (o *IndividualDetail) SetGivenNames(v string) {
	o.GivenNames = &v
}

// GetGivenNamesNonRoman returns the GivenNamesNonRoman field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndividualDetail) GetGivenNamesNonRoman() string {
	if o == nil || utils.IsNil(o.GivenNamesNonRoman.Get()) {
		var ret string
		return ret
	}
	return *o.GivenNamesNonRoman.Get()
}

// GetGivenNamesNonRomanOk returns a tuple with the GivenNamesNonRoman field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndividualDetail) GetGivenNamesNonRomanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GivenNamesNonRoman.Get(), o.GivenNamesNonRoman.IsSet()
}

// HasGivenNamesNonRoman returns a boolean if a field has been set.
func (o *IndividualDetail) HasGivenNamesNonRoman() bool {
	if o != nil && o.GivenNamesNonRoman.IsSet() {
		return true
	}

	return false
}

// SetGivenNamesNonRoman gets a reference to the given NullableString and assigns it to the GivenNamesNonRoman field.
func (o *IndividualDetail) SetGivenNamesNonRoman(v string) {
	o.GivenNamesNonRoman.Set(&v)
}
// SetGivenNamesNonRomanNil sets the value for GivenNamesNonRoman to be an explicit nil
func (o *IndividualDetail) SetGivenNamesNonRomanNil() {
	o.GivenNamesNonRoman.Set(nil)
}

// UnsetGivenNamesNonRoman ensures that no value is present for GivenNamesNonRoman, not even an explicit nil
func (o *IndividualDetail) UnsetGivenNamesNonRoman() {
	o.GivenNamesNonRoman.Unset()
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndividualDetail) GetMiddleName() string {
	if o == nil || utils.IsNil(o.MiddleName.Get()) {
		var ret string
		return ret
	}
	return *o.MiddleName.Get()
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndividualDetail) GetMiddleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MiddleName.Get(), o.MiddleName.IsSet()
}

// HasMiddleName returns a boolean if a field has been set.
func (o *IndividualDetail) HasMiddleName() bool {
	if o != nil && o.MiddleName.IsSet() {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given NullableString and assigns it to the MiddleName field.
func (o *IndividualDetail) SetMiddleName(v string) {
	o.MiddleName.Set(&v)
}
// SetMiddleNameNil sets the value for MiddleName to be an explicit nil
func (o *IndividualDetail) SetMiddleNameNil() {
	o.MiddleName.Set(nil)
}

// UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
func (o *IndividualDetail) UnsetMiddleName() {
	o.MiddleName.Unset()
}

// GetSurname returns the Surname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndividualDetail) GetSurname() string {
	if o == nil || utils.IsNil(o.Surname.Get()) {
		var ret string
		return ret
	}
	return *o.Surname.Get()
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndividualDetail) GetSurnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Surname.Get(), o.Surname.IsSet()
}

// HasSurname returns a boolean if a field has been set.
func (o *IndividualDetail) HasSurname() bool {
	if o != nil && o.Surname.IsSet() {
		return true
	}

	return false
}

// SetSurname gets a reference to the given NullableString and assigns it to the Surname field.
func (o *IndividualDetail) SetSurname(v string) {
	o.Surname.Set(&v)
}
// SetSurnameNil sets the value for Surname to be an explicit nil
func (o *IndividualDetail) SetSurnameNil() {
	o.Surname.Set(nil)
}

// UnsetSurname ensures that no value is present for Surname, not even an explicit nil
func (o *IndividualDetail) UnsetSurname() {
	o.Surname.Unset()
}

// GetSurnameNonRoman returns the SurnameNonRoman field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndividualDetail) GetSurnameNonRoman() string {
	if o == nil || utils.IsNil(o.SurnameNonRoman.Get()) {
		var ret string
		return ret
	}
	return *o.SurnameNonRoman.Get()
}

// GetSurnameNonRomanOk returns a tuple with the SurnameNonRoman field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndividualDetail) GetSurnameNonRomanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SurnameNonRoman.Get(), o.SurnameNonRoman.IsSet()
}

// HasSurnameNonRoman returns a boolean if a field has been set.
func (o *IndividualDetail) HasSurnameNonRoman() bool {
	if o != nil && o.SurnameNonRoman.IsSet() {
		return true
	}

	return false
}

// SetSurnameNonRoman gets a reference to the given NullableString and assigns it to the SurnameNonRoman field.
func (o *IndividualDetail) SetSurnameNonRoman(v string) {
	o.SurnameNonRoman.Set(&v)
}
// SetSurnameNonRomanNil sets the value for SurnameNonRoman to be an explicit nil
func (o *IndividualDetail) SetSurnameNonRomanNil() {
	o.SurnameNonRoman.Set(nil)
}

// UnsetSurnameNonRoman ensures that no value is present for SurnameNonRoman, not even an explicit nil
func (o *IndividualDetail) UnsetSurnameNonRoman() {
	o.SurnameNonRoman.Unset()
}

// GetMotherMaidenName returns the MotherMaidenName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndividualDetail) GetMotherMaidenName() string {
	if o == nil || utils.IsNil(o.MotherMaidenName.Get()) {
		var ret string
		return ret
	}
	return *o.MotherMaidenName.Get()
}

// GetMotherMaidenNameOk returns a tuple with the MotherMaidenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndividualDetail) GetMotherMaidenNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MotherMaidenName.Get(), o.MotherMaidenName.IsSet()
}

// HasMotherMaidenName returns a boolean if a field has been set.
func (o *IndividualDetail) HasMotherMaidenName() bool {
	if o != nil && o.MotherMaidenName.IsSet() {
		return true
	}

	return false
}

// SetMotherMaidenName gets a reference to the given NullableString and assigns it to the MotherMaidenName field.
func (o *IndividualDetail) SetMotherMaidenName(v string) {
	o.MotherMaidenName.Set(&v)
}
// SetMotherMaidenNameNil sets the value for MotherMaidenName to be an explicit nil
func (o *IndividualDetail) SetMotherMaidenNameNil() {
	o.MotherMaidenName.Set(nil)
}

// UnsetMotherMaidenName ensures that no value is present for MotherMaidenName, not even an explicit nil
func (o *IndividualDetail) UnsetMotherMaidenName() {
	o.MotherMaidenName.Unset()
}

// GetGender returns the Gender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndividualDetail) GetGender() string {
	if o == nil || utils.IsNil(o.Gender.Get()) {
		var ret string
		return ret
	}
	return *o.Gender.Get()
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndividualDetail) GetGenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gender.Get(), o.Gender.IsSet()
}

// HasGender returns a boolean if a field has been set.
func (o *IndividualDetail) HasGender() bool {
	if o != nil && o.Gender.IsSet() {
		return true
	}

	return false
}

// SetGender gets a reference to the given NullableString and assigns it to the Gender field.
func (o *IndividualDetail) SetGender(v string) {
	o.Gender.Set(&v)
}
// SetGenderNil sets the value for Gender to be an explicit nil
func (o *IndividualDetail) SetGenderNil() {
	o.Gender.Set(nil)
}

// UnsetGender ensures that no value is present for Gender, not even an explicit nil
func (o *IndividualDetail) UnsetGender() {
	o.Gender.Unset()
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndividualDetail) GetDateOfBirth() string {
	if o == nil || utils.IsNil(o.DateOfBirth.Get()) {
		var ret string
		return ret
	}
	return *o.DateOfBirth.Get()
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndividualDetail) GetDateOfBirthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateOfBirth.Get(), o.DateOfBirth.IsSet()
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *IndividualDetail) HasDateOfBirth() bool {
	if o != nil && o.DateOfBirth.IsSet() {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given NullableString and assigns it to the DateOfBirth field.
func (o *IndividualDetail) SetDateOfBirth(v string) {
	o.DateOfBirth.Set(&v)
}
// SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil
func (o *IndividualDetail) SetDateOfBirthNil() {
	o.DateOfBirth.Set(nil)
}

// UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
func (o *IndividualDetail) UnsetDateOfBirth() {
	o.DateOfBirth.Unset()
}

// GetNationality returns the Nationality field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndividualDetail) GetNationality() string {
	if o == nil || utils.IsNil(o.Nationality.Get()) {
		var ret string
		return ret
	}
	return *o.Nationality.Get()
}

// GetNationalityOk returns a tuple with the Nationality field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndividualDetail) GetNationalityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nationality.Get(), o.Nationality.IsSet()
}

// HasNationality returns a boolean if a field has been set.
func (o *IndividualDetail) HasNationality() bool {
	if o != nil && o.Nationality.IsSet() {
		return true
	}

	return false
}

// SetNationality gets a reference to the given NullableString and assigns it to the Nationality field.
func (o *IndividualDetail) SetNationality(v string) {
	o.Nationality.Set(&v)
}
// SetNationalityNil sets the value for Nationality to be an explicit nil
func (o *IndividualDetail) SetNationalityNil() {
	o.Nationality.Set(nil)
}

// UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
func (o *IndividualDetail) UnsetNationality() {
	o.Nationality.Unset()
}

// GetPlaceOfBirth returns the PlaceOfBirth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndividualDetail) GetPlaceOfBirth() string {
	if o == nil || utils.IsNil(o.PlaceOfBirth.Get()) {
		var ret string
		return ret
	}
	return *o.PlaceOfBirth.Get()
}

// GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndividualDetail) GetPlaceOfBirthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlaceOfBirth.Get(), o.PlaceOfBirth.IsSet()
}

// HasPlaceOfBirth returns a boolean if a field has been set.
func (o *IndividualDetail) HasPlaceOfBirth() bool {
	if o != nil && o.PlaceOfBirth.IsSet() {
		return true
	}

	return false
}

// SetPlaceOfBirth gets a reference to the given NullableString and assigns it to the PlaceOfBirth field.
func (o *IndividualDetail) SetPlaceOfBirth(v string) {
	o.PlaceOfBirth.Set(&v)
}
// SetPlaceOfBirthNil sets the value for PlaceOfBirth to be an explicit nil
func (o *IndividualDetail) SetPlaceOfBirthNil() {
	o.PlaceOfBirth.Set(nil)
}

// UnsetPlaceOfBirth ensures that no value is present for PlaceOfBirth, not even an explicit nil
func (o *IndividualDetail) UnsetPlaceOfBirth() {
	o.PlaceOfBirth.Unset()
}

// GetEmployment returns the Employment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndividualDetail) GetEmployment() EmploymentDetail {
	if o == nil || utils.IsNil(o.Employment.Get()) {
		var ret EmploymentDetail
		return ret
	}
	return *o.Employment.Get()
}

// GetEmploymentOk returns a tuple with the Employment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndividualDetail) GetEmploymentOk() (*EmploymentDetail, bool) {
	if o == nil {
		return nil, false
	}
	return o.Employment.Get(), o.Employment.IsSet()
}

// HasEmployment returns a boolean if a field has been set.
func (o *IndividualDetail) HasEmployment() bool {
	if o != nil && o.Employment.IsSet() {
		return true
	}

	return false
}

// SetEmployment gets a reference to the given NullableEmploymentDetail and assigns it to the Employment field.
func (o *IndividualDetail) SetEmployment(v EmploymentDetail) {
	o.Employment.Set(&v)
}
// SetEmploymentNil sets the value for Employment to be an explicit nil
func (o *IndividualDetail) SetEmploymentNil() {
	o.Employment.Set(nil)
}

// UnsetEmployment ensures that no value is present for Employment, not even an explicit nil
func (o *IndividualDetail) UnsetEmployment() {
	o.Employment.Unset()
}

func (o IndividualDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndividualDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.GivenNames) {
		toSerialize["given_names"] = o.GivenNames
	}
	if o.GivenNamesNonRoman.IsSet() {
		toSerialize["given_names_non_roman"] = o.GivenNamesNonRoman.Get()
    }
	if o.MiddleName.IsSet() {
		toSerialize["middle_name"] = o.MiddleName.Get()
    }
	if o.Surname.IsSet() {
		toSerialize["surname"] = o.Surname.Get()
    }
	if o.SurnameNonRoman.IsSet() {
		toSerialize["surname_non_roman"] = o.SurnameNonRoman.Get()
    }
	if o.MotherMaidenName.IsSet() {
		toSerialize["mother_maiden_name"] = o.MotherMaidenName.Get()
    }
	if o.Gender.IsSet() {
		toSerialize["gender"] = o.Gender.Get()
        if o.Gender.Get() != nil && (*o.Gender.Get() != "MALE" && *o.Gender.Get() != "FEMALE" && *o.Gender.Get() != "OTHER") {
            toSerialize["gender"] = nil
            return toSerialize, utils.NewError("invalid value for Gender when marshalling to JSON, must be one of MALE, FEMALE, OTHER")
        }
    }
	if o.DateOfBirth.IsSet() {
		toSerialize["date_of_birth"] = o.DateOfBirth.Get()
    }
	if o.Nationality.IsSet() {
		toSerialize["nationality"] = o.Nationality.Get()
    }
	if o.PlaceOfBirth.IsSet() {
		toSerialize["place_of_birth"] = o.PlaceOfBirth.Get()
    }
	if o.Employment.IsSet() {
		toSerialize["employment"] = o.Employment.Get()
    }
	return toSerialize, nil
}

type NullableIndividualDetail struct {
	value *IndividualDetail
	isSet bool
}

func (v NullableIndividualDetail) Get() *IndividualDetail {
	return v.value
}

func (v *NullableIndividualDetail) Set(val *IndividualDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableIndividualDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableIndividualDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndividualDetail(val *IndividualDetail) *NullableIndividualDetail {
	return &NullableIndividualDetail{value: val, isSet: true}
}

func (v NullableIndividualDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndividualDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


