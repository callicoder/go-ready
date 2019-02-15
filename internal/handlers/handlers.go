package handlers

import (
	"net/http"

	"github.com/callicoder/go-ready/pkg/errors"
)

type Handler struct {
	HandleFunc             func(*Context, http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r)

	if h.RequiresAuthentication && c.Session == nil {
		c.Error(errors.NewUnauthorizedError(errors.New("Sorry, You're not authorized to access this resource")))
		return
	}

	h.HandleFunc(c, w, r)
}

func ApiHandler(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &Handler{
		HandleFunc:             h,
		RequiresAuthentication: false,
	}
}

func ApiAuthenticatedHandler(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &Handler{
		HandleFunc:             h,
		RequiresAuthentication: true,
	}
}
