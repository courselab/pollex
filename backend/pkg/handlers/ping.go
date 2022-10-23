package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
