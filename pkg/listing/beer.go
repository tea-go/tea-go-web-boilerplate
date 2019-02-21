package listing

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Beer defines the properties of a beer to be listed
type Beer struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Brewery   string    `json:"brewery"`
	Abv       float32   `json:"abv"`
	ShortDesc string    `json:"short_description"`
	Created   time.Time `json:"created"`
}

// BeerModel defines the properties of a beer model to be listed
type BeerModel struct {
	gorm.Model
	Name      string  `json:"name"`
	Brewery   string  `json:"brewery"`
	Abv       float32 `json:"abv"`
	ShortDesc string  `json:"short_description"`
}
