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

	p.Router.GET("/ping", h.ping)
	p.Router.GET("/users", h.getUsers)
	p.Router.GET("/users/:id", h.getUser)
	p.Router.POST("/users", h.createUser)
	p.Router.PUT("/users/:id", h.updateUser)
	p.Router.DELETE("/users/:id", h.deleteUser)
	p.Router.PATCH("/users/:id", h.patchUser)

	return h
}
