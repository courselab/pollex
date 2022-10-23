package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/courselab/pollex/pollex-backend/pkg/controllers"
	"github.com/courselab/pollex/pollex-backend/pkg/controllers/mocks"
	"github.com/courselab/pollex/pollex-backend/pkg/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupUserTest() (*gin.Engine, *handler) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	h := NewHandler(&Params{
		Router: router,
		User:   controllers.NewUserController(&controllers.Params{}),
	})

	return router, h
}

func TestGetUsers(t *testing.T) {

	router, handler := setupUserTest()

	t.Run("Success", func(t *testing.T) {
		mockResponse := []domain.User{
			{
				Id:          12,
				Name:        "Igor Takeo Passenger",
				Nickname:    "igortakeo_passenger",
				IsDriver:    false,
				DriverStats: nil,
				PassengerStats: domain.PassengerStats{
					RatingAvg:   10,
					RatingCount: 10,
					TripCount:   50,
				},
				Car: nil,
			},
		}
		mockUserController := new(mocks.User)
		mockUserController.On("GetUsers").Return(mockResponse)

		handler.user = mockUserController

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
		router.ServeHTTP(w, req)

		var response []domain.User
		json.Unmarshal(w.Body.Bytes(), &response)

		fmt.Println(mockResponse)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, mockResponse, response)
	})
}

func TestGetUser(t *testing.T) {

	router, handler := setupUserTest()

	t.Run("Success", func(t *testing.T) {
		mockResponse := domain.User{
			Id:          10,
			Name:        "",
			Nickname:    "",
			IsDriver:    false,
			DriverStats: nil,
			PassengerStats: domain.PassengerStats{
				RatingAvg:   0,
				RatingCount: 0,
				TripCount:   0,
			},
			Car: nil,
		}
		mockUserController := new(mocks.User)
		mockUserController.On("GetUser", int32(10)).Return(mockResponse, nil)

		handler.user = mockUserController

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/users/10", nil)
		router.ServeHTTP(w, req)

		var response domain.User
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, mockResponse, response)
	})

	t.Run("Param invalid", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/users/fail", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestCreateUser(t *testing.T) {
	router, handler := setupUserTest()

	t.Run("Success", func(t *testing.T) {
		mockResponse := domain.User{
			Id:          10,
			Name:        "Igor",
			Nickname:    "igorteste",
			IsDriver:    false,
			DriverStats: nil,
			PassengerStats: domain.PassengerStats{
				RatingAvg:   10,
				RatingCount: 10,
				TripCount:   10,
			},
			Car: nil,
		}
		mockUserController := new(mocks.User)
		mockUserController.On("CreateUser", mockResponse).Return(&mockResponse, nil)

		handler.user = mockUserController

		w := httptest.NewRecorder()
		jsonObject, _ := json.Marshal(mockResponse)
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonObject))
		router.ServeHTTP(w, req)

		var response domain.User
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, mockResponse, response)
	})

	t.Run("Invalid input", func(t *testing.T) {
		requestBody := `{"invalid"}`

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer([]byte(requestBody)))
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Driver conditions not satisfied", func(t *testing.T) {
		requestBody := domain.User{
			Id:          10,
			Name:        "Igor",
			Nickname:    "igorteste",
			IsDriver:    true,
			DriverStats: nil,
			PassengerStats: domain.PassengerStats{
				RatingAvg:   10,
				RatingCount: 10,
				TripCount:   10,
			},
			Car: nil,
		}

		w := httptest.NewRecorder()
		jsonObject, _ := json.Marshal(requestBody)
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonObject))
		router.ServeHTTP(w, req)

		var response domain.User
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Create user error", func(t *testing.T) {
		mockInput := domain.User{
			Id:          10,
			Name:        "Igor",
			Nickname:    "igorteste",
			IsDriver:    false,
			DriverStats: nil,
			PassengerStats: domain.PassengerStats{
				RatingAvg:   10,
				RatingCount: 10,
				TripCount:   10,
			},
			Car: nil,
		}
		mockUserController := new(mocks.User)
		mockUserController.On("CreateUser", mock.Anything).Return(nil, errors.New("create user error"))
		handler.user = mockUserController

		w := httptest.NewRecorder()
		jsonObject, _ := json.Marshal(mockInput)
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonObject))
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
