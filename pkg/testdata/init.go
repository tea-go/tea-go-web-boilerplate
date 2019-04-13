package testdata

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/app"
)

var (
	// DB a db instance of mysql
	DB *gorm.DB
)

func init() {
	// the test may be started from the home directory or a subdirectory
	err := app.LoadConfig("../config")

	if err != nil {
		panic(err)
	}

	dbHost := "127.0.0.1"
	dbPort := "3306"
	dbName := "teagotest"
	dbUser := "root"
	dbPass := "12345678"

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	params := url.Values{}
	params.Add("parseTime", "1")
	params.Add("loc", "Asia/Shanghai")

	dsn := fmt.Sprintf("%s?%s", connection, params.Encode())

	DB, err = gorm.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}
}

// ResetDB re-create the database schema and re-populate the initial data using the SQL statements in db.sql.
// This method is mainly used in tests.
func ResetDB() *gorm.DB {
	if err := runSQLFile(DB, getSQLFile()); err != nil {
		panic(fmt.Errorf("Error while initializing test database: %s", err))
	}
	return DB
}

func getSQLFile() string {
	file, _ := filepath.Abs("../testdata/db.sql")

	if _, err := os.Stat(file); err == nil {
		return file
	}
	return "testdata/db.sql"
}

func runSQLFile(db *gorm.DB, file string) error {
	s, err := ioutil.ReadFile(file)

	if err != nil {
		return err
	}

	lines := strings.Split(string(s), ";")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		db.Exec(line)
	}

	return nil
}
