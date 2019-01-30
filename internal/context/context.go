package context

import (
	"context"
	"net"

	"github.com/callicoder/go-ready/internal/model"
)

type contextKey string

const (
	requestIDKey = contextKey("request-id")
	userIPKey    = contextKey("ip-address")
	sessionKey   = contextKey("session")
)

func (ctxKey contextKey) String() string {
	return string(ctxKey)
}

// Extract RequestID from context
func RequestID(ctx context.Context) string {
	h, _ := ctx.Value(requestIDKey).(string)
	return h
}

// Associate RequestID with context
func WithRequestID(ctx context.Context, reqID string) context.Context {
	return context.WithValue(ctx, requestIDKey, reqID)
}

// Extract IPAddress from context
func IPAddress(ctx context.Context) net.IP {
	userIP, _ := ctx.Value(userIPKey).(net.IP)
	return userIP
}

// Associate IPAddress with Context
func WithIPAddress(ctx context.Context, userIP net.IP) context.Context {
	return context.WithValue(ctx, userIPKey, userIP)
}

// Get User Session from context
func Session(ctx context.Context) *model.Session {
	session, _ := ctx.Value(sessionKey).(*model.Session)
	return session
}

// Associate User Session with Context
func WithUserSession(ctx context.Context, session *model.Session) context.Context {
	return context.WithValue(ctx, sessionKey, session)
}
