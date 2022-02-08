package transaction

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/utils/validator"
)

// Client is the client used to invoke invoice API.
type Client struct {
	Opt          *xendit.Option
	APIRequester xendit.APIRequester
}

// GetTransaction gets one transaction
func (c *Client) GetTransaction(data *GetTransactionParams) (*xendit.Transaction, *xendit.Error) {
	return c.GetTransactionnWithContext(context.Background(), data)
}

// GetTransactionnWithContext gets one transaction with context
func (c *Client) GetTransactionnWithContext(ctx context.Context, data *GetTransactionParams) (*xendit.Transaction, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.Transaction{}

	header := http.Header{}
	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/transactions/%s", c.Opt.XenditURL, data.TransactionID),
		c.Opt.SecretKey,
		header,
		nil,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetListTransaction gets list transaction
func (c *Client) GetListTransaction(data *GetListTransactionParams) (*xendit.ListTransactions, *xendit.Error) {
	return c.GetListTransactionWithContext(context.Background(), data)
}

// GetListTransactionWithContext gets list transaction with context
func (c *Client) GetListTransactionWithContext(ctx context.Context, data *GetListTransactionParams) (*xendit.ListTransactions, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.ListTransactions{}
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	qs, errParseQuery := query.Values(data)
	if errParseQuery != nil {
		return nil, xendit.FromGoErr(errParseQuery)
	}
	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/transactions?%s", c.Opt.XenditURL, qs.Encode()),
		c.Opt.SecretKey,
		header,
		nil,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}
