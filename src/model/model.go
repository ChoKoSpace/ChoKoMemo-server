package model

import (
	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	LoginId        string `gorm:"not null;unique"`
	HashedPassword string `gorm:"not null;"`
	Salt           string `gorm:"not null;"`
}

type Memo struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
}
