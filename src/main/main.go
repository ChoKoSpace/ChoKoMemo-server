package main

import (
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
	case "/signup":
		api.Signup(w, r)
	case "/signin":
		api.Signin(w, r)
	case "/all-memos":
		api.AllMemo(w, r)
	case "/memo":
		api.Memo(w, r)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	model.InitializeDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/status", healthCheck)
	mux.HandleFunc("/api/chokomemo/", apiRounter)
	http.ListenAndServe(config.PORT, mux)
}
