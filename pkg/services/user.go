package services

import (
	"github.com/tea-go/tea-go-web-boilerplate/pkg/app"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/models"
)

type userDAO interface {
	Query(rs app.RequestScope, offset, limit int) []models.User
	Count(rs app.RequestScope) int
	Get(rs app.RequestScope, id int) *models.User
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
