package storage

import (
	"fmt"
	"net/url"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_errors "github.com/tea-go/tea-go-web-boilerplate/pkg/errors"
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

// AutoMigrate automatically migrate schemas
func (s *Storage) AutoMigrate() *Storage {
	s.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Beer{}, &Review{})
	return s
}

// AddBeer add a beer to mysql
func (s *Storage) AddBeer(b Beer) (*Beer, error) {
	existingBeers := s.GetAllBeers()

	for _, e := range existingBeers {
		if (b.Abv == e.Abv) &&
			b.Brewery == e.Brewery &&
			b.Name == e.Name {
			return nil, _errors.ErrDuplicate
		}
	}

	newB := Beer{
		Name:      b.Name,
		Brewery:   b.Brewery,
		Abv:       b.Abv,
		ShortDesc: b.ShortDesc,
	}

	s.db.Create(&newB)

	return &newB, nil
}

// AddReview add bree's review
func (s *Storage) AddReview(r Review) (*Review, error) {
	var beer Beer

	s.db.First(&beer, r.BeerID)

	if &beer == nil {
		return nil, _errors.ErrNotFound
	}

	newR := Review{
		BeerID:    r.BeerID,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Score:     r.Score,
		Text:      r.Text,
	}

	s.db.Create(&newR)

	return &newR, nil
}

// GetBeer get a user by id
func (s *Storage) GetBeer(id int) (Beer, error) {
	var beer Beer

	var resource = strconv.Itoa(id)

	s.db.First(&beer, resource)

	return beer, nil
}

// GetAllBeers get all beers from mysql
func (s *Storage) GetAllBeers() []Beer {
	var beers []Beer

	s.db.Find(&beers)

	return beers
}

// GetAllReviews get all special id beer's reviews
func (s *Storage) GetAllReviews(beeID int) []Review {
	var reviews []Review

	s.db.Where("beerID = ?", beeID).Find(&reviews)

	return reviews
}
