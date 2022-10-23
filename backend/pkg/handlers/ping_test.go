package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupPingTest() (*gin.Engine, *handler) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	h := NewHandler(&Params{
		Router: router,
	})

	return router, h
}

func TestPing(t *testing.T) {

	router, _ := setupPingTest()

	t.Run("ping endpoint success", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "pong", w.Body.String())
	})
}
