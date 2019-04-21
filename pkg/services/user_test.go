package services

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/tea-go/tea-go-web-boilerplate/pkg/app"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/models"
)

func TestNewUserService(t *testing.T) {
	dao := newMockUserDAO()
	s := NewUserService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestUserService_Get(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	user, err := s.Get(nil, 1)

	if assert.Nil(t, err) && assert.NotNil(t, user) {
		assert.Equal(t, "liuyi", user.Name)
	}

	user, err = s.Get(nil, 100)
	assert.NotNil(t, err)
}

func TestUserService_Query(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	users, err := s.Query(nil, 1, 2)

	if assert.Nil(t, err) {
		assert.Equal(t, 2, len(users))
	}
}

func TestNewUserService_Create(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	user, err := s.Create(nil, &models.User{
		Name:  "lisi",
		Email: "lisi@qq.com",
		Role:  "member",
	})

	if assert.Nil(t, err) && assert.NotNil(t, user) {
		assert.Equal(t, 4, user.ID)
		assert.Equal(t, "lisi", user.Name)
	}

	var u1 models.User
	u1.ID = 100
	u1.Name = "aaa"

	_, err = s.Create(nil, &u1)
	assert.NotNil(t, err)

	var u2 models.User
	u2.ID = 100
	u2.Name = "test"
	u2.Email = "test"

	_, err = s.Create(nil, &u2)

	assert.NotNil(t, err)
}

func TestNewUserService_Update(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	var u1 models.User
	u1.Name = "ddd"
	u1.Email = "liuer@qq.com"
	u1.Status = "enabled"
	u1.Role = "member"

	user, err := s.Update(nil, 2, &u1)

	if assert.Nil(t, err) && assert.NotNil(t, user) {
		assert.Equal(t, 2, user.ID)
		assert.Equal(t, "ddd", user.Name)
	}

	var u2 models.User
	u2.Name = "ddd"

	_, err = s.Update(nil, 100, &u2)
	assert.NotNil(t, err)

	var u3 models.User
	u3.Name = ""

	_, err = s.Update(nil, 2, &u3)

	assert.NotNil(t, err)
}

func TestNewUserService_Delete(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	user, err := s.Delete(nil, 2)

	if assert.Nil(t, err) && assert.NotNil(t, user) {
		assert.Equal(t, 2, user.ID)
		assert.Equal(t, "chener", user.Name)
	}

	_, err = s.Delete(nil, 2)
	assert.NotNil(t, err)
}

func newMockUserDAO() userDAO {
	var u1 models.User

	u1.ID = 1
	u1.Name = "liuyi"
	u1.Email = "liuyi@qq.com"
	u1.Status = "enabled"
	u1.Role = "admin"
	u1.CreatedAt = time.Now()
	u1.UpdatedAt = time.Now()
	u1.IsDeleted = "no"

	var u2 models.User

	u2.ID = 2
	u2.Name = "chener"
	u2.Email = "chener@qq.com"
	u2.Status = "enabled"
	u2.Role = "member"
	u2.CreatedAt = time.Now()
	u2.UpdatedAt = time.Now()
	u2.IsDeleted = "no"

	var u3 models.User

	u3.ID = 3
	u3.Name = "zhangsan"
	u3.Email = "zhangsan@qq.com"
	u3.Status = "enabled"
	u3.Role = "member"
	u3.CreatedAt = time.Now()
	u3.UpdatedAt = time.Now()
	u3.IsDeleted = "no"

	return &mockUserDAO{
		records: []models.User{u1, u2, u3},
	}
}

type mockUserDAO struct {
	records []models.User
}

// Query query all users
func (m *mockUserDAO) Query(rs app.RequestScope, offset, limit int) ([]models.User, error) {
	return m.records[offset : offset+limit], nil
}

// Get get a user detail
func (m *mockUserDAO) Get(rs app.RequestScope, id int) (*models.User, error) {
	for _, record := range m.records {
		if record.ID == id {
			return &record, nil
		}
	}

	return nil, errors.New("user not found")
}

// Count count all users
func (m *mockUserDAO) Count(rs app.RequestScope) (int, error) {
	return len(m.records), nil
}

// Create create a user
func (m *mockUserDAO) Create(rs app.RequestScope, user *models.User) error {
	if user.ID != 0 {
		return errors.New("ID cannot be set")
	}

	user.ID = len(m.records) + 1
	m.records = append(m.records, *user)

	return nil
}

// Update update a user
func (m *mockUserDAO) Update(rs app.RequestScope, id int, user *models.User) error {
	user.ID = id

	for i, record := range m.records {
		if record.ID == id {
			m.records[i] = *user
			return nil
		}
	}

	return errors.New("user not found")
}

// GetByEmail get a user by email
func (m *mockUserDAO) GetByEmail(rs app.RequestScope, email string) (*models.User, error) {
	for _, record := range m.records {
		if record.Email == email {
			return &record, nil
		}
	}

	return nil, errors.New("user not found")
}

// Delete delete a user by id
func (m *mockUserDAO) Delete(rs app.RequestScope, id int) error {
	for i, record := range m.records {
		if record.ID == id {
			m.records = append(m.records[:i], m.records[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}
