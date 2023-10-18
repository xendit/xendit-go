package refund

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


type RefundApi interface {

	/*
	CreateRefund Method for CreateRefund

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateRefundRequest
	*/
	CreateRefund(ctx context.Context) ApiCreateRefundRequest

	// CreateRefundExecute executes the request
	//  @return Refund
	CreateRefundExecute(r ApiCreateRefundRequest) (*Refund, *http.Response, *common.XenditSdkError)

	/*
	GetRefund Method for GetRefund

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param refundID
	@return ApiGetRefundRequest
	*/
	GetRefund(ctx context.Context, refundID string) ApiGetRefundRequest

	// GetRefundExecute executes the request
	//  @return Refund
	GetRefundExecute(r ApiGetRefundRequest) (*Refund, *http.Response, *common.XenditSdkError)

	/*
	GetAllRefunds Method for GetAllRefunds

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAllRefundsRequest
	*/
	GetAllRefunds(ctx context.Context) ApiGetAllRefundsRequest

	// GetAllRefundsExecute executes the request
	//  @return RefundList
	GetAllRefundsExecute(r ApiGetAllRefundsRequest) (*RefundList, *http.Response, *common.XenditSdkError)

	/*
	CancelRefund Method for CancelRefund

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param refundID
	@return ApiCancelRefundRequest
	*/
	CancelRefund(ctx context.Context, refundID string) ApiCancelRefundRequest

	// CancelRefundExecute executes the request
	//  @return Refund
	CancelRefundExecute(r ApiCancelRefundRequest) (*Refund, *http.Response, *common.XenditSdkError)
}

// RefundApiService RefundApi service
type RefundApiService struct {
	client common.IClient
}

// NewRefundApi Create a new RefundApi service
func NewRefundApi (client common.IClient) RefundApi {
	return &RefundApiService{
		client: client,
	}
}


type ApiCreateRefundRequest struct {
	ctx context.Context
	ApiService RefundApi
	idempotencyKey *string
	forUserId *string
	createRefund *CreateRefund
}

func (r ApiCreateRefundRequest) IdempotencyKey(idempotencyKey string) ApiCreateRefundRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiCreateRefundRequest) ForUserId(forUserId string) ApiCreateRefundRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiCreateRefundRequest) CreateRefund(createRefund CreateRefund) ApiCreateRefundRequest {
	r.createRefund = &createRefund
	return r
}

func (r ApiCreateRefundRequest) Execute() (*Refund, *http.Response, *common.XenditSdkError) {
	return r.ApiService.CreateRefundExecute(r)
}

/*
CreateRefund Method for CreateRefund

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateRefundRequest
*/
func (a *RefundApiService) CreateRefund(ctx context.Context) ApiCreateRefundRequest {
	return ApiCreateRefundRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Refund
func (a *RefundApiService) CreateRefundExecute(r ApiCreateRefundRequest) (*Refund, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *Refund
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "RefundApiService.CreateRefund")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: RefundApiService.CreateRefundExecute")
	}

	localVarPath := localBasePath + "/refunds"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.idempotencyKey != nil {
		utils.ParameterAddToHeaderOrQuery(localVarHeaderParams, "idempotency-key", r.idempotencyKey, "")
	}
	if r.forUserId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarHeaderParams, "for-user-id", r.forUserId, "")
	}
	// body params
	localVarPostBody = r.createRefund
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: RefundApiService.CreateRefundExecute")
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

type ApiGetRefundRequest struct {
	ctx context.Context
	ApiService RefundApi
	refundID string
	idempotencyKey *string
	forUserId *string
}

func (r ApiGetRefundRequest) IdempotencyKey(idempotencyKey string) ApiGetRefundRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiGetRefundRequest) ForUserId(forUserId string) ApiGetRefundRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiGetRefundRequest) Execute() (*Refund, *http.Response, *common.XenditSdkError) {
	return r.ApiService.GetRefundExecute(r)
}

/*
GetRefund Method for GetRefund

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param refundID
 @return ApiGetRefundRequest
*/
func (a *RefundApiService) GetRefund(ctx context.Context, refundID string) ApiGetRefundRequest {
	return ApiGetRefundRequest{
		ApiService: a,
		ctx: ctx,
		refundID: refundID,
	}
}

// Execute executes the request
//  @return Refund
func (a *RefundApiService) GetRefundExecute(r ApiGetRefundRequest) (*Refund, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *Refund
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "RefundApiService.GetRefund")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: RefundApiService.GetRefundExecute")
	}

	localVarPath := localBasePath + "/refunds/{refundID}"
	localVarPath = strings.Replace(localVarPath, "{"+"refundID"+"}", url.PathEscape(utils.ParameterValueToString(r.refundID, "refundID")), -1)

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
	if r.idempotencyKey != nil {
		utils.ParameterAddToHeaderOrQuery(localVarHeaderParams, "idempotency-key", r.idempotencyKey, "")
	}
	if r.forUserId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarHeaderParams, "for-user-id", r.forUserId, "")
	}
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: RefundApiService.GetRefundExecute")
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

type ApiGetAllRefundsRequest struct {
	ctx context.Context
	ApiService RefundApi
	forUserId *string
	paymentRequestId *string
	invoiceId *string
	paymentMethodType *string
	channelCode *string
	limit *float32
	afterId *string
	beforeId *string
}

func (r ApiGetAllRefundsRequest) ForUserId(forUserId string) ApiGetAllRefundsRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiGetAllRefundsRequest) PaymentRequestId(paymentRequestId string) ApiGetAllRefundsRequest {
	r.paymentRequestId = &paymentRequestId
	return r
}

func (r ApiGetAllRefundsRequest) InvoiceId(invoiceId string) ApiGetAllRefundsRequest {
	r.invoiceId = &invoiceId
	return r
}

func (r ApiGetAllRefundsRequest) PaymentMethodType(paymentMethodType string) ApiGetAllRefundsRequest {
	r.paymentMethodType = &paymentMethodType
	return r
}

func (r ApiGetAllRefundsRequest) ChannelCode(channelCode string) ApiGetAllRefundsRequest {
	r.channelCode = &channelCode
	return r
}

func (r ApiGetAllRefundsRequest) Limit(limit float32) ApiGetAllRefundsRequest {
	r.limit = &limit
	return r
}

func (r ApiGetAllRefundsRequest) AfterId(afterId string) ApiGetAllRefundsRequest {
	r.afterId = &afterId
	return r
}

func (r ApiGetAllRefundsRequest) BeforeId(beforeId string) ApiGetAllRefundsRequest {
	r.beforeId = &beforeId
	return r
}

func (r ApiGetAllRefundsRequest) Execute() (*RefundList, *http.Response, *common.XenditSdkError) {
	return r.ApiService.GetAllRefundsExecute(r)
}

/*
GetAllRefunds Method for GetAllRefunds

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllRefundsRequest
*/
func (a *RefundApiService) GetAllRefunds(ctx context.Context) ApiGetAllRefundsRequest {
	return ApiGetAllRefundsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RefundList
func (a *RefundApiService) GetAllRefundsExecute(r ApiGetAllRefundsRequest) (*RefundList, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *RefundList
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "RefundApiService.GetAllRefunds")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: RefundApiService.GetAllRefundsExecute")
	}

	localVarPath := localBasePath + "/refunds"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.paymentRequestId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "payment_request_id", r.paymentRequestId, "")
	}
	if r.invoiceId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "invoice_id", r.invoiceId, "")
	}
	if r.paymentMethodType != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "payment_method_type", r.paymentMethodType, "")
	}
	if r.channelCode != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "channel_code", r.channelCode, "")
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
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: RefundApiService.GetAllRefundsExecute")
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

type ApiCancelRefundRequest struct {
	ctx context.Context
	ApiService RefundApi
	refundID string
	idempotencyKey *string
	forUserId *string
}

func (r ApiCancelRefundRequest) IdempotencyKey(idempotencyKey string) ApiCancelRefundRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiCancelRefundRequest) ForUserId(forUserId string) ApiCancelRefundRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiCancelRefundRequest) Execute() (*Refund, *http.Response, *common.XenditSdkError) {
	return r.ApiService.CancelRefundExecute(r)
}

/*
CancelRefund Method for CancelRefund

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param refundID
 @return ApiCancelRefundRequest
*/
func (a *RefundApiService) CancelRefund(ctx context.Context, refundID string) ApiCancelRefundRequest {
	return ApiCancelRefundRequest{
		ApiService: a,
		ctx: ctx,
		refundID: refundID,
	}
}

// Execute executes the request
//  @return Refund
func (a *RefundApiService) CancelRefundExecute(r ApiCancelRefundRequest) (*Refund, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *Refund
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "RefundApiService.CancelRefund")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: RefundApiService.CancelRefundExecute")
	}

	localVarPath := localBasePath + "/refunds/{refundID}/cancel"
	localVarPath = strings.Replace(localVarPath, "{"+"refundID"+"}", url.PathEscape(utils.ParameterValueToString(r.refundID, "refundID")), -1)

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
	if r.idempotencyKey != nil {
		utils.ParameterAddToHeaderOrQuery(localVarHeaderParams, "idempotency-key", r.idempotencyKey, "")
	}
	if r.forUserId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarHeaderParams, "for-user-id", r.forUserId, "")
	}
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: RefundApiService.CancelRefundExecute")
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
