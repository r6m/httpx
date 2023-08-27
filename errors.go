package httpx

import (
	"fmt"
	"net/http"
)

// Error is a http error
type Error struct {
	Internal error  `xml:"-" json:"-"`
	Status   int    `xml:"status" json:"status"`
	Code     string `xml:"code" json:"code,omitempty"`
	Message  string `xml:"message,omitempty" json:"message,omitempty"`
	Data     any    `xml:"data,omitempty" json:"data,omitempty"`
}

// Error implements error
func (e *Error) Error() string {
	if e.Internal != nil {
		return e.Internal.Error()
	}

	return fmt.Sprintf("%d: %s", e.Status, e.Message)
}

// WithInternal adds error
func (e *Error) WithInternal(err error) *Error {
	e.Internal = err
	return e
}

// WithData adds data to body
// you can pass validation errors in data
func (e *Error) WithData(data any) *Error {
	e.Data = data
	return e
}

// NewError returns an http error
func NewError(status int, format string, args ...any) *Error {
	err := &Error{
		Status:  status,
		Code:    http.StatusText(status),
		Message: fmt.Sprintf(format, args...),
	}

	return err
}

// NotFoundError returns Error with status not found
func NotFoundError(format string, args ...any) *Error {
	return NewError(http.StatusNotFound, format, args...)
}

// BadRequestError returns Error with status BadRequest
func BadRequestError(format string, args ...any) *Error {
	return NewError(http.StatusBadRequest, format, args...)
}

// UnauthorizedError returns Error with status Unauthorized
func UnauthorizedError(format string, args ...any) *Error {
	return NewError(http.StatusUnauthorized, format, args...)
}

// InternalServerError returns Error with status Internal
func InternalServerError(format string, args ...any) *Error {
	return NewError(http.StatusInternalServerError, format, args...)
}

// PaymentRequiredError returns Error with status PaymentRequired
func PaymentRequiredError(format string, args ...any) *Error {
	return NewError(http.StatusPaymentRequired, format, args...)
}

// TooManyRequestsError returns Error with status TooManyRequests
func TooManyRequestsError(format string, args ...any) *Error {
	return NewError(http.StatusTooManyRequests, format, args...)
}

// ConflictError returns Error with status Conflict
func ConflictError(format string, args ...any) *Error {
	return NewError(http.StatusConflict, format, args...)
}

// NotImplementedError returns Error with status NotImplemented
func NotImplementedError(format string, args ...any) *Error {
	return NewError(http.StatusNotImplemented, format, args...)
}
