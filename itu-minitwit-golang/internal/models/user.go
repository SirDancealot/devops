package models

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"primaryKey"`
	Username  string     `gorm:"not null"`
	Email     string     `gorm:"not null"`
	PwHash    string     `gorm:"not null"`
	Following []User     `gorm:"many2many:follower"`
	Messages  []Message  `gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Override default table name
func (User) TableName() string {
    return "user"
}
