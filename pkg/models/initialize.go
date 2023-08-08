package models

import (
	"github.com/rewrking/go-practice-api/pkg/config"
	"gorm.io/gorm"
)

func Initialize(dialector gorm.Dialector) {
	config.Connect(dialector)

	db := config.GetDB()
	db.AutoMigrate(Book{})
}
