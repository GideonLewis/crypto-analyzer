package ginrouter

import "github.com/gin-gonic/gin"

type MapAPI map[string]map[string]func()

type Middleware gin.HandlerFunc

type RouterConfig struct {
	APIs        MapAPI
	Middlewares []gin.HandlerFunc
}
