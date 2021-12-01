package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go/payout"
)

func payoutTest() {
	createData := payout.CreateParams{
		ExternalID: time.Now().String(),
		Amount:     10000,
		Email:      "test@example.com",
	}

	resp, err := payout.Create(&createData)
	if err != nil {
		log.Panic(err)
	}

	resp, err = payout.Get(&payout.GetParams{
		ID: resp.ID,
	})
	if err != nil {
		log.Panic(err)
	}

	_, err = payout.Void(&payout.VoidParams{
		ID: resp.ID,
	})
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Payout integration tests done!")
}
