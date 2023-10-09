/*
Transaction Service V4 API

# Introduction This specification describes how to use the Transaction Service V4 API.   **Transaction Service** is the service that records the customer transactions and is responsible to calculate their balance.  All products that move customer money around whether it is money-in, money-out, or transfer will interact with the Transaction Service on its flow. Transaction Service is the source of truth of Xendit and Customer regarding how much money that customer has that is stored in Xendit. Transaction Service is the source that is used for both our internal and customer financial reconciliation. Internally, the Transaction Service data structure is similar to how double-entry accounting works.  ## How Xendit teams/services do integrate with Transaction Service V4   **Channel product team/service** They interact with the Transaction Service when they want to record the transactions. This transaction can be money-in (balance added), money-out (balance deducted), transfer, refund/void/reversal, or other kind of transaction that affects customer balance. Product team also interacts with the Transaction Service for getting information about the transaction or balance.  **Billing/Fee team/service** They interact with Transaction Service either as the dependency of Transaction Service for getting the correct fee calculation/settings. Or using Transaction Service for getting the transaction/fee information to calculate the bill for the customer.  **NUX team/service** They interact with the Transaction Service to set up the customer ledger_account that is used to record their transactions.  **Finance team/service** They interact with the Transaction Service to get the transaction and balance data for each customer to do reconciliation.  **Dashboard/API team/service** They interact with the Transaction Service as a proxy to show the data to the Customer.  ## Prerequisites  Before staring to use **Transaction Service API** you need to complete a few things: 1. Find out **Base URL** for the API. Every endpoint definition in this document contains list of available servers (local, staging, production) 2. Set up ledger accounts using business id and currency. **Ledger Account** represents the account of the customer that will be used to associate with ledger lines. Each business may have at least 1 ledger account group (a group consists of a few accounts of types such as cash, liability, holding), and the money movement of their ledger will revolve around those ledger accounts. **Ledger Lines** that show a debit or credit transaction for a ledger account. We’re using the double-entry principle in accounting where we should post 2 lines every time we make a transaction, 1 to debit an account and 1 to credit another account. See how to call <a href=\"#operation/setupLedgerAccounts\">Create cash, liability, holding, and tax account for a business (api/ledger-accounts/setup)</a> section of this document 3. To be able to create payments with fee/VAT the Product rate settings and VAT rate settings should be created using Transaction Fee Service. See <a href=\"https://docs.google.com/document/d/1HrrA4GhWD1DaJS5dn0dh9VyMhLR6TOUMW1qhRUZ9d_k/edit?pli=1#heading=h.518me3lwf8rb\">Fee Service Documentation</a> for details about how to create Product/VAT rate settings.   ## Transaction flows  To integrate with the Transaction Service you should decide what types of transaction flows your integration will be using. Transaction flow is set by the transaction `type` during transaction creation  1. Money In flows     1. Payment from credit card           `type: CREDIT_CARD_PAYMENT`           3. Payment from other sources without fee/VAT           `type: DEPOSIT, FOREX_DEPOSIT, ISSUING_FUNDING_REFUND, BNPL_PARTNER_SETTLEMENT_CREDIT, PROMO_FEE_CASHBACK, PROMO_VAT_CASHBACK, BATCH_VA_PAYMENT`           4. Payment from other sources with fee/VAT           `type: VA_PAYMENT, IM_ESCROW_VA_PAYMENT, IM_DEPOSIT, RO_PAYMENT, EWALLET_PAYMENT, CARDLESS_CREDIT_PAYMENT, IM_REMITTANCE_VA_PAYMENT, PAYLATER_PAYMENT, INVOICE, QR_CODE_PAYMENT, DIRECT_DEBIT_PAYMENT, DIRECT_BANK_TRANSFER, ACH_PAYMENT, CRYPTO_PAYMENT`           5. Billing deposit from cash           `type: BILLING_DEPOSIT`           6. Billing deposit from other sources           `type: BILLING_DIRECT_DEPOSIT, BILLING_VA_DIRECT_DEPOSIT`       2. Money out flows     1. Instant payment           `type: simple money out types`              `status: COMPLETED`           2. Simple payment without fee/VAT           `type: CHARGEBACK_DEDUCTION, FRAUD_DEDUCTION, LOAN_REPAYMENT, FOREX_DEDUCTION, BNPL_PARTNER_SETTLEMENT_DEBIT, WITHDRAWAL`       3. Simple payment with fee/VAT           `type: ISSUING_FUNDING, BATCH_DISBURSEMENT, CASH_DISBURSEMENT, DISBURSEMENT, REMITTANCE, REMITTANCE_PAYOUT, TAX_DISBURSEMENT`           4. Billing withdraw to cash           `type: BILLING_WITHDRAWAL`           4. Billing withdraw to other destinations           `type: BILL_PAYMENT`       3. Reversal flow      Some of transactions could be reversed. See <a href=\"#section/Introduction/Reversible-non-reversible-transaction-types\">Reversible / non reversible transaction types</a> section of this document. To reverse transaction you should call <a href=\"#operation/updateTransaction\">Update transaction (/api/transactions/:id)</a>  endpoint with the transaction status `REVERSED`.    4. Void/Cancellation Flow      Transaction in the `PENDING_SETTLEMENT` status could be canceled. To do that you should call <a href=\"#operation/updateTransaction\">Update transaction (/api/transactions/:id)</a>  endpoint with the transaction status `VOIDED`.       5. Switcher flow      Switchers are transactions that do not affect the customer balance. These are transactions that goes directly to the customers’ account and simply passes through Xendit. Therefore, it will not impact the customer balance and we will only charge Fee and VAT. To create switcher flow you should set `is_switcher_payment` field to `true`.       ## Instant/non instant settlement  Transactions can be performed instantly (instant settlement) or with delay (non instant settlement).  Some of the transaction types are only instantly processed, some of them support both instant and non instant settlement and some of them have only non instant settlement. If settlement is instant than balance will be changed instantly. In opposite case the transaction status has to be set into PENDING_SETTLEMENT and settlement date should be provided.   1. Instant settlement Money In transaction types      `DEPOSIT, BATCH_VA_PAYMENT, FOREX_DEPOSIT, IM_DEPOSIT, CARDLESS_CREDIT_PAYMENT, ISSUING_FUNDING_REFUND, BNPL_PARTNER_SETTLEMENT_CREDIT, PROMO_FEE_CASHBACK, PROMO_VAT_CASHBACK, REMITTANCE_VA_PAYMENT_CLAIM`    2. Both instant and non instant Money In transaction types      `DIRECT_DEBIT_PAYMENT, DIRECT_BANK_TRANSFER, ACH_PAYMENT, RO_PAYMENT, EWALLET_PAYMENT, QR_CODE_PAYMENT, VA_PAYMENT, INVOICE, PAYLATER_PAYMENT`  3. Non Instant settlement Money In transaction types      `CREDIT_CARD_PAYMENT`    4. Instant settlement Money Out transaction types      `LOAN_REPAYMENT, FOREX_DEDUCTION, BILL_PAYMENT, ISSUING_FUNDING, BNPL_PARTNER_SETTLEMENT_DEBIT, FRAUD_DEDUCTION`  5. Both instant and non instant settlement supported Money Out transaction types      `CHARGEBACK_DEDUCTION`  6. Non Instant settlement Money Out transaction types      All other money out types are non instant settlement  ## Reversible / non reversible transaction types  Some transactions can be reversed. Here are the list of transaction types that could be reversed:   `CASH_DISBURSEMENT, DISBURSEMENT, BATCH_DISBURSEMENT, REMITTANCE, REMITTANCE_PAYOUT, TAX_DISBURSEMENT, WITHDRAWAL, DEPOSIT, FOREX_DEPOSIT, FOREX_DEDUCTION, VA_PAYMENT, BATCH_VA_PAYMENT, IM_REMITTANCE_VA_PAYMENT, IM_ESCROW_VA_PAYMENT, IM_DEPOSIT, REMITTANCE_VA_PAYMENT, REMITTANCE_VA_PAYMENT_CLAIM, RO_PAYMENT, CARDLESS_CREDIT_PAYMENT, PAYLATER_PAYMENT, INVOICE, QR_CODE_PAYMENT, CREDIT_CARD_PAYMENT, EWALLET_PAYMENT, DIRECT_DEBIT_PAYMENT, DIRECT_BANK_TRANSFER, ACH_PAYMENT, CHARGEBACK_DEDUCTION, FRAUD_DEDUCTION, LOAN_REPAYMENT, ISSUING_FUNDING, ISSUING_FUNDING_REFUND, BNPL_PARTNER_SETTLEMENT_DEBIT, BNPL_PARTNER_SETTLEMENT_CREDIT, BILLING_DEPOSIT, BILLING_DIRECT_DEPOSIT, BILLING_VA_DIRECT_DEPOSIT, BILLING_WITHDRAWAL, BILL_PAYMENT, PROMO_FEE_CASHBACK, PROMO_VAT_CASHBACK`       ## How to create transaction  After you created or already have the `BUSINESS_CASH` ledger account ID (See <a href=\"#section/Introduction/Prerequisites\">Prerequisites</a> section) and you know what transaction flows are going to be used  you can create the new transaction using POST request to the  <a href=\"#operation/createTransaction\">Create a new transaction (/api/transactions)</a>  endpoint  ## How to update transaction  To update transaction you should do  PATCH request to the  <a href=\"#operation/updateTransaction\">Update transaction (/api/transactions/::id)</a>  endpoint   

API version: 3.4.3
*/


package balance_and_transaction

import (
	"encoding/json"
	
	"fmt"
)

// Currency the model 'Currency'
type Currency string

// List of Currency
const (
	CURRENCY_IDR Currency = "IDR"
	CURRENCY_PHP Currency = "PHP"
	CURRENCY_USD Currency = "USD"
	CURRENCY_JPY Currency = "JPY"
	CURRENCY_VND Currency = "VND"
	CURRENCY_SGD Currency = "SGD"
	CURRENCY_AED Currency = "AED"
	CURRENCY_AFN Currency = "AFN"
	CURRENCY_ALL Currency = "ALL"
	CURRENCY_AMD Currency = "AMD"
	CURRENCY_ANG Currency = "ANG"
	CURRENCY_AOA Currency = "AOA"
	CURRENCY_ARS Currency = "ARS"
	CURRENCY_AUD Currency = "AUD"
	CURRENCY_AWG Currency = "AWG"
	CURRENCY_AZN Currency = "AZN"
	CURRENCY_BAM Currency = "BAM"
	CURRENCY_BBD Currency = "BBD"
	CURRENCY_BDT Currency = "BDT"
	CURRENCY_BGN Currency = "BGN"
	CURRENCY_BHD Currency = "BHD"
	CURRENCY_BIF Currency = "BIF"
	CURRENCY_BMD Currency = "BMD"
	CURRENCY_BND Currency = "BND"
	CURRENCY_BOB Currency = "BOB"
	CURRENCY_BRL Currency = "BRL"
	CURRENCY_BSD Currency = "BSD"
	CURRENCY_BTN Currency = "BTN"
	CURRENCY_BWP Currency = "BWP"
	CURRENCY_BYN Currency = "BYN"
	CURRENCY_BZD Currency = "BZD"
	CURRENCY_CAD Currency = "CAD"
	CURRENCY_CDF Currency = "CDF"
	CURRENCY_CHF Currency = "CHF"
	CURRENCY_CLP Currency = "CLP"
	CURRENCY_CNY Currency = "CNY"
	CURRENCY_COP Currency = "COP"
	CURRENCY_CRC Currency = "CRC"
	CURRENCY_CUC Currency = "CUC"
	CURRENCY_CUP Currency = "CUP"
	CURRENCY_CVE Currency = "CVE"
	CURRENCY_CZK Currency = "CZK"
	CURRENCY_DJF Currency = "DJF"
	CURRENCY_DKK Currency = "DKK"
	CURRENCY_DOP Currency = "DOP"
	CURRENCY_DZD Currency = "DZD"
	CURRENCY_EGP Currency = "EGP"
	CURRENCY_ERN Currency = "ERN"
	CURRENCY_ETB Currency = "ETB"
	CURRENCY_EUR Currency = "EUR"
	CURRENCY_FJD Currency = "FJD"
	CURRENCY_FKP Currency = "FKP"
	CURRENCY_GBP Currency = "GBP"
	CURRENCY_GEL Currency = "GEL"
	CURRENCY_GGP Currency = "GGP"
	CURRENCY_GHS Currency = "GHS"
	CURRENCY_GIP Currency = "GIP"
	CURRENCY_GMD Currency = "GMD"
	CURRENCY_GNF Currency = "GNF"
	CURRENCY_GTQ Currency = "GTQ"
	CURRENCY_GYD Currency = "GYD"
	CURRENCY_HKD Currency = "HKD"
	CURRENCY_HNL Currency = "HNL"
	CURRENCY_HRK Currency = "HRK"
	CURRENCY_HTG Currency = "HTG"
	CURRENCY_HUF Currency = "HUF"
	CURRENCY_ILS Currency = "ILS"
	CURRENCY_IMP Currency = "IMP"
	CURRENCY_INR Currency = "INR"
	CURRENCY_IQD Currency = "IQD"
	CURRENCY_IRR Currency = "IRR"
	CURRENCY_ISK Currency = "ISK"
	CURRENCY_JEP Currency = "JEP"
	CURRENCY_JMD Currency = "JMD"
	CURRENCY_JOD Currency = "JOD"
	CURRENCY_KES Currency = "KES"
	CURRENCY_KGS Currency = "KGS"
	CURRENCY_KHR Currency = "KHR"
	CURRENCY_KMF Currency = "KMF"
	CURRENCY_KPW Currency = "KPW"
	CURRENCY_KRW Currency = "KRW"
	CURRENCY_KWD Currency = "KWD"
	CURRENCY_KYD Currency = "KYD"
	CURRENCY_KZT Currency = "KZT"
	CURRENCY_LAK Currency = "LAK"
	CURRENCY_LBP Currency = "LBP"
	CURRENCY_LKR Currency = "LKR"
	CURRENCY_LRD Currency = "LRD"
	CURRENCY_LSL Currency = "LSL"
	CURRENCY_LYD Currency = "LYD"
	CURRENCY_MAD Currency = "MAD"
	CURRENCY_MDL Currency = "MDL"
	CURRENCY_MGA Currency = "MGA"
	CURRENCY_MKD Currency = "MKD"
	CURRENCY_MMK Currency = "MMK"
	CURRENCY_MNT Currency = "MNT"
	CURRENCY_MOP Currency = "MOP"
	CURRENCY_MRU Currency = "MRU"
	CURRENCY_MUR Currency = "MUR"
	CURRENCY_MVR Currency = "MVR"
	CURRENCY_MWK Currency = "MWK"
	CURRENCY_MXN Currency = "MXN"
	CURRENCY_MYR Currency = "MYR"
	CURRENCY_MZN Currency = "MZN"
	CURRENCY_NAD Currency = "NAD"
	CURRENCY_NGN Currency = "NGN"
	CURRENCY_NIO Currency = "NIO"
	CURRENCY_NOK Currency = "NOK"
	CURRENCY_NPR Currency = "NPR"
	CURRENCY_NZD Currency = "NZD"
	CURRENCY_OMR Currency = "OMR"
	CURRENCY_PAB Currency = "PAB"
	CURRENCY_PEN Currency = "PEN"
	CURRENCY_PGK Currency = "PGK"
	CURRENCY_PKR Currency = "PKR"
	CURRENCY_PLN Currency = "PLN"
	CURRENCY_PYG Currency = "PYG"
	CURRENCY_QAR Currency = "QAR"
	CURRENCY_RON Currency = "RON"
	CURRENCY_RSD Currency = "RSD"
	CURRENCY_RUB Currency = "RUB"
	CURRENCY_RWF Currency = "RWF"
	CURRENCY_SAR Currency = "SAR"
	CURRENCY_SBD Currency = "SBD"
	CURRENCY_SCR Currency = "SCR"
	CURRENCY_SDG Currency = "SDG"
	CURRENCY_SEK Currency = "SEK"
	CURRENCY_SHP Currency = "SHP"
	CURRENCY_SLL Currency = "SLL"
	CURRENCY_SOS Currency = "SOS"
	CURRENCY_SPL Currency = "SPL"
	CURRENCY_SRD Currency = "SRD"
	CURRENCY_STN Currency = "STN"
	CURRENCY_SVC Currency = "SVC"
	CURRENCY_SYP Currency = "SYP"
	CURRENCY_SZL Currency = "SZL"
	CURRENCY_THB Currency = "THB"
	CURRENCY_TJS Currency = "TJS"
	CURRENCY_TMT Currency = "TMT"
	CURRENCY_TND Currency = "TND"
	CURRENCY_TOP Currency = "TOP"
	CURRENCY_TRY Currency = "TRY"
	CURRENCY_TTD Currency = "TTD"
	CURRENCY_TVD Currency = "TVD"
	CURRENCY_TWD Currency = "TWD"
	CURRENCY_TZS Currency = "TZS"
	CURRENCY_UAH Currency = "UAH"
	CURRENCY_UGX Currency = "UGX"
	CURRENCY_UYU Currency = "UYU"
	CURRENCY_UZS Currency = "UZS"
	CURRENCY_VEF Currency = "VEF"
	CURRENCY_VUV Currency = "VUV"
	CURRENCY_WST Currency = "WST"
	CURRENCY_XAF Currency = "XAF"
	CURRENCY_XCD Currency = "XCD"
	CURRENCY_XDR Currency = "XDR"
	CURRENCY_XOF Currency = "XOF"
	CURRENCY_XPF Currency = "XPF"
	CURRENCY_YER Currency = "YER"
	CURRENCY_ZAR Currency = "ZAR"
	CURRENCY_ZMW Currency = "ZMW"
	CURRENCY_ZWD Currency = "ZWD"
)

// All allowed values of Currency enum
var AllowedCurrencyEnumValues = []Currency{
	"IDR",
	"PHP",
	"USD",
	"JPY",
	"VND",
	"SGD",
	"AED",
	"AFN",
	"ALL",
	"AMD",
	"ANG",
	"AOA",
	"ARS",
	"AUD",
	"AWG",
	"AZN",
	"BAM",
	"BBD",
	"BDT",
	"BGN",
	"BHD",
	"BIF",
	"BMD",
	"BND",
	"BOB",
	"BRL",
	"BSD",
	"BTN",
	"BWP",
	"BYN",
	"BZD",
	"CAD",
	"CDF",
	"CHF",
	"CLP",
	"CNY",
	"COP",
	"CRC",
	"CUC",
	"CUP",
	"CVE",
	"CZK",
	"DJF",
	"DKK",
	"DOP",
	"DZD",
	"EGP",
	"ERN",
	"ETB",
	"EUR",
	"FJD",
	"FKP",
	"GBP",
	"GEL",
	"GGP",
	"GHS",
	"GIP",
	"GMD",
	"GNF",
	"GTQ",
	"GYD",
	"HKD",
	"HNL",
	"HRK",
	"HTG",
	"HUF",
	"ILS",
	"IMP",
	"INR",
	"IQD",
	"IRR",
	"ISK",
	"JEP",
	"JMD",
	"JOD",
	"KES",
	"KGS",
	"KHR",
	"KMF",
	"KPW",
	"KRW",
	"KWD",
	"KYD",
	"KZT",
	"LAK",
	"LBP",
	"LKR",
	"LRD",
	"LSL",
	"LYD",
	"MAD",
	"MDL",
	"MGA",
	"MKD",
	"MMK",
	"MNT",
	"MOP",
	"MRU",
	"MUR",
	"MVR",
	"MWK",
	"MXN",
	"MYR",
	"MZN",
	"NAD",
	"NGN",
	"NIO",
	"NOK",
	"NPR",
	"NZD",
	"OMR",
	"PAB",
	"PEN",
	"PGK",
	"PKR",
	"PLN",
	"PYG",
	"QAR",
	"RON",
	"RSD",
	"RUB",
	"RWF",
	"SAR",
	"SBD",
	"SCR",
	"SDG",
	"SEK",
	"SHP",
	"SLL",
	"SOS",
	"SPL",
	"SRD",
	"STN",
	"SVC",
	"SYP",
	"SZL",
	"THB",
	"TJS",
	"TMT",
	"TND",
	"TOP",
	"TRY",
	"TTD",
	"TVD",
	"TWD",
	"TZS",
	"UAH",
	"UGX",
	"UYU",
	"UZS",
	"VEF",
	"VUV",
	"WST",
	"XAF",
	"XCD",
	"XDR",
	"XOF",
	"XPF",
	"YER",
	"ZAR",
	"ZMW",
	"ZWD",
}

func (v *Currency) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Currency(value)
	for _, existing := range AllowedCurrencyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Currency", value)
}

// NewCurrencyFromValue returns a pointer to a valid Currency
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCurrencyFromValue(v string) (*Currency, error) {
	ev := Currency(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Currency: valid values are %v", v, AllowedCurrencyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Currency) IsValid() bool {
	for _, existing := range AllowedCurrencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v Currency) String() string {
	return string(v)
}

// Ptr returns reference to Currency value
func (v Currency) Ptr() *Currency {
	return &v
}

type NullableCurrency struct {
	value *Currency
	isSet bool
}

func (v NullableCurrency) Get() *Currency {
	return v.value
}

func (v *NullableCurrency) Set(val *Currency) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrency) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrency(val *Currency) *NullableCurrency {
	return &NullableCurrency{value: val, isSet: true}
}

func (v NullableCurrency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
