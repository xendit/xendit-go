package account

import (
	"context"
	"fmt"
	"net/http"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/utils/validator"
)

// Client is the client used to invoke account API.
type Client struct {
	Opt          *xendit.Option
	APIRequester xendit.APIRequester
}

// Create creates new invoice
func (c *Client) Create(data *CreateParams) (*xendit.Account, *xendit.Error) {
	return c.CreateWithContext(context.Background(), data)
}

// CreateWithContext creates new account with context
func (c *Client) CreateWithContext(ctx context.Context, data *CreateParams) (*xendit.Account, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.Account{}
	header := http.Header{}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/v2/accounts", c.Opt.XenditURL),
		c.Opt.SecretKey,
		header,
		data,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Get gets one account
func (c *Client) Get(data *GetParams) (*xendit.Account, *xendit.Error) {
	return c.GetWithContext(context.Background(), data)
}

// GetWithContext gets one account with context
func (c *Client) GetWithContext(ctx context.Context, data *GetParams) (*xendit.Account, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.Account{}
	header := http.Header{}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/v2/accounts/%s", c.Opt.XenditURL, data.ID),
		c.Opt.SecretKey,
		header,
		nil,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}
