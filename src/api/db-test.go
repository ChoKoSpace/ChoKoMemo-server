package api

import (
	"github.com/ChoKoSpace/ChoKoMemo-server/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Memo struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
}

var db *gorm.DB

func InitDB() {
	var err error = nil
	db, err = gorm.Open(mysql.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Exec("CREATE DATABASE IF NOT EXISTS CHOKO_MEMO").Exec("USE CHOKO_MEMO")
	db.AutoMigrate(&Memo{})
}

func GetFromDB() {
}
