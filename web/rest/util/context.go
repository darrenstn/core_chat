package util

import (
	"context"
	"net/http"
)

type contextKey string

const identifierKey contextKey = "identifier"

func WithIdentifier(ctx context.Context, identifier string) context.Context {
	return context.WithValue(ctx, identifierKey, identifier)
}

func GetIdentifier(r *http.Request) (string, bool) {
	val := r.Context().Value(identifierKey)
	identifier, ok := val.(string)
	return identifier, ok
}
