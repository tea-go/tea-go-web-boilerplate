package listing

import "github.com/tea-go/tea-go-web-boilerplate/pkg/storage"

// Repository provides access to the beer and review storage.
type Repository interface {
	// GetBeer returns the beer with given ID.
	// GetBeer(int) (Beer, error)
	GetBeer(id int) (storage.Beer, error)
	// GetAllBeers returns all beers saved in storage.
	// GetAllBeers() []Beer
	GetAllBeers() []storage.Beer
	// GetAllReviews returns a list of all reviews for a given beer ID.
	// GetAllReviews(int) []Review
	GetAllReviews(beeID int) []storage.Review
}

// Service provides beer and review listing operations.
type Service interface {
	GetBeer(int) (storage.Beer, error)
	GetBeers() []storage.Beer
	GetBeerReviews(int) []storage.Review
}

type service struct {
	r Repository
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// GetBeers returns all beers
func (s *service) GetBeers() []storage.Beer {
	return s.r.GetAllBeers()
}

// GetBeer returns a beer
func (s *service) GetBeer(id int) (storage.Beer, error) {
	return s.r.GetBeer(id)
}

// GetBeerReviews returns all requests for a beer
func (s *service) GetBeerReviews(beerID int) []storage.Review {
	return s.r.GetAllReviews(beerID)
}
