package handlers

import (
    "net/url"
    "os"

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
    h.authenticatedRoutes(p.Router)

	return h
}

func (h *handler) routePing(router *gin.Engine) {
	router.GET("/ping", h.ping)
}

func (h *handler) authenticatedRoutes(router *gin.Engine) {
    group := router.Group("/")

    //TODO: figure out a better place for this configuration
    base := os.Getenv("AUTH_SERVICE_URL")
    if base != "UNIT_TEST" {
        baseUrl, err := url.Parse(base)
        if len(base) == 0 || err != nil {
            panic("Invalid or missing auth service base url")
        }

        group.Use(checkAuth(baseUrl))
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
