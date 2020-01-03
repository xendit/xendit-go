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
func (a *API) Init(xdAPIRequester *xendit.XdAPIRequester) {
	if xdAPIRequester == nil {
		xdAPIRequesterObj := xendit.GetXdAPIRequester()
		xdAPIRequester = &xdAPIRequesterObj
	}

	a.Invoice = &invoice.Client{Opt: &a.opt, XdAPIRequester: *xdAPIRequester}
}

// New creates a new Xendit API client
func New(publicKey string, secretKey string, xenditURL string, xdAPIRequester *xendit.XdAPIRequester) *API {
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
	api.Init(xdAPIRequester)

	return &api
}
