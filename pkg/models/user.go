package models

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
