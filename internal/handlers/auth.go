package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/callicoder/go-ready/internal/service"
	"github.com/callicoder/go-ready/pkg/contracts"
	"github.com/callicoder/go-ready/pkg/logger"
	"github.com/gorilla/mux"
	"google.golang.org/api/oauth2/v2"
)

type AuthHandler struct {
	userService *service.UserService
}

func InitAuthHandler(router *mux.Router, userService *service.UserService) {
	authHandler := &AuthHandler{
		userService: userService,
	}

	router.Handle("/auth/tokensignin", ApiHandler(authHandler.tokenSignin)).Methods("POST")
}

func (h *AuthHandler) tokenSignin(c *Context, w http.ResponseWriter, r *http.Request) {
	var tokenSigninRequest contracts.TokenSigninRequest
	if err := json.NewDecoder(r.Body).Decode(&tokenSigninRequest); err != nil {

	}

	httpClient := &http.Client{
		Timeout: 3 * time.Second,
	}

	oauth2Service, err := oauth2.New(httpClient)
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(tokenSigninRequest.IdToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {

	}

	logger.Info(tokenInfo)
}
