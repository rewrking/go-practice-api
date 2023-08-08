package config

import (
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect(dialector gorm.Dialector) {
	database, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = database
}

func GetDB() *gorm.DB {
	return db
}
