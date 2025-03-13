package customer

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
)


type CustomerApi interface {

	/*
	CreateCustomer Create Customer

	Function to create a customer that you may use in your Invoice or Payment Requests. For detail explanations, see this link: https://developers.xendit.co/api-reference/#create-customer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateCustomerRequest
	*/
	CreateCustomer(ctx context.Context) ApiCreateCustomerRequest

	// CreateCustomerExecute executes the request
	//  @return Customer
	CreateCustomerExecute(r ApiCreateCustomerRequest) (*Customer, *http.Response, *common.XenditSdkError)

	/*
	GetCustomer Get Customer By ID

	Retrieves a single customer object For detail explanations, see this link: https://developers.xendit.co/api-reference/#get-customer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id End customer resource id
	@return ApiGetCustomerRequest
	*/
	GetCustomer(ctx context.Context, id string) ApiGetCustomerRequest

	// GetCustomerExecute executes the request
	//  @return Customer
	GetCustomerExecute(r ApiGetCustomerRequest) (*Customer, *http.Response, *common.XenditSdkError)

	/*
	GetCustomerByReferenceID GET customers by reference id

	Retrieves an array with a customer object that matches the provided reference_id - the identifier provided by you For detail explanations, see this link: https://developers.xendit.co/api-reference/#get-customer-by-reference-id

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetCustomerByReferenceIDRequest
	*/
	GetCustomerByReferenceID(ctx context.Context) ApiGetCustomerByReferenceIDRequest

	// GetCustomerByReferenceIDExecute executes the request
	//  @return GetCustomerByReferenceID200Response
	GetCustomerByReferenceIDExecute(r ApiGetCustomerByReferenceIDRequest) (*GetCustomerByReferenceID200Response, *http.Response, *common.XenditSdkError)

	/*
	UpdateCustomer Update End Customer Resource

	Function to update an existing customer. For a detailed explanation For detail explanations, see this link: https://developers.xendit.co/api-reference/#update-customer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id End customer resource id
	@return ApiUpdateCustomerRequest
	*/
	UpdateCustomer(ctx context.Context, id string) ApiUpdateCustomerRequest

	// UpdateCustomerExecute executes the request
	//  @return Customer
	UpdateCustomerExecute(r ApiUpdateCustomerRequest) (*Customer, *http.Response, *common.XenditSdkError)
}

// CustomerApiService CustomerApi service
type CustomerApiService struct {
	client common.IClient
}

// NewCustomerApi Create a new CustomerApi service
func NewCustomerApi (client common.IClient) CustomerApi {
	return &CustomerApiService{
		client: client,
	}
}


type ApiCreateCustomerRequest struct {
	ctx context.Context
	ApiService CustomerApi
	idempotencyKey *string
	forUserId *string
	customerRequest *CustomerRequest
}

// A unique key to prevent processing duplicate requests.
func (r ApiCreateCustomerRequest) IdempotencyKey(idempotencyKey string) ApiCreateCustomerRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

// The sub-account user-id that you want to make this transaction for.
func (r ApiCreateCustomerRequest) ForUserId(forUserId string) ApiCreateCustomerRequest {
	r.forUserId = &forUserId
	return r
}

// Request object for end customer object
func (r ApiCreateCustomerRequest) CustomerRequest(customerRequest CustomerRequest) ApiCreateCustomerRequest {
	r.customerRequest = &customerRequest
	return r
}

func (r ApiCreateCustomerRequest) Execute() (*Customer, *http.Response, *common.XenditSdkError) {
	return r.ApiService.CreateCustomerExecute(r)
}

/*
CreateCustomer Create Customer

Function to create a customer that you may use in your Invoice or Payment Requests. For detail explanations, see this link: https://developers.xendit.co/api-reference/#create-customer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateCustomerRequest
*/
func (a *CustomerApiService) CreateCustomer(ctx context.Context) ApiCreateCustomerRequest {
	return ApiCreateCustomerRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Customer
func (a *CustomerApiService) CreateCustomerExecute(r ApiCreateCustomerRequest) (*Customer, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *Customer
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "CustomerApiService.CreateCustomer")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: CustomerApiService.CreateCustomerExecute")
	}

	localVarPath := localBasePath + "/customers"

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
	localVarPostBody = r.customerRequest
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: CustomerApiService.CreateCustomerExecute")
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
        if err != nil {
            return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", fmt.Sprintf("Error calling API: CustomerApiService.CreateCustomerExecute: %v", err))
        }

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
        if err != nil {
            return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", fmt.Sprintf("Error parsing HTTP response: CustomerApiService.CreateCustomerExecute: %v", err))
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

type ApiGetCustomerRequest struct {
	ctx context.Context
	ApiService CustomerApi
	id string
	forUserId *string
}

// The sub-account user-id that you want to make this transaction for.
func (r ApiGetCustomerRequest) ForUserId(forUserId string) ApiGetCustomerRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiGetCustomerRequest) Execute() (*Customer, *http.Response, *common.XenditSdkError) {
	return r.ApiService.GetCustomerExecute(r)
}

/*
GetCustomer Get Customer By ID

Retrieves a single customer object For detail explanations, see this link: https://developers.xendit.co/api-reference/#get-customer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id End customer resource id
 @return ApiGetCustomerRequest
*/
func (a *CustomerApiService) GetCustomer(ctx context.Context, id string) ApiGetCustomerRequest {
	return ApiGetCustomerRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Customer
func (a *CustomerApiService) GetCustomerExecute(r ApiGetCustomerRequest) (*Customer, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *Customer
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "CustomerApiService.GetCustomer")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: CustomerApiService.GetCustomerExecute")
	}

	localVarPath := localBasePath + "/customers/{id}"
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
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: CustomerApiService.GetCustomerExecute")
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
        if err != nil {
            return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", fmt.Sprintf("Error calling API: CustomerApiService.GetCustomerExecute: %v", err))
        }

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
        if err != nil {
            return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", fmt.Sprintf("Error parsing HTTP response: CustomerApiService.GetCustomerExecute: %v", err))
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

type ApiGetCustomerByReferenceIDRequest struct {
	ctx context.Context
	ApiService CustomerApi
	referenceId *string
	forUserId *string
}

// Merchant&#39;s reference of end customer
func (r ApiGetCustomerByReferenceIDRequest) ReferenceId(referenceId string) ApiGetCustomerByReferenceIDRequest {
	r.referenceId = &referenceId
	return r
}

// The sub-account user-id that you want to make this transaction for.
func (r ApiGetCustomerByReferenceIDRequest) ForUserId(forUserId string) ApiGetCustomerByReferenceIDRequest {
	r.forUserId = &forUserId
	return r
}

func (r ApiGetCustomerByReferenceIDRequest) Execute() (*GetCustomerByReferenceID200Response, *http.Response, *common.XenditSdkError) {
	return r.ApiService.GetCustomerByReferenceIDExecute(r)
}

/*
GetCustomerByReferenceID GET customers by reference id

Retrieves an array with a customer object that matches the provided reference_id - the identifier provided by you For detail explanations, see this link: https://developers.xendit.co/api-reference/#get-customer-by-reference-id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCustomerByReferenceIDRequest
*/
func (a *CustomerApiService) GetCustomerByReferenceID(ctx context.Context) ApiGetCustomerByReferenceIDRequest {
	return ApiGetCustomerByReferenceIDRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetCustomerByReferenceID200Response
func (a *CustomerApiService) GetCustomerByReferenceIDExecute(r ApiGetCustomerByReferenceIDRequest) (*GetCustomerByReferenceID200Response, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *GetCustomerByReferenceID200Response
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "CustomerApiService.GetCustomerByReferenceID")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: CustomerApiService.GetCustomerByReferenceIDExecute")
	}

	localVarPath := localBasePath + "/customers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.referenceId == nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "referenceId is required and must be specified")
	}
	if utils.Strlen(*r.referenceId) > 255 {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "referenceId must have less than 255 elements")
	}

	utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "reference_id", r.referenceId, "")
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
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: CustomerApiService.GetCustomerByReferenceIDExecute")
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
        if err != nil {
            return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", fmt.Sprintf("Error calling API: CustomerApiService.GetCustomerByReferenceIDExecute: %v", err))
        }

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
        if err != nil {
            return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", fmt.Sprintf("Error parsing HTTP response: CustomerApiService.GetCustomerByReferenceIDExecute: %v", err))
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

type ApiUpdateCustomerRequest struct {
	ctx context.Context
	ApiService CustomerApi
	id string
	forUserId *string
	patchCustomer *PatchCustomer
}

// The sub-account user-id that you want to make this transaction for.
func (r ApiUpdateCustomerRequest) ForUserId(forUserId string) ApiUpdateCustomerRequest {
	r.forUserId = &forUserId
	return r
}

// Update Request for end customer object
func (r ApiUpdateCustomerRequest) PatchCustomer(patchCustomer PatchCustomer) ApiUpdateCustomerRequest {
	r.patchCustomer = &patchCustomer
	return r
}

func (r ApiUpdateCustomerRequest) Execute() (*Customer, *http.Response, *common.XenditSdkError) {
	return r.ApiService.UpdateCustomerExecute(r)
}

/*
UpdateCustomer Update End Customer Resource

Function to update an existing customer. For a detailed explanation For detail explanations, see this link: https://developers.xendit.co/api-reference/#update-customer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id End customer resource id
 @return ApiUpdateCustomerRequest
*/
func (a *CustomerApiService) UpdateCustomer(ctx context.Context, id string) ApiUpdateCustomerRequest {
	return ApiUpdateCustomerRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Customer
func (a *CustomerApiService) UpdateCustomerExecute(r ApiUpdateCustomerRequest) (*Customer, *http.Response, *common.XenditSdkError) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []common.FormFile
		localVarReturnValue  *Customer
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "CustomerApiService.UpdateCustomer")
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: CustomerApiService.UpdateCustomerExecute")
	}

	localVarPath := localBasePath + "/customers/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(utils.ParameterValueToString(r.id, "id")), -1)

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
	localVarPostBody = r.patchCustomer
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", "Error creating HTTP request: CustomerApiService.UpdateCustomerExecute")
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
        if err != nil {
            return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", fmt.Sprintf("Error calling API: CustomerApiService.UpdateCustomerExecute: %v", err))
        }

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
        if err != nil {
            return localVarReturnValue, nil, common.NewXenditSdkError(nil, "", fmt.Sprintf("Error parsing HTTP response: CustomerApiService.UpdateCustomerExecute: %v", err))
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
