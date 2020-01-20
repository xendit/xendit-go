package virtualaccount

import (
	"context"

	"github.com/xendit/xendit-go"
)

// CreateFixed creates new invoice
func CreateFixed(data *CreateFixedParams) (*xendit.VirtualAccount, *xendit.Error) {
	return CreateFixedWithContext(context.Background(), data)
}

// CreateFixedWithContext creates new invoice
func CreateFixedWithContext(ctx context.Context, data *CreateFixedParams) (*xendit.VirtualAccount, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreateFixedWithContext(ctx, data)
}

func getClient() (*Client, *xendit.Error) {
	return &Client{
		Opt:          &xendit.Opt,
		APIRequester: xendit.GetAPIRequester(),
	}, nil
}
