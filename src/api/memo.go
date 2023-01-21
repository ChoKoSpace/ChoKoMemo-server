package api

import "net/http"

//type GetMemoRequest

func Memo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		http.NotFound(w, r)

	case http.MethodGet:
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodDelete:
	}
}
