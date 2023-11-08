package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	var err error
	DBConn, err = gorm.Open("postgres", "host=localhost port=5432 user=root dbname=postgres password=root_pwd sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection opened to database")
}
