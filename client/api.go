// Package client provides a Xendit client for invoking APIs across all products
package client

import (
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/balance"
	"github.com/xendit/xendit-go/card"
	"github.com/xendit/xendit-go/ewallet"
	"github.com/xendit/xendit-go/invoice"
	"github.com/xendit/xendit-go/payout"
	"github.com/xendit/xendit-go/recurringpayment"
	"github.com/xendit/xendit-go/retailoutlet"
	"github.com/xendit/xendit-go/virtualaccount"
)

// API is the Xendit client which contains all products
type API struct {
	opt            xendit.Option
	Invoice        *invoice.Client
	EWallet        *ewallet.Client
	Balance        *balance.Client
	RetailOutlet   *retailoutlet.Client
	VirtualAccount *virtualaccount.Client
	Card           *card.Client
	Payout         *payout.Client
	RecurringPayment *recurringpayment.Client

}

// Init initiates all the products of the API client
func (a *API) Init(apiRequester *xendit.APIRequester) {
	if apiRequester == nil {
		apiRequesterObj := xendit.GetAPIRequester()
		apiRequester = &apiRequesterObj
	}

	a.Invoice = &invoice.Client{Opt: &a.opt, APIRequester: *apiRequester}
	a.EWallet = &ewallet.Client{Opt: &a.opt, APIRequester: *apiRequester}
	a.Balance = &balance.Client{Opt: &a.opt, APIRequester: *apiRequester}
	a.RetailOutlet = &retailoutlet.Client{Opt: &a.opt, APIRequester: *apiRequester}
	a.VirtualAccount = &virtualaccount.Client{Opt: &a.opt, APIRequester: *apiRequester}
	a.Card = &card.Client{Opt: &a.opt, APIRequester: *apiRequester}
	a.Payout = &payout.Client{Opt: &a.opt, APIRequester: *apiRequester}
	a.RecurringPayment = &recurringpayment.Client{Opt: &a.opt, APIRequester: *apiRequester}
}

// New creates a new Xendit API client
func New(secretKey string, xenditURL string, apiRequester *xendit.APIRequester) *API {
	if xenditURL == "" {
		xenditURL = "https://api.xendit.co"
	}

	api := API{
		opt: xendit.Option{
			SecretKey: secretKey,
			XenditURL: xenditURL,
		},
	}
	api.Init(apiRequester)

	return &api
}
