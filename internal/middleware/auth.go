package middleware

import (
	"net/http"
	"strings"

	"github.com/callicoder/go-ready/internal/context"
	"github.com/callicoder/go-ready/internal/service"
	"github.com/callicoder/go-ready/pkg/logger"
)

func AuthHandler(next http.Handler, tokenService *service.TokenService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := getAuthTokenFromRequest(r)
		if token != "" {
			session, err := tokenService.GetUserSessionFromToken(token)
			logger.Info(session)
			if err != nil {
				logger.Infof("Could not get user from token %s", err.Error())
			} else {
				ctx := context.WithUserSession(r.Context(), session)
				r = r.WithContext(ctx)
			}
		}

		next.ServeHTTP(w, r)
	})
}

func getAuthTokenFromRequest(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")

	// Parse the token from the header
	if len(authHeader) > 6 && strings.ToUpper(authHeader[0:6]) == "BEARER" {
		return authHeader[7:]
	}

	return ""
}
