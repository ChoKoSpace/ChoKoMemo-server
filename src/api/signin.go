package api

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ChoKoSpace/ChoKoMemo-server/src/model"
	"github.com/ChoKoSpace/ChoKoMemo-server/src/session"
)

type SigninRequestJson struct {
	AccountId string `json:"accountId"`
	Password  string `json:"password"`
}

type SigninResponseJson struct {
	Error  *ErrorObject `json:"error,omitempty"`
	UserId *string      `json:"userId,omitempty"`
	Token  *string      `json:"token,omitempty"`
}

func Signin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		http.NotFound(w, r)

	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		var Request SigninRequestJson
		decoder.Decode(&Request)

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		var Response = SigninResponseJson{}

		var targetUser model.UserInfo
		db := model.GetDB()
		db.First(&targetUser, "account_id = ?", Request.AccountId)

		saltBytes, err := hex.DecodeString(targetUser.Salt)
		if err != nil {
			panic("Failed to decode salf from db")
		}
		sha := sha512.New()
		passwordBytes := append([]byte(Request.Password), saltBytes...)
		sha.Write(passwordBytes)
		hash := sha.Sum(nil)

		//비밀번호가 맞는지 확인
		if hex.EncodeToString(hash) != targetUser.HashedPassword {
			errorObj := ErrorObject{}
			errorObj.Message = "Invalid user"
			Response.Error = &errorObj
		} else {
			userIdStr := fmt.Sprintf("%d", targetUser.ID)
			Response.UserId = &userIdStr

			_token, err := session.CreateSession(userIdStr)
			if err != nil {
				errorObj := ErrorObject{}
				errorObj.Message = fmt.Sprintf("Error creating session (%v)", err)
				Response.Error = &errorObj
			} else {
				Response.Token = &_token
			}
		}
		data, _ := json.MarshalIndent(Response, "", "    ")
		fmt.Fprintf(w, string(data))
	}
}
