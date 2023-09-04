# ReceiptNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailTo** | Pointer to **[]string** | Valid email address to send the payout receipt | [optional] 
**EmailCc** | Pointer to **[]string** | Valid email address to cc the payout receipt | [optional] 
**EmailBcc** | Pointer to **[]string** | Valid email address to bcc the payout receipt | [optional] 

## Methods

### NewReceiptNotification

`func NewReceiptNotification() *ReceiptNotification`

NewReceiptNotification instantiates a new ReceiptNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptNotificationWithDefaults

`func NewReceiptNotificationWithDefaults() *ReceiptNotification`

NewReceiptNotificationWithDefaults instantiates a new ReceiptNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailTo

`func (o *ReceiptNotification) GetEmailTo() []string`

GetEmailTo returns the EmailTo field if non-nil, zero value otherwise.

### GetEmailToOk

`func (o *ReceiptNotification) GetEmailToOk() (*[]string, bool)`

GetEmailToOk returns a tuple with the EmailTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTo

`func (o *ReceiptNotification) SetEmailTo(v []string)`

SetEmailTo sets EmailTo field to given value.

### HasEmailTo

`func (o *ReceiptNotification) HasEmailTo() bool`

HasEmailTo returns a boolean if a field has been set.

### SetEmailToNil

`func (o *ReceiptNotification) SetEmailToNil(b bool)`

 SetEmailToNil sets the value for EmailTo to be an explicit nil

### UnsetEmailTo
`func (o *ReceiptNotification) UnsetEmailTo()`

UnsetEmailTo ensures that no value is present for EmailTo, not even an explicit nil
### GetEmailCc

`func (o *ReceiptNotification) GetEmailCc() []string`

GetEmailCc returns the EmailCc field if non-nil, zero value otherwise.

### GetEmailCcOk

`func (o *ReceiptNotification) GetEmailCcOk() (*[]string, bool)`

GetEmailCcOk returns a tuple with the EmailCc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCc

`func (o *ReceiptNotification) SetEmailCc(v []string)`

SetEmailCc sets EmailCc field to given value.

### HasEmailCc

`func (o *ReceiptNotification) HasEmailCc() bool`

HasEmailCc returns a boolean if a field has been set.

### SetEmailCcNil

`func (o *ReceiptNotification) SetEmailCcNil(b bool)`

 SetEmailCcNil sets the value for EmailCc to be an explicit nil

### UnsetEmailCc
`func (o *ReceiptNotification) UnsetEmailCc()`

UnsetEmailCc ensures that no value is present for EmailCc, not even an explicit nil
### GetEmailBcc

`func (o *ReceiptNotification) GetEmailBcc() []string`

GetEmailBcc returns the EmailBcc field if non-nil, zero value otherwise.

### GetEmailBccOk

`func (o *ReceiptNotification) GetEmailBccOk() (*[]string, bool)`

GetEmailBccOk returns a tuple with the EmailBcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailBcc

`func (o *ReceiptNotification) SetEmailBcc(v []string)`

SetEmailBcc sets EmailBcc field to given value.

### HasEmailBcc

`func (o *ReceiptNotification) HasEmailBcc() bool`

HasEmailBcc returns a boolean if a field has been set.

### SetEmailBccNil

`func (o *ReceiptNotification) SetEmailBccNil(b bool)`

 SetEmailBccNil sets the value for EmailBcc to be an explicit nil

### UnsetEmailBcc
`func (o *ReceiptNotification) UnsetEmailBcc()`

UnsetEmailBcc ensures that no value is present for EmailBcc, not even an explicit nil

[[Back to README]](../../README.md)


