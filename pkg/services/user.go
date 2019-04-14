package services

import (
	"github.com/tea-go/tea-go-web-boilerplate/pkg/app"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/models"
)

type userDAO interface {
	Query(rs app.RequestScope, offset, limit int) []models.User
	Count(rs app.RequestScope) int
	Get(rs app.RequestScope, id int) *models.User
	GetByEmail(rs app.RequestScope, email string) *models.User
	Create(rs app.RequestScope, user *models.User) *models.User
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
func (bs *UserService) Query(rs app.RequestScope, offset, limit int) []models.User {
	return bs.dao.Query(rs, offset, limit)
}

// Count get the count of user
func (bs *UserService) Count(rs app.RequestScope) int {
	return bs.dao.Count(rs)
}

// Get get all users from dao
func (bs *UserService) Get(rs app.RequestScope, id int) *models.User {
	return bs.dao.Get(rs, id)
}

// Create create a user
func (bs *UserService) Create(rs app.RequestScope, user *models.User) (*models.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if data := bs.dao.GetByEmail(rs, user.Email); data != nil {
		return data, nil
	}

	data := bs.dao.Create(rs, user)

	return data, nil
}
