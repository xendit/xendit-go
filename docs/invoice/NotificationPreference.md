# NotificationPreference
An object representing notification preferences for different invoice events.

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **InvoiceCreated** | Pointer to [**NotificationChannel[]**](NotificationChannel.md) |  | Notification channels for when an invoice is created. |  |
| **InvoiceReminder** | Pointer to [**NotificationChannel[]**](NotificationChannel.md) |  | Notification channels for invoice reminders. |  |
| **InvoicePaid** | Pointer to [**NotificationChannel[]**](NotificationChannel.md) |  | Notification channels for when an invoice is paid. |  |

## Methods

### NewNotificationPreference

`func NewNotificationPreference() *NotificationPreference`

NewNotificationPreference instantiates a new NotificationPreference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationPreferenceWithDefaults

`func NewNotificationPreferenceWithDefaults() *NotificationPreference`

NewNotificationPreferenceWithDefaults instantiates a new NotificationPreference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceCreated

`func (o *NotificationPreference) GetInvoiceCreated() []NotificationChannel`

GetInvoiceCreated returns the InvoiceCreated field if non-nil, zero value otherwise.

### GetInvoiceCreatedOk

`func (o *NotificationPreference) GetInvoiceCreatedOk() (*[]NotificationChannel, bool)`

GetInvoiceCreatedOk returns a tuple with the InvoiceCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCreated

`func (o *NotificationPreference) SetInvoiceCreated(v []NotificationChannel)`

SetInvoiceCreated sets InvoiceCreated field to given value.

### HasInvoiceCreated

`func (o *NotificationPreference) HasInvoiceCreated() bool`

HasInvoiceCreated returns a boolean if a field has been set.

### GetInvoiceReminder

`func (o *NotificationPreference) GetInvoiceReminder() []NotificationChannel`

GetInvoiceReminder returns the InvoiceReminder field if non-nil, zero value otherwise.

### GetInvoiceReminderOk

`func (o *NotificationPreference) GetInvoiceReminderOk() (*[]NotificationChannel, bool)`

GetInvoiceReminderOk returns a tuple with the InvoiceReminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceReminder

`func (o *NotificationPreference) SetInvoiceReminder(v []NotificationChannel)`

SetInvoiceReminder sets InvoiceReminder field to given value.

### HasInvoiceReminder

`func (o *NotificationPreference) HasInvoiceReminder() bool`

HasInvoiceReminder returns a boolean if a field has been set.

### GetInvoicePaid

`func (o *NotificationPreference) GetInvoicePaid() []NotificationChannel`

GetInvoicePaid returns the InvoicePaid field if non-nil, zero value otherwise.

### GetInvoicePaidOk

`func (o *NotificationPreference) GetInvoicePaidOk() (*[]NotificationChannel, bool)`

GetInvoicePaidOk returns a tuple with the InvoicePaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePaid

`func (o *NotificationPreference) SetInvoicePaid(v []NotificationChannel)`

SetInvoicePaid sets InvoicePaid field to given value.

### HasInvoicePaid

`func (o *NotificationPreference) HasInvoicePaid() bool`

HasInvoicePaid returns a boolean if a field has been set.


[[Back to README]](../../README.md)


