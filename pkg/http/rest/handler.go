package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/listing"
)

// Handler handler all urls
func Handler(l listing.Service) http.Handler {
	router := gin.Default()

	router.GET("/beers", getBeers(l))
	router.GET("/beers/:id", getBeer(l))
	router.GET("/beers/:id/reviews", getBeerReviews(l))

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
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%s is not a valid beer ID, it must be a number.", c.Param("id"))})
			return
		}

		beer, err := s.GetBeer(ID)
		if err == listing.ErrNotFound {
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
