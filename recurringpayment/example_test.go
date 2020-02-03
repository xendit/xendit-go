package recurringpayment_test

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/recurringpayment"
)

func ExampleCreate() {
	xendit.Opt.SecretKey = "examplesecretkey"

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
}

func ExampleGet() {
	xendit.Opt.SecretKey = "examplesecretkey"

	getData := recurringpayment.GetParams{
		ID: "123",
	}

	resp, err := recurringpayment.Get(&getData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved recurring payment: %+v\n", resp)
}

func ExampleEdit() {
	xendit.Opt.SecretKey = "examplesecretkey"

	editData := recurringpayment.EditParams{
		ID:            "123",
		Amount:        500000,
		Interval:      xendit.RecurringPaymentIntervalMonth,
		IntervalCount: 3,
	}

	resp, err := recurringpayment.Edit(&editData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("edited recurring payment: %+v\n", resp)
}

func ExampleStop() {
	xendit.Opt.SecretKey = "examplesecretkey"

	stopData := recurringpayment.StopParams{
		ID: "123",
	}

	resp, err := recurringpayment.Stop(&stopData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("stopped recurring payment: %+v\n", resp)
}

func ExamplePause() {
	xendit.Opt.SecretKey = "examplesecretkey"

	pauseData := recurringpayment.PauseParams{
		ID: "123",
	}

	resp, err := recurringpayment.Pause(&pauseData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("paused recurring payment: %+v\n", resp)
}

func ExampleResume() {
	xendit.Opt.SecretKey = "examplesecretkey"

	resumeData := recurringpayment.ResumeParams{
		ID: "123",
	}

	resp, err := recurringpayment.Resume(&resumeData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("resumed recurring payment: %+v\n", resp)
}
