package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/recurringpayment"
)

func recurringpaymentTest() {
	createData := recurringpayment.CreateParams{
		ExternalID:    "recurringpayment-" + time.Now().String(),
		Amount:        200000,
		PayerEmail:    "customer@customer.com",
		Description:   "recurringpayment #1",
		Interval:      xendit.RecurringPaymentIntervalDay,
		IntervalCount: 3,
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
