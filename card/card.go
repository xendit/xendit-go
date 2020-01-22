package card

import (
	"context"

	"github.com/xendit/xendit-go"
)

// CreateAuthorization creates new card authorization
func CreateAuthorization(data *CreateAuthorizationParams) (*xendit.CardCreateAuthorizationResponse, *xendit.Error) {
	return CreateAuthorizationWithContext(context.Background(), data)
}

// CreateAuthorizationWithContext creates new card authorization with context
func CreateAuthorizationWithContext(ctx context.Context, data *CreateAuthorizationParams) (*xendit.CardCreateAuthorizationResponse, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreateAuthorizationWithContext(ctx, data)
}

// ReverseAuthorization reverses a card authorization
func ReverseAuthorization(data *ReverseAuthorizationParams) (*xendit.CardReverseAuthorizationResponse, *xendit.Error) {
	return ReverseAuthorizationWithContext(context.Background(), data)
}

// ReverseAuthorizationWithContext reverses a card authorization with context
func ReverseAuthorizationWithContext(ctx context.Context, data *ReverseAuthorizationParams) (*xendit.CardReverseAuthorizationResponse, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.ReverseAuthorizationWithContext(ctx, data)
}

func getClient() (*Client, *xendit.Error) {
	return &Client{
		Opt:          &xendit.Opt,
		APIRequester: xendit.GetAPIRequester(),
	}, nil
}
