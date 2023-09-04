# CardParametersCardInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardNumber** | **string** |  | 
**ExpiryMonth** | **string** | Card expiry month in MM format | 
**ExpiryYear** | **string** | Card expiry month in YY format | 
**CardholderName** | Pointer to **NullableString** | Cardholder name | [optional] 
**Cvv** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCardParametersCardInformation

`func NewCardParametersCardInformation(cardNumber string, expiryMonth string, expiryYear string, ) *CardParametersCardInformation`

NewCardParametersCardInformation instantiates a new CardParametersCardInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardParametersCardInformationWithDefaults

`func NewCardParametersCardInformationWithDefaults() *CardParametersCardInformation`

NewCardParametersCardInformationWithDefaults instantiates a new CardParametersCardInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardNumber

`func (o *CardParametersCardInformation) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *CardParametersCardInformation) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *CardParametersCardInformation) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.


### GetExpiryMonth

`func (o *CardParametersCardInformation) GetExpiryMonth() string`

GetExpiryMonth returns the ExpiryMonth field if non-nil, zero value otherwise.

### GetExpiryMonthOk

`func (o *CardParametersCardInformation) GetExpiryMonthOk() (*string, bool)`

GetExpiryMonthOk returns a tuple with the ExpiryMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryMonth

`func (o *CardParametersCardInformation) SetExpiryMonth(v string)`

SetExpiryMonth sets ExpiryMonth field to given value.


### GetExpiryYear

`func (o *CardParametersCardInformation) GetExpiryYear() string`

GetExpiryYear returns the ExpiryYear field if non-nil, zero value otherwise.

### GetExpiryYearOk

`func (o *CardParametersCardInformation) GetExpiryYearOk() (*string, bool)`

GetExpiryYearOk returns a tuple with the ExpiryYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryYear

`func (o *CardParametersCardInformation) SetExpiryYear(v string)`

SetExpiryYear sets ExpiryYear field to given value.


### GetCardholderName

`func (o *CardParametersCardInformation) GetCardholderName() string`

GetCardholderName returns the CardholderName field if non-nil, zero value otherwise.

### GetCardholderNameOk

`func (o *CardParametersCardInformation) GetCardholderNameOk() (*string, bool)`

GetCardholderNameOk returns a tuple with the CardholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderName

`func (o *CardParametersCardInformation) SetCardholderName(v string)`

SetCardholderName sets CardholderName field to given value.

### HasCardholderName

`func (o *CardParametersCardInformation) HasCardholderName() bool`

HasCardholderName returns a boolean if a field has been set.

### SetCardholderNameNil

`func (o *CardParametersCardInformation) SetCardholderNameNil(b bool)`

 SetCardholderNameNil sets the value for CardholderName to be an explicit nil

### UnsetCardholderName
`func (o *CardParametersCardInformation) UnsetCardholderName()`

UnsetCardholderName ensures that no value is present for CardholderName, not even an explicit nil
### GetCvv

`func (o *CardParametersCardInformation) GetCvv() string`

GetCvv returns the Cvv field if non-nil, zero value otherwise.

### GetCvvOk

`func (o *CardParametersCardInformation) GetCvvOk() (*string, bool)`

GetCvvOk returns a tuple with the Cvv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvv

`func (o *CardParametersCardInformation) SetCvv(v string)`

SetCvv sets Cvv field to given value.

### HasCvv

`func (o *CardParametersCardInformation) HasCvv() bool`

HasCvv returns a boolean if a field has been set.

### SetCvvNil

`func (o *CardParametersCardInformation) SetCvvNil(b bool)`

 SetCvvNil sets the value for Cvv to be an explicit nil

### UnsetCvv
`func (o *CardParametersCardInformation) UnsetCvv()`

UnsetCvv ensures that no value is present for Cvv, not even an explicit nil

[[Back to README]](../../README.md)


