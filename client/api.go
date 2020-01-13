package client

import (
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

// API is the Xendit client which contains all products
type API struct {
	opt     xendit.Option
	Invoice *invoice.Client
}

// Init initiate all the products of the API client
func (a *API) Init(apiRequester *xendit.APIRequester) {
	if apiRequester == nil {
		apiRequesterObj := xendit.GetAPIRequester()
		apiRequester = &apiRequesterObj
	}

	a.Invoice = &invoice.Client{Opt: &a.opt, APIRequester: *apiRequester}
}

// New creates a new Xendit API client
func New(publicKey string, secretKey string, xenditURL string, apiRequester *xendit.APIRequester) *API {
	if xenditURL == "" {
		xenditURL = "https://api.xendit.co"
	}

	api := API{
		opt: xendit.Option{
			PublicKey: publicKey,
			SecretKey: secretKey,
			XenditURL: xenditURL,
		},
	}
	api.Init(apiRequester)

	return &api
}
