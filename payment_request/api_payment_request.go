package payment_request

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
	"reflect"
)


type PaymentRequestApi interface {

	/*
	CreatePaymentRequest Create Payment Request

	Create Payment Request

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreatePaymentRequestRequest
	*/
	CreatePaymentRequest(ctx context.Context) ApiCreatePaymentRequestRequest

	// CreatePaymentRequestExecute executes the request
	//  @return PaymentRequest
	CreatePaymentRequestExecute(r ApiCreatePaymentRequestRequest) (*PaymentRequest, *http.Response, *common.XenditSdkError)

	/*
	GetPaymentRequestByID Get payment request by ID

	Get payment request by ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param paymentRequestId
	@return ApiGetPaymentRequestByIDRequest
	*/
	GetPaymentRequestByID(ctx context.Context, paymentRequestId string) ApiGetPaymentRequestByIDRequest

	// GetPaymentRequestByIDExecute executes the request
	//  @return PaymentRequest
	GetPaymentRequestByIDExecute(r ApiGetPaymentRequestByIDRequest) (*PaymentRequest, *http.Response, *common.XenditSdkError)

	/*
	GetPaymentRequestCaptures Get Payment Request Capture

	Get Payment Request Capture

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param paymentRequestId
	@return ApiGetPaymentRequestCapturesRequest
	*/
	GetPaymentRequestCaptures(ctx context.Context, paymentRequestId string) ApiGetPaymentRequestCapturesRequest

	// GetPaymentRequestCapturesExecute executes the request
	//  @return CaptureListResponse
	GetPaymentRequestCapturesExecute(r ApiGetPaymentRequestCapturesRequest) (*CaptureListResponse, *http.Response, *common.XenditSdkError)

	/*
	GetAllPaymentRequests Get all payment requests by filter

	Get all payment requests by filter

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAllPaymentRequestsRequest
	*/
	GetAllPaymentRequests(ctx context.Context) ApiGetAllPaymentRequestsRequest

	// GetAllPaymentRequestsExecute executes the request
	//  @return PaymentRequestListResponse
	GetAllPaymentRequestsExecute(r ApiGetAllPaymentRequestsRequest) (*PaymentRequestListResponse, *http.Response, *common.XenditSdkError)

	/*
	CapturePaymentRequest Payment Request Capture

	Payment Request Capture

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param paymentRequestId
	@return ApiCapturePaymentRequestRequest
	*/
	CapturePaymentRequest(ctx context.Context, paymentRequestId string) ApiCapturePaymentRequestRequest

	// CapturePaymentRequestExecute executes the request
	//  @return Capture
	CapturePaymentRequestExecute(r ApiCapturePaymentRequestRequest) (*Capture, *http.Response, *common.XenditSdkError)

	/*
	AuthorizePaymentRequest Payment Request Authorize

	Payment Request Authorize

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param paymentRequestId
	@return ApiAuthorizePaymentRequestRequest
	*/
	AuthorizePaymentRequest(ctx context.Context, paymentRequestId string) ApiAuthorizePaymentRequestRequest

	// AuthorizePaymentRequestExecute executes the request
	//  @return PaymentRequest
	AuthorizePaymentRequestExecute(r ApiAuthorizePaymentRequestRequest) (*PaymentRequest, *http.Response, *common.XenditSdkError)

	/*
	ResendPaymentRequestAuth Payment Request Resend Auth

	Payment Request Resend Auth

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param paymentRequestId
	@return ApiResendPaymentRequestAuthRequest
	*/
	ResendPaymentRequestAuth(ctx context.Context, paymentRequestId string) ApiResendPaymentRequestAuthRequest

	// ResendPaymentRequestAuthExecute executes the request
	//  @return PaymentRequest
	ResendPaymentRequestAuthExecute(r ApiResendPaymentRequestAuthRequest) (*PaymentRequest, *http.Response, *common.XenditSdkError)
}

// PaymentRequestApiService PaymentRequestApi service
type PaymentRequestApiService struct {
	client common.IClient
}

// NewPaymentRequestApi Create a new PaymentRequestApi service
func NewPaymentRequestApi (client common.IClient) PaymentRequestApi {
	return &PaymentRequestApiService{
		client: client,
	}
}


type ApiCreatePaymentRequestRequest struct {
	ctx context.Context
	ApiService PaymentRequestApi
	idempotencyKey *string
	forUserId *string
	paymentRequestParameters *PaymentRequestParameters
}

func (r ApiCreatePaymentRequestRequest) IdempotencyKey(idempotencyKey string) ApiCreatePaymentRequestRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiCreatePaymentRequestRequest) ForUserId(forUserId string) ApiCreatePaymentRequestRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiCreatePaymentRequestRequest) PaymentRequestParameters(paymentRequestParameters PaymentRequestParameters) ApiCreatePaymentRequestRequest {
	r.paymentRequestParameters = &paymentRequestParameters
	return r
}

func (r ApiCreatePaymentRequestRequest) Execute() (*PaymentRequest, *http.Response, *common.XenditSdkError) {
	return r.ApiService.CreatePaymentRequestExecute(r)
}

/*
CreatePaymentRequest Create Payment Request

Create Payment Request

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreatePaymentRequestRequest
*/
func (a *PaymentRequestApiService) CreatePaymentRequest(ctx context.Context) ApiCreatePaymentRequestRequest {
	return ApiCreatePaymentRequestRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaymentRequest
func (a *PaymentRequestApiService) CreatePaymentRequestExecute(r ApiCreatePaymentRequestRequest) (*PaymentRequest, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *PaymentRequest
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "PaymentRequestApiService.CreatePaymentRequest")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentRequestApiService.CreatePaymentRequestExecute")
	}

	localVarPath := localBasePath + "/payment_requests"

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
	localVarPostBody = r.paymentRequestParameters
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentRequestApiService.CreatePaymentRequestExecute")
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

type ApiGetPaymentRequestByIDRequest struct {
	ctx context.Context
	ApiService PaymentRequestApi
	paymentRequestId string
	forUserId *string
}

func (r ApiGetPaymentRequestByIDRequest) ForUserId(forUserId string) ApiGetPaymentRequestByIDRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiGetPaymentRequestByIDRequest) Execute() (*PaymentRequest, *http.Response, *common.XenditSdkError) {
	return r.ApiService.GetPaymentRequestByIDExecute(r)
}

/*
GetPaymentRequestByID Get payment request by ID

Get payment request by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentRequestId
 @return ApiGetPaymentRequestByIDRequest
*/
func (a *PaymentRequestApiService) GetPaymentRequestByID(ctx context.Context, paymentRequestId string) ApiGetPaymentRequestByIDRequest {
	return ApiGetPaymentRequestByIDRequest{
		ApiService: a,
		ctx: ctx,
		paymentRequestId: paymentRequestId,
	}
}

// Execute executes the request
//  @return PaymentRequest
func (a *PaymentRequestApiService) GetPaymentRequestByIDExecute(r ApiGetPaymentRequestByIDRequest) (*PaymentRequest, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *PaymentRequest
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "PaymentRequestApiService.GetPaymentRequestByID")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentRequestApiService.GetPaymentRequestByIDExecute")
	}

	localVarPath := localBasePath + "/payment_requests/{paymentRequestId}"
	localVarPath = strings.Replace(localVarPath, "{"+"paymentRequestId"+"}", url.PathEscape(utils.ParameterValueToString(r.paymentRequestId, "paymentRequestId")), -1)

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
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentRequestApiService.GetPaymentRequestByIDExecute")
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

type ApiGetPaymentRequestCapturesRequest struct {
	ctx context.Context
	ApiService PaymentRequestApi
	paymentRequestId string
	forUserId *string
	limit *int32
}

func (r ApiGetPaymentRequestCapturesRequest) ForUserId(forUserId string) ApiGetPaymentRequestCapturesRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiGetPaymentRequestCapturesRequest) Limit(limit int32) ApiGetPaymentRequestCapturesRequest {
	r.limit = &limit
	return r
}

func (r ApiGetPaymentRequestCapturesRequest) Execute() (*CaptureListResponse, *http.Response, *common.XenditSdkError) {
	return r.ApiService.GetPaymentRequestCapturesExecute(r)
}

/*
GetPaymentRequestCaptures Get Payment Request Capture

Get Payment Request Capture

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentRequestId
 @return ApiGetPaymentRequestCapturesRequest
*/
func (a *PaymentRequestApiService) GetPaymentRequestCaptures(ctx context.Context, paymentRequestId string) ApiGetPaymentRequestCapturesRequest {
	return ApiGetPaymentRequestCapturesRequest{
		ApiService: a,
		ctx: ctx,
		paymentRequestId: paymentRequestId,
	}
}

// Execute executes the request
//  @return CaptureListResponse
func (a *PaymentRequestApiService) GetPaymentRequestCapturesExecute(r ApiGetPaymentRequestCapturesRequest) (*CaptureListResponse, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *CaptureListResponse
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "PaymentRequestApiService.GetPaymentRequestCaptures")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentRequestApiService.GetPaymentRequestCapturesExecute")
	}

	localVarPath := localBasePath + "/payment_requests/{paymentRequestId}/captures"
	localVarPath = strings.Replace(localVarPath, "{"+"paymentRequestId"+"}", url.PathEscape(utils.ParameterValueToString(r.paymentRequestId, "paymentRequestId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentRequestApiService.GetPaymentRequestCapturesExecute")
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

type ApiGetAllPaymentRequestsRequest struct {
	ctx context.Context
	ApiService PaymentRequestApi
	forUserId *string
	referenceId *[]string
	id *[]string
	customerId *[]string
	limit *int32
	beforeId *string
	afterId *string
}

func (r ApiGetAllPaymentRequestsRequest) ForUserId(forUserId string) ApiGetAllPaymentRequestsRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiGetAllPaymentRequestsRequest) ReferenceId(referenceId []string) ApiGetAllPaymentRequestsRequest {
	r.referenceId = &referenceId
	return r
}

func (r ApiGetAllPaymentRequestsRequest) Id(id []string) ApiGetAllPaymentRequestsRequest {
	r.id = &id
	return r
}

func (r ApiGetAllPaymentRequestsRequest) CustomerId(customerId []string) ApiGetAllPaymentRequestsRequest {
	r.customerId = &customerId
	return r
}

func (r ApiGetAllPaymentRequestsRequest) Limit(limit int32) ApiGetAllPaymentRequestsRequest {
	r.limit = &limit
	return r
}

func (r ApiGetAllPaymentRequestsRequest) BeforeId(beforeId string) ApiGetAllPaymentRequestsRequest {
	r.beforeId = &beforeId
	return r
}

func (r ApiGetAllPaymentRequestsRequest) AfterId(afterId string) ApiGetAllPaymentRequestsRequest {
	r.afterId = &afterId
	return r
}

func (r ApiGetAllPaymentRequestsRequest) Execute() (*PaymentRequestListResponse, *http.Response, *common.XenditSdkError) {
	return r.ApiService.GetAllPaymentRequestsExecute(r)
}

/*
GetAllPaymentRequests Get all payment requests by filter

Get all payment requests by filter

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllPaymentRequestsRequest
*/
func (a *PaymentRequestApiService) GetAllPaymentRequests(ctx context.Context) ApiGetAllPaymentRequestsRequest {
	return ApiGetAllPaymentRequestsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaymentRequestListResponse
func (a *PaymentRequestApiService) GetAllPaymentRequestsExecute(r ApiGetAllPaymentRequestsRequest) (*PaymentRequestListResponse, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *PaymentRequestListResponse
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "PaymentRequestApiService.GetAllPaymentRequests")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentRequestApiService.GetAllPaymentRequestsExecute")
	}

	localVarPath := localBasePath + "/payment_requests"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.referenceId != nil {
		t := *r.referenceId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "reference_id", s.Index(i), "multi")
			}
		} else {
			utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "reference_id", t, "multi")
		}
	}
	if r.id != nil {
		t := *r.id
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "id", s.Index(i), "multi")
			}
		} else {
			utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "id", t, "multi")
		}
	}
	if r.customerId != nil {
		t := *r.customerId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "customer_id", s.Index(i), "multi")
			}
		} else {
			utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "customer_id", t, "multi")
		}
	}
	if r.limit != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.beforeId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "before_id", r.beforeId, "")
	}
	if r.afterId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "after_id", r.afterId, "")
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
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentRequestApiService.GetAllPaymentRequestsExecute")
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

type ApiCapturePaymentRequestRequest struct {
	ctx context.Context
	ApiService PaymentRequestApi
	paymentRequestId string
	forUserId *string
	captureParameters *CaptureParameters
}

func (r ApiCapturePaymentRequestRequest) ForUserId(forUserId string) ApiCapturePaymentRequestRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiCapturePaymentRequestRequest) CaptureParameters(captureParameters CaptureParameters) ApiCapturePaymentRequestRequest {
	r.captureParameters = &captureParameters
	return r
}

func (r ApiCapturePaymentRequestRequest) Execute() (*Capture, *http.Response, *common.XenditSdkError) {
	return r.ApiService.CapturePaymentRequestExecute(r)
}

/*
CapturePaymentRequest Payment Request Capture

Payment Request Capture

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentRequestId
 @return ApiCapturePaymentRequestRequest
*/
func (a *PaymentRequestApiService) CapturePaymentRequest(ctx context.Context, paymentRequestId string) ApiCapturePaymentRequestRequest {
	return ApiCapturePaymentRequestRequest{
		ApiService: a,
		ctx: ctx,
		paymentRequestId: paymentRequestId,
	}
}

// Execute executes the request
//  @return Capture
func (a *PaymentRequestApiService) CapturePaymentRequestExecute(r ApiCapturePaymentRequestRequest) (*Capture, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *Capture
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "PaymentRequestApiService.CapturePaymentRequest")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentRequestApiService.CapturePaymentRequestExecute")
	}

	localVarPath := localBasePath + "/payment_requests/{paymentRequestId}/captures"
	localVarPath = strings.Replace(localVarPath, "{"+"paymentRequestId"+"}", url.PathEscape(utils.ParameterValueToString(r.paymentRequestId, "paymentRequestId")), -1)

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
	if r.forUserId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarHeaderParams, "for-user-id", r.forUserId, "")
	}
	// body params
	localVarPostBody = r.captureParameters
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentRequestApiService.CapturePaymentRequestExecute")
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

type ApiAuthorizePaymentRequestRequest struct {
	ctx context.Context
	ApiService PaymentRequestApi
	paymentRequestId string
	forUserId *string
	paymentRequestAuthParameters *PaymentRequestAuthParameters
}

func (r ApiAuthorizePaymentRequestRequest) ForUserId(forUserId string) ApiAuthorizePaymentRequestRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiAuthorizePaymentRequestRequest) PaymentRequestAuthParameters(paymentRequestAuthParameters PaymentRequestAuthParameters) ApiAuthorizePaymentRequestRequest {
	r.paymentRequestAuthParameters = &paymentRequestAuthParameters
	return r
}

func (r ApiAuthorizePaymentRequestRequest) Execute() (*PaymentRequest, *http.Response, *common.XenditSdkError) {
	return r.ApiService.AuthorizePaymentRequestExecute(r)
}

/*
AuthorizePaymentRequest Payment Request Authorize

Payment Request Authorize

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentRequestId
 @return ApiAuthorizePaymentRequestRequest
*/
func (a *PaymentRequestApiService) AuthorizePaymentRequest(ctx context.Context, paymentRequestId string) ApiAuthorizePaymentRequestRequest {
	return ApiAuthorizePaymentRequestRequest{
		ApiService: a,
		ctx: ctx,
		paymentRequestId: paymentRequestId,
	}
}

// Execute executes the request
//  @return PaymentRequest
func (a *PaymentRequestApiService) AuthorizePaymentRequestExecute(r ApiAuthorizePaymentRequestRequest) (*PaymentRequest, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *PaymentRequest
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "PaymentRequestApiService.AuthorizePaymentRequest")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentRequestApiService.AuthorizePaymentRequestExecute")
	}

	localVarPath := localBasePath + "/payment_requests/{paymentRequestId}/auth"
	localVarPath = strings.Replace(localVarPath, "{"+"paymentRequestId"+"}", url.PathEscape(utils.ParameterValueToString(r.paymentRequestId, "paymentRequestId")), -1)

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
	if r.forUserId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarHeaderParams, "for-user-id", r.forUserId, "")
	}
	// body params
	localVarPostBody = r.paymentRequestAuthParameters
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentRequestApiService.AuthorizePaymentRequestExecute")
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

type ApiResendPaymentRequestAuthRequest struct {
	ctx context.Context
	ApiService PaymentRequestApi
	paymentRequestId string
	forUserId *string
}

func (r ApiResendPaymentRequestAuthRequest) ForUserId(forUserId string) ApiResendPaymentRequestAuthRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiResendPaymentRequestAuthRequest) Execute() (*PaymentRequest, *http.Response, *common.XenditSdkError) {
	return r.ApiService.ResendPaymentRequestAuthExecute(r)
}

/*
ResendPaymentRequestAuth Payment Request Resend Auth

Payment Request Resend Auth

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentRequestId
 @return ApiResendPaymentRequestAuthRequest
*/
func (a *PaymentRequestApiService) ResendPaymentRequestAuth(ctx context.Context, paymentRequestId string) ApiResendPaymentRequestAuthRequest {
	return ApiResendPaymentRequestAuthRequest{
		ApiService: a,
		ctx: ctx,
		paymentRequestId: paymentRequestId,
	}
}

// Execute executes the request
//  @return PaymentRequest
func (a *PaymentRequestApiService) ResendPaymentRequestAuthExecute(r ApiResendPaymentRequestAuthRequest) (*PaymentRequest, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *PaymentRequest
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "PaymentRequestApiService.ResendPaymentRequestAuth")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentRequestApiService.ResendPaymentRequestAuthExecute")
	}

	localVarPath := localBasePath + "/payment_requests/{paymentRequestId}/auth/resend"
	localVarPath = strings.Replace(localVarPath, "{"+"paymentRequestId"+"}", url.PathEscape(utils.ParameterValueToString(r.paymentRequestId, "paymentRequestId")), -1)

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
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentRequestApiService.ResendPaymentRequestAuthExecute")
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
