package app

import (
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// Init set context in gin.Context
func Init(logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		now := time.Now()

		ac := newRequestScope(now, logger, ctx)
		ctx.Set("Context", ac)

		elapsed := float64(time.Now().Sub(now).Nanoseconds()) / 1e6
		requestLine := fmt.Sprintf("%s %s %s", ctx.Request.Method, ctx.Request.URL.Path, ctx.Request.Proto)

		ac.Infof("[%.3fms] %s", elapsed, requestLine)

		ctx.Next()
	}
}

// GetRequestScope returns the RequestScope of the current request.
func GetRequestScope(c *gin.Context) RequestScope {
	value, _ := c.Get("Context")
	return value.(RequestScope)
}
