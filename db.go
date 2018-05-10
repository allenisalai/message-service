package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func init() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=alai dbname=alai sslmode=disable")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&Message{})
}

func GetDb() *gorm.DB {
	return db
}
