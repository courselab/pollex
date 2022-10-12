package handlers

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {
	router.GET("/ping", ping)
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)
	router.POST("/users", createUser)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)
	router.PATCH("/users/:id", patchUser)
}
