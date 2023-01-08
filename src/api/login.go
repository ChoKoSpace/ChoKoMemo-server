package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginRequestJson struct {
	LoginId string `json:"loginId"`
}

type LoginResponseJson struct {
	Error  *ErrorObject `json:"error,omitempty"`
	UserId *string      `json:"userId,omitempty"`
	Token  *string      `json:"token,omitempty"`
}

func Login(w http.ResponseWriter, r *http.Request) {
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
