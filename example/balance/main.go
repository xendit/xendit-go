package main

import (
	"fmt"
	"log"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/balance"
)

func main() {
	xendit.Opt.SecretKey = "xnd_development_REt02KJzkM6AootfKnDrMw1Sse4LlzEDHfKzXoBocqIEiH4bqjHUJXbl6Cfaab"

	getData := balance.GetParams{
		AccountType: "CASH",
	}

	resp, err := balance.Get(&getData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("balance: %+v\n", resp)
}
