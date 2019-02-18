package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/listing"
)

// Handler handler all urls
func Handler(l listing.Service) http.Handler {
	router := httprouter.New()

	router.GET("/beers", getBeers(l))
	router.GET("/beers/:id", getBeer(l))
	router.GET("/beers/:id/reviews", getBeerReviews(l))

	return router
}

// getBeers returns a handler for GET /beers requests
func getBeers(s listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		list := s.GetBeers()
		json.NewEncoder(w).Encode(list)
	}
}

// getBeer returns a handler for GET /beers/:id requests
func getBeer(s listing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ID, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid beer ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}

		beer, err := s.GetBeer(ID)
		if err == listing.ErrNotFound {
			http.Error(w, "The beer you requested does not exist.", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(beer)
	}
}

// getBeerReviews returns a handler for GET /beers/:id/reviews requests
func getBeerReviews(s listing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ID, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid beer ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}

		reviews := s.GetBeerReviews(ID)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reviews)
	}
}
