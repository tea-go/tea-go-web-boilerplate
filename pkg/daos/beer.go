package daos

import (
	"github.com/tea-go/tea-go-web-boilerplate/pkg/app"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/models"
)

// BeerDao a struct of beer dao
type BeerDao struct{}

// NewBeerDao create a instance of beer dao
func NewBeerDao() *BeerDao {
	return &BeerDao{}
}

// Query get all beers from db
func (dao *BeerDao) Query(rs app.RequestScope, offset, limit int) ([]models.Beer, error) {
	var beers []models.Beer

	rs.Tx().Where("isDelete = ?", "no").Find(&beers).Offset(offset).Limit(limit)

	return beers, nil
}

// Count count all beers
func (dao *BeerDao) Count(rs app.RequestScope) int {
	var beers []models.Beer
	var count int

	rs.Tx().Where("isDelete = ?", "no").Find(&beers).Count(&count)

	return count
}
