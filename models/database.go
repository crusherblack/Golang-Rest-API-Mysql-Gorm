package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func ConnectDatabase() {
	database, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_golang?charset=utf8&parseTime=True")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Booking{})

	DB = database
}
