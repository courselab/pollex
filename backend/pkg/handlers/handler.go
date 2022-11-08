package handlers

import (
	"github.com/courselab/pollex/pollex-backend/pkg/controllers"
	"github.com/gin-gonic/gin"
)

type handler struct {
	user   controllers.User
	travel controllers.Travel
}

type Params struct {
	Router *gin.Engine
	User   controllers.User
	Travel controllers.Travel
}

func NewHandler(p *Params) *handler {
	h := &handler{
		user:   p.User,
		travel: p.Travel,
	}

	h.routePing(p.Router)
	h.routeUsers(p.Router)
	h.routeTravels(p.Router)

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

func (h *handler) routeTravels(router *gin.Engine) {
	router.GET("/travels", h.getTravels)
	router.GET("/travels/:id", h.getTravel)
	router.POST("/travels", h.createTravel)
	router.PUT("/travels/:id", h.updateTravel)
	router.DELETE("/travels/:id", h.deleteTravel)
	router.PATCH("/travels/:id", h.patchTravel)
}
