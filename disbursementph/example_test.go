package disbursementph_test

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/disbursementph"
)

func ExampleCreate() {
	xendit.Opt.SecretKey = "examplesecretkey"

	createData := disbursementph.CreateParams{
		IdempotencyKey: "disbursement" + time.Now().String(),
		ReferenceID:    "disbursement-reference",
		ChannelCode:    "PH_BDO",
		AccountName:    "Michael Jackson",
		AccountNumber:  "1234567890",
		Description:    "Disbursement from Go",
		Amount:         200000,
		Beneficiary: xendit.Beneficiary{
			Type:         "INDIVIDUAL",
			GivenNames:   "John",
			MiddleName:   "Michel",
			Surname:      "Doe",
			BusinessName: "Example Business",
			StreetLine1:  "Sultan Street",
			StreetLine2:  "BCG",
			City:         "Manila",
			Province:     "Manila",
			State:        "Manila",
			Country:      "Philippines",
			ZipCode:      "12345",
			MobileNumber: "9876543210",
			PhoneNumber:  "987654321",
			Email:        "example@test.com",
		},
		ReceiptNotification: xendit.ReceiptNotification{
			EmailTo:  []string{"example@test.com"},
			EmailCC:  []string{"example@test.com"},
			EmailBCC: []string{"example@test.com"},
		},
		Metadata: map[string]interface{}{
			"meta": "data",
		},
	}

	resp, err := disbursementph.Create(&createData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created disbursement: %+v\n", resp)
}

func ExampleGetByID() {
	xendit.Opt.SecretKey = "examplesecretkey"

	getByIDData := disbursementph.GetByIDParams{
		DisbursementID: "123",
	}

	resp, err := disbursementph.GetByID(&getByIDData)
	if err != nil {
		log.Fatal(err.ErrorCode, err.Message, err.GetStatus())
	}

	fmt.Printf("retrieved disbursement: %+v\n", resp)
}

func ExampleGetByReferenceID() {
	xendit.Opt.SecretKey = "examplesecretkey"

	getByReferenceIDData := disbursementph.GetByReferenceIDParams{
		ReferenceID: "disbursement-reference-id",
	}

	resps, err := disbursementph.GetByReferenceID(&getByReferenceIDData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved disbursements: %+v\n", resps)
}

func ExampleGetDisbursementChannels() {
	xendit.Opt.SecretKey = "examplesecretkey"

	disbursementChannels, err := disbursementph.GetDisbursementChannels()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("disbursement channels: %+v\n", disbursementChannels)
}

func ExampleGetDisbursementChannelsByChannelCategory() {
	xendit.Opt.SecretKey = "examplesecretkey"

	getByChannelCategoryData := disbursementph.GetByChannelCategoryParams{
		ChannelCategory: "BANK",
	}

	disbursementChannels, err := disbursementph.GetDisbursementChannelsByChannelCategory(&getByChannelCategoryData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("disbursement channels: %+v\n", disbursementChannels)
}

func ExampleGetDisbursementChannelsByChannelCode() {
	xendit.Opt.SecretKey = "examplesecretkey"

	getByChannelCodeData := disbursementph.GetByChannelCodeParams{
		ChannelCode: "BANK",
	}

	disbursementChannels, err := disbursementph.GetDisbursementChannelsByChannelCode(&getByChannelCodeData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("disbursement channels: %+v\n", disbursementChannels)
}
