package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/directdebit/directdebitpayment"
	"github.com/xendit/xendit-go/directdebit/linkedaccount"
	"github.com/xendit/xendit-go/directdebit/paymentmethod"
)

func main() {
	godotenvErr := godotenv.Load()
	if godotenvErr != nil {
		log.Fatal(godotenvErr)
	}
	xendit.Opt.SecretKey = os.Getenv("SECRET_KEY")

	properties := map[string]interface{}{
		"account_mobile_number": "+62818555988",
		"card_last_four":        "8888",
		"card_expiry":           "06/24",
		"account_email":         "test.email@xendit.co",
	}

	metadata := map[string]interface{}{
		"meta": "data",
	}

	customerID := "74965fb0-0446-47a9-a7d7-0db06b70af98"

	initializeTokenizationData := linkedaccount.InitializeLinkedAccountTokenizationParams{
		CustomerID:  customerID,
		ChannelCode: xendit.DC_BRI,
		Properties:  properties,
		Metadata:    metadata,
	}

	initializedLinkedAccount, err := linkedaccount.InitializeLinkedAccountTokenization(&initializeTokenizationData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("initialized linked account tokenization: %+v\n\n", initializedLinkedAccount)

	validateOTPForLinkedAccountData := linkedaccount.ValidateOTPForLinkedAccountParams{
		LinkedAccountTokenID: initializedLinkedAccount.ID,
		OTPCode:              "333000",
	}

	validatedLinkedAccount, err := linkedaccount.ValidateOTPForLinkedAccount(&validateOTPForLinkedAccountData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("validated linked account: %+v\n\n", validatedLinkedAccount)

	retrieveAccessibleLinkedAccountsData := linkedaccount.RetrieveAccessibleLinkedAccountParams{
		LinkedAccountTokenID: initializedLinkedAccount.ID,
	}

	accessibleLinkedAccounts, err := linkedaccount.RetrieveAccessibleLinkedAccounts(&retrieveAccessibleLinkedAccountsData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retrieved accessible linked accounts: %+v\n\n", accessibleLinkedAccounts)

	unbindLinkedAccountTokenData := linkedaccount.UnbindLinkedAccountTokenParams{
		LinkedAccountTokenID: initializedLinkedAccount.ID,
	}

	unbindedLinkedAccount, err := linkedaccount.UnbindLinkedAccountToken(&unbindLinkedAccountTokenData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("unbinded linked account: %+v\n\n", unbindedLinkedAccount)

	properties = map[string]interface{}{
		"id": accessibleLinkedAccounts[0].ID,
	}

	createPaymentMethodData := paymentmethod.CreatePaymentMethodParams{
		CustomerID: customerID,
		Type:       xendit.DEBIT_CARD,
		Properties: properties,
		Metadata:   metadata,
	}

	paymentMethod, err := paymentmethod.CreatePaymentMethod(&createPaymentMethodData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created payment method: %+v\n\n", paymentMethod)

	getPaymentMethodsByCustomerIDData := paymentmethod.GetPaymentMethodsByCustomerIDParams{
		CustomerID: customerID,
	}

	paymentMethods, err := paymentmethod.GetPaymentMethodsByCustomerID(&getPaymentMethodsByCustomerIDData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retrieved payment methods: %+v\n\n", paymentMethods)

	createDirectDebitPaymentData := directdebitpayment.CreateDirectDebitPaymentParams{
		ReferenceID:     "test-direct-debit-ref-0100",
		PaymentMethodID: paymentMethod.ID,
		Currency:        "IDR",
		Amount:          15000,
		CallbackURL:     "http://webhook.site/",
		EnableOTP:       true,
		Description:     "test description",
		Basket: []xendit.DirectDebitBasketItem{
			{
				ReferenceID: "basket-product-ref-id",
				Name:        "product-name",
				Category:    "mechanics",
				Market:      "ID",
				Price:       50000,
				Quantity:    5,
				Type:        "product type",
				SubCategory: "product sub category",
				Description: "product description",
				URL:         "https://product.url",
			},
		},
		Device: xendit.DirectDebitDevice{
			ID:        "device-id",
			IPAddress: "0.0.0.0",
			UserAgent: "user-agent",
			ADID:      "ad-id",
			Imei:      "123a456b789c",
		},
		SuccessRedirectURL: "https://success-redirect.url",
		FailureRedirectURL: "https://failure-redirect.url",
		Metadata:           metadata,
		IdempotencyKey:     time.Now().String(),
	}

	directDebitPayment, err := directdebitpayment.CreateDirectDebitPayment(&createDirectDebitPaymentData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created direct debit payment: %+v\n\n", directDebitPayment)

	validateOTPForDirectDebitPaymentData := directdebitpayment.ValidateOTPForDirectDebitPaymentParams{
		DirectDebitID: directDebitPayment.ID,
		OTPCode:       "222000",
	}

	directDebitPayment, err = directdebitpayment.ValidateOTPForDirectDebitPayment(&validateOTPForDirectDebitPaymentData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("validated direct debit payment: %+v\n\n", directDebitPayment)

	getDirectDebitPaymentStatusByIDData := directdebitpayment.GetDirectDebitPaymentStatusByIDParams{
		ID: directDebitPayment.ID,
	}

	directDebitPayment, err = directdebitpayment.GetDirectDebitPaymentStatusByID(&getDirectDebitPaymentStatusByIDData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retrieved direct debit payment: %+v\n\n", directDebitPayment)

	getDirectDebitPaymentStatusByReferenceIDData := directdebitpayment.GetDirectDebitPaymentStatusByReferenceIDParams{
		ReferenceID: directDebitPayment.ReferenceID,
	}

	directDebitPayments, err := directdebitpayment.GetDirectDebitPaymentStatusByReferenceID(&getDirectDebitPaymentStatusByReferenceIDData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retrieved direct debit payments: %+v\n\n", directDebitPayments)
}
