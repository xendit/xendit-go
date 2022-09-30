package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/recurringpayment"
)

func recurringpaymentTest() {

	customerAddress := xendit.RecurringPaymenCustomerAddress{
		Country:     "ID",
		StreetLine1: "Jl. 123",
		StreetLine2: "Jl. 456",
		City:        "Jakarta Selatan",
		State:       "DKI JAKARTA",
		PostalCode:  "12345",
	}

	customerData := xendit.RecurringPaymentCustomer{
		GivenNames:   "customer 1",
		Email:        "customer@customer.com",
		MobileNumber: "+6281212345678",
		Address:      []xendit.RecurringPaymenCustomerAddress{customerAddress},
	}

	customerNotificationPreference := xendit.CustomerNotificationPreference{
		InvoiceCreated:  []xendit.CustomerNotificationChannelEnum{xendit.CustomerNotificationChannelEmail, xendit.CustomerNotificationChannelWhatsApp},
		InvoiceReminder: []xendit.CustomerNotificationChannelEnum{xendit.CustomerNotificationChannelSMS, xendit.CustomerNotificationChannelEmail},
		InvoicePaid:     []xendit.CustomerNotificationChannelEnum{xendit.CustomerNotificationChannelViber, xendit.CustomerNotificationChannelEmail},
		InvoiceExpired:  []xendit.CustomerNotificationChannelEnum{xendit.CustomerNotificationChannelWhatsApp, xendit.CustomerNotificationChannelEmail},
	}

	createData := recurringpayment.CreateParams{
		ExternalID:                     "recurringpayment-" + time.Now().String(),
		Amount:                         200000,
		PayerEmail:                     "customer@customer.com",
		Description:                    "recurringpayment #1",
		Interval:                       xendit.RecurringPaymentIntervalDay,
		IntervalCount:                  3,
		Customer:                       customerData,
		CustomerNotificationPreference: customerNotificationPreference,
	}

	resp, err := recurringpayment.Create(&createData)
	if err != nil {
		log.Panic(err)
	}

	getData := recurringpayment.GetParams{
		ID: resp.ID,
	}
	resp, err = recurringpayment.Get(&getData)
	if err != nil {
		log.Panic(err)
	}

	editData := recurringpayment.EditParams{
		ID:     resp.ID,
		Amount: 500000,
	}
	resp, err = recurringpayment.Edit(&editData)
	if err != nil {
		log.Panic(err)
	}

	pauseData := recurringpayment.PauseParams{
		ID: resp.ID,
	}
	resp, err = recurringpayment.Pause(&pauseData)
	if err != nil {
		log.Panic(err)
	}

	resumeData := recurringpayment.ResumeParams{
		ID: resp.ID,
	}
	resp, err = recurringpayment.Resume(&resumeData)
	if err != nil {
		log.Panic(err)
	}

	stopData := recurringpayment.StopParams{
		ID: resp.ID,
	}
	_, err = recurringpayment.Stop(&stopData)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Recurring Payment integration tests done!")
}
