package main

import (
	"fmt"
	"log"

	"github.com/xendit/xendit-go/transaction"
)

func transactionTest() {

	respList, err := transaction.GetListTransaction(&transaction.GetListTransactionParams{})
	if err != nil {
		log.Panic(err)
	}

	_, err = transaction.GetTransaction(&transaction.GetTransactionParams{
		TransactionID: respList.Data[0].ID,
	})

	if err != nil {
		log.Panic(err)
	}

	fmt.Println("transaction integration tests done!")
}
