package handlers

import (
	"github.com/courselab/pollex/pollex-backend/pkg/controllers"
	"github.com/gin-gonic/gin"
)

type handler struct {
	user      controllers.User
	locations controllers.Locations
}

type Params struct {
	Router    *gin.Engine
	User      controllers.User
	Locations controllers.Locations
}

func NewHandler(p *Params) *handler {
	h := &handler{
		user:      p.User,
		locations: p.Locations,
	}

	h.routePing(p.Router)
    h.routeLogin(p.Router)
    h.authenticatedRoutes(p.Router)

	return h
}

func (h *handler) routePing(router *gin.Engine) {
	router.GET("/ping", h.ping)
}

func (h *handler) routeLogin(router *gin.Engine) {
    router.GET("/login/google", h.googleLogin)
    router.POST("/login/google/callback", h.googleLoginCallback)
}

func (h *handler) authenticatedRoutes(router *gin.Engine) {
    group := router.Group("/")

    //nil only in unit tests
    if authBaseUrl != nil {
        group.Use(checkAuth(authBaseUrl))
    }

    h.routeUsers(group)
	h.routeLocations(group)
}

func (h *handler) routeUsers(router *gin.RouterGroup) {
	router.GET("/users", h.getUsers)
	router.GET("/users/:id", h.getUser)
	router.POST("/users", h.createUser)
	router.PUT("/users/:id", h.updateUser)
	router.DELETE("/users/:id", h.deleteUser)
	router.PATCH("/users/:id", h.patchUser)
}

func (h *handler) routeLocations(router *gin.RouterGroup) {
	router.GET("/locations", h.getLocations)
}
