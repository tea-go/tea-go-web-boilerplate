package mysql

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Review defines the storage form of a beer review
type Review struct {
	gorm.Model
	BeerID    int       `json:"beer_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Score     int       `json:"score"`
	Text      string    `json:"text"`
	Created   time.Time `json:"created"`
}
