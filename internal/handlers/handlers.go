package handlers

import (
	"net/http"

	"github.com/callicoder/go-ready/pkg/errors"
)

type Handler struct {
	HandleFunc             func(Context)
	RequiresAuthentication bool
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r)

	if h.RequiresAuthentication && c.Session == nil {
		c.Error(errors.NewUnauthorizedError(errors.New("Sorry, You're not authorized to access this resource")))
		return
	}

	h.HandleFunc(c)
}

func ApiHandler(h func(Context)) http.Handler {
	return &Handler{
		HandleFunc:             h,
		RequiresAuthentication: false,
	}
}

func ApiAuthenticatedHandler(h func(Context)) http.Handler {
	return &Handler{
		HandleFunc:             h,
		RequiresAuthentication: true,
	}
}
