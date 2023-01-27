package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ChoKoSpace/ChoKoMemo-server/src/model"
	"github.com/ChoKoSpace/ChoKoMemo-server/src/session"
)

type GetAllMemoRequestJson struct {
	UserId string `json:"userId"`
	Token  string `json:"token"`
}

type MemoListInfo struct {
	MemoId int    `json:"memoId"`
	Title  string `json:"title"`
}

type GetAllMemoResponseJson struct {
	Error    *ErrorObject   `json:"error,omitempty"`
	MemoList []MemoListInfo `json:"memoList"`
}

func AllMemo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		http.NotFound(w, r)

	case http.MethodGet:
		decoder := json.NewDecoder(r.Body)
		var Request GetAllMemoRequestJson
		errorObj := ErrorObject{}
		decoder.Decode(&Request)

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		var Response = GetAllMemoResponseJson{}

		if session.IsValidToken(Request.UserId, Request.Token) {
			session.RefreshSession(Request.UserId)

			db := model.GetDB()
			if db != nil {
				var count int64
				db.Model(&model.Memo{}).Where(&model.Memo{UserId: Request.UserId}).Count(&count)
				memos := make([]model.Memo, count)
				db.Where(&model.Memo{UserId: Request.UserId}).Find(&memos)

				var i int64
				for i = 0; i < count; i++ {
					Response.MemoList = append(Response.MemoList, MemoListInfo{MemoId: int(memos[i].ID), Title: memos[i].Title})
				}
			}
		} else {
			errorObj.Message = append(errorObj.Message, "Invalid token")
		}

		if len(errorObj.Message) > 0 {
			Response.Error = &errorObj
		}
		data, _ := json.MarshalIndent(Response, "", "    ")
		fmt.Fprintf(w, string(data))
	}
}
