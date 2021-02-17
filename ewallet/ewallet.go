package ewallet

import (
	"context"

	"github.com/xendit/xendit-go"
)

// CreatePayment creates new payment
func CreatePayment(data *CreatePaymentParams) (*xendit.EWallet, *xendit.Error) {
	return CreatePaymentWithContext(context.Background(), data)
}

// CreatePaymentWithContext creates new payment
func CreatePaymentWithContext(ctx context.Context, data *CreatePaymentParams) (*xendit.EWallet, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreatePaymentWithContext(ctx, data)
}

// GetPaymentStatus gets one payment with its status
func GetPaymentStatus(data *GetPaymentStatusParams) (*xendit.EWallet, *xendit.Error) {
	return GetPaymentStatusWithContext(context.Background(), data)
}

// GetPaymentStatusWithContext gets one payment with its status
func GetPaymentStatusWithContext(ctx context.Context, data *GetPaymentStatusParams) (*xendit.EWallet, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetPaymentStatusWithContext(ctx, data)
}

// CreateEWalletCharge creates new e-wallet charge
func CreateEWalletCharge(data *CreateEWalletChargeParams) (*xendit.EWalletCharge, *xendit.Error) {
	return CreateEWalletChargeWithContext(context.Background(), data)
}

// CreateEWalletChargeWithContext creates new e-wallet charge
func CreateEWalletChargeWithContext(ctx context.Context, data *CreateEWalletChargeParams) (*xendit.EWalletCharge, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreateEWalletChargeWithContext(ctx, data)
}

// GetEWalletChargeStatus gets one e-wallet charge with its status
func GetEWalletChargeStatus(data *GetEWalletChargeStatusParams) (*xendit.EWalletCharge, *xendit.Error) {
	return GetEWalletChargeStatusWithContext(context.Background(), data)
}

// GetEWalletChargeStatusWithContext gets one e-wallet with its status
func GetEWalletChargeStatusWithContext(ctx context.Context, data *GetEWalletChargeStatusParams) (*xendit.EWalletCharge, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetEWalletChargeStatusWithContext(ctx, data)
}

func getClient() (*Client, *xendit.Error) {
	return &Client{
		Opt:          &xendit.Opt,
		APIRequester: xendit.GetAPIRequester(),
	}, nil
}
