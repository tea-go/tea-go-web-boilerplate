package daos

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/app"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/models"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/testdata"
)

func TestUserDAO(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewUserDAO()

	{
		// Count
		testDBCall(db, func(rs app.RequestScope) {
			count, err := dao.Count(rs)
			assert.Nil(t, nil, err)
			assert.Equal(t, 2, count)
		})
	}

	{
		// Query
		testDBCall(db, func(rs app.RequestScope) {
			users, err := dao.Query(rs, 1, 10)
			assert.Nil(t, nil, err)
			assert.Equal(t, 2, len(users))
		})
	}

	{
		// Get
		testDBCall(db, func(rs app.RequestScope) {
			user, err := dao.Get(rs, 1)
			assert.Nil(t, nil, err)
			assert.Equal(t, 1, user.ID)
			assert.Equal(t, "baiyu", user.Name)
			assert.Equal(t, "baiyu@qq.com", user.Email)
		})
	}

	{
		// Create
		testDBCall(db, func(rs app.RequestScope) {
			user := &models.User{
				Name:  "zhangsan",
				Email: "zhangsan@qq.com",
				Role:  "member",
			}

			err := dao.Create(rs, user)

			assert.Nil(t, nil, err)
			assert.Equal(t, 3, user.ID)
			assert.Equal(t, "zhangsan", user.Name)
			assert.Equal(t, "zhangsan@qq.com", user.Email)
		})
	}

	{
		// Update
		testDBCall(db, func(rs app.RequestScope) {
			user := &models.User{
				Name:  "jason1",
				Email: "jason@qq.com",
				Role:  "member",
			}

			err := dao.Update(rs, 2, user)

			assert.Nil(t, nil, err)
			assert.Equal(t, 2, user.ID)
			assert.Equal(t, "jason1", user.Name)
			assert.Equal(t, "jason@qq.com", user.Email)
		})
	}
}
