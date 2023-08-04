package headerutils

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

const (
	// XRequestIDKey defines x-request-id key string.
	XRequestIDKey = "x-request-id"
	// UsernameKey defines username key string.
	UsernameKey = "username"
)

// GetRequestID request id from header
func GetRequestID(ctx context.Context) string {
	id := GetHeaderFirst(ctx, XRequestIDKey)
	if id == "" {
		id = NewRequestID()
	}
	return id
}

// GetUsername get username from header
func GetUsername(ctx context.Context) string {
	return GetHeaderFirst(ctx, UsernameKey)
}

// GetHeaderFirst get header first value
func GetHeaderFirst(ctx context.Context, key string) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if values := md.Get(key); len(values) > 0 {
			return values[0]
		}
	}
	return ""
}

// NewRequestID generate a RequestId
func NewRequestID() string {
	return uuid.New().String()
}
