package handlers

import (
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

func (h *GroupHandler) create(c Context) {

}

func (h *GroupHandler) retrieve(c Context) {

}

func (h *GroupHandler) update(c Context) {

}
