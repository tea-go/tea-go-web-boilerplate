package models

import (
	"fmt"
	"net/url"

	"github.com/tea-go/tea-go-web-boilerplate/pkg/app"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// stores data in mysql db
type db struct {
	db *gorm.DB
}

// NewDB return a new mysql db
func NewDB() (*db, error) {
	var err error

	s := new(db)

	dbHost := app.Config.DBHost
	dbPort := app.Config.DBPort
	dbName := app.Config.DBName
	dbUser := app.Config.DBUser
	dbPass := app.Config.DBPass

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	params := url.Values{}
	params.Add("parseTime", "1")
	params.Add("loc", "Asia/Shanghai")

	dsn := fmt.Sprintf("%s?%s", connection, params.Encode())

	s.db, err = gorm.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return s, nil
}

// AutoMigrate automatically migrate schemas
func (s *db) AutoMigrate() *db {
	s.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	return s
}

func (s *db) Close() error {
	return s.db.Close()
}

func (s *db) DB() *gorm.DB {
	return s.db
}
