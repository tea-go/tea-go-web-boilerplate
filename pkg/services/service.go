package services

import (
	"github.com/tea-go/tea-go-web-boilerplate/pkg/app"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/models"
)

type beerDao interface {
	Query(rs app.RequestScope, offset, limit int) ([]models.Beer, error)
	Count(rs app.RequestScope) int
}

// BeerService a struct of beer service
type BeerService struct {
	dao beerDao
}

// NewBeerService create a instance of beer service
func NewBeerService(dao beerDao) *BeerService {
	return &BeerService{dao}
}

// Query get all beers from dao
func (bs *BeerService) Query(rs app.RequestScope, offset, limit int) ([]models.Beer, error) {
	return bs.dao.Query(rs, offset, limit)
}

// Count count all beers from
func (bs *BeerService) Count(rs app.RequestScope) int {
	return bs.dao.Count(rs)
}
