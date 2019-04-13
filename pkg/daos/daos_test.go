package daos

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/app"
)

func testDBCall(db *gorm.DB, f func(rs app.RequestScope)) {
	rs := mockRequestScope(db)

	defer func() {
		rs.Tx().Rollback()
	}()

	f(rs)
}

type requestScope struct {
	app.Logger
	tx *gorm.DB
}

func mockRequestScope(db *gorm.DB) app.RequestScope {
	tx := db.Begin()
	return &requestScope{
		tx: tx,
	}
}

func (rs *requestScope) UserID() string {
	return "tester"
}

func (rs *requestScope) SetUserID(id string) {
}

func (rs *requestScope) RequestID() string {
	return "test"
}

func (rs *requestScope) Tx() *gorm.DB {
	return rs.tx
}

func (rs *requestScope) SetTx(tx *gorm.DB) {
	rs.tx = tx
}

func (rs *requestScope) Rollback() bool {
	return false
}

func (rs *requestScope) SetRollback(v bool) {
}

func (rs *requestScope) Now() time.Time {
	return time.Now()
}
