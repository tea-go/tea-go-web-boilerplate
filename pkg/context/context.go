package context

import (
	"github.com/gin-gonic/gin"
)



func NewContext() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		ctx.Set("context", )
	}
}