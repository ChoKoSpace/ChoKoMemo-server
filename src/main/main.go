package main

import (
	"net/http"

	"github.com/ChoKoSpace/ChoKoMemo-server/src/api"
	"github.com/ChoKoSpace/ChoKoMemo-server/src/config"
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
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	api.InitDB()
	api.GetFromDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/status", healthCheck)
	mux.HandleFunc("/api/chokomemo/", apiRounter)
	http.ListenAndServe(config.Port, mux)
}
