package model

import (
	"fmt"

	"github.com/ChoKoSpace/ChoKoMemo-server/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	if db == nil {
		panic("DB is not initialized.")
	}
	return db
}

func InitializeDB() {
	var err error
	db, err = gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the DB")
	}

	db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", config.NAME_DATABASE))
	db.Exec(fmt.Sprintf("USE %s;", config.NAME_DATABASE))

	db.AutoMigrate(&UserInfo{})
	db.AutoMigrate(&Session{})
	db.AutoMigrate(&Memo{})
}
