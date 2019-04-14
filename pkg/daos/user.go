package daos

import (
	"time"

	"github.com/tea-go/tea-go-web-boilerplate/pkg/app"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/models"
)

// UserDAO a struct of User dao
type UserDAO struct{}

// NewUserDAO create a instance of User dao
func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

// Get get a user detail
func (dao *UserDAO) Get(rs app.RequestScope, id int) *models.User {
	var user models.User

	rs.Tx().Where("id = ?", id).First(&user)

	return &user
}

// GetByEmail get a user by email
func (dao *UserDAO) GetByEmail(rs app.RequestScope, email string) *models.User {
	var user models.User

	rs.Tx().Where("email = ?", email).First(&user)

	return &user
}

// Query get all Users from db
func (dao *UserDAO) Query(rs app.RequestScope, offset, limit int) []models.User {
	var users []models.User

	rs.Tx().Where("is_deleted = ?", "no").Find(&users).Offset(offset).Limit(limit)

	return users
}

// Count count all Users
func (dao *UserDAO) Count(rs app.RequestScope) int {
	var Users []models.User
	var count int

	rs.Tx().Where("is_deleted = ?", "no").Find(&Users).Count(&count)

	return count
}

// Create create a user
func (dao *UserDAO) Create(rs app.RequestScope, user *models.User) *models.User {
	user.ID = 0

	now := time.Now()

	user.CreatedAt = now
	user.UpdatedAt = now

	if user.IsDeleted == "" {
		user.IsDeleted = "no"
	}

	if user.Status == "" {
		user.Status = "enabled"
	}

	rs.Tx().Create(user)

	return user
}
