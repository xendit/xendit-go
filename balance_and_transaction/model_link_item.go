/*
Transaction Service V4 API

# Introduction This specification describes how to use the Transaction Service V4 API.   **Transaction Service** is the service that records the customer transactions and is responsible to calculate their balance.  All products that move customer money around whether it is money-in, money-out, or transfer will interact with the Transaction Service on its flow. Transaction Service is the source of truth of Xendit and Customer regarding how much money that customer has that is stored in Xendit. Transaction Service is the source that is used for both our internal and customer financial reconciliation. Internally, the Transaction Service data structure is similar to how double-entry accounting works.  ## How Xendit teams/services do integrate with Transaction Service V4   **Channel product team/service** They interact with the Transaction Service when they want to record the transactions. This transaction can be money-in (balance added), money-out (balance deducted), transfer, refund/void/reversal, or other kind of transaction that affects customer balance. Product team also interacts with the Transaction Service for getting information about the transaction or balance.  **Billing/Fee team/service** They interact with Transaction Service either as the dependency of Transaction Service for getting the correct fee calculation/settings. Or using Transaction Service for getting the transaction/fee information to calculate the bill for the customer.  **NUX team/service** They interact with the Transaction Service to set up the customer ledger_account that is used to record their transactions.  **Finance team/service** They interact with the Transaction Service to get the transaction and balance data for each customer to do reconciliation.  **Dashboard/API team/service** They interact with the Transaction Service as a proxy to show the data to the Customer.  ## Prerequisites  Before staring to use **Transaction Service API** you need to complete a few things: 1. Find out **Base URL** for the API. Every endpoint definition in this document contains list of available servers (local, staging, production) 2. Set up ledger accounts using business id and currency. **Ledger Account** represents the account of the customer that will be used to associate with ledger lines. Each business may have at least 1 ledger account group (a group consists of a few accounts of types such as cash, liability, holding), and the money movement of their ledger will revolve around those ledger accounts. **Ledger Lines** that show a debit or credit transaction for a ledger account. We’re using the double-entry principle in accounting where we should post 2 lines every time we make a transaction, 1 to debit an account and 1 to credit another account. See how to call <a href=\"#operation/setupLedgerAccounts\">Create cash, liability, holding, and tax account for a business (api/ledger-accounts/setup)</a> section of this document 3. To be able to create payments with fee/VAT the Product rate settings and VAT rate settings should be created using Transaction Fee Service. See <a href=\"https://docs.google.com/document/d/1HrrA4GhWD1DaJS5dn0dh9VyMhLR6TOUMW1qhRUZ9d_k/edit?pli=1#heading=h.518me3lwf8rb\">Fee Service Documentation</a> for details about how to create Product/VAT rate settings.   ## Transaction flows  To integrate with the Transaction Service you should decide what types of transaction flows your integration will be using. Transaction flow is set by the transaction `type` during transaction creation  1. Money In flows     1. Payment from credit card           `type: CREDIT_CARD_PAYMENT`           3. Payment from other sources without fee/VAT           `type: DEPOSIT, FOREX_DEPOSIT, ISSUING_FUNDING_REFUND, BNPL_PARTNER_SETTLEMENT_CREDIT, PROMO_FEE_CASHBACK, PROMO_VAT_CASHBACK, BATCH_VA_PAYMENT`           4. Payment from other sources with fee/VAT           `type: VA_PAYMENT, IM_ESCROW_VA_PAYMENT, IM_DEPOSIT, RO_PAYMENT, EWALLET_PAYMENT, CARDLESS_CREDIT_PAYMENT, IM_REMITTANCE_VA_PAYMENT, PAYLATER_PAYMENT, INVOICE, QR_CODE_PAYMENT, DIRECT_DEBIT_PAYMENT, DIRECT_BANK_TRANSFER, ACH_PAYMENT, CRYPTO_PAYMENT`           5. Billing deposit from cash           `type: BILLING_DEPOSIT`           6. Billing deposit from other sources           `type: BILLING_DIRECT_DEPOSIT, BILLING_VA_DIRECT_DEPOSIT`       2. Money out flows     1. Instant payment           `type: simple money out types`              `status: COMPLETED`           2. Simple payment without fee/VAT           `type: CHARGEBACK_DEDUCTION, FRAUD_DEDUCTION, LOAN_REPAYMENT, FOREX_DEDUCTION, BNPL_PARTNER_SETTLEMENT_DEBIT, WITHDRAWAL`       3. Simple payment with fee/VAT           `type: ISSUING_FUNDING, BATCH_DISBURSEMENT, CASH_DISBURSEMENT, DISBURSEMENT, REMITTANCE, REMITTANCE_PAYOUT, TAX_DISBURSEMENT`           4. Billing withdraw to cash           `type: BILLING_WITHDRAWAL`           4. Billing withdraw to other destinations           `type: BILL_PAYMENT`       3. Reversal flow      Some of transactions could be reversed. See <a href=\"#section/Introduction/Reversible-non-reversible-transaction-types\">Reversible / non reversible transaction types</a> section of this document. To reverse transaction you should call <a href=\"#operation/updateTransaction\">Update transaction (/api/transactions/:id)</a>  endpoint with the transaction status `REVERSED`.    4. Void/Cancellation Flow      Transaction in the `PENDING_SETTLEMENT` status could be canceled. To do that you should call <a href=\"#operation/updateTransaction\">Update transaction (/api/transactions/:id)</a>  endpoint with the transaction status `VOIDED`.       5. Switcher flow      Switchers are transactions that do not affect the customer balance. These are transactions that goes directly to the customers’ account and simply passes through Xendit. Therefore, it will not impact the customer balance and we will only charge Fee and VAT. To create switcher flow you should set `is_switcher_payment` field to `true`.       ## Instant/non instant settlement  Transactions can be performed instantly (instant settlement) or with delay (non instant settlement).  Some of the transaction types are only instantly processed, some of them support both instant and non instant settlement and some of them have only non instant settlement. If settlement is instant than balance will be changed instantly. In opposite case the transaction status has to be set into PENDING_SETTLEMENT and settlement date should be provided.   1. Instant settlement Money In transaction types      `DEPOSIT, BATCH_VA_PAYMENT, FOREX_DEPOSIT, IM_DEPOSIT, CARDLESS_CREDIT_PAYMENT, ISSUING_FUNDING_REFUND, BNPL_PARTNER_SETTLEMENT_CREDIT, PROMO_FEE_CASHBACK, PROMO_VAT_CASHBACK, REMITTANCE_VA_PAYMENT_CLAIM`    2. Both instant and non instant Money In transaction types      `DIRECT_DEBIT_PAYMENT, DIRECT_BANK_TRANSFER, ACH_PAYMENT, RO_PAYMENT, EWALLET_PAYMENT, QR_CODE_PAYMENT, VA_PAYMENT, INVOICE, PAYLATER_PAYMENT`  3. Non Instant settlement Money In transaction types      `CREDIT_CARD_PAYMENT`    4. Instant settlement Money Out transaction types      `LOAN_REPAYMENT, FOREX_DEDUCTION, BILL_PAYMENT, ISSUING_FUNDING, BNPL_PARTNER_SETTLEMENT_DEBIT, FRAUD_DEDUCTION`  5. Both instant and non instant settlement supported Money Out transaction types      `CHARGEBACK_DEDUCTION`  6. Non Instant settlement Money Out transaction types      All other money out types are non instant settlement  ## Reversible / non reversible transaction types  Some transactions can be reversed. Here are the list of transaction types that could be reversed:   `CASH_DISBURSEMENT, DISBURSEMENT, BATCH_DISBURSEMENT, REMITTANCE, REMITTANCE_PAYOUT, TAX_DISBURSEMENT, WITHDRAWAL, DEPOSIT, FOREX_DEPOSIT, FOREX_DEDUCTION, VA_PAYMENT, BATCH_VA_PAYMENT, IM_REMITTANCE_VA_PAYMENT, IM_ESCROW_VA_PAYMENT, IM_DEPOSIT, REMITTANCE_VA_PAYMENT, REMITTANCE_VA_PAYMENT_CLAIM, RO_PAYMENT, CARDLESS_CREDIT_PAYMENT, PAYLATER_PAYMENT, INVOICE, QR_CODE_PAYMENT, CREDIT_CARD_PAYMENT, EWALLET_PAYMENT, DIRECT_DEBIT_PAYMENT, DIRECT_BANK_TRANSFER, ACH_PAYMENT, CHARGEBACK_DEDUCTION, FRAUD_DEDUCTION, LOAN_REPAYMENT, ISSUING_FUNDING, ISSUING_FUNDING_REFUND, BNPL_PARTNER_SETTLEMENT_DEBIT, BNPL_PARTNER_SETTLEMENT_CREDIT, BILLING_DEPOSIT, BILLING_DIRECT_DEPOSIT, BILLING_VA_DIRECT_DEPOSIT, BILLING_WITHDRAWAL, BILL_PAYMENT, PROMO_FEE_CASHBACK, PROMO_VAT_CASHBACK`       ## How to create transaction  After you created or already have the `BUSINESS_CASH` ledger account ID (See <a href=\"#section/Introduction/Prerequisites\">Prerequisites</a> section) and you know what transaction flows are going to be used  you can create the new transaction using POST request to the  <a href=\"#operation/createTransaction\">Create a new transaction (/api/transactions)</a>  endpoint  ## How to update transaction  To update transaction you should do  PATCH request to the  <a href=\"#operation/updateTransaction\">Update transaction (/api/transactions/::id)</a>  endpoint   

API version: 5.4.0
*/


package balance_and_transaction

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
)

// checks if the LinkItem type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &LinkItem{}

// LinkItem struct for LinkItem
type LinkItem struct {
	// URI of target, this will be to the next link.
	Href string `json:"href"`
	// The relationship between source and target. The value will be `next`.
	Rel string `json:"rel"`
	// The HTTP method, the value will be `GET`.
	Method string `json:"method"`
}

// NewLinkItem instantiates a new LinkItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkItem(href string, rel string, method string) *LinkItem {
	this := LinkItem{}
	this.Href = href
	this.Rel = rel
	this.Method = method
	return &this
}

// NewLinkItemWithDefaults instantiates a new LinkItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkItemWithDefaults() *LinkItem {
	this := LinkItem{}
	return &this
}

// GetHref returns the Href field value
func (o *LinkItem) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *LinkItem) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *LinkItem) SetHref(v string) {
	o.Href = v
}

// GetRel returns the Rel field value
func (o *LinkItem) GetRel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rel
}

// GetRelOk returns a tuple with the Rel field value
// and a boolean to check if the value has been set.
func (o *LinkItem) GetRelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rel, true
}

// SetRel sets field value
func (o *LinkItem) SetRel(v string) {
	o.Rel = v
}

// GetMethod returns the Method field value
func (o *LinkItem) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *LinkItem) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *LinkItem) SetMethod(v string) {
	o.Method = v
}

func (o LinkItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href
	toSerialize["rel"] = o.Rel
	toSerialize["method"] = o.Method
	return toSerialize, nil
}

type NullableLinkItem struct {
	value *LinkItem
	isSet bool
}

func (v NullableLinkItem) Get() *LinkItem {
	return v.value
}

func (v *NullableLinkItem) Set(val *LinkItem) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkItem) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkItem(val *LinkItem) *NullableLinkItem {
	return &NullableLinkItem{value: val, isSet: true}
}

func (v NullableLinkItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


