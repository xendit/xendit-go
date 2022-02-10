package main

import (
	"os"
	"sync"

	"github.com/xendit/xendit-go"
)

//array of functions to test
var testFunctions = []func(){
	balanceTest,
	cardTest,
	cardlesscreditTest,
	disbursementTest,
	ewalletTest,
	invoiceTest,
	payoutTest,
	recurringpaymentTest,
	retailoutletTest,
	virtualaccountTest,
	promotionTest,
	customerTest,
	directDebitTest,
	qrcodeTest,
	transactionTest,
	accountTest,
}

func main() {
	xendit.Opt.SecretKey = os.Getenv("SECRET_KEY")

	wg := sync.WaitGroup{}
	//Correctly add the wait group, instead of 'magic number'
	wg.Add(len(testFunctions))
	for _, f := range testFunctions {
		//Scope Capture
		go func(f func()) {
			f()
			wg.Done()
		}(f)
	}
	wg.Wait()
}
