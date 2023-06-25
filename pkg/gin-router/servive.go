package ginrouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouterService(rc RouterConfig) *gin.Engine {
	r := gin.New()
	if len(rc.Middlewares) != 0 {
		for _, middleware := range rc.Middlewares {
			r.Use(middleware)
		}
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// for _, api := range rc.APIs {

	// }
	return r
}
