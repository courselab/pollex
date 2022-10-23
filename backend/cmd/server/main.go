package main

import (
	"fmt"
	"log"

	"github.com/courselab/pollex/pollex-backend/pkg/controllers"
	"github.com/courselab/pollex/pollex-backend/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	handlers.NewHandler(&handlers.Params{
		Router: router,
		User:   controllers.NewUserController(&controllers.Params{}),
	})

	fmt.Println("Starting server on http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}
