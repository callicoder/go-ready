package app

import (
	"net/http"

	"github.com/callicoder/go-ready/internal/config"
	"github.com/callicoder/go-ready/internal/handlers"
	"github.com/callicoder/go-ready/internal/middleware"
	gHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Routes struct {
	Root  *mux.Router // ''
	User  *mux.Router // '/users'
	Group *mux.Router // '/groups'
}

func NewRouter(config *config.Config, deps *Dependencies) http.Handler {
	// Create root router
	router := mux.NewRouter().PathPrefix(config.Server.ContextPath).Subrouter()

	// Attach API Handlers
	handlers.InitUserHandler(router, deps.UserService)
	handlers.InitGroupHandler(router, deps.GroupService)

	// Prepare the Http handler
	var handler http.Handler = router

	// Request Logging
	handler = middleware.LoggingHandler(handler)

	// Request Tracing
	handler = middleware.TracingHandler(handler)

	// Auth
	handler = middleware.AuthHandler(handler, deps.TokenService)

	// CORS support
	allowedMethods := gHandlers.AllowedMethods([]string{"HEAD", "OPTIONS", "GET", "POST", "PUT", "DELETE"})
	allowedHeaders := gHandlers.AllowedHeaders([]string{"*"})
	allowedOrigins := gHandlers.AllowedOrigins([]string{"*"})
	maxAge := gHandlers.MaxAge(86400)
	handler = gHandlers.CORS(allowedMethods, allowedHeaders, allowedOrigins, maxAge)(handler)

	// Error Recovery
	handler = gHandlers.RecoveryHandler(gHandlers.PrintRecoveryStack(true))(handler)

	// Compression
	handler = gHandlers.CompressHandler(handler)

	return handler
}
