package main

import (
	"fmt"
	"log"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/ewallet"
)

func ewalletTest() {

	createEWalletChargeMinimalFields := ewallet.CreateEWalletChargeParams{
		ReferenceID:    "test-reference-id-2",
		Currency:       "IDR",
		Amount:         10000,
		CheckoutMethod: "ONE_TIME_PAYMENT",
		ChannelCode:    "ID_SHOPEEPAY",
		ChannelProperties: map[string]string{
			"success_redirect_url": "https://yourwebsite.com/order/123",
			"failure_redirect_url": "https://yourwebsite.com/failure",
		},
	}

	createChargeResponse, err := ewallet.CreateEWalletCharge(&createEWalletChargeMinimalFields)
	if err != nil {
		log.Panic(err)
	}

	getChargeResponse := ewallet.GetEWalletChargeStatusParams{
		ChargeID: createChargeResponse.ID,
	}

	_, err = ewallet.GetEWalletChargeStatus(&getChargeResponse)
	if err != nil {
		log.Panic(err)
	}

	createEWalletChargeAdditionalFields := ewallet.CreateEWalletChargeParams{
		ReferenceID:    "test-reference-id-additional-fields-2",
		Currency:       "IDR",
		Amount:         10000,
		CheckoutMethod: "ONE_TIME_PAYMENT",
		ChannelCode:    "ID_SHOPEEPAY",
		ChannelProperties: map[string]string{
			"success_redirect_url": "https://yourwebsite.com/order/123",
			"failure_redirect_url": "https://yourwebsite.com/failure",
			"cancel_redirect_url":  "https://yourwebsite.com/cancel",
		},
		CaptureNow:      true,
		CustomerID:      "",
		PaymentMethodID: "",
		Customer: &xendit.EwalletCustomer{
			ReferenceId:            func(i string) *string { return &i }("sample_customer_reference_id"),
			GivenNames:             func(i string) *string { return &i }("sample_given_name"),
			Surname:                func(i string) *string { return &i }("sample_surname"),
			Email:                  func(i string) *string { return &i }("sample_email"),
			MobileNumber:           func(i string) *string { return &i }("+639192999999"),
			DomicileOfRegistration: func(i string) *string { return &i }("ID"),
			DateOfRegistration:     func(i string) *string { return &i }("2022-11-01"),
		},
		ShippingInformation: &xendit.ShippingInformation{
			Country:       "ID",
			StreetLine1:   "sample_street_line",
			StreetLine2:   func(i string) *string { return &i }("sample_street_line_2"),
			City:          "sample_city",
			ProvinceState: "sample_province",
			PostalCode:    "sample_postal_code",
		},
		Basket:   []xendit.EWalletBasketItem{},
		Metadata: map[string]interface{}{},
	}

	createEWalletChargeAdditionalFieldsResponse, err := ewallet.CreateEWalletCharge(&createEWalletChargeAdditionalFields)
	if err != nil {
		log.Panic(err)
	}

	getEWalletChargeAdditionalFieldsResponse := ewallet.GetEWalletChargeStatusParams{
		ChargeID: createEWalletChargeAdditionalFieldsResponse.ID,
	}

	_, err = ewallet.GetEWalletChargeStatus(&getEWalletChargeAdditionalFieldsResponse)
	if err != nil {
		log.Panic(err)
	}

	createEWalletChargeTokenized := ewallet.CreateEWalletChargeParams{
		ReferenceID:    "test-reference-id-tokenized_2",
		Currency:       "IDR",
		Amount:         10000,
		CheckoutMethod: "TOKENIZED_PAYMENT",
		ChannelCode:    "ID_DANA",
		ChannelProperties: map[string]string{
			"success_redirect_url": "https://yourwebsite.com/order/123",
			"failure_redirect_url": "https://yourwebsite.com/failure",
			"cancel_redirect_url":  "https://yourwebsite.com/cancel",
		},
		CaptureNow:      true,
		CustomerID:      "cust-d9630db4-0c3d-41b8-bf36-77d7cb0b5dee",
		PaymentMethodID: "pm-d69f5d53-2f9c-43dd-a210-cc9c0cdc28d4",
	}

	createEWalletChargeTokenizedResponse, err := ewallet.CreateEWalletCharge(&createEWalletChargeTokenized)
	if err != nil {
		log.Panic(err)
	}

	getEWalletChargeTokenizedResponse := ewallet.GetEWalletChargeStatusParams{
		ChargeID: createEWalletChargeTokenizedResponse.ID,
	}

	_, err = ewallet.GetEWalletChargeStatus(&getEWalletChargeTokenizedResponse)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("EWallet integration tests done!")
}
