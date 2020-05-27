package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	c := gin.Default()
	c.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	c.Run()
}
