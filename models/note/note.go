package note

import (
	"gorm.io/gorm"
)

type Note struct {
	ID 		uint 	`json:"id" gorm:"primarykay"`
	Title	string	`json:"title"`
	Content string	`json:"content"`
}

func AutoMigrate(tx *gorm.DB) error {
	return tx.AutoMigrate(&Note{})
}