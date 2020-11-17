package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setUpRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello cicd")
	})
	return r
}
func main() {
	router := setUpRouter()
	router.Run(":8080")
}
