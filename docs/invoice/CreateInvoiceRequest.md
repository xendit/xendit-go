# CreateInvoiceRequest
An object representing for an invoice creation request.

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **ExternalId** | **string** | ☑️ | The external ID of the invoice. |  |
| **Amount** | **float64** | ☑️ | The invoice amount. |  |
| **PayerEmail** | Pointer to **string** |  | The email address of the payer. |  |
| **Description** | Pointer to **string** |  | A description of the payment. |  |
| **InvoiceDuration** | Pointer to **string** |  | The duration of the invoice. |  |
| **CallbackVirtualAccountId** | Pointer to **string** |  | The ID of the callback virtual account. |  |
| **ShouldSendEmail** | Pointer to **bool** |  | Indicates whether email notifications should be sent. |  |
| **Customer** | Pointer to [**CustomerObject**](CustomerObject.md) |  |  |  |
| **CustomerNotificationPreference** | Pointer to [**NotificationPreference**](NotificationPreference.md) |  |  |  |
| **SuccessRedirectUrl** | Pointer to **string** |  | The URL to redirect to on successful payment. |  |
| **FailureRedirectUrl** | Pointer to **string** |  | The URL to redirect to on payment failure. |  |
| **PaymentMethods** | Pointer to **string[]** |  | An array of available payment methods. |  |
| **MidLabel** | Pointer to **string** |  | The middle label. |  |
| **ShouldAuthenticateCreditCard** | Pointer to **bool** |  | Indicates whether credit card authentication is required. |  |
| **Currency** | Pointer to **string** |  | The currency of the invoice. |  |
| **ReminderTime** | Pointer to **float32** |  | The reminder time. |  |
| **Locale** | Pointer to **string** |  | The default language to display. |  |
| **ReminderTimeUnit** | Pointer to **string** |  | The unit of the reminder time. |  |
| **Items** | Pointer to [**InvoiceItem[]**](InvoiceItem.md) |  | An array of items included in the invoice. |  |
| **Fees** | Pointer to [**InvoiceFee[]**](InvoiceFee.md) |  | An array of fees associated with the invoice. |  |
| **ChannelProperties** | Pointer to [**ChannelProperties**](ChannelProperties.md) |  |  |  |
| **Metadata** | Pointer to **map[string]interface{}** |  | A free-format JSON for additional information that you may use. Object can be up to 50 keys, with key names up to 40 characters long and values up to 500 characters long. |  |

## Methods

### NewCreateInvoiceRequest

`func NewCreateInvoiceRequest(externalId string, amount float64, ) *CreateInvoiceRequest`

NewCreateInvoiceRequest instantiates a new CreateInvoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInvoiceRequestWithDefaults

`func NewCreateInvoiceRequestWithDefaults() *CreateInvoiceRequest`

NewCreateInvoiceRequestWithDefaults instantiates a new CreateInvoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *CreateInvoiceRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateInvoiceRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateInvoiceRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetAmount

`func (o *CreateInvoiceRequest) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateInvoiceRequest) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateInvoiceRequest) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetPayerEmail

`func (o *CreateInvoiceRequest) GetPayerEmail() string`

GetPayerEmail returns the PayerEmail field if non-nil, zero value otherwise.

### GetPayerEmailOk

`func (o *CreateInvoiceRequest) GetPayerEmailOk() (*string, bool)`

GetPayerEmailOk returns a tuple with the PayerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerEmail

`func (o *CreateInvoiceRequest) SetPayerEmail(v string)`

SetPayerEmail sets PayerEmail field to given value.

### HasPayerEmail

`func (o *CreateInvoiceRequest) HasPayerEmail() bool`

HasPayerEmail returns a boolean if a field has been set.

### GetDescription

`func (o *CreateInvoiceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateInvoiceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateInvoiceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateInvoiceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInvoiceDuration

`func (o *CreateInvoiceRequest) GetInvoiceDuration() string`

GetInvoiceDuration returns the InvoiceDuration field if non-nil, zero value otherwise.

### GetInvoiceDurationOk

`func (o *CreateInvoiceRequest) GetInvoiceDurationOk() (*string, bool)`

GetInvoiceDurationOk returns a tuple with the InvoiceDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDuration

`func (o *CreateInvoiceRequest) SetInvoiceDuration(v string)`

SetInvoiceDuration sets InvoiceDuration field to given value.

### HasInvoiceDuration

`func (o *CreateInvoiceRequest) HasInvoiceDuration() bool`

HasInvoiceDuration returns a boolean if a field has been set.

### GetCallbackVirtualAccountId

`func (o *CreateInvoiceRequest) GetCallbackVirtualAccountId() string`

GetCallbackVirtualAccountId returns the CallbackVirtualAccountId field if non-nil, zero value otherwise.

### GetCallbackVirtualAccountIdOk

`func (o *CreateInvoiceRequest) GetCallbackVirtualAccountIdOk() (*string, bool)`

GetCallbackVirtualAccountIdOk returns a tuple with the CallbackVirtualAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackVirtualAccountId

`func (o *CreateInvoiceRequest) SetCallbackVirtualAccountId(v string)`

SetCallbackVirtualAccountId sets CallbackVirtualAccountId field to given value.

### HasCallbackVirtualAccountId

`func (o *CreateInvoiceRequest) HasCallbackVirtualAccountId() bool`

HasCallbackVirtualAccountId returns a boolean if a field has been set.

### GetShouldSendEmail

`func (o *CreateInvoiceRequest) GetShouldSendEmail() bool`

GetShouldSendEmail returns the ShouldSendEmail field if non-nil, zero value otherwise.

### GetShouldSendEmailOk

`func (o *CreateInvoiceRequest) GetShouldSendEmailOk() (*bool, bool)`

GetShouldSendEmailOk returns a tuple with the ShouldSendEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldSendEmail

`func (o *CreateInvoiceRequest) SetShouldSendEmail(v bool)`

SetShouldSendEmail sets ShouldSendEmail field to given value.

### HasShouldSendEmail

`func (o *CreateInvoiceRequest) HasShouldSendEmail() bool`

HasShouldSendEmail returns a boolean if a field has been set.

### GetCustomer

`func (o *CreateInvoiceRequest) GetCustomer() CustomerObject`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CreateInvoiceRequest) GetCustomerOk() (*CustomerObject, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CreateInvoiceRequest) SetCustomer(v CustomerObject)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *CreateInvoiceRequest) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetCustomerNotificationPreference

`func (o *CreateInvoiceRequest) GetCustomerNotificationPreference() NotificationPreference`

GetCustomerNotificationPreference returns the CustomerNotificationPreference field if non-nil, zero value otherwise.

### GetCustomerNotificationPreferenceOk

`func (o *CreateInvoiceRequest) GetCustomerNotificationPreferenceOk() (*NotificationPreference, bool)`

GetCustomerNotificationPreferenceOk returns a tuple with the CustomerNotificationPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNotificationPreference

`func (o *CreateInvoiceRequest) SetCustomerNotificationPreference(v NotificationPreference)`

SetCustomerNotificationPreference sets CustomerNotificationPreference field to given value.

### HasCustomerNotificationPreference

`func (o *CreateInvoiceRequest) HasCustomerNotificationPreference() bool`

HasCustomerNotificationPreference returns a boolean if a field has been set.

### GetSuccessRedirectUrl

`func (o *CreateInvoiceRequest) GetSuccessRedirectUrl() string`

GetSuccessRedirectUrl returns the SuccessRedirectUrl field if non-nil, zero value otherwise.

### GetSuccessRedirectUrlOk

`func (o *CreateInvoiceRequest) GetSuccessRedirectUrlOk() (*string, bool)`

GetSuccessRedirectUrlOk returns a tuple with the SuccessRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRedirectUrl

`func (o *CreateInvoiceRequest) SetSuccessRedirectUrl(v string)`

SetSuccessRedirectUrl sets SuccessRedirectUrl field to given value.

### HasSuccessRedirectUrl

`func (o *CreateInvoiceRequest) HasSuccessRedirectUrl() bool`

HasSuccessRedirectUrl returns a boolean if a field has been set.

### GetFailureRedirectUrl

`func (o *CreateInvoiceRequest) GetFailureRedirectUrl() string`

GetFailureRedirectUrl returns the FailureRedirectUrl field if non-nil, zero value otherwise.

### GetFailureRedirectUrlOk

`func (o *CreateInvoiceRequest) GetFailureRedirectUrlOk() (*string, bool)`

GetFailureRedirectUrlOk returns a tuple with the FailureRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureRedirectUrl

`func (o *CreateInvoiceRequest) SetFailureRedirectUrl(v string)`

SetFailureRedirectUrl sets FailureRedirectUrl field to given value.

### HasFailureRedirectUrl

`func (o *CreateInvoiceRequest) HasFailureRedirectUrl() bool`

HasFailureRedirectUrl returns a boolean if a field has been set.

### GetPaymentMethods

`func (o *CreateInvoiceRequest) GetPaymentMethods() []string`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *CreateInvoiceRequest) GetPaymentMethodsOk() (*[]string, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *CreateInvoiceRequest) SetPaymentMethods(v []string)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *CreateInvoiceRequest) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetMidLabel

`func (o *CreateInvoiceRequest) GetMidLabel() string`

GetMidLabel returns the MidLabel field if non-nil, zero value otherwise.

### GetMidLabelOk

`func (o *CreateInvoiceRequest) GetMidLabelOk() (*string, bool)`

GetMidLabelOk returns a tuple with the MidLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMidLabel

`func (o *CreateInvoiceRequest) SetMidLabel(v string)`

SetMidLabel sets MidLabel field to given value.

### HasMidLabel

`func (o *CreateInvoiceRequest) HasMidLabel() bool`

HasMidLabel returns a boolean if a field has been set.

### GetShouldAuthenticateCreditCard

`func (o *CreateInvoiceRequest) GetShouldAuthenticateCreditCard() bool`

GetShouldAuthenticateCreditCard returns the ShouldAuthenticateCreditCard field if non-nil, zero value otherwise.

### GetShouldAuthenticateCreditCardOk

`func (o *CreateInvoiceRequest) GetShouldAuthenticateCreditCardOk() (*bool, bool)`

GetShouldAuthenticateCreditCardOk returns a tuple with the ShouldAuthenticateCreditCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldAuthenticateCreditCard

`func (o *CreateInvoiceRequest) SetShouldAuthenticateCreditCard(v bool)`

SetShouldAuthenticateCreditCard sets ShouldAuthenticateCreditCard field to given value.

### HasShouldAuthenticateCreditCard

`func (o *CreateInvoiceRequest) HasShouldAuthenticateCreditCard() bool`

HasShouldAuthenticateCreditCard returns a boolean if a field has been set.

### GetCurrency

`func (o *CreateInvoiceRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateInvoiceRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateInvoiceRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateInvoiceRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetReminderTime

`func (o *CreateInvoiceRequest) GetReminderTime() float32`

GetReminderTime returns the ReminderTime field if non-nil, zero value otherwise.

### GetReminderTimeOk

`func (o *CreateInvoiceRequest) GetReminderTimeOk() (*float32, bool)`

GetReminderTimeOk returns a tuple with the ReminderTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderTime

`func (o *CreateInvoiceRequest) SetReminderTime(v float32)`

SetReminderTime sets ReminderTime field to given value.

### HasReminderTime

`func (o *CreateInvoiceRequest) HasReminderTime() bool`

HasReminderTime returns a boolean if a field has been set.

### GetLocale

`func (o *CreateInvoiceRequest) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *CreateInvoiceRequest) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *CreateInvoiceRequest) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *CreateInvoiceRequest) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetReminderTimeUnit

`func (o *CreateInvoiceRequest) GetReminderTimeUnit() string`

GetReminderTimeUnit returns the ReminderTimeUnit field if non-nil, zero value otherwise.

### GetReminderTimeUnitOk

`func (o *CreateInvoiceRequest) GetReminderTimeUnitOk() (*string, bool)`

GetReminderTimeUnitOk returns a tuple with the ReminderTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderTimeUnit

`func (o *CreateInvoiceRequest) SetReminderTimeUnit(v string)`

SetReminderTimeUnit sets ReminderTimeUnit field to given value.

### HasReminderTimeUnit

`func (o *CreateInvoiceRequest) HasReminderTimeUnit() bool`

HasReminderTimeUnit returns a boolean if a field has been set.

### GetItems

`func (o *CreateInvoiceRequest) GetItems() []InvoiceItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CreateInvoiceRequest) GetItemsOk() (*[]InvoiceItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CreateInvoiceRequest) SetItems(v []InvoiceItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *CreateInvoiceRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetFees

`func (o *CreateInvoiceRequest) GetFees() []InvoiceFee`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *CreateInvoiceRequest) GetFeesOk() (*[]InvoiceFee, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *CreateInvoiceRequest) SetFees(v []InvoiceFee)`

SetFees sets Fees field to given value.

### HasFees

`func (o *CreateInvoiceRequest) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetChannelProperties

`func (o *CreateInvoiceRequest) GetChannelProperties() ChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *CreateInvoiceRequest) GetChannelPropertiesOk() (*ChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *CreateInvoiceRequest) SetChannelProperties(v ChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *CreateInvoiceRequest) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateInvoiceRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateInvoiceRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateInvoiceRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateInvoiceRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to README]](../../README.md)


