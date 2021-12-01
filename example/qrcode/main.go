package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/qrcode"
)

func main() {
	godotenvErr := godotenv.Load()
	if godotenvErr != nil {
		log.Fatal(godotenvErr)
	}
	xendit.Opt.SecretKey = os.Getenv("SECRET_KEY")

	createData := qrcode.CreateQRCodeParams{
		ExternalID:  time.Now().String(),
		Type:        xendit.DynamicQRCode,
		CallbackURL: "https://httpstat.us/200",
		Amount:      10000,
	}

	resp, err := qrcode.CreateQRCode(&createData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created QR code: %+v\n", resp)

	resp, err = qrcode.GetQRCode(&qrcode.GetQRCodeParams{
		ExternalID: resp.ExternalID,
	})

	fmt.Printf("retrieved QR code: %+v\n", resp)

	if err != nil {
		log.Panic(err)
	}

}
