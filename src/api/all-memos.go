package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ChoKoSpace/ChoKoMemo-server/src/session"
)

type MemoListInfo struct {
	MemoId int    `json:"memoId"`
	Title  string `json:"title"`
}

type GetAllMemoRequestJson struct {
	UserId string `json:"userId"`
	Token  string `json:"token"`
}

type GetAllMemoResponseJson struct {
	Error    *ErrorObject    `json:"error,omitempty"`
	MemoList *[]MemoListInfo `json:"memoList,omitempty"`
}

func AllMemo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		http.NotFound(w, r)

	case http.MethodGet:
		decoder := json.NewDecoder(r.Body)
		var Request GetAllMemoRequestJson
		decoder.Decode(&Request)

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		var Response = GetAllMemoResponseJson{}

		if session.IsValidToken(Request.UserId, Request.Token) {
			session.RefreshSession(Request.UserId)
			//test response
			countOfMemos := 3
			var newMemoList = make([]MemoListInfo, countOfMemos)
			for i := 0; i < countOfMemos; i++ {
				newMemoList[i].MemoId = i
				newMemoList[i].Title = "memo-title"
			}
			Response.MemoList = &newMemoList
		} else {
			errorObj := ErrorObject{}
			errorObj.Message = "Invalid token"
			Response.Error = &errorObj
			data, _ := json.MarshalIndent(Response, "", "    ")
			fmt.Fprintf(w, string(data))
		}
	}
}
