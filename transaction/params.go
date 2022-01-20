package transaction

// GetTransactionParams contains parameters for GetTransaction
type GetTransactionParams struct {
	ForUserID     string `json:"-"`
	TransactionID string `json:"transaction_id" validate:"required"`
}

// GetListTransactionParams contains parameters for GetListTransaction
type GetListTransactionParams struct {
	ForUserID         string   `url:"-"`
	Types             []string `url:"types,omitempty"`
	Statuses          []string `url:"statuses,omitempty"`
	ChannelCategories []string `url:"channel_categories,omitempty"`
	ReferenceID       string   `url:"reference_id,omitempty"`
	ProductID         string   `url:"product_id,omitempty"`
	AccountIdentifier string   `url:"account_identifier,omitempty"`
	Currency          string   `url:"currency,omitempty"`
	Amount            float64  `url:"Amount,omitempty"`
	CreatedGte        string   `url:"created[gte],omitempty"`
	CreatedLte        string   `url:"created[lte],omitempty"`
	UpdatedGte        string   `url:"updated[gte],omitempty"`
	UpdatedLte        string   `url:"updated[lte],omitempty"`
	Limit             int      `url:"limit,omitempty"`
	AfterID           string   `url:"after_id,omitempty"`
	BeforeID          string   `url:"before_id,omitempty"`
}
