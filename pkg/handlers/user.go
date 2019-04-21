package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/app"
	Error "github.com/tea-go/tea-go-web-boilerplate/pkg/json/error"
	Fail "github.com/tea-go/tea-go-web-boilerplate/pkg/json/fail"
	Success "github.com/tea-go/tea-go-web-boilerplate/pkg/json/success"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/models"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/router"
)

type (
	userService interface {
		Query(rs app.RequestScope, offset, limit int) ([]models.User, error)
		Count(rs app.RequestScope) (int, error)
		Get(rs app.RequestScope, id int) (*models.User, error)
		Create(rs app.RequestScope, user *models.User) (*models.User, error)
		Update(rs app.RequestScope, id int, user *models.User) (*models.User, error)
		Delete(rs app.RequestScope, id int) (*models.User, error)
	}

	userResource struct {
		service userService
	}
)

// HanldeUserResource handler all user's resource
func HanldeUserResource(r router.Routes, service userService) router.Routes {
	resource := &userResource{service}

	r.LIST("user", "", resource.query)
	r.DETAIL("user", "", resource.get)
	r.POST("user", "", resource.create)
	r.PATCH("user", "", resource.update)
	r.DELETE("user", "", resource.delete)

	return r
}

// users return all users
func (r *userResource) query(c *gin.Context) {
	rs := app.GetRequestScope(c)

	count, err := r.service.Count(rs)

	if err != nil {
		h := Error.InternalServerError(err.Error(), 0)
		h["statuCode"] = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, h)
		return
	}

	if count == 0 {
		h := Success.OK("ok", 0)
		h["statusCode"] = http.StatusOK
		h["count"] = count
		h["data"] = nil
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

	users, err := r.service.Query(rs, offset, limit)

	if err != nil {
		h := Error.InternalServerError(err.Error(), 0)
		h["statuCode"] = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, h)
		return
	}

	h := Success.OK("ok", 0)
	h["statusCode"] = http.StatusOK
	h["count"] = count
	h["data"] = users

	c.JSON(http.StatusOK, h)
}

// user return a user's detail
func (r *userResource) get(c *gin.Context) {
	rs := app.GetRequestScope(c)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		h := Fail.InvalidParameter("id is invalid", 0)
		h["statusCode"] = http.StatusUnprocessableEntity
		c.JSON(http.StatusUnprocessableEntity, h)
		return
	}

	user, err := r.service.Get(rs, id)

	if err != nil {
		h := Fail.NotFound(err.Error(), 0)
		h["statusCode"] = http.StatusNotFound
		c.JSON(http.StatusNotFound, h)
		return
	}

	h := Success.OK("ok", 0)
	h["statusCode"] = http.StatusOK
	h["data"] = user

	c.JSON(http.StatusOK, h)
}

// create create a user
func (r *userResource) create(c *gin.Context) {
	rs := app.GetRequestScope(c)

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		h := Fail.BindJSONFail(err.Error(), 0)
		h["statusCode"] = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, h)
		return
	}

	data, err := r.service.Create(rs, &user)

	if err != nil {
		h := Error.InternalServerError(err.Error(), 0)
		h["statusCode"] = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, h)
		return
	}

	h := Success.OK("ok", 0)
	h["statusCode"] = http.StatusOK
	h["data"] = data

	c.JSON(http.StatusOK, h)
}

// update update a user
func (r *userResource) update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		h := Fail.InvalidID(err.Error(), 0)
		h["statusCode"] = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, h)
		return
	}

	rs := app.GetRequestScope(c)

	if _, err := r.service.Get(rs, id); err != nil {
		h := Fail.NotFound(err.Error(), 0)
		h["statusCode"] = http.StatusNotFound
		c.JSON(http.StatusNotFound, h)
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		h := Fail.BindJSONFail(err.Error(), 0)
		h["statusCode"] = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, h)
		return
	}

	updated, err := r.service.Update(rs, id, &user)

	if err != nil {
		h := Error.InternalServerError(err.Error(), 0)
		h["statusCode"] = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, h)
		return
	}

	h := Success.OK("ok", 0)
	h["statusCode"] = http.StatusOK
	h["data"] = updated

	c.JSON(http.StatusOK, h)
}

// update update a user
func (r *userResource) delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		h := Fail.InvalidID(err.Error(), 0)
		h["statusCode"] = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, h)
		return
	}

	rs := app.GetRequestScope(c)

	if _, err := r.service.Delete(rs, id); err != nil {
		h := Error.InternalServerError(err.Error(), 0)
		h["statusCode"] = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, h)
		return
	}

	h := Success.NoContent("", 0)
	h["statusCode"] = http.StatusNoContent
	c.JSON(http.StatusNoContent, h)
}
