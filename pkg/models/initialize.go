package models

import (
	"github.com/rewrking/go-practice-api/pkg/config"
)

func Initialize(dbFile string) {
	config.Connect(dbFile)

	db := config.GetDB()
	db.AutoMigrate(&Book{})
}
