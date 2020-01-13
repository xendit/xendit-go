package xendit

import (
	"encoding/json"
	"net/http"
)

const (
	// APIValidationErrCode error code for parameters validation
	APIValidationErrCode string = "API_VALIDATION_ERROR"
	// GoErrCode error code for errors happen inside Go code
	GoErrCode string = "GO_ERROR"
)

// Error conventional Xendit error
type Error struct {
	Status    int    `json:"status,omitempty"`
	ErrorCode string `json:"error_code,omitempty"`
	Message   string `json:"message,omitempty"`
}

// Error return error message
// this enables xendit.Error to comply with Go error interface
func (e *Error) Error() string {
	return e.Message
}

// GetErrorCode returns error code coming from xendit backend
func (e *Error) GetErrorCode() string {
	return e.ErrorCode
}

// GetStatus returns http status code
func (e *Error) GetStatus() int {
	return e.Status
}

// FromGoErr generate xendit.Error from generic go errors
func FromGoErr(err error) *Error {
	return &Error{
		Status:    http.StatusTeapot,
		ErrorCode: GoErrCode,
		Message:   err.Error(),
	}
}

// FromHTTPErr generate xendit.Error from http errors with non 2xx status
func FromHTTPErr(status int, respBody []byte) *Error {
	var httpError *Error
	if err := json.Unmarshal(respBody, &httpError); err != nil {
		return FromGoErr(err)
	}
	httpError.Status = status

	return httpError
}
