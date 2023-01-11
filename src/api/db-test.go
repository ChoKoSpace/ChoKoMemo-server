package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ChoKoSpace/ChoKoMemo-server/src/model"
)

type TempResponseJson struct {
	Id      uint   `json:"memoId"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func Db_test(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("DB-Test %v\n", r.Method)

	db := model.GetDB()
	if db == nil {
		panic("no DB")
	}

	switch r.Method {
	default:
		http.NotFound(w, r)

	case http.MethodGet:
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		var memos []model.Memo

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
		var Request model.Memo
		decoder.Decode(&Request)

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		newMemo := model.Memo{Title: Request.Title, Content: Request.Content}
		db.Create(&newMemo)
		fmt.Printf("[DB-TEST] Create Memo %v", newMemo)

	case http.MethodDelete:
	}
}
