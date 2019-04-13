package daos

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/app"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/testdata"
)

func TestUserDAO(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewUserDAO()
	{
		// Query
		testDBCall(db, func(rs app.RequestScope) {
			users := dao.Query(rs, 1, 10)
			assert.Equal(t, 2, len(users))
		})
	}
}
