package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ChoKoSpace/ChoKoMemo-server/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Memo struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
}

type TempResponseJson struct {
	Id      uint   `json:"memoId"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var db *gorm.DB

func InitDB() {
	var err error = nil
	db, err = gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Exec("CREATE DATABASE IF NOT EXISTS CHOKO_MEMO").Exec("USE CHOKO_MEMO")
	db.AutoMigrate(&Memo{})
}

func Db_test(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("DB-Test %v\n", r.Method)

	switch r.Method {
	default:
		http.NotFound(w, r)

	case http.MethodGet:
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		var memos []Memo

		result := db.Find(&memos)
		var Response = make([]TempResponseJson, result.RowsAffected)
		for i := 0; i < len(memos); i++ {
			Response[i].Id = memos[i].ID
			Response[i].Title = memos[i].Title
			Response[i].Content = memos[i].Content
			//fmt.Printf("%s / %s\n", memos[i].Title, memos[i].Content)
		}

		data, _ := json.MarshalIndent(Response, "", "    ")
		fmt.Fprintf(w, string(data))

	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		var Request Memo
		decoder.Decode(&Request)

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		newMemo := Memo{Title: Request.Title, Content: Request.Content}
		db.Create(&newMemo)
		fmt.Printf("[DB-TEST] Create Memo %v", newMemo)

	case http.MethodDelete:
	}
}
