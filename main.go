package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorObject struct {
	Message string `json:"message"`
}

type LoginRequestJson struct {
	LoginId string `json:"loginId"`
}

type LoginResponseJson struct {
	Error  *ErrorObject `json:"error,omitempty"`
	UserId *string      `json:"userId,omitempty"`
	Token  *string      `json:"token,omitempty"`
}

type MemoListInfo struct {
	MemoId int    `json:"memoId"`
	Title  string `json:"title"`
}

type GetAllMemoListRequestJson struct {
	LoginId string `json:"loginId"`
	Token   string `json:"token"`
}

type GetAllMemoListResponseJson struct {
	Error    *ErrorObject    `json:"error,omitempty"`
	MemoList *[]MemoListInfo `json:"memoList,omitempty"`
}

func mainRounter(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[14:]

	switch url {
	default:
		http.NotFound(w, r)
	case "/login":
		login(w, r)
	case "/all-memos":
		getAllMemoList(w, r)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		http.NotFound(w, r)

	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		var Request LoginRequestJson
		decoder.Decode(&Request)

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		var Response = LoginResponseJson{}

		if Request.LoginId != "1" {
			errorObj := ErrorObject{}
			errorObj.Message = "Invalid user"
			Response.Error = &errorObj
		} else {
			_userId := "1234"
			Response.UserId = &_userId
			_token := "temp-token"
			Response.Token = &_token
		}
		data, _ := json.MarshalIndent(Response, "", "    ")
		fmt.Fprintf(w, string(data))
	}
}

func getAllMemoList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		http.NotFound(w, r)

	case http.MethodGet:
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		var Response = GetAllMemoListResponseJson{}
		//test response
		countOfMemos := 3
		var newMemoList = make([]MemoListInfo, countOfMemos)
		for i := 0; i < countOfMemos; i++ {
			newMemoList[i].MemoId = i
			newMemoList[i].Title = "memo-title"
		}
		Response.MemoList = &newMemoList

		data, _ := json.MarshalIndent(Response, "", "    ")
		fmt.Fprintf(w, string(data))
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome ChoKoSpace Server")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexPage)
	mux.HandleFunc("/status", healthCheck)
	mux.HandleFunc("/api/chokomemo/", mainRounter)
	http.ListenAndServe(":3901", mux)
}
