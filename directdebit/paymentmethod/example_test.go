package paymentmethod_test

import (
	"fmt"
	"log"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/directdebit/paymentmethod"
)

func ExampleCreatePaymentMethod() {
	xendit.Opt.SecretKey = "examplesecretkey"

	properties := map[string]interface{}{
		"id": "test-la-id",
	}

	metadata := map[string]interface{}{
		"meta": "data",
	}

	data := paymentmethod.CreatePaymentMethodParams{
		CustomerID: "test-cust-id",
		Type:       xendit.DEBIT_CARD,
		Properties: properties,
		Metadata:   metadata,
	}

	resp, err := paymentmethod.CreatePaymentMethod(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created payment method: %+v\n", resp)
}

func ExampleGetPaymentMethodsByCustomerID() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := paymentmethod.GetPaymentMethodsByCustomerIDParams{
		CustomerID: "test-cust-id",
	}

	resp, err := paymentmethod.GetPaymentMethodsByCustomerID(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved payment methods: %+v\n", resp)
}
