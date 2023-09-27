package common

import (
    "context"
    "encoding/json"
    "net/http"
    "net/url"
    "fmt"
)

type IConfiguration interface {
	AddDefaultHeader(key string, value string)
	ServerURL(index int, variables map[string]string) (string, error)
	ServerURLWithContext(ctx context.Context, endpoint string) (string, error)
}

type IClient interface {
	GetConfig() IConfiguration
	PrepareRequest(
		ctx context.Context,
		path string, method string,
		postBody interface{},
		headerParams map[string]string,
		queryParams url.Values,
		formParams url.Values,
		formFiles []FormFile) (localVarRequest *http.Request, err error)
	CallAPI(request *http.Request) (*http.Response, error)
	Decode(v interface{}, b []byte, contentType string) (err error)
}

type FormFile struct {
	FileBytes []byte
	FileName string
	FormFileName string
}

type XenditSdkError struct {
	rawResponse map[string]interface{}
	status string
	errorCode string
	errorMessage string
}

func NewXenditSdkError(response *[]byte, paramStatus string, paramStatusText string) *XenditSdkError {
	var _rawResponse map[string]interface{};

	err := json.Unmarshal(*response,&_rawResponse)
	if err != nil {
		_rawResponse = map[string]interface{}{}
	}

	_status := paramStatus

	if _status == "" {
		if val, ok := _rawResponse["status"]; ok {
			_status = fmt.Sprintf("%v", val)
		}
	}
	if _status == "" {
		if val, ok := _rawResponse["status_code"]; ok {
			_status = fmt.Sprintf("%v", val)
		}
	}
	if _status == "" {
		if val, ok := _rawResponse["statusCode"]; ok {
			_status = fmt.Sprintf("%v", val)
		}
	}

	_errorCode := ""
	if _errorCode == "" {
		if val, ok := _rawResponse["error"]; ok {
			_errorCode = fmt.Sprintf("%v", val)
		}
	}
	if _errorCode == "" {
		if val, ok := _rawResponse["error_code"]; ok {
			_errorCode = fmt.Sprintf("%v", val)
		}
	}
	if _errorCode == "" {
		if val, ok := _rawResponse["errorCode"]; ok {
			_errorCode = fmt.Sprintf("%v", val)
		}
	}

	_errorMessage := ""
	if _errorMessage == "" {
		if val, ok := _rawResponse["message"]; ok {
			_errorMessage = fmt.Sprintf("%v", val)
		}
	}
	if _errorMessage == "" {
		if val, ok := _rawResponse["error_message"]; ok {
			_errorMessage = fmt.Sprintf("%v", val)
		}
	}
	if _errorMessage == "" {
		if val, ok := _rawResponse["errorMessage"]; ok {
			_errorMessage = fmt.Sprintf("%v", val)
		}
	}
	if _errorMessage == "" {
		_errorMessage = paramStatusText
	}


	return &XenditSdkError{
		rawResponse: _rawResponse,
		status: _status,
		errorCode: _errorCode,
		errorMessage: _errorMessage,
	}
}

func (e XenditSdkError) Error() string {
	return e.errorMessage
}

func (e XenditSdkError) ErrorCode() string {
	return e.errorCode
}

func (e XenditSdkError) RawResponse() map[string]interface{} {
	return e.rawResponse
}

func (e XenditSdkError) Status() string {
	return e.status
}

func (e XenditSdkError) FullError() map[string]interface{} {
	return map[string]interface{}{
		"rawResponse": e.rawResponse,
		"status": e.status,
		"errorCode": e.errorCode,
		"errorMessage": e.errorMessage,
	}
}
