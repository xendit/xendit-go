package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/customer"
	"github.com/xendit/xendit-go/directdebit/directdebitpayment"
	"github.com/xendit/xendit-go/directdebit/linkedaccount"
	"github.com/xendit/xendit-go/directdebit/paymentmethod"
)

func directDebitTest() {
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
		ReferenceID: 	time.Now().String(),
		Email:			"tes@tes.com",
		GivenNames:     "Given Names",
		Nationality: 	"ID",
		DateOfBirth: 	"1995-12-30",
		Addresses:		[]xendit.CustomerAddress{customerAddress},
		Metadata:		metadata,
	}

	customer, err := customer.CreateCustomer(&createCustomerData)
	if err != nil {
		log.Panic(err)
	}

	properties := map[string]interface{}{
		"account_mobile_number": "+62818555988",
		"card_last_four": "8888",
		"card_expiry": "06/24",
		"account_email": "test.email@xendit.co",
	}

	customerID := customer.ID

	initializeTokenizationData := linkedaccount.InitializeLinkedAccountTokenizationParams{
		CustomerID:		customerID,
		ChannelCode:	xendit.DC_BRI,
		Properties:		properties,
		Metadata:		metadata,
	}

	resp, err := linkedaccount.InitializeLinkedAccountTokenization(&initializeTokenizationData)
	if err != nil {
		log.Panic(err)
	}

	validateOTPForLinkedAccountData := linkedaccount.ValidateOTPForLinkedAccountParams{
		LinkedAccountTokenID: 	resp.ID,
		OTPCode:				"333000",
	}

	_, err = linkedaccount.ValidateOTPForLinkedAccount(&validateOTPForLinkedAccountData)
	if err != nil {
		log.Panic(err)
	}

	retrieveAccessibleLinkedAccountsData := linkedaccount.RetrieveAccessibleLinkedAccountParams{
		LinkedAccountTokenID:	resp.ID,
	}

	resps, err := linkedaccount.RetrieveAccessibleLinkedAccounts(&retrieveAccessibleLinkedAccountsData)
	if err != nil {
		log.Panic(err)
	}

	unbindLinkedAccountTokenData := linkedaccount.UnbindLinkedAccountTokenParams{
		LinkedAccountTokenID:	resp.ID,
	}

	_, err = linkedaccount.UnbindLinkedAccountToken(&unbindLinkedAccountTokenData)
	if err != nil {
		log.Panic(err)
	}

	properties = map[string]interface{}{
		"id": resps[0].ID,
	}

	createPaymentMethodData := paymentmethod.CreatePaymentMethodParams{
		CustomerID:	customerID,
		Type:		xendit.DEBIT_CARD,
		Properties: properties,
		Metadata:	metadata,	
	}

	resppm, err := paymentmethod.CreatePaymentMethod(&createPaymentMethodData)
	if err != nil {
		log.Panic(err)
	}

	getPaymentMethodsByCustomerIDData := paymentmethod.GetPaymentMethodsByCustomerIDParams{
		CustomerID: customerID,
	}

	_, err = paymentmethod.GetPaymentMethodsByCustomerID(&getPaymentMethodsByCustomerIDData)
	if err != nil {
		log.Panic(err)
	}

	createDirectDebitPaymentData := directdebitpayment.CreateDirectDebitPaymentParams{
		ReferenceID:		"test-direct-debit-ref-0100",
		PaymentMethodID: 	resppm.ID,
		Currency:			"IDR",
		Amount:				15000,
		CallbackURL:		"http://webhook.site/",
		EnableOTP:			true,
		Description:		"test description",
		Basket:				[]xendit.DirectDebitBasketItem{
			{
				ReferenceID:	"basket-product-ref-id",
				Name:			"product-name",
				Category: 		"mechanics",
				Market: 		"ID",
				Price: 			50000,
				Quantity: 		5,
				Type: 			"product type",
				SubCategory:	"product sub category",
				Description: 	"product description",
				URL:			"https://product.url",
			},
		},
		Device:				xendit.DirectDebitDevice{
			ID:			"device-id",
			IPAddress:	"0.0.0.0",
			UserAgent:	"user-agent",
			ADID: 		"ad-id",
			Imei: 		"123a456b789c",
		},
		SuccessRedirectURL: "https://success-redirect.url",
		FailureRedirectURL: "https://failure-redirect.url",
		Metadata: 			metadata,
		IdempotencyKey: 	time.Now().String(),
	}

	respdd, err := directdebitpayment.CreateDirectDebitPayment(&createDirectDebitPaymentData)
	if err != nil {
		log.Panic(err)
	}

	validateOTPForDirectDebitPaymentData := directdebitpayment.ValidateOTPForDirectDebitPaymentParams{
		DirectDebitID:	respdd.ID,
		OTPCode:		"222000",
	}

	_, err = directdebitpayment.ValidateOTPForDirectDebitPayment(&validateOTPForDirectDebitPaymentData)
	if err != nil {
		log.Panic(err)
	}

	getDirectDebitPaymentStatusByIDData := directdebitpayment.GetDirectDebitPaymentStatusByIDParams{
		ID:	respdd.ID,
	}

	_, err = directdebitpayment.GetDirectDebitPaymentStatusByID(&getDirectDebitPaymentStatusByIDData)
	if err != nil {
		log.Panic(err)
	}

	getDirectDebitPaymentStatusByReferenceIDData := directdebitpayment.GetDirectDebitPaymentStatusByReferenceIDParams{
		ReferenceID: respdd.ReferenceID,
	}

	_, err = directdebitpayment.GetDirectDebitPaymentStatusByReferenceID(&getDirectDebitPaymentStatusByReferenceIDData)
	if err != nil {
		log.Panic(err)
	}
	
	fmt.Println("Direct debit integration tests done!")
}
