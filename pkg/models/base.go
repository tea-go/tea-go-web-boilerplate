package models

import "time"

// Base the base of model
type Base struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsDelete  string    `json:"isDeleted"`
}
