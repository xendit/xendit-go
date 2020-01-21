package balance

import (
	"context"
	"fmt"

	"github.com/xendit/xendit-go"
)

// Client is the client used to invoke balance API.
type Client struct {
	Opt          *xendit.Option
	APIRequester xendit.APIRequester
}

// Get gets one balance
func (c *Client) Get(data *GetParams) (*xendit.Balance, *xendit.Error) {
	return c.GetWithContext(context.Background(), data)
}

// GetWithContext gets one balance with context
func (c *Client) GetWithContext(ctx context.Context, data *GetParams) (*xendit.Balance, *xendit.Error) {
	var queryString string

	if data != nil {
		queryString = data.QueryString()
	}

	response := &xendit.Balance{}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/balance?%s", c.Opt.XenditURL, queryString),
		c.Opt.SecretKey,
		nil,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}
