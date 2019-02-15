package handlers

import (
	"net/http"
)

type Handler struct {
	HandleFunc             func(*Context, http.ResponseWriter, *http.Request)
	IsStatic               bool
	RequiresAuthentication bool
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r)

	if h.IsStatic {
		// Instruct the browser not to display us in an iframe unless is the same origin for anti-clickjacking
		w.Header().Set("X-Frame-Options", "SAMEORIGIN")
		w.Header().Set("Content-Security-Policy", "frame-ancestors 'self'")
	} else {
		// All api response bodies will be JSON formatted by default
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodGet {
			w.Header().Set("Expires", "0")
		}
	}

	if h.RequiresAuthentication && c.Session == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	h.HandleFunc(c, w, r)
}

func ApiHandler(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &Handler{
		HandleFunc:             h,
		IsStatic:               false,
		RequiresAuthentication: false,
	}
}

func ApiAuthenticatedHandler(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &Handler{
		HandleFunc:             h,
		IsStatic:               false,
		RequiresAuthentication: true,
	}
}
