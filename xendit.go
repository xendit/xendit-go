package xendit

import (
	invoice "github.com/xendit/xendit-go/invoice"
	option "github.com/xendit/xendit-go/option"
)

// Xendit is ...
type Xendit struct {
	Opt     option.Option
	Invoice invoice.Invoice
}

// New is constructor of Xendit
func New(publicKey string, secretKey string, xenditURL string) Xendit {

	if xenditURL == "" {
		xenditURL = "https://api.xendit.co"
	}

	xendit := Xendit{
		Opt: option.Option{
			PublicKey: publicKey,
			SecretKey: secretKey,
			XenditURL: xenditURL,
		},
	}

	xendit.Invoice = invoice.Invoice{
		Opt: &xendit.Opt,
	}

	return xendit
}
