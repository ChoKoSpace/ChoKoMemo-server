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

func mainRounter(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[14:]

	switch url {
	default:
		http.NotFound(w, r)
	case "/login":
		login(w, r)
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
		data, _ := json.Marshal(Response)
		fmt.Fprintf(w, string(data))
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/status", healthCheck)
	mux.HandleFunc("/api/chokomemo/", mainRounter)
	http.ListenAndServe(":3901", mux)
}
