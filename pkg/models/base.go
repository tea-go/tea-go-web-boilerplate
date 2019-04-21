package models

import "time"

// Base the base of model
type Base struct {
	ID        int       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsDeleted string    `json:"isDeleted"`
}
