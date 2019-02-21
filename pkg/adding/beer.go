package adding

import (
	"github.com/jinzhu/gorm"
)

// Beer defines the properties of a beer to be added
type Beer struct {
	Name      string  `json:"name"`
	Brewery   string  `json:"brewery"`
	Abv       float32 `json:"abv"`
	ShortDesc string  `json:"short_description"`
}

// BeerModel defines the properties of a beer model to be added
type BeerModel struct {
	gorm.Model
	Name      string  `json:"name"`
	Brewery   string  `json:"brewery"`
	Abv       float32 `json:"abv"`
	ShortDesc string  `json:"short_description"`
}
