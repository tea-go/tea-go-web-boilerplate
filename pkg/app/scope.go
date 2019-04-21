package app

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

// RequestScope contains the application-specific information that are carried around in a request.
type RequestScope interface {
	Logger
	// UserID returns the ID of the user for the current request
	UserID() string
	// UserRole returns the role of the user for the current request
	UserRole() string
	// SetUserID sets the ID of the currently authenticated user
	SetUserID(id string)
	// RequestID returns the ID of the current request
	RequestID() string
	// Tx returns the currently active database transaction that can be used for DB query purpose
	Tx() *gorm.DB
	// SetTx sets the database transaction
	SetTx(tx *gorm.DB)
	// Rollback returns a value indicating whether the current database transaction should be rolled back
	Rollback() bool
	// SetRollback sets a value indicating whether the current database transaction should be rolled back
	SetRollback(bool)
	// Now returns the timestamp representing the time when the request is being processed
	Now() time.Time
}

type requestScope struct {
	Logger              // the logger tagged with the current request information
	now       time.Time // the time when the request is being processed
	requestID string    // an ID identifying one or multiple correlated HTTP requests
	userID    string    // an ID identifying the current user
	userRole  string    // role of the current user
	rollback  bool      // whether to roll back the current transaction
	tx        *gorm.DB  // the currently active transaction
}

func (rs *requestScope) UserID() string {
	return rs.userID
}

func (rs *requestScope) UserRole() string {
	return rs.userRole
}

func (rs *requestScope) SetUserID(id string) {
	rs.Logger.SetField("UserID", id)
	rs.userID = id
}

func (rs *requestScope) RequestID() string {
	return rs.requestID
}

func (rs *requestScope) Tx() *gorm.DB {
	return rs.tx
}

func (rs *requestScope) SetTx(tx *gorm.DB) {
	rs.tx = tx
}

func (rs *requestScope) Rollback() bool {
	return rs.rollback
}

func (rs *requestScope) SetRollback(v bool) {
	rs.rollback = v
}

func (rs *requestScope) Now() time.Time {
	return rs.now
}

// newRequestScope creates a new RequestScope with the current request information.
func newRequestScope(now time.Time, logger *logrus.Logger, ctx *gin.Context) RequestScope {
	l := NewLogger(logger, logrus.Fields{})

	requestID := ctx.Request.Header.Get("X-Request-Id")

	if requestID != "" {
		l.SetField("RequestID", requestID)
	}
	return &requestScope{
		Logger:    l,
		now:       now,
		requestID: requestID,
	}
}
