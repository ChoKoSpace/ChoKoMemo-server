package model

import (
	"time"

	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	AccountId      string `gorm:"not null;unique"`
	HashedPassword string `gorm:"not null;"`
	Salt           string `gorm:"not null;"`
}

type Session struct {
	gorm.Model
	UserID    int       `gorm:"not null" json:"user_id"`
	Token     string    `gorm:"not null" json:"token"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
}

type Memo struct {
	gorm.Model
	UserId  string `gorm:"not null"`
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
}
