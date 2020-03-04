package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/cardlesscredit"
)

func cardlesscreditTest() {
	createPaymentData := cardlesscredit.CreatePaymentParams{
		CardlessCreditType: xendit.CardlessCreditTypeEnumKREDIVO,
		ExternalID:         "cardless-credit-" + time.Now().String(),
		Amount:             200000,
		PaymentType:        xendit.PaymentTypeEnum3Months,
		Items: []cardlesscredit.Item{
			{
				ID:       "123",
				Name:     "Laptop Asus Ila",
				Price:    200000,
				Type:     "Laptop",
				URL:      "http://asus-ila.com",
				Quantity: 1,
			},
		},
		CustomerDetails: cardlesscredit.CustomerDetails{
			FirstName: "Michael",
			LastName:  "Belajarrock",
			Email:     "michaelbelajarrock@mail.com",
			Phone:     "08123123123",
		},
		ShippingAddress: cardlesscredit.ShippingAddress{
			FirstName:   "Michael",
			LastName:    "Belajarjazz",
			Address:     "Jalan Teknologi No. 12",
			City:        "Jakarta",
			PostalCode:  "40000",
			Phone:       "08123123123",
			CountryCode: "IDN",
		},
		RedirectURL: "https://google.com",
		CallbackURL: "https://google.com",
	}
	_, err := cardlesscredit.CreatePayment(&createPaymentData)
	if err != nil && err.GetErrorCode() != "MERCHANT_NOT_FOUND_ERROR" {
		log.Fatal(err)
	}

	fmt.Println("Cardless Credit integration tests done!")
}
