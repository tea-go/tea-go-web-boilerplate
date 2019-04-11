package beer

import "github.com/tea-go/tea-go-web-boilerplate/pkg/storage"

// Repository provides access to the beer and review storage.
type Repository interface {
	GetBeers() []storage.Beer
}

// Service provides beer and review listing operations.
type Service interface {
	GetBeers() []storage.Beer
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
	return s.r.GetBeers()
}
