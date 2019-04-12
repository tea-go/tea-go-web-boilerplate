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
	beerService interface {
		Query(rs app.RequestScope, offset, limit int) ([]models.Beer, error)
		Count(rs app.RequestScope) int
	}

	beerResource struct {
		service beerService
	}
)

// HanldeBeerResource handler all beer's resource
func HanldeBeerResource(r router.Routes, service beerService) router.Routes {
	resource := &beerResource{service}

	r.LIST("beer", "", resource.beers)

	return r
}

// get beers returns a handler for GET /beers requests
func (r *beerResource) beers(c *gin.Context) {
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

	beers, _ := r.service.Query(rs, offset, limit)

	c.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"statusCode": http.StatusOK,
		"count":      count,
		"message":    "ok",
		"data":       beers,
	})
}
