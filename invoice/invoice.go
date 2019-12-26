package invoice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	form "github.com/xendit/xendit-go/form"
	option "github.com/xendit/xendit-go/option"
)

// Invoice is ...
type Invoice struct {
	Opt *option.Option
}

// bank is data that contained in API response of invoice in available_banks
type bank struct {
	BankCode          string `json:"bank_code"`
	CollectionType    string `json:"collection_type"`
	BankAccountNumber string `json:"bank_account_number"`
	TransferAmount    int64  `json:"transfer_amount"`
	BankBranch        string `json:"bank_branch"`
	AccountHolderName string `json:"account_holder_name"`
	IdentityAmount    int    `json:"identity_amount"`
}

// eWallet is data that contained in response at availableEWallets
type eWallet struct {
	EWalletType string `json:"ewallet_type"`
}

// retailOutlet is data that contained in response at availableEWallets
type retailOutlet struct {
	RetailOutletName string `json:"retail_outlet_name"`
	PaymentCode      string `json:"payment_code"`
	TransferAmount   int64  `json:"transfer_amount"`
	MerchantName     string `json:"merchant_name,omitempty"`
}

//
type item struct {
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int    `json:"quantity"`
}

// Response is data that contained in API response of invoice related request
type Response struct {
	ID                        string         `json:"id"`
	ExternalID                string         `json:"external_id"`
	UserID                    string         `json:"user_id"`
	PayerEmail                string         `json:"payer_email"`
	Description               string         `json:"description"`
	Amount                    int64          `json:"amount"`
	MerchantName              string         `json:"merchant_name"`
	MerchantProfilePictureURL string         `json:"merchant_profile_picture_url"`
	InvoiceURL                string         `json:"invoice_url"`
	ExpiryDate                string         `json:"expiry_date"`
	AvailableBanks            []bank         `json:"available_banks,omitempty"`
	AvailableEWallets         []eWallet      `json:"available_ewallets,omitempty"`
	AvailableRetailOutlets    []retailOutlet `json:"available_retail_outlets,omitempty"`
	ShouldExcludeCreditCard   bool           `json:"should_exclude_credit_card"`
	ShouldSendEmail           bool           `json:"should_send_email"`
	Created                   string         `json:"created"`
	Updated                   string         `json:"updated"`
	BankCode                  string         `json:"bank_code,omitempty"`
	PaidAmount                string         `json:"paid_amount,omitempty"`
	AdjustedReceivedAmount    string         `json:"adjusted_received_amount,omitempty"`
	RecurringPaymentID        string         `json:"recurring_payment_id,omitempty"`
	CreditCardChargeID        string         `json:"credit_card_charge_id,omitempty"`
	Currency                  string         `json:"currency,omitempty"`
	InitialCurrency           string         `json:"initial_currency,omitempty"`
	InitialAmount             string         `json:"initial_amount,omitempty"`
	PaidAt                    string         `json:"paid_at,omitempty"`
	MidLabel                  string         `json:"mid_label,omitempty"`
	PaymentChannel            string         `json:"payment_channel,omitempty"`
	PaymentMethod             string         `json:"payment_method,omitempty"`
	PaymentDestination        string         `json:"payment_destination,omitempty"`
	SuccessRedirectURL        string         `json:"success_redirect_url,omitempty"`
	FailureRedirectURL        string         `json:"failure_redirect_url,omitempty"`
	Items                     string         `json:"items,omitempty"`
	FixedVA                   string         `json:"fixed_va,omitempty"`
	ErrorCode                 string         `json:"error_code,omitempty"`
	ErrorMessage              string         `json:"message,omitempty"`
}

// TODO: implement context

// CreateInvoice creates new invoice
func (i Invoice) CreateInvoice(data form.CreateInvoiceData) (Response, error) {
	// TODO: validate the input data

	var response Response

	reqBody, err := json.Marshal(data)
	if err != nil {
		return response, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/v2/invoices", i.Opt.XenditURL),
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return response, err
	}
	req.SetBasicAuth(i.Opt.SecretKey, "")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return response, err
	}

	return response, nil
}

// GetInvoice gets one invoice
func (i Invoice) GetInvoice(invoiceID string) (Response, error) {

	var response Response

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/v2/invoices/%s", i.Opt.XenditURL, invoiceID),
		nil,
	)
	if err != nil {
		return response, err
	}
	req.SetBasicAuth(i.Opt.SecretKey, "")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return response, err
	}

	return response, nil
}

// ExpireInvoice expire the created invoice
func (i Invoice) ExpireInvoice(invoiceID string) (Response, error) {

	var response Response

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/invoices/%s/expire!", i.Opt.XenditURL, invoiceID),
		nil,
	)
	if err != nil {
		return response, err
	}
	req.SetBasicAuth(i.Opt.SecretKey, "")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return response, err
	}

	return response, nil
}

func removeFromList(slice []string, idx int) []string {
	return append(slice[:idx], slice[idx+1:]...)
}

func convertJSONStringToStringList(stringParams string) [][]string {
	var stringList [][]string

	stringParamsList := strings.Split(stringParams, ",")

	loopEnd := len(stringParamsList)
	for i := 0; i < loopEnd; i++ {
		if strings.Contains(stringParamsList[i], "[") {
			stringParamsList[i] = stringParamsList[i] + stringParamsList[i+1]
			stringParamsList = removeFromList(stringParamsList, i+1)
			loopEnd--
		}

		stringList = append(stringList, strings.SplitN(stringParamsList[i], ":", 2))
	}

	return stringList
}

// GetAllInvoices gets all invoices with conditions
func (i Invoice) GetAllInvoices(data form.GetAllInvoicesData) ([]Response, error) {
	// TODO: validate the input data
	var responses []Response

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/v2/invoices", i.Opt.XenditURL),
		nil,
	)
	if err != nil {
		return responses, err
	}
	req.SetBasicAuth(i.Opt.SecretKey, "")

	reqParamsByte, err := json.Marshal(data)
	if err != nil {
		return responses, err
	}
	paramList := convertJSONStringToStringList(string(reqParamsByte)[1 : len(reqParamsByte)-1])

	queryParams := req.URL.Query()
	for _, param := range paramList {
		queryParams.Add(param[0], param[1])
	}
	req.URL.RawQuery = queryParams.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return responses, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return responses, err
	}

	if err := json.Unmarshal(body, &responses); err != nil {
		return responses, err
	}

	return responses, nil
}
