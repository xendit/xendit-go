package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/ewallet"
)

func ewalletTest() {
	createPaymentData := ewallet.CreatePaymentParams{
		ExternalID:  "ovo-" + time.Now().String(),
		Amount:      3200,
		Phone:       "08123123123",
		EWalletType: xendit.EWalletTypeOVO,
	}

	resp, err := ewallet.CreatePayment(&createPaymentData)
	if err != nil {
		log.Panic(err)
	}
	getPaymentStatusData := ewallet.GetPaymentStatusParams{
		ExternalID:  resp.ExternalID,
		EWalletType: resp.EWalletType,
	}

	_, err = ewallet.GetPaymentStatus(&getPaymentStatusData)
	if err != nil {
		log.Panic(err)
	}

	ewalletBasketItem := xendit.EWalletBasketItem{
		ReferenceID: "basket-product-ref-id",
		Name:        "product name",
		Category:    "mechanics",
		Currency:    "IDR",
		Price:       50000,
		Quantity:    5,
		Type:        "wht",
		SubCategory: "evr",
		Metadata: map[string]interface{}{
			"meta": "data",
		},
	}

	createEWalletChargeData := ewallet.CreateEWalletChargeParams{
		ReferenceID:    time.Now().String(),
		Currency:       "IDR",
		Amount:         100,
		CheckoutMethod: "ONE_TIME_PAYMENT",
		ChannelCode:    "ID_OVO",
		ChannelProperties: map[string]string{
			"mobile_number": "+6281234567890",
		},
		Basket: []xendit.EWalletBasketItem{
			ewalletBasketItem,
		},
		Metadata: map[string]interface{}{
			"meta2": "data2",
		},
	}

	charge, chargeErr := ewallet.CreateEWalletCharge(&createEWalletChargeData)
	if chargeErr != nil {
		log.Panic(chargeErr)
	}

	getEWalletChargeStatusData := ewallet.GetEWalletChargeStatusParams{
		ChargeID: charge.ID,
	}

	_, chargeErr = ewallet.GetEWalletChargeStatus(&getEWalletChargeStatusData)
	if chargeErr != nil {
		log.Panic(chargeErr)
	}

	createEWalletChargeAdditionalFields := ewallet.CreateEWalletChargeParams{
		ReferenceID:    "test-reference-id",
		Currency:       "IDR",
		Amount:         10000,
		CheckoutMethod: "TOKENIZED_PAYMENT",
		ChannelCode:    "ID_OVO",
		ChannelProperties: map[string]string{
			"success_redirect_url": "https://yourwebsite.com/order/123",
			"failure_redirect_url": "https://yourwebsite.com/failure",
			"cancel_redirect_url":  "https://yourwebsite.com/cancel",
		},
		CaptureNow:      true,
		CustomerID:      "f3925450-5c54-4777-98c1-fcf22b0d1e1c",
		PaymentMethodID: "pm-f3925450-5c54-4777-98c1-fcf22b0d1e1c",
		Customer: xendit.EwalletCustomer{
			ReferenceId:            func(i string) *string { return &i }("sample_customer_reference_id"),
			GivenNames:             func(i string) *string { return &i }("sample_given_name"),
			Surname:                func(i string) *string { return &i }("sample_surname"),
			Email:                  func(i string) *string { return &i }("sample_email"),
			MobileNumber:           func(i string) *string { return &i }("sample_mobile_number"),
			DomicileOfRegistration: func(i string) *string { return &i }("sample_domicile"),
			DateOfRegistration:     func(i string) *string { return &i }("sample_dor"),
		},
		ShippingInformation: xendit.ShippingInformation{
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

	createChargeResponse, err := ewallet.CreateEWalletCharge(&createEWalletChargeAdditionalFields)
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

	fmt.Println("EWallet integration tests done!")
}
