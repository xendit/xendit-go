package invoice

import (
	"encoding/json"
	"fmt"

	"github.com/xendit/xendit-go"

	validator "github.com/go-playground/validator"
)

// Client is the client used to invoke invoice API.
type Client struct {
	Opt           *xendit.Option
	HTTPRequester xendit.HTTPRequester
}

// CreateInvoice creates new invoice
func CreateInvoice(data *xendit.CreateInvoiceParams) (*xendit.Invoice, error) {
	return getClient().CreateInvoice(data)
}

// CreateInvoice creates new invoice
func (c Client) CreateInvoice(data *xendit.CreateInvoiceParams) (*xendit.Invoice, error) {
	var response *xendit.Invoice

	v := validator.New()
	if err := v.Struct(data); err != nil {
		return nil, err
	}

	reqBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	respBody, err := c.HTTPRequester.Call(
		"POST",
		fmt.Sprintf("%s/v2/invoices", c.Opt.XenditURL),
		c.Opt.SecretKey,
		string(reqBody),
	)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetInvoice gets one invoice
func GetInvoice(invoiceID string) (*xendit.Invoice, error) {
	return getClient().GetInvoice(invoiceID)
}

// GetInvoice gets one invoice
func (c Client) GetInvoice(invoiceID string) (*xendit.Invoice, error) {
	var response *xendit.Invoice

	respBody, err := c.HTTPRequester.Call(
		"GET",
		fmt.Sprintf("%s/v2/invoices/%s", c.Opt.XenditURL, invoiceID),
		c.Opt.SecretKey,
		"",
	)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// ExpireInvoice expire the created invoice
func ExpireInvoice(invoiceID string) (*xendit.Invoice, error) {
	return getClient().ExpireInvoice(invoiceID)
}

// ExpireInvoice expire the created invoice
func (c Client) ExpireInvoice(invoiceID string) (*xendit.Invoice, error) {
	var response *xendit.Invoice

	respBody, err := c.HTTPRequester.Call(
		"POST",
		fmt.Sprintf("%s/invoices/%s/expire!", c.Opt.XenditURL, invoiceID),
		c.Opt.SecretKey,
		"",
	)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetAllInvoices gets all invoices with conditions
func GetAllInvoices(data *xendit.GetAllInvoicesParams) ([]xendit.Invoice, error) {
	return getClient().GetAllInvoices(data)
}

// GetAllInvoices gets all invoices with conditions
func (c Client) GetAllInvoices(data *xendit.GetAllInvoicesParams) ([]xendit.Invoice, error) {
	var responses []xendit.Invoice

	reqParams, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	respBody, err := c.HTTPRequester.Call(
		"GET",
		fmt.Sprintf("%s/v2/invoices", c.Opt.XenditURL),
		c.Opt.SecretKey,
		string(reqParams),
	)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(respBody, &responses); err != nil {
		return nil, err
	}

	return responses, nil
}

func getClient() Client {
	return Client{
		Opt: &xendit.Option{
			PublicKey: xendit.PublicKey,
			SecretKey: xendit.SecretKey,
			XenditURL: xendit.XenditURL,
		},
		HTTPRequester: xendit.GetHTTPRequester(),
	}
}
