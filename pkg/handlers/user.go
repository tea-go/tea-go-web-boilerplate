package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/app"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/models"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/router"
)

type (
	userService interface {
		Query(rs app.RequestScope, offset, limit int) []models.User
		Count(rs app.RequestScope) int
		Get(rs app.RequestScope, id int) *models.User
	}

	userResource struct {
		service userService
	}
)

// HanldeUserResource handler all user's resource
func HanldeUserResource(r router.Routes, service userService) router.Routes {
	resource := &userResource{service}

	r.LIST("user", "", resource.users)
	r.DETAIL("user", "", resource.user)

	return r
}

// get users returns a handler for GET /users requests
func (r *userResource) users(c *gin.Context) {
	rs := app.GetRequestScope(c)

	count := r.service.Count(rs)

	if count == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":     "success",
			"statusCode": http.StatusOK,
			"count":      count,
			"message":    "ok",
			"data":       nil,
		})
		return
	}

	c.DefaultQuery("offset", "0")
	c.DefaultQuery("limit", "10")

	offset, err := strconv.Atoi(c.Query("offset"))

	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.Query("limit"))

	if err != nil {
		limit = 10
	}

	users := r.service.Query(rs, offset, limit)

	c.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"statusCode": http.StatusOK,
		"count":      count,
		"message":    "ok",
		"data":       users,
	})
}

// get users returns a handler for GET /users requests
func (r *userResource) user(c *gin.Context) {
	rs := app.GetRequestScope(c)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":     "failed",
			"statusCode": http.StatusUnprocessableEntity,
			"message":    "id is a invalid id",
			"data":       nil,
		})
		return
	}

	user := r.service.Get(rs, id)

	c.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"statusCode": http.StatusOK,
		"message":    "ok",
		"data":       user,
	})
	return

}
