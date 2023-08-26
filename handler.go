package httpx

import (
	"net/http"
)

var (
	// RequestKey is http request context key
	RequestKey = &ContextKey{"req"}
)

// HandlerFunc wrapps http handler func with error
// e.g.   func GetUser(w http.ResponseWriter, r * http.Request) error
func HandlerFunc(handler func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			handleError(w, r, err)
		}
	}
}

// GenericHandlerFunc is a wapper to implement generic handler func
// e.g.   func GetUser(w http.ResponseWriter, r * http.Request, in *ExampleRequest) error
func GenericHandlerFunc[T any](handler func(w http.ResponseWriter, r *http.Request, in *T) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		in, err := DecodeTo[T](r)
		if err != nil {
			Respond(w, r, BadRequestError("can't decode body").WithInternal(err))
			return
		}

		err = handler(w, r, in)
		if err != nil {
			handleError(w, r, err)
		}
	}
}
