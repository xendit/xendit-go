package account

import (
	"context"

	"github.com/xendit/xendit-go"
)

// Create creates new account
func Create(data *CreateParams) (*xendit.Account, *xendit.Error) {
	return CreateWithContext(context.Background(), data)
}

// CreateWithContext creates new account with context
func CreateWithContext(ctx context.Context, data *CreateParams) (*xendit.Account, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreateWithContext(ctx, data)
}

// Get gets one account
func Get(data *GetParams) (*xendit.Account, *xendit.Error) {
	return GetWithContext(context.Background(), data)
}

// GetWithContext gets one account with context
func GetWithContext(ctx context.Context, data *GetParams) (*xendit.Account, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetWithContext(ctx, data)
}

func getClient() (*Client, *xendit.Error) {
	return &Client{
		Opt:          &xendit.Opt,
		APIRequester: xendit.GetAPIRequester(),
	}, nil
}
