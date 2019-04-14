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
		Create(rs app.RequestScope, user *models.User) (*models.User, error)
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
	r.POST("user", "", resource.create)

	return r
}

// users return all users
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

// user return a user's detail
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
}

// create create a user
func (r *userResource) create(c *gin.Context) {
	rs := app.GetRequestScope(c)

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":     "failed",
			"statusCode": http.StatusBadRequest,
			"message":    err.Error(),
			"data":       nil,
		})
		return
	}

	data, err := r.service.Create(rs, &user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":     "failed",
			"statusCode": http.StatusBadRequest,
			"message":    err.Error(),
			"data":       nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"statusCode": http.StatusOK,
		"message":    "ok",
		"data":       data,
	})
}
