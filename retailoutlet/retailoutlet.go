package retailoutlet

import (
	"context"

	"github.com/xendit/xendit-go"
)

// CreateFixedPaymentCode creates new retail outlet fixed payment code
func CreateFixedPaymentCode(data *CreateFixedPaymentCodeParams) (*xendit.RetailOutlet, *xendit.Error) {
	return CreateFixedPaymentCodeWithContext(context.Background(), data)
}

// CreateFixedPaymentCodeWithContext creates new retail outlet fixed payment code with context
func CreateFixedPaymentCodeWithContext(ctx context.Context, data *CreateFixedPaymentCodeParams) (*xendit.RetailOutlet, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreateFixedPaymentCodeWithContext(ctx, data)
}

// GetFixedPaymentCode gets one retail outlet fixed payment code
func GetFixedPaymentCode(data *GetFixedPaymentCodeParams) (*xendit.RetailOutlet, *xendit.Error) {
	return GetFixedPaymentCodeWithContext(context.Background(), data)
}

// GetFixedPaymentCodeWithContext gets one retail outlet fixed payment code with context
func GetFixedPaymentCodeWithContext(ctx context.Context, data *GetFixedPaymentCodeParams) (*xendit.RetailOutlet, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetFixedPaymentCodeWithContext(ctx, data)
}

// GetPaymentByFixedPaymentCode gets one retail outlet fixed payment code
func GetPaymentByFixedPaymentCode(data *GetPaymentByFixedPaymentCodeParams) (*xendit.RetailOutletPayments, *xendit.Error) {
	return GetPaymentByFixedPaymentCodeWithContext(context.Background(), data)
}

// GetPaymentByFixedPaymentCodeWithContext gets one retail outlet fixed payment code with context
func GetPaymentByFixedPaymentCodeWithContext(ctx context.Context, data *GetPaymentByFixedPaymentCodeParams) (*xendit.RetailOutletPayments, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetPaymentByFixedPaymentCodeWithContext(ctx, data)
}

// UpdateFixedPaymentCode updates a retail outlet fixed payment code
func UpdateFixedPaymentCode(data *UpdateFixedPaymentCodeParams) (*xendit.RetailOutlet, *xendit.Error) {
	return UpdateFixedPaymentCodeWithContext(context.Background(), data)
}

// UpdateFixedPaymentCodeWithContext updates a retail outlet fixed payment code with context
func UpdateFixedPaymentCodeWithContext(ctx context.Context, data *UpdateFixedPaymentCodeParams) (*xendit.RetailOutlet, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.UpdateFixedPaymentCodeWithContext(ctx, data)
}

func getClient() (*Client, *xendit.Error) {
	return &Client{
		Opt:          &xendit.Opt,
		APIRequester: xendit.GetAPIRequester(),
	}, nil
}
