package transaction_test

import (
	"fmt"
	"log"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/transaction"
)

func ExampleGetTransaction() {
	xendit.Opt.SecretKey = "examplesecretkey"

	getTransaction := transaction.GetTransactionParams{
		TransactionID: "txn_13dd178d-41fa-40b7-8fd3-f83675d6f413",
	}

	resp, err := transaction.GetTransaction(&getTransaction)
	if err != nil {
		log.Fatal(err.ErrorCode)
	}

	fmt.Printf("retrieved transaction %+v\n", resp)
}

func ExampleGetListTransaction() {
	xendit.Opt.SecretKey = "examplesecretkey"

	getListTransaction := transaction.GetListTransactionParams{
		Limit: 10,
	}

	resp, err := transaction.GetListTransaction(&getListTransaction)
	if err != nil {
		log.Fatal(err.ErrorCode)
	}

	fmt.Printf("retrieved list transaction %+v\n", resp)
}
