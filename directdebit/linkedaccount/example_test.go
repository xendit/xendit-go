package linkedaccount_test

import (
	"fmt"
	"log"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/directdebit/linkedaccount"
)

func ExampleInitializeLinkedAccountTokenization() {
	xendit.Opt.SecretKey = "examplesecretkey"

	properties := map[string]interface{}{
		"account_mobile_number": "+62818555988",
		"card_last_four": "8888",
		"card_expiry": "06/24",
		"account_email": "test.email@xendit.co",
	}

	metadata := map[string]interface{}{
		"meta": "data",
	}

	data := linkedaccount.InitializeLinkedAccountTokenizationParams{
		CustomerID: 	"test-cust-id",
		ChannelCode:	xendit.DC_BRI,
		Properties: 	properties,
		Metadata:		metadata,
	}

	resp, err := linkedaccount.InitializeLinkedAccountTokenization(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("initialized linked account tokenization: %+v\n", resp)
}

func ExampleValidateOTPForLinkedAccount() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := linkedaccount.ValidateOTPForLinkedAccountParams{
		LinkedAccountTokenID: 	"test-lat-id",
		OTPCode:				"333000",
	}

	resp, err := linkedaccount.ValidateOTPForLinkedAccount(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("validated linked account: %+v\n", resp)
}

func ExampleRetrieveAccessibleLinkedAccounts() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := linkedaccount.RetrieveAccessibleLinkedAccountParams{
		LinkedAccountTokenID:	"test-lat-id",
	}

	resp, err := linkedaccount.RetrieveAccessibleLinkedAccounts(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved accessible linked accounts: %+v\n", resp)
}

func ExampleUnbindLinkedAccountToken() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := linkedaccount.UnbindLinkedAccountTokenParams{
		LinkedAccountTokenID:	"test-lat-id",
	}

	resp, err := linkedaccount.UnbindLinkedAccountToken(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("unbinded linked account token: %+v\n", resp)
}
