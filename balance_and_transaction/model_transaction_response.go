/*
Transaction Service V4 API

# Introduction This specification describes how to use the Transaction Service V4 API.   **Transaction Service** is the service that records the customer transactions and is responsible to calculate their balance.  All products that move customer money around whether it is money-in, money-out, or transfer will interact with the Transaction Service on its flow. Transaction Service is the source of truth of Xendit and Customer regarding how much money that customer has that is stored in Xendit. Transaction Service is the source that is used for both our internal and customer financial reconciliation. Internally, the Transaction Service data structure is similar to how double-entry accounting works.  ## How Xendit teams/services do integrate with Transaction Service V4   **Channel product team/service** They interact with the Transaction Service when they want to record the transactions. This transaction can be money-in (balance added), money-out (balance deducted), transfer, refund/void/reversal, or other kind of transaction that affects customer balance. Product team also interacts with the Transaction Service for getting information about the transaction or balance.  **Billing/Fee team/service** They interact with Transaction Service either as the dependency of Transaction Service for getting the correct fee calculation/settings. Or using Transaction Service for getting the transaction/fee information to calculate the bill for the customer.  **NUX team/service** They interact with the Transaction Service to set up the customer ledger_account that is used to record their transactions.  **Finance team/service** They interact with the Transaction Service to get the transaction and balance data for each customer to do reconciliation.  **Dashboard/API team/service** They interact with the Transaction Service as a proxy to show the data to the Customer.  ## Prerequisites  Before staring to use **Transaction Service API** you need to complete a few things: 1. Find out **Base URL** for the API. Every endpoint definition in this document contains list of available servers (local, staging, production) 2. Set up ledger accounts using business id and currency. **Ledger Account** represents the account of the customer that will be used to associate with ledger lines. Each business may have at least 1 ledger account group (a group consists of a few accounts of types such as cash, liability, holding), and the money movement of their ledger will revolve around those ledger accounts. **Ledger Lines** that show a debit or credit transaction for a ledger account. We’re using the double-entry principle in accounting where we should post 2 lines every time we make a transaction, 1 to debit an account and 1 to credit another account. See how to call <a href=\"#operation/setupLedgerAccounts\">Create cash, liability, holding, and tax account for a business (api/ledger-accounts/setup)</a> section of this document 3. To be able to create payments with fee/VAT the Product rate settings and VAT rate settings should be created using Transaction Fee Service. See <a href=\"https://docs.google.com/document/d/1HrrA4GhWD1DaJS5dn0dh9VyMhLR6TOUMW1qhRUZ9d_k/edit?pli=1#heading=h.518me3lwf8rb\">Fee Service Documentation</a> for details about how to create Product/VAT rate settings.   ## Transaction flows  To integrate with the Transaction Service you should decide what types of transaction flows your integration will be using. Transaction flow is set by the transaction `type` during transaction creation  1. Money In flows     1. Payment from credit card           `type: CREDIT_CARD_PAYMENT`           3. Payment from other sources without fee/VAT           `type: DEPOSIT, FOREX_DEPOSIT, ISSUING_FUNDING_REFUND, BNPL_PARTNER_SETTLEMENT_CREDIT, PROMO_FEE_CASHBACK, PROMO_VAT_CASHBACK, BATCH_VA_PAYMENT`           4. Payment from other sources with fee/VAT           `type: VA_PAYMENT, IM_ESCROW_VA_PAYMENT, IM_DEPOSIT, RO_PAYMENT, EWALLET_PAYMENT, CARDLESS_CREDIT_PAYMENT, IM_REMITTANCE_VA_PAYMENT, PAYLATER_PAYMENT, INVOICE, QR_CODE_PAYMENT, DIRECT_DEBIT_PAYMENT, DIRECT_BANK_TRANSFER, ACH_PAYMENT, CRYPTO_PAYMENT`           5. Billing deposit from cash           `type: BILLING_DEPOSIT`           6. Billing deposit from other sources           `type: BILLING_DIRECT_DEPOSIT, BILLING_VA_DIRECT_DEPOSIT`       2. Money out flows     1. Instant payment           `type: simple money out types`              `status: COMPLETED`           2. Simple payment without fee/VAT           `type: CHARGEBACK_DEDUCTION, FRAUD_DEDUCTION, LOAN_REPAYMENT, FOREX_DEDUCTION, BNPL_PARTNER_SETTLEMENT_DEBIT, WITHDRAWAL`       3. Simple payment with fee/VAT           `type: ISSUING_FUNDING, BATCH_DISBURSEMENT, CASH_DISBURSEMENT, DISBURSEMENT, REMITTANCE, REMITTANCE_PAYOUT, TAX_DISBURSEMENT`           4. Billing withdraw to cash           `type: BILLING_WITHDRAWAL`           4. Billing withdraw to other destinations           `type: BILL_PAYMENT`       3. Reversal flow      Some of transactions could be reversed. See <a href=\"#section/Introduction/Reversible-non-reversible-transaction-types\">Reversible / non reversible transaction types</a> section of this document. To reverse transaction you should call <a href=\"#operation/updateTransaction\">Update transaction (/api/transactions/:id)</a>  endpoint with the transaction status `REVERSED`.    4. Void/Cancellation Flow      Transaction in the `PENDING_SETTLEMENT` status could be canceled. To do that you should call <a href=\"#operation/updateTransaction\">Update transaction (/api/transactions/:id)</a>  endpoint with the transaction status `VOIDED`.       5. Switcher flow      Switchers are transactions that do not affect the customer balance. These are transactions that goes directly to the customers’ account and simply passes through Xendit. Therefore, it will not impact the customer balance and we will only charge Fee and VAT. To create switcher flow you should set `is_switcher_payment` field to `true`.       ## Instant/non instant settlement  Transactions can be performed instantly (instant settlement) or with delay (non instant settlement).  Some of the transaction types are only instantly processed, some of them support both instant and non instant settlement and some of them have only non instant settlement. If settlement is instant than balance will be changed instantly. In opposite case the transaction status has to be set into PENDING_SETTLEMENT and settlement date should be provided.   1. Instant settlement Money In transaction types      `DEPOSIT, BATCH_VA_PAYMENT, FOREX_DEPOSIT, IM_DEPOSIT, CARDLESS_CREDIT_PAYMENT, ISSUING_FUNDING_REFUND, BNPL_PARTNER_SETTLEMENT_CREDIT, PROMO_FEE_CASHBACK, PROMO_VAT_CASHBACK, REMITTANCE_VA_PAYMENT_CLAIM`    2. Both instant and non instant Money In transaction types      `DIRECT_DEBIT_PAYMENT, DIRECT_BANK_TRANSFER, ACH_PAYMENT, RO_PAYMENT, EWALLET_PAYMENT, QR_CODE_PAYMENT, VA_PAYMENT, INVOICE, PAYLATER_PAYMENT`  3. Non Instant settlement Money In transaction types      `CREDIT_CARD_PAYMENT`    4. Instant settlement Money Out transaction types      `LOAN_REPAYMENT, FOREX_DEDUCTION, BILL_PAYMENT, ISSUING_FUNDING, BNPL_PARTNER_SETTLEMENT_DEBIT, FRAUD_DEDUCTION`  5. Both instant and non instant settlement supported Money Out transaction types      `CHARGEBACK_DEDUCTION`  6. Non Instant settlement Money Out transaction types      All other money out types are non instant settlement  ## Reversible / non reversible transaction types  Some transactions can be reversed. Here are the list of transaction types that could be reversed:   `CASH_DISBURSEMENT, DISBURSEMENT, BATCH_DISBURSEMENT, REMITTANCE, REMITTANCE_PAYOUT, TAX_DISBURSEMENT, WITHDRAWAL, DEPOSIT, FOREX_DEPOSIT, FOREX_DEDUCTION, VA_PAYMENT, BATCH_VA_PAYMENT, IM_REMITTANCE_VA_PAYMENT, IM_ESCROW_VA_PAYMENT, IM_DEPOSIT, REMITTANCE_VA_PAYMENT, REMITTANCE_VA_PAYMENT_CLAIM, RO_PAYMENT, CARDLESS_CREDIT_PAYMENT, PAYLATER_PAYMENT, INVOICE, QR_CODE_PAYMENT, CREDIT_CARD_PAYMENT, EWALLET_PAYMENT, DIRECT_DEBIT_PAYMENT, DIRECT_BANK_TRANSFER, ACH_PAYMENT, CHARGEBACK_DEDUCTION, FRAUD_DEDUCTION, LOAN_REPAYMENT, ISSUING_FUNDING, ISSUING_FUNDING_REFUND, BNPL_PARTNER_SETTLEMENT_DEBIT, BNPL_PARTNER_SETTLEMENT_CREDIT, BILLING_DEPOSIT, BILLING_DIRECT_DEPOSIT, BILLING_VA_DIRECT_DEPOSIT, BILLING_WITHDRAWAL, BILL_PAYMENT, PROMO_FEE_CASHBACK, PROMO_VAT_CASHBACK`       ## How to create transaction  After you created or already have the `BUSINESS_CASH` ledger account ID (See <a href=\"#section/Introduction/Prerequisites\">Prerequisites</a> section) and you know what transaction flows are going to be used  you can create the new transaction using POST request to the  <a href=\"#operation/createTransaction\">Create a new transaction (/api/transactions)</a>  endpoint  ## How to update transaction  To update transaction you should do  PATCH request to the  <a href=\"#operation/updateTransaction\">Update transaction (/api/transactions/::id)</a>  endpoint   

API version: 3.4.3
*/


package balance_and_transaction

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
	"time"
)

// checks if the TransactionResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TransactionResponse{}

// TransactionResponse struct for TransactionResponse
type TransactionResponse struct {
	// The unique id of a transaction. It will have `txn_` as prefix
	Id string `json:"id"`
	// The product_id of the transaction. Product id will have a different prefix for each product. You can use this id to match the transaction from this API to each product API.
	ProductId string `json:"product_id"`
	Type TransactionResponseType `json:"type"`
	Status TransactionStatuses `json:"status"`
	ChannelCategory ChannelsCategories `json:"channel_category"`
	// The channel of the transaction that is used. See [channel codes](https://docs.xendit.co/xendisburse/channel-codes) for the list of available per channel categories.
	ChannelCode NullableString `json:"channel_code"`
	// Account identifier of transaction. The format will be different from each channel.
	AccountIdentifier NullableString `json:"account_identifier"`
	// customer supplied reference/external_id
	ReferenceId string `json:"reference_id"`
	Currency Currency `json:"currency"`
	// The transaction amount. The number of decimal places will be different for each currency according to ISO 4217.
	Amount float32 `json:"amount"`
	// Representing whether the transaction is money in or money out For transfer, the transfer out side it will shows up as money out and on transfer in side in will shows up as money-in. Available values are `MONEY_IN` for money in and `MONEY_OUT` for money out.
	Cashflow string `json:"cashflow"`
	// The settlement status of the transaction. `PENDING` - Transaction amount has not been settled to merchant's balance. `SETTLED` - Transaction has been settled to merchant's balance
	SettlementStatus NullableString `json:"settlement_status,omitempty"`
	// Estimated settlement time will only apply to money-in transactions. For money-out transaction, the value will be `NULL`. Estimated settlement time in which transaction amount will be settled to merchant's balance.
	EstimatedSettlementTime NullableTime `json:"estimated_settlement_time,omitempty"`
	// The id of business where this transaction belong to
	BusinessId string `json:"business_id"`
	Fee FeeResponse `json:"fee"`
	// Transaction created timestamp (UTC+0)
	Created time.Time `json:"created"`
	// Transaction updated timestamp (UTC+0)
	Updated time.Time `json:"updated"`
}

// NewTransactionResponse instantiates a new TransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionResponse(id string, productId string, type_ TransactionResponseType, status TransactionStatuses, channelCategory ChannelsCategories, channelCode NullableString, accountIdentifier NullableString, referenceId string, currency Currency, amount float32, cashflow string, businessId string, fee FeeResponse, created time.Time, updated time.Time) *TransactionResponse {
	this := TransactionResponse{}
	this.Id = id
	this.ProductId = productId
	this.Type = type_
	this.Status = status
	this.ChannelCategory = channelCategory
	this.ChannelCode = channelCode
	this.AccountIdentifier = accountIdentifier
	this.ReferenceId = referenceId
	this.Currency = currency
	this.Amount = amount
	this.Cashflow = cashflow
	this.BusinessId = businessId
	this.Fee = fee
	this.Created = created
	this.Updated = updated
	return &this
}

// NewTransactionResponseWithDefaults instantiates a new TransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionResponseWithDefaults() *TransactionResponse {
	this := TransactionResponse{}
	return &this
}

// GetId returns the Id field value
func (o *TransactionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TransactionResponse) SetId(v string) {
	o.Id = v
}

// GetProductId returns the ProductId field value
func (o *TransactionResponse) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *TransactionResponse) SetProductId(v string) {
	o.ProductId = v
}

// GetType returns the Type field value
func (o *TransactionResponse) GetType() TransactionResponseType {
	if o == nil {
		var ret TransactionResponseType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetTypeOk() (*TransactionResponseType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransactionResponse) SetType(v TransactionResponseType) {
	o.Type = v
}

// GetStatus returns the Status field value
func (o *TransactionResponse) GetStatus() TransactionStatuses {
	if o == nil {
		var ret TransactionStatuses
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetStatusOk() (*TransactionStatuses, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TransactionResponse) SetStatus(v TransactionStatuses) {
	o.Status = v
}

// GetChannelCategory returns the ChannelCategory field value
func (o *TransactionResponse) GetChannelCategory() ChannelsCategories {
	if o == nil {
		var ret ChannelsCategories
		return ret
	}

	return o.ChannelCategory
}

// GetChannelCategoryOk returns a tuple with the ChannelCategory field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetChannelCategoryOk() (*ChannelsCategories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelCategory, true
}

// SetChannelCategory sets field value
func (o *TransactionResponse) SetChannelCategory(v ChannelsCategories) {
	o.ChannelCategory = v
}

// GetChannelCode returns the ChannelCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransactionResponse) GetChannelCode() string {
	if o == nil || o.ChannelCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.ChannelCode.Get()
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionResponse) GetChannelCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelCode.Get(), o.ChannelCode.IsSet()
}

// SetChannelCode sets field value
func (o *TransactionResponse) SetChannelCode(v string) {
	o.ChannelCode.Set(&v)
}

// GetAccountIdentifier returns the AccountIdentifier field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransactionResponse) GetAccountIdentifier() string {
	if o == nil || o.AccountIdentifier.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccountIdentifier.Get()
}

// GetAccountIdentifierOk returns a tuple with the AccountIdentifier field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionResponse) GetAccountIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountIdentifier.Get(), o.AccountIdentifier.IsSet()
}

// SetAccountIdentifier sets field value
func (o *TransactionResponse) SetAccountIdentifier(v string) {
	o.AccountIdentifier.Set(&v)
}

// GetReferenceId returns the ReferenceId field value
func (o *TransactionResponse) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *TransactionResponse) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetCurrency returns the Currency field value
func (o *TransactionResponse) GetCurrency() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetCurrencyOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *TransactionResponse) SetCurrency(v Currency) {
	o.Currency = v
}

// GetAmount returns the Amount field value
func (o *TransactionResponse) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransactionResponse) SetAmount(v float32) {
	o.Amount = v
}

// GetCashflow returns the Cashflow field value
func (o *TransactionResponse) GetCashflow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cashflow
}

// GetCashflowOk returns a tuple with the Cashflow field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetCashflowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cashflow, true
}

// SetCashflow sets field value
func (o *TransactionResponse) SetCashflow(v string) {
	o.Cashflow = v
}

// GetSettlementStatus returns the SettlementStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionResponse) GetSettlementStatus() string {
	if o == nil || utils.IsNil(o.SettlementStatus.Get()) {
		var ret string
		return ret
	}
	return *o.SettlementStatus.Get()
}

// GetSettlementStatusOk returns a tuple with the SettlementStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionResponse) GetSettlementStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SettlementStatus.Get(), o.SettlementStatus.IsSet()
}

// HasSettlementStatus returns a boolean if a field has been set.
func (o *TransactionResponse) HasSettlementStatus() bool {
	if o != nil && o.SettlementStatus.IsSet() {
		return true
	}

	return false
}

// SetSettlementStatus gets a reference to the given NullableString and assigns it to the SettlementStatus field.
func (o *TransactionResponse) SetSettlementStatus(v string) {
	o.SettlementStatus.Set(&v)
}
// SetSettlementStatusNil sets the value for SettlementStatus to be an explicit nil
func (o *TransactionResponse) SetSettlementStatusNil() {
	o.SettlementStatus.Set(nil)
}

// UnsetSettlementStatus ensures that no value is present for SettlementStatus, not even an explicit nil
func (o *TransactionResponse) UnsetSettlementStatus() {
	o.SettlementStatus.Unset()
}

// GetEstimatedSettlementTime returns the EstimatedSettlementTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionResponse) GetEstimatedSettlementTime() time.Time {
	if o == nil || utils.IsNil(o.EstimatedSettlementTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EstimatedSettlementTime.Get()
}

// GetEstimatedSettlementTimeOk returns a tuple with the EstimatedSettlementTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionResponse) GetEstimatedSettlementTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedSettlementTime.Get(), o.EstimatedSettlementTime.IsSet()
}

// HasEstimatedSettlementTime returns a boolean if a field has been set.
func (o *TransactionResponse) HasEstimatedSettlementTime() bool {
	if o != nil && o.EstimatedSettlementTime.IsSet() {
		return true
	}

	return false
}

// SetEstimatedSettlementTime gets a reference to the given NullableTime and assigns it to the EstimatedSettlementTime field.
func (o *TransactionResponse) SetEstimatedSettlementTime(v time.Time) {
	o.EstimatedSettlementTime.Set(&v)
}
// SetEstimatedSettlementTimeNil sets the value for EstimatedSettlementTime to be an explicit nil
func (o *TransactionResponse) SetEstimatedSettlementTimeNil() {
	o.EstimatedSettlementTime.Set(nil)
}

// UnsetEstimatedSettlementTime ensures that no value is present for EstimatedSettlementTime, not even an explicit nil
func (o *TransactionResponse) UnsetEstimatedSettlementTime() {
	o.EstimatedSettlementTime.Unset()
}

// GetBusinessId returns the BusinessId field value
func (o *TransactionResponse) GetBusinessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetBusinessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessId, true
}

// SetBusinessId sets field value
func (o *TransactionResponse) SetBusinessId(v string) {
	o.BusinessId = v
}

// GetFee returns the Fee field value
func (o *TransactionResponse) GetFee() FeeResponse {
	if o == nil {
		var ret FeeResponse
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetFeeOk() (*FeeResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *TransactionResponse) SetFee(v FeeResponse) {
	o.Fee = v
}

// GetCreated returns the Created field value
func (o *TransactionResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *TransactionResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetUpdated returns the Updated field value
func (o *TransactionResponse) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *TransactionResponse) SetUpdated(v time.Time) {
	o.Updated = v
}

func (o TransactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["product_id"] = o.ProductId
	toSerialize["type"] = o.Type
	toSerialize["status"] = o.Status
	toSerialize["channel_category"] = o.ChannelCategory
	toSerialize["channel_code"] = o.ChannelCode.Get()

	toSerialize["account_identifier"] = o.AccountIdentifier.Get()

	toSerialize["reference_id"] = o.ReferenceId
	toSerialize["currency"] = o.Currency
	toSerialize["amount"] = o.Amount
	toSerialize["cashflow"] = o.Cashflow
    if o.Cashflow != "MONEY_IN" && o.Cashflow != "MONEY_OUT" {
        toSerialize["cashflow"] = nil
        return toSerialize, utils.NewError("invalid value for Cashflow when marshalling to JSON, must be one of MONEY_IN, MONEY_OUT")
    }
	if o.SettlementStatus.IsSet() {
		toSerialize["settlement_status"] = o.SettlementStatus.Get()
        if o.SettlementStatus.Get() != nil && (*o.SettlementStatus.Get() != "PENDING" && *o.SettlementStatus.Get() != "SETTLED") {
            toSerialize["settlement_status"] = nil
            return toSerialize, utils.NewError("invalid value for SettlementStatus when marshalling to JSON, must be one of PENDING, SETTLED")
        }
    }
	if o.EstimatedSettlementTime.IsSet() {
		toSerialize["estimated_settlement_time"] = o.EstimatedSettlementTime.Get()
    }
	toSerialize["business_id"] = o.BusinessId
	toSerialize["fee"] = o.Fee
	toSerialize["created"] = o.Created
	toSerialize["updated"] = o.Updated
	return toSerialize, nil
}

type NullableTransactionResponse struct {
	value *TransactionResponse
	isSet bool
}

func (v NullableTransactionResponse) Get() *TransactionResponse {
	return v.value
}

func (v *NullableTransactionResponse) Set(val *TransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionResponse(val *TransactionResponse) *NullableTransactionResponse {
	return &NullableTransactionResponse{value: val, isSet: true}
}

func (v NullableTransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


