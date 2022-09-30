package xendit

import "time"

// MissedPaymentActionEnum constants are the available recurring payment missed payment actions
type MissedPaymentActionEnum string

// This consists the values that MissedPaymentActionEnum can take
const (
	MissedPaymentActionIgnore MissedPaymentActionEnum = "IGNORE"
	MissedPaymentActionStop   MissedPaymentActionEnum = "STOP"
)

// RecurringPaymentIntervalEnum constants are the available recurring payment intervals
type RecurringPaymentIntervalEnum string

// This consists the values that RecurringPaymentIntervalEnum can take
const (
	RecurringPaymentIntervalDay   RecurringPaymentIntervalEnum = "DAY"
	RecurringPaymentIntervalWeek  RecurringPaymentIntervalEnum = "WEEK"
	RecurringPaymentIntervalMonth RecurringPaymentIntervalEnum = "MONTH"
)

// RecurringPayment contains data from Xendit's API response of recurring payment related requests.
// For more details see https://xendit.github.io/apireference/?bash#recurring-payments.
// For documentation of subpackage recurringpayment, checkout https://pkg.go.dev/github.com/xendit/xendit-go/recurringpayment
type RecurringPayment struct {
	ID                             string                         `json:"id"`
	ExternalID                     string                         `json:"external_id"`
	UserID                         string                         `json:"user_id"`
	PayerEmail                     string                         `json:"payer_email"`
	Description                    string                         `json:"description"`
	Status                         string                         `json:"status"`
	Amount                         float64                        `json:"amount"`
	ShouldSendEmail                bool                           `json:"should_send_email"`
	Customer                       RecurringPaymentCustomer       `json:"customer,omitempty"`
	CustomerNotificationPreference CustomerNotificationPreference `json:"customer_notification_preference,omitempty"`
	Interval                       RecurringPaymentIntervalEnum   `json:"interval"`
	IntervalCount                  int                            `json:"interval_count"`
	MissedPaymentAction            MissedPaymentActionEnum        `json:"missed_payment_action"`
	Created                        *time.Time                     `json:"created"`
	Updated                        *time.Time                     `json:"updated"`
	InvoiceDuration                int                            `json:"invoice_duration,omitempty"`
	StartDate                      *time.Time                     `json:"start_date,omitempty"`
	LastCreatedInvoiceURL          string                         `json:"last_created_invoice_url,omitempty"`
	CreditCardToken                string                         `json:"credit_card_token,omitempty"`
	SuccessRedirectURL             string                         `json:"success_redirect_url,omitempty"`
	FailureRedirectURL             string                         `json:"failure_redirect_url,omitempty"`
	TotalRecurrence                int                            `json:"total_recurrence,omitempty"`
	RecurrenceProgress             int                            `json:"recurrence_progress,omitempty"`
	Recharge                       bool                           `json:"recharge,omitempty"`
	ChargeImmediately              bool                           `json:"charge_immediately,omitempty"`
	Currency                       string                         `json:"currency,omitempty"`
}

// RecurringPaymentCustomer is data that contained in `RecurringPayment` at Customer
type RecurringPaymentCustomer struct {
	GivenNames   string                           `json:"given_names,omitempty"`
	Email        string                           `json:"email,omitempty"`
	MobileNumber string                           `json:"mobile_number,omitempty"`
	Address      []RecurringPaymenCustomerAddress `json:"address,omitempty"`
}

// RecurringPaymenCustomerAddress is data that contained in `RecurringPaymentCustomer` at Address
type RecurringPaymenCustomerAddress struct {
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	State       string `json:"state,omitempty"`
	StreetLine1 string `json:"street_line1,omitempty"`
	StreetLine2 string `json:"street_line2,omitempty"`
}

// CustomerNotificationChannelEnum constants are the available customer notification channels
type CustomerNotificationChannelEnum string

// This consists the values that CustomerNotificationChannelEnum can take
const (
	CustomerNotificationChannelWhatsApp CustomerNotificationChannelEnum = "whatsapp"
	CustomerNotificationChannelSMS      CustomerNotificationChannelEnum = "sms"
	CustomerNotificationChannelEmail    CustomerNotificationChannelEnum = "email"
	CustomerNotificationChannelViber    CustomerNotificationChannelEnum = "viber"
)

// CustomerNotificationPreference is data that contained in `RecurringPayment` at CustomerNotificationPreference
type CustomerNotificationPreference struct {
	InvoiceCreated  []CustomerNotificationChannelEnum `json:"invoice_created,omitempty"`
	InvoiceReminder []CustomerNotificationChannelEnum `json:"invoice_reminder,omitempty"`
	InvoicePaid     []CustomerNotificationChannelEnum `json:"invoice_paid,omitempty"`
	InvoiceExpired  []CustomerNotificationChannelEnum `json:"invoice_expired,omitempty"`
}
