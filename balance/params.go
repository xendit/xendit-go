package balance

import (
	"net/url"
)

// GetParams contains parameters for Get
type GetParams struct {
	AccountType string `json:"account_type"`
}

// QueryString creates query string from GetAllParams, ignores nil values
func (p *GetParams) QueryString() string {
	urlValues := &url.Values{}

	if p.AccountType != "" {
		urlValues.Add("account_type", p.AccountType)
	}

	return urlValues.Encode()
}
