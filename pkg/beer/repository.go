package beer

import (
	"github.com/jinzhu/gorm"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/storage"
)

type BeerRepository interface {
	GetBeers() []storage.Beer
}

type bs struct {
	db *gorm.DB
}

func NewBeerRepository(db *gorm.DB) BeerRepository {
	return &bs{
		db: db,
	}
}

func (bs *bs) GetBeers() []storage.Beer {
	return nil
}
