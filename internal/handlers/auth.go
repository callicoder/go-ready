package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/callicoder/go-ready/internal/config"
	"github.com/callicoder/go-ready/internal/model"
	"github.com/callicoder/go-ready/internal/service"
	"github.com/callicoder/go-ready/pkg/contracts"
	googleAuthIDTokenVerifier "github.com/futurenda/google-auth-id-token-verifier"
	"github.com/gorilla/mux"
)

type AuthHandler struct {
	userService service.UserService
	config      config.AuthConfig
}

func InitAuthHandler(router *mux.Router, userService service.UserService, config config.AuthConfig) {
	authHandler := &AuthHandler{
		userService: userService,
		config:      config,
	}

	router.Handle("/auth/tokensignin", ApiHandler(authHandler.tokenSignin)).Methods("POST")
}

func (h *AuthHandler) tokenSignin(c *Context, w http.ResponseWriter, r *http.Request) {
	var tokenSigninRequest contracts.TokenSigninRequest
	if err := json.NewDecoder(r.Body).Decode(&tokenSigninRequest); err != nil {
		return
	}

	verifier := googleAuthIDTokenVerifier.Verifier{}
	err := verifier.VerifyIDToken(tokenSigninRequest.IdToken, []string{
		h.config.GoogleClientId,
	})
	if err != nil {
		return
	}
	claimSet, err := googleAuthIDTokenVerifier.Decode(tokenSigninRequest.IdToken)
	if err != nil {
		return
	}

	user := &model.User{
		Name:     claimSet.Name,
		Email:    claimSet.Email,
		ImageUrl: claimSet.Picture,
	}

	h.userService.Create(user)
	return
}
