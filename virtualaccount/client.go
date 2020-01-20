package virtualaccount

import (
	"context"
	"fmt"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/utils/validator"
)

// Client is the client used to invoke invoice API.
type Client struct {
	Opt          *xendit.Option
	APIRequester xendit.APIRequester
}

// CreateFixed creates new fixed virtual account
func (c Client) CreateFixed(data *CreateFixedParams) (*xendit.VirtualAccount, *xendit.Error) {
	return c.CreateFixedWithContext(context.Background(), data)
}

// CreateFixedWithContext creates new fixed virtual account with context
func (c *Client) CreateFixedWithContext(ctx context.Context, data *CreateFixedParams) (*xendit.VirtualAccount, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.VirtualAccount{}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/callback_virtual_accounts", c.Opt.XenditURL),
		c.Opt.SecretKey,
		data,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}
