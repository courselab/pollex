package main

import (
  "fmt"
  "net/http"
  "log"
  "github.com/gin-gonic/gin"
)

func main() {
  Router.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong")
  })
  fmt.Println("Starting server on http://localhost:8080")
  log.Fatal(Router.Run(":8080"))
}
