package payout

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strconv"

	common "github.com/xendit/xendit-go/v3/common"
	utils "github.com/xendit/xendit-go/v3/utils"
	"strings"
)


type PayoutApi interface {

	/*
	CancelPayout API to cancel requested payouts that have not yet been sent to partner banks and e-wallets. Cancellation is possible if the payout has not been sent out via our partner and when payout status is ACCEPTED.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Payout id returned from the response of /v2/payouts
	@return ApiCancelPayoutRequest
	*/
	CancelPayout(ctx context.Context, id string) ApiCancelPayoutRequest

	// CancelPayoutExecute executes the request
	//  @return GetPayouts200ResponseDataInner
	CancelPayoutExecute(r ApiCancelPayoutRequest) (*GetPayouts200ResponseDataInner, *http.Response, *common.XenditSdkError)

	/*
	CreatePayout API to send money at scale to bank accounts & eWallets

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreatePayoutRequest
	*/
	CreatePayout(ctx context.Context) ApiCreatePayoutRequest

	// CreatePayoutExecute executes the request
	//  @return GetPayouts200ResponseDataInner
	CreatePayoutExecute(r ApiCreatePayoutRequest) (*GetPayouts200ResponseDataInner, *http.Response, *common.XenditSdkError)

	/*
	GetPayoutById API to fetch the current status, or details of the payout

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Payout id returned from the response of /v2/payouts
	@return ApiGetPayoutByIdRequest
	*/
	GetPayoutById(ctx context.Context, id string) ApiGetPayoutByIdRequest

	// GetPayoutByIdExecute executes the request
	//  @return GetPayouts200ResponseDataInner
	GetPayoutByIdExecute(r ApiGetPayoutByIdRequest) (*GetPayouts200ResponseDataInner, *http.Response, *common.XenditSdkError)

	/*
	GetPayoutChannels API providing the current list of banks and e-wallets we support for payouts for both regions

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetPayoutChannelsRequest
	*/
	GetPayoutChannels(ctx context.Context) ApiGetPayoutChannelsRequest

	// GetPayoutChannelsExecute executes the request
	//  @return []Channel
	GetPayoutChannelsExecute(r ApiGetPayoutChannelsRequest) ([]Channel, *http.Response, *common.XenditSdkError)

	/*
	GetPayouts API to retrieve all matching payouts with reference ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetPayoutsRequest
	*/
	GetPayouts(ctx context.Context) ApiGetPayoutsRequest

	// GetPayoutsExecute executes the request
	//  @return GetPayouts200Response
	GetPayoutsExecute(r ApiGetPayoutsRequest) (*GetPayouts200Response, *http.Response, *common.XenditSdkError)
}

// PayoutApiService PayoutApi service
type PayoutApiService struct {
	client common.IClient
}

// NewPayoutApi Create a new PayoutApi service
func NewPayoutApi (client common.IClient) PayoutApi {
	return &PayoutApiService{
		client: client,
	}
}


type ApiCancelPayoutRequest struct {
	ctx context.Context
	ApiService PayoutApi
	id string
}

func (r ApiCancelPayoutRequest) Execute() (*GetPayouts200ResponseDataInner, *http.Response, *common.XenditSdkError) {
	return r.ApiService.CancelPayoutExecute(r)
}

/*
CancelPayout API to cancel requested payouts that have not yet been sent to partner banks and e-wallets. Cancellation is possible if the payout has not been sent out via our partner and when payout status is ACCEPTED.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Payout id returned from the response of /v2/payouts
 @return ApiCancelPayoutRequest
*/
func (a *PayoutApiService) CancelPayout(ctx context.Context, id string) ApiCancelPayoutRequest {
	return ApiCancelPayoutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GetPayouts200ResponseDataInner
func (a *PayoutApiService) CancelPayoutExecute(r ApiCancelPayoutRequest) (*GetPayouts200ResponseDataInner, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *GetPayouts200ResponseDataInner
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "PayoutApiService.CancelPayout")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PayoutApiService.CancelPayoutExecute")
	}

	localVarPath := localBasePath + "/v2/payouts/{id}/cancel"
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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PayoutApiService.CancelPayoutExecute")
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

type ApiCreatePayoutRequest struct {
	ctx context.Context
	ApiService PayoutApi
	idempotencyKey *string
	forUserId *string
	createPayoutRequest *CreatePayoutRequest
}

// A unique key to prevent duplicate requests from pushing through our system. No expiration.
func (r ApiCreatePayoutRequest) IdempotencyKey(idempotencyKey string) ApiCreatePayoutRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

// The sub-account user-id that you want to make this transaction for. This header is only used if you have access to xenPlatform. See xenPlatform for more information.
func (r ApiCreatePayoutRequest) ForUserId(forUserId string) ApiCreatePayoutRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiCreatePayoutRequest) CreatePayoutRequest(createPayoutRequest CreatePayoutRequest) ApiCreatePayoutRequest {
	r.createPayoutRequest = &createPayoutRequest
	return r
}

func (r ApiCreatePayoutRequest) Execute() (*GetPayouts200ResponseDataInner, *http.Response, *common.XenditSdkError) {
	return r.ApiService.CreatePayoutExecute(r)
}

/*
CreatePayout API to send money at scale to bank accounts & eWallets

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreatePayoutRequest
*/
func (a *PayoutApiService) CreatePayout(ctx context.Context) ApiCreatePayoutRequest {
	return ApiCreatePayoutRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetPayouts200ResponseDataInner
func (a *PayoutApiService) CreatePayoutExecute(r ApiCreatePayoutRequest) (*GetPayouts200ResponseDataInner, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *GetPayouts200ResponseDataInner
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "PayoutApiService.CreatePayout")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PayoutApiService.CreatePayoutExecute")
	}

	localVarPath := localBasePath + "/v2/payouts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.idempotencyKey == nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "idempotencyKey is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	utils.ParameterAddToHeaderOrQuery(localVarHeaderParams, "idempotency-key", r.idempotencyKey, "")
	if r.forUserId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarHeaderParams, "for-user-id", r.forUserId, "")
	}
	// body params
	localVarPostBody = r.createPayoutRequest
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PayoutApiService.CreatePayoutExecute")
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

type ApiGetPayoutByIdRequest struct {
	ctx context.Context
	ApiService PayoutApi
	id string
}

func (r ApiGetPayoutByIdRequest) Execute() (*GetPayouts200ResponseDataInner, *http.Response, *common.XenditSdkError) {
	return r.ApiService.GetPayoutByIdExecute(r)
}

/*
GetPayoutById API to fetch the current status, or details of the payout

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Payout id returned from the response of /v2/payouts
 @return ApiGetPayoutByIdRequest
*/
func (a *PayoutApiService) GetPayoutById(ctx context.Context, id string) ApiGetPayoutByIdRequest {
	return ApiGetPayoutByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GetPayouts200ResponseDataInner
func (a *PayoutApiService) GetPayoutByIdExecute(r ApiGetPayoutByIdRequest) (*GetPayouts200ResponseDataInner, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *GetPayouts200ResponseDataInner
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "PayoutApiService.GetPayoutById")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PayoutApiService.GetPayoutByIdExecute")
	}

	localVarPath := localBasePath + "/v2/payouts/{id}"
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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PayoutApiService.GetPayoutByIdExecute")
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

type ApiGetPayoutChannelsRequest struct {
	ctx context.Context
	ApiService PayoutApi
	currency *string
	channelCategory *[]ChannelCategory
	channelCode *string
}

// Filter channels by currency from ISO-4217 values
func (r ApiGetPayoutChannelsRequest) Currency(currency string) ApiGetPayoutChannelsRequest {
	r.currency = &currency
	return r
}

// Filter channels by category
func (r ApiGetPayoutChannelsRequest) ChannelCategory(channelCategory []ChannelCategory) ApiGetPayoutChannelsRequest {
	r.channelCategory = &channelCategory
	return r
}

// Filter channels by channel code, prefixed by ISO-3166 country code
func (r ApiGetPayoutChannelsRequest) ChannelCode(channelCode string) ApiGetPayoutChannelsRequest {
	r.channelCode = &channelCode
	return r
}

func (r ApiGetPayoutChannelsRequest) Execute() ([]Channel, *http.Response, *common.XenditSdkError) {
	return r.ApiService.GetPayoutChannelsExecute(r)
}

/*
GetPayoutChannels API providing the current list of banks and e-wallets we support for payouts for both regions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPayoutChannelsRequest
*/
func (a *PayoutApiService) GetPayoutChannels(ctx context.Context) ApiGetPayoutChannelsRequest {
	return ApiGetPayoutChannelsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Channel
func (a *PayoutApiService) GetPayoutChannelsExecute(r ApiGetPayoutChannelsRequest) ([]Channel, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  []Channel
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "PayoutApiService.GetPayoutChannels")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PayoutApiService.GetPayoutChannelsExecute")
	}

	localVarPath := localBasePath + "/payouts_channels"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.currency != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "currency", r.currency, "")
	}
	if r.channelCategory != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "channel_category", r.channelCategory, "csv")
	}
	if r.channelCode != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "channel_code", r.channelCode, "")
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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PayoutApiService.GetPayoutChannelsExecute")
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

type ApiGetPayoutsRequest struct {
	ctx context.Context
	ApiService PayoutApi
	referenceId *string
	limit *float32
	afterId *string
	beforeId *string
}

// Reference_id provided when creating the payout
func (r ApiGetPayoutsRequest) ReferenceId(referenceId string) ApiGetPayoutsRequest {
	r.referenceId = &referenceId
	return r
}

// Number of records to fetch per API call
func (r ApiGetPayoutsRequest) Limit(limit float32) ApiGetPayoutsRequest {
	r.limit = &limit
	return r
}

// Used to fetch record after this payout unique id
func (r ApiGetPayoutsRequest) AfterId(afterId string) ApiGetPayoutsRequest {
	r.afterId = &afterId
	return r
}

// Used to fetch record before this payout unique id
func (r ApiGetPayoutsRequest) BeforeId(beforeId string) ApiGetPayoutsRequest {
	r.beforeId = &beforeId
	return r
}

func (r ApiGetPayoutsRequest) Execute() (*GetPayouts200Response, *http.Response, *common.XenditSdkError) {
	return r.ApiService.GetPayoutsExecute(r)
}

/*
GetPayouts API to retrieve all matching payouts with reference ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPayoutsRequest
*/
func (a *PayoutApiService) GetPayouts(ctx context.Context) ApiGetPayoutsRequest {
	return ApiGetPayoutsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetPayouts200Response
func (a *PayoutApiService) GetPayoutsExecute(r ApiGetPayoutsRequest) (*GetPayouts200Response, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *GetPayouts200Response
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "PayoutApiService.GetPayouts")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PayoutApiService.GetPayoutsExecute")
	}

	localVarPath := localBasePath + "/v2/payouts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.referenceId == nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "referenceId is required and must be specified")
	}

	utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "reference_id", r.referenceId, "")
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
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PayoutApiService.GetPayoutsExecute")
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
