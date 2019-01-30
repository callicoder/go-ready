package middleware

import (
	"net/http"
	"strings"

	"github.com/callicoder/go-ready/internal/context"
	"github.com/callicoder/go-ready/pkg/requestutil"
	uuid "github.com/satori/go.uuid"
)

const (
	headerRequestID = "X-Request-ID"
)

func TracingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := strings.TrimSpace(r.Header.Get(headerRequestID))
		if len(requestID) == 0 {
			requestID = uuid.NewV4().String()
		}

		ctx := context.WithRequestID(r.Context(), requestID)
		ctx = context.WithIPAddress(r.Context(), requestutil.GetIpAddress(r))

		r = r.WithContext(ctx)
		w.Header().Set(headerRequestID, requestID)
		next.ServeHTTP(w, r)
	})
}
