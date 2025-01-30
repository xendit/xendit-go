package payment_method

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
	"strings"
	"reflect"
	"time"
)


type PaymentMethodApi interface {

	/*
	CreatePaymentMethod Creates payment method

	This endpoint initiates creation of payment method

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreatePaymentMethodRequest
	*/
	CreatePaymentMethod(ctx context.Context) ApiCreatePaymentMethodRequest

	// CreatePaymentMethodExecute executes the request
	//  @return PaymentMethod
	CreatePaymentMethodExecute(r ApiCreatePaymentMethodRequest) (*PaymentMethod, *http.Response, *common.XenditSdkError)

	/*
	GetPaymentMethodByID Get payment method by ID

	Get payment method by ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param paymentMethodId
	@return ApiGetPaymentMethodByIDRequest
	*/
	GetPaymentMethodByID(ctx context.Context, paymentMethodId string) ApiGetPaymentMethodByIDRequest

	// GetPaymentMethodByIDExecute executes the request
	//  @return PaymentMethod
	GetPaymentMethodByIDExecute(r ApiGetPaymentMethodByIDRequest) (*PaymentMethod, *http.Response, *common.XenditSdkError)

	/*
	GetPaymentsByPaymentMethodId Returns payments with matching PaymentMethodID.

	Returns payments with matching PaymentMethodID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param paymentMethodId
	@return ApiGetPaymentsByPaymentMethodIdRequest
	*/
	GetPaymentsByPaymentMethodId(ctx context.Context, paymentMethodId string) ApiGetPaymentsByPaymentMethodIdRequest

	// GetPaymentsByPaymentMethodIdExecute executes the request
	//  @return map[string]interface{}
	GetPaymentsByPaymentMethodIdExecute(r ApiGetPaymentsByPaymentMethodIdRequest) (map[string]interface{}, *http.Response, *common.XenditSdkError)

	/*
	PatchPaymentMethod Patch payment methods

	This endpoint is used to toggle the ```status``` of an e-Wallet or a Direct Debit payment method to ```ACTIVE``` or ```INACTIVE```. This is also used to update the details of an Over-the-Counter or a Virtual Account payment method.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param paymentMethodId
	@return ApiPatchPaymentMethodRequest
	*/
	PatchPaymentMethod(ctx context.Context, paymentMethodId string) ApiPatchPaymentMethodRequest

	// PatchPaymentMethodExecute executes the request
	//  @return PaymentMethod
	PatchPaymentMethodExecute(r ApiPatchPaymentMethodRequest) (*PaymentMethod, *http.Response, *common.XenditSdkError)

	/*
	GetAllPaymentMethods Get all payment methods by filters

	Get all payment methods by filters

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAllPaymentMethodsRequest
	*/
	GetAllPaymentMethods(ctx context.Context) ApiGetAllPaymentMethodsRequest

	// GetAllPaymentMethodsExecute executes the request
	//  @return PaymentMethodList
	GetAllPaymentMethodsExecute(r ApiGetAllPaymentMethodsRequest) (*PaymentMethodList, *http.Response, *common.XenditSdkError)

	/*
	ExpirePaymentMethod Expires a payment method

	This endpoint expires a payment method and performs unlinking if necessary

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param paymentMethodId
	@return ApiExpirePaymentMethodRequest
	*/
	ExpirePaymentMethod(ctx context.Context, paymentMethodId string) ApiExpirePaymentMethodRequest

	// ExpirePaymentMethodExecute executes the request
	//  @return PaymentMethod
	ExpirePaymentMethodExecute(r ApiExpirePaymentMethodRequest) (*PaymentMethod, *http.Response, *common.XenditSdkError)

	/*
	AuthPaymentMethod Validate a payment method's linking OTP

	This endpoint validates a payment method linking OTP

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param paymentMethodId
	@return ApiAuthPaymentMethodRequest
	*/
	AuthPaymentMethod(ctx context.Context, paymentMethodId string) ApiAuthPaymentMethodRequest

	// AuthPaymentMethodExecute executes the request
	//  @return PaymentMethod
	AuthPaymentMethodExecute(r ApiAuthPaymentMethodRequest) (*PaymentMethod, *http.Response, *common.XenditSdkError)

	/*
	SimulatePayment Makes payment with matching PaymentMethodID.

	Makes payment with matching PaymentMethodID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param paymentMethodId
	@return ApiSimulatePaymentRequest
	*/
	SimulatePayment(ctx context.Context, paymentMethodId string) ApiSimulatePaymentRequest

	// SimulatePaymentExecute executes the request
	SimulatePaymentExecute(r ApiSimulatePaymentRequest) (*http.Response, *common.XenditSdkError)
}

// PaymentMethodApiService PaymentMethodApi service
type PaymentMethodApiService struct {
	client common.IClient
}

// NewPaymentMethodApi Create a new PaymentMethodApi service
func NewPaymentMethodApi (client common.IClient) PaymentMethodApi {
	return &PaymentMethodApiService{
		client: client,
	}
}


type ApiCreatePaymentMethodRequest struct {
	ctx context.Context
	ApiService PaymentMethodApi
	forUserId *string
	paymentMethodParameters *PaymentMethodParameters
}

func (r ApiCreatePaymentMethodRequest) ForUserId(forUserId string) ApiCreatePaymentMethodRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiCreatePaymentMethodRequest) PaymentMethodParameters(paymentMethodParameters PaymentMethodParameters) ApiCreatePaymentMethodRequest {
	r.paymentMethodParameters = &paymentMethodParameters
	return r
}

func (r ApiCreatePaymentMethodRequest) Execute() (*PaymentMethod, *http.Response, *common.XenditSdkError) {
	return r.ApiService.CreatePaymentMethodExecute(r)
}

/*
CreatePaymentMethod Creates payment method

This endpoint initiates creation of payment method

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreatePaymentMethodRequest
*/
func (a *PaymentMethodApiService) CreatePaymentMethod(ctx context.Context) ApiCreatePaymentMethodRequest {
	return ApiCreatePaymentMethodRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaymentMethod
func (a *PaymentMethodApiService) CreatePaymentMethodExecute(r ApiCreatePaymentMethodRequest) (*PaymentMethod, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *PaymentMethod
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "PaymentMethodApiService.CreatePaymentMethod")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentMethodApiService.CreatePaymentMethodExecute")
	}

	localVarPath := localBasePath + "/v2/payment_methods"

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
	localVarPostBody = r.paymentMethodParameters
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentMethodApiService.CreatePaymentMethodExecute")
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
    if err != nil {
        return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", fmt.Sprintf("Error creating HTTP request: PaymentMethodApiService.CreatePaymentMethodExecute: %v", err))
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

type ApiGetPaymentMethodByIDRequest struct {
	ctx context.Context
	ApiService PaymentMethodApi
	paymentMethodId string
	forUserId *string
}

func (r ApiGetPaymentMethodByIDRequest) ForUserId(forUserId string) ApiGetPaymentMethodByIDRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiGetPaymentMethodByIDRequest) Execute() (*PaymentMethod, *http.Response, *common.XenditSdkError) {
	return r.ApiService.GetPaymentMethodByIDExecute(r)
}

/*
GetPaymentMethodByID Get payment method by ID

Get payment method by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentMethodId
 @return ApiGetPaymentMethodByIDRequest
*/
func (a *PaymentMethodApiService) GetPaymentMethodByID(ctx context.Context, paymentMethodId string) ApiGetPaymentMethodByIDRequest {
	return ApiGetPaymentMethodByIDRequest{
		ApiService: a,
		ctx: ctx,
		paymentMethodId: paymentMethodId,
	}
}

// Execute executes the request
//  @return PaymentMethod
func (a *PaymentMethodApiService) GetPaymentMethodByIDExecute(r ApiGetPaymentMethodByIDRequest) (*PaymentMethod, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *PaymentMethod
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "PaymentMethodApiService.GetPaymentMethodByID")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentMethodApiService.GetPaymentMethodByIDExecute")
	}

	localVarPath := localBasePath + "/v2/payment_methods/{paymentMethodId}"
	localVarPath = strings.Replace(localVarPath, "{"+"paymentMethodId"+"}", url.PathEscape(utils.ParameterValueToString(r.paymentMethodId, "paymentMethodId")), -1)

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
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentMethodApiService.GetPaymentMethodByIDExecute")
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
    if err != nil {
        return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", fmt.Sprintf("Error creating HTTP request: PaymentMethodApiService.GetPaymentMethodByIDExecute: %v", err))
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

type ApiGetPaymentsByPaymentMethodIdRequest struct {
	ctx context.Context
	ApiService PaymentMethodApi
	paymentMethodId string
	forUserId *string
	paymentRequestId *[]string
	paymentMethodId2 *[]string
	referenceId *[]string
	paymentMethodType *[]PaymentMethodType
	channelCode *[]string
	status *[]string
	currency *[]string
	createdGte *time.Time
	createdLte *time.Time
	updatedGte *time.Time
	updatedLte *time.Time
	limit *int32
}

func (r ApiGetPaymentsByPaymentMethodIdRequest) ForUserId(forUserId string) ApiGetPaymentsByPaymentMethodIdRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiGetPaymentsByPaymentMethodIdRequest) PaymentRequestId(paymentRequestId []string) ApiGetPaymentsByPaymentMethodIdRequest {
	r.paymentRequestId = &paymentRequestId
	return r
}

func (r ApiGetPaymentsByPaymentMethodIdRequest) PaymentMethodId2(paymentMethodId2 []string) ApiGetPaymentsByPaymentMethodIdRequest {
	r.paymentMethodId2 = &paymentMethodId2
	return r
}

func (r ApiGetPaymentsByPaymentMethodIdRequest) ReferenceId(referenceId []string) ApiGetPaymentsByPaymentMethodIdRequest {
	r.referenceId = &referenceId
	return r
}

func (r ApiGetPaymentsByPaymentMethodIdRequest) PaymentMethodType(paymentMethodType []PaymentMethodType) ApiGetPaymentsByPaymentMethodIdRequest {
	r.paymentMethodType = &paymentMethodType
	return r
}

func (r ApiGetPaymentsByPaymentMethodIdRequest) ChannelCode(channelCode []string) ApiGetPaymentsByPaymentMethodIdRequest {
	r.channelCode = &channelCode
	return r
}

func (r ApiGetPaymentsByPaymentMethodIdRequest) Status(status []string) ApiGetPaymentsByPaymentMethodIdRequest {
	r.status = &status
	return r
}

func (r ApiGetPaymentsByPaymentMethodIdRequest) Currency(currency []string) ApiGetPaymentsByPaymentMethodIdRequest {
	r.currency = &currency
	return r
}

func (r ApiGetPaymentsByPaymentMethodIdRequest) CreatedGte(createdGte time.Time) ApiGetPaymentsByPaymentMethodIdRequest {
	r.createdGte = &createdGte
	return r
}

func (r ApiGetPaymentsByPaymentMethodIdRequest) CreatedLte(createdLte time.Time) ApiGetPaymentsByPaymentMethodIdRequest {
	r.createdLte = &createdLte
	return r
}

func (r ApiGetPaymentsByPaymentMethodIdRequest) UpdatedGte(updatedGte time.Time) ApiGetPaymentsByPaymentMethodIdRequest {
	r.updatedGte = &updatedGte
	return r
}

func (r ApiGetPaymentsByPaymentMethodIdRequest) UpdatedLte(updatedLte time.Time) ApiGetPaymentsByPaymentMethodIdRequest {
	r.updatedLte = &updatedLte
	return r
}

func (r ApiGetPaymentsByPaymentMethodIdRequest) Limit(limit int32) ApiGetPaymentsByPaymentMethodIdRequest {
	r.limit = &limit
	return r
}

func (r ApiGetPaymentsByPaymentMethodIdRequest) Execute() (map[string]interface{}, *http.Response, *common.XenditSdkError) {
	return r.ApiService.GetPaymentsByPaymentMethodIdExecute(r)
}

/*
GetPaymentsByPaymentMethodId Returns payments with matching PaymentMethodID.

Returns payments with matching PaymentMethodID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentMethodId
 @return ApiGetPaymentsByPaymentMethodIdRequest
*/
func (a *PaymentMethodApiService) GetPaymentsByPaymentMethodId(ctx context.Context, paymentMethodId string) ApiGetPaymentsByPaymentMethodIdRequest {
	return ApiGetPaymentsByPaymentMethodIdRequest{
		ApiService: a,
		ctx: ctx,
		paymentMethodId: paymentMethodId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *PaymentMethodApiService) GetPaymentsByPaymentMethodIdExecute(r ApiGetPaymentsByPaymentMethodIdRequest) (map[string]interface{}, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "PaymentMethodApiService.GetPaymentsByPaymentMethodId")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentMethodApiService.GetPaymentsByPaymentMethodIdExecute")
	}

	localVarPath := localBasePath + "/v2/payment_methods/{paymentMethodId}/payments"
	localVarPath = strings.Replace(localVarPath, "{"+"paymentMethodId"+"}", url.PathEscape(utils.ParameterValueToString(r.paymentMethodId, "paymentMethodId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.paymentRequestId != nil {
		t := *r.paymentRequestId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "payment_request_id", s.Index(i), "multi")
			}
		} else {
			utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "payment_request_id", t, "multi")
		}
	}
	if r.paymentMethodId2 != nil {
		t := *r.paymentMethodId2
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "payment_method_id", s.Index(i), "multi")
			}
		} else {
			utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "payment_method_id", t, "multi")
		}
	}
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
	if r.paymentMethodType != nil {
		t := *r.paymentMethodType
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "payment_method_type", s.Index(i), "multi")
			}
		} else {
			utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "payment_method_type", t, "multi")
		}
	}
	if r.channelCode != nil {
		t := *r.channelCode
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "channel_code", s.Index(i), "multi")
			}
		} else {
			utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "channel_code", t, "multi")
		}
	}
	if r.status != nil {
		t := *r.status
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "status", s.Index(i), "multi")
			}
		} else {
			utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "status", t, "multi")
		}
	}
	if r.currency != nil {
		t := *r.currency
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "currency", s.Index(i), "multi")
			}
		} else {
			utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "currency", t, "multi")
		}
	}
	if r.createdGte != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "created[gte]", r.createdGte, "")
	}
	if r.createdLte != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "created[lte]", r.createdLte, "")
	}
	if r.updatedGte != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "updated[gte]", r.updatedGte, "")
	}
	if r.updatedLte != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "updated[lte]", r.updatedLte, "")
	}
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
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentMethodApiService.GetPaymentsByPaymentMethodIdExecute")
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
    if err != nil {
        return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", fmt.Sprintf("Error creating HTTP request: PaymentMethodApiService.GetPaymentsByPaymentMethodIdExecute: %v", err))
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

type ApiPatchPaymentMethodRequest struct {
	ctx context.Context
	ApiService PaymentMethodApi
	paymentMethodId string
	forUserId *string
	paymentMethodUpdateParameters *PaymentMethodUpdateParameters
}

func (r ApiPatchPaymentMethodRequest) ForUserId(forUserId string) ApiPatchPaymentMethodRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiPatchPaymentMethodRequest) PaymentMethodUpdateParameters(paymentMethodUpdateParameters PaymentMethodUpdateParameters) ApiPatchPaymentMethodRequest {
	r.paymentMethodUpdateParameters = &paymentMethodUpdateParameters
	return r
}

func (r ApiPatchPaymentMethodRequest) Execute() (*PaymentMethod, *http.Response, *common.XenditSdkError) {
	return r.ApiService.PatchPaymentMethodExecute(r)
}

/*
PatchPaymentMethod Patch payment methods

This endpoint is used to toggle the ```status``` of an e-Wallet or a Direct Debit payment method to ```ACTIVE``` or ```INACTIVE```. This is also used to update the details of an Over-the-Counter or a Virtual Account payment method.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentMethodId
 @return ApiPatchPaymentMethodRequest
*/
func (a *PaymentMethodApiService) PatchPaymentMethod(ctx context.Context, paymentMethodId string) ApiPatchPaymentMethodRequest {
	return ApiPatchPaymentMethodRequest{
		ApiService: a,
		ctx: ctx,
		paymentMethodId: paymentMethodId,
	}
}

// Execute executes the request
//  @return PaymentMethod
func (a *PaymentMethodApiService) PatchPaymentMethodExecute(r ApiPatchPaymentMethodRequest) (*PaymentMethod, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *PaymentMethod
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "PaymentMethodApiService.PatchPaymentMethod")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentMethodApiService.PatchPaymentMethodExecute")
	}

	localVarPath := localBasePath + "/v2/payment_methods/{paymentMethodId}"
	localVarPath = strings.Replace(localVarPath, "{"+"paymentMethodId"+"}", url.PathEscape(utils.ParameterValueToString(r.paymentMethodId, "paymentMethodId")), -1)

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
	localVarPostBody = r.paymentMethodUpdateParameters
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentMethodApiService.PatchPaymentMethodExecute")
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
    if err != nil {
        return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", fmt.Sprintf("Error creating HTTP request: PaymentMethodApiService.PatchPaymentMethodExecute: %v", err))
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

type ApiGetAllPaymentMethodsRequest struct {
	ctx context.Context
	ApiService PaymentMethodApi
	forUserId *string
	id *[]string
	type_ *[]string
	status *[]PaymentMethodStatus
	reusability *PaymentMethodReusability
	customerId *string
	referenceId *string
	afterId *string
	beforeId *string
	limit *int32
}

func (r ApiGetAllPaymentMethodsRequest) ForUserId(forUserId string) ApiGetAllPaymentMethodsRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiGetAllPaymentMethodsRequest) Id(id []string) ApiGetAllPaymentMethodsRequest {
	r.id = &id
	return r
}

func (r ApiGetAllPaymentMethodsRequest) Type_(type_ []string) ApiGetAllPaymentMethodsRequest {
	r.type_ = &type_
	return r
}

func (r ApiGetAllPaymentMethodsRequest) Status(status []PaymentMethodStatus) ApiGetAllPaymentMethodsRequest {
	r.status = &status
	return r
}

func (r ApiGetAllPaymentMethodsRequest) Reusability(reusability PaymentMethodReusability) ApiGetAllPaymentMethodsRequest {
	r.reusability = &reusability
	return r
}

func (r ApiGetAllPaymentMethodsRequest) CustomerId(customerId string) ApiGetAllPaymentMethodsRequest {
	r.customerId = &customerId
	return r
}

func (r ApiGetAllPaymentMethodsRequest) ReferenceId(referenceId string) ApiGetAllPaymentMethodsRequest {
	r.referenceId = &referenceId
	return r
}

func (r ApiGetAllPaymentMethodsRequest) AfterId(afterId string) ApiGetAllPaymentMethodsRequest {
	r.afterId = &afterId
	return r
}

func (r ApiGetAllPaymentMethodsRequest) BeforeId(beforeId string) ApiGetAllPaymentMethodsRequest {
	r.beforeId = &beforeId
	return r
}

func (r ApiGetAllPaymentMethodsRequest) Limit(limit int32) ApiGetAllPaymentMethodsRequest {
	r.limit = &limit
	return r
}

func (r ApiGetAllPaymentMethodsRequest) Execute() (*PaymentMethodList, *http.Response, *common.XenditSdkError) {
	return r.ApiService.GetAllPaymentMethodsExecute(r)
}

/*
GetAllPaymentMethods Get all payment methods by filters

Get all payment methods by filters

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllPaymentMethodsRequest
*/
func (a *PaymentMethodApiService) GetAllPaymentMethods(ctx context.Context) ApiGetAllPaymentMethodsRequest {
	return ApiGetAllPaymentMethodsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaymentMethodList
func (a *PaymentMethodApiService) GetAllPaymentMethodsExecute(r ApiGetAllPaymentMethodsRequest) (*PaymentMethodList, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *PaymentMethodList
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "PaymentMethodApiService.GetAllPaymentMethods")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentMethodApiService.GetAllPaymentMethodsExecute")
	}

	localVarPath := localBasePath + "/v2/payment_methods"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.type_ != nil {
		t := *r.type_
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", s.Index(i), "multi")
			}
		} else {
			utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", t, "multi")
		}
	}
	if r.status != nil {
		t := *r.status
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "status", s.Index(i), "multi")
			}
		} else {
			utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "status", t, "multi")
		}
	}
	if r.reusability != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "reusability", r.reusability, "")
	}
	if r.customerId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "customer_id", r.customerId, "")
	}
	if r.referenceId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "reference_id", r.referenceId, "")
	}
	if r.afterId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "after_id", r.afterId, "")
	}
	if r.beforeId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "before_id", r.beforeId, "")
	}
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
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentMethodApiService.GetAllPaymentMethodsExecute")
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
    if err != nil {
        return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", fmt.Sprintf("Error creating HTTP request: PaymentMethodApiService.GetAllPaymentMethodsExecute: %v", err))
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

type ApiExpirePaymentMethodRequest struct {
	ctx context.Context
	ApiService PaymentMethodApi
	paymentMethodId string
	forUserId *string
	paymentMethodExpireParameters *PaymentMethodExpireParameters
}

func (r ApiExpirePaymentMethodRequest) ForUserId(forUserId string) ApiExpirePaymentMethodRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiExpirePaymentMethodRequest) PaymentMethodExpireParameters(paymentMethodExpireParameters PaymentMethodExpireParameters) ApiExpirePaymentMethodRequest {
	r.paymentMethodExpireParameters = &paymentMethodExpireParameters
	return r
}

func (r ApiExpirePaymentMethodRequest) Execute() (*PaymentMethod, *http.Response, *common.XenditSdkError) {
	return r.ApiService.ExpirePaymentMethodExecute(r)
}

/*
ExpirePaymentMethod Expires a payment method

This endpoint expires a payment method and performs unlinking if necessary

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentMethodId
 @return ApiExpirePaymentMethodRequest
*/
func (a *PaymentMethodApiService) ExpirePaymentMethod(ctx context.Context, paymentMethodId string) ApiExpirePaymentMethodRequest {
	return ApiExpirePaymentMethodRequest{
		ApiService: a,
		ctx: ctx,
		paymentMethodId: paymentMethodId,
	}
}

// Execute executes the request
//  @return PaymentMethod
func (a *PaymentMethodApiService) ExpirePaymentMethodExecute(r ApiExpirePaymentMethodRequest) (*PaymentMethod, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *PaymentMethod
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "PaymentMethodApiService.ExpirePaymentMethod")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentMethodApiService.ExpirePaymentMethodExecute")
	}

	localVarPath := localBasePath + "/v2/payment_methods/{paymentMethodId}/expire"
	localVarPath = strings.Replace(localVarPath, "{"+"paymentMethodId"+"}", url.PathEscape(utils.ParameterValueToString(r.paymentMethodId, "paymentMethodId")), -1)

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
	localVarPostBody = r.paymentMethodExpireParameters
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentMethodApiService.ExpirePaymentMethodExecute")
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
    if err != nil {
        return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", fmt.Sprintf("Error creating HTTP request: PaymentMethodApiService.ExpirePaymentMethodExecute: %v", err))
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

type ApiAuthPaymentMethodRequest struct {
	ctx context.Context
	ApiService PaymentMethodApi
	paymentMethodId string
	forUserId *string
	paymentMethodAuthParameters *PaymentMethodAuthParameters
}

func (r ApiAuthPaymentMethodRequest) ForUserId(forUserId string) ApiAuthPaymentMethodRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiAuthPaymentMethodRequest) PaymentMethodAuthParameters(paymentMethodAuthParameters PaymentMethodAuthParameters) ApiAuthPaymentMethodRequest {
	r.paymentMethodAuthParameters = &paymentMethodAuthParameters
	return r
}

func (r ApiAuthPaymentMethodRequest) Execute() (*PaymentMethod, *http.Response, *common.XenditSdkError) {
	return r.ApiService.AuthPaymentMethodExecute(r)
}

/*
AuthPaymentMethod Validate a payment method's linking OTP

This endpoint validates a payment method linking OTP

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentMethodId
 @return ApiAuthPaymentMethodRequest
*/
func (a *PaymentMethodApiService) AuthPaymentMethod(ctx context.Context, paymentMethodId string) ApiAuthPaymentMethodRequest {
	return ApiAuthPaymentMethodRequest{
		ApiService: a,
		ctx: ctx,
		paymentMethodId: paymentMethodId,
	}
}

// Execute executes the request
//  @return PaymentMethod
func (a *PaymentMethodApiService) AuthPaymentMethodExecute(r ApiAuthPaymentMethodRequest) (*PaymentMethod, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *PaymentMethod
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "PaymentMethodApiService.AuthPaymentMethod")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentMethodApiService.AuthPaymentMethodExecute")
	}

	localVarPath := localBasePath + "/v2/payment_methods/{paymentMethodId}/auth"
	localVarPath = strings.Replace(localVarPath, "{"+"paymentMethodId"+"}", url.PathEscape(utils.ParameterValueToString(r.paymentMethodId, "paymentMethodId")), -1)

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
	localVarPostBody = r.paymentMethodAuthParameters
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentMethodApiService.AuthPaymentMethodExecute")
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
    if err != nil {
        return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", fmt.Sprintf("Error creating HTTP request: PaymentMethodApiService.AuthPaymentMethodExecute: %v", err))
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

type ApiSimulatePaymentRequest struct {
	ctx context.Context
	ApiService PaymentMethodApi
	paymentMethodId string
	simulatePaymentRequest *SimulatePaymentRequest
}

func (r ApiSimulatePaymentRequest) SimulatePaymentRequest(simulatePaymentRequest SimulatePaymentRequest) ApiSimulatePaymentRequest {
	r.simulatePaymentRequest = &simulatePaymentRequest
	return r
}

func (r ApiSimulatePaymentRequest) Execute() (*http.Response, *common.XenditSdkError) {
	return r.ApiService.SimulatePaymentExecute(r)
}

/*
SimulatePayment Makes payment with matching PaymentMethodID.

Makes payment with matching PaymentMethodID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentMethodId
 @return ApiSimulatePaymentRequest
*/
func (a *PaymentMethodApiService) SimulatePayment(ctx context.Context, paymentMethodId string) ApiSimulatePaymentRequest {
	return ApiSimulatePaymentRequest{
		ApiService: a,
		ctx: ctx,
		paymentMethodId: paymentMethodId,
	}
}

// Execute executes the request
func (a *PaymentMethodApiService) SimulatePaymentExecute(r ApiSimulatePaymentRequest) (*http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []common.FormFile
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "PaymentMethodApiService.SimulatePayment")
	if err != nil {
		return nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentMethodApiService.SimulatePaymentExecute")
	}

	localVarPath := localBasePath + "/v2/payment_methods/{paymentMethodId}/payments/simulate"
	localVarPath = strings.Replace(localVarPath, "{"+"paymentMethodId"+"}", url.PathEscape(utils.ParameterValueToString(r.paymentMethodId, "paymentMethodId")), -1)

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
	// body params
	localVarPostBody = r.simulatePaymentRequest
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: PaymentMethodApiService.SimulatePaymentExecute")
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
    if err != nil {
        return nil, common.NewXenditSdkError(nil, "", fmt.Sprintf("Error creating HTTP request: PaymentMethodApiService.SimulatePaymentExecute: %v", err))
    }
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))

    

	if err != nil || localVarHTTPResponse.StatusCode < 200 || localVarHTTPResponse.StatusCode >= 300 {
		xenditSdkError := common.NewXenditSdkError(&localVarBody, strconv.Itoa(localVarHTTPResponse.StatusCode), localVarHTTPResponse.Status)

		return localVarHTTPResponse, xenditSdkError
	}

	return localVarHTTPResponse, nil
}
