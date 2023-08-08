package models

import (
	"github.com/rewrking/go-practice-api/pkg/config"
)

type Book struct {
	BaseModel
	Name        string `gorm:"name" json:"name"`
	Author      string `gorm:"author" json:"author"`
	Publication string `gorm:"publication" json:"publication"`
}

func (book *Book) Update(updates *Book) {
	if updates.Name != "" {
		book.Name = updates.Name
	}
	if updates.Author != "" {
		book.Author = updates.Author
	}
	if updates.Publication != "" {
		book.Publication = updates.Publication
	}

	config.GetDB().Save(book)
}
