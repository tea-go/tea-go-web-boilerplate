package reviewing

import "github.com/jinzhu/gorm"

// Review defines a beer review
type Review struct {
	BeerID    int    `json:"beer_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Score     int    `json:"score"`
	Text      string `json:"text"`
}

// Review defines a beer review
type ReviewModel struct {
	gorm.Model
	BeerID    int    `json:"beer_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Score     int    `json:"score"`
	Text      string `json:"text"`
}
