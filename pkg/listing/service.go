package listing

import "github.com/tea-go/tea-go-web-boilerplate/pkg/storage/mysql"

// Repository provides access to the beer and review storage.
type Repository interface {
	// GetBeer returns the beer with given ID.
	// GetBeer(int) (Beer, error)
	GetBeer(id int) (mysql.Beer, error)
	// GetAllBeers returns all beers saved in storage.
	// GetAllBeers() []Beer
	GetAllBeers() []mysql.Beer
	// GetAllReviews returns a list of all reviews for a given beer ID.
	// GetAllReviews(int) []Review
	GetAllReviews(beeID int) []mysql.Review
}

// Service provides beer and review listing operations.
type Service interface {
	GetBeer(int) (mysql.Beer, error)
	GetBeers() []mysql.Beer
	GetBeerReviews(int) []mysql.Review
}

type service struct {
	r Repository
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// GetBeers returns all beers
func (s *service) GetBeers() []mysql.Beer {
	return s.r.GetAllBeers()
}

// GetBeer returns a beer
func (s *service) GetBeer(id int) (mysql.Beer, error) {
	return s.r.GetBeer(id)
}

// GetBeerReviews returns all requests for a beer
func (s *service) GetBeerReviews(beerID int) []mysql.Review {
	return s.r.GetAllReviews(beerID)
}
