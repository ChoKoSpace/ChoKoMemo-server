package model

import (
	"time"

	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	LoginId        string `gorm:"not null;unique"`
	HashedPassword string `gorm:"not null;"`
	Salt           string `gorm:"not null;"`
}

type Session struct {
	gorm.Model
	UserID    int       `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

type Memo struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
}
