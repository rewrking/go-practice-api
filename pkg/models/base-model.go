package models

import (
	"time"

	"github.com/rewrking/go-practice-api/pkg/config"
)

type BaseModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	// Enables "soft delete"
	// DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type DeletedModel struct {
	ID        uint      `json:"id"`
	DeletedAt time.Time `json:"deletedAt"`
}

func Create[T any](data *T) *T {
	db := config.GetDB()
	db.Create(data)
	return data
}

func GetAll[T any]() []T {
	db := config.GetDB()
	var data []T
	db.Find(&data)
	return data
}

func GetById[T any](ID int64) *T {
	var data T
	db := config.GetDB()
	count := int64(0)
	db.Where("ID = ?", ID).Find(&data).Count(&count)
	if count == 0 {
		return nil
	}
	return &data
}

func DeleteById[T any](ID int64) *DeletedModel {
	data := GetById[T](ID)
	if data == nil {
		return nil
	}

	db := config.GetDB()
	db.Delete(&data)

	var result DeletedModel
	result.ID = uint(ID)
	result.DeletedAt = time.Now()
	return &result
}
