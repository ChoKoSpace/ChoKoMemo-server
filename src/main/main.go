package main

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"net/http"

	"github.com/ChoKoSpace/ChoKoMemo-server/src/api"
	"github.com/ChoKoSpace/ChoKoMemo-server/src/config"
	"github.com/ChoKoSpace/ChoKoMemo-server/src/model"
)

func apiRounter(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[14:]

	switch url {
	default:
		http.NotFound(w, r)
	case "/login":
		api.Login(w, r)
	case "/all-memos":
		api.GetAllMemoList(w, r)
	case "/db-test":
		api.Db_test(w, r)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func temp_AddUser(loginId string, password string) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		panic("Failed to generate salt")
	}

	sha := sha512.New()
	passwordBytes := append([]byte(password), salt...)
	sha.Write(passwordBytes)
	hash := sha.Sum(nil)

	newUserInfo := model.UserInfo{}
	newUserInfo.LoginId = loginId
	newUserInfo.HashedPassword = hex.EncodeToString(hash)
	newUserInfo.Salt = hex.EncodeToString(salt)

	model.GetDB().Create(&newUserInfo)
}

func main() {
	model.InitializeDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/status", healthCheck)
	mux.HandleFunc("/api/chokomemo/", apiRounter)
	http.ListenAndServe(config.PORT, mux)
}
