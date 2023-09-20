# PaymentRequest

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **Id** | **string** |  |  |
| **Created** | **string** |  |  |
| **Updated** | **string** |  |  |
| **ReferenceId** | **string** |  |  |
| **BusinessId** | **string** |  |  |
| **CustomerId** | Pointer to **NullableString** |  | [optional]  |
| **Customer** | Pointer to **map[string]interface{}** |  | [optional]  |
| **Amount** | Pointer to **float64** |  | [optional]  |
| **MinAmount** | Pointer to **NullableFloat64** |  | [optional]  |
| **MaxAmount** | Pointer to **NullableFloat64** |  | [optional]  |
| **Country** | Pointer to [**PaymentRequestCountry**](PaymentRequestCountry.md) |  | [optional]  |
| **Currency** | [**PaymentRequestCurrency**](PaymentRequestCurrency.md) |  |  |
| **PaymentMethod** | [**PaymentMethod**](PaymentMethod.md) |  |  |
| **Description** | Pointer to **NullableString** |  | [optional]  |
| **FailureCode** | Pointer to **NullableString** |  | [optional]  |
| **CaptureMethod** | Pointer to [**NullablePaymentRequestCaptureMethod**](PaymentRequestCaptureMethod.md) |  | [optional]  |
| **Initiator** | Pointer to [**NullablePaymentRequestInitiator**](PaymentRequestInitiator.md) |  | [optional]  |
| **CardVerificationResults** | Pointer to [**NullablePaymentRequestCardVerificationResults**](PaymentRequestCardVerificationResults.md) |  | [optional]  |
| **Status** | [**PaymentRequestStatus**](PaymentRequestStatus.md) |  |  |
| **Actions** | Pointer to [**PaymentRequestAction[]**](PaymentRequestAction.md) |  | [optional]  |
| **Metadata** | Pointer to **map[string]interface{}** |  | [optional]  |
| **ShippingInformation** | Pointer to [**NullablePaymentRequestShippingInformation**](PaymentRequestShippingInformation.md) |  | [optional]  |
| **Items** | Pointer to [**PaymentRequestBasketItem[]**](PaymentRequestBasketItem.md) |  | [optional]  |

## Methods

### NewPaymentRequest

`func NewPaymentRequest(id string, created string, updated string, referenceId string, businessId string, currency PaymentRequestCurrency, paymentMethod PaymentMethod, status PaymentRequestStatus, ) *PaymentRequest`

NewPaymentRequest instantiates a new PaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestWithDefaults

`func NewPaymentRequestWithDefaults() *PaymentRequest`

NewPaymentRequestWithDefaults instantiates a new PaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetCreated

`func (o *PaymentRequest) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PaymentRequest) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PaymentRequest) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *PaymentRequest) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PaymentRequest) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PaymentRequest) SetUpdated(v string)`

SetUpdated sets Updated field to given value.


### GetReferenceId

`func (o *PaymentRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *PaymentRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *PaymentRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetBusinessId

`func (o *PaymentRequest) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *PaymentRequest) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *PaymentRequest) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.


### GetCustomerId

`func (o *PaymentRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PaymentRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PaymentRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *PaymentRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *PaymentRequest) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *PaymentRequest) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetCustomer

`func (o *PaymentRequest) GetCustomer() map[string]interface{}`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *PaymentRequest) GetCustomerOk() (*map[string]interface{}, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *PaymentRequest) SetCustomer(v map[string]interface{})`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *PaymentRequest) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### SetCustomerNil

`func (o *PaymentRequest) SetCustomerNil(b bool)`

 SetCustomerNil sets the value for Customer to be an explicit nil

### UnsetCustomer
`func (o *PaymentRequest) UnsetCustomer()`

UnsetCustomer ensures that no value is present for Customer, not even an explicit nil
### GetAmount

`func (o *PaymentRequest) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentRequest) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentRequest) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetMinAmount

`func (o *PaymentRequest) GetMinAmount() float64`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *PaymentRequest) GetMinAmountOk() (*float64, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *PaymentRequest) SetMinAmount(v float64)`

SetMinAmount sets MinAmount field to given value.

### HasMinAmount

`func (o *PaymentRequest) HasMinAmount() bool`

HasMinAmount returns a boolean if a field has been set.

### SetMinAmountNil

`func (o *PaymentRequest) SetMinAmountNil(b bool)`

 SetMinAmountNil sets the value for MinAmount to be an explicit nil

### UnsetMinAmount
`func (o *PaymentRequest) UnsetMinAmount()`

UnsetMinAmount ensures that no value is present for MinAmount, not even an explicit nil
### GetMaxAmount

`func (o *PaymentRequest) GetMaxAmount() float64`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *PaymentRequest) GetMaxAmountOk() (*float64, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *PaymentRequest) SetMaxAmount(v float64)`

SetMaxAmount sets MaxAmount field to given value.

### HasMaxAmount

`func (o *PaymentRequest) HasMaxAmount() bool`

HasMaxAmount returns a boolean if a field has been set.

### SetMaxAmountNil

`func (o *PaymentRequest) SetMaxAmountNil(b bool)`

 SetMaxAmountNil sets the value for MaxAmount to be an explicit nil

### UnsetMaxAmount
`func (o *PaymentRequest) UnsetMaxAmount()`

UnsetMaxAmount ensures that no value is present for MaxAmount, not even an explicit nil
### GetCountry

`func (o *PaymentRequest) GetCountry() PaymentRequestCountry`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PaymentRequest) GetCountryOk() (*PaymentRequestCountry, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PaymentRequest) SetCountry(v PaymentRequestCountry)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PaymentRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentRequest) GetCurrency() PaymentRequestCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentRequest) GetCurrencyOk() (*PaymentRequestCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentRequest) SetCurrency(v PaymentRequestCurrency)`

SetCurrency sets Currency field to given value.


### GetPaymentMethod

`func (o *PaymentRequest) GetPaymentMethod() PaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PaymentRequest) GetPaymentMethodOk() (*PaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PaymentRequest) SetPaymentMethod(v PaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetDescription

`func (o *PaymentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PaymentRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PaymentRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFailureCode

`func (o *PaymentRequest) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *PaymentRequest) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *PaymentRequest) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *PaymentRequest) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### SetFailureCodeNil

`func (o *PaymentRequest) SetFailureCodeNil(b bool)`

 SetFailureCodeNil sets the value for FailureCode to be an explicit nil

### UnsetFailureCode
`func (o *PaymentRequest) UnsetFailureCode()`

UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil
### GetCaptureMethod

`func (o *PaymentRequest) GetCaptureMethod() PaymentRequestCaptureMethod`

GetCaptureMethod returns the CaptureMethod field if non-nil, zero value otherwise.

### GetCaptureMethodOk

`func (o *PaymentRequest) GetCaptureMethodOk() (*PaymentRequestCaptureMethod, bool)`

GetCaptureMethodOk returns a tuple with the CaptureMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureMethod

`func (o *PaymentRequest) SetCaptureMethod(v PaymentRequestCaptureMethod)`

SetCaptureMethod sets CaptureMethod field to given value.

### HasCaptureMethod

`func (o *PaymentRequest) HasCaptureMethod() bool`

HasCaptureMethod returns a boolean if a field has been set.

### SetCaptureMethodNil

`func (o *PaymentRequest) SetCaptureMethodNil(b bool)`

 SetCaptureMethodNil sets the value for CaptureMethod to be an explicit nil

### UnsetCaptureMethod
`func (o *PaymentRequest) UnsetCaptureMethod()`

UnsetCaptureMethod ensures that no value is present for CaptureMethod, not even an explicit nil
### GetInitiator

`func (o *PaymentRequest) GetInitiator() PaymentRequestInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *PaymentRequest) GetInitiatorOk() (*PaymentRequestInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *PaymentRequest) SetInitiator(v PaymentRequestInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *PaymentRequest) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### SetInitiatorNil

`func (o *PaymentRequest) SetInitiatorNil(b bool)`

 SetInitiatorNil sets the value for Initiator to be an explicit nil

### UnsetInitiator
`func (o *PaymentRequest) UnsetInitiator()`

UnsetInitiator ensures that no value is present for Initiator, not even an explicit nil
### GetCardVerificationResults

`func (o *PaymentRequest) GetCardVerificationResults() PaymentRequestCardVerificationResults`

GetCardVerificationResults returns the CardVerificationResults field if non-nil, zero value otherwise.

### GetCardVerificationResultsOk

`func (o *PaymentRequest) GetCardVerificationResultsOk() (*PaymentRequestCardVerificationResults, bool)`

GetCardVerificationResultsOk returns a tuple with the CardVerificationResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardVerificationResults

`func (o *PaymentRequest) SetCardVerificationResults(v PaymentRequestCardVerificationResults)`

SetCardVerificationResults sets CardVerificationResults field to given value.

### HasCardVerificationResults

`func (o *PaymentRequest) HasCardVerificationResults() bool`

HasCardVerificationResults returns a boolean if a field has been set.

### SetCardVerificationResultsNil

`func (o *PaymentRequest) SetCardVerificationResultsNil(b bool)`

 SetCardVerificationResultsNil sets the value for CardVerificationResults to be an explicit nil

### UnsetCardVerificationResults
`func (o *PaymentRequest) UnsetCardVerificationResults()`

UnsetCardVerificationResults ensures that no value is present for CardVerificationResults, not even an explicit nil
### GetStatus

`func (o *PaymentRequest) GetStatus() PaymentRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentRequest) GetStatusOk() (*PaymentRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentRequest) SetStatus(v PaymentRequestStatus)`

SetStatus sets Status field to given value.


### GetActions

`func (o *PaymentRequest) GetActions() []PaymentRequestAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PaymentRequest) GetActionsOk() (*[]PaymentRequestAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PaymentRequest) SetActions(v []PaymentRequestAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *PaymentRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetMetadata

`func (o *PaymentRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaymentRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaymentRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PaymentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PaymentRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PaymentRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetShippingInformation

`func (o *PaymentRequest) GetShippingInformation() PaymentRequestShippingInformation`

GetShippingInformation returns the ShippingInformation field if non-nil, zero value otherwise.

### GetShippingInformationOk

`func (o *PaymentRequest) GetShippingInformationOk() (*PaymentRequestShippingInformation, bool)`

GetShippingInformationOk returns a tuple with the ShippingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingInformation

`func (o *PaymentRequest) SetShippingInformation(v PaymentRequestShippingInformation)`

SetShippingInformation sets ShippingInformation field to given value.

### HasShippingInformation

`func (o *PaymentRequest) HasShippingInformation() bool`

HasShippingInformation returns a boolean if a field has been set.

### SetShippingInformationNil

`func (o *PaymentRequest) SetShippingInformationNil(b bool)`

 SetShippingInformationNil sets the value for ShippingInformation to be an explicit nil

### UnsetShippingInformation
`func (o *PaymentRequest) UnsetShippingInformation()`

UnsetShippingInformation ensures that no value is present for ShippingInformation, not even an explicit nil
### GetItems

`func (o *PaymentRequest) GetItems() []PaymentRequestBasketItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PaymentRequest) GetItemsOk() (*[]PaymentRequestBasketItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PaymentRequest) SetItems(v []PaymentRequestBasketItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *PaymentRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *PaymentRequest) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *PaymentRequest) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to README]](../../README.md)


