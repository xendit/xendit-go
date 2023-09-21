/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.5.0
*/


package invoice

import (
	"encoding/json"
	
	"fmt"
)

// NotificationChannel Representing a notification channel for sending messages.
type NotificationChannel string

// List of NotificationChannel
const (
	NOTIFICATIONCHANNEL_EMAIL NotificationChannel = "email"
	NOTIFICATIONCHANNEL_SMS NotificationChannel = "sms"
	NOTIFICATIONCHANNEL_WHATSAPP NotificationChannel = "whatsapp"
	NOTIFICATIONCHANNEL_VIBER NotificationChannel = "viber"
)

// All allowed values of NotificationChannel enum
var AllowedNotificationChannelEnumValues = []NotificationChannel{
	"email",
	"sms",
	"whatsapp",
	"viber",
}

func (v *NotificationChannel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NotificationChannel(value)
	for _, existing := range AllowedNotificationChannelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NotificationChannel", value)
}

// NewNotificationChannelFromValue returns a pointer to a valid NotificationChannel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotificationChannelFromValue(v string) (*NotificationChannel, error) {
	ev := NotificationChannel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotificationChannel: valid values are %v", v, AllowedNotificationChannelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotificationChannel) IsValid() bool {
	for _, existing := range AllowedNotificationChannelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v NotificationChannel) String() string {
	return string(v)
}

// Ptr returns reference to NotificationChannel value
func (v NotificationChannel) Ptr() *NotificationChannel {
	return &v
}

type NullableNotificationChannel struct {
	value *NotificationChannel
	isSet bool
}

func (v NullableNotificationChannel) Get() *NotificationChannel {
	return v.value
}

func (v *NullableNotificationChannel) Set(val *NotificationChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationChannel(val *NotificationChannel) *NullableNotificationChannel {
	return &NullableNotificationChannel{value: val, isSet: true}
}

func (v NullableNotificationChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
