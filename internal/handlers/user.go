package handlers

import (
	"net/http"

	"github.com/callicoder/go-ready/internal/service"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	userService service.UserService
}

func InitUserHandler(router *mux.Router, userService service.UserService) {
	userHandler := &UserHandler{
		userService: userService,
	}

	router.Handle("/users/{userId}", ApiAuthenticatedHandler(userHandler.find)).Methods("GET")
	router.Handle("/users/{userId}", ApiAuthenticatedHandler(userHandler.update)).Methods("PUT")
}

func (h *UserHandler) find(c *Context, w http.ResponseWriter, r *http.Request) {
	
}

func (h *UserHandler) update(c *Context, w http.ResponseWriter, r *http.Request) {

}
