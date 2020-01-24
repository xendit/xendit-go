// Package client provides a Xendit client for invoking APIs across all products
package client

import (
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/ewallet"
	"github.com/xendit/xendit-go/invoice"
	"github.com/xendit/xendit-go/retailoutlet"
)

// API is the Xendit client which contains all products
type API struct {
	opt          xendit.Option
	Invoice      *invoice.Client
	EWallet      *ewallet.Client
	RetailOutlet *retailoutlet.Client
}

// Init initiates all the products of the API client
func (a *API) Init(apiRequester *xendit.APIRequester) {
	if apiRequester == nil {
		apiRequesterObj := xendit.GetAPIRequester()
		apiRequester = &apiRequesterObj
	}

	a.Invoice = &invoice.Client{Opt: &a.opt, APIRequester: *apiRequester}
	a.EWallet = &ewallet.Client{Opt: &a.opt, APIRequester: *apiRequester}
	a.RetailOutlet = &retailoutlet.Client{Opt: &a.opt, APIRequester: *apiRequester}
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
