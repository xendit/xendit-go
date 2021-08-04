package directdebitpayment_test

import (
	"fmt"
	"log"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/directdebit/directdebitpayment"
)

func ExampleCreateDirectDebitPayment() {
	xendit.Opt.SecretKey = "examplesecretkey"

	metadata := map[string]interface{}{
		"meta": "data",
	}

	data := directdebitpayment.CreateDirectDebitPaymentParams{
		IdempotencyKey:  "idem-key",
		ReferenceID:     "test-ref-id",
		PaymentMethodID: "test-pm-id",
		Currency:        "IDR",
		Amount:          15000,
		CallbackURL:     "http://webhook.site",
		EnableOTP:       true,
		Description:     "Test description",
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
	}

	resp, err := directdebitpayment.CreateDirectDebitPayment(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created direct debit payment: %+v\n", resp)
}

func ExampleValidateOTPForDirectDebitPayment() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := directdebitpayment.ValidateOTPForDirectDebitPaymentParams{
		DirectDebitID: "test-ddpy-id",
		OTPCode:       "333000",
	}

	resp, err := directdebitpayment.ValidateOTPForDirectDebitPayment(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("validated direct debit payment: %+v\n", resp)
}

func ExampleGetDirectDebitPaymentStatusByID() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := directdebitpayment.GetDirectDebitPaymentStatusByIDParams{
		ID: "test-ddpy-id",
	}

	resp, err := directdebitpayment.GetDirectDebitPaymentStatusByID(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved direct debit payment: %+v\n", resp)
}

func ExampleGetDirectDebitPaymentStatusByReferenceID() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := directdebitpayment.GetDirectDebitPaymentStatusByReferenceIDParams{
		ReferenceID: "test-ddpy-ref-id",
	}

	resp, err := directdebitpayment.GetDirectDebitPaymentStatusByReferenceID(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved direct debit payments: %+v\n", resp)
}
