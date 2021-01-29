package qrcode

import (
	"context"

	"github.com/xendit/xendit-go"
)

// CreateQRCode creates new QR Code.
func CreateQRCode(data *CreateQRCodeParams) (*xendit.QRCode, *xendit.Error) {
	return CreateQRCodeWithContext(context.Background(), data)
}

// CreateQRCodeWithContext creates new QR Code with context.
func CreateQRCodeWithContext(ctx context.Context, data *CreateQRCodeParams) (*xendit.QRCode, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreateQRCodeWithContext(ctx, data)
}

// GetQRCode gets a single QR Code.
func GetQRCode(data *GetQRCodeParams) (*xendit.QRCode, *xendit.Error) {
	return GetQRCodeWithContext(context.Background(), data)
}

// GetQRCodeWithContext gets a single QR Code with context.
func GetQRCodeWithContext(ctx context.Context, data *GetQRCodeParams) (*xendit.QRCode, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetQRCodeWithContext(ctx, data)
}

// GetQRCodePayments gets a list of QR Code payments.
func GetQRCodePayments(data *GetQRCodePaymentsParams) ([]xendit.QRCodePayments, *xendit.Error) {
	return GetQRCodePaymentsWithContext(context.Background(), data)
}

// GetQRCodePaymentsWithContext gets a list of QR Code payments with context.
func GetQRCodePaymentsWithContext(ctx context.Context, data *GetQRCodePaymentsParams) ([]xendit.QRCodePayments, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetQRCodePaymentsWithContext(ctx, data)
}

func getClient() (*Client, *xendit.Error) {
	return &Client{
		Opt:          &xendit.Opt,
		APIRequester: xendit.GetAPIRequester(),
	}, nil
}
