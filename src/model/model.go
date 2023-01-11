package model

import "gorm.io/gorm"

type Memo struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
}
