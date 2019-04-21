package services

import (
	"github.com/tea-go/tea-go-web-boilerplate/pkg/app"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/models"
)

type userDAO interface {
	Query(rs app.RequestScope, offset, limit int) ([]models.User, error)
	Count(rs app.RequestScope) (int, error)
	Get(rs app.RequestScope, id int) (*models.User, error)
	GetByEmail(rs app.RequestScope, email string) (*models.User, error)
	Create(rs app.RequestScope, user *models.User) error
	Update(rs app.RequestScope, id int, user *models.User) error
	Delete(rs app.RequestScope, id int) error
}

// UserService a struct of user service
type UserService struct {
	dao userDAO
}

// NewUserService create a instance of user service
func NewUserService(dao userDAO) *UserService {
	return &UserService{dao}
}

// Query get all users from dao
func (bs *UserService) Query(rs app.RequestScope, offset, limit int) ([]models.User, error) {
	return bs.dao.Query(rs, offset, limit)
}

// Count get the count of user
func (bs *UserService) Count(rs app.RequestScope) (int, error) {
	return bs.dao.Count(rs)
}

// Get get all users from dao
func (bs *UserService) Get(rs app.RequestScope, id int) (*models.User, error) {
	return bs.dao.Get(rs, id)
}

// Create create a user
func (bs *UserService) Create(rs app.RequestScope, user *models.User) (*models.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	existedUser, err := bs.dao.GetByEmail(rs, user.Email)

	if err == nil {
		return existedUser, nil
	}

	if err := bs.dao.Create(rs, user); err != nil {
		return nil, err
	}

	return user, nil
}

// Update update a user
func (bs *UserService) Update(rs app.RequestScope, id int, user *models.User) (*models.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := bs.dao.Update(rs, id, user); err != nil {
		return nil, err
	}

	return user, nil
}

// Delete delete a user by id
func (bs *UserService) Delete(rs app.RequestScope, id int) (*models.User, error) {
	user, err := bs.Get(rs, id)

	if err != nil {
		return nil, err
	}

	err = bs.dao.Delete(rs, id)

	return user, err
}
