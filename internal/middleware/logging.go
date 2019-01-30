package middleware

import (
	"net/http"

	"github.com/callicoder/go-ready/internal/context"
	"github.com/callicoder/go-ready/pkg/logger"
)

func LoggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := context.RequestID(r.Context())
		ipAddress := context.IPAddress(r.Context())

		logger.WithFields(logger.Fields{
			"path":       r.URL.Path,
			"request_id": requestID,
			"method":     r.Method,
			"ip_address": ipAddress,
		}).Info("Http Request")

		next.ServeHTTP(w, r)
	})
}
