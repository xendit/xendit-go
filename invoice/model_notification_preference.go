/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.9.0
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
)

// checks if the NotificationPreference type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NotificationPreference{}

// NotificationPreference An object representing notification preferences for different invoice events.
type NotificationPreference struct {
	// Notification channels for when an invoice is created.
	InvoiceCreated []NotificationChannel `json:"invoice_created,omitempty"`
	// Notification channels for invoice reminders.
	InvoiceReminder []NotificationChannel `json:"invoice_reminder,omitempty"`
	// Notification channels for when an invoice is paid.
	InvoicePaid []NotificationChannel `json:"invoice_paid,omitempty"`
}

// NewNotificationPreference instantiates a new NotificationPreference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationPreference() *NotificationPreference {
	this := NotificationPreference{}
	return &this
}

// NewNotificationPreferenceWithDefaults instantiates a new NotificationPreference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationPreferenceWithDefaults() *NotificationPreference {
	this := NotificationPreference{}
	return &this
}

// GetInvoiceCreated returns the InvoiceCreated field value if set, zero value otherwise.
func (o *NotificationPreference) GetInvoiceCreated() []NotificationChannel {
	if o == nil || utils.IsNil(o.InvoiceCreated) {
		var ret []NotificationChannel
		return ret
	}
	return o.InvoiceCreated
}

// GetInvoiceCreatedOk returns a tuple with the InvoiceCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationPreference) GetInvoiceCreatedOk() ([]NotificationChannel, bool) {
	if o == nil || utils.IsNil(o.InvoiceCreated) {
		return nil, false
	}
	return o.InvoiceCreated, true
}

// HasInvoiceCreated returns a boolean if a field has been set.
func (o *NotificationPreference) HasInvoiceCreated() bool {
	if o != nil && !utils.IsNil(o.InvoiceCreated) {
		return true
	}

	return false
}

// SetInvoiceCreated gets a reference to the given []NotificationChannel and assigns it to the InvoiceCreated field.
func (o *NotificationPreference) SetInvoiceCreated(v []NotificationChannel) {
	o.InvoiceCreated = v
}

// GetInvoiceReminder returns the InvoiceReminder field value if set, zero value otherwise.
func (o *NotificationPreference) GetInvoiceReminder() []NotificationChannel {
	if o == nil || utils.IsNil(o.InvoiceReminder) {
		var ret []NotificationChannel
		return ret
	}
	return o.InvoiceReminder
}

// GetInvoiceReminderOk returns a tuple with the InvoiceReminder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationPreference) GetInvoiceReminderOk() ([]NotificationChannel, bool) {
	if o == nil || utils.IsNil(o.InvoiceReminder) {
		return nil, false
	}
	return o.InvoiceReminder, true
}

// HasInvoiceReminder returns a boolean if a field has been set.
func (o *NotificationPreference) HasInvoiceReminder() bool {
	if o != nil && !utils.IsNil(o.InvoiceReminder) {
		return true
	}

	return false
}

// SetInvoiceReminder gets a reference to the given []NotificationChannel and assigns it to the InvoiceReminder field.
func (o *NotificationPreference) SetInvoiceReminder(v []NotificationChannel) {
	o.InvoiceReminder = v
}

// GetInvoicePaid returns the InvoicePaid field value if set, zero value otherwise.
func (o *NotificationPreference) GetInvoicePaid() []NotificationChannel {
	if o == nil || utils.IsNil(o.InvoicePaid) {
		var ret []NotificationChannel
		return ret
	}
	return o.InvoicePaid
}

// GetInvoicePaidOk returns a tuple with the InvoicePaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationPreference) GetInvoicePaidOk() ([]NotificationChannel, bool) {
	if o == nil || utils.IsNil(o.InvoicePaid) {
		return nil, false
	}
	return o.InvoicePaid, true
}

// HasInvoicePaid returns a boolean if a field has been set.
func (o *NotificationPreference) HasInvoicePaid() bool {
	if o != nil && !utils.IsNil(o.InvoicePaid) {
		return true
	}

	return false
}

// SetInvoicePaid gets a reference to the given []NotificationChannel and assigns it to the InvoicePaid field.
func (o *NotificationPreference) SetInvoicePaid(v []NotificationChannel) {
	o.InvoicePaid = v
}

func (o NotificationPreference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationPreference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.InvoiceCreated) {
		toSerialize["invoice_created"] = o.InvoiceCreated
	}
	if !utils.IsNil(o.InvoiceReminder) {
		toSerialize["invoice_reminder"] = o.InvoiceReminder
	}
	if !utils.IsNil(o.InvoicePaid) {
		toSerialize["invoice_paid"] = o.InvoicePaid
	}
	return toSerialize, nil
}

type NullableNotificationPreference struct {
	value *NotificationPreference
	isSet bool
}

func (v NullableNotificationPreference) Get() *NotificationPreference {
	return v.value
}

func (v *NullableNotificationPreference) Set(val *NotificationPreference) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationPreference) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationPreference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationPreference(val *NotificationPreference) *NullableNotificationPreference {
	return &NullableNotificationPreference{value: val, isSet: true}
}

func (v NullableNotificationPreference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationPreference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


