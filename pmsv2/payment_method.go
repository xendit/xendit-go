package pmsv2

import (
	"context"

	"github.com/xendit/xendit-go"
)

func CreatePaymentMethod(data *CreatePaymentMethodParams) (*xendit.PaymentMethodResponse, *xendit.Error) {
	return CreatePaymentMethodWithContext(context.Background(), data)
}

func CreatePaymentMethodWithContext(ctx context.Context, data *CreatePaymentMethodParams) (*xendit.PaymentMethodResponse, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreatePaymentMethodWithContext(ctx, data)
}

func AuthPaymentMethod(data *ValidateOTPRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {
	return AuthPaymentMethodWithContext(context.Background(), data)
}

func AuthPaymentMethodWithContext(ctx context.Context, data *ValidateOTPRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.AuthPaymentMethodWithContext(ctx, data)
}

func ExpirePaymentMethod(data *ExpireRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {
	return ExpirePaymentMethodWithContext(context.Background(), data)
}

func ExpirePaymentMethodWithContext(ctx context.Context, data *ExpireRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.ExpirePaymentMethodWithContext(ctx, data)
}

func RetrievePaymentMethod(data *RetrievePaymentMethodRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {
	return RetrievePaymentMethodWithContext(context.Background(), data)
}

func RetrievePaymentMethodWithContext(ctx context.Context, data *RetrievePaymentMethodRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.RetrievePaymentMethodWithContext(ctx, data)
}

func RetrieveAllPaymentMethods(data *RetrieveAllPaymentMethodsRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {
	return RetrieveAllPaymentMethodsWithContext(context.Background(), data)
}

func RetrieveAllPaymentMethodsWithContext(ctx context.Context, data *RetrieveAllPaymentMethodsRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.RetrieveAllPaymentMethodsWithContext(ctx, data)
}

func UpdatePaymentMethod(data *UpdateRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {
	return UpdatePaymentMethodWithContext(context.Background(), data)
}

func UpdatePaymentMethodWithContext(ctx context.Context, data *UpdateRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.UpdatePaymentMethodWithContext(ctx, data)
}

func RetrievePayments(data *RetrievePaymentsRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {
	return RetrievePaymentsWithContext(context.Background(), data)
}

func RetrievePaymentsWithContext(ctx context.Context, data *RetrievePaymentsRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.RetrievePaymentsWithContext(ctx, data)
}

func getClient() (*Client, *xendit.Error) {
	return &Client{
		Opt:          &xendit.Opt,
		APIRequester: xendit.GetAPIRequester(),
	}, nil
}
