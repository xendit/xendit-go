package xendit

import "time"

// Transaction contains data from Xendit's API response of invoice related request.
// For more details see https://developers.xendit.co/api-reference/#transactions.
// For documentation of subpackage payout, checkout https://pkg.go.dev/github.com/xendit/xendit-go/transaction
type Transaction struct {
	ID                string         `json:"id"`
	ProductID         string         `json:"product_id"`
	Type              string         `json:"type"`
	ChannelCode       string         `json:"channel_code,omitempty"`
	ReferenceID       string         `json:"reference_id"`
	AccountIdentifier string         `json:"account_identifier,omitempty"`
	Currency          string         `json:"currency,omitempty"`
	Amount            float64        `json:"amount"`
	NetAmount         float64        `json:"net_amount"`
	Cashflow          string         `json:"cashflow"`
	Status            string         `json:"status"`
	ChannelCategory   string         `json:"channel_category"`
	BusinessID        string         `json:"business_id"`
	Created           *time.Time     `json:"created"`
	Updated           *time.Time     `json:"updated"`
	Fee               TransactionFee `json:"fee"`
}

type TransactionFee struct {
	XenditFee                float64 `json:"xendit_fee,omitempty"`
	ValueAdded_tax           float64 `json:"value_added_tax,omitempty"`
	XenditWithholdingTax     float64 `json:"xendit_withholding_tax,omitempty"`
	ThirdPartyWithholdingTax float64 `json:"third_party_withholding_tax,omitempty"`
	Status                   string  `json:"status,omitempty"`
}

type ListTransactions struct {
	Data    []Transaction         `json:"data"`
	HasMore bool                  `json:"has_more"`
	Links   ListTransactionsLinks `json:"links"`
}

// ListTransactionsLinks is data that contained in `ListTransactions` at Links field.
type ListTransactionsLinks struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}
