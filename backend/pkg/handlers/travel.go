package handlerss

import (
	"net/http"

	"github.com/courselab/pollex/pollex-backend/pkg/domain"
	"github.com/gin-gonic/gin"
)

// Return all travels
func (h *handler) getTravels(c *gin.Context) {
	travels := h.travel.getTravels()
	c.IndentedJSON(http.StatusOK, travels)
}

func (h *handler) getTravel(c *gin.Context) {
	param := c.Param("id")
	travelId, err := h.paramToInt(param)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	travel, err := h.user.GetTravel(int32(*travelId))
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.IndentedJSON(http.StatusOK, &travel)
}

func (h *handler) createTravel(c *gin.Context) {
	travel := domain.Travel{}
	if err := c.BindJSON(&travel); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	travelCreated, err := h.travel.CreateTravel(travel)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, &travelCreated)
}

func (h *handler) updateTravel(c *gin.Conext) {
	param := c.Param("id")
	travelId, err := h.paramToInt(param)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	travel := domain.Travel{}
	if err := c.BindJSON(&travel); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	travelUpdated, err := h.user.UpdateTravel(int(*travel), travel)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, &travelUpdated)
}

func (h *handler) deleteTravel(c *gin.Context) {
	param := c.Param("id")
	travelId, err := h.paramToInt(param)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err = h.user.DeleteTravel(int32(*userId)); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.String(http.StatusOK, "")
}

func (h *handler) patchTravel(c *gin.Context) {
	param := c.Param("id")
	travelId, err := h.paramToInt(param)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	travel := domain.Travel{}
	if err := c.BindJSON(&travel); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	travelUpdated, err := h.travel.PatchTravel(int32(*travelId), travel)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, &travelUpdated)
}
