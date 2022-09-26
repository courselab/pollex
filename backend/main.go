package main

import (
  "fmt"
  "log"
  "github.com/gin-gonic/gin"
  "courselab/pollex/backend/handlers"
)

func main() {
  router := gin.Default()
  handlers.SetRoutes(router)
  fmt.Println("Starting server on http://localhost:8080")
  log.Fatal(router.Run(":8080"))
}
