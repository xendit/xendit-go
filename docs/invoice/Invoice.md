# Invoice
An object representing details for an invoice.

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Id** | Pointer to **string** |  | The unique identifier for the invoice. |  |
| **ExternalId** | **string** | ☑️ | The external identifier for the invoice. |  |
| **UserId** | **string** | ☑️ | The user ID associated with the invoice. |  |
| **PayerEmail** | Pointer to **string** |  | The email address of the payer. |  |
| **Description** | Pointer to **string** |  | A description of the invoice. |  |
| **PaymentMethod** | Pointer to [**InvoicePaymentMethod**](InvoicePaymentMethod.md) |  |  |  |
| **Status** | [**InvoiceStatus**](InvoiceStatus.md) | ☑️ |  |  |
| **MerchantName** | **string** | ☑️ | The name of the merchant. |  |
| **MerchantProfilePictureUrl** | **string** | ☑️ | The URL of the merchant&#39;s profile picture. |  |
| **Locale** | Pointer to **string** |  | The locale or language used for the invoice. |  |
| **Amount** | **float64** | ☑️ | The total amount of the invoice. |  |
| **ExpiryDate** | **time.Time** | ☑️ | Representing a date and time in ISO 8601 format. |  |
| **InvoiceUrl** | **string** | ☑️ | The URL to view the invoice. |  |
| **AvailableBanks** | [**Bank[]**](Bank.md) | ☑️ | An array of available banks for payment. |  |
| **AvailableRetailOutlets** | [**RetailOutlet[]**](RetailOutlet.md) | ☑️ | An array of available retail outlets for payment. |  |
| **AvailableEwallets** | [**Ewallet[]**](Ewallet.md) | ☑️ | An array of available e-wallets for payment. |  |
| **AvailableQrCodes** | [**QrCode[]**](QrCode.md) | ☑️ | An array of available QR codes for payment. |  |
| **AvailableDirectDebits** | [**DirectDebit[]**](DirectDebit.md) | ☑️ | An array of available direct debit options for payment. |  |
| **AvailablePaylaters** | [**Paylater[]**](Paylater.md) | ☑️ | An array of available pay-later options for payment. |  |
| **ShouldExcludeCreditCard** | Pointer to **bool** |  | Indicates whether credit card payments should be excluded. |  |
| **ShouldSendEmail** | **bool** | ☑️ | Indicates whether email notifications should be sent. |  |
| **Created** | **time.Time** | ☑️ | Representing a date and time in ISO 8601 format. |  |
| **Updated** | **time.Time** | ☑️ | Representing a date and time in ISO 8601 format. |  |
| **SuccessRedirectUrl** | Pointer to **string** |  | The URL to redirect to on successful payment. |  |
| **FailureRedirectUrl** | Pointer to **string** |  | The URL to redirect to on payment failure. |  |
| **ShouldAuthenticateCreditCard** | Pointer to **bool** |  | Indicates whether credit card authentication is required. |  |
| **Currency** | Pointer to [**InvoiceCurrency**](InvoiceCurrency.md) |  |  |  |
| **Items** | Pointer to [**InvoiceItem[]**](InvoiceItem.md) |  | An array of items included in the invoice. |  |
| **FixedVa** | Pointer to **bool** |  | Indicates whether the virtual account is fixed. |  |
| **ReminderDate** | Pointer to **time.Time** |  | Representing a date and time in ISO 8601 format. |  |
| **Customer** | Pointer to [**CustomerObject**](CustomerObject.md) |  |  |  |
| **CustomerNotificationPreference** | Pointer to [**NotificationPreference**](NotificationPreference.md) |  |  |  |
| **Fees** | Pointer to [**InvoiceFee[]**](InvoiceFee.md) |  | An array of fees associated with the invoice. |  |

## Methods

### NewInvoice

`func NewInvoice(externalId string, userId string, status InvoiceStatus, merchantName string, merchantProfilePictureUrl string, amount float64, expiryDate time.Time, invoiceUrl string, availableBanks []Bank, availableRetailOutlets []RetailOutlet, availableEwallets []Ewallet, availableQrCodes []QrCode, availableDirectDebits []DirectDebit, availablePaylaters []Paylater, shouldSendEmail bool, created time.Time, updated time.Time, ) *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Invoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invoice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Invoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalId

`func (o *Invoice) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Invoice) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Invoice) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetUserId

`func (o *Invoice) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Invoice) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Invoice) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetPayerEmail

`func (o *Invoice) GetPayerEmail() string`

GetPayerEmail returns the PayerEmail field if non-nil, zero value otherwise.

### GetPayerEmailOk

`func (o *Invoice) GetPayerEmailOk() (*string, bool)`

GetPayerEmailOk returns a tuple with the PayerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerEmail

`func (o *Invoice) SetPayerEmail(v string)`

SetPayerEmail sets PayerEmail field to given value.

### HasPayerEmail

`func (o *Invoice) HasPayerEmail() bool`

HasPayerEmail returns a boolean if a field has been set.

### GetDescription

`func (o *Invoice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Invoice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Invoice) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Invoice) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *Invoice) GetPaymentMethod() InvoicePaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *Invoice) GetPaymentMethodOk() (*InvoicePaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *Invoice) SetPaymentMethod(v InvoicePaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *Invoice) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetStatus

`func (o *Invoice) GetStatus() InvoiceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Invoice) GetStatusOk() (*InvoiceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Invoice) SetStatus(v InvoiceStatus)`

SetStatus sets Status field to given value.


### GetMerchantName

`func (o *Invoice) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *Invoice) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *Invoice) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.


### GetMerchantProfilePictureUrl

`func (o *Invoice) GetMerchantProfilePictureUrl() string`

GetMerchantProfilePictureUrl returns the MerchantProfilePictureUrl field if non-nil, zero value otherwise.

### GetMerchantProfilePictureUrlOk

`func (o *Invoice) GetMerchantProfilePictureUrlOk() (*string, bool)`

GetMerchantProfilePictureUrlOk returns a tuple with the MerchantProfilePictureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProfilePictureUrl

`func (o *Invoice) SetMerchantProfilePictureUrl(v string)`

SetMerchantProfilePictureUrl sets MerchantProfilePictureUrl field to given value.


### GetLocale

`func (o *Invoice) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *Invoice) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *Invoice) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *Invoice) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetAmount

`func (o *Invoice) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Invoice) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Invoice) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetExpiryDate

`func (o *Invoice) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *Invoice) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *Invoice) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.


### GetInvoiceUrl

`func (o *Invoice) GetInvoiceUrl() string`

GetInvoiceUrl returns the InvoiceUrl field if non-nil, zero value otherwise.

### GetInvoiceUrlOk

`func (o *Invoice) GetInvoiceUrlOk() (*string, bool)`

GetInvoiceUrlOk returns a tuple with the InvoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceUrl

`func (o *Invoice) SetInvoiceUrl(v string)`

SetInvoiceUrl sets InvoiceUrl field to given value.


### GetAvailableBanks

`func (o *Invoice) GetAvailableBanks() []Bank`

GetAvailableBanks returns the AvailableBanks field if non-nil, zero value otherwise.

### GetAvailableBanksOk

`func (o *Invoice) GetAvailableBanksOk() (*[]Bank, bool)`

GetAvailableBanksOk returns a tuple with the AvailableBanks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBanks

`func (o *Invoice) SetAvailableBanks(v []Bank)`

SetAvailableBanks sets AvailableBanks field to given value.


### GetAvailableRetailOutlets

`func (o *Invoice) GetAvailableRetailOutlets() []RetailOutlet`

GetAvailableRetailOutlets returns the AvailableRetailOutlets field if non-nil, zero value otherwise.

### GetAvailableRetailOutletsOk

`func (o *Invoice) GetAvailableRetailOutletsOk() (*[]RetailOutlet, bool)`

GetAvailableRetailOutletsOk returns a tuple with the AvailableRetailOutlets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableRetailOutlets

`func (o *Invoice) SetAvailableRetailOutlets(v []RetailOutlet)`

SetAvailableRetailOutlets sets AvailableRetailOutlets field to given value.


### GetAvailableEwallets

`func (o *Invoice) GetAvailableEwallets() []Ewallet`

GetAvailableEwallets returns the AvailableEwallets field if non-nil, zero value otherwise.

### GetAvailableEwalletsOk

`func (o *Invoice) GetAvailableEwalletsOk() (*[]Ewallet, bool)`

GetAvailableEwalletsOk returns a tuple with the AvailableEwallets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableEwallets

`func (o *Invoice) SetAvailableEwallets(v []Ewallet)`

SetAvailableEwallets sets AvailableEwallets field to given value.


### GetAvailableQrCodes

`func (o *Invoice) GetAvailableQrCodes() []QrCode`

GetAvailableQrCodes returns the AvailableQrCodes field if non-nil, zero value otherwise.

### GetAvailableQrCodesOk

`func (o *Invoice) GetAvailableQrCodesOk() (*[]QrCode, bool)`

GetAvailableQrCodesOk returns a tuple with the AvailableQrCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableQrCodes

`func (o *Invoice) SetAvailableQrCodes(v []QrCode)`

SetAvailableQrCodes sets AvailableQrCodes field to given value.


### GetAvailableDirectDebits

`func (o *Invoice) GetAvailableDirectDebits() []DirectDebit`

GetAvailableDirectDebits returns the AvailableDirectDebits field if non-nil, zero value otherwise.

### GetAvailableDirectDebitsOk

`func (o *Invoice) GetAvailableDirectDebitsOk() (*[]DirectDebit, bool)`

GetAvailableDirectDebitsOk returns a tuple with the AvailableDirectDebits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableDirectDebits

`func (o *Invoice) SetAvailableDirectDebits(v []DirectDebit)`

SetAvailableDirectDebits sets AvailableDirectDebits field to given value.


### GetAvailablePaylaters

`func (o *Invoice) GetAvailablePaylaters() []Paylater`

GetAvailablePaylaters returns the AvailablePaylaters field if non-nil, zero value otherwise.

### GetAvailablePaylatersOk

`func (o *Invoice) GetAvailablePaylatersOk() (*[]Paylater, bool)`

GetAvailablePaylatersOk returns a tuple with the AvailablePaylaters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePaylaters

`func (o *Invoice) SetAvailablePaylaters(v []Paylater)`

SetAvailablePaylaters sets AvailablePaylaters field to given value.


### GetShouldExcludeCreditCard

`func (o *Invoice) GetShouldExcludeCreditCard() bool`

GetShouldExcludeCreditCard returns the ShouldExcludeCreditCard field if non-nil, zero value otherwise.

### GetShouldExcludeCreditCardOk

`func (o *Invoice) GetShouldExcludeCreditCardOk() (*bool, bool)`

GetShouldExcludeCreditCardOk returns a tuple with the ShouldExcludeCreditCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldExcludeCreditCard

`func (o *Invoice) SetShouldExcludeCreditCard(v bool)`

SetShouldExcludeCreditCard sets ShouldExcludeCreditCard field to given value.

### HasShouldExcludeCreditCard

`func (o *Invoice) HasShouldExcludeCreditCard() bool`

HasShouldExcludeCreditCard returns a boolean if a field has been set.

### GetShouldSendEmail

`func (o *Invoice) GetShouldSendEmail() bool`

GetShouldSendEmail returns the ShouldSendEmail field if non-nil, zero value otherwise.

### GetShouldSendEmailOk

`func (o *Invoice) GetShouldSendEmailOk() (*bool, bool)`

GetShouldSendEmailOk returns a tuple with the ShouldSendEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldSendEmail

`func (o *Invoice) SetShouldSendEmail(v bool)`

SetShouldSendEmail sets ShouldSendEmail field to given value.


### GetCreated

`func (o *Invoice) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Invoice) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Invoice) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *Invoice) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Invoice) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Invoice) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### GetSuccessRedirectUrl

`func (o *Invoice) GetSuccessRedirectUrl() string`

GetSuccessRedirectUrl returns the SuccessRedirectUrl field if non-nil, zero value otherwise.

### GetSuccessRedirectUrlOk

`func (o *Invoice) GetSuccessRedirectUrlOk() (*string, bool)`

GetSuccessRedirectUrlOk returns a tuple with the SuccessRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRedirectUrl

`func (o *Invoice) SetSuccessRedirectUrl(v string)`

SetSuccessRedirectUrl sets SuccessRedirectUrl field to given value.

### HasSuccessRedirectUrl

`func (o *Invoice) HasSuccessRedirectUrl() bool`

HasSuccessRedirectUrl returns a boolean if a field has been set.

### GetFailureRedirectUrl

`func (o *Invoice) GetFailureRedirectUrl() string`

GetFailureRedirectUrl returns the FailureRedirectUrl field if non-nil, zero value otherwise.

### GetFailureRedirectUrlOk

`func (o *Invoice) GetFailureRedirectUrlOk() (*string, bool)`

GetFailureRedirectUrlOk returns a tuple with the FailureRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureRedirectUrl

`func (o *Invoice) SetFailureRedirectUrl(v string)`

SetFailureRedirectUrl sets FailureRedirectUrl field to given value.

### HasFailureRedirectUrl

`func (o *Invoice) HasFailureRedirectUrl() bool`

HasFailureRedirectUrl returns a boolean if a field has been set.

### GetShouldAuthenticateCreditCard

`func (o *Invoice) GetShouldAuthenticateCreditCard() bool`

GetShouldAuthenticateCreditCard returns the ShouldAuthenticateCreditCard field if non-nil, zero value otherwise.

### GetShouldAuthenticateCreditCardOk

`func (o *Invoice) GetShouldAuthenticateCreditCardOk() (*bool, bool)`

GetShouldAuthenticateCreditCardOk returns a tuple with the ShouldAuthenticateCreditCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldAuthenticateCreditCard

`func (o *Invoice) SetShouldAuthenticateCreditCard(v bool)`

SetShouldAuthenticateCreditCard sets ShouldAuthenticateCreditCard field to given value.

### HasShouldAuthenticateCreditCard

`func (o *Invoice) HasShouldAuthenticateCreditCard() bool`

HasShouldAuthenticateCreditCard returns a boolean if a field has been set.

### GetCurrency

`func (o *Invoice) GetCurrency() InvoiceCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Invoice) GetCurrencyOk() (*InvoiceCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Invoice) SetCurrency(v InvoiceCurrency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Invoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetItems

`func (o *Invoice) GetItems() []InvoiceItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Invoice) GetItemsOk() (*[]InvoiceItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Invoice) SetItems(v []InvoiceItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *Invoice) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetFixedVa

`func (o *Invoice) GetFixedVa() bool`

GetFixedVa returns the FixedVa field if non-nil, zero value otherwise.

### GetFixedVaOk

`func (o *Invoice) GetFixedVaOk() (*bool, bool)`

GetFixedVaOk returns a tuple with the FixedVa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedVa

`func (o *Invoice) SetFixedVa(v bool)`

SetFixedVa sets FixedVa field to given value.

### HasFixedVa

`func (o *Invoice) HasFixedVa() bool`

HasFixedVa returns a boolean if a field has been set.

### GetReminderDate

`func (o *Invoice) GetReminderDate() time.Time`

GetReminderDate returns the ReminderDate field if non-nil, zero value otherwise.

### GetReminderDateOk

`func (o *Invoice) GetReminderDateOk() (*time.Time, bool)`

GetReminderDateOk returns a tuple with the ReminderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderDate

`func (o *Invoice) SetReminderDate(v time.Time)`

SetReminderDate sets ReminderDate field to given value.

### HasReminderDate

`func (o *Invoice) HasReminderDate() bool`

HasReminderDate returns a boolean if a field has been set.

### GetCustomer

`func (o *Invoice) GetCustomer() CustomerObject`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *Invoice) GetCustomerOk() (*CustomerObject, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *Invoice) SetCustomer(v CustomerObject)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *Invoice) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetCustomerNotificationPreference

`func (o *Invoice) GetCustomerNotificationPreference() NotificationPreference`

GetCustomerNotificationPreference returns the CustomerNotificationPreference field if non-nil, zero value otherwise.

### GetCustomerNotificationPreferenceOk

`func (o *Invoice) GetCustomerNotificationPreferenceOk() (*NotificationPreference, bool)`

GetCustomerNotificationPreferenceOk returns a tuple with the CustomerNotificationPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNotificationPreference

`func (o *Invoice) SetCustomerNotificationPreference(v NotificationPreference)`

SetCustomerNotificationPreference sets CustomerNotificationPreference field to given value.

### HasCustomerNotificationPreference

`func (o *Invoice) HasCustomerNotificationPreference() bool`

HasCustomerNotificationPreference returns a boolean if a field has been set.

### GetFees

`func (o *Invoice) GetFees() []InvoiceFee`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *Invoice) GetFeesOk() (*[]InvoiceFee, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *Invoice) SetFees(v []InvoiceFee)`

SetFees sets Fees field to given value.

### HasFees

`func (o *Invoice) HasFees() bool`

HasFees returns a boolean if a field has been set.


[[Back to README]](../../README.md)


