package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go/invoice"
)

func invoiceTest() {
	data := invoice.CreateParams{
		ExternalID:  "invoice-" + time.Now().String(),
		Amount:      200000,
		PayerEmail:  "customer@customer.com",
		Description: "invoice  #1",
	}
	resp, err := invoice.Create(&data)
	if err != nil {
		log.Panic(err)

	}
	resp, err = invoice.Get(&invoice.GetParams{
		ID: resp.ID,
	})
	if err != nil {
		log.Panic(err)
	}
	_, err = invoice.Expire(&invoice.ExpireParams{
		ID: resp.ID,
	})
	if err != nil {
		log.Panic(err)
	}
	_, err = invoice.GetAll(nil)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Invoice integration tests done!")
}
