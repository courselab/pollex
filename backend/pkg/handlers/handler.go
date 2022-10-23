package handlers

import (
	"github.com/courselab/pollex/pollex-backend/pkg/controllers"
	"github.com/gin-gonic/gin"
)

type handler struct {
	user controllers.User
}

type Params struct {
	Router *gin.Engine
	User   controllers.User
}

func NewHandler(p *Params) *handler {
	h := &handler{
		user: p.User,
	}

	h.routePing(p.Router)
	h.routeUsers(p.Router)

	return h
}

func (h *handler) routePing(router *gin.Engine) {
	router.GET("/ping", h.ping)
}

func (h *handler) routeUsers(router *gin.Engine) {
	router.GET("/users", h.getUsers)
	router.GET("/users/:id", h.getUser)
	router.POST("/users", h.createUser)
	router.PUT("/users/:id", h.updateUser)
	router.DELETE("/users/:id", h.deleteUser)
	router.PATCH("/users/:id", h.patchUser)
}
