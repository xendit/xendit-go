package balance_and_transaction

import (
	"bytes"
	"context"
    "fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	common "github.com/xendit/xendit-go/v6/common"
	utils "github.com/xendit/xendit-go/v6/utils"
	"time"
)


type BalanceApi interface {

	/*
	GetBalance Retrieves balances for a business, default to CASH type

	Retrieves balance for your business, defaults to CASH type

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetBalanceRequest
	*/
	GetBalance(ctx context.Context) ApiGetBalanceRequest

	// GetBalanceExecute executes the request
	//  @return Balance
	GetBalanceExecute(r ApiGetBalanceRequest) (*Balance, *http.Response, *common.XenditSdkError)
}

// BalanceApiService BalanceApi service
type BalanceApiService struct {
	client common.IClient
}

// NewBalanceApi Create a new BalanceApi service
func NewBalanceApi (client common.IClient) BalanceApi {
	return &BalanceApiService{
		client: client,
	}
}


type ApiGetBalanceRequest struct {
	ctx context.Context
	ApiService BalanceApi
	accountType *string
	currency *string
	atTimestamp *time.Time
	forUserId *string
}

// The selected balance type
func (r ApiGetBalanceRequest) AccountType(accountType string) ApiGetBalanceRequest {
	r.accountType = &accountType
	return r
}

// Currency for filter for customers with multi currency accounts
func (r ApiGetBalanceRequest) Currency(currency string) ApiGetBalanceRequest {
	r.currency = &currency
	return r
}

// The timestamp you want to use as the limit for balance retrieval
func (r ApiGetBalanceRequest) AtTimestamp(atTimestamp time.Time) ApiGetBalanceRequest {
	r.atTimestamp = &atTimestamp
	return r
}

// The sub-account user-id that you want to make this transaction for. This header is only used if you have access to xenPlatform. See xenPlatform for more information
func (r ApiGetBalanceRequest) ForUserId(forUserId string) ApiGetBalanceRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiGetBalanceRequest) Execute() (*Balance, *http.Response, *common.XenditSdkError) {
	return r.ApiService.GetBalanceExecute(r)
}

/*
GetBalance Retrieves balances for a business, default to CASH type

Retrieves balance for your business, defaults to CASH type

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetBalanceRequest
*/
func (a *BalanceApiService) GetBalance(ctx context.Context) ApiGetBalanceRequest {
	return ApiGetBalanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Balance
func (a *BalanceApiService) GetBalanceExecute(r ApiGetBalanceRequest) (*Balance, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *Balance
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "BalanceApiService.GetBalance")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: BalanceApiService.GetBalanceExecute")
	}

	localVarPath := localBasePath + "/balance"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.accountType != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "account_type", r.accountType, "")
	}
	if r.currency != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "currency", r.currency, "")
	}
	if r.atTimestamp != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "at_timestamp", r.atTimestamp, "")
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
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: BalanceApiService.GetBalanceExecute")
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
    if err != nil {
        return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", fmt.Sprintf("Error creating HTTP request: BalanceApiService.GetBalanceExecute: %v", err))
    }
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))

    err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil || localVarHTTPResponse.StatusCode < 200 || localVarHTTPResponse.StatusCode >= 300 {
		xenditSdkError := common.NewXenditSdkError(&localVarBody, strconv.Itoa(localVarHTTPResponse.StatusCode), localVarHTTPResponse.Status)

		return localVarReturnValue, localVarHTTPResponse, xenditSdkError
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
