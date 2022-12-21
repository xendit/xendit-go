package constant

type PaymentMethodTypeEnum string
type ReusabilityEnum string
type CountryEnum string
type CurrencyEnum string
type TokenStatus string

// This consists the values that PaymentMethodTypeEnum can take
const (
	PaymentMethodTypeEwallet        PaymentMethodTypeEnum = "EWALLET"
	PaymentMethodTypeDirectDebit    PaymentMethodTypeEnum = "DIRECT_DEBIT"
	PaymentMethodTypeOverTheCounter PaymentMethodTypeEnum = "OVER_THE_COUNTER"
	PaymentMethodTypeQRCode         PaymentMethodTypeEnum = "QR_CODE"
	PaymentMethodTypeVirtualAccount PaymentMethodTypeEnum = "VIRTUAL_ACCOUNT"
	PaymentMethodTypeCard           PaymentMethodTypeEnum = "CARD"

	ReusabilityOneTimeUse  ReusabilityEnum = "ONE_TIME_USE"
	ReusabilityMultipleUse ReusabilityEnum = "MULTIPLE_USE"

	CountryPH CountryEnum = "PH"
	CountryID CountryEnum = "ID"
	CountryVN CountryEnum = "VN"
	CountryTH CountryEnum = "TH"

	CurrencyPH  CurrencyEnum = "PHP"
	CurrencyIDR CurrencyEnum = "IDR"
	CurrencyVND CurrencyEnum = "VND"
	CurrencyTHB CurrencyEnum = "THB"
)

type PaymentMethodStatusEnum string

const (
	Active         PaymentMethodStatusEnum = "ACTIVE"
	Expired        PaymentMethodStatusEnum = "EXPIRED"
	Inactive       PaymentMethodStatusEnum = "INACTIVE"
	Pending        PaymentMethodStatusEnum = "PENDING"
	RequiresAction PaymentMethodStatusEnum = "REQUIRES_ACTION"
	Failed         PaymentMethodStatusEnum = "FAILED"
)
