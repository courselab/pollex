package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/courselab/pollex/pollex-backend/pkg/controllers"
	"github.com/courselab/pollex/pollex-backend/pkg/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/courselab/pollex/pollex-backend/pkg/controllers/mocks"
)

func setupLocationsTest() (*gin.Engine, *handler) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	h := NewHandler(&Params{
		Router:    router,
		Locations: controllers.NewLocationsController(&controllers.LocationsParams{}),
	})

	return router, h
}

func TestGetLocations(t *testing.T) {

	router, handler := setupLocationsTest()

	t.Run("locations endpoint success", func(t *testing.T) {
		mockResponse := []domain.Location{
			{
				Id:        1,
				Name:      "MOCKLocationTest",
				Coords:    "MOCKCoordsTest",
				Thumbnail: "MOCKTHumbnailTest",
			},
		}
		mockLocationController := new(mocks.Locations)
		mockLocationController.On("GetLocations").Return(mockResponse)

		handler.locations = mockLocationController

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/locations", nil)
		router.ServeHTTP(w, req)

		var response []domain.Location
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, mockResponse, response)
	})
}
