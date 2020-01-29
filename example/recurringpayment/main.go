package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/recurringpayment"
)

func main() {
	xendit.Opt.SecretKey = "xnd_development_REt02KJzkM6AootfKnDrMw1Sse4LlzEDHfKzXoBocqIEiH4bqjHUJXbl6Cfaab"

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
		log.Fatal(err)
	}
	fmt.Printf("created recurring payment: %+v\n", resp)

	getData := recurringpayment.GetParams{
		ID: resp.ID,
	}

	resp, err = recurringpayment.Get(&getData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retrieved recurring payment: %+v\n", resp)

	editData := recurringpayment.EditParams{
		ID:     resp.ID,
		Amount: 500000,
	}

	resp, err = recurringpayment.Edit(&editData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("edited recurring payment: %+v\n", resp)

	pauseData := recurringpayment.PauseParams{
		ID: resp.ID,
	}

	resp, err = recurringpayment.Pause(&pauseData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("paused recurring payment: %+v\n", resp)

	resumeData := recurringpayment.ResumeParams{
		ID: resp.ID,
	}

	resp, err = recurringpayment.Resume(&resumeData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("resumed recurring payment: %+v\n", resp)

	stopData := recurringpayment.StopParams{
		ID: resp.ID,
	}

	resp, err = recurringpayment.Stop(&stopData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("stopped recurring payment: %+v\n", resp)
}
