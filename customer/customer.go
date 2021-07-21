package customer

import (
	"context"

	"github.com/xendit/xendit-go"
)

// CreateCustomer creates new customer
func CreateCustomer(data *CreateCustomerParams) (*xendit.Customer, *xendit.Error) {
	return CreateCustomerWithContext(context.Background(), data)
}

// CreateCustomerWithContext creates new payment
func CreateCustomerWithContext(ctx context.Context, data *CreateCustomerParams) (*xendit.Customer, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreateCustomerWithContext(ctx, data)
}

// GetCustomerByReferenceID gets customer by reference ID
func GetCustomerByReferenceID(data *GetCustomerByReferenceIDParams) ([]xendit.Customer, *xendit.Error) {
	return GetCustomerByReferenceIDWithContext(context.Background(), data)
}

// GetCustomerByReferenceIDWithContext gets customer by reference ID
func GetCustomerByReferenceIDWithContext(ctx context.Context, data *GetCustomerByReferenceIDParams) ([]xendit.Customer, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetCustomerByReferenceIDWithContext(ctx, data)
}

func getClient() (*Client, *xendit.Error) {
	return &Client{
		Opt:          &xendit.Opt,
		APIRequester: xendit.GetAPIRequester(),
	}, nil
}
