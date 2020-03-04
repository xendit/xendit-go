package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go/virtualaccount"
)

func virtualaccountTest() {
	availableBanks, err := virtualaccount.GetAvailableBanks()
	if err != nil {
		log.Fatal(err)
	}

	createFixedVAData := virtualaccount.CreateFixedVAParams{
		ExternalID: "va-" + time.Now().String(),
		BankCode:   availableBanks[0].Code,
		Name:       "Michael Jackson",
	}
	resp, err := virtualaccount.CreateFixedVA(&createFixedVAData)
	if err != nil {
		log.Fatal(err)
	}

	getFixedVAData := virtualaccount.GetFixedVAParams{
		ID: resp.ID,
	}
	resp, err = virtualaccount.GetFixedVA(&getFixedVAData)
	if err != nil {
		log.Fatal(err)
	}

	expirationDate := time.Now().AddDate(0, 0, 1)
	updateFixedVAData := virtualaccount.UpdateFixedVAParams{
		ID:             resp.ID,
		ExpirationDate: &expirationDate,
	}
	_, err = virtualaccount.UpdateFixedVA(&updateFixedVAData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Virtual Account integration tests done!")
}
