package api

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ChoKoSpace/ChoKoMemo-server/src/model"
)

type SignupRequestJson struct {
	AccountId string `json:"accountId"`
	Password  string `json:"password"`
}

type SignupResponseJson struct {
	Error     *ErrorObject `json:"error,omitempty"`
	IsSuccess bool         `json:"isSuccess"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		http.NotFound(w, r)

	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		var Request SignupRequestJson
		decoder.Decode(&Request)

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		var Response = SignupResponseJson{}
		errObj := ErrorObject{}
		Response.IsSuccess = false

		salt := make([]byte, 16)
		_, err := rand.Read(salt)
		if err != nil {
			errObj.Message = "Failed to generate salt"
		} else {
			sha := sha512.New()
			passwordBytes := append([]byte(Request.Password), salt...)
			sha.Write(passwordBytes)
			hash := sha.Sum(nil)

			newUserInfo := model.UserInfo{}
			newUserInfo.AccountId = Request.AccountId
			newUserInfo.HashedPassword = hex.EncodeToString(hash)
			newUserInfo.Salt = hex.EncodeToString(salt)

			if err := model.GetDB().Create(&newUserInfo).Error; err == nil {
				Response.IsSuccess = true
			} else {
				errObj.Message = err.Error()
			}
		}

		if len(errObj.Message) != 0 {
			Response.Error = &errObj
		}
		data, _ := json.MarshalIndent(Response, "", "    ")
		fmt.Fprintf(w, string(data))
	}
}
