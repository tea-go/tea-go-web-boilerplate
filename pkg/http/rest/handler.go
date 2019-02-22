package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_errors "github.com/tea-go/tea-go-web-boilerplate/pkg/errors"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/listing"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/middleware"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/router"
)

// Handler handler all urls
func Handler(l listing.Service) http.Handler {
	// 1. create a custom router
	r := router.NewRouter()

	router := r.Router()

	// 2. add middlewares
	middleware.NewMiddleware(router)

	/* beer apis */
	r.LIST("beer", "", getBeers(l))
	r.DETAIL("beer", "", getBeer(l))

	/* beer's review apis */
	r.LIST("review", "beer", getBeerReviews(l))

	return router
}

// getBeers returns a handler for GET /beers requests
func getBeers(s listing.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		list := s.GetBeers()
		c.JSON(http.StatusOK, list)
	}
}

// getBeer returns a handler for GET /beers/:id requests
func getBeer(s listing.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println("getBeer")
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%s is not a valid beer ID, it must be a number.", c.Param("id"))})
			return
		}

		beer, err := s.GetBeer(ID)
		if err == _errors.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "The beer you requested does not exist."})
			return
		}

		c.Header("Content-Type", "application/json")

		c.JSON(http.StatusOK, beer)
	}
}

// getBeerReviews returns a handler for GET /beers/:id/reviews requests
func getBeerReviews(s listing.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%s is not a valid beer ID, it must be a number.", c.Param("id"))})
			return
		}

		reviews := s.GetBeerReviews(ID)

		c.Header("Content-Type", "application/json")

		c.JSON(http.StatusOK, reviews)
	}
}
