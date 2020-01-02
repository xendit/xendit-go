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
	api.Init(nil)

	assert.Equal(t, *api.Invoice.Opt, api.opt)
}

func TestAPINew(t *testing.T) {
	api, err := New("", "sk_123", "", nil)

	assert.Nil(t, err)
	assert.Equal(t, "sk_123", api.opt.SecretKey)
	assert.Equal(t, "https://api.xendit.co", api.opt.XenditURL)
}

func TestAPINewWithEmptySecretKey(t *testing.T) {
	api, err := New("", "", "", nil)

	assert.NotNil(t, err)
	assert.Nil(t, api)
}
