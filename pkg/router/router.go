package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Routes defines all methods
type Routes interface {
	LIST(res string, controller string, handlers ...gin.HandlerFunc) gin.IRoutes
	DETAIL(res string, _path string, handlers ...gin.HandlerFunc) gin.IRoutes
	POST(res string, _path string, handlers ...gin.HandlerFunc) gin.IRoutes
	PUT(res string, _path string, handlers ...gin.HandlerFunc) gin.IRoutes
	PATCH(res string, _path string, handlers ...gin.HandlerFunc) gin.IRoutes
	DELETE(res string, _path string, handlers ...gin.HandlerFunc) gin.IRoutes
	Use(middlewares []gin.HandlerFunc) *gin.Engine
	Router() *gin.Engine
}

// router interval router struct
type router struct {
	engine *gin.Engine
}

// NewRouter create a new router
func NewRouter() Routes {
	return &router{
		engine: gin.New(),
	}
}

// Router get a gin.engine pointer
func (r *router) Router() *gin.Engine {
	return r.engine
}

func (r *router) Use(middlewares []gin.HandlerFunc) *gin.Engine {
	if middlewares != nil && len(middlewares) > 0 {
		r.engine.Use(middlewares...)
	}
	return r.engine
}

// LIST define a list method
func (r *router) LIST(res string, controller string, handlers ...gin.HandlerFunc) gin.IRoutes {
	relativePath := fmt.Sprintf("/%s", res+"s")

	if controller != "" {
		relativePath = fmt.Sprintf("/%s/:%s/%s", controller+"s", "id", res+"s")
	}

	return r.engine.GET(relativePath, handlers...)
}

// DETAIL define a detail method
func (r *router) DETAIL(res string, _path string, handlers ...gin.HandlerFunc) gin.IRoutes {
	relativePath := fmt.Sprintf("/%s/:%s", res+"s", "id")

	if _path != "" {
		relativePath = _path
	}

	return r.engine.GET(relativePath, handlers...)
}

// POST define a post method
func (r *router) POST(res string, _path string, handlers ...gin.HandlerFunc) gin.IRoutes {
	relativePath := fmt.Sprintf("/%s", res+"s")

	if _path != "" {
		relativePath = _path
	}

	return r.engine.POST(relativePath, handlers...)
}

// PUT define a put method
func (r *router) PUT(res string, _path string, handlers ...gin.HandlerFunc) gin.IRoutes {
	relativePath := fmt.Sprintf("/%s/:%s", res+"s", "id")

	if _path != "" {
		relativePath = _path
	}

	return r.engine.PUT(relativePath, handlers...)
}

// PATCH define a patch method
func (r *router) PATCH(res string, _path string, handlers ...gin.HandlerFunc) gin.IRoutes {
	relativePath := fmt.Sprintf("/%s/:%s", res+"s", "id")

	if _path != "" {
		relativePath = _path
	}
	return r.engine.PATCH(relativePath, handlers...)
}

// DELETE define a delete method
func (r *router) DELETE(res string, _path string, handlers ...gin.HandlerFunc) gin.IRoutes {
	relativePath := fmt.Sprintf("/%s/:%s", res+"s", "id")

	if _path != "" {
		relativePath = _path
	}
	return r.engine.DELETE(relativePath, handlers...)
}
