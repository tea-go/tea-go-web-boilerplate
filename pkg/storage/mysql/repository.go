package mysql

import (
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/adding"
)

// Storage stores beer data in JSON files
type Storage struct {
	db *gorm.DB
}

// NewStorage return a new mysql storage
func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	dbUser := "root"
	dbPass := "12345678"
	dbHost := "127.0.0.1"
	dbPort := "3306"
	dbName := "teagodev"

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	params := url.Values{}
	params.Add("parseTime", "1")
	params.Add("loc", "Asia/Shanghai")

	dsn := fmt.Sprintf("%s?%s", connection, params.Encode())

	s.db, err = gorm.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Storage) AddBeer(b adding.BeerModel) error {
	existingBeers := s.GetAllBeers()
	for _, e := range existingBeers {
		if (b.Abv == e.Abv) &&
		b.Brewery == e.Brewery &&
		b.Name == e.Name {
			return adding.ErrDuplicate
		}
	}

	newB := Beer{
		Name: b.Name,
		Brewery: b.Brewery,
		Abv: b.Abv,
		ShortDesc: b.ShortDesc,
	}

	s.db.Create(&newB)
}

func (s *Storage) GetAllBeers()[]listing.BeerModel {
	return []
}
