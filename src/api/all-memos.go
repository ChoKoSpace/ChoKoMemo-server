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

type GetAllMemoListRequestJson struct {
	UserId string `json:"userId"`
	Token  string `json:"token"`
}

type GetAllMemoListResponseJson struct {
	Error    *ErrorObject    `json:"error,omitempty"`
	MemoList *[]MemoListInfo `json:"memoList,omitempty"`
}

func GetAllMemoList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		http.NotFound(w, r)

	case http.MethodGet:
		decoder := json.NewDecoder(r.Body)
		var Request GetAllMemoListRequestJson
		decoder.Decode(&Request)

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		var Response = GetAllMemoListResponseJson{}

		if !session.IsValidToken(Request.UserId, Request.Token) {
			errorObj := ErrorObject{}
			errorObj.Message = "Invalid token"
			Response.Error = &errorObj
		} else {
			session.RefreshSession(Request.UserId)
			//test response
			countOfMemos := 3
			var newMemoList = make([]MemoListInfo, countOfMemos)
			for i := 0; i < countOfMemos; i++ {
				newMemoList[i].MemoId = i
				newMemoList[i].Title = "memo-title"
			}
			Response.MemoList = &newMemoList
		}

		data, _ := json.MarshalIndent(Response, "", "    ")
		fmt.Fprintf(w, string(data))
	}
}
