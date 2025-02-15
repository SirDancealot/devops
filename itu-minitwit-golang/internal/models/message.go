package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	AuthorID uint
	User     User
}
