package customer_test

import (
	"fmt"
	"log"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/customer"
)

func ExampleCreateCustomer() {
	xendit.Opt.SecretKey = "examplesecretkey"

	customerAddress := xendit.CustomerAddress{
		Country:     "ID",
		StreetLine1: "Jl. 123",
		StreetLine2: "Jl. 456",
		City:        "Jakarta Selatan",
		Province:    "DKI Jakarta",
		State:       "-",
		PostalCode:  "12345",
	}

	metadata := map[string]interface{}{
		"meta": "data",
	}

	data := customer.CreateCustomerParams{
		ReferenceID: "test-reference-id-002",
		Email:       "tes@tes.com",
		GivenNames:  "Given Names",
		Nationality: "ID",
		DateOfBirth: "1992-12-30",
		Addresses:   []xendit.CustomerAddress{customerAddress},
		Metadata:    metadata,
	}

	resp, err := customer.CreateCustomer(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created customer: %+v\n", resp)
}

func ExampleGetCustomerByReferenceID() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := customer.GetCustomerByReferenceIDParams{
		ReferenceID: "test-reference-id-002",
	}

	resp, err := customer.GetCustomerByReferenceID(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved customer: %+v\n", resp)
}
