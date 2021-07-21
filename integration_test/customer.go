package main

import (
	"fmt"
	"log"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/customer"
)

func customerTest() {
	customerAddress := xendit.CustomerAddress{
		Country:		"ID",
		StreetLine1:	"Jl. 123",
		StreetLine2:    "Jl. 456",
		City:			"Jakarta Selatan",
		Province:       "DKI Jakarta",
		State:			"-",
		PostalCode:     "12345",
	}

	metadata := map[string]interface{}{
		"meta": "data",
	}

	createCustomerData := customer.CreateCustomerParams{
		ReferenceID: 	"test-reference-id-004",
		Email:			"tes@tes.com",
		GivenNames:     "Given Names",
		Nationality: 	"ID",
		DateOfBirth: 	"1995-12-30",
		Addresses:		[]xendit.CustomerAddress{customerAddress},
		Metadata:		metadata,
	}

	resp, err := customer.CreateCustomer(&createCustomerData)
	if err != nil {
		log.Panic(err)
	}
	getCustomerByReferenceIDData := customer.GetCustomerByReferenceIDParams{
		ReferenceID: resp.ReferenceID,
	}

	_, err = customer.GetCustomerByReferenceID(&getCustomerByReferenceIDData)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Customer integration tests done!")
}
