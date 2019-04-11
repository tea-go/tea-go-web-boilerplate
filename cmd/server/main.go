package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tea-go/tea-go-web-boilerplate/pkg/beer"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/router"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/storage"
)

func main() {
	// create a mysql storage
	s, _ := storage.NewStorage()
	s.AutoMigrate()

	// create a new router and add default middleware
	r := router.NewRouter()
	r.Use([]gin.HandlerFunc{
		gin.Logger(), gin.Recovery(),
	})

	// handle all beer's resources
	beer.NewBeerRepository(s)

	fmt.Println("The beer server is on tap now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r.Router()))
}
