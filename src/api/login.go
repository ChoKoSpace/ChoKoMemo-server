package api

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ChoKoSpace/ChoKoMemo-server/src/model"
)

type LoginRequestJson struct {
	LoginId  string `json:"loginId"`
	Password string `json:"password"`
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

		fmt.Printf("[Login]-> id: %s, password: %s \n", Request.LoginId, Request.Password)

		var targetUser model.UserInfo
		db := model.GetDB()
		db.First(&targetUser, "login_Id = ?", Request.LoginId)

		fmt.Printf("[Login]-> %v \n", targetUser)
		saltBytes, err := hex.DecodeString(targetUser.Salt)
		if err != nil {
			panic("Failed to decode salf from db")
		}
		sha := sha512.New()
		passwordBytes := append([]byte(Request.Password), saltBytes...)
		sha.Write(passwordBytes)
		hash := sha.Sum(nil)

		fmt.Printf("[Login]-> %s \n", hex.EncodeToString(hash))

		if hex.EncodeToString(hash) != targetUser.HashedPassword {
			errorObj := ErrorObject{}
			errorObj.Message = "Invalid user"
			Response.Error = &errorObj
		} else {
			ret_UserId := fmt.Sprintf("%d", targetUser.ID)
			Response.UserId = &ret_UserId
			_token := "temp-token"
			Response.Token = &_token
		}
		data, _ := json.MarshalIndent(Response, "", "    ")
		fmt.Fprintf(w, string(data))
	}
}
