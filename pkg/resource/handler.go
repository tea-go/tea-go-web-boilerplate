package resource

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/beer"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/router"
)

// Handler handler all beer's resource
func HandleBeer(r router.Routes, bs beer.Service) router.Routes {
	/* beer apis */
	r.LIST("beer", "", getBeers(bs))

	return r
}

// getBeers returns a handler for GET /beers requests
func getBeers(s beer.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		list := s.GetBeers()
		c.JSON(http.StatusOK, list)
	}
}
