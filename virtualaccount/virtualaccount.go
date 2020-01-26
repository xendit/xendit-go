package virtualaccount

import (
	"context"

	"github.com/xendit/xendit-go"
)

// CreateFixed creates new fixed virtual account
func CreateFixed(data *CreateFixedParams) (*xendit.VirtualAccount, *xendit.Error) {
	return CreateFixedWithContext(context.Background(), data)
}

// CreateFixedWithContext creates new fixed virtual account with context
func CreateFixedWithContext(ctx context.Context, data *CreateFixedParams) (*xendit.VirtualAccount, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreateFixedWithContext(ctx, data)
}

// GetFixed gets a fixed virtual account
func GetFixed(data *GetFixedParams) (*xendit.VirtualAccount, *xendit.Error) {
	return GetFixedWithContext(context.Background(), data)
}

// GetFixedWithContext gets a fixed virtual account with context
func GetFixedWithContext(ctx context.Context, data *GetFixedParams) (*xendit.VirtualAccount, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetFixedWithContext(ctx, data)
}

// UpdateFixed updates a fixed virtual account
func UpdateFixed(data *UpdateFixedParams) (*xendit.VirtualAccount, *xendit.Error) {
	return UpdateFixedWithContext(context.Background(), data)
}

// UpdateFixedWithContext updates a fixed virtual account with context
func UpdateFixedWithContext(ctx context.Context, data *UpdateFixedParams) (*xendit.VirtualAccount, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.UpdateFixedWithContext(ctx, data)
}

// GetAvailableBanks gets available virtual account banks
func GetAvailableBanks() ([]xendit.VirtualAccountBank, *xendit.Error) {
	return GetAvailableBanksWithContext(context.Background())
}

// GetAvailableBanksWithContext gets available virtual account banks with context
func GetAvailableBanksWithContext(ctx context.Context) ([]xendit.VirtualAccountBank, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetAvailableBanksWithContext(ctx)
}

// GetPayment gets one fixed virtual account payment
func GetPayment(data *GetPaymentParams) (*xendit.VirtualAccountPayment, *xendit.Error) {
	return GetPaymentWithContext(context.Background(), data)
}

// GetPaymentWithContext gets one fixed virtual account payment with context
func GetPaymentWithContext(ctx context.Context, data *GetPaymentParams) (*xendit.VirtualAccountPayment, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetPaymentWithContext(ctx, data)
}

func getClient() (*Client, *xendit.Error) {
	return &Client{
		Opt:          &xendit.Opt,
		APIRequester: xendit.GetAPIRequester(),
	}, nil
}
