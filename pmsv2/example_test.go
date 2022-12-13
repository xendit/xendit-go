package pmsv2_test

import (
	"fmt"
	"log"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/pmsv2"
)

func ExampleCreatePaymentMethod() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := pmsv2.CreatePaymentMethodParams{}

	resp, err := pmsv2.CreatePaymentMethod(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created payment: %+v\n", resp)
}

func ExampleAuthPaymentMethod() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := pmsv2.ValidateOTPRequest{}

	resp, err := pmsv2.AuthPaymentMethod(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created payment: %+v\n", resp)
}

func ExampleExpirePaymentMethod() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := pmsv2.ExpireRequest{}

	resp, err := pmsv2.ExpirePaymentMethod(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created payment: %+v\n", resp)
}

func ExampleRetrievePaymentMethod() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := pmsv2.RetrievePaymentMethodRequest{}

	resp, err := pmsv2.RetrievePaymentMethod(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created payment: %+v\n", resp)
}

func ExampleRetrieveAllPaymentMethods() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := pmsv2.RetrieveAllPaymentMethodsRequest{}

	resp, err := pmsv2.RetrieveAllPaymentMethods(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created payment: %+v\n", resp)
}

func ExampleUpdatePaymentMethod() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := pmsv2.UpdateRequest{}

	resp, err := pmsv2.UpdatePaymentMethod(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created payment: %+v\n", resp)
}

func ExampleRetrievePayments() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := pmsv2.RetrievePaymentsRequest{}

	resp, err := pmsv2.RetrievePayments(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created payment: %+v\n", resp)
}
