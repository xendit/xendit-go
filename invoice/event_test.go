package invoice_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

func TestInvoiceCallback_EWallet(t *testing.T) {
	testPayload := `{
		"id": "12345",
		"fees": [
			{
				"type": "example fee",
				"value": 100
			}
		],
		"amount": 1000,
		"status": "PAID",
		"created": "2023-05-14T07:32:42.646Z",
		"is_high": false,
		"paid_at": "2023-05-14T07:35:14.000Z",
		"updated": "2023-05-14T07:35:15.810Z",
		"user_id": "example",
		"currency": "IDR",
		"payment_id": "qrpy_example",
		"external_id": "example",
		"paid_amount": 1000,
		"merchant_name": "Example Merchant",
		"payment_method": "QR_CODE",
		"payment_channel": "QRIS",
		"payment_details": {
			"source": "Ovo",
			"receipt_id": "example"
		},
		"payment_method_id": "pm-example-123"
	}`
	expectedPayload := invoice.InvoiceCallback{
		ID:             "12345",
		Fees:           []xendit.InvoiceFee{{Type: "example fee", Value: 100}},
		Amount:         1000,
		Status:         "PAID",
		Created:        "2023-05-14T07:32:42.646Z",
		IsHigh:         false,
		PaidAt:         "2023-05-14T07:35:14.000Z",
		Updated:        "2023-05-14T07:35:15.810Z",
		UserID:         "example",
		Currency:       "IDR",
		PaymentID:      "qrpy_example",
		ExternalID:     "example",
		PaidAmount:     1000,
		MerchantName:   "Example Merchant",
		PaymentMethod:  "QR_CODE",
		PaymentChannel: "QRIS",
		PaymentDetails: &xendit.InvoicePaymentDetail{
			Source:    "Ovo",
			ReceiptID: "example",
		},
		PaymentMethodID: "pm-example-123",
	}
	var actualPayload invoice.InvoiceCallback
	err := json.Unmarshal([]byte(testPayload), &actualPayload)
	assert.NoError(t, err)
	assert.Equal(t, expectedPayload, actualPayload)
}

func TestInvoiceCallback_Bank(t *testing.T) {
	testPayload := `{
		"id": "593f4ed1c3d3bb7f39733d83",
		"external_id": "testing-invoice",
		"user_id": "5848fdf860053555135587e7",
		"is_high": false,
		"payment_method": "BANK_TRANSFER",
		"status": "PAID",
		"merchant_name": "Xendit",
		"amount": 2000000,
		"paid_amount": 2000000,
		"bank_code": "TESTBANK",
		"paid_at": "2020-01-14T02:32:50.912Z",
		"payer_email": "test@xendit.co",
		"description": "Invoice callback test",
		"created": "2020-01-13T02:32:49.827Z",
		"updated": "2020-01-13T02:32:50.912Z",
		"currency": "IDR",
		"payment_channel": "TEST",
		"payment_destination": "8458478548758748"
	}`
	expectedPayload := invoice.InvoiceCallback{
		ID:                 "593f4ed1c3d3bb7f39733d83",
		ExternalID:         "testing-invoice",
		UserID:             "5848fdf860053555135587e7",
		IsHigh:             false,
		PaymentMethod:      "BANK_TRANSFER",
		Status:             "PAID",
		MerchantName:       "Xendit",
		Amount:             2000000,
		PaidAmount:         2000000,
		BankCode:           "TESTBANK",
		PaidAt:             "2020-01-14T02:32:50.912Z",
		PayerEmail:         "test@xendit.co",
		Description:        "Invoice callback test",
		Created:            "2020-01-13T02:32:49.827Z",
		Updated:            "2020-01-13T02:32:50.912Z",
		Currency:           "IDR",
		PaymentChannel:     "TEST",
		PaymentDestination: "8458478548758748",
	}
	var actualPayload invoice.InvoiceCallback
	err := json.Unmarshal([]byte(testPayload), &actualPayload)
	assert.NoError(t, err)
	assert.Equal(t, expectedPayload, actualPayload)
}

func TestInvoiceCallback_RetailOutlet(t *testing.T) {
	testPayload := `{
		"id": "593f4ed1c3d3bb7f39733d83",
		"external_id": "testing-invoice",
		"user_id": "5848fdf860053555135587e7",
		"is_high": false,
		"payment_method": "RETAIL_OUTLET",
		"status": "PAID",
		"merchant_name": "Xendit",
		"amount": 2000000,
		"paid_amount": 2000000,
		"paid_at": "2020-01-14T02:32:50.912Z",
		"payer_email": "test@xendit.co",
		"description": "Invoice callback test",
		"created": "2020-01-13T02:32:49.827Z",
		"updated": "2020-01-13T02:32:50.912Z",
		"currency": "IDR",
		"payment_channel": "TESTMART",
		"payment_destination": "TEST815"
	}`
	expectedPayload := invoice.InvoiceCallback{
		ID:                 "593f4ed1c3d3bb7f39733d83",
		ExternalID:         "testing-invoice",
		UserID:             "5848fdf860053555135587e7",
		IsHigh:             false,
		PaymentMethod:      "RETAIL_OUTLET",
		Status:             "PAID",
		MerchantName:       "Xendit",
		Amount:             2000000,
		PaidAmount:         2000000,
		PaidAt:             "2020-01-14T02:32:50.912Z",
		PayerEmail:         "test@xendit.co",
		Description:        "Invoice callback test",
		Created:            "2020-01-13T02:32:49.827Z",
		Updated:            "2020-01-13T02:32:50.912Z",
		Currency:           "IDR",
		PaymentChannel:     "TESTMART",
		PaymentDestination: "TEST815",
	}

	var actualPayload invoice.InvoiceCallback
	err := json.Unmarshal([]byte(testPayload), &actualPayload)
	assert.NoError(t, err)
	assert.Equal(t, expectedPayload, actualPayload)
}

func TestInvoiceCallback_Expired(t *testing.T) {
	testPayload := `{
		"id": "621887f17d9cdaa199d6e787",
		"user_id": "61d3c21692594a88b0dad56b",
		"external_id": "yumin-invoice-1645774832",
		"is_high": false,
		"status": "EXPIRED",
		"merchant_name": "fintech",
		"amount": 1000,
		"created": "2022-02-25T07:40:33.922Z",
		"updated": "2022-02-25T07:42:43.872Z",
		"description": "Invoice Demo #123",
		"currency": "IDR",
		"success_redirect_url": "https://www.example.com",
		"failure_redirect_url": "https://www.example.com"
	}`
	expectedPayload := invoice.InvoiceCallback{
		ID:                 "621887f17d9cdaa199d6e787",
		UserID:             "61d3c21692594a88b0dad56b",
		ExternalID:         "yumin-invoice-1645774832",
		IsHigh:             false,
		Status:             "EXPIRED",
		MerchantName:       "fintech",
		Amount:             1000,
		Created:            "2022-02-25T07:40:33.922Z",
		Updated:            "2022-02-25T07:42:43.872Z",
		Description:        "Invoice Demo #123",
		Currency:           "IDR",
		SuccessRedirectURL: "https://www.example.com",
		FailureRedirectURL: "https://www.example.com",
	}

	var actualPayload invoice.InvoiceCallback
	err := json.Unmarshal([]byte(testPayload), &actualPayload)
	assert.NoError(t, err)
	assert.Equal(t, expectedPayload, actualPayload)
}
