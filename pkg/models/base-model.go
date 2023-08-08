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

func CreateOne[T any](data *T) *T {
	db := config.GetDB()
	db.Create(data)
	return data
}

func ReadAll[T any]() []T {
	db := config.GetDB()
	var data []T
	db.Find(&data)
	return data
}

func ReadOne[T any](ID int64) *T {
	var data T
	db := config.GetDB()
	count := int64(0)
	db.Where("ID = ?", ID).Find(&data).Count(&count)
	if count == 0 {
		return nil
	}
	return &data
}

func DeleteOne[T any](ID int64) *DeletedModel {
	data := ReadOne[T](ID)
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
