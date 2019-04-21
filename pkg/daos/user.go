package daos

import (
	"errors"
	"fmt"
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
func (dao *UserDAO) Get(rs app.RequestScope, id int) (*models.User, error) {
	var user models.User

	if notFound := rs.Tx().Where("id = ?", id).First(&user).RecordNotFound(); notFound == true {
		return nil, errors.New("user not found by id")
	}

	return &user, nil
}

// GetByEmail get a user by email
func (dao *UserDAO) GetByEmail(rs app.RequestScope, email string) (*models.User, error) {
	var user models.User

	if notFound := rs.Tx().Where("email = ?", email).First(&user).RecordNotFound(); notFound == true {
		return nil, errors.New("user not found by email")
	}

	return &user, nil
}

// Query get all Users from db
func (dao *UserDAO) Query(rs app.RequestScope, offset, limit int) ([]models.User, error) {
	var users []models.User

	if err := rs.Tx().Where("is_deleted = ?", "no").Find(&users).Offset(offset).Limit(limit).Error; err != nil {
		return nil, errors.New("query all users error from db")
	}

	return users, nil
}

// Count count all Users
func (dao *UserDAO) Count(rs app.RequestScope) (int, error) {
	var count int

	if err := rs.Tx().Table("user").Where("is_deleted = ?", "no").Count(&count).Error; err != nil {
		return 0, errors.New("count all users error from db")
	}

	return count, nil
}

// Create create a user
func (dao *UserDAO) Create(rs app.RequestScope, user *models.User) error {
	user.ID = 0

	now := time.Now()

	user.CreatedAt = now
	user.UpdatedAt = now
	user.IsDeleted = "no"

	if user.Status == "" {
		user.Status = "enabled"
	}

	if user.Language == "" {
		user.Language = "zh"
	}

	if err := rs.Tx().Create(user).Error; err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// Update update a user
func (dao *UserDAO) Update(rs app.RequestScope, id int, user *models.User) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}

	user.ID = id
	user.UpdatedAt = time.Now()

	if user.IsDeleted == "" {
		user.IsDeleted = "no"
	}

	if user.Status == "" {
		user.Status = "enabled"
	}

	if user.Language == "" {
		user.Status = "zh"
	}

	if err := rs.Tx().Save(&user).Error; err != nil {
		return err
	}

	return nil
}

// Delete delete a user by id
func (dao *UserDAO) Delete(rs app.RequestScope, id int) error {
	user, err := dao.Get(rs, id)

	if err != nil {
		return err
	}

	user.IsDeleted = "yes"

	if err := rs.Tx().Save(&user).Error; err != nil {
		return err
	}

	return nil
}
