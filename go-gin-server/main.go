package main

import (
	"github.com/gin-gonic/gin"
	"github.com/povia/go-study/go-gin-server/controllers"
	"net/http"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	order := new(controllers.OrderController)
	v1.POST("/place-order", order.Place)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
