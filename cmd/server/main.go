package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/app"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/daos"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/handlers"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/models"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/router"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/services"
)

func main() {

	// load application configurations
	if err := app.LoadConfig("./pkg/config"); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	gin.SetMode(app.Config.MODE)

	// create a mysql storage
	s, _ := models.NewDB()
	s.AutoMigrate()

	// create the logger
	logger := logrus.New()

	// create a new router and add default middleware
	r := router.NewRouter()
	r.Use([]gin.HandlerFunc{
		app.Init(logger), app.Transaction(s.DB()),
	})

	// handle all beer's resources
	beerDao := daos.NewBeerDao()
	beerService := services.NewBeerService(beerDao)
	handlers.HanldeBeerResource(r, beerService)

	port := fmt.Sprintf(":%d", app.Config.ServerPort)

	fmt.Println("The beer server is on tap now: http://localhost", port)

	log.Fatal(http.ListenAndServe(port, r.Router()))
}
