package main

import (
	"fmt"
	"log"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/virtualaccount"
)

func main() {
	xendit.Opt.SecretKey = "xnd_development_REt02KJzkM6AootfKnDrMw1Sse4LlzEDHfKzXoBocqIEiH4bqjHUJXbl6Cfaab"

	createFixedData := virtualaccount.CreateFixedParams{
		ExternalID: "example-externalid",
		BankCode:   "MANDIRI",
		Name:       "Michael Jackson",
		// ExpirationDate: time.Now().Local().Add(60 * time.Hour),
	}

	resp, err := virtualaccount.CreateFixed(&createFixedData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created fixed va: %+v\n", resp)
}
