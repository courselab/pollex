package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/courselab/pollex/pollex-backend/pkg/controllers"
	"github.com/courselab/pollex/pollex-backend/pkg/domain"
	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {
	users := controllers.GetUsers()

	c.IndentedJSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	param := c.Param("id")
	userId, err := paramToInt(param)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := controllers.GetUser(int32(*userId))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func createUser(c *gin.Context) {

	/*
		Input example:

		{
			"id": 1223,
			"name": "Igor Takeo Passenger",
			"nickname": "igortakeo_passenger",
			"isDriver": false,
			"driverStats": null,
			"passengerStats": {
				"ratingAvg": 10,
				"ratingCount": 10,
				"tripCount": 50
			},
			"car": null
		}

	*/

	user := domain.User{}
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := validateDriverConditions(user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userCreated, err := controllers.CreateUser(user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, userCreated)
}

func updateUser(c *gin.Context) {
	param := c.Param("id")
	userId, err := paramToInt(param)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user := domain.User{}
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := validateDriverConditions(user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userUpdated, err := controllers.UpdateUser(int32(*userId), user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, userUpdated)
}

func deleteUser(c *gin.Context) {
	param := c.Param("id")
	userId, err := paramToInt(param)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err = controllers.DeleteUser(int32(*userId)); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.String(http.StatusOK, "")
}

func patchUser(c *gin.Context) {
	param := c.Param("id")
	userId, err := paramToInt(param)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user := domain.User{}
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := validateDriverConditions(user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userUpdated, err := controllers.PatchUser(int32(*userId), user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, userUpdated)
}

func paramToInt(param string) (*int, error) {
	paramToInt, err := strconv.Atoi(param)
	if err != nil {
		return nil, err
	}

	return &paramToInt, nil
}

func validateDriverConditions(user domain.User) error {
	if user.IsDriver && (user.DriverStats == nil || user.Car == nil) {
		var err error

		if user.DriverStats == nil {
			err = errors.New("driverStats is required")
		} else if user.Car == nil {
			err = errors.New("car is required")
		}

		return err
	}

	return nil
}
