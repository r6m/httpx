package httpx

import (
	"net/http"
)

// ContextKey warpps context key name
type ContextKey struct {
	Name string
}

// FromContext returns specified context value
func FromContext[T any](r *http.Request, key any) *T {
	value := r.Context().Value(key)
	if value != nil {
		return value.(*T)
	}

	return nil
}
