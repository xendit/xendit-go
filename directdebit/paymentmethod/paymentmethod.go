package paymentmethod

import (
	"context"

	"github.com/xendit/xendit-go"
)

// CreatePaymentMethod creates new payment method
func CreatePaymentMethod(data *CreatePaymentMethodParams) (*xendit.PaymentMethod, *xendit.Error) {
	return CreatePaymentMethodWithContext(context.Background(), data)
}

// CreatePaymentMethodWithContext creates new payment method
func CreatePaymentMethodWithContext(ctx context.Context, data *CreatePaymentMethodParams) (*xendit.PaymentMethod, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreatePaymentMethodWithContext(ctx, data)
}

// GetPaymentMethodsByCustomerID gets payment methods by customer ID
func GetPaymentMethodsByCustomerID(data *GetPaymentMethodsByCustomerIDParams) ([]xendit.PaymentMethod, *xendit.Error) {
	return GetPaymentMethodsByCustomerIDWithContext(context.Background(), data)
}

// GetPaymentMethodsByCustomerIDWithContext gets payment methods by customer ID
func GetPaymentMethodsByCustomerIDWithContext(ctx context.Context, data *GetPaymentMethodsByCustomerIDParams) ([]xendit.PaymentMethod, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetPaymentMethodsByCustomerIDWithContext(ctx, data)
}

func getClient() (*Client, *xendit.Error) {
	return &Client{
		Opt:          &xendit.Opt,
		APIRequester: xendit.GetAPIRequester(),
	}, nil
}
