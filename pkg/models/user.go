package models

import (
	"crypto/md5"
	"fmt"
	"io"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/util"
)

// User defines the model of a user
type User struct {
	Base
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
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

// BeforeCreate check user validation before create
func (u *User) BeforeCreate() error {
	salt := util.RandStr(10, "normal")

	h := md5.New()

	io.WriteString(h, salt)

	u.Salt = salt

	password := fmt.Sprintf("%x", h.Sum(nil))

	fmt.Println(salt, password)

	return nil
}
