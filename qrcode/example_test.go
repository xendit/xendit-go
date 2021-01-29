package qrcode_test

import (
	"fmt"
	"log"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/qrcode"
)

func ExampleCreateQRCode() {
	xendit.Opt.SecretKey = "examplesecretkey"

	createQRCodeData := qrcode.CreateQRCodeParams{
		ExternalID: "external_id",
		CallbackURL: "http://webhook-site",
		Type: xendit.DynamicQRCode,
		Amount: 10000,
	}

	qrCodeResponse, err := qrcode.CreateQRCode(&createQRCodeData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created QR code: %+v\n", qrCodeResponse)
}

func ExampleGetQRCode() {
	xendit.Opt.SecretKey = "examplesecretkey"

	qrCodeResponse, err := qrcode.GetQRCode(&qrcode.GetQRCodeParams{
		ExternalID: "external_id",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved qr code: %v\n", qrCodeResponse)
}

func ExampleGetQRCodePayments() {
	xendit.Opt.SecretKey = "examplesecretkey"

	getQRCodePaymentsdata := qrcode.GetQRCodePaymentsParams{
		ExternalID: "external_id",
		Limit: 20,
	}

	getQRCodePaymentsRes, err := qrcode.GetQRCodePayments(&getQRCodePaymentsdata)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved QR Code payments: %v\n", getQRCodePaymentsRes)
}
