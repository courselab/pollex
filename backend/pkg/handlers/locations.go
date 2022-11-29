package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) getLocations(c *gin.Context) {
	locations := h.locations.GetLocations()
	c.IndentedJSON(http.StatusOK, &locations)
}
