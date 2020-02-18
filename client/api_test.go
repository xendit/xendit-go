package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xendit/xendit-go"
)

func TestAPIInit(t *testing.T) {
	api := API{
		opt: xendit.Option{
			SecretKey: "sk_123",
		},
	}
	api.init()

	assert.Equal(t, *api.Invoice.Opt, api.opt)
}

func TestAPINew(t *testing.T) {
	api := New("sk_123")

	assert.Equal(t, "sk_123", api.opt.SecretKey)
	assert.Equal(t, "https://api.xendit.co", api.opt.XenditURL)
}
