package httpx

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/ritwickdey/querydecoder"
)

// Decode parses query and body into give struct
func Decode(r *http.Request, v any) (err error) {
	err = DecodeQuery(r, v)
	if err != nil {
		return err
	}

	switch render.GetRequestContentType(r) {
	case render.ContentTypeJSON:
		err = render.DecodeJSON(r.Body, v)
	case render.ContentTypeXML:
		err = render.DecodeXML(r.Body, v)
	case render.ContentTypeForm:
		err = render.DecodeForm(r.Body, v)
	}

	return err
}

// DecodeTo decodes to generic type
// T must be inferred
func DecodeTo[T any](r *http.Request) (*T, error) {
	in := new(T)
	if err := Decode(r, in); err != nil {
		return nil, err
	}

	return in, nil
}

// DecodeQuery decodes query params
func DecodeQuery(r *http.Request, v any) error {
	return querydecoder.New(r.URL.Query()).Decode(v)
}
