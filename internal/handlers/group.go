package handlers

import (
	"net/http"

	"github.com/callicoder/go-ready/internal/service"
	"github.com/gorilla/mux"
)

type GroupHandler struct {
	groupService service.GroupService
}

func InitGroupHandler(router *mux.Router, groupService service.GroupService) {
	groupHandler := &GroupHandler{
		groupService: groupService,
	}

	router.Handle("/groups", ApiHandler(groupHandler.create)).Methods("POST")
	router.Handle("/groups/{groupId}", ApiHandler(groupHandler.retrieve)).Methods("GET")
	router.Handle("/groups/{groupId}", ApiHandler(groupHandler.update)).Methods("PUT")
}

func (h *GroupHandler) create(c *Context, w http.ResponseWriter, r *http.Request) {

}

func (h *GroupHandler) retrieve(c *Context, w http.ResponseWriter, r *http.Request) {

}

func (h *GroupHandler) update(c *Context, w http.ResponseWriter, r *http.Request) {

}
