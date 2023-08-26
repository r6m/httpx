package httpx

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

// Respond parses error and encodes response based on request content-type
func Respond(w http.ResponseWriter, r *http.Request, v any) error {
	switch vv := v.(type) {
	case error:
		handleError(w, r, vv)
	default:
		render.Respond(w, r, vv)
	}

	return nil
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	switch e := err.(type) {
	case *Error:
		if e.Status >= http.StatusInternalServerError {
			logrus.WithField("status", http.StatusInternalServerError).WithField("internal", e.Internal).WithError(e).Error(e.Message)
		}
		w.WriteHeader(e.Status)
		render.Respond(w, r, e)
	default:
		logrus.WithField("status", "internal").WithError(err).Error("internal server error")
		render.Respond(w, r, InternalServerError("internal server error"))
		return
	}
}

// NotFound reponds NotFound error
func NotFound(w http.ResponseWriter, r *http.Request, msg string) {
	Respond(w, r, NewError(http.StatusNotFound, msg))
}

// BadRequest reponds BadRequest error
func BadRequest(w http.ResponseWriter, r *http.Request, msg string) {
	Respond(w, r, NewError(http.StatusBadRequest, msg))
}

// Unauthorized reponds Unauthorized error
func Unauthorized(w http.ResponseWriter, r *http.Request, msg string) {
	Respond(w, r, NewError(http.StatusUnauthorized, msg))
}

// PaymentRequired reponds PaymentRequired error
func PaymentRequired(w http.ResponseWriter, r *http.Request, msg string) {
	Respond(w, r, NewError(http.StatusPaymentRequired, msg))
}

// TooManyRequests reponds TooManyRequests error
func TooManyRequests(w http.ResponseWriter, r *http.Request, msg string) {
	Respond(w, r, NewError(http.StatusTooManyRequests, msg))
}

// Conflict reponds Conflict error
func Conflict(w http.ResponseWriter, r *http.Request, msg string) {
	Respond(w, r, NewError(http.StatusConflict, msg))
}

// InternalServer reponds InternalServer error
func InternalServer(w http.ResponseWriter, r *http.Request, msg string) {
	Respond(w, r, NewError(http.StatusInternalServerError, msg))
}

// NotImplemented reponds NotImplemented error
func NotImplemented(w http.ResponseWriter, r *http.Request, msg string) {
	Respond(w, r, NewError(http.StatusNotImplemented, msg))
}
