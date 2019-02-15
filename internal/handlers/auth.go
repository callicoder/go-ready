package handlers

import (
	"net/http"

	"github.com/callicoder/go-ready/internal/config"
	"github.com/callicoder/go-ready/internal/model"
	"github.com/callicoder/go-ready/internal/service"
	"github.com/callicoder/go-ready/pkg/contracts"
	googleAuthIDTokenVerifier "github.com/futurenda/google-auth-id-token-verifier"
	"github.com/gorilla/mux"
)

type AuthHandler struct {
	userService  service.UserService
	tokenService service.TokenService
	config       config.AuthConfig
}

func InitAuthHandler(router *mux.Router, userService service.UserService, tokenService service.TokenService, config config.AuthConfig) {
	authHandler := &AuthHandler{
		userService: userService,
		config:      config,
	}

	router.Handle("/auth/tokensignin", ApiHandler(authHandler.tokenSignin)).Methods("POST")
}

func (h *AuthHandler) tokenSignin(c *Context, w http.ResponseWriter, r *http.Request) {
	var tokenSigninRequest contracts.TokenSigninRequest
	if err := c.BindJSON(&tokenSigninRequest); err != nil {
		c.Error(err)
		return
	}

	verifier := googleAuthIDTokenVerifier.Verifier{}
	err := verifier.VerifyIDToken(tokenSigninRequest.IdToken, []string{
		h.config.GoogleClientId,
	})
	if err != nil {
		c.Error(err)
		return
	}

	claimSet, err := googleAuthIDTokenVerifier.Decode(tokenSigninRequest.IdToken)
	if err != nil {
		c.Error(err)
		return
	}

	user := &model.User{
		Name:     claimSet.Name,
		Email:    claimSet.Email,
		ImageUrl: claimSet.Picture,
	}

	h.userService.Save(user)

	authToken, err := h.tokenService.CreateToken(user)

	if err != nil {
		c.Error(err)
		return
	}

	authResponse := contracts.AuthResponse{
		AuthToken: authToken,
		TokenType: "Bearer",
	}
	c.JSON(http.StatusOK, authResponse)
}
