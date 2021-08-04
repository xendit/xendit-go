package main

import (
	"os"
	"sync"

	"github.com/xendit/xendit-go"
)

func main() {
	xendit.Opt.SecretKey = os.Getenv("SECRET_KEY")

	wg := sync.WaitGroup{}
	wg.Add(13)
	go func() {
		balanceTest()
		wg.Done()
	}()
	go func() {
		cardTest()
		wg.Done()
	}()
	go func() {
		cardlesscreditTest()
		wg.Done()
	}()
	go func() {
		disbursementTest()
		wg.Done()
	}()
	go func() {
		ewalletTest()
		wg.Done()
	}()
	go func() {
		invoiceTest()
		wg.Done()
	}()
	go func() {
		payoutTest()
		wg.Done()
	}()
	go func() {
		recurringpaymentTest()
		wg.Done()
	}()
	go func() {
		retailoutletTest()
		wg.Done()
	}()
	go func() {
		virtualaccountTest()
		wg.Done()
	}()
	go func() {
		promotionTest()
		wg.Done()
	}()
	go func() {
		customerTest()
		wg.Done()
	}()
	go func() {
		directDebitTest()
		wg.Done()
	}()

	wg.Wait()
}
