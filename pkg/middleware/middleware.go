package middleware

import "github.com/gin-gonic/gin"

// NewMiddleware create middlewares
func NewMiddleware(engine *gin.Engine) *gin.Engine {
	engine.Use(gin.Logger(), gin.Recovery())

	return engine
}
