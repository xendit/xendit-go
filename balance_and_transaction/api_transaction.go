package balance_and_transaction

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strconv"

	common "github.com/xendit/xendit-go/v6/common"
	utils "github.com/xendit/xendit-go/v6/utils"
	"strings"
	"reflect"
)


type TransactionApi interface {

	/*
	GetTransactionByID Get a transaction based on its id

	Get single specific transaction by transaction id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiGetTransactionByIDRequest
	*/
	GetTransactionByID(ctx context.Context, id string) ApiGetTransactionByIDRequest

	// GetTransactionByIDExecute executes the request
	//  @return TransactionResponse
	GetTransactionByIDExecute(r ApiGetTransactionByIDRequest) (*TransactionResponse, *http.Response, *common.XenditSdkError)

	/*
	GetAllTransactions Get a list of transactions

	Get a list of all transactions based on filter and search parameters.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAllTransactionsRequest
	*/
	GetAllTransactions(ctx context.Context) ApiGetAllTransactionsRequest

	// GetAllTransactionsExecute executes the request
	//  @return TransactionsResponse
	GetAllTransactionsExecute(r ApiGetAllTransactionsRequest) (*TransactionsResponse, *http.Response, *common.XenditSdkError)
}

// TransactionApiService TransactionApi service
type TransactionApiService struct {
	client common.IClient
}

// NewTransactionApi Create a new TransactionApi service
func NewTransactionApi (client common.IClient) TransactionApi {
	return &TransactionApiService{
		client: client,
	}
}


type ApiGetTransactionByIDRequest struct {
	ctx context.Context
	ApiService TransactionApi
	id string
	forUserId *string
}

// The sub-account user-id that you want to make this transaction for. This header is only used if you have access to xenPlatform. See xenPlatform for more information
func (r ApiGetTransactionByIDRequest) ForUserId(forUserId string) ApiGetTransactionByIDRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiGetTransactionByIDRequest) Execute() (*TransactionResponse, *http.Response, *common.XenditSdkError) {
	return r.ApiService.GetTransactionByIDExecute(r)
}

/*
GetTransactionByID Get a transaction based on its id

Get single specific transaction by transaction id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetTransactionByIDRequest
*/
func (a *TransactionApiService) GetTransactionByID(ctx context.Context, id string) ApiGetTransactionByIDRequest {
	return ApiGetTransactionByIDRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return TransactionResponse
func (a *TransactionApiService) GetTransactionByIDExecute(r ApiGetTransactionByIDRequest) (*TransactionResponse, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *TransactionResponse
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "TransactionApiService.GetTransactionByID")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: TransactionApiService.GetTransactionByIDExecute")
	}

	localVarPath := localBasePath + "/transactions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(utils.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := utils.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := utils.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.forUserId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarHeaderParams, "for-user-id", r.forUserId, "")
	}
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: TransactionApiService.GetTransactionByIDExecute")
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)

	localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))

    err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil || localVarHTTPResponse.StatusCode < 200 || localVarHTTPResponse.StatusCode >= 300 {
		xenditSdkError := common.NewXenditSdkError(&localVarBody, strconv.Itoa(localVarHTTPResponse.StatusCode), localVarHTTPResponse.Status)

		return localVarReturnValue, localVarHTTPResponse, xenditSdkError
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAllTransactionsRequest struct {
	ctx context.Context
	ApiService TransactionApi
	forUserId *string
	types *[]TransactionTypes
	statuses *[]TransactionStatuses
	channelCategories *[]ChannelsCategories
	referenceId *string
	productId *string
	accountIdentifier *string
	amount *float32
	currency *Currency
	created *DateRangeFilter
	updated *DateRangeFilter
	limit *float32
	afterId *string
	beforeId *string
}

// The sub-account user-id that you want to make this transaction for. This header is only used if you have access to xenPlatform. See xenPlatform for more information
func (r ApiGetAllTransactionsRequest) ForUserId(forUserId string) ApiGetAllTransactionsRequest {
	r.forUserId = &forUserId
	return r
}

// Transaction types that will be included in the result. Default is to include all transaction types
func (r ApiGetAllTransactionsRequest) Types(types []TransactionTypes) ApiGetAllTransactionsRequest {
	r.types = &types
	return r
}

// Status of the transaction. Default is to include all status.
func (r ApiGetAllTransactionsRequest) Statuses(statuses []TransactionStatuses) ApiGetAllTransactionsRequest {
	r.statuses = &statuses
	return r
}

// Payment channels in which the transaction is carried out. Default is to include all channels.
func (r ApiGetAllTransactionsRequest) ChannelCategories(channelCategories []ChannelsCategories) ApiGetAllTransactionsRequest {
	r.channelCategories = &channelCategories
	return r
}

// To filter the result for transactions with matching reference given (case sensitive)
func (r ApiGetAllTransactionsRequest) ReferenceId(referenceId string) ApiGetAllTransactionsRequest {
	r.referenceId = &referenceId
	return r
}

// To filter the result for transactions with matching product_id (a.k.a payment_id) given (case sensitive)
func (r ApiGetAllTransactionsRequest) ProductId(productId string) ApiGetAllTransactionsRequest {
	r.productId = &productId
	return r
}

// Account identifier of transaction. The format will be different from each channel. For example, on &#x60;BANK&#x60; channel it will be account number and on &#x60;CARD&#x60; it will be masked card number.
func (r ApiGetAllTransactionsRequest) AccountIdentifier(accountIdentifier string) ApiGetAllTransactionsRequest {
	r.accountIdentifier = &accountIdentifier
	return r
}

// Specific transaction amount to search for
func (r ApiGetAllTransactionsRequest) Amount(amount float32) ApiGetAllTransactionsRequest {
	r.amount = &amount
	return r
}

func (r ApiGetAllTransactionsRequest) Currency(currency Currency) ApiGetAllTransactionsRequest {
	r.currency = &currency
	return r
}

// Filter time of transaction by created date. If not specified will list all dates.
func (r ApiGetAllTransactionsRequest) Created(created DateRangeFilter) ApiGetAllTransactionsRequest {
	r.created = &created
	return r
}

// Filter time of transaction by updated date. If not specified will list all dates.
func (r ApiGetAllTransactionsRequest) Updated(updated DateRangeFilter) ApiGetAllTransactionsRequest {
	r.updated = &updated
	return r
}

// number of items in the result per page. Another name for \&quot;results_per_page\&quot;
func (r ApiGetAllTransactionsRequest) Limit(limit float32) ApiGetAllTransactionsRequest {
	r.limit = &limit
	return r
}

func (r ApiGetAllTransactionsRequest) AfterId(afterId string) ApiGetAllTransactionsRequest {
	r.afterId = &afterId
	return r
}

func (r ApiGetAllTransactionsRequest) BeforeId(beforeId string) ApiGetAllTransactionsRequest {
	r.beforeId = &beforeId
	return r
}

func (r ApiGetAllTransactionsRequest) Execute() (*TransactionsResponse, *http.Response, *common.XenditSdkError) {
	return r.ApiService.GetAllTransactionsExecute(r)
}

/*
GetAllTransactions Get a list of transactions

Get a list of all transactions based on filter and search parameters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllTransactionsRequest
*/
func (a *TransactionApiService) GetAllTransactions(ctx context.Context) ApiGetAllTransactionsRequest {
	return ApiGetAllTransactionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TransactionsResponse
func (a *TransactionApiService) GetAllTransactionsExecute(r ApiGetAllTransactionsRequest) (*TransactionsResponse, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *TransactionsResponse
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "TransactionApiService.GetAllTransactions")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: TransactionApiService.GetAllTransactionsExecute")
	}

	localVarPath := localBasePath + "/transactions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.types != nil {
		t := *r.types
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "types", s.Index(i), "multi")
			}
		} else {
			utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "types", t, "multi")
		}
	}
	if r.statuses != nil {
		t := *r.statuses
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "statuses", s.Index(i), "multi")
			}
		} else {
			utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "statuses", t, "multi")
		}
	}
	if r.channelCategories != nil {
		t := *r.channelCategories
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "channel_categories", s.Index(i), "multi")
			}
		} else {
			utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "channel_categories", t, "multi")
		}
	}
	if r.referenceId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "reference_id", r.referenceId, "")
	}
	if r.productId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "product_id", r.productId, "")
	}
	if r.accountIdentifier != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "account_identifier", r.accountIdentifier, "")
	}
	if r.amount != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "")
	}
	if r.currency != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "currency", r.currency, "")
	}
	if r.created != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "created", r.created, "")
	}
	if r.updated != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "updated", r.updated, "")
	}
	if r.limit != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.afterId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "after_id", r.afterId, "")
	}
	if r.beforeId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "before_id", r.beforeId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := utils.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := utils.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.forUserId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarHeaderParams, "for-user-id", r.forUserId, "")
	}
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: TransactionApiService.GetAllTransactionsExecute")
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)

	localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))

    err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil || localVarHTTPResponse.StatusCode < 200 || localVarHTTPResponse.StatusCode >= 300 {
		xenditSdkError := common.NewXenditSdkError(&localVarBody, strconv.Itoa(localVarHTTPResponse.StatusCode), localVarHTTPResponse.Status)

		return localVarReturnValue, localVarHTTPResponse, xenditSdkError
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
