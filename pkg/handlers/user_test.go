package handlers

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tea-go/tea-go-web-boilerplate/pkg/daos"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/models"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/services"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/testdata"
)

type UserListResponseBody struct {
	ListResponseBody
	Data []models.User
}

type UserDetailResponseBody struct {
	DetailResponseBody
	Data models.User
}

func TestUser(t *testing.T) {
	testdata.ResetDB()

	r := newRouter()

	userDao := daos.NewUserDAO()
	userService := services.NewUserService(userDao)

	HanldeUserResource(r, userService)

	runAPITests(t, r, []apiTestCase{
		{
			"t1 - get a list of users",
			"GET",
			"/users",
			"",
			http.StatusOK,
			func(body string) {
				data := &UserListResponseBody{}
				if err := json.Unmarshal([]byte(body), data); err != nil {
					assert.Equal(t, 2, data.Count)
					assert.Equal(t, "success", data.Status)
					assert.Equal(t, 200, data.StatusCode)
					assert.Equal(t, 2, len(data.Data))
				}
			},
		},
		{
			"t2 - get a list of users by limit 1",
			"GET",
			"/users?limit=1",
			"",
			http.StatusOK,
			func(body string) {
				data := &UserListResponseBody{}
				if err := json.Unmarshal([]byte(body), data); err != nil {
					assert.Equal(t, 2, data.Count)
					assert.Equal(t, "success", data.Status)
					assert.Equal(t, 200, data.StatusCode)
					assert.Equal(t, 1, len(data.Data))
				}
			},
		},
		{
			"t3 - get a user by id is 1",
			"GET",
			"/users/1",
			"",
			http.StatusOK,
			func(body string) {
				data := &UserDetailResponseBody{}
				if err := json.Unmarshal([]byte(body), data); err != nil {
					assert.Equal(t, "success", data.Status)
					assert.Equal(t, 200, data.StatusCode)
					assert.Equal(t, 1, data.Data.ID)
					assert.Equal(t, "baiyu", data.Data.Name)
					assert.Equal(t, "baiyu@qq.com", data.Data.Email)
					assert.Equal(t, "role", data.Data.Role)
					assert.Equal(t, "enabled", data.Data.Status)
					assert.Equal(t, "no", data.Data.IsDeleted)

				}
			},
		},
	})
}
