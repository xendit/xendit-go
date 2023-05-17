package xendit_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xendit/xendit-go"
)

func TestEventInvoicePaid_EWallet(t *testing.T) {
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
	expectedPayload := xendit.EventInvoicePaid{
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
		PaymentDetails: xendit.InvoicePaymentDetail{
			Source:    "Ovo",
			ReceiptID: "example",
		},
		PaymentMethodID: "pm-example-123",
	}
	var actualPayload xendit.EventInvoicePaid
	err := json.Unmarshal([]byte(testPayload), &actualPayload)
	assert.NoError(t, err)
	assert.Equal(t, expectedPayload, actualPayload)
}
