package handlers

import (
	"net/http"

	"github.com/callicoder/go-ready/internal/model"
	"github.com/callicoder/go-ready/internal/service"
	"github.com/callicoder/go-ready/pkg/contracts"
	"github.com/gorilla/mux"
)

type GroupHandler struct {
	groupService service.GroupService
}

func InitGroupHandler(router *mux.Router, groupService service.GroupService) {
	groupHandler := &GroupHandler{
		groupService: groupService,
	}

	router.Handle("/groups", ApiAuthenticatedHandler(groupHandler.create)).Methods("POST")
	router.Handle("/groups/{groupId}", ApiAuthenticatedHandler(groupHandler.retrieve)).Methods("GET")
	router.Handle("/groups/{groupId}", ApiAuthenticatedHandler(groupHandler.update)).Methods("PUT")
}

func (h *GroupHandler) create(c Context) {
	var groupRequest contracts.GroupRequest
	if err := c.BindJSON(&groupRequest); err != nil {
		c.Error(err)
		return
	}

	group := &model.Group{
		Name:        groupRequest.Name,
		Description: groupRequest.Description,
		ImageUrl:    groupRequest.ImageUrl,
		CreatedBy:   c.Session().Id,
		UpdatedBy:   c.Session().Id,
	}

	h.groupService.Save(group)
	c.JSON(http.StatusCreated, group)
}

func (h *GroupHandler) retrieve(c Context) {

}

func (h *GroupHandler) update(c Context) {

}
