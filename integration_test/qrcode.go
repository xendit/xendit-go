package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/qrcode"
)

func qrcodeTest() {
	createData := qrcode.CreateQRCodeParams{
		ExternalID:  time.Now().String(),
		Type:        xendit.DynamicQRCode,
		CallbackURL: "https://httpstat.us/200",
		Amount:      10000,
	}

	resp, err := qrcode.CreateQRCode(&createData)
	if err != nil {
		log.Panic(err)
	}

	_, err = qrcode.GetQRCode(&qrcode.GetQRCodeParams{
		ExternalID: resp.ExternalID,
	})

	if err != nil {
		log.Panic(err)
	}

	fmt.Println("QR Code integration tests done!")
}
