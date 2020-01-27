package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go/payout"

	"github.com/xendit/xendit-go"
)

func main() {
	xendit.Opt.SecretKey = "xnd_development_REt02KJzkM6AootfKnDrMw1Sse4LlzEDHfKzXoBocqIEiH4bqjHUJXbl6Cfaab"

	createData := payout.CreateParams{
		ExternalID: "payout-" + time.Now().String(),
		Amount:     200000,
	}

	resp, err := payout.Create(&createData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created payout: %+v\n", resp)

	resp, err = payout.Get(&payout.GetParams{
		ID: resp.ID,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retrieved payout: %+v\n", resp)

	resp, err = payout.Void(&payout.VoidParams{
		ID: resp.ID,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("voidd payout: %+v\n", resp)
}
