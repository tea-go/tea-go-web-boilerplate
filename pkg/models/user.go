package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// User defines the model of a user
type User struct {
	Base
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Status   string `json:"status"`
	Language string `json:"language"`
}

// TableName reset a name of user
func (User) TableName() string {
	return "user"
}

// Validate validate user
func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Role, validation.Required, validation.In("admin", "member")),
	)
}
