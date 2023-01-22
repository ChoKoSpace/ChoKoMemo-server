package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ChoKoSpace/ChoKoMemo-server/src/model"
	"github.com/ChoKoSpace/ChoKoMemo-server/src/session"
)

func Memo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		http.NotFound(w, r)

	case http.MethodGet:
		GetMemo(w, r)
	case http.MethodPost:
		PostMemo(w, r)
	case http.MethodPut:
		PutMemo(w, r)
	case http.MethodDelete:
		DeleteMemo(w, r)
	}
}

type GetMemoRequestJson struct {
	UserId string `json:"userId"`
	Token  string `json:"token"`
	MemoId uint   `json:"memoId"`
}

type GetMemoResponseJson struct {
	Error     *ErrorObject `json:"error,omitempty"`
	Title     string       `json:"title"`
	Content   string       `json:"content"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
}

func GetMemo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	var Request GetMemoRequestJson
	var Response = GetMemoResponseJson{}
	errorObj := ErrorObject{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&Request)

	if session.IsValidToken(Request.UserId, Request.Token) {
		session.RefreshSession(Request.UserId)

		db := model.GetDB()
		searchTarget := model.Memo{}
		searchTarget.ID = Request.MemoId
		searchTarget.UserId = Request.UserId
		var foundMemo model.Memo
		if err := db.Model(&model.Memo{}).Where(&searchTarget).First(&foundMemo).Error; err == nil {
			Response.Title = foundMemo.Title
			Response.Content = foundMemo.Content
			Response.CreatedAt = foundMemo.CreatedAt
			Response.UpdatedAt = foundMemo.UpdatedAt
		} else {
			errorObj.Message = "Not found memo"
		}
	} else {
		errorObj.Message = "Invalid user"
	}

	if len(errorObj.Message) > 0 {
		Response.Error = &errorObj
	}
	data, _ := json.MarshalIndent(Response, "", "    ")
	fmt.Fprintf(w, string(data))
}

type PostMemoRequestJson struct {
	UserId  string `json:"userId"`
	Token   string `json:"token"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostMemoResponseJson struct {
	Error     *ErrorObject `json:"error,omitempty"`
	IsSuccess bool         `json:"isSuccess"`
	MemoId    uint         `json:"memoId"`
}

func PostMemo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	var Request PostMemoRequestJson
	var Response = PostMemoResponseJson{}
	errorObj := ErrorObject{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&Request)

	if session.IsValidToken(Request.UserId, Request.Token) {
		session.RefreshSession(Request.UserId)

		db := model.GetDB()
		newMemo := model.Memo{}
		newMemo.UserId = Request.UserId
		newMemo.Title = Request.Title
		newMemo.Content = Request.Content

		if err := db.Create(&newMemo).Error; err == nil {
			Response.IsSuccess = true
			Response.MemoId = newMemo.ID
		} else {
			errorObj.Message = "Failed to create memo"
		}
	} else {
		errorObj.Message = "Invalid user"
	}

	if len(errorObj.Message) > 0 {
		Response.Error = &errorObj
	}
	data, _ := json.MarshalIndent(Response, "", "    ")
	fmt.Fprintf(w, string(data))
}

type PutMemoRequestJson struct {
	UserId  string `json:"userId"`
	Token   string `json:"token"`
	MemoId  uint   `json:"memoId"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PutMemoResponseJson struct {
	Error     *ErrorObject `json:"error,omitempty"`
	IsSuccess bool         `json:"isSuccess"`
}

func PutMemo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	var Request PutMemoRequestJson
	var Response = PutMemoResponseJson{}
	errorObj := ErrorObject{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&Request)

	if session.IsValidToken(Request.UserId, Request.Token) {
		session.RefreshSession(Request.UserId)

		db := model.GetDB()
		searchTarget := model.Memo{}
		searchTarget.ID = Request.MemoId
		if row := db.Model(&model.Memo{}).Where(&searchTarget); row != nil {
			if err := row.Updates(model.Memo{Title: Request.Title, Content: Request.Content}).Error; err == nil {
				Response.IsSuccess = true
			} else {
				errorObj.Message = err.Error()
			}
		}
	} else {
		errorObj.Message = "Invalid user"
	}

	if len(errorObj.Message) > 0 {
		Response.Error = &errorObj
	}
	data, _ := json.MarshalIndent(Response, "", "    ")
	fmt.Fprintf(w, string(data))
}

type DeleteMemoRequestJson struct {
	UserId  string `json:"userId"`
	Token   string `json:"token"`
	MemoIds []uint `json:"memoIds"`
}

type DeleteMemoResponseJson struct {
	Error          *ErrorObject `json:"error,omitempty"`
	DeletedMemoIds []uint       `json:"deletedMemoIds"`
}

func DeleteMemo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	var Request DeleteMemoRequestJson
	var Response = DeleteMemoResponseJson{}
	errorObj := ErrorObject{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&Request)

	if session.IsValidToken(Request.UserId, Request.Token) {
		session.RefreshSession(Request.UserId)

		db := model.GetDB()
		for _, memoId := range Request.MemoIds {
			targetMemo := model.Memo{}
			targetMemo.ID = memoId
			targetMemo.UserId = Request.UserId
			deletedRow := db.Model(&targetMemo).Delete(&targetMemo)
			if err := deletedRow.Error; err == nil {
				if deletedRow.RowsAffected > 0 {
					Response.DeletedMemoIds = append(Response.DeletedMemoIds, memoId)
				} else {
					errorObj.Message = "no data"
				}
			} else {
				errorObj.Message = err.Error()
			}
		}
	} else {
		errorObj.Message = "Invalid user"
	}

	if len(errorObj.Message) > 0 {
		Response.Error = &errorObj
	}
	data, _ := json.MarshalIndent(Response, "", "    ")
	fmt.Fprintf(w, string(data))
}
