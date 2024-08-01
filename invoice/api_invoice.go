package invoice

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
	"time"
)


type InvoiceApi interface {

	/*
	CreateInvoice Create an invoice

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateInvoiceRequest
	*/
	CreateInvoice(ctx context.Context) ApiCreateInvoiceRequest

	// CreateInvoiceExecute executes the request
	//  @return Invoice
	CreateInvoiceExecute(r ApiCreateInvoiceRequest) (*Invoice, *http.Response, *common.XenditSdkError)

	/*
	GetInvoiceById Get invoice by invoice id

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param invoiceId Invoice ID
	@return ApiGetInvoiceByIdRequest
	*/
	GetInvoiceById(ctx context.Context, invoiceId string) ApiGetInvoiceByIdRequest

	// GetInvoiceByIdExecute executes the request
	//  @return Invoice
	GetInvoiceByIdExecute(r ApiGetInvoiceByIdRequest) (*Invoice, *http.Response, *common.XenditSdkError)

	/*
	GetInvoices Get all Invoices

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetInvoicesRequest
	*/
	GetInvoices(ctx context.Context) ApiGetInvoicesRequest

	// GetInvoicesExecute executes the request
	//  @return []Invoice
	GetInvoicesExecute(r ApiGetInvoicesRequest) ([]Invoice, *http.Response, *common.XenditSdkError)

	/*
	ExpireInvoice Manually expire an invoice

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param invoiceId Invoice ID to be expired
	@return ApiExpireInvoiceRequest
	*/
	ExpireInvoice(ctx context.Context, invoiceId string) ApiExpireInvoiceRequest

	// ExpireInvoiceExecute executes the request
	//  @return Invoice
	ExpireInvoiceExecute(r ApiExpireInvoiceRequest) (*Invoice, *http.Response, *common.XenditSdkError)
}

// InvoiceApiService InvoiceApi service
type InvoiceApiService struct {
	client common.IClient
}

// NewInvoiceApi Create a new InvoiceApi service
func NewInvoiceApi (client common.IClient) InvoiceApi {
	return &InvoiceApiService{
		client: client,
	}
}


type ApiCreateInvoiceRequest struct {
	ctx context.Context
	ApiService InvoiceApi
	createInvoiceRequest *CreateInvoiceRequest
	forUserId *string
}

func (r ApiCreateInvoiceRequest) CreateInvoiceRequest(createInvoiceRequest CreateInvoiceRequest) ApiCreateInvoiceRequest {
	r.createInvoiceRequest = &createInvoiceRequest
	return r
}

// Business ID of the sub-account merchant (XP feature)
func (r ApiCreateInvoiceRequest) ForUserId(forUserId string) ApiCreateInvoiceRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiCreateInvoiceRequest) Execute() (*Invoice, *http.Response, *common.XenditSdkError) {
	return r.ApiService.CreateInvoiceExecute(r)
}

/*
CreateInvoice Create an invoice

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateInvoiceRequest
*/
func (a *InvoiceApiService) CreateInvoice(ctx context.Context) ApiCreateInvoiceRequest {
	return ApiCreateInvoiceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Invoice
func (a *InvoiceApiService) CreateInvoiceExecute(r ApiCreateInvoiceRequest) (*Invoice, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *Invoice
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "InvoiceApiService.CreateInvoice")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: InvoiceApiService.CreateInvoiceExecute")
	}

	localVarPath := localBasePath + "/v2/invoices/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createInvoiceRequest == nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "createInvoiceRequest is required and must be specified")
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
	if r.forUserId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarHeaderParams, "for-user-id", r.forUserId, "")
	}
	// body params
	localVarPostBody = r.createInvoiceRequest
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: InvoiceApiService.CreateInvoiceExecute")
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

type ApiGetInvoiceByIdRequest struct {
	ctx context.Context
	ApiService InvoiceApi
	invoiceId string
	forUserId *string
}

// Business ID of the sub-account merchant (XP feature)
func (r ApiGetInvoiceByIdRequest) ForUserId(forUserId string) ApiGetInvoiceByIdRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiGetInvoiceByIdRequest) Execute() (*Invoice, *http.Response, *common.XenditSdkError) {
	return r.ApiService.GetInvoiceByIdExecute(r)
}

/*
GetInvoiceById Get invoice by invoice id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param invoiceId Invoice ID
 @return ApiGetInvoiceByIdRequest
*/
func (a *InvoiceApiService) GetInvoiceById(ctx context.Context, invoiceId string) ApiGetInvoiceByIdRequest {
	return ApiGetInvoiceByIdRequest{
		ApiService: a,
		ctx: ctx,
		invoiceId: invoiceId,
	}
}

// Execute executes the request
//  @return Invoice
func (a *InvoiceApiService) GetInvoiceByIdExecute(r ApiGetInvoiceByIdRequest) (*Invoice, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *Invoice
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "InvoiceApiService.GetInvoiceById")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: InvoiceApiService.GetInvoiceByIdExecute")
	}

	localVarPath := localBasePath + "/v2/invoices/{invoice_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"invoice_id"+"}", url.PathEscape(utils.ParameterValueToString(r.invoiceId, "invoiceId")), -1)

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
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: InvoiceApiService.GetInvoiceByIdExecute")
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

type ApiGetInvoicesRequest struct {
	ctx context.Context
	ApiService InvoiceApi
	forUserId *string
	externalId *string
	statuses *[]InvoiceStatus
	limit *float32
	createdAfter *time.Time
	createdBefore *time.Time
	paidAfter *time.Time
	paidBefore *time.Time
	expiredAfter *time.Time
	expiredBefore *time.Time
	lastInvoice *string
	clientTypes *[]InvoiceClientType
	paymentChannels *[]string
	onDemandLink *string
	recurringPaymentId *string
}

// Business ID of the sub-account merchant (XP feature)
func (r ApiGetInvoicesRequest) ForUserId(forUserId string) ApiGetInvoicesRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiGetInvoicesRequest) ExternalId(externalId string) ApiGetInvoicesRequest {
	r.externalId = &externalId
	return r
}

func (r ApiGetInvoicesRequest) Statuses(statuses []InvoiceStatus) ApiGetInvoicesRequest {
	r.statuses = &statuses
	return r
}

func (r ApiGetInvoicesRequest) Limit(limit float32) ApiGetInvoicesRequest {
	r.limit = &limit
	return r
}

func (r ApiGetInvoicesRequest) CreatedAfter(createdAfter time.Time) ApiGetInvoicesRequest {
	r.createdAfter = &createdAfter
	return r
}

func (r ApiGetInvoicesRequest) CreatedBefore(createdBefore time.Time) ApiGetInvoicesRequest {
	r.createdBefore = &createdBefore
	return r
}

func (r ApiGetInvoicesRequest) PaidAfter(paidAfter time.Time) ApiGetInvoicesRequest {
	r.paidAfter = &paidAfter
	return r
}

func (r ApiGetInvoicesRequest) PaidBefore(paidBefore time.Time) ApiGetInvoicesRequest {
	r.paidBefore = &paidBefore
	return r
}

func (r ApiGetInvoicesRequest) ExpiredAfter(expiredAfter time.Time) ApiGetInvoicesRequest {
	r.expiredAfter = &expiredAfter
	return r
}

func (r ApiGetInvoicesRequest) ExpiredBefore(expiredBefore time.Time) ApiGetInvoicesRequest {
	r.expiredBefore = &expiredBefore
	return r
}

func (r ApiGetInvoicesRequest) LastInvoice(lastInvoice string) ApiGetInvoicesRequest {
	r.lastInvoice = &lastInvoice
	return r
}

func (r ApiGetInvoicesRequest) ClientTypes(clientTypes []InvoiceClientType) ApiGetInvoicesRequest {
	r.clientTypes = &clientTypes
	return r
}

func (r ApiGetInvoicesRequest) PaymentChannels(paymentChannels []string) ApiGetInvoicesRequest {
	r.paymentChannels = &paymentChannels
	return r
}

func (r ApiGetInvoicesRequest) OnDemandLink(onDemandLink string) ApiGetInvoicesRequest {
	r.onDemandLink = &onDemandLink
	return r
}

func (r ApiGetInvoicesRequest) RecurringPaymentId(recurringPaymentId string) ApiGetInvoicesRequest {
	r.recurringPaymentId = &recurringPaymentId
	return r
}

func (r ApiGetInvoicesRequest) Execute() ([]Invoice, *http.Response, *common.XenditSdkError) {
	return r.ApiService.GetInvoicesExecute(r)
}

/*
GetInvoices Get all Invoices

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetInvoicesRequest
*/
func (a *InvoiceApiService) GetInvoices(ctx context.Context) ApiGetInvoicesRequest {
	return ApiGetInvoicesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Invoice
func (a *InvoiceApiService) GetInvoicesExecute(r ApiGetInvoicesRequest) ([]Invoice, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  []Invoice
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "InvoiceApiService.GetInvoices")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: InvoiceApiService.GetInvoicesExecute")
	}

	localVarPath := localBasePath + "/v2/invoices"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.externalId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "external_id", r.externalId, "")
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
	if r.limit != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.createdAfter != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "created_after", r.createdAfter, "")
	}
	if r.createdBefore != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "created_before", r.createdBefore, "")
	}
	if r.paidAfter != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "paid_after", r.paidAfter, "")
	}
	if r.paidBefore != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "paid_before", r.paidBefore, "")
	}
	if r.expiredAfter != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "expired_after", r.expiredAfter, "")
	}
	if r.expiredBefore != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "expired_before", r.expiredBefore, "")
	}
	if r.lastInvoice != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "last_invoice", r.lastInvoice, "")
	}
	if r.clientTypes != nil {
		t := *r.clientTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "client_types", s.Index(i), "multi")
			}
		} else {
			utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "client_types", t, "multi")
		}
	}
	if r.paymentChannels != nil {
		t := *r.paymentChannels
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "payment_channels", s.Index(i), "multi")
			}
		} else {
			utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "payment_channels", t, "multi")
		}
	}
	if r.onDemandLink != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "on_demand_link", r.onDemandLink, "")
	}
	if r.recurringPaymentId != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "recurring_payment_id", r.recurringPaymentId, "")
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
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: InvoiceApiService.GetInvoicesExecute")
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

type ApiExpireInvoiceRequest struct {
	ctx context.Context
	ApiService InvoiceApi
	invoiceId string
	forUserId *string
}

// Business ID of the sub-account merchant (XP feature)
func (r ApiExpireInvoiceRequest) ForUserId(forUserId string) ApiExpireInvoiceRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiExpireInvoiceRequest) Execute() (*Invoice, *http.Response, *common.XenditSdkError) {
	return r.ApiService.ExpireInvoiceExecute(r)
}

/*
ExpireInvoice Manually expire an invoice

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param invoiceId Invoice ID to be expired
 @return ApiExpireInvoiceRequest
*/
func (a *InvoiceApiService) ExpireInvoice(ctx context.Context, invoiceId string) ApiExpireInvoiceRequest {
	return ApiExpireInvoiceRequest{
		ApiService: a,
		ctx: ctx,
		invoiceId: invoiceId,
	}
}

// Execute executes the request
//  @return Invoice
func (a *InvoiceApiService) ExpireInvoiceExecute(r ApiExpireInvoiceRequest) (*Invoice, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *Invoice
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "InvoiceApiService.ExpireInvoice")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: InvoiceApiService.ExpireInvoiceExecute")
	}

	localVarPath := localBasePath + "/invoices/{invoice_id}/expire!"
	localVarPath = strings.Replace(localVarPath, "{"+"invoice_id"+"}", url.PathEscape(utils.ParameterValueToString(r.invoiceId, "invoiceId")), -1)

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
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: InvoiceApiService.ExpireInvoiceExecute")
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
