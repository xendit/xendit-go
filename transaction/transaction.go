package transaction

import (
	"context"

	"github.com/xendit/xendit-go"
)

// GetTransaction gets one retail outlet fixed payment code
func GetTransaction(data *GetTransactionParams) (*xendit.Transaction, *xendit.Error) {
	return GetTransactionnWithContext(context.Background(), data)
}

// GetTransactionnWithContext gets one retail outlet fixed payment code with context
func GetTransactionnWithContext(ctx context.Context, data *GetTransactionParams) (*xendit.Transaction, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetTransactionnWithContext(ctx, data)
}

// GetListTransaction gets one retail outlet fixed payment code
func GetListTransaction(data *GetListTransactionParams) (*xendit.ListTransactions, *xendit.Error) {
	return GetListTransactionnWithContext(context.Background(), data)
}

// GetListTransactionnWithContext gets one retail outlet fixed payment code with context
func GetListTransactionnWithContext(ctx context.Context, data *GetListTransactionParams) (*xendit.ListTransactions, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetListTransactionWithContext(ctx, data)
}

func getClient() (*Client, *xendit.Error) {
	return &Client{
		Opt:          &xendit.Opt,
		APIRequester: xendit.GetAPIRequester(),
	}, nil
}
