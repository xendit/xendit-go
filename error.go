package xendit

import "net/http"

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

// FromGoErr generate xendit.Error from generic go errors
func FromGoErr(err error) *Error {
	return &Error{
		Status:    http.StatusTeapot,
		ErrorCode: GoErrCode,
		Message:   err.Error(),
	}
}
